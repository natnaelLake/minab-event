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
                v-for="event in paginatedEvents"
                :key="event.id"
                class="hover:bg-gray-50"
              >
                <td class="p-4 truncate w-1/5">{{ event.title }}</td>
                <td class="p-4 truncate w-1/5">
                  {{
                    formatEventDate(event.event_date, event.event_start_time)
                  }}
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
                      v-for="tag in event.tags"
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
                      event.status === 'active'
                        ? 'bg-green-100 text-green-800'
                        : '',
                      event.status === 'blocked'
                        ? 'bg-red-100 text-red-800'
                        : '',
                      event.status === 'closed'
                        ? 'bg-gray-200 text-gray-800'
                        : '',
                      'px-2 py-1 rounded text-xs font-semibold',
                    ]"
                    >{{ event.status }}</span
                  >
                </td>
                <td class="p-4 w-1/5">
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
import { ref, computed, watch, onMounted } from "vue";
import { useQuery } from "@vue/apollo-composable";
import { useRouter, useRoute } from "vue-router";
import GetEvents from "~/graphql/query/GetEvents.gql";
import { format } from "date-fns";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";

const router = useRouter();
const route = useRoute();
const user = useAuthStore();
const events = ref([]);

// Reactive State
onMounted(async () => {
  await fetchEvents();
});

const fetchEvents = async () => {
  try {
    const { onResult, onError, refetch } = useQuery(GetEvents, {
      limit: 10,
      offset: 0,
      order_by: [{ created_at: "desc" }],
      where: {},
    });

    onResult((result) => {
      if (result.data) {
        events.value = result.data.events;
      }
    });

    onError((error) => {
      console.error("Error fetching events: ", error.message);
      toast.error("Something went wrong, try again", {
        transition: toast.TRANSITIONS.FLIP,
        position: toast.POSITION.TOP_RIGHT,
      });
    });
  } catch (error) {
    console.error("Error during fetching events: ", error);
    toast.error("Failed to load events.");
  }
};

// Pagination
const currentPage = ref(1);
const itemsPerPage = 5;

const totalPages = computed(() =>
  Math.ceil(events.value.length / itemsPerPage)
);

const paginatedEvents = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  return events.value.slice(start, start + itemsPerPage);
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

// Block Event Function
const blockEvent = (eventId) => {
  alert(`Event with ID ${eventId} has been blocked.`);
};

const formatEventDate = (date, eventTime) => {
  return format(new Date(date), "'Time: 'MMM do yyyy") + " at " + eventTime;
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
