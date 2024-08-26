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
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">Date</label>
      <input
        v-model="dateRange"
        type="date"
        class="w-full p-3 border rounded"
      />
    </div>

    <!-- Category -->
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">Category</label>
      <select v-model="selectedCategory" class="w-full p-3 border rounded">
        <option value="">All Categories</option>
        <option
          v-for="category in categories"
          :key="category"
          :value="category"
        >
          {{ category }}
        </option>
      </select>
    </div>

    <!-- Tags -->
    <div class="mb-4">
      <label class="block text-sm font-medium mb-2">Tags</label>
      <select v-model="selectedTags" class="w-full p-3 border rounded">
        <option value="">All Tags</option>
        <option v-for="tag in tags" :key="tag" :value="tag">
          {{ tag }}
        </option>
      </select>
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

// Filter fields
const venue = ref("");
const dateRange = ref("");
const selectedCategory = ref("");
const selectedTags = ref([]);
const minPriceRange = ref(0);
const maxPriceRange = ref(1000);
const preparationTime = ref("");
const mapLocation = ref({ lat: null, lng: null });

// Categories and Tags arrays
const categories = ["Music", "Art", "Social", "Education", "Sports"];
const tags = ["Rock", "Classical", "Pop", "Tech", "DIY", "Workshop"];

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
  console.log("mmmmmm", mapLocation.value);
  if (venue.value) {
    where.venue = { _ilike: `%${venue.value}%` };
  }
  if (dateRange.value) {
    where.date = { _eq: dateRange.value };
  }
  if (selectedCategory.value) {
    where.category = { _eq: selectedCategory.value };
  }
  if (selectedTags.length > 0) {
    where.tags = { _has_key_any: selectedTags };
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
      _st_d_within: {
        from: {
          type: "Point",
          coordinates: [mapLocation.value.lng, mapLocation.value.lat],
        },
        distance: 5000, // Example distance in meters
      },
    };
  }

  return where;
};

const applyFilters = () => {
  emitFilterData();
};

const resetFilters = () => {
  venue.value = "";
  dateRange.value = "";
  selectedCategory.value = "";
  selectedTags.value = [];
  minPriceRange.value = 0;
  maxPriceRange.value = 1000;
  preparationTime.value = "";
  mapLocation.value = { lat: null, lng: null };
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
