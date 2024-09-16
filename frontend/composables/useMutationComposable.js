import { useMutation } from "@vue/apollo-composable";
import { useAuthStore } from "~/store";

/**
 * Reusable composable for Apollo mutations with default auth context.
 * 
 * @param {Object} mutationDocument - The GraphQL mutation to execute.
 * @param {Object} additionalOptions - Additional options (e.g., headers, fetchPolicy, etc.).
 * 
 * @returns {Object} - Returns mutate function, done handler, loading state, and error handler.
 */
export default function useMutationComposable(mutationDocument, additionalOptions = {}) {
  const user = useAuthStore();
  const currentUser = user.id;
  const currentUserRole = user.role;
  const token = user.token;

  // Default context with headers
  const defaultContext = {
    headers: {
      "x-hasura-user-id": currentUser,
      "x-hasura-role": currentUserRole,
      Authorization: `Bearer ${token}`,
    },
  };

  // Merge default context with additional options
  const mergedOptions = {
    fetchPolicy: "network-only",
    context: {
      ...defaultContext,
      ...(additionalOptions.context || {}),
    },
  };

  // Use Apollo mutation with merged options
  const { mutate, onDone, loading, error } = useMutation(mutationDocument, () => mergedOptions);

  return {
    mutate,
    onDone,
    loading,
    error,
  };
}
