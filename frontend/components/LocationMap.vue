
<script setup>
import { L } from "leaflet";
import { ref } from "vue";
import axios from "axios";

const zoom = ref(6);
const center = ref([9.030651, 38.740494]);
const selectedLocation = ref(null);
const locationName = ref("");
const searchQuery = ref("");

// Emit event with location data
const emitLocationData = () => {
  const event = new CustomEvent("location-selected", {
    detail: {
      locationName: locationName.value,
      location: selectedLocation.value,
    },
  });
  window.dispatchEvent(event);
};

const onMapClick = (event) => {
  selectedLocation.value = [event.latlng.lat, event.latlng.lng];
  fetchPlaceName(event.latlng.lat, event.latlng.lng);
};

const searchLocation = async () => {
  try {
    const response = await axios.get(
      "https://nominatim.openstreetmap.org/search",
      {
        params: {
          q: searchQuery.value,
          format: "json",
          limit: 1,
        },
      }
    );

    if (response.data.length > 0) {
      const location = response.data[0];
      selectedLocation.value = [location.lat, location.lon];
      locationName.value = location.display_name;
      center.value = [location.lat, location.lon];
    } else {
      console.error("No results found for the search query");
    }
  } catch (error) {
    console.error("Failed to search location: ", error);
  } finally {
    emitLocationData();
  }
};

const fetchPlaceName = async (lat, lon) => {
  try {
    const response = await axios.get(
      "https://nominatim.openstreetmap.org/reverse",
      {
        params: {
          lat: lat,
          lon: lon,
          format: "json",
        },
      }
    );

    locationName.value = response.data.display_name;
    emitLocationData();
  } catch (error) {
    console.error("Failed to fetch place name: ", error);
  }finally{
    emitLocationData();
  }
};
</script>

<template>
  <div class="h-96 w-full">
    <div class="p-4">
      <input
        type="text"
        v-model="searchQuery"
        @keyup.enter="searchLocation"
        placeholder="Search for a location..."
        class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300"
      />
    </div>
    <LMap ref="map" :zoom="zoom" :center="center" @click="onMapClick">
      <LTileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&amp;copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors'
      />
      <LMarker v-if="selectedLocation" :lat-lng="selectedLocation" />
      <LMarker v-else :lat-lng="center" />
    </LMap>
    <!-- <p v-if="locationName">Selected Location: {{ locationName }}</p> -->
  </div>
</template>
