<template>
    <div class="p-4 bg-white shadow-md rounded-lg">
      <h2 class="text-lg font-semibold mb-4">Filter Events</h2>
      
      <!-- Search -->
      <!-- <div class="mb-4">
        <input 
          v-model="searchQuery" 
          type="text" 
          placeholder="Search by title, description, or tags" 
          class="w-full p-3 border rounded" 
        />
      </div>
       -->
      <!-- Location -->
      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">Location</label>
        <div class="relative">
          <input 
            v-model="location" 
            type="text" 
            placeholder="Enter location" 
            class="w-full p-3 border rounded"
          />
          <!-- Map Component Integration (if any) -->
        </div>
      </div>
      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">Location</label>
        <div class="relative">
          <input 
            v-model="location" 
            type="text" 
            placeholder="Enter location" 
            class="w-full p-3 border rounded"
          />
          <!-- Map Component Integration (if any) -->
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
        <select 
          v-model="selectedCategory" 
          class="w-full p-3 border rounded"
        >
          <option value="">All Categories</option>
          <option v-for="category in categories" :key="category" :value="category">
            {{ category }}
          </option>
        </select>
      </div>
  
      <!-- Price -->
      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">Price</label>
        <input 
          v-model="priceRange" 
          type="range" 
          min="0" 
          max="1000" 
          step="10" 
          class="w-full"
        />
        <div class="flex justify-between text-sm mt-2">
          <span>$0</span>
          <span>$1000</span>
        </div>
      </div>
  
      <!-- Apply Button -->
      <button 
        @click="applyFilters" 
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
      >
        Apply Filters
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const searchQuery = ref('')
  const location = ref('')
  const dateRange = ref('')
  const selectedCategory = ref('')
  const priceRange = ref('')
  
  const categories = ['Music', 'Art', 'Social', 'Education', 'Sports']
  
  const applyFilters = () => {
    // Emit or process the filter data
    console.log({
      searchQuery: searchQuery.value,
      location: location.value,
      dateRange: dateRange.value,
      category: selectedCategory.value,
      priceRange: priceRange.value,
    })
  }
  </script>
  
  <style scoped>
  /* Optional: Customize filter section */
  </style>
  