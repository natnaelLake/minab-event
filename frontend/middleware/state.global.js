import { useAuthStore } from "@/store";
export default defineNuxtRouteMiddleware((to, from, next) => {
  const authStore = useAuthStore();
  if (authStore.isAuthenticated) {
    authStore.autoLogin();
  }
  return;
});
