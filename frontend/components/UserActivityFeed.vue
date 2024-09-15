<template>
  <div class="bg-gray-50 p-4 rounded-lg">
    <h2 class="text-lg font-semibold">User Activity</h2>
    <ul>
      <li
        v-for="activity in activities"
        :key="activity.id"
        class="py-2 border-b"
      >
        <p class="text-sm text-gray-700">{{ activity.description }}</p>
        <span class="text-xs text-gray-500">{{
          formatDistance(new Date(activity.timestamp), new Date(), {
            addSuffix: true,
          })
        }}</span>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { format, formatDistance } from "date-fns";
import { ref } from "vue";
import GET_ACTIVITIES from "~/graphql/query/getActivities.gql";

const activities = ref([]);

const { onResult, onError } = useQueryComposable(GET_ACTIVITIES);

onResult((result) => {
  if (result.data) {
    activities.value = result.data.activities;
  }
});
</script>
