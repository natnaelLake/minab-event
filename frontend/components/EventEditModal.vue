<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from "vue";
import * as yup from "yup";
import UpdateEvent from "~/graphql/mutations/UpdateEvent.gql";
import DeleteEvent from "~/graphql/mutations/DeleteEvent.gql";
import UploadImageMutation from "~/graphql/mutations/UploadImageMutation.gql";
import LocationMap from "~/components/LocationMap.vue";
import DynamicForm from "~/components/DynamicForm.vue";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import { useRouter } from "vue-router";
import GetAllCategories from "~/graphql/query/GetAllCategories.gql";
import GetAllTags from "~/graphql/query/GetAllTags.gql";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable"

const router = useRouter();
const user = useAuthStore();
const selectedLocation = ref(null);
const selectedLocationName = ref(null);
const locationText = ref("");
const showLocationModal = ref(false);
const images = ref([]);
const currentUser = user.id;
const currentUserRole = user.role;
const currentEventId = ref(null);
const searchQuery = ref("");
const selectedTags = ref([]);
const selectedTagValues = ref([]);
const dropdownOpen = ref(false);
const selectedCategory = ref(null);
const selectedCategoryValue = ref("");
const categoryDropdownOpen = ref(false);
const categorySearchQuery = ref("");
const tagsError = ref("");
const categoryError = ref("");
const locationError = ref("");
const tagOptions = ref([]);
const categoryOptions = ref([]);
const props = defineProps({
  show: Boolean,
  event: Object,
  mode: {
    type: String,
    default: "update",
  },
});
const emit = defineEmits(["close", "update", "delete"]);

const isDeleteMode = computed(() => props.mode === "delete");

const formData = ref({
  title: "",
  description: "",
  location: "",
  venue: "",
  price: 0,
  quantity: 0,
  preparation_time: "",
  event_start_time: "",
  event_end_time: "",
  tags: "",
  categories: "",
  featured_image_url: "",
});

const { mutate: updateEvent } = useMutationComposable(UpdateEvent);
const { mutate: deleteEvent } = useMutationComposable(DeleteEvent);
const {
  mutate: uploadImages,
  loading: uploadImageLoading,
  onDone: onUploadImageDone,
  onError: onUploadImageError,
} = useMutationComposable(UploadImageMutation);

watch(
  () => props.event,
  (newEvent) => {
    if (newEvent) {
      currentEventId.value = newEvent.id;
      // Update form data with event values
      formData.value = {
        title: newEvent.title || "",
        description: newEvent.description || "",
        location: newEvent.location || "",
        venue: newEvent.venue || "",
        price: newEvent.price || 0,
        quantity: newEvent.quantity || 0,
        preparation_time: newEvent.preparation_time || "",
        event_start_time: newEvent.event_start_time || "",
        event_end_time: newEvent.event_end_time || "",
        tags: newEvent.tags || "",
        categories: newEvent.categories || "",
        featured_image_url: "",
      };
      if (newEvent.featured_image_url) {
        images.value = newEvent.featured_image_url
          .replace(/[{}]/g, "") // Remove curly braces
          .split(",")
          .map((url) => url.trim());
      }
      if (newEvent.location) {
        const locationArr = newEvent.location.split(",");
        selectedLocation.value = [
          parseFloat(locationArr[0]),
          parseFloat(locationArr[1]),
        ];
        locationText.value = `Lat: ${selectedLocation.value[0]}, Lng: ${selectedLocation.value[1]}`;
      }
      if (newEvent.tags) {
        selectedTags.value = newEvent.tags
          .replace(/[{}]/g, "")
          .split(",")
          .map((tag) => ({
            value: tag.trim(),
            name: tag.trim(),
          }));
        selectedTagValues.value = newEvent.tags.replace(/[{}]/g, "").split(",");
      }
      if (newEvent.categories) {
        selectedCategoryValue.value = newEvent.categories;
      }
    }
  },
  { immediate: true }
);
const { onResult: tagResult, refetch: refetchTags } = useQueryComposable(GetAllTags, {
  limit: 100,
  offset: 0,
  order_by: [{ created_at: "desc" }],
},{
    context: {
      headers: {
        "x-hasura-user-id": currentUser,
        "x-hasura-role": currentUserRole,
        Authorization: `Bearer ${user.token}`,
      },
    },
  });
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
const updateLocation = (event) => {
  const { locationName, location } = event.detail;
  console.log("Received location data:", locationName, location);
  selectedLocation.value = location;
  selectedLocationName.value = locationName;
  locationText.value = `Lat: ${location[0]}, Lng: ${location[1]}`;
};
const closeModal = () => emit("close");
const openLocationModal = () => {
  showLocationModal.value = true;
};

