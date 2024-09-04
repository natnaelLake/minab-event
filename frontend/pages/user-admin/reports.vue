<template>
    <div class="p-6">
      <h1 class="text-3xl font-bold mb-6">Reports</h1>
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-4">
        <ChartContainer title="Revenue Trends" />
        <ChartContainer title="Event Popularity" />
        <ChartContainer title="User Growth" />
      </div>
    </div>
  </template>
  
  <script setup>
  import ChartContainer from '~/components/ChartContainer.vue'
  </script>
  