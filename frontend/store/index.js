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
    isAdmin: (state) => state.role === "user-admin",
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
      if (this.token) {
        try {
          const decoded = jwt_decode(this.token);
          const userId =
            decoded["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
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
