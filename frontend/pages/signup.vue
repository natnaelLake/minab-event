<template>
  <div class="flex h-screen items-center justify-center bg-gray-100 dark:bg-gray-100">
    <form
      @submit.prevent="onSubmit"
      class="bg-white dark:bg-gray-800 w-[380px] p-8 rounded-xl shadow-lg transform transition-all duration-300"
    >
      <h2 class="text-center text-2xl font-semibold text-gray-700 dark:text-gray-200 mb-6">Create an Account</h2>

      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">First Name</label>
        <input
          type="text"
          placeholder="Enter your first name"
          v-bind="first_name"
          class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300"
          :class="{ 'border-red-500 focus:ring-red-500': errors.first_name }"
        />
        <span v-if="errors.first_name" class="text-red-500 text-sm mt-1">{{ errors.first_name }}</span>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Last Name</label>
        <input
          type="text"
          placeholder="Enter your last name"
          v-bind="last_name"
          class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300"
          :class="{ 'border-red-500 focus:ring-red-500': errors.last_name }"
        />
        <span v-if="errors.last_name" class="text-red-500 text-sm mt-1">{{ errors.last_name }}</span>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Email</label>
        <input
          type="email"
          placeholder="Enter your email"
          v-bind="email"
          class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300"
          :class="{ 'border-red-500 focus:ring-red-500': errors.email }"
        />
        <span v-if="errors.email" class="text-red-500 text-sm mt-1">{{ errors.email }}</span>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Password</label>
        <input
          type="password"
          placeholder="Enter your password"
          v-bind="password"
          class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300"
          :class="{ 'border-red-500 focus:ring-red-500': errors.password }"
        />
        <span v-if="errors.password" class="text-red-500 text-sm mt-1">{{ errors.password }}</span>
      </div>

      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Confirm Password</label>
        <input
          type="password"
          placeholder="Confirm your password"
          v-bind="confirm_password"
          class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300"
          :class="{ 'border-red-500 focus:ring-red-500': errors.confirm_password }"
        />
        <span v-if="errors.confirm_password" class="text-red-500 text-sm mt-1">{{ errors.confirm_password }}</span>
      </div>

      <button
        type="submit"
        class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300"
        :disabled="loading"
        :class="{ 'opacity-50 cursor-not-allowed': loading }"
      >
        {{ loading ? "Signing Up..." : "Sign Up" }}
      </button>

      <div class="mt-4 text-center">
        <span class="text-sm text-gray-600 dark:text-gray-400">
          Already have an account?
          <NuxtLink class="text-blue-500 hover:underline dark:text-blue-400" to="/auth/login">Log in</NuxtLink>
        </span>
      </div>
    </form>
  </div>
</template>

<script setup>
import signup from "~/graphql/mutations/SignupMutation.gql";
import { useAuthStore } from "~/store";
import Cookies from "js-cookie";
import { toast } from "vue3-toastify";
import * as yup from "yup";

const authStore = useAuthStore();

const { mutate, onDone, loading, onError } = authentication(signup);

const { defineInputBinds, handleSubmit, errors, setFieldError } = useForm({
  validationSchema: yup.object({
    first_name: yup.string().required("First name is required"),
    last_name: yup.string().required("Last name is required"),
    email: yup
      .string()
      .email("Invalid email address")
      .required("Email is required"),
    password: yup
      .string()
      .min(6, "Password must be at least 6 characters")
      .required("Password is required"),
    confirm_password: yup
      .string()
      .oneOf([yup.ref("password"), null], "Passwords must match")
      .required("Password confirmation is required"),
  }),
});

const first_name = defineInputBinds("first_name");
const last_name = defineInputBinds("last_name");
const email = defineInputBinds("email");
const password = defineInputBinds("password");
const confirm_password = defineInputBinds("confirm_password");

const onSubmit = handleSubmit((values, { setFieldError }) => {
  const input = {
    first_name: values.first_name,
    last_name: values.last_name,
    email: values.email,
    password: values.password,
    confirm_password: values.confirm_password,
  };
  mutate(input);
});

onDone((result) => {
  toast.success("User signed up successfully!", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
  Cookies.set("auth_token", result.data.signup.token, { expires: 7 });

  authStore.setToken(result.data.signup.token);
  authStore.setId(result.data.signup.id);
  authStore.setUser(result.data.signup.id);
  authStore.setRole(result.data.signup.role);
  navigateTo("/");
});

onError((error) => {
  toast.error("Something went wrong", {
    transition: toast.TRANSITIONS.FLIP,
    position: toast.POSITION.TOP_RIGHT,
  });
});
</script>

<style scoped>
/* Focused border color for dark mode */
input:focus {
  outline: none;
}

/* Transition for hover and focus states */
input,
button {
  transition: all 0.3s ease;
}
</style>