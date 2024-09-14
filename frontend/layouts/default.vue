<template>
  <div> 
    <header>
      <Navigation />
    </header>
    <main class="h-screen">
      <slot/>
    </main>
    <!-- <footer>
      <p>&copy; 2024 My Task Organizer</p>
    </footer> -->
  </div>
</template>

<script setup>
import Navigation from "~/components/Navigation.vue";
</script>

<style scoped>
main {
  background-color: #e2e5eb;
}
</style>
