<script>
import GetAllTags from "~/graphql/query/GetAllTags.gql";

import {
  Navigation,
  Pagination,
  Scrollbar,
  A11y,
  EffectCoverflow,
  Autoplay,
} from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";

import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "swiper/css/scrollbar";
import "swiper/css/effect-coverflow";
import "swiper/css/autoplay";

export default {
  components: {
    Swiper,
    SwiperSlide,
  },
  setup() {
    const { onResult: tagResult, refetch: refetchTags } = useQuery(GetAllTags, {
      limit: 100,
      offset: 0,
      order_by: [{ created_at: "desc" }],
    });
    const tag = ref("");
    const tags = ref([]);
    const onSwiper = (swiper) => {
      console.log(swiper);
    };

    const onSlideChange = () => {
      console.log("Slide changed");
    };
    const handleSlideClick = (eventTag) => {
      tag.value = eventTag.currentTarget.dataset.title;
      const where = {};
      where.tags = { _ilike: `%${tag.value}%` };
      const event = new CustomEvent("selected-tag", {
        detail: { where },
      });
      window.dispatchEvent(event);
      console.log("Slide clicked:", tag);
    };
    tagResult((result)=>{
      if(result.data){
        tags.value = result.data.tags;
        console.log('d;::::::::::::::',tags)
      }
    })

    return {
      onSwiper,
      onSlideChange,
      handleSlideClick,
      modules: [
        Navigation,
        Pagination,
        Scrollbar,
        A11y,
        EffectCoverflow,
        Autoplay,
      ],
      tags,
    };
  },
};
</script>
<template>
  <div v-if="tags"class="w-full event-carousel">
    <swiper
      :modules="modules"
      :slides-per-view="3"
      :space-between="50"
      loop
      centered-slides
      effect="coverflow"
      :coverflow-effect="{
        rotate: 50,
        stretch: 0,
        depth: 100,
        modifier: 1,
        slideShadows: true,
      }"
      :breakpoints="{
        640: { slidesPerView: 1 },
        768: { slidesPerView: 2 },
        1024: { slidesPerView: 3 },
      }"
      :pagination="{ clickable: true }"
      :navigation="true"
      :autoplay="{ delay: 2500, disableOnInteraction: false }"
      @swiper="onSwiper"
      @slideChange="onSlideChange"
    >
      <swiper-slide
        v-for="(tag, index) in tags"
        :key="index"
        @click="handleSlideClick"
        :data-title="tag.name"
      >
        <div class="event-slide">
          <img :src="tag.tag_image || 'https://via.placeholder.com/150'" :alt="tag.title" />
          <div class="event-info">
            <h3>{{ tag.name }}</h3>
            <p>{{ tag.description }}</p>
          </div>
        </div>
      </swiper-slide>
    </swiper>
  </div>
</template>

<style scoped>
.event-carousel {
  width: 100%;
  max-width: 1050px;
  margin: auto;
  padding: 20px;
  background-color: transparent;
  border-radius: 15px;
}

.event-slide {
  position: relative;
  display: flex;
  cursor: pointer;
  flex-direction: column;
  align-items: center;
  border-radius: 10px;
  overflow: hidden;
  transition: transform 0.3s ease-in-out;
  border: 1px solid #ddd;
}

.event-slide img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 10px;
}

.event-info {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 15px;
  background: rgba(0, 0, 0, 0.5);
  color: #fff;
  text-align: center;
}

.event-info h3 {
  margin: 0;
  font-size: 1.2rem;
  font-weight: bold;
}

.event-info p {
  font-size: 0.9rem;
  margin-top: 5px;
}

.swiper-slide-active {
  transform: scale(1.2);
}

.swiper-button-next,
.swiper-button-prev {
  color: white;
}

.swiper-pagination-bullet-active {
  background-color: white;
}
</style>
