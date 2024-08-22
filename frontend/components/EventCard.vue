<template>
  <div
    :class="[
      'bg-white dark:bg-gray-100 rounded-lg shadow-lg overflow-hidden',
      cardClass,
    ]"
  >
    <div v-if="header" :class="['flex items-center p-4', headerClass]">
      <img
        v-if="header.avatar"
        :src="header.avatar"
        alt="Avatar"
        class="w-10 h-10 rounded-full mr-4"
      />
      <div v-if="header.name" class="flex-1">
        <h3 class="text-lg font-semibold dark:text-black">{{ header.name }}</h3>
      </div>
      <div v-if="header.actions" class="flex space-x-4">
        <slot name="header-actions">
          <button
            v-for="(action, idx) in header.actions"
            :key="idx"
            :class="action.class"
            @click="action.handler"
          >
            {{ action.icon }}
          </button>
        </slot>
      </div>
    </div>

    <img
      v-if="image"
      :src="image.src"
      :alt="image.alt"
      :class="['w-full object-cover', imageClass]"
      @click.stop="$emit('navigate')"
    />

    <div
      v-if="title || description || address || price || deadline"
      :class="['p-4', contentClass]"
      @click.stop="$emit('navigate')"
    >
      <h2 v-if="title" class="text-xl font-bold dark:text-black">
        {{ title }}
      </h2>
      <p
        v-if="description"
        class="mt-2 text-gray-600 dark:text-gray-500"
        :class="{ truncate: truncateDescription }"
      >
        {{ description }}
      </p>
      <p v-if="address" class="mt-2 text-gray-500 dark:text-gray-400">
        {{ address }}
      </p>
      <p
        v-if="price"
        class="mt-2 text-gray-800 dark:text-gray-300 font-semibold"
      >
        {{ price }}
      </p>
      <p v-if="deadline" class="mt-2 text-gray-600 dark:text-gray-500">
        {{ deadline }}
      </p>
    </div>

    <div
      v-if="footer"
      :class="['flex items-center justify-between p-4', footerClass]"
    >
      <slot name="footer-left">
        <button
          v-if="footer.reserve"
          @click="footer.reserve.handler"
          :class="footer.reserve.class"
        >
          Reserve
        </button>
      </slot>
      <slot name="footer-right">
        <span
          v-if="footer.postTime"
          :class="footer.postTime.class"
          class="dark:text-gray-500"
          >{{ footer.postTime.text }}</span
        >
      </slot>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from "vue";

const props = defineProps({
  header: {
    type: Object,
    default: () => ({}),
  },
  image: {
    type: Object,
    default: () => ({}),
  },
  title: {
    type: String,
    default: "",
  },
  description: {
    type: String,
    default: "",
  },
  footer: {
    type: Object,
    default: () => ({}),
  },
  address: {
    type: String,
    default: "",
  },
  price: {
    type: String,
    default: "",
  },
  deadline: {
    type: String,
    default: "",
  },
  cardClass: {
    type: String,
    default: "",
  },
  headerClass: {
    type: String,
    default: "",
  },
  imageClass: {
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
  truncateDescription: {
    type: Boolean,
    default: true,
  },
});

const emits = defineEmits(["header-action", "footer-action"]);
</script>

<style scoped>
.truncate {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
