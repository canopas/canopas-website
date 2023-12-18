<template>
  <section class="mt-16 mb-16 lg:mt-60 lg:mb-0 3xl:outer-container">
    <div class="container">
      <div class="flex flex-col gap-4 lg:gap-6 text-center">
        <h2 class="header-2 text-black-core/[0.87]">Case studies</h2>
        <p class="sub-h1-regular px-4 text-black-core/[0.6]">
          Explore our case studies to witness how we've transformed
          <br class="hidden lg:flex" />our client's ideas into successful apps.
        </p>
      </div>
      <!-- ---------------------MOBILE UI START------------------ -->
      <div
        v-for="item in cases"
        :key="item"
        class="mt-6 relative lg:hidden"
        :class="item.className"
      >
        <a
          class="cursor-pointer"
          :href="item.url"
          @click.native="$mixpanel.track('tap_backend_app_portfolio')"
          :target="item.target"
        >
          <img
            :src="item.images[0]"
            :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
            alt="casestudy-image"
            class="w-full h-full object-cover"
            loading="lazy"
          />

          <div
            :class="item.textalign"
            class="absolute bottom-4 flex flex-col gap-2"
          >
            <h2 class="header-2 text-black-core/[0.87]">
              {{ item.title }}
            </h2>
            <p
              class="sub-h3-regular text-black-core/[0.6]"
              v-html="item.content"
            ></p>
          </div>
        </a>
      </div>
    </div>
    <!-- ------------------------ MOBILE UI END---------------------- -->

    <!-- ---------------------DESKTOP UI START------------------ -->
    <div class="hidden lg:block bg-[#F8F8F8] mt-[4.5rem]">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :loop="true"
        :speed="4000"
        :autoplay="{
          delay: 4000,
          disableOnInteraction: false,
        }"
        :spaceBetween="20"
        :navigation="true"
        :modules="modules"
        @autoplayTimeLeft="onAutoplayTimeLeft"
        class="swiper-container xll:container"
      >
        <swiper-slide v-for="item in cases" :key="item">
          <div
            class="flex flex-row py-[8.5rem] container px-10 xl:px-16 mx-auto"
          >
            <div class="flex flex-col gap-8 my-auto text-left flex-1">
              <div class="white-btn">
                <span class="header-3">
                  {{ item.title }}
                </span>
              </div>
              <div class="flex flex-col gap-4">
                <a
                  class="cursor-pointer"
                  :href="item.url"
                  @click.native="$mixpanel.track('tap_backend_app_portfolio')"
                  :target="item.target"
                >
                  <h3
                    class="header-3 text-black-core/[0.87]"
                    v-html="item.deskcontent"
                  ></h3
                ></a>
                <p
                  class="text-black-core/[0.6] sub-h3-medium"
                  v-html="item.description"
                ></p>
              </div>
            </div>
            <div class="flex-1 flex justify-end">
              <a
                class="cursor-pointer"
                :href="item.url"
                @click.native="$mixpanel.track('tap_backend_app_portfolio')"
                :target="item.target"
              >
                <img
                  :src="item.images[2]"
                  :srcset="`${item.images[2]} 400w,${item.images[3]} 800w`"
                  class="w-[29.79663rem] h-[28.95819rem] object-contain"
                  :alt="item.title"
                  loading="lazy"
              /></a>
            </div>
          </div>
        </swiper-slide>
        <div class="absolute top-0 left-0 w-full h-1 bg-black-core/[0.08]">
          <div
            class="progressbar h-1 bg-gradient-to-r from-[#FF835B] to-[#F2709C] rounded-xl"
            ref="progressBar"
          ></div>
        </div>
      </swiper>
    </div>
    <!-- ---------------------DESKTOP UI END------------------ -->
  </section>
</template>

<script setup type="module">
import justlymobile_400w from "@/assets/images/backend-development/casestudy/p1-400w.webp";
import justlymobile_800w from "@/assets/images/backend-development/casestudy/p1-800w.webp";
import luxeradiomobile_400w from "@/assets/images/backend-development/casestudy/p2-400w.webp";
import luxeradiomobile_800w from "@/assets/images/backend-development/casestudy/p2-800w.webp";
import tognessmobile_400w from "@/assets/images/backend-development/casestudy/p3-400w.webp";
import tognessmobile_800w from "@/assets/images/backend-development/casestudy/p3-800w.webp";
import justly_400w from "@/assets/images/backend-development/casestudy/desktop/p1-400w.webp";
import justly_800w from "@/assets/images/backend-development/casestudy/desktop/p1-800w.webp";
import luxeradio_400w from "@/assets/images/backend-development/casestudy/desktop/p2-400w.webp";
import luxeradio_800w from "@/assets/images/backend-development/casestudy/desktop/p2-800w.webp";
import togness_400w from "@/assets/images/backend-development/casestudy/desktop/p3-400w.webp";
import togness_800w from "@/assets/images/backend-development/casestudy/desktop/p3-800w.webp";
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";

const progressBar = ref(null);
const onAutoplayTimeLeft = (s, time, progress) => {
  progressBar.value.style.setProperty("--progress", `${progress * 100}%`);
};
const { $mixpanel } = useNuxtApp();
const modules = [Autoplay];
const cases = [
  {
    title: "Justly",
    url: "/portfolio/justly",
    target: "_self",
    content: "Create systems to build your dream life",
    deskcontent: `Overcome depression by <br> building habits`,
    description: `Justly aims to tackle loneliness, depression, <br>and mental health through innovative <br> solutions.`,
    images: [justlymobile_400w, justlymobile_800w, justly_400w, justly_800w],
    className: "text-left items-start",
    textalign: "w-[53%] ml-4",
  },
  {
    title: "Luxeradio",
    url: "/portfolio/luxeradio",
    target: "_self",
    content: "Radio, music and podcast",
    deskcontent: `A radio, music and podcast <br> platform`,
    description: `Luxe Radio displays the best of Moroccan and <br>international creation, emphasizing taste, <br>elegance, and refinement.`,
    images: [
      luxeradiomobile_400w,
      luxeradiomobile_800w,
      luxeradio_400w,
      luxeradio_800w,
    ],
    className: "text-left items-start",
    textalign: "w-[41%] right-4",
  },
  {
    title: "Togness",
    url: "/portfolio/togness",
    target: "_self",
    content: "Create systems to build your dream life",
    deskcontent: `Photo editor and video <br> maker app`,
    description: `Togness are create slideshows,and add <br> quotes and your own music for every special <br> occasion of your life!`,
    images: [
      tognessmobile_400w,
      tognessmobile_800w,
      togness_400w,
      togness_800w,
    ],
    className: "text-left items-start ",
    textalign: "w-[53%] ml-4 ",
  },
];
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";
.progressbar {
  width: calc(100% - var(--progress, 0%));
  transition: width 0.04s ease;
}
</style>
