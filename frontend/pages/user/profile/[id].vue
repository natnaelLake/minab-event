<template>
  <div class="container mx-auto p-4">
    <div class="bg-white shadow rounded-lg p-28 mt-20">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <img
            class="w-20 h-20 rounded-full object-cover"
            :src="user.avatar"
            alt="User Avatar"
          />
          <div class="ml-4">
            <h2 class="text-xl font-bold">{{ user.name }}</h2>
            <p class="text-gray-600">{{ user.title }}</p>
          </div>
        </div>
        <!-- <div>
          <button
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
          >
            Edit Profile
          </button>
        </div> -->
      </div>

      <div class="mt-6 flex justify-around">
        <button
          class="text-blue-500 font-semibold hover:underline"
          @click="redirectToFollowing"
        >
          {{ user.following }} Following
        </button>
        <button
          class="text-blue-500 font-semibold hover:underline"
          @click="redirectToFollowers"
        >
          {{ user.followers }} Followers
        </button>
      </div>

      <div class="mt-6">
        <h3 class="text-lg font-semibold">About</h3>
        <p class="text-gray-700 mt-2">{{ user.about }}</p>
      </div>

      <div class="mt-6">
        <h3 class="text-lg font-semibold">Skills</h3>
        <ul class="list-disc pl-5 text-gray-700 mt-2">
          <li v-for="skill in user.skills" :key="skill">{{ skill }}</li>
        </ul>
      </div>

      <div class="mt-6">
        <h3 class="text-lg font-semibold">Events Created</h3>
        <ul class="mt-4">
          <li
            v-for="event in user.events.slice(0, 5)"
            :key="event.id"
            class="mb-4 p-4 bg-gray-100 rounded-lg shadow"
          >
            <h4 class="text-lg font-bold">{{ event.title }}</h4>
            <p class="text-gray-600">{{ event.description }}</p>
            <p class="text-sm text-gray-500">Date: {{ event.date }}</p>
          </li>
        </ul>
      </div>

      <!-- Add more sections here as needed -->
    </div>
  </div>
</template>

<script setup>
const user = {
  avatar:
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTafDGOCQJda4LjZXx9U4ZGlFdjYyVmZNRrwA&s",
  name: "John Doe",
  title: "Software Engineer at Company",
  following: 123,
  followers: 456,
  about:
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
  skills: ["JavaScript", "Vue.js", "Nuxt.js", "GraphQL"],
  events: [
    {
      id: 1,
      title: "Tech Conference 2024",
      description:
        "An annual conference focused on the latest trends in technology.",
      date: "2024-09-10",
    },
    {
      id: 2,
      title: "Vue.js Workshop",
      description: "A hands-on workshop to learn and master Vue.js.",
      date: "2024-08-25",
    },
    {
      id: 3,
      title: "Startup Pitch Night",
      description: "An event where startups pitch their ideas to investors.",
      date: "2024-10-05",
    },
    {
      id: 4,
      title: "Blockchain Summit",
      description: "A summit discussing the future of blockchain technology.",
      date: "2024-11-12",
    },
    {
      id: 5,
      title: "AI in Healthcare",
      description: "Exploring the impact of AI in the healthcare industry.",
      date: "2024-12-15",
    },
  ],
};
const redirectToFollowing = () => {
  navigateTo("/user/follows");
};
const redirectToFollowers = () => {
  navigateTo("/user/follows");
};
</script>

<style scoped>
.container {
  max-width: 800px;
}
</style>
