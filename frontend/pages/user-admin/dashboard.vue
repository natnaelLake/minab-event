<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold mb-6">Admin Dashboard</h1>

    <!-- Stat Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <StatCard title="Total Events" :stat="stats.totalEvents" />
      <StatCard title="Active Users" :stat="stats.activeUsers" />
      <StatCard title="Active Events" :stat="stats.activeEvents" />
      <StatCard title="Total Users" :stat="stats.totalUsers" />
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
import { ref, onMounted } from "vue";
import { useQuery } from "@vue/apollo-composable";
import StatCard from "~/components/StatCard.vue";
import GetDashboardStats from "~/graphql/query/GetDashboardStats.gql";
import RecentEventsTable from "~/components/RecentEventsTable.vue";
import UserActivityFeed from "~/components/UserActivityFeed.vue";
import ChartContainer from "~/components/ChartContainer.vue";
import TaskList from "~/components/TaskList.vue";
import Announcements from "~/components/Announcements.vue";
import { useAuthStore } from "~/store";
const stats = ref({
  totalEvents: 0,
  activeUsers: 0,
  activeEvents: 0,
  totalUsers: 0,
});
const user = useAuthStore();
const currentUserId = user.id;
const currentUserRole = user.role;
const currentUserToken = user.token;

const { onResult, onError } = useQuery(
  GetDashboardStats,{},
  {
    context: {
      headers: {
        "x-hasura-user-id": currentUserId,
        "x-hasura-role": currentUserRole,
        Authorization: `Bearer ${currentUserToken}`,
      },
    },
  }
);

onResult((result) => {
  if (result.data) {
    stats.value = {
      totalEvents: result.data.totalEvents.aggregate.count,
      activeUsers: result.data.activeUsers.aggregate.count,
      activeEvents: result.data.activeEvents.aggregate.count,
      totalUsers: result.data.totalUsers.aggregate.count,
    };
  }
});

onError((error) => {
  console.error("Error fetching dashboard stats:", error.message);
});
</script>
