// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  ssr: false,
  modules: [
    "@nuxtjs/tailwindcss",
    "@nuxtjs/apollo",
    "@pinia/nuxt",
    "nuxt-rating",
    "nuxt-icon",
    "@vee-validate/nuxt",
    "@vueuse/nuxt",
    "@nuxtjs/color-mode",
    "@nuxt/image",
    "@nuxtjs/leaflet",
  ],
  leaflet: {
    markerCluster: true,
  },
  colorMode: {
    classSuffix: "",
  },

  apollo: {
    autoImports: true,
    authType: "Bearer",
    authHeader: "Authorization",
    tokenStorage: "cookie",
    proxyCookies: true,
    clients: {
      default: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
      },
      authClient: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
        tokenName: "auth_token",
      },
    },
  },

  veeValidate: {
    autoImports: true,
    componentNames: {
      Form: "VeeForm",
      Field: "VeeField",
      FieldArray: "VeeFieldArray",
      ErrorMessage: "VeeErrorMessage",
    },
  },
  image: {
    inject: true,
    quality: 80,
  },
  plugins: [{ src: "~/plugins/vue3-toastify.ts" }],
  compatibilityDate: "2024-08-01",
});
