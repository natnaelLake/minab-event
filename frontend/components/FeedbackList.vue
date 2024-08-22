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
  const feedbacks = [
    { id: 1, title: "Website Issue", description: "The homepage is slow to load.", status: "Resolved" },
    { id: 2, title: "Feature Request", description: "Add dark mode option.", status: "Pending" },
    // Add more feedback items here
  ];
  </script>
  
  <style scoped>
  /* Optional: Customize table scroll height */
  .overflow-x-auto {
    max-height: 400px;
  }
  </style>
  