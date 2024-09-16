<template>
  <div class="flex">
    <svg
      v-for="n in 5"
      :key="n"
      @click="rate(n)"
      xmlns="http://www.w3.org/2000/svg"
      class="w-6 h-6 cursor-pointer"
      fill="currentColor"
      :class="{
        'text-yellow-500': n <= rating,
        'text-gray-300': n > rating,
      }"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 2.1l2.9 8.1H22l-6.5 4.6 2.5 7.9L12 17.9 7.5 22.7l2.5-7.9L3.5 10.2h7L12 2.1z"
      />
    </svg>
  </div>
</template>

<script setup>
const rate = (value) => {
  const ratingValue = new CustomEvent("rating-selected", {
    detail: { rating: value },
  });
  window.dispatchEvent(ratingValue);
};
const props = defineProps({
  rating: {
    type: Number,
    default: 0,
  },
});
</script>
