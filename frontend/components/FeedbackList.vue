<template>
  <div class="bg-white shadow-md rounded-lg overflow-hidden">
    <h2 class="p-4 text-lg font-semibold bg-gray-100">Feedback List</h2>
    <div class="overflow-x-auto">
      <table class="min-w-full">
        <thead class="bg-gray-200">
          <tr>
            <th class="p-4 text-left">Title</th>
            <th class="p-4 text-left">Description</th>
            <th class="p-4 text-left">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="feedback in feedbacks" :key="feedback.id" class="hover:bg-gray-50">
            <td class="p-4">{{ feedback.title }}</td>
            <td class="p-4">{{ feedback.description }}</td>
            <td class="p-4">{{ feedback.status }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import GET_FEEDBACK_LIST from '~/graphql/query/GetFeedbackList.gql'


const { onResult, loading, onError } = useQueryComposable(GET_FEEDBACK_LIST);
const feedbacks = ref([]);

onResult((result) => {
  if (result.data) {
    feedbacks.value = result.data.feedbacks;
  }
});
</script>

<style scoped>
.overflow-x-auto {
  max-height: 400px;
}
</style>
