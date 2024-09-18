<script setup>
import DynamicForm from "@/components/DynamicForm.vue";
import Login from "~/graphql/mutations/LoginMutation.gql";
import Cookies from "js-cookie";
import { toast } from "vue3-toastify";
import * as yup from "yup";
import { useAuthStore } from "~/store";

const authStore = useAuthStore();
const { mutate, onDone, loading, onError } = authentication(Login);

const loginSchema = {
  fields: [
    {
      as: "input",
      name: "email",
      label: "Email",
      placeholder: "Enter your email",
      type: "email",
      rules: yup
        .string()
        .email("Invalid email address")
        .required("Email is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "password",
      label: "Password",
      placeholder: "Enter your password",
      type: "password",
      rules: yup
        .string()
        .min(6, "Password must be at least 6 characters")
        .required("Password is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-400 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
  ],
};

const onSubmit = async (values) => {
  const input = {
    email: values.email,
    password: values.password,
  };
  mutate(input);
};

onDone((result) => {
  if (result.data.status !== "Suspended") {
    Cookies.set("auth_token", result.data.login.token, { expires: 7 });
    authStore.setToken(result.data.login.token);
    authStore.setId(result.data.login.id);
    authStore.setUser(result.data.login.id);
    authStore.setRole(result.data.login.role);
    if (result.data.login.role === "user-admin") {
      navigateTo("/user-admin/dashboard");
    } else {
      navigateTo("/");
    }
  } else {
    toast.error("Something went wrong");
  }
});

onError((error) => {
  if (error.message.includes("Invalid")) {
    toast.error("Invalid email or password");
  } else {
    toast.error("Something went wrong");
  }
});
</script>

<template>
  <div
    class="flex h-screen items-center justify-center bg-gray-100 dark:bg-gray-100"
  >
    <DynamicForm :schema="loginSchema" :submitHandler="onSubmit">
      <button
        type="submit"
        class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-lg shadow-md focus:outline-none transition-all duration-300"
        :disabled="loading"
        :class="{ 'opacity-50 cursor-not-allowed': loading }"
      >
        {{ loading ? "Logging In..." : "Log in" }}
      </button>
      <div class="mt-4 text-center">
        <span class="text-sm text-gray-600 dark:text-gray-400">
          Don't have an account?
          <NuxtLink
            class="text-blue-500 hover:underline dark:text-blue-400"
            to="/signup"
          >
            Sign up
          </NuxtLink>
        </span>
      </div>
    </DynamicForm>
  </div>
</template>
