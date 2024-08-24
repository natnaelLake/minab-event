<template>
  <div
    class="flex flex-col justify-center items-center mx-auto max-w-3xl p-8 rounded-xl mt-20"
  >
    <DynamicForm :schema="eventSchema" :submitHandler="onSubmit">
      <!-- Category Selection -->
      <div class="mb-5">
        <label class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
          >Category</label
        >
        <div class="relative">
          <div
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300 cursor-pointer"
            @click="toggleCategoryDropdown"
          >
            <div class="flex flex-wrap gap-2">
              <span
                v-if="selectedCategories.length > 0"
                v-for="category in selectedCategories"
                :key="category.value"
                class="bg-green-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
              >
                {{ category.label }}
                <button
                  @click.stop="removeCategory(category)"
                  class="ml-2 text-white hover:text-gray-200"
                >
                  ×
                </button>
              </span>
              <span v-else class="text-gray-500">Select Category</span>
            </div>
          </div>
          <div
            v-if="categoryDropdownOpen"
            class="absolute z-10 mt-2 w-full bg-white border border-gray-300 rounded-lg shadow-lg mb-10"
          >
            <div class="px-4 py-2 border-b">
              <input
                type="text"
                v-model="categorySearchQuery"
                placeholder="Search categories..."
                class="w-full px-3 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300"
              />
            </div>
            <div class="max-h-40 overflow-y-auto">
              <div
                v-for="option in filteredCategoryOptions"
                :key="option.value"
                class="px-4 py-2 flex items-center cursor-pointer hover:bg-gray-100"
                @click="toggleCategory(option)"
              >
                <input
                  type="checkbox"
                  :value="option.value"
                  v-model="selectedCategoryValues"
                  class="mr-2"
                />
                {{ option.label }}
              </div>
            </div>
          </div>
        </div>
        <p v-if="categoryError" class="text-red-500 text-sm mt-2">
          {{ categoryError }}
        </p>
      </div>

      <!-- Tags Section -->
      <div class="mb-5">
        <label class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
          >Tags</label
        >
        <div class="relative">
          <div
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300 cursor-pointer"
            @click="toggleDropdown"
          >
            <div class="flex flex-wrap gap-2">
              <span
                v-if="selectedTags.length > 0"
                v-for="tag in selectedTags"
                :key="tag.value"
                class="bg-blue-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
              >
                {{ tag.label }}
                <button
                  @click.stop="removeTag(tag)"
                  class="ml-2 text-white hover:text-gray-200"
                >
                  ×
                </button>
              </span>
              <span v-else class="text-gray-500">Select Tags</span>
            </div>
          </div>
          <div
            v-if="dropdownOpen"
            class="absolute z-10 mt-2 w-full bg-white border border-gray-300 rounded-lg shadow-lg mb-10"
          >
            <div class="px-4 py-2 border-b">
              <input
                type="text"
                v-model="searchQuery"
                placeholder="Search tags..."
                class="w-full px-3 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300"
              />
            </div>
            <div class="max-h-40 overflow-y-auto">
              <div
                v-for="option in filteredTagOptions"
                :key="option.value"
                class="px-4 py-2 flex items-center cursor-pointer hover:bg-gray-100"
                @click="toggleTag(option)"
              >
                <input
                  type="checkbox"
                  :value="option.value"
                  v-model="selectedTagValues"
                  class="mr-2"
                />
                {{ option.label }}
              </div>
            </div>
          </div>
        </div>
        <p v-if="tagsError" class="text-red-500 text-sm mt-2">
          {{ tagsError }}
        </p>
      </div>

      <!-- Location Selection Button -->
      <div class="mb-5">
        <label class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
          >Event Location</label
        >
        <div class="relative">
          <input
            type="text"
            v-model="locationText"
            placeholder="Click to select location"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300 cursor-pointer"
            readonly
            @click="openLocationModal"
          />
        </div>
        <p v-if="locationError" class="text-red-500 text-sm mt-2">
          {{ locationError }}
        </p>
      </div>

      <!-- Modal Trigger -->
      <button
        type="submit"
        class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300"
        :disabled="createEventLoading"
        :class="{ 'opacity-50 cursor-not-allowed': createEventLoading }"
      >
        {{ createEventLoading ? "Creating Event..." : "Create Event" }}
      </button>
    </DynamicForm>

    <!-- Location Selection Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 z-50 flex items-center justify-center bg-gray-800 bg-opacity-75"
    >
      <div class="bg-white rounded-lg shadow-lg w-full max-w-3xl">
        <div class="p-4 flex justify-between items-center border-b">
          <h3 class="text-xl font-semibold">Select Event Location</h3>
          <button
            @click="closeLocationModal"
            class="text-gray-600 hover:text-gray-800"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
        <div class="p-4">
          <LocationMap @location-selected="updateLocation" />
        </div>
        <div class="p-4 border-t flex justify-end">
          <div class="p-4">
            {{ selectedLocationName }}
          </div>
          <button
            @click="closeLocationModal"
            class="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-lg shadow-md"
          >
            Confirm Location
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from "vue";
import { useForm } from "vee-validate";
import * as yup from "yup";
import DynamicForm from "@/components/DynamicForm.vue";
import createEventMutation from "~/graphql/mutations/CreateEventMutation.gql";
import UploadImageMutation from "~/graphql/mutations/UploadImageMutation.gql";
import { useAuthStore } from "~/store";
import { useMutation } from "@vue/apollo-composable";
import { toast } from "vue3-toastify";
import { useRouter } from "vue-router";
import LocationMap from "@/components/LocationMap.vue";

