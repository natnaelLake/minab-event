<template>
  <div class="h-96 w-full">
    <LMap ref="map" :zoom="zoom" :center="center" @click="onMapClick">
      <LTileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&amp;copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors'
      />
      <LMarker v-if="selectedLocation" :lat-lng="selectedLocation" />
      <LMarker v-else :lat-lng="center" />
    </LMap>
    <p v-if="locationName">Selected Location: {{ locationName }}</p>
  </div>
</template>

<script setup>
import { L } from "leaflet";
import { ref } from "vue";
import axios from "axios";

const zoom = ref(6);
const center = ref([9.030651, 38.740494]);
const selectedLocation = ref(null);
const locationName = ref("");

// Emit event with location data
const emitLocationData = () => {
  const event = new CustomEvent('location-selected', {
    detail: {
      locationName: locationName.value,
      location: selectedLocation.value,
    },
  });
  window.dispatchEvent(event);
};

const onMapClick = async (event) => {
  selectedLocation.value = [event.latlng.lat, event.latlng.lng];
  console.log("Selected location: ", ...selectedLocation.value);

  // Reverse geocoding request to Nominatim
  try {
    const response = await axios.get(
      "https://nominatim.openstreetmap.org/reverse",
      {
        params: {
          lat: event.latlng.lat,
          lon: event.latlng.lng,
          format: "json",
        },
      }
    );

    locationName.value = response.data.display_name;
    emitLocationData();
    console.log("Place name: ", locationName.value);
  } catch (error) {
    console.error("Failed to fetch place name: ", error);
  }
};
</script>
