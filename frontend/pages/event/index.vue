<template>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <Card
        v-for="event in events"
        :key="event.id"
        :header="{
          avatar: event.avatar,
          name: event.ownerName,
          bookmarkHandler: () => handleBookmark(event.id),
          followHandler: () => handleFollow(event.ownerId),
        }"
        :image="{ src: event.imageSrc, alt: event.title }"
        :title="event.title"
        :description="event.description"
        :actions="{
          reserveHandler: () => handleReserve(event.id),
          postTime: event.postTime,
        }"
      />
    </div>
  </template>
  
  <script setup>
import Card from '~/components/Card.vue';

  
  const events = [
    {
      id: 1,
      avatar: 'avatar1.jpg',
      ownerName: 'Owner 1',
      imageSrc: 'event1.jpg',
      title: 'Event 1',
      description: 'This is a brief description of Event 1.',
      postTime: '2 hours ago',
      ownerId: 101,
    },
  ];
  
  const handleBookmark = (eventId) => {
    alert(`Bookmark event ${eventId}`);
  };
  
  const handleFollow = (ownerId) => {
    alert(`Follow owner ${ownerId}`);
  };
  
  const handleReserve = (eventId) => {
    alert(`Reserve event ${eventId}`);
  };
  </script>
  