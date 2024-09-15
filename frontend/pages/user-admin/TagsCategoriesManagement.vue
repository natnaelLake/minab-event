<script setup>
import { ref, watch, onMounted } from "vue";
import { useQuery, useMutation } from "@vue/apollo-composable";
import GetAllCategories from "~/graphql/query/GetAllCategories.gql";
import GetAllTags from "~/graphql/query/GetAllTags.gql";
import DeleteSingleCategory from "~/graphql/mutations/DeleteCategory.gql";
import DeleteSingleTag from "~/graphql/mutations/DeleteTag.gql";
import UpdateSingleCategory from "~/graphql/mutations/UpdateCategory.gql";
import UpdateSingleTag from "~/graphql/mutations/UpdateTag.gql";
import CreateSingleCategory from "~/graphql/mutations/CreateCategoryMutation.gql";
import CreateSingleTag from "~/graphql/mutations/CreateTagMutation.gql";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";

const tags = ref([]);
const categories = ref([]);
const currentTab = ref("tags");
const showTagModal = ref(false);
const showCategoryModal = ref(false);
const modalTitle = ref("");
const modalButton = ref("");
const currentItem = ref({ name: "", description: "", tag_image: "" });
const isEditing = ref(false);
const user = useAuthStore();
const currentUserId = user.id;
const currentUserRole = user.role;
const currentUserToken = user.token;
const { onResult: tagResult, refetch: refetchTags } = useQuery(
  GetAllTags,
  {
    limit: 10,
    offset: 0,
    order_by: [{ created_at: "desc" }],
  },
  {
    context: {
      headers: {
        "x-hasura-user-id": currentUserId,
        "x-hasura-role": currentUserRole,
        Authorization: `Bearer ${currentUserToken}`,
      },
    },
  }
);
const { onResult: categoryResult, refetch: refetchCategories } = useQuery(
  GetAllCategories,
  {
    limit: 10,
    offset: 0,
    order_by: [{ created_at: "desc" }],
  },
  {
    context: {
      headers: {
        "x-hasura-user-id": currentUserId,
        "x-hasura-role": currentUserRole,
        Authorization: `Bearer ${currentUserToken}`,
      },
    },
  }
);
const { mutate: createTag } = useMutation(CreateSingleTag, {
  context: {
    headers: {
      "x-hasura-user-id": currentUserId,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${currentUserToken}`,
    },
  },
});
const { mutate: createCategory } = useMutation(CreateSingleCategory, {
  context: {
    headers: {
      "x-hasura-user-id": currentUserId,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${currentUserToken}`,
    },
  },
});
const { mutate: deleteTag } = useMutation(DeleteSingleTag, {
  context: {
    headers: {
      "x-hasura-user-id": currentUserId,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${currentUserToken}`,
    },
  },
});
const { mutate: deleteCategory } = useMutation(DeleteSingleCategory, {
  context: {
    headers: {
      "x-hasura-user-id": currentUserId,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${currentUserToken}`,
    },
  },
});
const { mutate: updateTag } = useMutation(UpdateSingleTag, {
  context: {
    headers: {
      "x-hasura-user-id": currentUserId,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${currentUserToken}`,
    },
  },
});
const { mutate: updateCategory } = useMutation(UpdateSingleCategory, {
  context: {
    headers: {
      "x-hasura-user-id": currentUserId,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${currentUserToken}`,
    },
  },
});
tagResult((result) => {
  if (result.data) {
    tags.value = result.data.tags;
  }
});
categoryResult((result) => {
  if (result.data) {
    categories.value = result.data.categories;
  }
});

// Watcher to refetch data when the tab is changed
watch(currentTab, async (newTab) => {
  if (newTab === "tags") {
    await refetchTags();
  } else {
    await refetchCategories();
  }
});

onMounted(async () => {
  await refetchTags();
  await refetchCategories();
});

const changeTab = (tab) => {
  currentTab.value = tab;
};

const openTagModal = () => {
  modalTitle.value = "Add New Tag";
  modalButton.value = "Save";
  currentItem.value = { name: "", description: "", tag_image: "" };
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

const saveItem = async () => {
  try {
    if (isEditing.value) {
      if (currentTab.value === "tags") {
        await updateTag({
          tag_id: currentItem.value.id,
          changes: { name: currentItem.value.name },
        });
        toast.success("Tag updated successfully.");
      } else {
        await updateCategory({
          category_id: currentItem.value.id,
          changes: { name: currentItem.value.name },
        });
        toast.success("Category updated successfully.");
      }
    } else {
      if (currentTab.value === "tags") {
        await createTag({ name: currentItem.value.name });
        toast.success("Tag created successfully.");
      } else {
        await createCategory({ name: currentItem.value.name });
        toast.success("Category created successfully.");
      }
    }
    await refetchData();
    closeModal();
  } catch (error) {
    console.error("Error saving data: ", error);
    toast.error("Failed to save data.");
  }
};

const handleDeleteTag = async (id) => {
  try {
    await deleteTag({ tag_id: id });
    toast.success("Tag deleted successfully.");
    await refetchTags();
  } catch (error) {
    console.error("Error deleting tag: ", error);
    toast.error("Failed to delete tag.");
  }
};

const handleDeleteCategory = async (id) => {
  try {
    await deleteCategory({ category_id: id });
    toast.success("Category deleted successfully.");
    await refetchCategories();
  } catch (error) {
    console.error("Error deleting category: ", error);
    toast.error("Failed to delete category.");
  }
};

const closeModal = () => {
  showTagModal.value = false;
  showCategoryModal.value = false;
  isEditing.value = false;
  currentItem.value = { name: "" };
};

const refetchData = async () => {
  if (currentTab.value === "tags") {
    await refetchTags();
  } else {
    await refetchCategories();
  }
};
</script>

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
          @click="changeTab('tags')"
          :class="{
            'bg-blue-500 text-white': currentTab === 'tags',
            'bg-gray-200': currentTab !== 'tags',
          }"
          class="py-2 px-4 rounded-t-lg font-semibold"
        >
          Tags
        </button>
        <button
          @click="changeTab('categories')"
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
              <button @click="handleDeleteTag(tag.id)" class="text-red-500">
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
              <button
                @click="handleDeleteCategory(category.id)"
                class="text-red-500"
              >
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
                <label for="name" class="block text-gray-700"
                  >Description</label
                >
                <textarea
                  v-model="currentItem.description"
                  type="text"
                  cols="4"
                  id="description"
                  class="w-full border border-gray-300 rounded-lg px-4 py-2 mt-1"
                  required
                />
                <label for="name" class="block text-gray-700">Tag Image</label>
                <input
                  v-model="currentItem.tag_image"
                  type="file"
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


<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
