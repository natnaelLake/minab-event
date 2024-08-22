<template>
  <div class="flex mt-20">
    <!-- Product Image and Thumbnails Section -->
    <div class="flex-grow max-w-6xl ml-20 p-6 bg-white shadow-lg rounded-lg">
      <div class="flex">
        <!-- Image Carousel with Thumbnails -->
        <div class="w-1/4 pr-4">
          <div
            ref="thumbnailContainer"
            class="flex flex-col space-y-2 overflow-y-auto h-full max-h-[calc(100vh-9rem)] scrollbar-hidden"
          >
            <img
              v-for="(image, index) in images"
              :key="index"
              :src="image"
              alt="Thumbnail"
              class="w-20 h-20 object-cover rounded-lg cursor-pointer border-2 transition-all duration-300 ease-in-out hover:border-blue-400"
              :class="{ 'border-blue-500': currentIndex === index }"
              @click="changeImage(index)"
            />
          </div>
        </div>

        <!-- Large Product Image -->
        <div class="w-3/4">
          <img
            :src="currentImage"
            alt="Product Image"
            class="w-full h-[30rem] object-cover rounded-lg shadow-lg"
          />
        </div>
      </div>
    </div>

    <!-- Product Owner Card -->
    <div class="w-1/4 pl-6">
      <div class="flex flex-col items-start p-4 bg-white shadow-lg rounded-lg">
        <!-- Avatar -->
        <img
          :src="avatarUrl"
          alt="Product Owner Avatar"
          class="w-24 h-24 rounded-full object-cover border-4 border-gray-300"
        />

        <!-- Product Owner Details -->
        <div class="mt-4">
          <h2 class="text-2xl font-bold text-gray-900">{{ ownerName }}</h2>
          <p class="text-sm text-gray-600 mt-1">Product Owner</p>
        </div>

        <!-- Follow Button -->
        <button
          @click="followOwner"
          class="mt-4 bg-blue-600 text-white px-5 py-2 rounded-lg font-semibold hover:bg-blue-700 transition-colors duration-300"
        >
          Follow
        </button>
      </div>
    </div>
  </div>

  <!-- Product Details Section -->
  <div class="flex max-w-6xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-8">
    <div class="flex-grow mr-8">
      <!-- Product Title and Type/Category -->
      <div class="mb-6">
        <h1 class="text-4xl font-bold text-gray-800">{{ productTitle }}</h1>
        <div class="mt-2 flex space-x-2">
          <span
            class="bg-blue-100 text-blue-800 text-sm font-medium px-3 py-1 rounded-full"
          >
            {{ productType }}
          </span>
          <span
            class="bg-green-100 text-green-800 text-sm font-medium px-3 py-1 rounded-full"
          >
            {{ productCategory }}
          </span>
        </div>
      </div>

      <!-- Tags -->
      <div class="mb-6 flex flex-wrap gap-2">
        <span
          v-for="(tag, index) in tags"
          :key="index"
          class="bg-gray-200 text-gray-800 text-xs font-medium px-3 py-1 rounded-full cursor-pointer hover:bg-gray-300"
        >
          {{ tag }}
        </span>
      </div>

      <!-- Description -->
      <div class="mb-6">
        <h2 class="text-xl font-semibold text-gray-700">Description</h2>
        <p class="text-gray-600 mt-2" :class="{ 'line-clamp-3': !isExpanded }">
          {{ productDescription }}
        </p>
        <button
          @click="toggleDescription"
          class="text-blue-600 mt-2 inline-block"
        >
          {{ isExpanded ? "Show Less" : "Show More" }}
        </button>
      </div>

      <!-- Additional Details -->
      <div class="mb-6">
        <div class="flex items-center justify-between mb-4">
          <span class="text-3xl font-bold text-gray-800">${{ price }}</span>
          <div class="text-gray-600">
            <p class="text-sm">Available: {{ availableAmount }}</p>
            <p class="text-sm">Sold: {{ soldAmount }}</p>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
                class="bg-blue-500 h-2 rounded-full"
                :style="{ width: soldPercentage + '%' }"
              ></div>
            </div>
          </div>
        </div>

        <div class="text-sm text-gray-500 mb-4">
          <p>Posted on: {{ postDate }}</p>
        </div>

        <div class="flex items-center space-x-4">
          <span class="text-gray-700">Likes: {{ likes }}</span>
          <a href="#" class="text-blue-600">Reviews ({{ reviews }})</a>
        </div>
      </div>
    </div>

    <!-- Checkout Section -->
    <div class="fixed right-0 bottom-0 p-6 bg-white shadow-lg rounded-lg border-t border-gray-200 z-40">
      <h2 class="text-xl font-bold text-gray-800 mb-4">Checkout</h2>
      <div class="flex items-center justify-between mb-4">
        <span class="text-2xl font-bold text-gray-800">${{ price }}</span>
        <button
          class="bg-blue-600 text-white px-5 py-2 rounded-lg font-semibold hover:bg-blue-700 transition-colors duration-300"
        >
          Buy Now
        </button>
      </div>
      <div class="text-sm text-gray-600">
        <p>Available: {{ availableAmount }}</p>
        <p>Sold: {{ soldAmount }}</p>
      </div>
    </div>
  </div>

  <!-- Related Products Section -->
  <div class="max-w-6xl mx-auto p-6 mt-12">
    <h2 class="text-3xl font-bold text-gray-800 mb-8">Related Products</h2>
    
    <!-- User's Other Products -->
    <div class="mb-12">
      <h3 class="text-2xl font-semibold text-gray-700 mb-4">User's Other Products</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="(product, index) in usersOtherProducts" :key="index" class="bg-white shadow-lg rounded-lg overflow-hidden transform hover:scale-105 transition-transform duration-300 ease-in-out">
          <img :src="product.image" alt="User Product" class="w-full h-48 object-cover"/>
          <div class="p-4">
            <h4 class="text-xl font-bold text-gray-800">{{ product.title }}</h4>
            <p class="text-gray-600 mt-2">${{ product.price }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Similar Products -->
    <div>
      <h3 class="text-2xl font-semibold text-gray-700 mb-4">Similar Products</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="(product, index) in similarProducts" :key="index" class="bg-white shadow-lg rounded-lg overflow-hidden transform hover:scale-105 transition-transform duration-300 ease-in-out">
          <img :src="product.image" alt="Similar Product" class="w-full h-48 object-cover"/>
          <div class="p-4">
            <h4 class="text-xl font-bold text-gray-800">{{ product.title }}</h4>
            <p class="text-gray-600 mt-2">${{ product.price }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// Data for related products
const usersOtherProducts = ref([
  { image: "https://via.placeholder.com/300", title: "Product A", price: "199.99" },
  { image: "https://via.placeholder.com/300", title: "Product B", price: "249.99" },
  { image: "https://via.placeholder.com/300", title: "Product C", price: "299.99" },
]);

const similarProducts = ref([
  { image: "https://via.placeholder.com/300", title: "Similar Product A", price: "179.99" },
  { image: "https://via.placeholder.com/300", title: "Similar Product B", price: "229.99" },
  { image: "https://via.placeholder.com/300", title: "Similar Product C", price: "279.99" },
]);

// Existing code remains unchanged
const images = [
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQUyAfXfniYfSTZ7Z2HjW2COSyC8WTH3TgkGw&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVNN58XFDLxdqtwwWRSE924NjtuSryXFGxjg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRezkhZEGyU-SbkR5X1RGxo8OxQFLfKonocyg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNNRqKoghIDo3Vefq59_pyxm3TMTFnMYzOzg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZO-2Hz9rOfjRXIwyKMj_QgNX-LktjjEW4Og&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQUyAfXfniYfSTZ7Z2HjW2COSyC8WTH3TgkGw&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVNN58XFDLxdqtwwWRSE924NjtuSryXFGxjg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRezkhZEGyU-SbkR5X1RGxo8OxQFLfKonocyg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNNRqKoghIDo3Vefq59_pyxm3TMTFnMYzOzg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZO-2Hz9rOfjRXIwyKMj_QgNX-LktjjEW4Og&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQUyAfXfniYfSTZ7Z2HjW2COSyC8WTH3TgkGw&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVNN58XFDLxdqtwwWRSE924NjtuSryXFGxjg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRezkhZEGyU-SbkR5X1RGxo8OxQFLfKonocyg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNNRqKoghIDo3Vefq59_pyxm3TMTFnMYzOzg&s",
  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZO-2Hz9rOfjRXIwyKMj_QgNX-LktjjEW4Og&s",
  // ... other images
];

const currentIndex = ref(0);
const currentImage = ref(images[currentIndex.value]);
const avatarUrl = ref("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVNN58XFDLxdqtwwWRSE924NjtuSryXFGxjg&s");
const ownerName = ref("John Doe");

function changeImage(index) {
  currentIndex.value = index;
  currentImage.value = images[index];
}

function followOwner() {
  alert("Followed " + ownerName.value);
}

const productTitle = ref("Amazing Product");
const productType = ref("Electronics");
const productCategory = ref("Gadgets");
const tags = ref(["Smart", "Portable", "Tech"]);
const productDescription = ref("This is a detailed description of the product. It can be very long, so we provide an option to expand or collapse the text to make the page cleaner and more user-friendly. The text describes all the features and benefits of the product in detail.");
const price = ref("299.99");
const availableAmount = ref("120");
const soldAmount = ref("45");
const soldPercentage = ref(37); // Percentage of sold amount
const postDate = ref("August 22, 2024");
const likes = ref("123");
const reviews = ref("45");
const isExpanded = ref(false);

function toggleDescription() {
  isExpanded.value = !isExpanded.value;
}
</script>

<style scoped>
/* Hide scrollbar */
.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
.scrollbar-hidden {
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>

  