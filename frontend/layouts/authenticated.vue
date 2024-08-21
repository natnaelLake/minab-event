<template>
  <div>
    <header class="absolute top-0 left-0 w-full z-10">
      <nav
        class="container mx-auto p-4 flex justify-between items-center bg-transparent"
      >
        <!-- Left side with Home icon and link -->
        <ul class="flex gap-4 items-center">
          <li class="flex items-center gap-2">
            <NuxtLink to="/" :class="linkClass('/')">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="2"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M3 9l9-7 9 7v11a2 2 0 01-2 2h-4a2 2 0 01-2-2V10H9v9a2 2 0 01-2 2H3a2 2 0 01-2-2V9z"
                />
              </svg>
              <span>Home</span>
            </NuxtLink>
          </li>
        </ul>

        <!-- Right side with Profile Button -->
        <ul class="flex gap-4 items-center">
          <li class="relative">
            <button
              @click="toggleDropdown"
              class="w-10 h-10 rounded-full bg-gray-300 flex items-center justify-center focus:outline-none focus:ring-4 focus:ring-blue-300"
            >
              <img src="" alt="Profile" class="w-full h-full rounded-full" />
            </button>
            <div
              v-if="dropdownOpen"
              class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg"
            >
              <ul class="py-1">
                <li>
                  <NuxtLink
                    to="/profile"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100"
                    >Profile</NuxtLink
                  >
                </li>
                <li>
                  <button
                    @click="openLogoutModal"
                    class="w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100"
                  >
                    Logout
                  </button>
                </li>
                <!-- Add more options as needed -->
              </ul>
            </div>
          </li>
        </ul>
      </nav>
    </header>
    <div class="container mx-auto p-4 pt-20">
      <slot />
    </div>

    <!-- Logout Confirmation Modal -->
    <div
      v-if="isLogoutModalOpen"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
    >
      <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full text-center">
        <h2 class="text-xl font-semibold mb-4">Confirm Logout</h2>
        <p class="mb-6">Are you sure you want to log out?</p>
        <div class="flex justify-center gap-4">
          <button
            @click="confirmLogout"
            class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
          >
            Yes
          </button>
          <button
            @click="closeLogoutModal"
            class="bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const visitedLinks = ref([]);
const route = useRoute();
const dropdownOpen = ref(false);
const isLogoutModalOpen = ref(false);

const isActive = (path) => route.path === path;

const linkClass = (path) => {
  return {
    "text-blue-500 font-bold": isActive(path),
    "text-gray-300": visitedLinks.value.includes(path),
  };
};

const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

const openLogoutModal = () => {
  isLogoutModalOpen.value = true;
};

const closeLogoutModal = () => {
  isLogoutModalOpen.value = false;
};

const confirmLogout = () => {
  // Add your logout logic here
  console.log("Logging out...");
  closeLogoutModal();
  // You might also want to navigate to a login page after logout
  // navigateTo('/login')
};
</script>

<style scoped>
/* Add any additional styling if necessary */
</style>
