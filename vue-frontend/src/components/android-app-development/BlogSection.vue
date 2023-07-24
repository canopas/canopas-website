<template>
  <section class="section tw-my-20 md:tw-mt-16">
    <!-- Mobile UI start -->
    <div
      class="tw-relative tw-z-[1] tw-h-[27rem] tw-overflow-hidden tw-py-[3rem] sm:tw-py-[2rem] md:tw-hidden"
    >
      <p
        class="tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.4375rem] tw-text-black-core/[0.87]"
      >
        Our Blogs
      </p>
      <div class="blog tw-mt-8 tw-block md:tw-hidden">
        <swiper
          :slidesPerView="2"
          :spaceBetween="10"
          :centeredSlides="true"
          class="tw-p-[1rem]"
        >
          <swiper-slide
            v-for="(blog, index) in blogs"
            :key="index"
            class="tw-relative tw-flex tw-flex-col"
          >
            <div class="tw-overflow-hidden">
              <div class="tw-h-full tw-w-full tw-cursor-pointer">
                <img
                  @click="openBlog(blog.link)"
                  :src="blog.image[0]"
                  :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
                  :alt="blog.title"
                  class="tw-h-full tw-w-full tw-rounded-[10px] tw-object-cover"
                  loading="lazy"
                />
              </div>
              <div class="title tw-hidden">
                <p
                  class="tw-mt-2 tw-pl-2 tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.3125rem] tw-text-black-core/[0.6]"
                >
                  {{ blog.publishDate }}
                </p>
                <a
                  @click.native="mixpanel.track('tap_android__app_blog_title')"
                  :href="blog.link"
                  target="_blank"
                >
                  <p
                    class="tw-mt-2 tw-cursor-pointer tw-pl-2 tw-font-inter-semibold tw-text-[1.125rem] tw-leading-[1.6875rem] tw-text-black-core/[0.87]"
                  >
                    {{ blog.title }}
                  </p>
                </a>
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div
        class="tw-absolute tw-inset-0 tw-z-[0] tw-mx-auto tw-h-full tw-w-[60%] tw-rounded-[10px] tw-border tw-border-black-core/[0.60]"
      ></div>
    </div>

    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div
      class="tw-container tw-flex tw-hidden tw-flex-col md:tw-block md:tw-min-h-[580px] lg:tw-py-0"
    >
      <p
        class="tw-mb-8 tw-text-center tw-font-inter-bold tw-text-[3.4375rem] tw-leading-[5.15625rem] tw-text-black-core/[0.87]"
      >
        Our Blogs
      </p>
      <div class="tw-flex tw-flex-row tw-space-x-5 tw-space-y-1">
        <div
          @click="openBlog(blog.link)"
          v-for="blog in blogs.slice(0, 1)"
          :key="blog"
          class="content tw-group tw-relative tw-mt-[1.5rem] tw-flex tw-basis-[65%] tw-cursor-pointer tw-flex-col tw-rounded-[20px]"
        >
          <img
            :src="blog.image[1]"
            :srcset="`${blog.image[1]} 800w`"
            class="tw-object-fit tw-h-full tw-w-full tw-cursor-pointer tw-rounded-t-[20px]"
            loading="lazy"
            :alt="blog.title"
          />

          <div
            class="tw-rounded-b-[20px] tw-bg-gradient-to-b tw-from-[#DADADA]/[10%] tw-to-[#E7E7E7]/[100%] tw-px-8 tw-pb-[0.5rem]"
          >
            <div
              class="tw-my-[0.5rem] tw-flex tw-flex-row tw-justify-between tw-font-inter-regular tw-text-[1rem] tw-leading-[1.6875rem] tw-text-black-core/[0.6] lg:tw-text-[1.125rem] lg:tw-leading-[1.6875rem] 2xl:tw-my-5"
            >
              <span class="tw-w-auto">
                {{ blog.publishDate }}
              </span>
              <span class="tw-w-auto"> 3 min read</span>
            </div>
            <div class="tw-my-[10px]">
              <span
                @click="openBlog(blog)"
                class="tw-cursor-pointer tw-text-black-core[0.87] tw-cursor-pointer tw-font-inter-semibold tw-text-[1.125rem] tw-leading-[2.1875rem] lg:tw-text-[1.5rem] lg:tw-leading-[3.1875rem] 2xl:tw-text-[2.125rem]"
              >
                {{ blog.title }}</span
              >
            </div>
          </div>
          <div
            class="overlay tw-rounded-[20px] tw-duration-700 tw-ease-in tw-absolute tw-bottom-0 tw-left-0 tw-right-0 tw-overflow-hidden tw-w-full tw-h-0 tw-opacity-0 group-hover:tw-opacity-[0.9]"
            :class="blog.bgColor"
          >
            <div
              class="tw-absolute tw-top-1/2 tw-left-1/2 tw-text-center tw--translate-x-1/2 tw--translate-y-1/2"
            >
              <span
                v-html="blog.hovertitle"
                @click="openBlog(blog.link)"
              ></span>
            </div>
          </div>
        </div>
        <div class="tw-mt-[0.3rem] tw-flex tw-basis-[45%] tw-flex-col">
          <div
            class="tw-cursor-pointer tw-group content tw-relative tw-mt-[1rem]"
            v-for="blog in blogs.slice(1, 3)"
            :key="blog"
            @click="openBlog(blog.link)"
            :class="blog.className"
          >
            <img
              :src="blog.image[1]"
              :srcset="`${blog.image[1]} 800w`"
              class="tw-object-fit tw-h-full tw-w-full tw-cursor-pointer tw-rounded-[20px]"
              loading="lazy"
              :alt="blog.title"
            />
            <div
              class="overlay tw-rounded-[20px] tw-duration-700 tw-ease-in tw-absolute tw-bottom-0 tw-left-0 tw-right-0 tw-overflow-hidden tw-w-full tw-h-0 tw-opacity-0 group-hover:tw-opacity-[0.9]"
              :class="blog.bgColor"
            >
              <div
                class="tw-absolute tw-top-1/2 tw-left-1/2 tw-text-center tw--translate-x-1/2 tw--translate-y-1/2 tw-w-[75%]"
              >
                <span
                  @click="openBlog(blog.link)"
                  class="tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[1] tw-leading-[1.5rem] lg:tw-text-[1.5rem] lg:tw-leading-[2.25rem] tw-text-white"
                  v-html="blog.hovertitle"
                ></span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="tw-mt-[0.3rem] tw-flex tw-gap-x-[1.3rem]">
        <div
          class="tw-cursor-pointer tw-group content tw-relative tw-mt-[1rem] tw-flex-col"
          v-for="blog in blogs.slice(3, 5)"
          :key="blog"
          @click="openBlog(blog.link)"
          :class="blog.className"
        >
          <img
            :src="blog.image[1]"
            :srcset="`${blog.image[1]} 800w`"
            class="tw-object-cover tw-h-full tw-w-full tw-cursor-pointer tw-rounded-[20px]"
            loading="lazy"
            :alt="blog.title"
          />
          <div
            class="overlay tw-rounded-[18px] tw-duration-700 tw-ease-in tw-opacity-0 tw-absolute tw-bottom-0 tw-left-0 tw-right-0 tw-overflow-hidden tw-w-full tw-h-0 group-hover:tw-opacity-[0.9]"
            :class="blog.bgColor"
          >
            <div
              class="tw-absolute tw-top-1/2 tw-left-1/2 tw-text-center tw--translate-x-1/2 tw--translate-y-1/2 tw-opacity-1"
            >
              <span
                @click="openBlog(blog.link)"
                class="tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[1rem] lg:tw-text-[1.5rem] tw-leading-[1.5rem] lg:tw-leading-[2.25rem] tw-text-black-core/[0.87]"
                v-html="blog.hovertitle"
              ></span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>

