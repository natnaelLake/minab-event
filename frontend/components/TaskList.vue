<template>
  <div class="bg-white shadow-md rounded-lg p-4">
    <h2 class="text-lg font-semibold mb-4">Tasks</h2>
    <ul>
      <li v-for="task in tasks" :key="task.id" class="flex items-center py-2">
        <input type="checkbox" class="mr-3" />
        <p class="text-gray-800">{{ task.description }}</p>
      </li>
    </ul>
    <button class="bg-green-500 text-white py-2 px-4 rounded hover:bg-green-600 mt-4">Add New Task</button>
  </div>
</template>

<script setup>
const tasks = [
  { id: 1, description: 'Review new event submissions' },
  { id: 2, description: 'Respond to user feedback' },
  // More tasks here...
]
</script>
