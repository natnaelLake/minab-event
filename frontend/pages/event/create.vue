<script setup>
import { ref, computed, onMounted, onUnmounted } from "vue";
import * as yup from "yup";
import DynamicForm from "@/components/DynamicForm.vue";
import createEventMutation from "~/graphql/mutations/CreateEventMutation.gql";
import UploadImageMutation from "~/graphql/mutations/UploadImageMutation.gql";
import { toast } from "vue3-toastify";
import { useRouter } from "vue-router";
import LocationMap from "@/components/LocationMap.vue";
import GetAllCategories from "~/graphql/query/GetAllCategories.gql";
import GetAllTags from "~/graphql/query/GetAllTags.gql";
import { VueSpinner } from "vue3-spinners";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";

const router = useRouter();
const {
  mutate: createEvent,
  loading: createEventLoading,
  onDone: onCreateEventDone,
  onError: onCreateEventError,
} = useMutationComposable(createEventMutation);
const {
  mutate: uploadImages,
  loading: uploadImageLoading,
  onDone: onUploadImageDone,
  onError: onUploadImageError,
} = useMutationComposable(UploadImageMutation);

const showModal = ref(false);
const selectedLocation = ref(null);
const selectedLocationName = ref(null);
const locationText = ref("");
const searchQuery = ref("");
const selectedTags = ref([]);
const selectedTagValues = ref([]);
const dropdownOpen = ref(false);
const selectedCategory = ref(null);
const selectedCategoryValue = ref(null);
const categoryDropdownOpen = ref(false);
const categorySearchQuery = ref("");
const tagsError = ref("");
const categoryError = ref("");
const locationError = ref("");
const tagOptions = ref([]);
const categoryOptions = ref([]);

const { onResult: tagResult, refetch: refetchTags } = useQueryComposable(
  GetAllTags,
  {
    limit: 100,
    offset: 0,
    order_by: [{ created_at: "desc" }],
  }
);
const { onResult: categoryResult, refetch: refetchCategories } = useQueryComposable(
  GetAllCategories,
  {
    limit: 100,
    offset: 0,
    order_by: [{ created_at: "desc" }],
  }
);
tagResult((result) => {
  if (result.data) {
    tagOptions.value = result.data.tags;
  }
});
categoryResult((result) => {
  if (result.data) {
    categoryOptions.value = result.data.categories;
  }
});
// Computed property for filtering categories
const filteredCategoryOptions = computed(() => {
  if (!categoryOptions.value) return [];
  if (categorySearchQuery.value) {
    return categoryOptions.value.filter((category) =>
      category.name
        .toLowerCase()
        .includes(categorySearchQuery.value.toLowerCase())
    );
  }
  return categoryOptions.value;
});

// Method to toggle the category dropdown
const toggleCategoryDropdown = () => {
  categoryDropdownOpen.value = !categoryDropdownOpen.value;
};

// Method to select a category
const selectCategory = (category) => {
  selectedCategory.value = category;
  selectedCategoryValue.value = category.name;
  categoryDropdownOpen.value = false;
};

// Method to remove a category
const removeCategory = () => {
  selectedCategory.value = null;
  selectedCategoryValue.value = null;
};
const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value;
};

const toggleTag = (tag) => {
  const index = selectedTags.value.findIndex((t) => t.value === tag.name);
  if (index === -1) {
    selectedTags.value.push(tag);
    selectedTagValues.value.push(tag.name);
  } else {
    selectedTags.value.splice(index, 1);
    selectedTagValues.value.splice(
      selectedTagValues.value.indexOf(tag.name),
      1
    );
  }
};

const removeTag = (tag) => {
  selectedTags.value = selectedTags.value.filter((t) => t.value !== tag.value);
  selectedTagValues.value = selectedTagValues.value.filter(
    (v) => v !== tag.name
  );
};

const filteredTagOptions = computed(() =>
  tagOptions.value.filter((option) =>
    option.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
);
const validateForm = () => {
  locationError.value = selectedLocation.value ? "" : "Location is required";
  tagsError.value =
    selectedTagValues.value.length > 0 ? "" : "At least one tag is required";
  if (!selectedCategory.value) {
    categoryError.value = "Please select a category.";
    isValid = false;
  } else {
    categoryError.value = "";
  }
  return !(locationError.value || tagsError.value || categoryError.value);
};
const onSubmit = async (values) => {
  // Ensure all selected images are assigned and reset any previous errors
  const files = values.featured_image_url;

  if (!files || files.length === 0) {
    toast.error("No files selected", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    return;
  }

  if (!validateForm()) return; // Prevent submission if there are errors

  try {
    createEventLoading.value = true;
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
    const base64Files = await Promise.all(
      files.map((file) => readFilesAsBase64(file))
    );

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
      location: selectedLocation.value
        ? `${selectedLocation.value[0]},${selectedLocation.value[1]}`
        : values.location,
      venue: values.venue,
      is_free: values.is_free || false,
      price: parseFloat(values.price) || 0,
      quantity: parseInt(values.quantity) || 1,
      preparation_time: values.preparation_time,
      event_start_time: values.event_start_time,
      event_end_time: values.event_end_time,
      tags: `{${selectedTagValues.value.join(",")}}`,
      categories: selectedCategoryValue.value,
      featured_image_url: `{${uploadedImages.join(",")}}`,
    };
    // Use the updated mutation hook
    await createEvent(input);
    toast.success("Event created successfully!", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    router.push("/");
  } catch (error) {
    toast.error("Error creating event: " + error.message, {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

// Listen for custom event from LocationMap
const updateLocation = (event) => {
  const { locationName, location } = event.detail;
  console.log("Received location data:", locationName, location);
  selectedLocation.value = location;
  selectedLocationName.value = locationName;
  locationText.value = `Lat: ${location[0]}, Lng: ${location[1]}`;
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
      name: "event_start_time",
      label: "Event Start Time",
      placeholder: "Select event end time",
      type: "datetime-local",
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
      type: "datetime-local",
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
      as: "checkbox",
      name: "is_free",
      label: "Is this event free?",
      rules: yup.mixed().optional(),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "price",
      label: "Price",
      placeholder: "Enter price",
      type: "number",
      rules: yup
        .number()
        .min(0, "Price cannot be negative")
        .when("is_free", {
          is: true,
          then: (schema) => schema.notRequired(),
          otherwise: (schema) => schema.required("Price is required"),
        }),
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
                v-if="selectedCategory"
                class="bg-green-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
              >
                {{ selectedCategory.name }}
                <button
                  @click.stop="removeCategory"
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
                :key="option.name"
                class="px-4 py-2 flex items-center cursor-pointer hover:bg-gray-100"
                @click="selectCategory(option)"
              >
                <input
                  type="radio"
                  :value="option.name"
                  v-model="selectedCategoryValue"
                  class="mr-2"
                />
                {{ option.name }}
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
                {{ tag.name }}
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
                :key="option.name"
                class="px-4 py-2 flex items-center cursor-pointer hover:bg-gray-100"
                @click="toggleTag(option)"
              >
                <input
                  type="checkbox"
                  :value="option.name"
                  v-model="selectedTagValues"
                  class="mr-2"
                />
                {{ option.name }}
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
        class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300 flex justify-center items-center"
        :disabled="createEventLoading"
        :class="{ 'opacity-50 cursor-not-allowed': createEventLoading }"
      >
        <VueSpinner v-if="createEventLoading" size="24" color="#ffffff" />
        <span v-else>Create Event</span>
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
        <div class="p-4 border-t flex justify-end mt-20">
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
