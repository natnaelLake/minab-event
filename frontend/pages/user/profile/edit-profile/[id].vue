<template>
    <div class="container mx-auto p-4 mt-20">
      <div class="bg-white shadow-md rounded-lg p-10 mt-10 max-w-3xl mx-auto">
        <h2 class="text-2xl font-bold text-center mb-6">Edit Profile</h2>
        
        <div class="flex flex-col items-center mb-6">
          <img
            class="w-24 h-24 rounded-full object-cover mb-4"
            :src="user.avatar"
            alt="User Avatar"
          />
          <input
            type="file"
            @change="updateAvatar"
            class="hidden"
            ref="avatarInput"
          />
          <button
            @click="triggerAvatarUpload"
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
          >
            Change Avatar
          </button>
        </div>
        
        <form @submit.prevent="updateProfile">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
            <div>
              <label class="block text-gray-700">Name</label>
              <input
                v-model="user.name"
                type="text"
                class="w-full p-2 border border-gray-300 rounded-lg mt-1"
                required
              />
            </div>
            <div>
              <label class="block text-gray-700">Title</label>
              <input
                v-model="user.title"
                type="text"
                class="w-full p-2 border border-gray-300 rounded-lg mt-1"
                required
              />
            </div>
          </div>
          
          <div class="mb-6">
            <label class="block text-gray-700">About</label>
            <textarea
              v-model="user.about"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
              rows="4"
              required
            ></textarea>
          </div>
          
          <div class="mb-6">
            <label class="block text-gray-700">Skills</label>
            <input
              v-model="newSkill"
              type="text"
              placeholder="Add a skill and press Enter"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
              @keyup.enter="addSkill"
            />
            <div class="mt-2">
              <span
                v-for="(skill, index) in user.skills"
                :key="index"
                class="inline-block bg-blue-100 text-blue-700 text-sm px-3 py-1 rounded-full mr-2 mt-2"
              >
                {{ skill }}
                <button
                  type="button"
                  @click="removeSkill(index)"
                  class="ml-2 text-red-500"
                >
                  &times;
                </button>
              </span>
            </div>
          </div>
          
          <div class="text-center">
            <button
              type="submit"
              class="bg-green-500 text-white px-6 py-3 rounded-lg hover:bg-green-600"
            >
              Save Changes
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const user = ref({
    avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTafDGOCQJda4LjZXx9U4ZGlFdjYyVmZNRrwA&s',
    name: 'John Doe',
    title: 'Software Engineer at Company',
    about: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    skills: ['JavaScript', 'Vue.js', 'Nuxt.js', 'GraphQL'],
  });
  
  const newSkill = ref('');
  
  const updateAvatar = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        user.value.avatar = e.target.result;
      };
      reader.readAsDataURL(file);
    }
  };
  
  const triggerAvatarUpload = () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/*';
    input.onchange = updateAvatar;
    input.click();
  };
  
  const addSkill = () => {
    if (newSkill.value.trim() && !user.value.skills.includes(newSkill.value.trim())) {
      user.value.skills.push(newSkill.value.trim());
    }
    newSkill.value = '';
  };
  
  const removeSkill = (index) => {
    user.value.skills.splice(index, 1);
  };
  
  const updateProfile = () => {
    // Logic to update the profile goes here.
    console.log('Profile updated:', user.value);
  };
  </script>
  
  <style scoped>
  .container {
    max-width: 800px;
  }
  </style>
  