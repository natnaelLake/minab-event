<template>
  <div
    class="min-h-screen bg-gray-50 flex flex-col justify-center items-center py-12 px-4 sm:px-6 lg:px-8"
  >
    <div class="bg-white shadow-xl rounded-lg p-8 max-w-lg w-full space-y-6">
      <div class="text-center">
        <img
          src="https://via.placeholder.com/150"
          alt="Company Logo"
          class="h-12 w-auto mx-auto"
        />
        <h2 class="mt-6 text-3xl font-extrabold text-gray-900">
          Email Verification
        </h2>
      </div>
      <div class="space-y-4">
        <p v-if="isVerifying" class="text-sm text-gray-600">
          Verifying your email...
        </p>
        <div
          v-if="verificationSuccess"
          class="flex flex-col items-center space-y-2"
        >
          <svg
            class="w-12 h-12 text-green-500"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M9 12l2 2 4-4M7 12l-4-4m0 8l4-4m4 4l4-4"
            ></path>
          </svg>
          <h3 class="text-lg font-medium text-gray-900">
            Email Verified Successfully!
          </h3>
          <p class="text-gray-600 text-sm">
            Your email has been verified. You can now proceed to login and enjoy
            our services.
          </p>
          <router-link
            to="/login"
            class="text-indigo-600 hover:text-indigo-900 transition-colors font-medium"
          >
            Go to Login
          </router-link>
        </div>
        <div
          v-else-if="!isVerifying && !verificationSuccess"
          class="text-center"
        >
          <svg
            class="w-12 h-12 text-red-500"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M13 17h8m0 0l-8-8m8 8H5m0 0l8-8m-8 8H3"
            ></path>
          </svg>
          <h3 class="text-lg font-medium text-gray-900">
            Email Verification Failed!
          </h3>
          <p class="text-gray-600 text-sm">
            The verification link is either invalid or expired. Please request a
            new verification email.
          </p>
          <router-link
            to="/resend-verification"
            class="text-indigo-600 hover:text-indigo-900 transition-colors font-medium"
          >
            Resend Verification Email
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "~/store";
import verifyEmail from "~/graphql/mutations/VerifyEmailMutation.gql";
import { toast } from "vue3-toastify";
import Cookies from "js-cookie";

const isVerifying = ref(true);
const verificationSuccess = ref(false);
const route = useRoute();
const authStore = useAuthStore();

const { mutate, onDone, loading, onError } = authentication(verifyEmail);
onDone((result) => {
  verificationSuccess.value = true;
  toast.success("User signed up successfully!");
  console.log('_______============',result.data);
  Cookies.set("auth_token", result.data.verifyEmail.token, { expires: 7 });

  authStore.setToken(result.data.verifyEmail.token);
  authStore.setId(result.data.verifyEmail.id);
  authStore.setUser(result.data.verifyEmail.id);
  authStore.setRole(result.data.verifyEmail.role);
  verificationSuccess.value = false;
  navigateTo("/");
});

onMounted(async () => {
  const token = route.query.token;

  if (!token) {
    verificationSuccess.value = false;
    isVerifying.value = false;
    return;
  }
  mutate({ token });
});
</script>

<style scoped>
/* Add custom styles here */
</style>
