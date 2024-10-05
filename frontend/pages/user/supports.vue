
<script setup>
import { ref } from "vue";
import { useForm, useField, Field, ErrorMessage, Form } from "vee-validate";
import * as yup from "yup";
import CREATE_SUPPORT_TICKET from "~/graphql/mutations/CreateSupportTicket.gql";
import GET_SUPPORT_TICKETS from "~/graphql/query/GetSupportTickets.gql";
import { toast } from "vue3-toastify";
import { formatDistance } from "date-fns";

// Define schema with yup
const schema = yup.object({
  title: yup.string().required("Title is required"),
  description: yup.string().required("Description is required"),
});

const tickets = ref([]);
// Define fields with vee-validate
const { handleSubmit } = useForm({ validationSchema: schema });
const { value: title } = useField("title");
const { value: description } = useField("description");

const { mutate: createTicket } = useMutationComposable(CREATE_SUPPORT_TICKET);
const { onResult: ticketsResult, refetch } =
  useQueryComposable(GET_SUPPORT_TICKETS);

ticketsResult((result) => {
  if (result.data) {
    tickets.value = result.data.support_tickets;
    console.log(';;;;;;;;;;;;',tickets)
  }
});

const submitTicket = handleSubmit(async (values) => {
  try {
    await createTicket({
      title: values.title,
      description: values.description,
    });
    toast.success("Support ticket created successfully!", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    title.value = "";
    description.value = "";
    refetch();
  } catch (error) {
    console.error("Error creating ticket:", error);
    toast.error("Error creating support ticket.", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
});

// Utility function to format date
// const formatDate = (date) => {
//   const options = { year: 'numeric', month: 'long', day: 'numeric' };
//   return new Date(date).toLocaleDateString(undefined, options);
// };
</script>

<template>
  <div class="p-6 max-w-6xl mx-auto bg-white shadow-md rounded-lg mt-20">
    <h1 class="text-2xl font-bold mb-6 text-gray-800">Support Tickets</h1>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Create Support Ticket Section -->
      <div class="bg-white shadow-md rounded-lg p-6">
        <h2 class="text-xl font-semibold mb-4 text-gray-800">
          Create Support Ticket
        </h2>
        <Form
          @submit="submitTicket"
          :validation-schema="schema"
          class="space-y-6"
        >
          <div class="mb-4">
            <label for="title" class="block text-sm font-medium text-gray-700"
              >Title</label
            >
            <Field
              v-model="title"
              name="title"
              type="text"
              id="title"
              class="mt-1 block w-full border h-10 border-gray-300 rounded-md sm:text-sm"
            />
            <ErrorMessage name="title" class="text-red-500 text-sm mt-1" />
          </div>

          <div class="mb-4">
            <label
              for="description"
              class="block text-sm font-medium text-gray-700"
              >Description</label
            >
            <Field
              v-model="description"
              name="description"
              as="textarea"
              id="description"
              rows="4"
              class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm sm:text-sm"
            />
            <ErrorMessage
              name="description"
              class="text-red-500 text-sm mt-1"
            />
          </div>

          <div class="flex justify-end">
            <button
              type="submit"
              class="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-blue-500 hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Submit Ticket
            </button>
          </div>
        </Form>
      </div>

      <!-- Support Ticket List Section -->
      <div class="bg-white shadow-md rounded-lg p-6">
        <h2 class="text-xl font-semibold mb-4 text-gray-800">
          Support Tickets List
        </h2>
        <div v-if="tickets.length" class="space-y-4">
          <div
            v-for="ticket in tickets"
            :key="ticket.id"
            class="p-4 border border-gray-200 rounded-lg"
          >
            <h3 class="text-lg font-semibold text-gray-800">
              {{ ticket.title }}
            </h3>
            <p class="text-gray-700">{{ ticket.description }}</p>
            <span class="text-sm text-gray-500">{{
              formatDistance(new Date(ticket.created_at), new Date(), {
                addSuffix: true,
              })
            }}</span>
          </div>
        </div>
        <div v-else class="text-gray-500">No support tickets found.</div>
      </div>
    </div>
  </div>
</template>


<style scoped>
/* Optional: Add additional styling for better visuals */
</style>
