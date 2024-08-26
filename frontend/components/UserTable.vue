<template>
  <div class="p-6 mt-20">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold">User Management</h1>
      <button
        v-if="isSuperAdmin"
        class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
        @click="openModal('add')"
      >
        Add Admin
      </button>
    </div>
    <div class="mb-6">
      <input
        type="text"
        placeholder="Search Users..."
        class="w-full p-4 border rounded"
      />
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
              <tr
                v-for="user in paginatedUsers"
                :key="user.id"
                class="hover:bg-gray-50"
              >
                <td class="p-4">
                  {{ user.first_name + " " + user.last_name }}
                </td>
                <td class="p-4">{{ user.email }}</td>
                <td class="p-4">{{ user.role }}</td>
                <td class="p-4">active</td>
                <td class="p-4 space-y-2">
                  <button
                    class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                    @click="suspendUser(user.id)"
                  >
                    Suspend
                  </button>
                  <button
                    v-if="isSuperAdmin && user.role === 'admin'"
                    class="bg-yellow-500 text-white px-4 ml-5 py-2 rounded hover:bg-yellow-600"
                    @click="openModal('edit', user)"
                  >
                    Edit Admin
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
    <div
      v-if="isModalVisible"
      class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50"
    >
      <div class="bg-white p-6 rounded-lg shadow-md w-full max-w-md">
        <h2 class="text-2xl font-bold mb-4">
          {{ modalAction === "add" ? "Add Admin" : "Edit Admin" }}
        </h2>
        <input
          type="text"
          v-model="modalData.first_name"
          placeholder="First Name"
          class="w-full p-2 mb-4 border rounded"
        />
        <input
          type="text"
          v-model="modalData.last_name"
          placeholder="Last Name"
          class="w-full p-2 mb-4 border rounded"
        />
        <input
          type="email"
          v-model="modalData.email"
          placeholder="Email"
          class="w-full p-2 mb-4 border rounded"
        />

        <div class="flex justify-end space-x-4">
          <button
            class="bg-gray-300 px-4 py-2 rounded hover:bg-gray-400"
            @click="closeModal"
          >
            Cancel
          </button>
          <button
            class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
            @click="submitModal"
          >
            {{ modalAction === "add" ? "Add" : "Save" }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, computed, onMounted } from "vue";
import { useQuery } from "@vue/apollo-composable";
import { useRouter, useRoute } from "vue-router";
import getAllUsers from "~/graphql/query/getAllUsers.gql";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const users = ref([]);

// Reactive State
onMounted(async () => {
  await fetchEvents();
});

const fetchEvents = async () => {
  try {
    const { onResult, onError } = useQuery(getAllUsers, {
      limit: 10,
      offset: 0,
      order_by: [{ created_at: "desc" }],
      where: {},
    });

    onResult((result) => {
      if (result.data) {
        console.log("Fetched Users: ", result.data.users);
        users.value = result.data.users;
      }
    });

    onError((error) => {
      console.error("Error fetching users: ", error.message);
      toast.error("Something went wrong, try again", {
        transition: toast.TRANSITIONS.FLIP,
        position: toast.POSITION.TOP_RIGHT,
      });
    });
  } catch (error) {
    console.error("Error during fetching users: ", error);
    toast.error("Failed to load users.");
  }
};

// Pagination
const currentPage = ref(1);
const itemsPerPage = 5;
const totalPages = computed(() => Math.ceil(users.value.length / itemsPerPage));
const paginatedUsers = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  return users.value.slice(start, start + itemsPerPage);
});

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value += 1;
  }
};

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value -= 1;
  }
};

// Check if the user is a super admin
const isSuperAdmin = computed(() => authStore.role === "admin");

// Modal State and Functions
const isModalVisible = ref(false);
const modalAction = ref(""); // "add" or "edit"
const modalData = ref({ first_name: "", last_name: "", email: "" });

const openModal = (action, user = null) => {
  modalAction.value = action;
  isModalVisible.value = true;
  if (action === "edit" && user) {
    modalData.value = { ...user };
  } else {
    modalData.value = { first_name: "", last_name: "", email: "" };
  }
};

const closeModal = () => {
  isModalVisible.value = false;
  modalData.value = { first_name: "", last_name: "", email: "" };
};

const submitModal = () => {
  if (modalAction.value === "add") {
    addAdmin();
  } else if (modalAction.value === "edit") {
    editAdmin(modalData.value.id);
  }
  closeModal();
};

// Admin management functions
const addAdmin = () => {
  console.log(`Adding admin: ${JSON.stringify(modalData.value)}`);
  // Implementation for adding an admin
};

const editAdmin = (userId) => {
  console.log(`Editing admin ${userId}`);
  // Implementation for editing an admin
};

// Suspend User Function
const suspendUser = (userId) => {
  const user = users.value.find((user) => user.id === userId);
  if (user) {
    user.status = "Suspended";
    toast.success(
      `User ${user.first_name} ${user.last_name} has been suspended.`
    );
  }
};
</script>

<style scoped>
/* Limit the height of the table body and make it scrollable */
.overflow-y-auto {
  max-height: 400px;
}

/* Ensure the table header stays fixed at the top while scrolling */
thead.sticky {
  position: sticky;
  z-index: 10; /* Lower z-index for sticky table header */
}

/* Modal backdrop to cover everything, including sticky headers */
.fixed.inset-0 {
  z-index: 50; /* Higher z-index for modal backdrop */
}

.bg-white.p-6.rounded-lg.shadow-md {
  z-index: 60; /* Higher z-index for modal content */
}
</style>
