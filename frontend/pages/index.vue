<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from "vue";
import EventCard from "~/components/EventCard.vue";
import Filter from "~/components/Filter.vue";
import GetEvents from "~/graphql/query/GetEvents.gql";
import RESERVE_TICKET from "~/graphql/mutations/ReserveTicket.gql";
import GetReservedTickets from "~/graphql/query/GetReservedTickets.gql";
import UNRESERVE_TICKET from "~/graphql/mutations/UnReserveTicket.gql";
import BookMarkEvent from "~/graphql/mutations/BookMarkEvent.gql";
import LikeEvent from "~/graphql/mutations/LikeEvent.gql";
import UNBookMarkEvent from "~/graphql/mutations/UNBookMarkEvent.gql";
import UNLikeEvent from "~/graphql/mutations/UNLikeEvent.gql";
import GetEventLike from "~/graphql/query/GetEventLike.gql";
import GetEventBookMark from "~/graphql/query/GetEventBookMark.gql";
import { formatDistance, format, parseISO } from "date-fns";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import SkeletonLoader from "~/components/SkeletonLoader.vue";
import { useRouter, useRoute } from "vue-router";
import SaveTransaction from "~/graphql/mutations/SaveTransaction.gql";
import CreatePaymentMutation from "~/graphql/mutations/CreatePaymentMutation.gql";
import { VueSpinner } from "vue3-spinners";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";

const user = useAuthStore();
const currentUser = user.id;
const router = useRouter();
const route = useRoute();
const events = ref([]);
const ticketQuantity = ref(1);
const selectedEvent = ref(null);
const searchTerm = ref("");
const showCheckout = ref(false);
const isEventReserved = ref(false);
const tickets = ref([]);
const bookMarks = ref([]);
const likes = ref([]);
const filters = ref({});
const isFetching = ref(false);
const page = ref(0);
const hasMoreTokens = ref(true);
const itemsPerPage = 9;
const lastElementRef = ref(null);
const totalPrice = computed(
  () => ticketQuantity.value * (selectedEvent.value?.price || 0)
);
const availableTickets = computed(
  () =>
    selectedEvent.value.quantity -
    tickets.value.reduce((sum, ticket) => sum + ticket.quantity, 0)
);
const soldTickets = computed(() =>
  selectedEvent.value
    ? selectedEvent.value.quantity - availableTickets.value
    : 0
);
const { mutate: saveTransaction } = useMutationComposable(SaveTransaction);
const { mutate: makePayment, loading } = useMutationComposable(
  CreatePaymentMutation
);
const { mutate: unLikeEvent } = useMutationComposable(UNLikeEvent);
const { mutate: likeEvent } = useMutationComposable(LikeEvent);
const { mutate: bookMarkEvent } = useMutationComposable(BookMarkEvent);
const { mutate: unBookMarkEvent } = useMutationComposable(UNBookMarkEvent);
const resetPagination = () => {
  events.value = [];
  page.value = 0;
  hasMoreTokens.value = true;
  isFetching.value = false;
};

watch(searchTerm, async (newSearchTerm) => {
  resetPagination();
  await fetchEvents();
});

watch(
  () => route.fullPath,
  async () => {
    resetPagination();
    await fetchEvents();
  }
);

onMounted(async () => {
  resetPagination();
  await fetchEvents();
});

// Methods
onMounted(async () => {
  await fetchEvents();
  observeLastElement();
});

const fetchEvents = async () => {
  if (isFetching.value || !hasMoreTokens.value) return;

  isFetching.value = true;
  await new Promise((resolve) => setTimeout(resolve, 1000));

  try {
    const { onResult, onError } = useQueryComposable(GetEvents,
      {
        limit: itemsPerPage,
        offset: page.value * itemsPerPage,
        order_by: [{ created_at: "desc" }],
        where: {
          _and: [
            {
              _or: [
                {
                  title: searchTerm.value
                    ? { _ilike: `%${searchTerm.value}%` }
                    : {},
                },
                {
                  description: searchTerm.value
                    ? { _ilike: `%${searchTerm.value}%` }
                    : {},
                },
                {
                  tags: searchTerm.value
                    ? {
                        _in: `{${searchTerm.value
                          .split(" ")
                          .map((tag) => tag.trim())
                          .filter((tag) => tag !== "")}}`,
                      }
                    : {},
                },
              ],
            },
            ...Object.entries(filters.value).map(([key, value]) => ({
              [key]: value,
            })),
          ],
        },
      }      
    );

    onResult(({ data }) => {
      if (data.events && data.events.length < itemsPerPage) {
        hasMoreTokens.value = false;
      }
      events.value.push(...data.events);
      if (currentUser) {
        data.events.forEach((event) => {
          checkBookmark(event);
          checkLike(event);
        });
      }
      page.value++;
      isFetching.value = false;
    });

    onError((error) => {
      console.error("Error fetching events:", error);
      isFetching.value = false;
    });
  } catch (error) {
    console.error("Error fetching events:", error);
    isFetching.value = false;
  }
};

