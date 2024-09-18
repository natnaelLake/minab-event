<script setup>
import { ref, computed, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import getAllUsers from "~/graphql/query/getAllUsers.gql";
import insertAdminMutation from "~/graphql/mutations/insertAdmin.gql";
import updateAdminMutation from "~/graphql/mutations/updateAdmin.gql";
import updateUserStatus from "~/graphql/mutations/updateUserStatus.gql";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";
import { hashPassword } from "~/utils/authUtils";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";

const router = useRouter();
const route = useRoute();
const user = useAuthStore();
const users = ref([]);
const itemsPerPage = 5;
const currentPage = ref(1);
const totalUsers = ref(0);
const searchTerm = ref("");

const getOffset = () => (currentPage.value - 1) * itemsPerPage;
const resetPagination = () => {
  currentPage.value = 1;
};
// Reactive State
const { onResult, onError, refetch } = useQueryComposable(getAllUsers, {
  limit: itemsPerPage,
  offset: getOffset(),
  order_by: [{ created_at: "desc" }],
  where: {
    _or: [
      {
        first_name: { _ilike: `%${searchTerm.value}%` },
      },
      {
        last_name: { _ilike: `%${searchTerm.value}%` },
      },
      {
        email: { _ilike: `%${searchTerm.value}%` },
      },
    ],
  },
});
watch(searchTerm, async (newSearchTerm) => {
  resetPagination();
  await refetch({
    limit: itemsPerPage,
    offset: getOffset(),
    where: {
      _or: [
        {
          first_name: { _ilike: `%${searchTerm.value}%` },
        },
        {
          last_name: { _ilike: `%${searchTerm.value}%` },
        },
        {
          email: { _ilike: `%${searchTerm.value}%` },
        },
      ],
    },
  });
});
watch(currentPage, (newPage) => {
  refetch({
    limit: itemsPerPage,
    offset: getOffset(),
    where: {
      _or: [
        {
          first_name: { _ilike: `%${searchTerm.value}%` },
        },
        {
          last_name: { _ilike: `%${searchTerm.value}%` },
        },
        {
          email: { _ilike: `%${searchTerm.value}%` },
        },
      ],
    },
  });
});
onResult((result) => {
  if (result.data) {
    users.value = result.data.users;
    totalUsers.value = result.data.users_aggregate?.aggregate?.count || 0;
  }
});

onError((error) => {
  console.error("Error fetching events: ", error.message);
  toast.error("Something went wrong, try again", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});

// Calculate total pages
const totalPages = computed(() => Math.ceil(totalUsers.value / itemsPerPage));

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value += 1;
    refetch(); // Refetch data for the next page
  }
};

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value -= 1;
    refetch(); // Refetch data for the previous page
  }
};

// Check if the user is a super admin
const isSuperAdmin = computed(() => user.role === "user-admin");

// Modal State and Functions
const isModalVisible = ref(false);
const modalAction = ref(""); // "add" or "edit"
const modalData = ref({ first_name: "", last_name: "", email: "", id: "" });

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

// Apollo mutations for adding and editing admins
const { mutate: insertUser } = useMutationComposable(insertAdminMutation);
const { mutate: updateUser } = useMutationComposable(updateAdminMutation);
const { mutate: updateUserStatusMutation } =
  useMutationComposable(updateUserStatus);

const submitModal = async () => {
  if (modalAction.value === "add") {
    await addAdmin();
  } else if (modalAction.value === "edit") {
    await editAdmin(modalData.value.id);
  }
  await refetch(); // Refetch users list after mutation
  closeModal();
};

// Admin management functions
const addAdmin = async () => {
  const input = {
    first_name: modalData.value.first_name,
    last_name: modalData.value.last_name,
    email: modalData.value.email,
    password: await hashPassword(modalData.value.password),
    role: "user-admin",
  };
  try {
    await insertUser(input);
    await refetch();
    toast.success("Admin added successfully");
  } catch (error) {
    console.error("Error adding admin: ", error.message);
    toast.error("Failed to add admin");
  }
};

const editAdmin = async (id) => {
  const input = {
    first_name: modalData.value.first_name,
    last_name: modalData.value.last_name,
    email: modalData.value.email,
  };
  try {
    await updateUser({ userId: id, changes: input });
    await refetch();
    toast.success("Admin updated successfully");
  } catch (error) {
    console.error("Error updating admin: ", error.message);
    toast.error("Failed to update admin");
  }
};

// User status management functions
const suspendUser = async (id) => {
  try {
    await updateUserStatusMutation({
      id,
      status: "Suspended",
    });
    toast.success("User suspended successfully");
    await refetch();
  } catch (error) {
    console.error("Error suspending user: ", error.message);
    toast.error("Failed to suspend user");
  }
};

const activateUser = async (id) => {
  try {
    await updateUserStatusMutation({
      id,
      status: "Active",
    });
    toast.success("User activated successfully");
    await refetch();
  } catch (error) {
    console.error("Error activating user: ", error.message);
    toast.error("Failed to activate user");
  }
};
</script>

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
        v-model="searchTerm"
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
              <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50">
                <td class="p-4">
                  {{ user.first_name + " " + user.last_name }}
                </td>
                <td class="p-4">{{ user.email }}</td>
                <td class="p-4">{{ user.role }}</td>
                <td class="p-4 truncate w-1/5">
                  <span
                    :class="[
                      user.status === 'Active'
                        ? 'bg-green-500 text-white px-4 py-2 rounded'
                        : '',
                      user.status === 'Suspended'
                        ? 'bg-red-500 text-white px-4 py-2 rounded'
                        : '',
                      'px-2 py-1 rounded text-xs font-semibold',
                    ]"
                    >{{ user.status }}</span
                  >
                </td>
                <td class="p-4 space-y-2">
                  <button
                    v-if="user.status === 'Active'"
                    class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                    @click="suspendUser(user.id)"
                  >
                    Suspend
                  </button>
                  <button
                    v-if="user.status === 'Suspended'"
                    class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
                    @click="activateUser(user.id)"
                  >
                    Activate
                  </button>
                  <button
                    v-if="isSuperAdmin && user.role === 'user-admin'"
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
        <input
          v-if="modalAction === 'add'"
          type="text"
          v-model="modalData.password"
          placeholder="Password ..."
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

<style scoped>
.bg-green-500 {
  background-color: #48bb78;
}

.hover\:bg-green-600:hover {
  background-color: #38a169;
}

.bg-yellow-500 {
  background-color: #ecc94b;
}

.hover\:bg-yellow-600:hover {
  background-color: #d69e2e;
}
</style>
