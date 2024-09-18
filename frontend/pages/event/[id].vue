<script setup>
import { ref, watch,onMounted,onUnmounted,computed } from "vue";
import { useRoute } from "vue-router";
import GET_EVENT_DETAILS from "~/graphql/query/getEventDetails.gql";
import GET_LATEST_USER_EVENTS from "~/graphql/query/getLatestUserEvents.gql";
import { format, formatDistance, parseISO } from "date-fns";
import { toast } from "vue3-toastify";
import FOLLOWS_USER from "~/graphql/mutations/FollowUser.gql";
import UNFOLLOWS_USER from "~/graphql/mutations/UnfollowUser.gql";
import GET_REVIEWS from "~/graphql/query/GetReviews.gql";
import ADD_REVIEW from "~/graphql/mutations/AddReview.gql";
import EventOwnerFollowers from "~/graphql/query/getEventOwnerFollowers.gql";
import GetReservedTickets from "~/graphql/query/GetReservedTickets.gql";
import SaveTransaction from "~/graphql/mutations/SaveTransaction.gql";
import CreatePaymentMutation from "~/graphql/mutations/CreatePaymentMutation.gql";
import { useAuthStore } from "~/store";
import { VueSpinner } from "vue3-spinners";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";

const user = useAuthStore();
const currentUser = user.id;
const route = useRoute();
const eventId = route.params.id;
const eventData = ref(null);
const userId = ref("");
const isFollowing = ref(false);
const isEventReserved = ref(false);
const images = ref([]);
const eventTags = ref([]);
const currentImage = ref("");
const currentIndex = ref(0);
const isExpanded = ref(false);
const latestEvents = ref([]);
const ticketQuantity = ref(1);
const tickets = ref([]);
const reviews = ref([]);
const averageRating = ref(0);
const userReview = ref("");
const userRating = ref(0); // Rating can be between 1 to 5

const { mutate: saveTransaction } = useMutationComposable(SaveTransaction);
const { mutate: makePayment, loading: PaymentLoading } = useMutationComposable(
  CreatePaymentMutation
);
const { mutate: followUser } = useMutationComposable(FOLLOWS_USER);
const { mutate: unfollowUser } = useMutationComposable(UNFOLLOWS_USER);

// Fetch event details
const { onResult, loading, onError } = useQueryComposable(GET_EVENT_DETAILS, {
  id: eventId,
});

onResult((result) => {
  eventData.value = result.data.events_by_pk;
  userId.value = eventData.value.user.id; // Set userId after fetching event details
  images.value = eventData.value.featured_image_url
    .replace(/{|}/g, "")
    .split(",");
  eventTags.value = eventData.value.tags.replace(/{|}/g, "").split(",");
  currentImage.value = images.value[0] || "";
});

onError((error) => {
  console.error(error);
  toast.error("Failed to load event details.");
});

const { onResult: getReviewsResult } = useQueryComposable(GET_REVIEWS, {
  event_id: eventId,
});
const calculateAverageRating = computed(() => {
  if (reviews.value.length === 0) {
    return userRating.value; // If no reviews, return the new rating directly
  }
  
  // Sum of all current ratings
  const currentRatingsSum = reviews.value.reduce((sum, review) => sum + review.rating, 0);

  // Total number of ratings including the new one
  const totalRatings = reviews.value.length + 1;

  // Calculate the new average including the new user rating
  const newAverage = (currentRatingsSum + userRating.value) / totalRatings;

  return newAverage.toFixed(1); // Keep one decimal for average rating
});
getReviewsResult((result) => {
  reviews.value = result.data.reviews;
});

const {mutate} = useMutationComposable(ADD_REVIEW);

