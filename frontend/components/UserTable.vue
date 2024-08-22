<template>
    <div class="p-6 mt-20">
      <h1 class="text-3xl font-bold mb-6">User Management</h1>
      <div class="mb-6">
        <input type="text" placeholder="Search Users..." class="w-full p-4 border rounded" />
      </div>
      <div class="bg-white shadow-md rounded-lg">
        <h2 class="p-4 text-lg font-semibold bg-gray-100">User List</h2>
        <div class="relative">
          <!-- Scrollable Table Container -->
          <div class="overflow-y-auto max-h-96">
            <table class="min-w-full">
              <thead class="bg-gray-200 sticky top-0">
                <tr>
                  <th class="p-4 text-left">User Name</th>
                  <th class="p-4 text-left">Email</th>
                  <th class="p-4 text-left">Role</th>
                  <th class="p-4 text-left">Status</th>
                  <th class="p-4 text-left">Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="user in paginatedUsers" :key="user.id" class="hover:bg-gray-50">
                  <td class="p-4">{{ user.name }}</td>
                  <td class="p-4">{{ user.email }}</td>
                  <td class="p-4">{{ user.role }}</td>
                  <td class="p-4">{{ user.status }}</td>
                  <td class="p-4">
                    <button
                      class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                      @click="suspendUser(user.id)"
                    >
                      Suspend
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
  
  const users = ref([
    { id: 1, name: 'John Doe', email: 'john@example.com', role: 'Admin', status: 'Active' },
    { id: 2, name: 'Jane Smith', email: 'jane@example.com', role: 'User', status: 'Active' },
    { id: 3, name: 'Emily Johnson', email: 'emily@example.com', role: 'User', status: 'Suspended' },
    { id: 1, name: 'John Doe', email: 'john@example.com', role: 'Admin', status: 'Active' },
    { id: 2, name: 'Jane Smith', email: 'jane@example.com', role: 'User', status: 'Active' },
    { id: 3, name: 'Emily Johnson', email: 'emily@example.com', role: 'User', status: 'Suspended' },
    { id: 1, name: 'John Doe', email: 'john@example.com', role: 'Admin', status: 'Active' },
    { id: 2, name: 'Jane Smith', email: 'jane@example.com', role: 'User', status: 'Active' },
    { id: 3, name: 'Emily Johnson', email: 'emily@example.com', role: 'User', status: 'Suspended' },
    // More user data here...
  ])
  
  // Pagination
  const currentPage = ref(1)
  const itemsPerPage = 5
  
  const totalPages = computed(() => Math.ceil(users.value.length / itemsPerPage))
  
  const paginatedUsers = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage
    return users.value.slice(start, start + itemsPerPage)
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
  
  // Suspend User Function
  const suspendUser = (userId) => {
    const user = users.value.find(user => user.id === userId)
    if (user) {
      user.status = 'Suspended'
      alert(`User ${user.name} has been suspended.`)
    }
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
  