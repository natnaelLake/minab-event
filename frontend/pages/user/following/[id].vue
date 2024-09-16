<script setup>
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useQuery, useMutation } from "@vue/apollo-composable";
import GetUserFollowings from "~/graphql/query/GetUserFollowings.gql";
import { useAuthStore } from "~/store";
import UNFOLLOWS_USER from "~/graphql/mutations/UnfollowUser.gql";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";

const router = useRouter();
const route = useRoute();
const followingUsers = ref([]);
const currentUser = route.params.id;
const user = useAuthStore();
const { mutate: unfollowUserMutate } = useMutationComposable(UNFOLLOWS_USER);
const goToUserProfile = (userId) => {
  router.push(`/user/profile/${userId}`);
};

const unfollowUser = async (userId) => {
  followingUsers.value = followingUsers.value.filter(
    (follow) => follow.user.id !== userId
  );
  await unfollowUserMutate({
    followerId: currentUser,
    userId,
  });
};

const { onResult: followingResult, refetch } = useQueryComposable(
  GetUserFollowings,
  {
    follower_id: currentUser,
  }
);

followingResult((result) => {
  if (result.data && result.data.follows) {
    followingUsers.value = result.data.follows;
  }
});
onMounted(async () => {
  await refetch();
});
</script>

<template>
  <div
    class="flex flex-col items-center w-full justify-center min-h-screen dark:bg-background-800 px-4 sm:px-6 lg:px-8 py-10 mt-20"
  >
    <div class="w-full max-w-3xl">
      <h3 class="text-2xl font-semibold mb-6 text-center">Following</h3>
      <ul class="space-y-4 mb-10">
        <li
          v-for="(follow, index) in followingUsers"
          :key="index"
          class="flex justify-between items-center bg-gray-200 dark:bg-white p-4 rounded-lg shadow-md"
        >
          <div
            class="flex items-center space-x-4 cursor-pointer"
            @click="goToUserProfile(follow.user.id)"
          >
            <img
              :src="
                follow.user.profile_image_url ||
                'https://via.placeholder.com/150'
              "
              alt="profile-pic"
              class="w-12 h-12 rounded-full"
            />
            <div class="flex flex-col">
              <span class="font-semibold hover:text-blue-400 hover:underline">
                {{ follow.user.first_name }} {{ follow.user.last_name }}
              </span>
              <span class="text-sm text-gray-500 dark:text-gray-400">
                {{ follow.user.bio }}
              </span>
            </div>
          </div>
          <button
            class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600"
            @click.stop="unfollowUser(follow.user.id)"
          >
            Unfollow
          </button>
        </li>
      </ul>
    </div>
  </div>
</template>
