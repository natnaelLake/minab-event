export default (id, getData) => {
	const { onResult, loading, refetch, onError } = useQuery(
		getData,
		() => ({
			id: id,
		}),
		() => ({
			fetchPolicy: "network-only",
		})
	);

	return {
		onResult,
		loading,
		refetch,
		onError,
	};
};
