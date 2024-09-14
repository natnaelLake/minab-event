import { useQuery } from "@vue/apollo-composable";
import { useAuthStore } from "~/store";

/**
 * 
 * 
 * @param {Object} variables 
 * @param {Object} queryDocument 
 * @param {Object} additionalOptions 
 * 
 * @returns {Object} 
 */
export default function useQueryComposable(queryDocument,variables, additionalOptions = {}) {
  const user = useAuthStore();
  const currentUser = user.id;
  const currentUserRole = user.role;
  const token = user.token;
  const defaultContext = {
    headers: {
      "x-hasura-user-id": currentUser || "",
      "x-hasura-role": currentUserRole || "anonymous",
      ...(token && { Authorization: `Bearer ${token}` }),
    },
  };

  // Merge default context with additional options
  const mergedOptions = {
    fetchPolicy: "network-only",
    clientId: additionalOptions.clientId || "default",  // Updated clientId to match config
    context: {
      ...defaultContext,
      ...(additionalOptions.context || {}),
    },
  };

  // Wrap variables in ref for reactivity
  const queryVariables = ref({ ...variables });

  // Use Apollo query with dynamic variables and merged options
  const { onResult, loading, refetch, onError } = useQuery(
    queryDocument,
    queryVariables,
    () => mergedOptions
  );

  // Handle errors
  onError((error) => {
    console.error("GraphQL Query Error:", error);
  });

  return {
    onError,
    loading,
    refetch,
    onResult,
  };
}