<script>
import { Swiper, SwiperSlide } from "swiper/vue";

import jacoco_400w from "@/assets/images/andriod-app-development/blog/1-400w.webp";
import jacoco_800w from "@/assets/images/andriod-app-development/blog/1-800w.webp";
import mvvm_400w from "@/assets/images/andriod-app-development/blog/2-400w.webp";
import mvvm_800w from "@/assets/images/andriod-app-development/blog/2-800w.webp";
import keyboard_400w from "@/assets/images/andriod-app-development/blog/3-400w.webp";
import keyboard_800w from "@/assets/images/andriod-app-development/blog/3-800w.webp";
import coroutine_400w from "@/assets/images/andriod-app-development/blog/4-400w.webp";
import coroutine_800w from "@/assets/images/andriod-app-development/blog/4-800w.webp";
import websocket_400w from "@/assets/images/andriod-app-development/blog/5-400w.webp";
import websocket_800w from "@/assets/images/andriod-app-development/blog/5-800w.webp";

export default {
  data() {
    return {
      blogs: [
        {
          id: 1,
          image: [jacoco_400w, jacoco_800w],
          publishDate: "Sep 27, 2021",
          title: "Android code coverage using JaCoCo",
          hovertitle: `<span class="tw-font-inter-bold tw-text-[1.5rem] tw-leading-[2.1875rem] lg:tw-text-[2.125rem] lg:tw-leading-[3.1875rem] tw-text-black-core/[0.87]">Android code<span><br><span class="tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[1.5rem] tw-leading-[2.25rem] tw-text-black-core/[0.87]">coverage using JaCoCo</span>`,
          link: "https://blog.canopas.com/android-code-coverage-using-jacoco-6639a1fc4293",
          bgColor:
            "tw-bg-gradient-to-b tw-from-[#E7E7E7]/[0.9] tw-to-[#DADADA]/[0.9]",
        },
        {
          id: 2,
          image: [mvvm_400w, mvvm_800w],
          publishDate: "Dec 3, 2021",
          title: "Jetpack Compose: MVVM State management in a simple way",
          hovertitle: `<span class="tw-font-inter-bold tw-text-[1.5rem] tw-leading-[2.1875rem] xl:tw-text-[2.125rem] xl:tw-leading-[3.1875rem] tw-text-white">Jetpack Compose</span><br> MVVM State management in a simple way`,
          link: "https://blog.canopas.com/jetpack-compose-mvvm-state-management-in-a-simple-way-4c632fa6f554",
          bgColor:
            "tw-bg-gradient-to-b tw-from-[#282828]/[0.9] tw-to-[#282828]/[0.9]",
        },
        {
          id: 3,
          image: [keyboard_400w, keyboard_800w],
          publishDate: "Apr 26, 2022",
          title: "Keyboard Handling In Jetpack Compose — All You Need To Know",
          hovertitle: `<span >Keyboard Handling In<br> <span class="tw-font-inter-bold tw-text-[1.5rem] tw-leading-[2.1875rem] xl:tw-text-[2.125rem] xl:tw-leading-[3.1875rem] tw-text-white">Jetpack Compose</span><br>All You Need To Know</span>`,
          link: "https://blog.canopas.com/keyboard-handling-in-jetpack-compose-all-you-need-to-know-3e6fddd30d9a",
          bgColor: "tw-bg-gradient-to-b tw-from-[#070710] tw-to-[#222241]",
        },
        {
          id: 4,
          image: [coroutine_400w, coroutine_800w],
          publishDate: "Feb 17, 2022",
          title:
            "Retrofit — Effective error handling with Kotlin Coroutine and Result API",
          hovertitle: `<span class="tw-font-inter-bold tw-text-[1.5rem] tw-leading-[2.1875rem] lg:tw-text-[2.125rem] lg:tw-leading-[3.1875rem] tw-text-white">Retrofit</span> <br><span class="tw-text-white">Effective error handling with Kotlin Coroutine and Result API</span>`,
          link: "https://blog.canopas.com/retrofit-effective-error-handling-with-kotlin-coroutine-and-result-api-405217e9a73d",
          className: "tw-basis-[54%]",
          bgColor:
            "tw-bg-gradient-to-b tw-from-[#282828]/[0.9] tw-to-[#282828]/[0.9]",
        },
        {
          id: 5,
          image: [websocket_400w, websocket_800w],
          publishDate: "Mar 10, 2022",
          title:
            "Android — Send live audio stream from client to server using WebSocket and OkHttp client",
          hovertitle: `<span class="tw-font-inter-bold tw-text-[1.5rem] tw-leading-[2.1875rem] lg:tw-text-[2.125rem] lg:tw-leading-[3.1875rem] tw-text-black-core/[0.80]">Android</span><br>Send live audio stream from client to server using WebSocket and OkHttp client`,
          link: "https://blog.canopas.com/android-send-live-audio-stream-from-client-to-server-using-websocket-and-okhttp-client-ecc9f28118d9",
          className: "tw-basis-[60%] ",
          bgColor:
            "tw-bg-gradient-to-b tw-from-[#FEE4DF]/[0.9] tw-to-[#FDCCBC]/[0.9]",
        },
      ],
    };
  },
  methods: {
    openBlog(link) {
      window.open(link, "_blank");
      this.mixpanel.track("tap_android_app_blog_section");
    },
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
};
</script>
<style lang="postcss">
@import "swiper/css";
.swiper-slide-active .title {
  @apply tw-block tw-animate-fadeIn tw-delay-500;
}

.section .blog:not(.blog) {
  @apply tw-container;
}

.content:hover .overlay {
  @apply tw-h-full;
}
</style>
