// for authentcation login , register and reset password and forgot password
// this mutation are executed with default client or client with no token
import { useMutation } from "@vue/apollo-composable";

export default function (mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "network-only",
  }));
  return {
    onError,
    mutate,
    loading,
    onDone,
  };
}
