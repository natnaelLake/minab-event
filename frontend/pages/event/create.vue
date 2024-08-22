<template>
  <div class="max-w-3xl mx-auto bg-white dark:bg-gray-800 p-8 rounded-xl shadow-lg mt-10">
    <DynamicForm :schema="eventSchema" :submitHandler="onSubmit">
      <button
        type="submit"
        class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300"
        :disabled="loading"
        :class="{ 'opacity-50 cursor-not-allowed': loading }"
      >
        {{ loading ? "Creating Event..." : "Create Event" }}
      </button>
    </DynamicForm>
  </div>
</template>

<script setup>
import { useForm } from "vee-validate";
import * as yup from "yup";
import DynamicForm from "@/components/DynamicForm.vue";
import createEventMutation from "~/graphql/mutations/CreateEventMutation.gql";
import { useAuthStore } from "~/store";
import { useMutation } from '@vue/apollo-composable';
import { toast } from "vue3-toastify";
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const { mutate, loading, onDone, onError } = useMutation(createEventMutation);

const eventSchema = {
  fields: [
    {
      as: "input",
      name: "title",
      label: "Event Title",
      placeholder: "Enter event title",
      type: "text",
      rules: yup.string().required("Event title is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "textarea",
      name: "description",
      label: "Event Description",
      placeholder: "Enter event description",
      rules: yup.string().required("Description is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "event_date",
      label: "Event Date",
      placeholder: "Select event date",
      type: "date",
      rules: yup.date().required("Event date is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "location",
      label: "Event Location",
      placeholder: "Enter location",
      type: "text",
      rules: yup.string().required("Location is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "venue",
      label: "Venue",
      placeholder: "Enter venue",
      type: "text",
      rules: yup.string().required("Venue is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "address",
      label: "Address",
      placeholder: "Enter address",
      type: "text",
      rules: yup.string().required("Address is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "price",
      label: "Price",
      placeholder: "Enter price (or leave blank for free event)",
      type: "number",
      rules: yup.number().min(0, "Price cannot be negative"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "preparation_time",
      label: "Preparation Time",
      placeholder: "Enter preparation time",
      type: "text",
      rules: yup.string().required("Preparation time is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "tags",
      label: "Tags",
      placeholder: "Enter tags separated by commas",
      type: "text",
      rules: yup.string().required("At least one tag is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "featured_image_url",
      label: "Featured Image",
      placeholder: "Upload featured image",
      type: "file",
      rules: yup.mixed().required("Featured image is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input: "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
  ],
};

const onSubmit = async (values) => {
  const input = {
    title: values.title,
    description: values.description,
    location: values.location,
    venue: values.venue,
    address: values.address,
    price: parseFloat(values.price) || 0,
    preparation_time: values.preparation_time, // Assuming this is an ISO 8601 interval string
    event_date: values.event_date,
    tags: `{${values.tags.split(',').map(tag => `"${tag.trim()}"`).join(',')}}`, // Formatting as a PostgreSQL array
    featured_image_url: "values.featured_image_url[0]", // Assuming single file upload
    user_id: authStore.user.id, // Assuming user ID is stored in the auth store
  };

  mutate(input);

  onDone((result) => {
    toast.success("Event created successfully!", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    router.push('/');
  });

  onError((error) => {
    toast.error("Error creating event: " + error.message, {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  });
};
</script>
