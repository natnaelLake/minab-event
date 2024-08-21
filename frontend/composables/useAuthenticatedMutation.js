// // for authentcated user
// export default function (mutation) {
// 	const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
// 		fetchPolicy: "network-only",
// 		clientId: "authClient",
// 		context: {
// 			headers: {
// 				"x-hasura-role": "user",
// 			},
// 		},
// 	}));

// 	return {
// 		onError,
// 		mutate,
// 		loading,
// 		onDone,
// 	};
// }
export default function (mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "network-only",
    context: {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`, // Get the token from wherever it is stored
        "x-hasura-role": "user", // Adjust this if your role is different or dynamic
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