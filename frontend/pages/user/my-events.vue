<template>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6 pt-28 p-10">
    <div
      v-for="event in events"
      :key="event.id"
      class="text-black overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer"
    >
      <EventCard
        :image="{
          src: event.featured_image_url.split(',')[0].replace('{', ''), // Use the first image and remove '{'
          alt: event.title,
        }"
        :title="event.title"
        :description="event.description"
        :footer="{
          postTime: {
            text: formatEventDate(event.event_date, event.event_start_time),
          },
        }"
        :actions="[
          {
            label: 'Edit',
            handler: () => openEditModal(event), // Trigger Edit Modal
            class: 'text-blue-500 hover:underline',
          },
          {
            label: 'Delete',
            handler: () => openDeleteModal(event.id),
            class: 'text-red-500 hover:underline ml-4',
          },
        ]"
        @navigate="goToEventDetail(event.id)"
        class="dark:bg-white text-black"
      />
    </div>
  </div>

  <!-- Edit Modal -->
  <EventEditModal
    v-if="editModal.show"
    :show="editModal.show"
    :event="editModal.event"
    @close="closeEditModal"
  />

  <!-- Delete Confirmation Modal -->
  <div
    v-if="deleteModal.show"
    class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
  >
    <div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full">
      <h2 class="text-lg font-semibold mb-4">Confirm Delete</h2>
      <p class="-mt-4 mb-12 text-black">
        Are you sure you want to delete this event?
      </p>
      <div class="flex justify-end space-x-4">
        <button
          @click="handleDeleteEvent(deleteModal.eventId)"
          class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition duration-300"
        >
          Delete
        </button>
        <button
          @click="closeDeleteModal"
          class="bg-gray-300 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-400 transition duration-300"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import EventCard from "~/components/EventCard.vue";
import EventEditModal from "~/components/EventEditModal.vue";
import GetMyEvents from "~/graphql/query/GetMyEvents.gql";
import DeleteEvent from "~/graphql/mutations/DeleteEvent.gql";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import { format } from "date-fns";

const router = useRouter();
const user = useAuthStore();
const currentUser = user.id;
const events = ref([]);
const editModal = ref({ show: false, event: null });
const deleteModal = ref({ show: false, eventId: null }); // State for delete modal

onMounted(async () => {
  await fetchMyEvents();
});

const fetchMyEvents = async () => {
  try {
    const { onResult, onError } = useQuery(GetMyEvents, {
      limit: 10,
      offset: 0,
      order_by: [{ created_at: "desc" }],
      user_id: currentUser,
    });

    onResult((result) => {
      if (result.data) {
        events.value = result.data.events;
      }
    });

    onError((error) => {
      console.error("Error fetching events: ", error.message);
      toast.error("Something went wrong, try again");
    });
  } catch (error) {
    console.error("Error during fetching events: ", error);
    toast.error("Failed to load events.");
  }
};

const openEditModal = (event) => {
  editModal.value = { show: true, event };
};

const closeEditModal = () => {
  editModal.value.show = false;
};

// const handleUpdateEvent = (updatedEvent) => {
//   const index = events.value.findIndex((e) => e.id === updatedEvent.id);
//   if (index !== -1) {
//     events.value[index] = updatedEvent;
//     toast.success("Event updated successfully");
//   }
//   closeEditModal();
// };

const openDeleteModal = (eventId) => {
  deleteModal.value = { show: true, eventId };
};

const closeDeleteModal = () => {
  deleteModal.value.show = false;
};

const handleDeleteEvent = async (eventId) => {
  try {
    await deleteEvent({ event_id: eventId }); // Use event_id as per your mutation variable
    closeDeleteModal();
  } catch (error) {
    console.error("Error deleting event: ", error.message);
    toast.error("Failed to delete the event.");
  }
};

const { mutate: deleteEvent } = useMutation(DeleteEvent, {
  variables: { event_id: "" }, // Set default or placeholder value for event_id
  update(cache, { data }) {
    if (data.delete_events.affected_rows > 0) { // Check if the mutation was successful
      events.value = events.value.filter(
        (event) => event.id !== deleteModal.value.eventId
      );
      toast.success("Event deleted successfully");
    } else {
      toast.error("Failed to delete the event.");
    }
  },
  onError(error) {
    console.error("Error deleting event: ", error.message);
    toast.error("Failed to delete the event.");
  },
});

const goToEventDetail = (eventId) => {
  router.push(`/event/${eventId}`);
};

const formatEventDate = (date, eventTime) => {
  return format(new Date(date), "'Time: 'MMM do yyyy") + " at " + eventTime;
};
</script>
