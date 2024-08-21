import { useAuthStore } from "@/store";
export default defineNuxtRouteMiddleware((to, from, next) => {
  const authStore = useAuthStore();
  console.log('+++++++++++++++++',authStore);
  console.log("setup.global.js this is the global middleware");
  if (authStore.isAuthenticated) {
    authStore.autoLogin();
  }
  return;
});
