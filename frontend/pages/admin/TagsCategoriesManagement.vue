<template>
    <div class="p-6 bg-gray-100 min-h-screen mt-20">
      <div class="max-w-6xl mx-auto">
        <div class="mb-12">
          <h1 class="text-3xl font-bold text-gray-800">
            Manage Tags & Categories
          </h1>
        </div>
  
        <!-- Tabs for Navigation -->
        <div class="flex mb-8">
          <button
            @click="currentTab = 'tags'"
            :class="{
              'bg-blue-500 text-white': currentTab === 'tags',
              'bg-gray-200': currentTab !== 'tags',
            }"
            class="py-2 px-4 rounded-t-lg font-semibold"
          >
            Tags
          </button>
          <button
            @click="currentTab = 'categories'"
            :class="{
              'bg-blue-500 text-white': currentTab === 'categories',
              'bg-gray-200': currentTab !== 'categories',
            }"
            class="py-2 px-4 rounded-t-lg font-semibold"
          >
            Categories
          </button>
        </div>
  
        <!-- Tags Management -->
        <div
          v-if="currentTab === 'tags'"
          class="bg-white shadow-lg rounded-lg p-6"
        >
          <h2 class="text-2xl font-bold text-gray-800 mb-4">Tags</h2>
          <button
            @click="openTagModal"
            class="mb-4 bg-blue-500 text-white px-4 py-2 rounded"
          >
            Add New Tag
          </button>
  
          <div v-if="tags.length > 0" class="space-y-4">
            <div
              v-for="tag in tags"
              :key="tag.id"
              class="flex justify-between items-center bg-gray-50 p-4 rounded-lg shadow"
            >
              <span class="text-gray-800">{{ tag.name }}</span>
              <div class="flex space-x-2">
                <button @click="editTag(tag)" class="text-blue-500">Edit</button>
                <button @click="deleteTag(tag.id)" class="text-red-500">
                  Delete
                </button>
              </div>
            </div>
          </div>
          <p v-else class="text-gray-600">No tags available.</p>
        </div>
  
        <!-- Categories Management -->
        <div
          v-if="currentTab === 'categories'"
          class="bg-white shadow-lg rounded-lg p-6"
        >
          <h2 class="text-2xl font-bold text-gray-800 mb-4">Categories</h2>
          <button
            @click="openCategoryModal"
            class="mb-4 bg-blue-500 text-white px-4 py-2 rounded"
          >
            Add New Category
          </button>
  
          <div v-if="categories.length > 0" class="space-y-4">
            <div
              v-for="category in categories"
              :key="category.id"
              class="flex justify-between items-center bg-gray-50 p-4 rounded-lg shadow"
            >
              <span class="text-gray-800">{{ category.name }}</span>
              <div class="flex space-x-2">
                <button @click="editCategory(category)" class="text-blue-500">
                  Edit
                </button>
                <button @click="deleteCategory(category.id)" class="text-red-500">
                  Delete
                </button>
              </div>
            </div>
          </div>
          <p v-else class="text-gray-600">No categories available.</p>
        </div>
  
        <!-- Modals for Creating and Editing -->
        <Transition name="fade">
          <div
            v-if="showTagModal || showCategoryModal"
            class="fixed inset-0 bg-gray-800 bg-opacity-50 flex items-center justify-center"
          >
            <div class="bg-white p-6 rounded-lg shadow-lg w-96">
              <h3 class="text-xl font-semibold mb-4">{{ modalTitle }}</h3>
              <form @submit.prevent="saveItem">
                <div class="mb-4">
                  <label for="name" class="block text-gray-700">Name</label>
                  <input
                    v-model="currentItem.name"
                    type="text"
                    id="name"
                    class="w-full border border-gray-300 rounded-lg px-4 py-2 mt-1"
                    required
                  />
                </div>
                <div class="flex justify-end space-x-2">
                  <button
                    @click="closeModal"
                    type="button"
                    class="bg-red-500 text-white px-4 py-2 rounded"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    class="bg-blue-500 text-white px-4 py-2 rounded"
                  >
                    {{ modalButton }}
                  </button>
                </div>
              </form>
            </div>
          </div>
        </Transition>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from "vue";
  
  const tags = ref([
    { id: 1, name: "Technology" },
    { id: 2, name: "Health" },
    { id: 3, name: "Education" }
  ]);
  
  const categories = ref([
    { id: 1, name: "Workshops" },
    { id: 2, name: "Webinars" },
    { id: 3, name: "Conferences" }
  ]);
  
  const currentTab = ref("tags");
  const showTagModal = ref(false);
  const showCategoryModal = ref(false);
  const modalTitle = ref("");
  const modalButton = ref("");
  const currentItem = ref({ name: "" });
  const isEditing = ref(false);
  
  const openTagModal = () => {
    modalTitle.value = "Add New Tag";
    modalButton.value = "Save";
    currentItem.value = { name: "" };
    showTagModal.value = true;
  };
  
  const openCategoryModal = () => {
    modalTitle.value = "Add New Category";
    modalButton.value = "Save";
    currentItem.value = { name: "" };
    showCategoryModal.value = true;
  };
  
  const editTag = (tag) => {
    modalTitle.value = "Edit Tag";
    modalButton.value = "Update";
    currentItem.value = { ...tag };
    isEditing.value = true;
    showTagModal.value = true;
  };
  
  const editCategory = (category) => {
    modalTitle.value = "Edit Category";
    modalButton.value = "Update";
    currentItem.value = { ...category };
    isEditing.value = true;
    showCategoryModal.value = true;
  };
  
  const saveItem = () => {
    if (isEditing.value) {
      if (currentTab.value === "tags") {
        // Update tag logic
        alert(`Tag updated: ${currentItem.value.name}`);
      } else {
        // Update category logic
        alert(`Category updated: ${currentItem.value.name}`);
      }
    } else {
      if (currentTab.value === "tags") {
        // Create tag logic
        alert(`New tag created: ${currentItem.value.name}`);
      } else {
        // Create category logic
        alert(`New category created: ${currentItem.value.name}`);
      }
    }
    closeModal();
  };
  
  const deleteTag = (id) => {
    // Delete tag logic
    alert(`Tag with ID ${id} deleted`);
  };
  
  const deleteCategory = (id) => {
    // Delete category logic
    alert(`Category with ID ${id} deleted`);
  };
  
  const closeModal = () => {
    showTagModal.value = false;
    showCategoryModal.value = false;
    isEditing.value = false;
  };
  </script>
  
  <style scoped>
  /* Basic styles for modals and tabs */
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.3s;
  }
  .fade-enter,
  .fade-leave-to {
    opacity: 0;
  }
  
  button {
    cursor: pointer;
  }
  </style>
  