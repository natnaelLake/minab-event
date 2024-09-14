<template>
  <div
    :class="[
      'bg-white dark:bg-gray-100 rounded-lg shadow-lg overflow-hidden animate-pulse',
      cardClass,
      'w-full' // Ensure it takes the full width of the grid column
    ]"
  >
    <div v-if="header" :class="['flex items-center p-4', headerClass]">
      <div class="w-10 h-10 rounded-full bg-gray-300 dark:bg-gray-400 mr-4"></div>
      <div class="flex-1">
        <div class="h-4 bg-gray-300 dark:bg-gray-400 rounded w-3/4"></div>
      </div>
      <div class="flex space-x-4">
        <div class="w-5 h-5 bg-gray-300 dark:bg-gray-400 rounded"></div>
        <div class="w-5 h-5 bg-gray-300 dark:bg-gray-400 rounded"></div>
      </div>
    </div>

    <div class="w-full h-48 bg-gray-300 dark:bg-gray-400"></div>

    <div :class="['p-4', contentClass]">
      <div class="h-5 bg-gray-300 dark:bg-gray-400 rounded w-1/2 mb-2"></div>
      <div class="h-4 bg-gray-300 dark:bg-gray-400 rounded w-full mb-2"></div>
      <div class="h-4 bg-gray-300 dark:bg-gray-400 rounded w-2/3 mb-2"></div>
      <div class="h-4 bg-gray-300 dark:bg-gray-400 rounded w-1/4"></div>
    </div>

    <div :class="['flex items-center justify-between p-4', footerClass]">
      <div class="h-8 bg-gray-300 dark:bg-gray-400 rounded w-20"></div>
      <div class="h-4 bg-gray-300 dark:bg-gray-400 rounded w-1/4"></div>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from "vue";

const props = defineProps({
  cardClass: {
    type: String,
    default: "",
  },
  headerClass: {
    type: String,
    default: "",
  },
  contentClass: {
    type: String,
    default: "",
  },
  footerClass: {
    type: String,
    default: "",
  },
});
</script>
