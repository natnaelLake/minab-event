<template>
  <div class="container mx-auto p-4 mt-20">
    <div class="bg-white shadow-md rounded-lg p-10 mt-10 max-w-3xl mx-auto">
      <h2 class="text-2xl font-bold text-center mb-6">Edit Profile</h2>

      <div class="flex flex-col items-center mb-6">
        <img
          class="w-24 h-24 rounded-full object-cover mb-4"
          :src="user.profile_image_url"
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
            <label class="block text-gray-700">First Name</label>
            <input
              v-model="user.first_name"
              type="text"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
              required
            />
          </div>
          <div>
            <label class="block text-gray-700">Last Name</label>
            <input
              v-model="user.last_name"
              type="text"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
              required
            />
          </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
          <div>
            <label class="block text-gray-700">Email</label>
            <input
              v-model="user.email"
              type="email"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
              required
            />
          </div>
          <div>
            <label class="block text-gray-700">Password</label>
            <input
              v-model="user.password"
              type="password"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
              disabled="true"
              required
            />
          </div>
        </div>

        <div class="mb-6">
          <label class="block text-gray-700">About</label>
          <textarea
            v-model="user.bio"
            class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            rows="4"
            required
          ></textarea>
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
import getSingleUser from "~/graphql/query/getSingleUser.gql";
import UpdateSingleUser from "~/graphql/mutations/UpdateSingleUser.gql";
import UploadImageMutation from "~/graphql/mutations/UploadImageMutation.gql";
import { ref, onMounted } from "vue";
import { useMutation, useQuery } from "@vue/apollo-composable";
import { useAuthStore } from "~/store";
import { toast } from "vue3-toastify";

const userStore = useAuthStore();
const currentUser = userStore.id;
const user = ref({});
const {
  mutate: uploadImages,
  loading: uploadImageLoading,
  onDone: onUploadImageDone,
  onError: onUploadImageError,
} = useMutation(UploadImageMutation, {
  context: {
    headers: {
      "x-hasura-user-id": currentUser,
      "x-hasura-role": user.role,
      Authorization: `Bearer ${user.token}`,
    },
  },
});
onMounted(async () => {
  await fetchUser();
});

const fetchUser = async () => {
  try {
    const { onResult } = useQuery(
      getSingleUser,
      {
        id: currentUser,
      },
      {
        context: {
          headers: {
            "x-hasura-user-id": currentUser,
            "x-hasura-role": user.role,
            Authorization: `Bearer ${user.token}`,
          },
        },
      }
    );

    onResult((result) => {
      if (result.data) {
        user.value = result.data.users_by_pk;
      }
    });
  } catch (error) {
    console.error("Error fetching user: ", error.message);
    toast.error("Something went wrong, try again", {
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};
const readFilesAsBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onloadend = () => {
      if (reader.result) {
        resolve(reader.result.split(",")[1]); // Return only the base64 part
      } else {
        reject(new Error("Failed to read file"));
      }
    };
    reader.readAsDataURL(file);
  });
};

const updateAvatar = async (e) => {
  const file = e.target.files[0];
  if (file) {
    try {
      const base64File = await readFilesAsBase64(file);

      const uploadImagesInput = { input: { files: base64File } };

      const { data } = await uploadImages(uploadImagesInput);

      const uploadedImageUrl = data.uploadImages.imageUrls[0];

      user.value = {
        ...user.value,
        profile_image_url: uploadedImageUrl,
      };
    } catch (error) {
      console.error("Error uploading image:", error);
    }
  }
};

const triggerAvatarUpload = () => {
  const avatarInput = document.createElement("input");
  avatarInput.type = "file";
  avatarInput.accept = "image/*";
  avatarInput.onchange = updateAvatar;
  avatarInput.click();
};

const { mutate: updateSingleUser } = useMutation(UpdateSingleUser, {
  context: {
    headers: {
      "x-hasura-user-id": currentUser,
      "x-hasura-role": user.role,
      Authorization: `Bearer ${user.token}`,
    },
  },
});

const updateProfile = async () => {
  try {
    const response = await updateSingleUser({
      user_id: user.value.id,
      changes: {
        first_name: user.value.first_name,
        last_name: user.value.last_name,
        email: user.value.email,
        password: user.value.password,
        bio: user.value.bio,
        profile_image_url: user.value.profile_image_url,
      },
    });

    if (response.data) {
      toast.success("Profile updated successfully!", {
        position: toast.POSITION.TOP_RIGHT,
      });
      await fetchUser();
    }
  } catch (error) {
    console.error("Error updating profile: ", error);
    toast.error("Failed to update profile. Please try again.", {
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};
</script>

<style scoped>
.container {
  max-width: 800px;
}
</style>
