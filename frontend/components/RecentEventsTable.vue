<script setup>
import { ref, computed, onMounted, watch } from "vue";
import GetEvents from "~/graphql/query/GetEvents.gql";
import updateEventStatus from "~/graphql/mutations/updateEventStatus.gql";
import { format, parseISO } from "date-fns";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";

const events = ref([]);
const itemsPerPage = 5;
const currentPage = ref(1);
const totalEvents = ref(0);
const user = useAuthStore();
const getOffset = () => (currentPage.value - 1) * itemsPerPage;
// Use Apollo's useQuery to fetch events with pagination
const { onResult, onError, refetch } = useQueryComposable(
  GetEvents,
  {
    limit: itemsPerPage,
    offset: getOffset(),
    order_by: [{ created_at: "desc" }],
    where: {},
  }
);
watch(currentPage, (newPage) => {
  refetch({ limit: itemsPerPage, offset: getOffset() });
});
onResult((result) => {
  if (result.data) {
    events.value = result.data.events;
    totalEvents.value = result.data.events_aggregate?.aggregate?.count || 0;
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
const totalPages = computed(() => Math.ceil(totalEvents.value / itemsPerPage));

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

// Apollo mutation for updating event status
const { mutate: updateEventStatusMutation } = useMutationComposable(
  updateEventStatus,
);

const blockEvent = async (id) => {
  try {
    await updateEventStatusMutation({ id, status: "Blocked" });
    toast.success("Event blocked successfully.");
    await refetch(); // Refetch events list after mutation
  } catch (error) {
    console.error("Error updating event status: ", error.message);
    toast.error("Failed to block event.");
  }
};
const parseTags = (tagsString) => {
  return tagsString.replace(/{|}/g, "").split(",");
};
const activateEvent = async (id) => {
  try {
    await updateEventStatusMutation({
      id,
      status: "Active",
    });
    toast.success("Event activated successfully.");
    await refetch(); // Refetch events list after mutation
  } catch (error) {
    console.error("Error updating event status: ", error.message);
    toast.error("Failed to activate event.");
  }
};

const formatEventDate = (date) => {
  if (!date) {
    console.error("Invalid date value:", date);
    return "Invalid Date"; // or you can return an empty string or a default date
  }

  try {
    const parsedDate = parseISO(date);
    return format(parsedDate, "PPpp");
  } catch (error) {
    console.error("Error parsing date:", error);
    return "Invalid Date"; // Handle the error case gracefully
  }
};
</script>
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
                <th class="p-4 text-left w-1/5">Event Name</th>
                <th class="p-4 text-left w-1/5">Date</th>
                <th class="p-4 text-left w-1/5">Category</th>
                <th class="p-4 text-left w-1/5">Tags</th>
                <th class="p-4 text-left w-1/5">Status</th>
                <th class="p-4 text-left w-1/5">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="event in events"
                :key="event.id"
                class="hover:bg-gray-50"
              >
                <td class="p-4 truncate w-1/5">{{ event.title }}</td>
                <td class="p-4 truncate w-1/5">
                  {{ formatEventDate(event.event_start_time) }}
                </td>
                <td class="p-4 truncate w-1/5">
                  <span
                    class="inline-block px-2 py-1 bg-blue-100 text-blue-800 text-xs font-semibold rounded"
                  >
                    {{ event.categories }}
                  </span>
                </td>
                <td class="p-4 truncate w-1/5">
                  <div class="flex flex-wrap gap-1">
                    <span
                      v-for="tag in parseTags(event.tags)"
                      :key="tag"
                      class="inline-block px-2 py-1 bg-green-200 text-green-800 text-xs font-semibold rounded"
                    >
                      {{ tag }}
                    </span>
                  </div>
                </td>
                <td class="p-4 truncate w-1/5">
                  <span
                    :class="[
                      event.status === 'Active'
                        ? 'bg-green-500 text-white px-4 py-2 rounded'
                        : '',
                      event.status === 'Blocked'
                        ? 'bg-red-500 text-white px-4 py-2 rounded'
                        : '',
                      event.status === 'Closed'
                        ? 'bg-gray-500 text-white px-4 py-2 rounded'
                        : '',
                      'px-2 py-1 rounded text-xs font-semibold',
                    ]"
                    >{{ event.status }}</span
                  >
                </td>
                <td class="p-4 w-1/5">
                  <button
                    v-if="event.status === 'Active'"
                    class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                    @click="blockEvent(event.id)"
                  >
                    Block
                  </button>
                  <button
                    v-if="event.status === 'Blocked'"
                    class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
                    @click="activateEvent(event.id)"
                  >
                    Activate
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

/* Add text truncation for overflow */
td.truncate {
  max-width: 120px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
