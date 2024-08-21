<template>
  <div>
    <h1>Create Event</h1>
    <!-- Form for creating an event -->
  </div>
</template>

<script setup>
// Requires authentication and the 'user' role
</script>

<script>
export default {
  middleware: "auth",
  meta: {
    requiresAuth: true,
    requiresRole: "user",
  },
};
</script>
