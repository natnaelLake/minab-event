<template>
  <div class="p-4 bg-white shadow-md rounded-lg">
    <h2 class="text-lg font-semibold mb-4">Filter Events</h2>

    <!-- Venue with Map Integration -->
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">Venue</label>
      <div class="relative">
        <input
          v-model="venue"
          type="text"
          placeholder="Enter venue"
          class="w-full p-3 border rounded"
        />
        <!-- Map Component Integration -->
        <div class="mt-2">
          <MapComponent @location-selected="handleLocationSelected" />
        </div>
      </div>
    </div>

    <!-- Date -->
    <div class="mb-4 mt-20">
      <label class="block text-sm font-medium mb-2">Start Date</label>
      <input
        v-model="start_date"
        type="datetime-local"
        class="w-full p-3 border rounded"
      />
    </div>
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">End Date</label>
      <input
        v-model="end_date"
        type="datetime-local"
        class="w-full p-3 border rounded"
      />
    </div>

    <!-- Category -->
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

    <!-- Price Range -->
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">Price Range</label>
      <div class="flex justify-between text-sm mb-2">
        <span>Min: ${{ minPrice }}</span>
        <span>Max: ${{ maxPrice }}</span>
      </div>
      <div class="flex space-x-2">
        <input
          v-model.number="minPriceRange"
          type="number"
          min="0"
          max="1000"
          step="10"
          class="w-full p-2 border rounded"
          placeholder="Min Price"
        />
        <input
          v-model.number="maxPriceRange"
          type="number"
          min="0"
          max="1000"
          step="10"
          class="w-full p-2 border rounded"
          placeholder="Max Price"
        />
      </div>
    </div>

    <!-- Preparation Time -->
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">Preparation Time</label>
      <input
        v-model="preparationTime"
        type="text"
        placeholder="Enter preparation time"
        class="w-full p-3 border rounded"
      />
    </div>

    <!-- Buttons -->
    <div class="flex space-x-4">
      <button
        @click="applyFilters"
        class="bg-gradient-to-r from-blue-400 to-blue-600 text-white px-6 py-3 rounded-full shadow-lg hover:from-blue-500 hover:to-blue-700 transition-transform transform hover:scale-105 focus:outline-none focus:ring-4 focus:ring-blue-300"
      >
        Apply
      </button>
      <button
        @click="resetFilters"
        class="bg-gradient-to-r from-gray-400 to-gray-600 text-white px-6 py-3 rounded-full shadow-lg hover:from-gray-500 hover:to-gray-700 transition-transform transform hover:scale-105 focus:outline-none focus:ring-4 focus:ring-gray-300"
      >
        Reset
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from "vue";
import MapComponent from "@/components/LocationMap.vue";
import GetAllCategories from "~/graphql/query/GetAllCategories.gql";
import GetAllTags from "~/graphql/query/GetAllTags.gql";

// Filter fields
const venue = ref("");
const start_date = ref("");
const end_date = ref("");
const minPriceRange = ref(0);
const maxPriceRange = ref(1000);
const preparationTime = ref("");
const mapLocation = ref({ lat: null, lng: null });

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

const tagOptions = ref([]);

const categoryOptions = ref([]);

const { onResult: tagResult, refetch: refetchTags } = useQuery(GetAllTags, {
  limit: 100,
  offset: 0,
  order_by: [{ created_at: "desc" }],
});
const { onResult: categoryResult, refetch: refetchCategories } = useQuery(
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
console.log("+++++++++++++++++++++++----------", categoryOptions);
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
// Compute min and max prices
const minPrice = computed(() => minPriceRange.value);
const maxPrice = computed(() => maxPriceRange.value);

// Watch for changes to ensure min and max values are valid
watch([minPriceRange, maxPriceRange], () => {
  if (minPriceRange.value > maxPriceRange.value) {
    minPriceRange.value = maxPriceRange.value;
  }
  if (maxPriceRange.value < minPriceRange.value) {
    maxPriceRange.value = minPriceRange.value;
  }
});

const handleLocationSelected = (event) => {
  mapLocation.value = event.detail.location;
};

const emitFilterData = () => {
  const whereClause = generateWhereClause();
  console.log("Filter whereClause: ", whereClause);
  const event = new CustomEvent("apply-filters", { detail: whereClause });
  window.dispatchEvent(event);
};

const generateWhereClause = () => {
  const where = {};
  console.log("mmmmmm", selectedTagValues.value);
  if (venue.value) {
    where.venue = { _ilike: `%${venue.value}%` };
  }
  if (start_date.value || end_date.value) {
    where.event_start_time = {};
    if (start_date.value) {
      where.event_start_time._gte = start_date.value;
    }
    if (end_date.value) {
      where.event_end_time._lte = end_date.value;
    }
  }
  if (selectedCategoryValue.value) {
    where.categories = { _eq: selectedCategoryValue.value };
  }
  if (selectedTagValues.value.length > 0) {
    where.tags = { _in: `{${selectedTagValues.value}}` };
  }
  if (minPrice.value || maxPrice.value) {
    where.price = {};
    if (minPrice.value) {
      where.price._gte = minPrice.value;
    }
    if (maxPrice.value) {
      where.price._lte = maxPrice.value;
    }
  }
  if (preparationTime.value) {
    where.preparation_time = { _eq: preparationTime.value };
  }
  if (mapLocation.value.lat && mapLocation.value.lng) {
    where.location = {
      _ilike: `%${(mapLocation.value.lat, mapLocation.value.lng)}%`,
    };
  }

  return where;
};

const applyFilters = () => {
  emitFilterData();
};

const resetFilters = () => {
  venue.value = "";
  start_date.value = "";
  end_date.value = "";
  selectedCategoryValue.value = "";
  selectedTagValues.value = [];
  minPriceRange.value = 0;
  maxPriceRange.value = 1000;
  preparationTime.value = "";
  mapLocation.value = { lat: null, lng: null };
  emitFilterData();
};

onMounted(() => {
  window.addEventListener("location-selected", handleLocationSelected);
});

onUnmounted(() => {
  window.removeEventListener("location-selected", handleLocationSelected);
});
</script>

<style scoped>
/* Optional: Customize select styles */
select[multiple] {
  height: auto;
  min-height: 100px;
}
button:focus {
  box-shadow: 0 0 0 3px rgba(29, 78, 216, 0.5); /* Blue focus ring for Apply Filters */
}

button:focus[aria-label="reset"] {
  box-shadow: 0 0 0 3px rgba(75, 85, 99, 0.5); /* Gray focus ring for Reset Filters */
}
</style>
