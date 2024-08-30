<template>
  <div class="container mx-auto p-4">
    <div class="bg-white shadow rounded-lg p-8 mt-20">
      <!-- Loading State -->
      <!-- <div v-if="loading">
        <p>Loading user data...</p>
      </div> -->

      <!-- User Data -->
      <div>
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <img
              class="w-20 h-20 rounded-full object-cover"
              :src="
                users.profile_image_url ||
                'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTafDGOCQJda4LjZXx9U4ZGlFdjYyVmZNRrwA&s'
              "
              alt="User Avatar"
            />
            <div class="ml-4">
              <h2 class="text-xl font-bold">
                {{ users.first_name || "Unknown" }}
                {{ users.last_name || "Unknown" }}
              </h2>
              <p class="text-gray-600">
                {{ "No title available" }}
              </p>
            </div>
          </div>
          <div>
            <button
              class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
              @click="redirectToEditProfile"
            >
              Edit Profile
            </button>
          </div>
        </div>

        <div class="mt-6 flex justify-around">
          <button
            class="text-blue-500 font-semibold hover:underline"
            @click="redirectFollowing"
          >
            {{ followingCounts || 0 }} Following
          </button>
          <button
            class="text-blue-500 font-semibold hover:underline"
            @click="redirectFollowers"
          >
            {{ followerCounts || 0 }} Followers
          </button>
        </div>

        <div class="mt-6">
          <h3 class="text-lg font-semibold">About</h3>
          <p class="text-gray-700 mt-2">
            {{ users.bio || "No bio available" }}
          </p>
        </div>

        <div class="mt-6">
          <h3 class="text-lg font-semibold">Personal Information</h3>
          <ul class="list-disc pl-5 text-gray-700 mt-2">
            <li>{{ users.email || "No email available" }}</li>
          </ul>
        </div>

        <!-- <div class="mt-6">
          <h3 class="text-lg font-semibold">Events Created</h3>
          <ul class="mt-4">
            <li
              v-for="event in userEvents.slice(0, 5)"
              :key="event.id"
              class="mb-4 p-4 bg-gray-100 rounded-lg shadow"
            >
              <h4 class="text-lg font-bold">
                {{ event.title || "Untitled Event" }}
              </h4>
              <p class="text-gray-600">
                {{ event.description || "No description available" }}
              </p>
              <p class="text-sm text-gray-500">
                Date: {{ event.event_start_time || "No date available" }}
              </p>
            </li>
          </ul>
        </div> -->
      </div>

      <!-- Empty or Error State -->
      <!-- <div v-else>
        <p>No user data available.</p>
      </div> -->
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from "vue";
import getSingleUser from "~/graphql/query/getSingleUser.gql";
import { useQuery } from "@vue/apollo-composable";
import { useRouter } from "vue-router";
import { formatDistance, format, parseISO } from "date-fns";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import GetUserFollowings from "~/graphql/query/GetUserFollowings.gql";
import GetUserFollowers from "~/graphql/query/GetUserFollowers.gql";

const router = useRouter();
const route = useRoute();
const user = useAuthStore();
const followingCounts = ref(0);
const followerCounts = ref(0);
const users = ref([]);
const currentUser = user.id;

watch(users, async (newUser) => {
  await fetchEvents();
});
// Methods
onMounted(async () => {
  await fetchEvents();
});
const {
  onResult: followingResult,
  onError: followingError,
  refetch: refetchFollowing,
} = useQuery(GetUserFollowings, {
  follower_id: currentUser,
});
const {
  onResult: followersResult,
  onError: followersError,
  refetch: refetchFollowers,
} = useQuery(GetUserFollowers, {
  user_id: currentUser,
});
followingResult((result) => {
  console.log("--------------===========", result.data);
  if (result.data) {
    console.log("--------------+++++++++++", result.data);

    followingCounts.value = result.data.follows.length;
  }
});
followersResult((result) => {
  if (result.data) {
    followerCounts.value = result.data.follows.length;
  }
});
const fetchEvents = async () => {
  try {
    const { onResult, onError, refetch } = useQuery(getSingleUser, {
      id: currentUser,
    });

    onResult((result) => {
      if (result.data) {
        const newEvents = result.data.events;
        users.value = result.data.users_by_pk;
      }
    });

    onError((error) => {
      console.error("Error fetching events: ", error.message);
      toast.error("Something went wrong, try again", {
        transition: toast.TRANSITIONS.FLIP,
        position: toast.POSITION.TOP_RIGHT,
      });
    });
  } catch (error) {
    console.error("Error during fetching events: ", error);
    toast.error("Failed to load events.");
  }
};

const redirectFollowing = () => {
  navigateTo("/user/following/" + currentUser);
};
const redirectFollowers = () => {
  navigateTo("/user/follows/" + currentUser);
};
const redirectToEditProfile = () => {
  navigateTo("/user/profile/edit-profile/" + currentUser);
};
</script>

<style scoped>
.container {
  max-width: 800px;
}
</style>
