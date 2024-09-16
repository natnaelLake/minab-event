<script setup>
import { ref, onMounted } from "vue";
import GetMyBookMarkedEvents from "~/graphql/query/GetMyEventBookMarks.gql";
import UNBookMarkEvent from "~/graphql/mutations/UNBookMarkEvent.gql";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { useRouter } from "vue-router";
import { toast } from "vue3-toastify";
import { useAuthStore } from "~/store";
import { format, formatDistance, parseISO } from "date-fns";
import useQueryComposable from "~/composables/useQueryComposable";
import useMutationComposable from "~/composables/useMutationComposable";

const { mutate: unBookMarkEvent } = useMutationComposable(UNBookMarkEvent);
const router = useRouter();
const user = useAuthStore();
const currentUser = user.id;
const bookmarks = ref([]);

onMounted(async () => {
  await fetchMyBookMarkedEvents();
});

const fetchMyBookMarkedEvents = async () => {
  try {
    const { onResult, onError } = useQueryComposable(GetMyBookMarkedEvents, {
      limit: 10,
      offset: 0,
      order_by: [{ created_at: "desc" }],
      user_id: currentUser,
    });

    onResult((result) => {
      if (result.data) {
        bookmarks.value = result.data.bookmarks;
      }
    });

    onError((error) => {
      console.error("Error fetching events: ", error.message);
      toast.error("Something went wrong, try again");
    });
  } catch (error) {
    console.error("Error during fetching events: ", error);
    toast.error("Failed to load events.");
  }
};
const formatEventDate = (date) => {
  if (!date) {
    console.error("Invalid date value:", date);
    return "Invalid Date";
  }

  try {
    const parsedDate = parseISO(date);
    return format(parsedDate, "PPpp");
  } catch (error) {
    console.error("Error parsing date:", error);
    return "Invalid Date";
  }
};
const goToEventDetail = (eventId) => {
  router.push(`/event/${eventId}`);
};
const handleFormatDistance = (date) => {
  return formatDistance(new Date(date), new Date(), { addSuffix: true });
};

const removeBookmark = async (eventId, bookMarkId) => {
  await unBookMarkEvent({
    event_id: eventId,
    user_id: currentUser,
  });
  bookmarks.value = bookmarks.value.filter(
    (bookmark) => bookmark.id !== bookMarkId
  );
  toast.success("Event UnBookMarked successfully.");
};
</script>

<template>
  <div
    class="flex flex-col items-center w-full justify-center min-h-screen px-4 sm:px-6 lg:px-8 py-10 mt-20"
  >
    <div class="w-full max-w-4xl">
      <h3
        class="text-3xl font-bold mb-8 text-center text-gray-900 dark:text-gray-900"
      >
        Your Bookmarks
      </h3>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div
          v-for="bookmark in bookmarks"
          :key="bookmark.id"
          class="bg-white dark:bg-gray-1'*00 p-5 rounded-lg shadow-lg transform transition-transform hover:-translate-y-2 hover:shadow-xl cursor-pointer"
          @click="goToEventDetail(bookmark.event.id)"
        >
          <img
            :src="
              bookmark.event.featured_image_url.split(',')[0].replace('{', '')
            "
            alt="bookmark-thumbnail"
            class="w-full h-32 sm:h-40 rounded-md object-cover mb-4"
          />
          <h4
            class="text-lg font-semibold text-gray-800 dark:text-gray-800 mb-2"
          >
            {{ bookmark.event.title }}
          </h4>
          <p class="text-sm text-gray-500 dark:text-gray-400 mb-4">
            {{ bookmark.event.description }}
          </p>
          <div class="flex justify-between items-center">
            <span class="text-xs text-gray-400 dark:text-gray-500">{{
              handleFormatDistance(bookmark.created_at)
            }}</span>
            <button
              class="text-red-500 hover:text-red-600 transition-colors duration-200"
              @click.stop="removeBookmark(bookmark.event.id, bookmark.id)"
            >
              Remove
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
