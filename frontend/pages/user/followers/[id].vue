<template>
  <div
    class="flex flex-col items-center w-full justify-center min-h-screen dark:bg-background-800 px-4 sm:px-6 lg:px-8 py-10 mt-20"
  >
    <div class="w-full max-w-3xl">
      <h3 class="text-2xl font-semibold mb-6 text-center">Followers</h3>
      <ul class="space-y-4 mb-10">
        <li
          v-for="(follower, index) in followerUsers"
          :key="index"
          class="flex justify-between items-center bg-gray-200 dark:bg-white p-4 rounded-lg shadow-md"
        >
          <div
            class="flex items-center space-x-4 cursor-pointer"
            @click="goToUserProfile(follower.userByFollowerId.id)"
          >
            <img
              :src="follower.userByFollowerId.profile_image_url || 'https://via.placeholder.com/150'"
              alt="profile-pic"
              class="w-12 h-12 rounded-full"
            />
            <div class="flex flex-col">
              <span class="font-semibold hover:text-blue-400 hover:underline">
                {{ follower.userByFollowerId.first_name }} {{ follower.userByFollowerId.last_name }}
              </span>
              <span class="text-sm text-gray-500 dark:text-gray-400">
                {{ follower.userByFollowerId.bio }}
              </span>
            </div>
          </div>
          
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useQuery } from "@vue/apollo-composable";
import GetUserFollowers from "~/graphql/query/GetUserFollowers.gql";
import { useAuthStore } from "~/store";

const router = useRouter();
const route = useRoute();
const followerUsers = ref([]);
const currentUser = route.params.id;

const goToUserProfile = (userId) => {
  router.push(`/user/profile/${userId}`);
};

const { onResult: followersResult,refetch } = useQuery(GetUserFollowers, {
  user_id: currentUser,
});

followersResult((result) => {
  if (result.data && result.data.follows) {
    followerUsers.value = result.data.follows;
  }
});
onMounted(async () => {
  await refetch();
});
</script>

<style scoped>
/* Additional styles if needed */
</style>
