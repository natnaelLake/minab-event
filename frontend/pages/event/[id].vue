<template>
  <div v-if="eventData && !loading" class="flex mt-20">
    <!-- Product Image and Thumbnails Section -->
    <div class="flex-grow max-w-6xl ml-20 p-6 bg-white shadow-lg rounded-lg">
      <div class="flex">
        <!-- Image Carousel with Thumbnails -->
        <div class="w-1/4 pr-4">
          <div
            ref="thumbnailContainer"
            class="flex flex-col space-y-2 overflow-y-auto h-full max-h-[calc(100vh-9rem)] scrollbar-hidden"
          >
            <img
              v-for="(image, index) in images"
              :key="index"
              :src="image"
              alt="Thumbnail"
              class="w-20 h-20 object-cover rounded-lg cursor-pointer border-2 transition-all duration-300 ease-in-out hover:border-blue-400"
              :class="{ 'border-blue-500': currentIndex === index }"
              @click="changeImage(index)"
            />
          </div>
        </div>

        <!-- Large Product Image -->
        <div class="w-3/4">
          <img
            :src="currentImage"
            alt="Product Image"
            class="w-full h-[30rem] object-cover rounded-lg shadow-lg"
          />
        </div>
      </div>
    </div>

    <!-- Product Owner Card -->
    <div class="w-1/4 pl-6">
      <div class="flex flex-col items-start p-4 bg-white shadow-lg rounded-lg">
        <!-- Avatar -->
        <img
          :src="eventData.user.profile_image_url || avatarUrl"
          alt="Product Owner Avatar"
          class="w-24 h-24 rounded-full object-cover border-4 border-gray-300"
        />

        <!-- Product Owner Details -->
        <div class="mt-4">
          <h2 class="text-2xl font-bold text-gray-900">
            {{ eventData.user.first_name }} {{ eventData.user.last_name }}
          </h2>
          <p class="text-sm text-gray-600 mt-1">Product Owner</p>
        </div>

        <!-- Follow Button -->
        <button
          v-if="currentUser !== eventData.user.id"
          @click="followOwner()"
          :class="{
            'bg-red-600': isFollowing,
            'bg-blue-600': !isFollowing,
          }"
          class="text-white px-5 py-2 rounded-lg font-semibold hover:opacity-90 transition-opacity duration-300"
        >
          {{ isFollowing ? "Unfollow" : "Follow" }}
        </button>
      </div>
    </div>
  </div>

  <!-- Product Details Section -->
  <div
    v-if="eventData && !loading"
    class="flex max-w-6xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-8"
  >
    <div class="flex-grow mr-8">
      <!-- Product Title and Type/Category -->
      <div class="mb-6">
        <h1 class="text-4xl font-bold text-gray-800">{{ eventData.title }}</h1>
        <p class="p-2">Category</p>
        <div class="mt-2 flex space-x-2">
          <span
            class="bg-blue-100 text-blue-800 text-sm font-medium px-3 py-1 rounded-full"
          >
            {{ eventData.categories.join(", ") }}
          </span>
        </div>
      </div>

      <!-- Tags -->
      <p class="p-2">Event Tags</p>
      <div class="mb-6 flex flex-wrap gap-2">
        <span
          v-for="(tag, index) in eventData.tags"
          :key="index"
          class="bg-green-100 text-green-800 text-sm font-medium px-3 py-1 rounded-full"
        >
          {{ tag }}
        </span>
      </div>

      <!-- Description -->
      <div class="mb-6">
        <h2 class="text-xl font-semibold text-gray-700">Description</h2>
        <p class="text-gray-600 mt-2" :class="{ 'line-clamp-3': !isExpanded }">
          {{ eventData.description }}
        </p>
        <button
          @click="toggleDescription"
          class="text-blue-600 mt-2 inline-block"
        >
          {{ isExpanded ? "Show Less" : "Show More" }}
        </button>
      </div>

      <!-- Additional Details -->
      <div class="mb-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex flex-col">
            <span class="text-3xl font-bold text-gray-800">
              Price Per Ticket ${{ eventData.price }}
            </span>
            <span class="text-2xl pt-3 text-gray-800">
              Total Capacity {{ eventData.quantity }} Person
            </span>
          </div>
          <div class="text-gray-600">
            <p class="text-sm flex items-center">
              <svg
                class="w-5 h-5 mr-2 text-blue-500"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 8v4l2 2m-2-2H4M12 16v-4l-2-2m2 2h8"
                ></path>
              </svg>
              Date:
              {{
                handleEventDay(
                  eventData.event_date,
                  eventData.event_start_time,
                  eventData.event_end_time
                )
              }}
            </p>
            <p class="text-sm flex items-center">
              <svg
                class="w-5 h-5 mr-2 text-green-500"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 6v6h4m-4 0H4m8-4V4m0 0L4 8m8-4l4 4m4 8v4a2 2 0 01-2 2h-8a2 2 0 01-2-2v-4m4 0h8"
                ></path>
              </svg>
              Location:
              <a
                :href="`https://www.google.com/maps?q=${eventData.location}`"
                class="text-blue-600 hover:underline"
                target="_blank"
                rel="noopener noreferrer"
              >
                {{ eventData.venue }}
              </a>
            </p>
            <p class="text-sm flex items-center">
              <svg
                class="w-5 h-5 mr-2 text-red-500"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13 16h-1v-2h-1V8h2v6h1v2zM6 16h1v-2H6v-6h1V8H6V6h2v2h1v2H6v2zM8 16h4v-4H8v4zM18 16h-4v-4h4v4z"
                ></path>
              </svg>
              Preparation Time: {{ eventData.preparation_time }}
            </p>
          </div>
        </div>

        <div class="text-sm text-gray-500 mb-4">
          <p>Posted on: {{ formatDate(eventData.created_at) }}</p>
        </div>

        <div class="flex items-center space-x-4">
          <span class="text-gray-700">Likes: {{ eventData.likes.length }}</span>
          <a href="#" class="text-blue-600"
            >Reviews ({{ eventData.reviews.length }})</a
          >
        </div>
      </div>
    </div>

    <!-- Checkout Section -->
    <div
      class="fixed right-0 bottom-0 w-full max-w-md p-6 bg-white shadow-2xl rounded-t-lg border-t border-gray-200 z-40"
    >
      <h2 class="text-2xl font-bold text-gray-800 mb-4">Checkout</h2>

      <!-- Ticket Price Section -->
      <div class="flex items-center justify-between mb-4">
        <div>
          <span class="text-lg font-medium text-gray-700"
            >Price Per Ticket:</span
          >
          <span class="text-2xl font-bold text-gray-800"
            >${{ eventData.price }}</span
          >
        </div>
      </div>

      <!-- Quantity and Total Price Section -->
      <div class="flex items-center justify-between mb-4">
        <div v-if="currentUser !== eventData.user.id" class="flex flex-col">
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
      <div class="mt-4" v-if="currentUser !== eventData.user.id">
        <button
          :class="
            isEventReserved
              ? 'bg-gray-500 text-white'
              : 'bg-blue-500 text-white hover:bg-blue-600'
          "
          @click="handleReserveEvent"
          class="w-full px-6 py-3 rounded-full shadow-lg transition duration-200 ease-in-out text-lg font-semibold"
          :disabled="isEventReserved"
        >
          {{ isEventReserved ? "Reserved" : "Buy Ticket" }}
        </button>
      </div>

      <!-- Ticket Availability Info -->
      <div class="text-sm text-gray-600 mt-4">
        <p>Available: {{ availableTickets }}</p>
        <p>Sold: {{ soldTickets }}</p>
      </div>
    </div>
  </div>

  <!-- Organized by this User Section -->
  <div
    v-if="latestEvents.length && !loadingEvents"
    class="max-w-6xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-8"
  >
    <h2 class="text-2xl font-bold text-gray-800 mb-4">
      Organized by {{ eventData.user.first_name }}
    </h2>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
      <div
        v-for="event in latestEvents"
        :key="event.id"
        class="bg-gray-100 rounded-lg p-4 shadow-lg"
      >
        <img
          :src="
            event.featured_image_url.split(',')[0].replace('{', '') ||
            'https://via.placeholder.com/150'
          "
          alt="Event Image"
          class="w-full h-40 object-cover rounded-lg mb-4"
        />
        <h3 class="text-xl font-bold text-gray-800">{{ event.title }}</h3>
        <p class="text-gray-600 mt-2">{{ event.description }}</p>
        <p class="text-gray-500 mt-2">
          Date: {{ formatDate(event.event_date) }}
        </p>
        <p class="text-gray-800 font-bold mt-2">${{ event.price }}</p>
        <a
          :href="`/event/${event.id}`"
          class="text-blue-600 mt-4 block hover:underline"
        >
          View Event Details
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from "vue";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { useRoute } from "vue-router";
import GET_EVENT_DETAILS from "~/graphql/query/getEventDetails.gql";
import GET_LATEST_USER_EVENTS from "~/graphql/query/getLatestUserEvents.gql";
import { format, formatDistance, parse } from "date-fns";
import { toast } from "vue3-toastify";
import FOLLOWS_USER from "~/graphql/mutations/FollowUser.gql";
import UNFOLLOWS_USER from "~/graphql/mutations/UnfollowUser.gql";
import EventOwnerFollowers from "~/graphql/query/getEventOwnerFollowers.gql";
import GetReservedTickets from "~/graphql/query/GetReservedTickets.gql";
import RESERVE_TICKET from "~/graphql/mutations/ReserveTicket.gql";
import UNRESERVE_TICKET from "~/graphql/mutations/UnReserveTicket.gql";
import { useAuthStore } from "~/store";
const user = useAuthStore();
const route = useRoute();
const eventId = route.params.id;
const eventData = ref(null);
const userId = ref("");
const currentUser = user.id;
const isFollowing = ref(false);
const isEventReserved = ref(false);
const images = ref([]);
const currentImage = ref("");
const currentIndex = ref(0);
const isExpanded = ref(false);
const latestEvents = ref([]);
const ticketQuantity = ref(1); // Default quantity
const tickets = ref([]);
// const soldTickets = ref(0);

