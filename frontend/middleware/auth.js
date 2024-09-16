import { useAuthStore } from "~/store";

export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();
  authStore.autoLogin();
  
  const meta = to.meta || {};
  
  // Redirect based on authentication and role
  if (authStore.isAuthenticated) {
    if (authStore.role === 'user-admin') {
      if (to.path === '/') {
        return navigateTo('/user-admin/dashboard'); // Redirect user-admin from '/' to dashboard
      }
    } else if (authStore.role === 'user') {
      if (to.path === '/user-admin/dashboard') {
        return navigateTo('/'); // Redirect user from user-admin dashboard to '/'
      }
    }
  }

  // Check if the route requires authentication
  if (meta.requiresAuth && !authStore.isAuthenticated) {
    return navigateTo("/login");
  }

  // Check if the route requires specific roles
  if (meta.allowedRoles && !meta.allowedRoles.includes(authStore.role)) {
    return navigateTo("/"); // Redirect to a default page or unauthorized page
  }
});


