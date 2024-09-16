<script setup>
import { ref, onMounted } from "vue";
import ChartContainer from "~/components/ChartContainer.vue";
import useQueryComposable from "~/composables/useQueryComposable";
import GetMonthlyUserGrowth from "~/graphql/query/GetMonthlyUserGrowth.gql";
import GetMonthlyRevenue from "~/graphql/query/GetMonthlyRevenue.gql";
import GetMostPopularEvents from "~/graphql/query/GetMostPopularEvents.gql";
import GetTopUsersByEventsAttended from "~/graphql/query/GetTopUsersByEventsAttended.gql";
import GetTotalBookingsPerEvent from "~/graphql/query/GetTotalBookingsPerEvent.gql";
import GetTotalTicketsSold from "~/graphql/query/GetTotalTicketsSold.gql";
import { format } from "date-fns";

const monthlyRevenueData = ref({});
const monthlyUserGrowthData = ref({});
const mostPopularEventsData = ref({});
const topUsersByEventsAttendedData = ref({});
const totalBookingsPerEventData = ref({});
const totalTicketsSoldData = ref({});

onMounted(async () => {
  const monthlyRevenue = ref([]);
  const { onResult, onError } = useQueryComposable(GetMonthlyRevenue);
  onResult((result) => {
    monthlyRevenue.value = result.data;
    monthlyRevenueData.value = {
      labels: monthlyRevenue.value.monthly_revenue.map((rev) => format(new Date(rev.month), "MMM yyyy")),
      datasets: [
        {
          label: "Revenue",
          backgroundColor: "#4CAF50",
          data: monthlyRevenue.value.monthly_revenue.map((rev) => rev.total_revenue),
        },
      ],
    };
  });

  const monthlyUserGrowth = ref([]);
  const { onResult: onUserGrowthResult, onError: onUserGrowthError } =
    useQueryComposable(GetMonthlyUserGrowth);
  onUserGrowthResult((result) => {
    if (result.data) {
      monthlyUserGrowth.value = result.data;
      monthlyUserGrowthData.value = {
        labels: monthlyUserGrowth.value.monthly_user_growth.map(
          (user) => format(new Date(user.month),"MMM yyyy")
        ),
        datasets: [
          {
            label: "User Growth",
            backgroundColor: "#FFCE56",
            data: monthlyUserGrowth.value.monthly_user_growth.map(
              (user) => user.user_count
            ),
          },
        ],
      };
    }
  });

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
  const {
    onResult: onResultByEventsAttended,
    onError: onErrorByEventsAttended,
  } = useQueryComposable(GetTopUsersByEventsAttended);
  onResultByEventsAttended((result) => {
    if (result.data) {
      topUsersByEventsAttended.value = result.data;
      topUsersByEventsAttendedData.value = {
        labels: topUsersByEventsAttended.value.users.map((user) => {return user.first_name + " " + user.last_name}),
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
  const {
    onResult: onResultBookingsPerEvent,
    onError: onErrorBookingsPerEvent,
  } = useQueryComposable(GetTotalBookingsPerEvent);
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

  const totalTicketsSold = ref({});
  const { onResult: onResultTicketsSold, onError: onErrorTicketsSold } =
    useQueryComposable(GetTotalTicketsSold);
  onResultTicketsSold((result) => {
    if (result.data) {
      totalTicketsSold.value = result.data;
      totalTicketsSoldData.value = {
        labels: ["Total Tickets Sold"],
        datasets: [
          {
            label: "Tickets Sold",
            backgroundColor: ["#FF6384"],
            data: [totalTicketsSold.value.tickets_aggregate.aggregate.count],
          },
        ],
      };
    }
  });
});
</script>

<template>
  <div class="p-6 mt-20">
    <h1 class="text-3xl font-bold mb-6 text-gray-800">Reports</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">

      <!-- Chart for Monthly Revenue -->
      <ChartContainer
        title="Monthly Revenue"
        chartType="line"
        :chartData="monthlyRevenueData"
      />

      <!-- Chart for Monthly User Growth -->
      <ChartContainer
        title="Monthly User Growth"
        chartType="line"
        :chartData="monthlyUserGrowthData"
      />

      <!-- Chart for Most Popular Events -->
      <ChartContainer
        title="Most Popular Events"
        chartType="bar"
        :chartData="mostPopularEventsData"
      />

      <!-- Chart for Top Users by Events Attended -->
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

      <!-- Chart for Total Tickets Sold -->
      <ChartContainer
        title="Total Tickets Sold"
        chartType="pie"
        :chartData="totalTicketsSoldData"
      />
    </div>
  </div>
</template>

<style scoped>
/* Add responsive spacing and styling for better UI */
h1 {
  color: #333;
  margin-bottom: 2rem;
}

.grid {
  gap: 1.5rem; /* Adjust grid spacing for better layout */
}

/* Customize padding and other styling for each card container */
.chart-container {
  padding: 1.5rem;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.chart-container:hover {
  transform: translateY(-5px);
}
</style>
