<template>
    <div class="bg-white shadow-md rounded-lg overflow-hidden">
      <h2 class="p-4 text-lg font-semibold bg-gray-100">Support Tickets</h2>
      <div class="overflow-x-auto">
        <table class="min-w-full">
          <thead class="bg-gray-200">
            <tr>
              <th class="p-4 text-left">Ticket ID</th>
              <th class="p-4 text-left">Title</th>
              <th class="p-4 text-left">Date</th>
              <th class="p-4 text-left">Status</th>
              <th class="p-4 text-left">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="ticket in tickets" :key="ticket.id" class="hover:bg-gray-50">
              <td class="p-4">{{ ticket.id }}</td>
              <td class="p-4">{{ ticket.title }}</td>
              <td class="p-4">{{ ticket.date }}</td>
              <td class="p-4">{{ ticket.status }}</td>
              <td class="p-4">
                <button
                  class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                  @click="viewTicket(ticket.id)"
                >
                  View
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </template>
  
  <script setup>
  const tickets = [
    { id: 101, title: "Login Issue", date: "2024-08-21", status: "Open" },
    { id: 102, title: "Payment Error", date: "2024-08-20", status: "Closed" },
    // Add more tickets here
  ];
  
  const viewTicket = (ticketId) => {
    // Logic to view ticket details
    alert(`Viewing ticket ID ${ticketId}`);
  };
  </script>
  
  <style scoped>
  /* Optional: Customize table scroll height */
  .overflow-x-auto {
    max-height: 400px;
  }
  </style>
  