<template>
  <div v-if="users" class="container mx-auto p-4">
    <div class="bg-white shadow rounded-lg p-28 mt-20">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <img
            class="w-20 h-20 rounded-full object-cover"
            :src="users.profile_image_url"
            alt="User Avatar"
          />
          <div class="ml-4">
            <h2 class="text-xl font-bold">
              {{ users.first_name }} {{ users.last_name }}
            </h2>
            <p class="text-gray-600">{{ user.title || "No user Title" }}</p>
          </div>
        </div>
        <div v-if="signedUserId === user.id">
          <button
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
          >
            Edit Profile
          </button>
        </div>
      </div>

      <div class="mt-6 flex justify-around">
        <button
          class="text-blue-500 font-semibold hover:underline"
          @click="redirectToFollowing"
        >
          {{ followingCounts }} Following
        </button>
        <button
          class="text-blue-500 font-semibold hover:underline"
          @click="redirectToFollowers"
        >
          {{ followerCounts }} Followers
        </button>
      </div>

      <div class="mt-6">
        <h3 class="text-lg font-semibold">About</h3>
        <p class="text-gray-700 mt-2">{{ user.bio }}</p>
      </div>
      <div class="mt-6">
        <h3 class="text-lg font-semibold">Personal Information</h3>
        <ul class="list-disc pl-5 text-gray-700 mt-2">
          <li>{{ users.email || "No email available" }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from "vue";
import getSingleUser from "~/graphql/query/getSingleUser.gql";
import { useQuery } from "@vue/apollo-composable";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import GetUserFollowings from "~/graphql/query/GetUserFollowings.gql";
import GetUserFollowers from "~/graphql/query/GetUserFollowers.gql";

const router = useRouter();
const route = useRoute();
const followingCounts = ref(0);
const followerCounts = ref(0);
const signedUserId = useAuthStore().user.id;
const user = useAuthStore();
const users = ref(null);
const currentUser = route.params.id;

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
} = useQuery(
  GetUserFollowings,
  {
    follower_id: currentUser,
  },
  {
    context: {
      headers: {
        "x-hasura-user-id": currentUser,
        "x-hasura-role": user.role,
        Authorization: `Bearer ${user.token}`,
      },
    },
  }
);
const {
  onResult: followersResult,
  onError: followersError,
  refetch: refetchFollowers,
} = useQuery(
  GetUserFollowers,
  {
    user_id: currentUser,
  },
  {
    context: {
      headers: {
        "x-hasura-user-id": currentUser,
        "x-hasura-role": user.role,
        Authorization: `Bearer ${user.token}`,
      },
    },
  }
);
followingResult((result) => {
  if (result.data) {
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
    const { onResult, onError, refetch } = useQuery(
      getSingleUser,
      {
        id: currentUser,
      },
      {
        context: {
          headers: {
            "x-hasura-user-id": currentUser,
            "x-hasura-role": user.role,
            Authorization: `Bearer ${user.token}`,
          },
        },
      }
    );

    onResult((result) => {
      if (result.data) {
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
const redirectToFollowing = () => {
  navigateTo("/user/following/" + currentUser);
};
const redirectToFollowers = () => {
  navigateTo("/user/followers/" + currentUser);
};
</script>

<style scoped>
.container {
  max-width: 800px;
}
</style>
