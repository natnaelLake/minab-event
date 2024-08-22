<template>
    <div class="bg-white p-6 rounded-lg shadow-md">
      <h2 class="text-lg font-semibold mb-4">{{ title }}</h2>
      <!-- Placeholder for Chart -->
      <div class="h-64 bg-gray-200 rounded-lg"></div>
    </div>
  </template>
  
  <script setup>
  const props = defineProps({
    title: String
  })
  </script>
  