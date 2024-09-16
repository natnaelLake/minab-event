export default function (mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "network-only",
    context: {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "x-hasura-role": "user", 
      },
    },
  }));

  return {
    mutate,
    onDone,
    loading,
    onError,
  };
}