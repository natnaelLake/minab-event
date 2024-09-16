<script setup>
import getSingleUser from "~/graphql/query/getSingleUser.gql";
import UpdateSingleUser from "~/graphql/mutations/UpdateSingleUser.gql";
import UploadImageMutation from "~/graphql/mutations/UploadImageMutation.gql";
import { ref, onMounted } from "vue";
import { useAuthStore } from "~/store";
import { toast } from "vue3-toastify";
import { hashPassword, comparePasswords } from "~/utils/authUtils"; // Import password utils
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

// User store for authentication
const userStore = useAuthStore();
const currentUser = userStore.id; // Get the current user's ID
const user = ref({
  first_name: "",
  last_name: "",
  email: "",
  bio: "",
  profile_image_url: "",
  current_password: "",
  new_password: "",
  confirm_password: "",
});

// Define validation schema using yup
const schema = yup.object({
  first_name: yup.string().required("First name is required"),
  last_name: yup.string().required("Last name is required"),
  email: yup.string().email("Invalid email").required("Email is required"),
  current_password: yup.string().optional(),
  new_password: yup
    .string()
    .optional()
    .min(6, "New password must be at least 6 characters"),
  confirm_password: yup
    .string()
    .optional()
    .oneOf([yup.ref("new_password")], "Passwords do not match"),
  bio: yup.string().required("Bio is required"),
});

// Upload mutation setup
const { mutate: uploadImages, loading: uploadImageLoading } =
  useMutationComposable(UploadImageMutation);

onMounted(async () => {
  await fetchUser();
});

// Fetch user details
const fetchUser = async () => {
  try {
    const { onResult } = useQueryComposable(getSingleUser, { id: currentUser });
    onResult((result) => {
      if (result.data) {
        user.value = { ...result.data.users_by_pk };
        console.log("***************", user);
      }
    });
  } catch (error) {
    toast.error("Something went wrong, try again", {
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};

// Convert file to base64
const readFilesAsBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onloadend = () => {
      if (reader.result) {
        resolve(reader.result.split(",")[1]);
      } else {
        reject(new Error("Failed to read file"));
      }
    };
    reader.readAsDataURL(file);
  });
};

// Update avatar
const updateAvatar = async (e) => {
  const file = e.target.files[0];
  if (file) {
    try {
      const base64File = await readFilesAsBase64(file);
      const uploadImagesInput = { input: { files: base64File } };
      const { data } = await uploadImages(uploadImagesInput);
      const uploadedImageUrl = data.uploadImages.imageUrls[0];
      user.value.profile_image_url = uploadedImageUrl; // Update the avatar URL
    } catch (error) {
      console.error("Error uploading image:", error);
    }
  }
};

// Trigger avatar file input
const triggerAvatarUpload = () => {
  const avatarInput = document.createElement("input");
  avatarInput.type = "file";
  avatarInput.accept = "image/*";
  avatarInput.onchange = updateAvatar;
  avatarInput.click();
};

// Mutation to update the user profile
const { mutate: updateSingleUser } = useMutationComposable(UpdateSingleUser);

// Update profile
const validatePasswords = (values) => {
  // Validate current_password
  if (values.current_password) {
    if (!values.new_password || !values.confirm_password) {
      throw new Error("New password and confirm password are required when current password is provided.");
    }

    if (values.new_password !== values.confirm_password) {
      throw new Error("New password and confirm password do not match.");
    }
  } else {
    if (values.new_password || values.confirm_password) {
      throw new Error("Current password is required to update the new password.");
    }
  }
};

const updateProfile = async (values) => {
  try {
    // Validate passwords
    validatePasswords(values);

    // Check if the current password matches
    if (values.current_password) {
      const validPassword = await comparePasswords(values.current_password, user.value.password);

      if (!validPassword) {
        toast.error("Incorrect current password.", {
          position: toast.POSITION.TOP_RIGHT,
        });
        return;
      }
    }

    // If a new password is provided, hash it
    let hashedNewPassword;
    if (values.new_password) {
      hashedNewPassword = await hashPassword(values.new_password);
    }

    // Proceed with profile update
    const response = await updateSingleUser({
      user_id: user.value.id,
      changes: {
        first_name: values.first_name,
        last_name: values.last_name,
        email: values.email,
        password: hashedNewPassword || user.value.password,
        bio: values.bio,
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
    toast.error(error.message || "Failed to update profile. Please try again.", {
      position: toast.POSITION.TOP_RIGHT,
    });
  }
};
</script>

<template>
  <div class="container mx-auto p-4 mt-20">
    <div class="bg-white shadow-md rounded-lg p-10 mt-10 max-w-3xl mx-auto">
      <h2 class="text-2xl font-bold text-center mb-6">Edit Profile</h2>

      <Form :validation-schema="schema" @submit="updateProfile">
        <div class="flex flex-col items-center mb-6">
          <img
            class="w-24 h-24 rounded-full object-cover mb-4"
            :src="user.profile_image_url"
            alt="User Avatar"
          />
          <button
            @click="triggerAvatarUpload"
            type="button"
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
          >
            Change Avatar
          </button>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
          <div>
            <label class="block text-gray-700">First Name</label>
            <Field
              :model-value="user.first_name"
              name="first_name"
              type="text"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            />
            <ErrorMessage name="first_name" class="text-red-500 mt-1" />
          </div>
          <div>
            <label class="block text-gray-700">Last Name</label>
            <Field
              :model-value="user.last_name"
              name="last_name"
              type="text"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            />
            <ErrorMessage name="last_name" class="text-red-500 mt-1" />
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
          <div>
            <label class="block text-gray-700">Email</label>
            <Field
              :model-value="user.email"
              name="email"
              type="email"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            />
            <ErrorMessage name="email" class="text-red-500 mt-1" />
          </div>
          <div>
            <label class="block text-gray-700">Current Password</label>
            <Field
              v-model="user.current_password"
              name="current_password"
              type="text"
              placeholder="Enter current password"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            />
            <ErrorMessage name="current_password" class="text-red-500 mt-1" />
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
          <div>
            <label class="block text-gray-700">New Password</label>
            <Field
              v-model="user.new_password"
              name="new_password"
              type="text"
              placeholder="Enter new password"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            />
            <ErrorMessage name="new_password" class="text-red-500 mt-1" />
          </div>
          <div>
            <label class="block text-gray-700">Confirm Password</label>
            <Field
              v-model="user.confirm_password"
              name="confirm_password"
              type="text"
              placeholder="Confirm new password"
              class="w-full p-2 border border-gray-300 rounded-lg mt-1"
            />
            <ErrorMessage name="confirm_password" class="text-red-500 mt-1" />
          </div>
        </div>

        <div class="mb-6">
          <label class="block text-gray-700">Bio</label>
          <Field
            v-model="user.bio"
            name="bio"
            as="textarea"
            class="w-full p-2 border border-gray-300 rounded-lg mt-1"
          />
          <ErrorMessage name="bio" class="text-red-500 mt-1" />
        </div>

        <div class="text-center">
          <button
            type="submit"
            class="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600"
          >
            Update Profile
          </button>
        </div>
      </Form>
    </div>
  </div>
</template>
