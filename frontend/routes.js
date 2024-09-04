export const routes = [
  {
    path: "/",
    component: () => import("@/pages/index.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user"],
    },
  },
  {
    name: "Login",
    path: "/login",
    component: () => import("@/pages/login.vue"),
    meta: {
      requiresAuth: false,
    },
  },
  {
    name: "Sign Up",
    path: "/signup",
    component: () => import("@/pages/signup.vue"),
    meta: {
      requiresAuth: false,
    },
  },
  {
    name: "My Events",
    path: "/user/my-events",
    component: () => import("@/pages/user/my-events.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user"],
    },
  },
  {
    name: "Ticket Reservation",
    path: "/user/tickets",
    component: () => import("@/pages/user/tickets.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user"],
    },
  },
  {
    name: "Bookmarks",
    path: "/user/bookmarks",
    component: () => import("@/pages/user/bookmarks.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user"],
    },
  },
  {
    name: "Dashboard",
    path: "/user-admin/dashboard",
    component: () => import("@/pages/user-admin/dashboard.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user-admin"],
    },
  },
  {
    name: "Event Management",
    path: "/user-admin/event-management",
    component: () => import("@/pages/user-admin/event-management.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user-admin"],
    },
  },
  {
    name: "Manage Users",
    path: "/user-admin/user-management",
    component: () => import("@/pages/user-admin/event-management.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user-admin"],
    },
  },
  {
    name: "Tasks Management",
    path: "/user-admin/TagsCategoriesManagement",
    component: () => import("@/pages/user-admin/TagsCategoriesManagement.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user-admin"],
    },
  },
  {
    name: "Supports",
    path: "/user-admin/supports",
    component: () => import("@/pages/user-admin/supports.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user-admin"],
    },
  },
  {
    name: "Reports",
    path: "/user-admin/reports",
    component: () => import("@/pages/user-admin/reports.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user-admin"],
    },
  },
  {
    name: "Settings",
    path: "/user/settings",
    component: () => import("@/pages/user/settings.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user", "user-admin"],
    },
  },
  {
    name: "Logout",
    path: "/logout",
    component: () => import("@/pages/logout.vue"),
    meta: {
      requiresAuth: true,
      allowedRoles: ["user", "user-admin"],
    },
  },
];

export function getNavigationRoutes(role, isAuthenticated) {
  return routes.filter((route) => {
    if (
      isAuthenticated &&
      (route.name === "Login" || route.name === "Sign Up")
    ) {
      return false;
    }
    if (route.meta.requiresAuth && !isAuthenticated) {
      return false;
    }
    if (route.meta.allowedRoles && !route.meta.allowedRoles.includes(role)) {
      return false;
    }
    return true;
  });
}