const authStore = useAuthStore();
const router = useRouter();
const { mutate: createEvent, loading: createEventLoading, onDone: onCreateEventDone, onError: onCreateEventError } = useMutation(createEventMutation);
const { mutate: uploadImages, loading: uploadImageLoading, onDone: onUploadImageDone, onError: onUploadImageError } = useMutation(UploadImageMutation);


const showModal = ref(false);
const selectedLocation = ref(null);
const selectedLocationName = ref(null);
const locationText = ref("");
const searchQuery = ref("");
const selectedTags = ref([]);
const selectedTagValues = ref([]);
const dropdownOpen = ref(false);
const selectedCategories = ref([]);
const selectedCategoryValues = ref([]);
const categoryDropdownOpen = ref(false);
const categorySearchQuery = ref("");
const tagsError = ref("");
const categoryError = ref("");
const locationError = ref("");
const image = ref(null);
const profileImage = (ref < File) | (null > null);

const tagOptions = [
  { value: "tech", label: "Tech" },
  { value: "music", label: "Music" },
  { value: "art", label: "Art" },
  // Add more tags as needed
];

const categoryOptions = [
  { value: "tech", label: "Tech" },
  { value: "music", label: "Music" },
  { value: "art", label: "Art" },
  // Add more categories as needed
];

const toggleCategoryDropdown = () => {
  categoryDropdownOpen.value = !categoryDropdownOpen.value;
};

const toggleCategory = (category) => {
  const index = selectedCategories.value.findIndex(
    (c) => c.value === category.value
  );
  if (index === -1) {
    selectedCategories.value.push(category);
    selectedCategoryValues.value.push(category.value);
  } else {
    selectedCategories.value.splice(index, 1);
    selectedCategoryValues.value.splice(
      selectedCategoryValues.value.indexOf(category.value),
      1
    );
  }
};

const removeCategory = (category) => {
  selectedCategories.value = selectedCategories.value.filter(
    (c) => c.value !== category.value
  );
  selectedCategoryValues.value = selectedCategoryValues.value.filter(
    (v) => v !== category.value
  );
};

const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

const toggleTag = (tag) => {
  const index = selectedTags.value.findIndex((t) => t.value === tag.value);
  if (index === -1) {
    selectedTags.value.push(tag);
    selectedTagValues.value.push(tag.value);
  } else {
    selectedTags.value.splice(index, 1);
    selectedTagValues.value.splice(
      selectedTagValues.value.indexOf(tag.value),
      1
    );
  }
};

const removeTag = (tag) => {
  selectedTags.value = selectedTags.value.filter((t) => t.value !== tag.value);
  selectedTagValues.value = selectedTagValues.value.filter(
    (v) => v !== tag.value
  );
};

