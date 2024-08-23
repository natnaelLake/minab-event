<template>
    <div class="flex flex-col items-center w-full justify-center min-h-screen px-4 sm:px-6 lg:px-8 py-10 mt-20">
      <div class="w-full max-w-4xl">
        <h3 class="text-3xl font-bold mb-8 text-center text-gray-900 dark:text-gray-900">Your Bookmarks</h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
          <div
            v-for="bookmark in bookmarks"
            :key="bookmark.id"
            class="bg-white dark:bg-gray-1'*00 p-5 rounded-lg shadow-lg transform transition-transform hover:-translate-y-2 hover:shadow-xl cursor-pointer"
            @click="goToCardDetail(bookmark.id)"
          >
            <img
              :src="bookmark.image"
              alt="bookmark-thumbnail"
              class="w-full h-32 sm:h-40 rounded-md object-cover mb-4"
            />
            <h4 class="text-lg font-semibold text-gray-800 dark:text-gray-100 mb-2">
              {{ bookmark.title }}
            </h4>
            <p class="text-sm text-gray-500 dark:text-gray-400 mb-4">
              {{ bookmark.description }}
            </p>
            <div class="flex justify-between items-center">
              <span class="text-xs text-gray-400 dark:text-gray-500">{{ bookmark.date }}</span>
              <button
                class="text-red-500 hover:text-red-600 transition-colors duration-200"
                @click.stop="removeBookmark(bookmark.id)"
              >
                Remove
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  <script setup>
  import { ref } from "vue";
  import { useRouter } from "vue-router";
  
  const router = useRouter();
  
  const bookmarks = ref([
    {
      id: 1,
      image: "https://via.placeholder.com/150",
      title: "Art Exhibition",
      description: "Explore the latest in contemporary art at the Art Expo.",
      date: "August 15, 2024",
    },
    {
      id: 2,
      image: "https://via.placeholder.com/150",
      title: "Music Festival",
      description: "Join us for a weekend of live music and entertainment.",
      date: "September 22, 2024",
    },
    {
      id: 1,
      image: "https://via.placeholder.com/150",
      title: "Art Exhibition",
      description: "Explore the latest in contemporary art at the Art Expo.",
      date: "August 15, 2024",
    },
    {
      id: 2,
      image: "https://via.placeholder.com/150",
      title: "Music Festival",
      description: "Join us for a weekend of live music and entertainment.",
      date: "September 22, 2024",
    },
    // Add more bookmarked items as needed
  ]);
  
  const goToCardDetail = (id) => {
    router.push(`/card-detail/${id}`);
  };
  
  const removeBookmark = (id) => {
    bookmarks.value = bookmarks.value.filter((bookmark) => bookmark.id !== id);
    // Optionally, add an API call to remove the bookmark from the backend
  };
  </script>
  