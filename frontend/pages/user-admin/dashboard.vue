<script setup>
import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import StatCard from "~/components/StatCard.vue";
import GetDashboardStats from "~/graphql/query/GetDashboardStats.gql";
import RecentEventsTable from "~/components/RecentEventsTable.vue";
import UserActivityFeed from "~/components/UserActivityFeed.vue";
import ChartContainer from "~/components/ChartContainer.vue";
import { useAuthStore } from "~/store";
import GetMostPopularEvents from "~/graphql/query/GetMostPopularEvents.gql";
import GetTopUsersByEventsAttended from "~/graphql/query/GetTopUsersByEventsAttended.gql";
import GetTotalBookingsPerEvent from "~/graphql/query/GetTotalBookingsPerEvent.gql";

const mostPopularEventsData = ref({});
const topUsersByEventsAttendedData = ref({});
const totalBookingsPerEventData = ref({});
const mostPopularEvents = ref();
const { onResult: onPopularEventsResult, onError: onPopularEventsError } =
  useQueryComposable(GetMostPopularEvents);
onPopularEventsResult((result) => {
  if (result.data) {
    mostPopularEvents.value = result.data;
    mostPopularEventsData.value = {
      labels: mostPopularEvents.value.events.map((event) => event.title),
      datasets: [
        {
          label: "Bookings",
          backgroundColor: "#FF6384",
          data: mostPopularEvents.value.events.map(
            (event) => event.total_tickets.aggregate.count
          ),
        },
      ],
    };
  }
});
const topUsersByEventsAttended = ref([]);
const { onResult: onResultByEventsAttended, onError: onErrorByEventsAttended } =
  useQueryComposable(GetTopUsersByEventsAttended);
onResultByEventsAttended((result) => {
  if (result.data) {
    topUsersByEventsAttended.value = result.data;
    topUsersByEventsAttendedData.value = {
      labels: topUsersByEventsAttended.value.users.map((user) => {
        return user.first_name + " " + user.last_name;
      }),
      datasets: [
        {
          label: "Events Attended",
          backgroundColor: "#36A2EB",
          data: topUsersByEventsAttended.value.users.map(
            (user) => user.total_events_attended.aggregate.count
          ),
        },
      ],
    };
  }
});

const totalBookingsPerEvent = ref();
const { onResult: onResultBookingsPerEvent, onError: onErrorBookingsPerEvent } =
  useQueryComposable(GetTotalBookingsPerEvent);
onResultBookingsPerEvent((result) => {
  if (result.data) {
    totalBookingsPerEvent.value = result.data;
    totalBookingsPerEventData.value = {
      labels: totalBookingsPerEvent.value.events.map((event) => event.title),
      datasets: [
        {
          label: "Bookings",
          backgroundColor: "#FF6384",
          data: totalBookingsPerEvent.value.events.map(
            (event) => event.total_bookings.aggregate.count
          ),
        },
      ],
    };
  }
});
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
  GetDashboardStats,
  {},
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
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-6">
      <ChartContainer
        title="Most Popular Events"
        chartType="bar"
        :chartData="mostPopularEventsData"
      />
      <ChartContainer
        title="Top Users by Events Attended"
        chartType="bar"
        :chartData="topUsersByEventsAttendedData"
      />
      <ChartContainer
        title="Total Bookings per Event"
        chartType="bar"
        :chartData="totalBookingsPerEventData"
      />
    </div>
  </div>
</template>