// GraphQL Mutations
const { mutate: reserveTicket } = useMutation(RESERVE_TICKET);
const { mutate: unReserveTicket } = useMutation(UNRESERVE_TICKET);
const { mutate: followUser } = useMutation(FOLLOWS_USER);
const { mutate: unfollowUser } = useMutation(UNFOLLOWS_USER);

// Fetch event details
const { onResult, loading, onError } = useQuery(GET_EVENT_DETAILS, {
  id: eventId,
});

onResult((result) => {
  eventData.value = result.data.events_by_pk;
  userId.value = eventData.value.user.id; // Set userId after fetching event details
  images.value = eventData.value.featured_image_url
    .replace(/{|}/g, "")
    .split(",");
  currentImage.value = images.value[0] || "";
});

onError((error) => {
  console.error(error);
  toast.error("Failed to load event details.");
});

// Watch userId and fetch latest events when it changes
watch(userId, (newUserId) => {
  if (newUserId) {
    const {
      onResult: onEventsResult,
      loading: loadingEvents,
      onError: onEventsError,
    } = useQuery(GET_LATEST_USER_EVENTS, { userId: newUserId });

    onEventsResult((result) => {
      latestEvents.value = result.data.events.filter(
        (event) => event.id !== eventId
      );
    });

    onEventsError((error) => {
      console.error(error);
      toast.error("Failed to load latest events.");
    });
  }
});

