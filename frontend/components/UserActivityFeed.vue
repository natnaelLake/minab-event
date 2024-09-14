<template>
    <div class="bg-gray-50 p-4 rounded-lg">
      <h2 class="text-lg font-semibold">User Activity</h2>
      <ul>
        <li v-for="activity in activities" :key="activity.id" class="py-2 border-b">
          <p class="text-sm text-gray-700">{{ activity.description }}</p>
          <span class="text-xs text-gray-500">{{ activity.timestamp }}</span>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  const activities = [
    { id: 1, description: 'User JohnDoe booked a ticket for Music Concert', timestamp: '2 hours ago' },
    { id: 2, description: 'User JaneDoe created a new event: Art Exhibition', timestamp: '1 day ago' },
    // More activities here...
  ]
  </script>
  