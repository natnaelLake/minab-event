import { ref } from "vue";
import { useApolloClient } from "@vue/apollo-composable";
import { useCookie } from "#app"; // Nuxt 3 cookies
import { useAuthStore } from "~/store";

// Define types for the GraphQL responses
type LoginResponse = {
  login: {
    token: string;
    id: string;
    email: string;
    full_name: string;
    profile_image_url: string | null;
  };
};

type SignupResponse = {
  signup: {
    token: string;
    id: string;
    email: string;
    full_name: string;
    profile_image_url: string | null;
  };
};

export function useAuth() {
  const tokenCookie = useCookie<string>("auth_token"); // Token managed in cookies
  const store = useAuthStore(); // Use the Pinia store
  const { client } = useApolloClient();

  // Login Mutation
  const login = async (
    mutation: any,
    UserData: { email: string; password: string }
  ) => {
    try {
      const { data } = await client.mutate<LoginResponse>({
        mutation: mutation,
        variables: UserData,
      });

      const token = data?.login.token;
      const id = data?.login.id;

      // Update the Pinia store
      store.setToken(token);
      store.setId(id);
      store.autoLogin(); // Optionally auto-login to set user details

      tokenCookie.value = token as string; // Set token in cookie
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  // Signup Mutation
  const signup = async (
    mutation: any,
    UserData: {
      email: string;
      password: string;
      full_name: string;
      profile_image_url: string;
    }
  ) => {
    try {
      const { data } = await client.mutate<SignupResponse>({
        mutation: mutation,
        variables: UserData,
      });

      const token = data?.signup.token;
      const id = data?.signup.id;

      // Update the Pinia store
      store.setToken(token);
      store.setId(id);
      store.autoLogin(); // Optionally auto-login to set user details

      tokenCookie.value = token as string; // Set token in cookie
    } catch (error) {
      console.error("Signup failed:", error);
    }
  };

  // Logout function
  const logout = () => {
    // Clear the Pinia store
    store.logout();

    // Clear token from cookie
    tokenCookie.value = "";
  };

  return {
    isAuthenticated: store.isAuthenticated, // Bind to the store's getter
    login,
    signup,
    logout,
  };
}
