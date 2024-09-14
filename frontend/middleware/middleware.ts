// middleware.ts
import authMiddleware from './auth';

export default defineNuxtRouteMiddleware((to, from) => {
  authMiddleware(to, from);
});
