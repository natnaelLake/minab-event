<template>
	<div class="flex h-screen items-center justify-center bg-gray-100 dark:bg-gray-100">
	  <form
		@submit.prevent="onSubmit"
		class="bg-white dark:bg-gray-800 w-[380px] p-8 rounded-xl shadow-lg transform transition-all duration-300"
	  >
		<h2 class="text-center text-2xl font-semibold text-gray-700 dark:text-gray-200 mb-6">Welcome Back</h2>
  
		<div class="mb-5">
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
  
		<div class="mb-6">
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
			<NuxtLink class="text-blue-500 hover:underline dark:text-blue-400" to="/signup">Sign up</NuxtLink>
		  </span>
		</div>
	  </form>
	</div>
  </template>
  
  <script setup>
  import Login from "~/graphql/mutations/LoginMutation.gql";
  import Cookies from "js-cookie";
  import { toast } from "vue3-toastify";
  import { useForm } from "vee-validate";
  import * as yup from "yup";
  import { useAuthStore } from "~/store";
  
  const authStore = useAuthStore();
  const { mutate, onDone, loading, onError } = authentication(Login);
  
  const { defineInputBinds, handleSubmit, errors, setFieldError } = useForm({
	validationSchema: yup.object({
	  email: yup.string().email().required(),
	  password: yup.string().min(6).required(),
	}),
  });
  
  const email = defineInputBinds("email");
  const password = defineInputBinds("password");
  
  const onSubmit = handleSubmit((values, { setFieldError }) => {
	console.log(values);
	const input = {
	  email: values.email,
	  password: values.password,
	};
	mutate(input);
  });
  
  onDone((result) => {
	Cookies.set("auth_token", result.data.login.token, { expires: 7 });
	authStore.setToken(result.data.login.token);
	authStore.setId(result.data.login.id);
	authStore.setUser(result.data.login.id);
	authStore.setRole(result.data.login.role);
	navigateTo("/");
  });
  
  onError((error) => {
	if (error.message.includes("Invalid")) {
	  setFieldError("email", "Invalid email or password");
	  setFieldError("password", "Invalid email or password");
	} else {
	  toast.error("Something went wrong", {
		transition: toast.TRANSITIONS.FLIP,
		position: toast.POSITION.TOP_RIGHT,
	  });
	}
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
  