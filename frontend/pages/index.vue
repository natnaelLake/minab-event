<template>
  <div class="flex mt-20">
    <!-- Filter Sidebar -->
    <div
      class="w-64 bg-gray-100 p-2 pb-16 fixed top-20 left-0 h-screen overflow-scroll"
    >
      <Filter />
    </div>

    <div class="flex-1 ml-64 p-6">
      <!-- Search Bar -->
      <div class="mb-6">
        <input
          type="text"
          placeholder="Search events..."
          class="w-full border border-gray-300 rounded-lg px-4 py-2"
        />
      </div>

      <!-- Event Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="event in events"
          :key="event.id"
          class="bg-white text-black shadow-lg rounded-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer"
        >
          <EventCard
            :header="{
              avatar: event.ownerAvatar,
              name: event.ownerName,
              actions: event.headerActions,
            }"
            :image="{ src: event.imageSrc, alt: event.imageAlt }"
            :title="event.title"
            :description="event.description"
            :address="event.address"
            :price="event.price"
            :deadline="event.deadline"
            :footer="{
              reserve: event.reserve,
              postTime: { text: event.postTime },
            }"
            @navigate="goToEventDetail(event.id)"
            class="dark:bg-white text-black"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import EventCard from "~/components/EventCard.vue";
import { useRouter } from "vue-router";
import Filter from "~/components/Filter.vue";
import '@fortawesome/fontawesome-free/css/all.css';

const router = useRouter();

const goToEventDetail = (eventId) => {
  router.push(`/event/${eventId}`);
};

const events = [
  {
    id: 1,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "John Doe",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 1",
    title: "Music Concert",
    description:
      "Join us for a night of amazing music and fun! fdsf dsfdsfdsfdsrev vdfeerf dfewresfds fdsrewfds ",
    price: "$50",
    address: "123 Music Lane, Cityville",
    deadline: "2024-09-15",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 2 hours ago",
  },
  {
    id: 2,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "John Doe",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 2",
    title: "Art Exhibition",
    description: "Explore a diverse range of artistic expressions.",
    price: "$30",
    address: "456 Art St, Gallery City",
    deadline: "2024-10-01",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 3 hours ago",
  },
  {
    id: 3,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "Jane Smith",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 3",
    title: "Food Festival",
    description: "Taste delicious dishes from around the world.",
    price: "$20",
    address: "789 Culinary Ave, Foodtown",
    deadline: "2024-09-30",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 4 hours ago",
  },
  {
    id: 4,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "Alex Johnson",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 4",
    title: "Tech Conference",
    description:
      "Join industry leaders in discussing the future of technology.",
    price: "$100",
    address: "101 Tech Park, Innovate City",
    deadline: "2024-11-15",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 5 hours ago",
  },
  {
    id: 5,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "Emily Davis",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 5",
    title: "Outdoor Adventure",
    description:
      "Experience the thrill of nature with guided outdoor activities.",
    price: "$80",
    address: "202 Adventure Rd, Nature City",
    deadline: "2024-10-20",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 6 hours ago",
  },
  {
    id: 6,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "Mark Lee",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 6",
    title: "Comedy Show",
    description: "Laugh the night away with our lineup of top comedians.",
    price: "$40",
    address: "303 Comedy Lane, Laugh City",
    deadline: "2024-09-25",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 7 hours ago",
  },
  {
    id: 7,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "Sarah Brown",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 7",
    title: "Yoga Retreat",
    description: "Rejuvenate your mind and body with our yoga retreat.",
    price: "$120",
    address: "404 Wellness Blvd, Serenity City",
    deadline: "2024-11-01",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 8 hours ago",
  },
  {
    id: 8,
    ownerAvatar:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s",
    ownerName: "Linda Martinez",
    headerActions: [
      {
        icon: "🔖",
        handler: () => alert("Bookmark clicked!"),
        class: "text-gray-600",
      },
      {
        icon: "⭐",
        handler: () => alert("Follow clicked!"),
        class: "text-yellow-500",
      },
    ],
    imageSrc:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s",
    imageAlt: "Event Image 8",
    title: "Craft Fair",
    description: "Discover unique handmade crafts and artworks.",
    price: "$25",
    address: "505 Craft Ave, Art City",
    deadline: "2024-09-22",
    reserve: {
      handler: () => alert("Reserve clicked!"),
      class: "bg-blue-500 text-white px-4 py-2 rounded",
    },
    postTime: "Posted 9 hours ago",
  },
];
</script>
<style scoped>
/* Make sure the filter and search bar are non-scrollable */
.fixed {
  position: fixed;
  height: calc(100vh - 20px); /* Full viewport height minus top margin */
  overflow-y: auto; /* Enable vertical scrolling */
}

.ml-64 {
  margin-left: 16rem; /* Adjust to the width of your filter */
}

.search-bar {
  position: relative;
  z-index: 10; /* Ensure it's above the event cards */
}

.event-cards {
  margin-left: 16rem; /* Same as filter width */
  padding: 1.5rem; /* Adjust padding as needed */
}

@media (max-width: 768px) {
  .ml-64 {
    margin-left: 0;
  }

  .fixed {
    position: relative;
    height: auto;
  }
}
</style>
