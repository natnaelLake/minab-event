<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
      <div v-if="showModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
        <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full">
          <h2 class="text-lg font-semibold mb-4">Confirm Logout</h2>
          <p class="mb-4">Are you sure you want to log out? Any unsaved changes will be lost.</p>
          <div class="flex justify-end space-x-4">
            <button @click="handleLogout" class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600">Logout</button>
            <button @click="cancelLogout" class="bg-gray-300 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-400">Cancel</button>
          </div>
        </div>
      </div>
      <button @click="openModal" class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
        Logout
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useAuthStore } from '~/store';
  
  const showModal = ref(false);
  const router = useRouter();
  const authStore = useAuthStore();
  
  const openModal = () => {
    showModal.value = true;
  };
  
  const cancelLogout = () => {
    showModal.value = false;
  };
  
  const handleLogout = () => {
    authStore.logout();
    router.push('/login');
  };
  </script>
  
  <style lang="scss" scoped>
  /* Custom styles can go here */
  </style>
  