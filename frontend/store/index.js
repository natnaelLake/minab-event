import { defineStore } from "pinia";
import Cookies from "js-cookie";
import jwt_decode from "jwt-decode";
import get_single_user from "@/graphql/query/getSingleUser.gql";
import authQuery from "@/composables/auth";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: Cookies.get("auth_token") || null,
    id: null,
    role: null,
    user: null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => state.role === "admin",
    isUser: (state) => state.role === "user", // Adjust according to your roles
  },

  actions: {
    setToken(token) {
      this.token = token;
    },

    setRole(role) {
      this.role = role;
    },

    setId(id) {
      console.log("*******iiiiiii", id);
      this.id = id;
    },

    logout() {
      this.token = null;
      this.user = null;
      this.id = null;
      this.role = null;
      Cookies.remove("auth_token");
    },

    async setUser(id) {
      const { onResult, onError } = authQuery(get_single_user, {
        id,
      });

      onResult((result) => {
        console.log("User fetched:", result);
        if (result.data) {
          this.user = { ...result.data.users_by_pk };
          this.role = this.user.role;
        }
      });

      onError((error) => {
        console.log("Error fetching user:", error);
      });
    },

    autoLogin() {
      console.log("Token:", this.token);
      if (this.token) {
        try {
          const decoded = jwt_decode(this.token);
          console.log("Decoded JWT:", decoded);

          // Extracting user ID from claims
          // Example: Extract user ID if it's in a specific format
          const userIdClaim =
            decoded["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

          // If userIdClaim contains additional characters or needs parsing, handle it here
          // For example, if the real ID is after '=' sign

          const userId =
            userIdClaim.split("=")[1].replace(/\)$/, "") || userIdClaim;

          console.log("User ID:", userId);

          // Validating token expiration
          if (decoded.exp * 1000 > Date.now()) {
            this.setId(userId);
            this.setUser(userId);
          } else {
            this.logout();
          }
        } catch (error) {
          console.error("Token decoding failed:", error);
          this.logout();
        }
      } else {
        this.logout();
      }
    },
  },
});