// Fetch followers and check if the current user is following
const {
  onResult: getFollowersResult,
  loading: getFollowersLoading,
  onError: getFollowersError,
} = useQuery(EventOwnerFollowers, {
  followerId: currentUser,
});

getFollowersResult((result) => {
  isFollowing.value = result.data.follows.some(
    (follow) => follow.follower_id === currentUser
  );
});

// Fetch reserved tickets to check if the event is reserved by the user
const { onResult: reservedTicketsResult } = useQuery(GetReservedTickets, {
  event_id: eventId,
});

reservedTicketsResult((result) => {
  tickets.value = result.data.tickets;
  console.log("-----------------++++++++++", tickets.value);
  isEventReserved.value = result.data.tickets.some(
    (ticket) => ticket.event_id === eventId
  );
});

// Change image in carousel
const changeImage = (index) => {
  currentImage.value = images.value[index];
  currentIndex.value = index;
};

// Toggle description visibility
const toggleDescription = () => {
  isExpanded.value = !isExpanded.value;
};

// Format date
const handleEventDay = (date, eventTime, eventEndTime) => {
  const startTimeDate = parse(eventTime, "HH:mm", new Date());
  const endTimeDate = parse(eventEndTime, "HH:mm", new Date());
  const formattedStartTime = format(startTimeDate, "hh:mm a");
  const formattedEndTime = format(endTimeDate, "hh:mm a");
  const dateTime =
    format(new Date(date), "MMM do yyyy") +
    " from " +
    formattedStartTime +
    " - " +
    formattedEndTime;
  return dateTime;
};
const formatDate = (dateString) => {
  return formatDistance(new Date(dateString), new Date(), {
    addSuffix: true,
  });
};

// Follow or unfollow event owner
const followOwner = async () => {
  try {
    if (isFollowing.value) {
      // Unfollow if currently following
      await unfollowUser({
        followerId: currentUser,
        userId: eventData.value.user.id,
      });
      isFollowing.value = false;
      toast.success("Unfollowed successfully.");
    } else {
      // Follow if not currently following
      await followUser({
        followerId: currentUser,
        userId: eventData.value.user.id,
      });
      isFollowing.value = true;
      toast.success("Followed successfully.");
    }
  } catch (error) {
    console.error("Error updating follow status:", error);
    toast.error("Failed to update follow status.");
  }
};
const totalPrice = computed(() => ticketQuantity.value * eventData.value.price);
const availableTickets = computed(
  () =>
    eventData.value.quantity -
    tickets.value.reduce((sum, ticket) => sum + ticket.quantity, 0)
);
const soldTickets = computed(
  () => eventData.value.quantity - availableTickets.value
);
// Handle event ticket reservation
const handleReserveEvent = async () => {
  try {
    if (isEventReserved.value) {
      await unReserveTicket({
        event_id: eventId,
        user_id: currentUser,
      });
      isEventReserved.value = false;
      toast.success("Event unreserved successfully.");
    } else {
      await reserveTicket({
        event_id: eventId,
        user_id: currentUser,
        quantity: ticketQuantity.value,
        total_price: totalPrice.value,
      });
      isEventReserved.value = true;
      toast.success("Event reserved successfully.");
    }
  } catch (error) {
    console.error("Error updating reservation status:", error);
    toast.error("Failed to update reservation status.");
  }
};
</script>

<style scoped>
.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
</style>
