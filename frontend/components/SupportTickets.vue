<script setup>
import { formatDistance } from "date-fns";
import { ref } from "vue";
import GET_SUPPORT_TICKETS from "~/graphql/query/GetSupportTickets.gql";

const { onResult, loading, onError } = useQueryComposable(GET_SUPPORT_TICKETS);
const tickets = ref([]);

onResult((result) => {
  if (result.data) {
    tickets.value = result.data.support_tickets.map((ticket) => ({
      ...ticket,
    }));
  }
});

const viewTicket = (ticketId) => {
  // Logic to view ticket details
  alert(`Viewing ticket ID ${ticketId}`);
};
</script>

<template>
  <div class="bg-white shadow-md rounded-lg overflow-hidden">
    <h2 class="p-4 text-lg font-semibold bg-gray-100">Support Tickets</h2>
    <div class="overflow-hidden">
      <table class="min-w-full">
        <thead class="bg-gray-200 sticky top-0">
          <tr>
            <th class="p-4 text-left">Ticket ID</th>
            <th class="p-4 text-left">Title</th>
            <th class="p-4 text-left">Description</th>
            <th class="p-4 text-left">Date</th>
            <th class="p-4 text-left">Status</th>
            <th class="p-4 text-left">Actions</th>
          </tr>
        </thead>
      </table>
      <!-- Scrollable Table Body -->
      <div class="overflow-y-auto max-h-60">
        <table class="min-w-full">
          <tbody>
            <tr
              v-for="ticket in tickets"
              :key="ticket.id"
              class="hover:bg-gray-50"
            >
              <td class="p-4">{{ ticket.id }}</td>
              <td class="p-4">{{ ticket.title }}</td>
              <td class="p-4">{{ ticket.description }}</td>
              <td class="p-4">
                {{
                  formatDistance(new Date(ticket.created_at), new Date(), {
                    addSuffix: true,
                  })
                }}
              </td>
              <td class="p-4">
                <span
                  :class="{
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-700 text-white': ticket.status === 'Open',
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-700 text-white': ticket.status === 'Closed'
                  }"
                >
                  {{ ticket.status }}
                </span>
              </td>
              <!-- <td class="p-4">
                <button
                  @click="viewTicket(ticket.id)"
                  class="text-blue-500 hover:underline"
                >
                  View
                </button>
              </td> -->
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Ensures that only the table body is scrollable */
.overflow-y-auto {
  max-height: 400px; /* Adjust this height based on your needs */
}

.sticky {
  position: sticky;
  top: 0;
  z-index: 10;
  background-color: #f7fafc; /* Matches the background color */
}
</style>