const closeLocationModal = () => {
  showLocationModal.value = false;
};
const validateForm = () => {
  locationError.value = selectedLocation.value ? "" : "Location is required";
  tagsError.value =
    selectedTagValues.value.length > 0 ? "" : "At least one tag is required";
  categoryError.value =
    selectedCategoryValue.value.length > 0 ? "" : "Category is required";

  return !(locationError.value || tagsError.value || categoryError.value);
};
const removeImage = async (imageUrl) => {
  try {
    images.value = images.value.filter((img) => img !== imageUrl);
  } catch (error) {
    console.error("Error deleting image:", error);
    toast.error("Failed to delete image", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};
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
const onSubmit = async (values) => {
  console.log("Submitted values:", values);

  // Ensure all selected images are assigned and reset any previous errors
  const newFiles = values.featured_image_url || [];
  const existingImages = images.value ? Array.from(images.value) : [];

  if (!newFiles.length && !existingImages.length) {
    toast.error("No files selected", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    return;
  }

  if (!validateForm()) return;

  try {
    let uploadedImages = [];

    if (newFiles.length > 0) {
      const base64NewFiles = await Promise.all(
        newFiles.map((file) => readFilesAsBase64(file))
      );

      const uploadImagesInput = { input: { files: base64NewFiles } };

      const { data } = await uploadImages(uploadImagesInput);

      if (!data || !data.uploadImages || !data.uploadImages.imageUrls) {
        throw new Error("Failed to upload images");
      }

      uploadedImages = data.uploadImages.imageUrls;
    }
    const removedImages = existingImages.filter(
      (img) => !uploadedImages.map((file) => file === img)
    );
    const finalImageUrls = [
      ...existingImages.filter((img) => !removedImages.includes(img)),
      ...uploadedImages,
    ];

    const input = {
      title: values.title,
      description: values.description,
      location: selectedLocation.value
        ? `${selectedLocation.value[0]},${selectedLocation.value[1]}`
        : values.location,
      venue: values.venue,
      price: parseFloat(values.price) || 0,
      quantity: parseInt(values.quantity) || 1,
      preparation_time: values.preparation_time,
      event_start_time: values.event_start_time,
      event_end_time: values.event_end_time,
      tags: `{${selectedTagValues.value.join(",")}}`,
      categories: selectedCategoryValue.value,
      featured_image_url: `{${finalImageUrls.join(",")}}`,
    };

    await updateEvent({ id: String(currentEventId.value), changes: input });
    toast.success("Event updated successfully!", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
    router.push("/user/my-events");
    emit("update");
    closeModal();
  } catch (error) {
    console.error("Error updating event:", error);
    toast.error("Failed to update event", {
      transition: toast.TRANSITIONS.FLIP,
      position: toast.POSITION.TOP_RIGHT,
    });
  }
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
      rules: yup
        .mixed()
        .test(
          "fileSize",
          "The file is too large",
          (value) =>
            !value ||
            (Array.isArray(value) &&
              value.every((file) => file.size <= 5000000)) // Example size limit: 5MB
        )
        .test(
          "fileType",
          "Unsupported file format",
          (value) =>
            !value ||
            (Array.isArray(value) &&
              value.every((file) =>
                ["image/jpeg", "image/png", "image/gif"].includes(file.type)
              ))
        ),
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
onMounted(() => {
  window.addEventListener("location-selected", updateLocation);
});

onUnmounted(() => {
  window.removeEventListener("location-selected", updateLocation);
});
</script>

<template>
  <div
    v-if="show"
    class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
  >
    <div
      class="dark:bg-gray-200 rounded-lg mt-20 p-6 w-96 shadow-lg max-h-full overflow-y-auto"
    >
      <h3 class="text-xl font-semibold mb-4">
        {{ isDeleteMode ? "Delete Event" : "Edit Event" }}
      </h3>

      <!-- Conditional rendering for delete confirmation -->
      <div v-if="isDeleteMode">
        <p>Are you sure you want to delete this event?</p>
        <div class="flex justify-end mt-4">
          <button
            @click="closeModal"
            class="mr-3 px-4 py-2 bg-gray-300 rounded"
          >
            Cancel
          </button>
          <button
            @click="handleDelete"
            class="px-4 py-2 bg-red-500 text-white rounded"
          >
            Block
          </button>
        </div>
      </div>

      <!-- Form for editing -->
      <div v-else>
        <div
          class="bg-white flex flex-col justify-center items-center max-w-3xl rounded-xl"
        >
          <DynamicForm
            :schema="eventSchema"
            :submitHandler="onSubmit"
            :initialValues="formData"
          >
            <div v-if="images.length > 0">
              <div class="image-preview-container">
                <div
                  v-for="(img, index) in images"
                  :key="index"
                  class="image-preview"
                >
                  <img
                    :src="img"
                    alt="Image"
                    class="w-24 h-24 object-cover rounded-md"
                  />
                  <button
                    @click="removeImage(img)"
                    class="remove-image-button text-red-500"
                  >
                    &times;
                    <!-- This is the "X" character for remove -->
                  </button>
                </div>
              </div>
            </div>
            <div v-else>
              <p>No images uploaded.</p>
            </div>

            <!-- Category Selection -->
            <div class="mb-5">
              <label
                class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
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
                      {{ selectedCategoryValue }}
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
              <label
                class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
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
              <label
                class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1"
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
            <div class="flex justify-end mt-4">
              <button
                @click="closeModal"
                class="mr-3 px-4 py-2 bg-gray-300 rounded"
              >
                Cancel
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-blue-500 text-white rounded"
              >
                Save
              </button>
            </div>
          </DynamicForm>

          <!-- Location Selection Modal -->
          <div
            v-if="showLocationModal"
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
      </div>
    </div>
  </div>
</template>

<style scoped>
.image-preview-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.image-preview {
  position: relative;
  display: inline-block;
}

.remove-image-button {
  position: absolute;
  top: 5px;
  right: 5px;
  background-color: rgba(255, 255, 255, 0.8);
  border: none;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}
</style>