const filteredTagOptions = computed(() =>
  tagOptions.filter((option) =>
    option.label.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
);

const filteredCategoryOptions = computed(() =>
  categoryOptions.filter((option) =>
    option.label.toLowerCase().includes(categorySearchQuery.value.toLowerCase())
  )
);

const validateForm = () => {
  locationError.value = selectedLocation.value ? "" : "Location is required";
  tagsError.value =
    selectedTagValues.value.length > 0 ? "" : "At least one tag is required";
  categoryError.value =
    selectedCategoryValues.value.length > 0 ? "" : "Category is required";

  return !(locationError.value || tagsError.value || categoryError.value);
};
const onSubmit = async (values) => {
  // Ensure all selected images are assigned and reset any previous errors
  const files = values.featured_image_url;
  console.log('[[[[[[[[[[[[[[[]]]]]]]]]]]]]]]', files);
  
  if (!files || files.length === 0) {
    toast.error("No files selected", { transition: toast.TRANSITIONS.FLIP, position: toast.POSITION.TOP_RIGHT });
    return;
  }

  if (!validateForm()) return; // Prevent submission if there are errors

  try {
    const readFilesAsBase64 = (file) => {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onloadend = () => {
          if (reader.result) {
            resolve(reader.result.split(",")[1]); // Return only the base64 part
          } else {
            reject(new Error("Failed to read file"));
          }
        };
        reader.readAsDataURL(file);
      });
    };

    // Convert all files to base64 in parallel
    const base64Files = await Promise.all(files.map(file => readFilesAsBase64(file)));

    // Prepare the input object as per mutation requirement
    const uploadImagesInput = { input: { files: base64Files } };

    // Perform the mutation to upload images
    const { data } = await uploadImages(uploadImagesInput);

    if (!data || !data.uploadImages || !data.uploadImages.imageUrls) {
      throw new Error("Failed to upload images");
    }

    const uploadedImages = data.uploadImages.imageUrls;

    // Prepare input for event creation
    const input = {
      title: values.title,
      description: values.description,
      location: selectedLocation.value ? `${selectedLocation.value[0]},${selectedLocation.value[1]}` : values.location,
      venue: values.venue,
      price: parseFloat(values.price) || 0,
      quantity: parseInt(values.quantity) || 1,
      preparation_time: values.preparation_time,
      event_date: values.event_date,
      event_start_time: values.event_start_time,
      event_end_time: values.event_end_time,
      tags: `{${selectedTagValues.value.join(",")}}`, 
      categories: `{${selectedCategoryValues.value.join(",")}}`, 
      featured_image_url: `{${uploadedImages.join(",")}}`, 
      user_id: authStore.user.id,
    };

    // Use the updated mutation hook
    await createEvent(input);
    toast.success("Event created successfully!", { transition: toast.TRANSITIONS.FLIP, position: toast.POSITION.TOP_RIGHT });
    router.push("/");
  } catch (error) {
    toast.error("Error creating event: " + error.message, { transition: toast.TRANSITIONS.FLIP, position: toast.POSITION.TOP_RIGHT });
  }
};



// Listen for custom event from LocationMap
const updateLocation = (event) => {
  const { locationName, location } = event.detail;
  console.log("Received location data:", locationName, location);
  selectedLocation.value = location;
  selectedLocationName.value = locationName;
  locationText.value = `Lat: ${location[0]}, Lng: ${location[1]}`;
  // closeLocationModal();
};

const openLocationModal = () => {
  showModal.value = true;
};

const closeLocationModal = () => {
  showModal.value = false;
};
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
        label: "text-sm font-medium text-gray-800 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
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
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
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
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "event_start_time",
      label: "Event Start Time",
      placeholder: "Select event end time",
      type: "time",
      rules: yup.string().required("Event start time is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "event_end_time",
      label: "Event End Time",
      placeholder: "Select event end time",
      type: "time",
      rules: yup.string().required("Event end time is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
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
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "price",
      label: "Price",
      placeholder: "Enter price (or leave blank for free event)",
      type: "number",
      rules: yup
        .number()
        .required("Price is required field")
        .min(0, "Price cannot be negative"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "quantity",
      label: "Total Capacity",
      placeholder: "Enter Total Holding capacity",
      type: "number",
      rules: yup
        .number()
        .required("Total Capacity is required field")
        .min(1, "Total Capacity cannot be  less than 1"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
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
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },

    {
      as: "input",
      name: "featured_image_url",
      label: "Featured Image",
      placeholder: "Upload featured image",
      type: "file",
      multiple: true,
      rules: yup.mixed().required("Featured image is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
  ],
};
// Set up event listener when component is mounted
onMounted(() => {
  window.addEventListener("location-selected", updateLocation);
});

onUnmounted(() => {
  window.removeEventListener("location-selected", updateLocation);
});
</script>
<style scoped>
/* Add any custom styling for Multiselect or other elements */
.w-full {
  max-width: 100%;
}

.cursor-pointer {
  cursor: pointer;
}

.border {
  border: 1px solid #ddd;
}

.max-h-60 {
  max-height: 15rem; /* Adjust as needed */
}

.overflow-y-auto {
  overflow-y: auto;
}
</style>
