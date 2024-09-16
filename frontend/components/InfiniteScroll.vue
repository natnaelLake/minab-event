<!-- components/InfiniteScroll.vue -->
<template>
    <div>
      <div v-for="post in posts" :key="post.id" class="card bg-white p-4 mb-4 shadow-md rounded">
        <h3 class="text-xl font-bold">{{ post.title }}</h3>
        <p class="text-gray-600">{{ post.body }}</p>
      </div>
      <infinite-loading
        @infinite="loadMore"
        :spinner="spinner"
        :distance="10"
      >
        <template #no-more>
          No more data
        </template>
        <template #error>
          Error loading data
        </template>
      </infinite-loading>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import InfiniteLoading from 'vue3-infinite-loading';
//   import 'vue3-infinite-loading/dist/vue3-infinite-loading.css';
  
  // Fetch posts from JSONPlaceholder
  const fetchPosts = (page, limit) => {
    return fetch(`https://jsonplaceholder.typicode.com/posts?_page=${page}&_limit=${limit}`)
      .then(response => response.json());
  };
  
  export default {
    components: {
      InfiniteLoading
    },
    data() {
      return {
        posts: [],  // Array to hold your data items
        page: 1,
        limit: 10
      };
    },
    methods: {
      async loadMore($state) {
        try {
          // Fetch data from JSONPlaceholder
          const newPosts = await fetchPosts(this.page, this.limit);
          
          if (newPosts.length) {
            this.posts.push(...newPosts);
            this.page++;
            $state.loaded();
          } else {
            $state.complete();
          }
        } catch (error) {
          $state.error();
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .card {
    @apply bg-white p-4 mb-4 shadow-md rounded;
  }
  </style>
  