const observeLastElement = () => {
  const observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting && !isFetching.value) {
        fetchEvents();
      }
    },
    {
      rootMargin: "100px",
    }
  );
  if (lastElementRef.value) {
    observer.observe(lastElementRef.value);
  }
};

const checkBookmark = (event) => {
  const { onResult: bookmarkResult } = useQueryComposable(GetEventBookMark, {
    event_id: event.id,
    user_id: currentUser,
  });
  bookmarkResult((result) => {
    if (result.data) {
      const isBookmarked = result.data.bookmarks.some(
        (bookmark) =>
          bookmark.event_id === event.id && bookmark.user_id === currentUser
      );
      if (isBookmarked) {
        bookMarks.value.push(event.id);
      }
    }
  });
};

const checkLike = (event) => {
  const { onResult: likeResult } = useQueryComposable(GetEventLike, {
    event_id: event.id,
    user_id: currentUser,
  });
  likeResult((result) => {
    if (result.data) {
      const isLiked = result.data.likes.some(
        (like) => like.event_id === event.id && like.user_id === currentUser
      );
      if (isLiked) {
        likes.value.push(event.id);
      }
    }
  });
};
const handleEventLike = async (event) => {
  try {
    const liked = likes.value.includes(event.id);
    if (liked) {
      await unLikeEvent({ event_id: event.id, user_id: currentUser });
      likes.value = likes.value.filter((id) => id !== event.id);
    } else {
      await likeEvent({ event_id: event.id });
      likes.value.push(event.id);
    }
  } catch (error) {
    console.error("Error liking/unliking event:", error);
  }
};

const handleEventBookMark = async (event) => {
  try {
    const bookmarked = bookMarks.value.includes(event.id);
    if (bookmarked) {
      await unBookMarkEvent({ event_id: event.id, user_id: currentUser });
      bookMarks.value = bookMarks.value.filter((id) => id !== event.id);
    } else {
      await bookMarkEvent({ event_id: event.id });
      bookMarks.value.push(event.id);
    }
  } catch (error) {
    console.error("Error bookmarking/unbookmarking event:", error);
  }
};

const formatEventDate = (date) => format(parseISO(date), "MMMM do, yyyy");

const handleFormatDistance = (date) =>
  formatDistance(new Date(date), new Date(), {
    addSuffix: true,
  });

const showCheckoutModal = (event) => {
  const { onResult: reservedTicketsResult } = useQueryComposable(GetReservedTickets,
    {
      event_id: event.id,
    }
  );

  reservedTicketsResult((result) => {
    tickets.value = result.data.tickets;
    isEventReserved.value = result.data.tickets.some(
      (ticket) => ticket.event_id === event.id
    );
  });

  selectedEvent.value = event;
  showCheckout.value = true;
};

const closeCheckoutModal = () => {
  showCheckout.value = false;
  ticketQuantity.value = 1;
};

// Handle event ticket reservation
const handleReserveEvent = async () => {
  try {
    loading.value = true;
    const transactionData = {
      total_price: totalPrice.value,
      tx_ref: `txn_${Date.now()}`,
      event_id: selectedEvent.value.id,
      quantity: ticketQuantity.value,
    };
    const input = {
      total_price: totalPrice.value,
      tx_ref: `txn_${Date.now()}`,
      event_id: selectedEvent.value.id,
      user_id: currentUser,
      quantity: ticketQuantity.value,
      email: user.user.email,
      first_name: user.user.first_name,
      last_name: user.user.last_name,
      return_url: `${
        window.location.origin
      }/payment-status?tx_ref=txn_${Date.now()}`,
    };
    await saveTransaction({ ...transactionData });

    const result = await makePayment({ input });

    console.log("Payment response:", result);

    // Extract the checkout URL from the response
    const checkoutUrl = result.data.processPayment.data.checkout_url;

    // Redirect the user to the payment checkout page
    if (checkoutUrl) {
      window.location.href = checkoutUrl;
    } else {
      console.error("No checkout URL found in the response");
    }
  } catch (error) {
    console.error("Error updating reservation status:", error);
    toast.error("Failed to update reservation status.");
  }
};
const updateFilters = async (event) => {
  filters.value = event.detail;
  resetPagination();
  await fetchEvents();
};

const goToEventDetail = (eventId) => {
  if (currentUser) {
    router.push(`/event/${eventId}`);
  }
};
onMounted(() => {
  window.addEventListener("apply-filters", updateFilters);
});

onUnmounted(() => {
  window.removeEventListener("apply-filters", updateFilters);
});
</script>

