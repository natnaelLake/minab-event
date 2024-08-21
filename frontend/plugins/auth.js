// // middleware/auth.js

// import { useAuthStore } from "~/store";

// export default defineNuxtRouteMiddleware((to, from) => {
//   const authStore = useAuthStore();

//   // Auto-login if token is present
//   authStore.autoLogin();

//   // Check if the route requires authentication
//   if (to.meta.requiresAuth && !authStore.isAuthenticated) {
//     return navigateTo("/login");
//   }

//   // Check if the route requires specific roles
//   if (to.meta.allowedRoles && !to.meta.allowedRoles.includes(authStore.role)) {
//     return navigateTo("/"); // Redirect to a default page or unauthorized page
//   }
// });
