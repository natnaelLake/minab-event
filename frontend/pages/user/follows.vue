<template>
    <div class="flex flex-col items-center w-full justify-center min-h-screen dark:bg-background-800 px-4 sm:px-6 lg:px-8 py-10 mt-20">
      <div class="w-full max-w-3xl">
        <h3 class="text-2xl font-semibold mb-6 text-center">Following</h3>
        <ul class="space-y-4 mb-10">
          <li
            v-for="(user, index) in followedUsers"
            :key="index"
            class="flex justify-between items-center bg-gray-200 dark:bg-white p-4 rounded-lg shadow-md"
          >
            <div class="flex items-center space-x-4 cursor-pointer" @click="goToUserProfile(user.id)">
              <img
                :src="user.avatar"
                alt="profile-pic"
                class="w-12 h-12 rounded-full"
              />
              <div class="flex flex-col">
                <span class="font-semibold hover:text-blue-400 hover:underline">
                  {{ user.name }}
                </span>
                <span class="text-sm text-gray-500 dark:text-gray-400">
                  {{ user.description }}
                </span>
              </div>
            </div>
            <button
              class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600"
              @click.stop="unfollowUser(user.id)"
            >
              Unfollow
            </button>
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from "vue";
  import { useRouter } from "vue-router";
  
  const router = useRouter();
  
  const followedUsers = ref([
    {
      id: 1,
      avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPfO37MK81JIyR1ptwqr_vYO3w4VR-iC2wqQ&s",
      name: "John Doe",
      description: "Software Developer at Tech Inc.",
    },
    {
      id: 2,
      avatar: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPfO37MK81JIyR1ptwqr_vYO3w4VR-iC2wqQ&s",
      name: "Jane Smith",
      description: "Designer at Creative Studio.",
    },
    // Add more users as needed
  ]);
  
  const goToUserProfile = (userId) => {
    router.push(`/user/profile/${userId}`);
  };
  
  const unfollowUser = (userId) => {
    followedUsers.value = followedUsers.value.filter(
      (user) => user.id !== userId
    );
    // API call can be made here to unfollow the user on the backend
  };
  </script>
  
  <style scoped>
  /* Additional styles if needed */
  </style>
  
