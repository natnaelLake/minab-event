<template>
  <nav
    class="fixed w-full top-0 bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 shadow-md p-4 flex items-center justify-between text-white z-50"
  >
    <!-- Left: MinabEvent and Logo -->
    <div class="flex items-center space-x-4">
      <img
        src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTSKbCFe_QYSVH-4FpaszXvakr2Eti9eAJpQ&s"
        alt="Profile Avatar"
        class="h-10 w-10 rounded-full border border-gray-300"
      />
      <div v-if="!isAdmin">
        <NuxtLink to="/" class="text-xl font-bold hover:text-pink-500">
          MinabEvent
        </NuxtLink>
      </div>
      <div v-else>
        <p class="text-xl font-bold">MinabEvent</p>
      </div>
    </div>

    <!-- Right: Auth Links or Profile Actions -->
    <div v-if="isAuthenticated" class="flex items-center space-x-6">
      <div v-if="!isAdmin">
        <NuxtLink
          to="/event/create"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-pink-700 transition duration-300"
        >
          Create Event
        </NuxtLink>
      </div>
      <img
        :src="
          user && user.profile_image_url ||
          'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTSKbCFe_QYSVH-4FpaszXvakr2Eti9eAJpQ&s'
        "
        alt="Profile Avatar"
        @click="toggleDrawer"
        class="h-10 w-10 rounded-full cursor-pointer border border-gray-300"
      />
    </div>
    <div v-else class="flex items-center space-x-4">
      <NuxtLink
        to="/login"
        class="bg-white text-blue-600 px-4 py-2 rounded-lg hover:bg-gray-100 transition duration-300"
      >
        Login
      </NuxtLink>
      <NuxtLink
        to="/signup"
        class="bg-white text-blue-600 px-4 py-2 rounded-lg hover:bg-gray-100 transition duration-300"
      >
        Signup
      </NuxtLink>
    </div>

    <!-- Drawer for Profile Actions -->
    <transition name="slide-right">
      <div
        v-if="drawerOpen"
        @click="closeDrawer"
        class="fixed inset-0 flex justify-end z-50 mt-16"
      >
        <div
          class="w-64 h-full bg-white shadow-lg p-6 text-white transform transition-transform duration-300 ease-in-out"
          @click.stop
        >
          <ul class="space-y-2">
            <li v-for="route in routes" :key="route.path">
              <template v-if="route.name !== 'Logout'">
                <NuxtLink
                  :to="route.path"
                  class="block w-full text-left p-2 text-black rounded-lg hover:bg-blue-400 hover:text-white hover:shadow-lg"
                >
                  {{ route.name }}
                </NuxtLink>
              </template>
              <template v-if="route.name === 'Logout'">
                <button
                  @click="openModal"
                  class="block w-full text-left p-2 text-black rounded-lg hover:bg-blue-400 hover:text-white hover:shadow-lg"
                >
                  {{ route.name }}
                </button>
              </template>
            </li>
          </ul>
        </div>
      </div>
    </transition>

    <!-- Logout Confirmation Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
    >
      <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full">
        <h2 class="text-lg font-semibold mb-4">Confirm Logout</h2>
        <p class="-mt-4 mb-12 text-black">Are you sure you want to log out?</p>
        <div class="flex justify-end space-x-4">
          <button
            @click="handleLogout"
            class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition duration-300"
          >
            Logout
          </button>
          <button
            @click="cancelLogout"
            class="bg-gray-300 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-400 transition duration-300"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, computed } from "vue";
import { useAuthStore } from "@/store";
import { useRouter } from "vue-router";
import { getNavigationRoutes } from "@/routes";

const authStore = useAuthStore();
const router = useRouter();
const user = authStore.user;
const showModal = ref(false);
const drawerOpen = ref(false);
const isAuthenticated = computed(() => authStore.isAuthenticated);
const isAdmin = computed(() => authStore.role === "user-admin"); // Check if the role is admin

const openModal = () => {
  showModal.value = true;
};

const cancelLogout = () => {
  showModal.value = false;
};

const handleLogout = () => {
  authStore.logout();
  showModal.value = false;
  drawerOpen.value = false;
  router.push("/");
};

const toggleDrawer = () => {
  drawerOpen.value = !drawerOpen.value;
};

const closeDrawer = () => {
  drawerOpen.value = false;
};

const routes = computed(() => {
  return getNavigationRoutes(authStore.role, isAuthenticated.value);
});
</script>

<style scoped>
/* Drawer Slide Animation */
.slide-right-enter-active,
.slide-right-leave-active {
  transition: transform 0.3s ease;
}
.slide-right-enter,
.slide-right-leave-to {
  transform: translateX(100%);
}

/* Enhanced navbar and drawer styles */
nav ul li a {
  font-weight: bold;
}

nav .bg-blue-600:hover {
  background-color: #2563eb;
}

nav .bg-pink-600:hover {
  background-color: #db2777;
}

.nav-logo {
  display: flex;
  align-items: center;
}

.nav-logo img {
  margin-left: 10px;
}

.drawer-content ul li {
  cursor: pointer;
}

.drawer-content ul li:hover {
  background-color: rgba(255, 255, 255, 0.1);
}
</style>
