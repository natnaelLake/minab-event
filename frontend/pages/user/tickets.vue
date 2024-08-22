<template>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 pt-28 p-10">
      <EventCard
        v-for="event in events"
        :key="event.id"
        :header="{ avatar: event.ownerAvatar, name: event.ownerName, actions: event.headerActions }"
        :image="{ src: event.imageSrc, alt: event.imageAlt }"
        :title="event.title"
        :description="event.description"
        :footer="{ reserve: event.reserve, postTime: { text: event.postTime } }"
      />
    </div>
  </template>
  <script setup>
  import EventCard from '~/components/EventCard.vue';
  
  
  const events = [
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun! fdsf dsfdsfdsfdsrev vdfeerf dfewresfds fdsrewfds ',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
    {
      id: 1,
      ownerAvatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5_JC2T02lxtB6JgjqZXbWkE-jiHGaejssAw&s',
      ownerName: 'John Doe',
      headerActions: [
        { icon: '🔖', handler: () => alert('Bookmark clicked!'), class: 'text-gray-600' },
        { icon: '⭐', handler: () => alert('Follow clicked!'), class: 'text-yellow-500' },
      ],
      imageSrc: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwC_a-pMFNFKn0INHjDNq_y1vggkf2VRRYNg&s',
      imageAlt: 'Event Image 1',
      title: 'Music Concert',
      description: 'Join us for a night of amazing music and fun!',
      reserve: { handler: () => alert('Reserve clicked!'), class: 'bg-blue-500 text-white px-4 py-2 rounded' },
      postTime: 'Posted 2 hours ago',
    },
  ];
  </script>
  