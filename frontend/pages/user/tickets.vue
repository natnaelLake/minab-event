<script setup>
import { ref, onMounted } from "vue";
import EventCard from "~/components/EventCard.vue";
import GetMyReservedTickets from "~/graphql/query/GetReservedTicket.gql";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import { format, formatDistance, parseISO } from "date-fns";
import useQueryComposable from "~/composables/useQueryComposable";

const router = useRouter();
const user = useAuthStore();
const currentUser = user.id;
const tickets = ref([]);

onMounted(async () => {
  await fetchMyReservedEvents();
});

const fetchMyReservedEvents = async () => {
  try {
    const { onResult, onError } = useQueryComposable(GetMyReservedTickets, {
      limit: 10,
      offset: 0,
      order_by: [{ purchased_at: "desc" }],
      user_id: currentUser,
    });

    onResult((result) => {
      if (result.data) {
        tickets.value = result.data.tickets;
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
const formatEventDate = (date) => {
  if (!date) {
    console.error("Invalid date value:", date);
    return "Invalid Date"; 
  }

  try {
    const parsedDate = parseISO(date);
    return format(parsedDate, "PPpp");
  } catch (error) {
    console.error("Error parsing date:", error);
    return "Invalid Date"; 
  }
};
const goToEventDetail = (eventId) => {
  router.push(`/event/${eventId}`);
};
const handleFormatDistance = (date) => {
  return formatDistance(new Date(date), new Date(), { addSuffix: true });
};
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6 pt-28 p-10">
    <div
      v-for="ticket in tickets"
      :key="ticket.id"
      class="overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer"
    >
      <EventCard
        :header="{
          avatar: ticket.event.user.profile_image_url,
          name: ticket.event.user.first_name,
        }"
        :image="{
          src: ticket.event.featured_image_url.split(',')[0].replace('{', ''),
          alt: ticket.event.title,
        }"
        :title="ticket.event.title"
        :description="ticket.event.description"
        :address="ticket.event.venue"
        :price="`Total ${ticket.quantity} tickets by $${ticket.total_price}`"
        :deadline="formatEventDate(ticket.event.event_start_time)"
        :footer="{
          postTime: {
            text: `Purchased : ${handleFormatDistance(ticket.purchased_at)} `,
          },
        }"
        @navigate="goToEventDetail(ticket.event.id)"
      />
    </div>
  </div>
</template>