<template>
  <div class="flex mt-20">
    <div
      class="w-64 bg-gray-100 p-2 pb-16 fixed top-20 left-0 h-screen overflow-scroll"
    >
      <Filter @apply-filters="updateFilters" />
    </div>
    <div class="flex-1 ml-64 p-6">
      <div class="mb-6">
        <input
          v-model="searchTerm"
          type="text"
          placeholder="Search events..."
          class="w-full border border-gray-300 rounded-lg px-4 py-2"
        />
      </div>

      <!-- Event Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="event in events"
          :key="event.id"
          class="overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer"
        >
          <EventCard
            :header="{
              avatar: event.user.profile_image_url || 'default-avatar.png',
              name: `${event.user.first_name} ${event.user.last_name}`,
              actions: [
                {
                  icon:
                    bookMarks && bookMarks.includes(event.id)
                      ? 'fa-solid fa-bookmark text-green-500'
                      : 'fa-regular fa-bookmark text-gray-500',
                  handler: () => handleEventBookMark(event),
                },
                {
                  icon:
                    likes && likes.includes(event.id)
                      ? 'fa-solid fa-heart text-red-500'
                      : 'fa-regular fa-heart text-gray-500',
                  handler: () => handleEventLike(event),
                },
              ],
            }"
            :image="{
              src: event.featured_image_url.split(',')[0].replace('{', ''), // Use the first image and remove '{'
              alt: event.title,
            }"
            :title="event.title"
            :description="event.description"
            :address="event.venue"
            :price="`$${event.price}`"
            :deadline="`Deadline ${formatEventDate(event.event_end_time)}`"
            :footer="{
              reserve: {
                handler: () => showCheckoutModal(event),
                class: 'bg-blue-500 text-white px-4 py-2 rounded',
              },
              postTime: {
                text: handleFormatDistance(event.created_at),
              },
            }"
            @navigate="goToEventDetail(event.id)"
            class="dark:bg-white text-black"
          />
        </div>

        <!-- Skeleton Loader -->
      </div>
      <div
        v-if="isFetching"
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
      >
        <div
          v-for="i in itemsPerPage"
          :key="`skeleton-${i}`"
          class="overflow-hidden"
        >
          <SkeletonLoader />
        </div>
      </div>

      <!-- Infinite Scroll -->
      <div ref="lastElementRef" class="h-10"></div>
    </div>
  </div>

  <!-- Checkout Modal -->
  <div
    v-if="showCheckout"
    class="fixed inset-0 z-50 flex items-center justify-center"
  >
    <div class="fixed inset-0 bg-black opacity-50"></div>
    <div
      class="bg-white w-full max-w-md p-6 rounded-lg shadow-2xl z-50 relative"
    >
      <button
        class="absolute top-2 right-2 text-gray-600"
        @click="closeCheckoutModal"
      >
        &times;
      </button>
      <h2 class="text-2xl font-bold text-gray-800 mb-4">Checkout</h2>

      <!-- Ticket Price Section -->
      <div class="flex items-center justify-between mb-4">
        <div>
          <span class="text-lg font-medium text-gray-700"
            >Price Per Ticket:</span
          >
          <span class="text-2xl font-bold text-gray-800"
            >${{ selectedEvent?.price }}</span
          >
        </div>
      </div>

      <!-- Quantity and Total Price Section -->
      <div class="flex items-center justify-between mb-4">
        <div class="flex flex-col">
          <label for="quantity" class="text-sm font-medium text-gray-600 mb-1"
            >Quantity</label
          >
          <input
            id="quantity"
            type="number"
            min="1"
            :max="availableTickets"
            v-model.number="ticketQuantity"
            class="border border-gray-300 rounded-lg px-3 py-2 text-gray-700 w-20"
          />
        </div>
        <div class="text-right">
          <span class="text-lg font-medium text-gray-600">Total Price:</span>
          <span class="text-xl font-bold text-gray-800">${{ totalPrice }}</span>
        </div>
      </div>

      <!-- Reserve Button Section -->
      <div class="mt-4">
        <button
          class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300 flex justify-center items-center"
          :disabled="loading"
          :class="{
            'bg-gray-500 text-white opacity-50 cursor-not-allowed': isEventReserved || loading,
            'bg-blue-500 text-white hover:bg-blue-600':
              !isEventReserved,
          }"
          @click="handleReserveEvent"
        >
          <VueSpinner v-if="loading" size="24" color="#ffffff" />

          <span v-else="loading">{{
            isEventReserved ? "Reserved" : "Buy Ticket"
          }}</span>
        </button>
      </div>

      <!-- Ticket Availability Info -->
      <div class="text-sm text-gray-600 mt-4">
        <p>Available: {{ availableTickets }}</p>
        <p>Sold: {{ soldTickets }}</p>
      </div>
    </div>
  </div>
</template>