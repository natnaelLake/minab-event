<template>
    <div class="p-6">
      <div class="bg-white shadow-md rounded-lg">
        <h2 class="p-4 text-lg font-semibold bg-gray-100">Recent Events</h2>
        <div class="relative">
          <!-- Scrollable Table Container -->
          <div class="overflow-y-auto max-h-96">
            <table class="min-w-full">
              <thead class="bg-gray-200 sticky top-0">
                <tr>
                  <th class="p-4 text-left">Event Name</th>
                  <th class="p-4 text-left">Date</th>
                  <th class="p-4 text-left">Category</th>
                  <th class="p-4 text-left">Status</th>
                  <th class="p-4 text-left">Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="event in paginatedEvents" :key="event.id" class="hover:bg-gray-50">
                  <td class="p-4">{{ event.name }}</td>
                  <td class="p-4">{{ event.date }}</td>
                  <td class="p-4">{{ event.category }}</td>
                  <td class="p-4">{{ event.status }}</td>
                  <td class="p-4">
                    <button
                      class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                      @click="blockEvent(event.id)"
                    >
                      Block
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
  
          <!-- Pagination Controls -->
          <div class="flex justify-end items-center p-4 bg-gray-100 space-x-5">
            <button
              @click="prevPage"
              :disabled="currentPage === 1"
              class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
            >
              Previous
            </button>
            <span>Page {{ currentPage }} of {{ totalPages }}</span>
            <button
              @click="nextPage"
              :disabled="currentPage === totalPages"
              class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, computed } from 'vue'
const events = ref([
  {
    id: 1,
    name: "Music Concert",
    date: "2024-09-01",
    category: "Music",
    status: "Active",
  },
  {
    id: 2,
    name: "Art Exhibition",
    date: "2024-09-10",
    category: "Art",
    status: "Pending",
  },
  {
    id: 3,
    name: "Illegal Gathering",
    date: "2024-09-15",
    category: "Social",
    status: "Illegal",
  },
  {
    id: 1,
    name: "Music Concert",
    date: "2024-09-01",
    category: "Music",
    status: "Active",
  },
  {
    id: 2,
    name: "Art Exhibition",
    date: "2024-09-10",
    category: "Art",
    status: "Pending",
  },
  {
    id: 3,
    name: "Illegal Gathering",
    date: "2024-09-15",
    category: "Social",
    status: "Illegal",
  },
  {
    id: 1,
    name: "Music Concert",
    date: "2024-09-01",
    category: "Music",
    status: "Active",
  },
  {
    id: 2,
    name: "Art Exhibition",
    date: "2024-09-10",
    category: "Art",
    status: "Pending",
  },
  {
    id: 3,
    name: "Illegal Gathering",
    date: "2024-09-15",
    category: "Social",
    status: "Illegal",
  },
  {
    id: 1,
    name: "Music Concert",
    date: "2024-09-01",
    category: "Music",
    status: "Active",
  },
  {
    id: 2,
    name: "Art Exhibition",
    date: "2024-09-10",
    category: "Art",
    status: "Pending",
  },
  {
    id: 3,
    name: "Illegal Gathering",
    date: "2024-09-15",
    category: "Social",
    status: "Illegal",
  },
  {
    id: 1,
    name: "Music Concert",
    date: "2024-09-01",
    category: "Music",
    status: "Active",
  },
  {
    id: 2,
    name: "Art Exhibition",
    date: "2024-09-10",
    category: "Art",
    status: "Pending",
  },
  {
    id: 3,
    name: "Illegal Gathering",
    date: "2024-09-15",
    category: "Social",
    status: "Illegal",
  },
  {
    id: 1,
    name: "Music Concert",
    date: "2024-09-01",
    category: "Music",
    status: "Active",
  },
  {
    id: 2,
    name: "Art Exhibition",
    date: "2024-09-10",
    category: "Art",
    status: "Pending",
  },
  {
    id: 3,
    name: "Illegal Gathering",
    date: "2024-09-15",
    category: "Social",
    status: "Illegal",
  },

  // More events here...
]);


// Pagination
const currentPage = ref(1)
const itemsPerPage = 5

const totalPages = computed(() => Math.ceil(events.value.length / itemsPerPage))

const paginatedEvents = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  return events.value.slice(start, start + itemsPerPage)
})

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value += 1
  }
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value -= 1
  }
}

// Block Event Function
const blockEvent = (eventId) => {
  // Logic to block the event
  alert(`Event with ID ${eventId} has been blocked.`)
}
</script>

<style scoped>
/* Limit the height of the table body and make it scrollable */
.overflow-y-auto {
  max-height: 400px;
}

/* Ensure the table header stays fixed at the top while scrolling */
thead.sticky {
  position: sticky;
  z-index: 10;
}
</style>
