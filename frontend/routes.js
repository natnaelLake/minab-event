export const routes = [
  // {
  //   name: 'MinabEvent',
  //   path: '/',
  //   component: () => import('@/pages/index.vue'),
  //   meta: {
  //     requiresAuth: false,
  //   },
  // },
  {
    name: 'Login',
    path: '/login',
    component: () => import('@/pages/login.vue'),
    meta: {
      requiresAuth: false,
    },
  },
  {
    name: 'Sign Up',
    path: '/signup',
    component: () => import('@/pages/signup.vue'),
    meta: {
      requiresAuth: false,
    },
  },
  {
    name: 'My Events',
    path: '/user/my-events',
    component: () => import('@/pages/user/my-events.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['user'],
    },
  },
  {
    name: 'Ticket Reservation',
    path: '/user/tickets',
    component: () => import('@/pages/user/tickets.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['user'],
    },
  },
  {
    name: 'Bookmarks',
    path: '/user/bookmarks',
    component: () => import('@/pages/user/bookmarks.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['user'],
    },
  },
  {
    name: 'Settings',
    path: '/user/settings',
    component: () => import('@/pages/user/settings.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['user', 'admin'],
    },
  },
  
  {
    name: 'Dashboard',
    path: '/admin/dashboard',
    component: () => import('@/pages/admin/dashboard.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin'],
    },
  },
  {
    name: 'Event Management',
    path: '/admin/events',
    component: () => import('@/pages/admin/events.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin'],
    },
  },
  {
    name: 'Manage Users',
    path: '/admin/users',
    component: () => import('@/pages/admin/users.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['admin'],
    },
  },
  {
    name: 'Logout',
    path: '/logout',
    component: () => import('@/pages/logout.vue'),
    meta: {
      requiresAuth: true,
      allowedRoles: ['user', 'admin'],
    },
  },
];

export function getNavigationRoutes(role, isAuthenticated) {
  return routes.filter(route => {
    console.log('*****************',role,isAuthenticated)
    // Exclude Login and Sign Up routes if the user is authenticated
    if (isAuthenticated && (route.name === 'Login' || route.name === 'Sign Up')) {
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
