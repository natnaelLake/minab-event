// middleware/auth.js
import { useAuthStore } from "~/store";

export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();

  // Auto-login if token is present
  authStore.autoLogin();

  // Handle undefined meta
  const meta = to.meta || {};
  
  // Check if the route requires authentication
  if (meta.requiresAuth && !authStore.isAuthenticated) {
    return navigateTo("/login");
  }

  // Check if the route requires specific roles
  if (meta.allowedRoles && !meta.allowedRoles.includes(authStore.role)) {
    return navigateTo("/"); // Redirect to a default page or unauthorized page
  }
});
