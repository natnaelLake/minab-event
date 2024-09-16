<script setup>
import DynamicForm from "@/components/DynamicForm.vue";
import signup from "~/graphql/mutations/SignupMutation.gql";
import { toast } from "vue3-toastify";
import * as yup from "yup";

const { mutate, onDone, loading, onError } = authentication(signup);

const signupSchema = {
  fields: [
    {
      as: "input",
      name: "first_name",
      label: "First Name",
      placeholder: "Enter your First Name",
      type: "text",
      rules: yup.string().required("First name is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-800 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "last_name",
      label: "Last Name",
      placeholder: "Enter your Last Name",
      type: "text",
      rules: yup.string().required("Last name is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-800 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
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
        label: "text-sm font-medium text-gray-600 dark:text-gray-800 mb-1",
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
        label: "text-sm font-medium text-gray-600 dark:text-gray-800 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
    {
      as: "input",
      name: "confirm_password",
      label: "Confirm Password",
      placeholder: "Confirm your password",
      type: "password",
      rules: yup
        .string()
        .min(6, "Passwords must match")
        .required("Confirm Password is required"),
      class: {
        wrapper: "mb-5",
        label: "text-sm font-medium text-gray-600 dark:text-gray-800 mb-1",
        input:
          "w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300",
        error: "text-red-500 text-sm mt-1",
      },
    },
  ],
};

const onSubmit = async (values) => {
  const input = {
    first_name: values.first_name,
    last_name: values.last_name,
    email: values.email,
    password: values.password,
    confirm_password: values.confirm_password,
  };
  mutate(input);
};
onDone((result) => {
  toast.success(`${result.data.signup.message}`);
});

onError((error) => {
  toast.error("Something went wrong");
});
</script>

<template>
  <div
    class="flex items-center justify-center bg-gray-100 dark:bg-gray-100 mt-16 pb-10 pt-10"
  >
    <DynamicForm :schema="signupSchema" :submitHandler="onSubmit">
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
          <NuxtLink
            class="text-blue-500 hover:underline dark:text-blue-400"
            to="/login"
          >
            Log in
          </NuxtLink>
        </span>
      </div>
    </DynamicForm>
  </div>
</template>
