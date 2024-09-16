<template>
    <div class="bg-yellow-100 p-4 rounded-lg shadow-inner">
      <h2 class="text-lg font-bold text-yellow-700">Announcements</h2>
      <div v-for="announcement in announcements" :key="announcement.id" class="mt-4">
        <h3 class="text-gray-800">{{ announcement.title }}</h3>
        <p class="text-gray-700">{{ announcement.content }}</p>
        <span class="text-sm text-gray-500">{{ announcement.date }}</span>
      </div>
    </div>
  </template>
  
  <script setup>
  const announcements = [
    { id: 1, title: 'Scheduled Maintenance', content: 'The site will be down for maintenance on 2024-09-05.', date: '2024-08-22' },
    // More announcements here...
  ]
  </script>
  