const submitReview = async () => {
  try {
    await mutate({
      event_id: eventId,
      comment: userReview.value,
      rating: userRating.value,
    });

    toast.success("Review submitted successfully.");
    // Clear the review input
    userReview.value = "";
    userRating.value = 0;

    // Refresh reviews after submitting
    const { onResult: refreshedReviewsResult } = useQueryComposable(
      GET_REVIEWS,
      { event_id: eventId }
    );
    refreshedReviewsResult((result) => {
      reviews.value = result.data.reviews;
    });
  } catch (error) {
    console.error("Error submitting review:", error);
    toast.error("Failed to submit review.");
  }
};
// Watch userId and fetch latest events when it changes
watch(userId, (newUserId) => {
  if (newUserId) {
    const {
      onResult: onEventsResult,
      loading: loadingEvents,
      onError: onEventsError,
    } = useQueryComposable(GET_LATEST_USER_EVENTS, { userId: newUserId });

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
} = useQueryComposable(EventOwnerFollowers, {
  followerId: currentUser,
});

getFollowersResult((result) => {
  isFollowing.value = result.data.follows.some(
    (follow) => follow.follower_id === currentUser
  );
});

// Fetch reserved tickets to check if the event is reserved by the user
const { onResult: reservedTicketsResult } = useQueryComposable(
  GetReservedTickets,
  {
    event_id: eventId,
  }
);

reservedTicketsResult((result) => {
  tickets.value = result.data.tickets;
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
const handleEventDay = (eventStartTime, eventEndTime) => {
  if (!eventStartTime || !eventEndTime) {
    console.error("Invalid date value:", eventStartTime, eventEndTime);
    return "Invalid Date"; // or you can return an empty string or a default date
  }

  try {
    const parsedStartDate = parseISO(eventStartTime);

    const parsedEndDate = parseISO(eventEndTime);
    const formattedStartTime = format(parsedStartDate, "PPpp");
    const formattedEndTime = format(parsedEndDate, "PPpp");
    return " from " + formattedStartTime + " - " + formattedEndTime;
  } catch (error) {
    console.error("Error parsing date:", error);
    return "Invalid Date"; // Handle the error case gracefully
  }
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
    PaymentLoading.value = true;
    const transactionData = {
      total_price: totalPrice.value,
      tx_ref: `txn_${Date.now()}`,
      event_id: eventId,
      quantity: ticketQuantity.value,
    };
    const input = {
      total_price: totalPrice.value,
      tx_ref: `txn_${Date.now()}`,
      event_id: eventId,
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

    // Log the response to ensure everything works as expected
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
const updateRating = (rating) => {
  console.log(rating);
  userRating.value = rating.detail.rating;
};

onMounted(() => {
  window.addEventListener("rating-selected", updateRating);
});
onUnmounted(() => {
  window.removeEventListener("rating-selected", updateRating);
});
</script>
<template>
  <div v-if="eventData && !loading" class="flex mt-20">
    <!-- Event Image and Thumbnails Section -->
    <div class="flex-grow max-w-6xl ml-20 p-6 bg-white shadow-lg rounded-lg">
      <div class="flex">
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
        <div class="w-3/4">
          <img
            :src="currentImage"
            alt="Product Image"
            class="w-full h-[30rem] object-cover rounded-lg shadow-lg"
          />
        </div>
      </div>
    </div>

    <div class="w-1/4 pl-6">
      <div class="flex flex-col items-start p-4 bg-white shadow-lg rounded-lg">
        <img
          :src="eventData.user.profile_image_url || avatarUrl"
          alt="Product Owner Avatar"
          class="w-24 h-24 rounded-full object-cover border-4 border-gray-300"
        />

        <div class="mt-4">
          <h2 class="text-2xl font-bold text-gray-900">
            {{ eventData.user.first_name }} {{ eventData.user.last_name }}
          </h2>
          <p class="text-sm text-gray-600 mt-1">Product Owner</p>
        </div>

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

  <div
    v-if="eventData && !loading"
    class="flex max-w-6xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-8"
  >
    <div class="flex-grow mr-8">
      <div class="mb-6">
        <h1 class="text-4xl font-bold text-gray-800">{{ eventData.title }}</h1>
        <p class="p-2">Category</p>
        <div class="mt-2 flex space-x-2">
          <span
            class="bg-blue-100 text-blue-800 text-sm font-medium px-3 py-1 rounded-full"
          >
            {{ eventData.categories }}
          </span>
        </div>
      </div>

      <p class="p-2">Event Tags</p>
      <div class="mb-6 flex flex-wrap gap-2">
        <span
          v-for="(tag, index) in eventTags"
          :key="index"
          class="bg-green-100 text-green-800 text-sm font-medium px-3 py-1 rounded-full"
        >
          {{ tag }}
        </span>
      </div>

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
              Event Date:
              {{
                handleEventDay(
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
            >Reviews ({{ calculateAverageRating }})</a
          >
        </div>
      </div>
    </div>
    <!-- Checkout Section -->
    <div
      v-if="!isEventReserved"
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
          :disabled="PaymentLoading"
          :class="{
            'bg-gray-500 text-white opacity-50 cursor-not-allowed':
              isEventReserved || PaymentLoading,
            'bg-blue-500 text-white hover:bg-blue-600': !isEventReserved,
          }"
          @click="handleReserveEvent"
        >
          <VueSpinner v-if="PaymentLoading" size="24" color="#ffffff" />

          <span v-else="PaymentLoading">{{
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
  <div
    class="flex flex-col max-w-6xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-8"
  >
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">
      Ratings and Reviews
    </h2>

    <!-- Average Rating -->
    <div class="flex items-center mb-6">
      <span class="text-4xl font-bold text-yellow-500">{{
        calculateAverageRating
      }}</span>
      <div class="ml-4">
        <div class="flex items-center">
          <RatingStars
            :rating="calculateAverageRating"
          />
          <span class="ml-2 text-gray-600">({{ reviews.length }} Reviews)</span>
        </div>
      </div>
    </div>

    <!-- Reviews List -->
    <div v-if="reviews.length > 0" class="space-y-4 mb-6">
      <div
        v-for="(review, index) in reviews"
        :key="index"
        class="flex items-start space-x-4"
      >
        <img
          :src="review.user.profile_image_url"
          alt="User Avatar"
          class="w-12 h-12 rounded-full object-cover"
        />
        <div>
          <h3 class="text-lg font-semibold">
            {{ review.user.first_name }} {{ review.user.last_name }}
          </h3>
          <div class="flex items-center">
            <RatingStars :rating="review.rating" />
            <span class="ml-2 text-sm text-gray-500">{{
              formatDate(review.created_at)
            }}</span>
          </div>
          <p class="mt-2 text-gray-700">{{ review.comment }}</p>
        </div>
      </div>
    </div>

    <!-- Add Review Form -->
    <div class="border-t pt-6">
      <h3 class="text-xl font-semibold text-gray-800 mb-4">Leave a Review</h3>
      <form @submit.prevent="submitReview">
        <!-- Rating Input -->
        <div class="mb-4">
          <label class="block text-gray-700">Rating:</label>
          <RatingStars :rating="userRating" @rating-selected="updateRating" />
        </div>

        <!-- Comment Input -->
        <div class="mb-4">
          <label class="block text-gray-700">Comment:</label>
          <textarea
            v-model="userReview"
            class="w-full px-4 py-2 border rounded-lg"
            rows="4"
            placeholder="Write your review..."
          ></textarea>
        </div>

        <button
          type="submit"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
        >
          Submit Review
        </button>
      </form>
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
          Date: {{ formatDate(event.event_start_time) }}
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
<style scoped>
.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
</style>
