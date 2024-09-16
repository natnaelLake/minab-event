<template>
  <div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
    <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full text-center">
      <h2 class="text-xl font-semibold mb-4">{{ title }}</h2>
      <p class="mb-6">{{ content }}</p>
      <div class="flex justify-center gap-4">
        <slot />
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  title: String,
  content: String,
});
</script>
