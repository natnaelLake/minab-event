<template>
    <div class="bg-blue-500 text-white rounded-lg p-4 shadow-lg">
      <h2 class="text-xl font-semibold">{{ title }}</h2>
      <p class="text-3xl font-bold mt-2">{{ stat }}</p>
    </div>
  </template>
  
  <script setup>
  const props = defineProps({
    title: String,
    stat: String | Number
  })
  </script>
  