<template>
  <div class="p-6 max-w-6xl mx-auto bg-white shadow-md rounded-lg mt-20">
    <h1 class="text-2xl font-bold mb-6 text-gray-800">Feedback & Support</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- Create Feedback Section -->
      <div class="bg-white shadow-md rounded-lg p-6">
        <h2 class="text-xl font-semibold mb-4 text-gray-800">
          Submit Feedback
        </h2>
        <Form
          @submit="submitFeedback"
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
              class="mt-1 h-10 block w-full border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
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
              class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
            />
            <ErrorMessage
              name="description"
              class="text-red-500 text-sm mt-1"
            />
          </div>

          <div class="flex justify-end">
            <button
              type="submit"
              class="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-green-500 hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            >
              Submit Feedback
            </button>
          </div>
        </Form>
      </div>

      <!-- Feedback List Section -->
      <div class="bg-white shadow-md rounded-lg p-6">
        <h2 class="text-xl font-semibold mb-4 text-gray-800">My Feedbacks</h2>
        <div v-if="feedbacks.length" class="space-y-4">
          <div
            v-for="feedback in feedbacks"
            :key="feedback.id"
            class="p-4 border border-gray-200 rounded-lg"
          >
            <h3 class="text-lg font-semibold text-gray-800">
              {{ feedback.title }}
            </h3>
            <p class="text-gray-700">{{ feedback.description }}</p>
            <span class="text-sm text-gray-500">{{ feedback.date }}</span>
          </div>
        </div>
        <div v-else class="text-gray-500">No feedbacks found.</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useForm, useField, Field, ErrorMessage, Form } from "vee-validate";
import * as yup from "yup";
import CREATE_FEEDBACK from "~/graphql/mutations/CreateFeedback.gql";
import GET_USER_FEEDBACKS from "~/graphql/query/GetFeedbackList.gql";
import useMutationComposable from "~/composables/useMutationComposable";
import { toast } from "vue3-toastify";

// Define schema with yup
const schema = yup.object({
  title: yup.string().required("Title is required"),
  description: yup.string().required("Description is required"),
});

// Define fields with vee-validate
const { handleSubmit } = useForm({ validationSchema: schema });
const { value: title } = useField("title");
const { value: description } = useField("description");
const feedbacks = ref([]);

const { mutate: createFeedback } = useMutationComposable(CREATE_FEEDBACK);
const { onResult: updatedFeedbackResult, refetch } =
  useQueryComposable(GET_USER_FEEDBACKS);
updatedFeedbackResult((result) => {
  if (result.data) {
    feedbacks.value = result.data.feedbacks;
    console.log('>>>>>>>>>>>>>>>>>',feedbacks)
  }
});

const submitFeedback = handleSubmit(async (values) => {
  try {
    await createFeedback({
      title: values.title,
      description: values.description,
    });
    toast.success("Feedback submitted successfully!", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    title.value = "";
    description.value = "";
    // Refresh feedback list after submission
    refetch();
  } catch (error) {
    console.error("Error submitting feedback:", error);
    toast.success("Error submitting feedback.", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
});
</script>

<style scoped>
/* Optional: Add additional styling for better visuals */
</style>
