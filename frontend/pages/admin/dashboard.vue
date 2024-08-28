<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold mb-6">Admin Dashboard</h1>
    
    <!-- Stat Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <StatCard title="Total Events" :stat="150" />
      <StatCard title="Active Users" :stat="1200" />
      <StatCard title="Active Events" :stat="45000" />
      <StatCard title="Total Users" :stat="34" />
    </div>

    <!-- Recent Activities -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 mb-6">
      <RecentEventsTable />
      <UserActivityFeed />
    </div>

    <!-- Charts & Reports -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-6">
      <ChartContainer title="Event Popularity" />
      <ChartContainer title="Revenue Trends" />
      <ChartContainer title="User Growth" />
    </div>

    <!-- Tasks & Announcements -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
      <TaskList />
      <Announcements />
    </div>
  </div>
</template>

<script setup>
import StatCard from '~/components/StatCard.vue'
import RecentEventsTable from '~/components/RecentEventsTable.vue'
import UserActivityFeed from '~/components/UserActivityFeed.vue'
import ChartContainer from '~/components/ChartContainer.vue'
import TaskList from '~/components/TaskList.vue'
import Announcements from '~/components/Announcements.vue'
</script>
