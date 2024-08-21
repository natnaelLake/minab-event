<template>
  <div>
    <nav class="bg-blue-600 p-4 text-white">
      <ul class="flex space-x-4">
        <li><NuxtLink to="/">Home</NuxtLink></li>
        <li><NuxtLink to="/tasks">Tasks</NuxtLink></li>
        <li><NuxtLink to="/login">Login</NuxtLink></li>
        <li><NuxtLink to="/signup">Sign Up</NuxtLink></li>
      </ul>
    </nav>
    <Nuxt />
  </div>
</template>
