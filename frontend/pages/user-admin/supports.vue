<template>
    <div class="p-6">
      <h1 class="text-3xl font-bold mb-6">Support & Feedback</h1>
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
        <FeedbackList />
        <SupportTickets />
      </div>
    </div>
  </template>
  
  <script setup>
  import FeedbackList from '~/components/FeedbackList.vue'
  import SupportTickets from '~/components/SupportTickets.vue'
  </script>
  