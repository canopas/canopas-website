<template>
  <section class="section tw-mt-16 md:tw-mt-40 xl:tw-mt-60">
    <!-- ---------------------MOBILE UI START------------------ -->
    <div
      class="tw-relative tw-z-[1] tw-h-[27rem] sm:tw-h-[30rem] tw-overflow-hidden lg:tw-hidden"
    >
      <p class="tw-text-center header-2 tw-text-black-core/[0.87]">Our blogs</p>
      <div class="blog tw-mt-6 tw-block lg:tw-hidden">
        <swiper
          :slidesPerView="1.425"
          :spaceBetween="0"
          :centeredSlides="true"
          class="tw-p-4"
        >
          <swiper-slide
            v-for="(blog, index) in blogs"
            :key="index"
            class="tw-relative tw-flex tw-flex-col"
          >
            <div class="tw-overflow-hidden">
              <div
                class="tw-scale-75 tw-transition-all tw-duration-200 tw-ease-out"
              >
                <img
                  @click="openBlog(blog.link, 'tap_android_app_blog_section')"
                  :src="blog.image[0]"
                  :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
                  :alt="blog.title"
                  class="tw-h-full tw-w-full tw-rounded-lg tw-object-cover"
                  loading="lazy"
                />
              </div>
              <div class="title tw-hidden">
                <p class="tw-mt-3 tw-text-black-core/[0.60]">
                  {{ blog.publishDate }}
                </p>
                <a
                  @click.native="mixpanel.track('tap_android__app_blog_title')"
                  :href="blog.link"
                  target="_blank"
                >
                  <p
                    class="tw-mt-2 tw-cursor-pointer sub-h1-semibold tw-text-black-core/[0.87]"
                  >
                    {{ blog.title }}
                  </p>
                </a>
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
    <!-- ---------------------MOBILE UI END------------------ -->

    <!-- ---------------------DSKTOP UI START------------------ -->
    <div
      class="tw-container tw-flex tw-hidden tw-flex-col lg:tw-block lg:tw-min-h-[580px] lg:tw-py-0"
    >
      <p class="tw-mb-8 tw-text-center header-2 tw-text-black-core/[0.87]">
        Our blogs
      </p>
      <div class="tw-flex tw-flex-row tw-space-x-5 tw-space-y-1">
        <div
          @click="openBlog(blog.link, 'tap_android_app_blog_section')"
          v-for="blog in blogs.slice(0, 1)"
          :key="blog"
          class="content tw-group tw-relative tw-mt-[1.3rem] tw-flex tw-basis-[54.3%] xl:tw-basis-[53.6%] tw-cursor-pointer tw-flex-col tw-rounded-[20px]"
        >
          <aspect-ratio height="96%">
            <img
              :src="blog.image[2]"
              :srcset="`${blog.image[2]} 800w`"
              class="tw-object-fill tw-h-full tw-w-full tw-cursor-pointer tw-rounded-[20px]"
              loading="lazy"
              :alt="blog.title"
            />
          </aspect-ratio>

          <div
            class="tw-ease tw-absolute tw-bottom-[0] tw-left-[0] tw-right-[0] tw-rounded-[18px] tw-transition-all tw-duration-700 tw-opacity-0 tw-overflow-hidden tw-w-full tw-h-0 group-hover:tw-opacity-100 group-hover:tw-h-full"
            :class="blog.bgColor"
          >
            <div
              class="tw-absolute tw-top-1/2 tw-left-1/2 tw-text-center tw--translate-x-1/2 tw--translate-y-1/2"
            >
              <span
                v-html="blog.hovertitle"
                @click="openBlog(blog.link, 'tap_android_app_blog_section')"
              ></span>
            </div>
          </div>
        </div>
        <div class="tw-mt-[0.3rem] tw-flex tw-basis-[45%] tw-flex-col">
          <div
            class="tw-cursor-pointer tw-group content tw-relative tw-mt-4"
            v-for="blog in blogs.slice(1, 3)"
            :key="blog"
            @click="openBlog(blog.link, 'tap_android_app_blog_section')"
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
              class="tw-ease tw-absolute tw-bottom-[0] tw-left-[0] tw-right-[0] tw-rounded-[18px] tw-transition-all tw-duration-700 tw-opacity-0 tw-overflow-hidden tw-w-full tw-h-0 group-hover:tw-opacity-100 group-hover:tw-h-full"
              :class="blog.bgColor"
            >
              <div
                class="tw-absolute tw-top-1/2 tw-left-1/2 tw-text-center tw--translate-x-1/2 tw--translate-y-1/2 tw-w-[75%]"
              >
                <span
                  @click="openBlog(blog.link, 'tap_android_app_blog_section')"
                  class="tw-font-inter-regular tw-text-black-core/[0.87] tw-text-base lg:tw-text-[1.5rem] lg:tw-leading-9 tw-text-white"
                  v-html="blog.hovertitle"
                ></span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="tw-mt-[0.3rem] tw-flex tw-gap-x-[1.3rem]">
        <div
          class="tw-cursor-pointer tw-group content tw-relative tw-mt-4 tw-flex-col"
          v-for="blog in blogs.slice(3, 5)"
          :key="blog"
          @click="openBlog(blog.link, 'tap_android_app_blog_section')"
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
            class="tw-ease tw-absolute tw-bottom-[0] tw-left-[0] tw-right-[0] tw-rounded-[18px] tw-transition-all tw-duration-700 tw-opacity-0 tw-overflow-hidden tw-w-full tw-h-0 group-hover:tw-opacity-100 group-hover:tw-h-full"
            :class="blog.bgColor"
          >
            <div
              class="tw-absolute tw-top-1/2 tw-left-1/2 tw-text-center tw--translate-x-1/2 tw--translate-y-1/2"
            >
              <span
                @click="openBlog(blog.link, 'tap_android_app_blog_section')"
                class="tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[1rem] lg:tw-text-[1.5rem] tw-leading-6 lg:tw-leading-9 tw-text-black-core/[0.87]"
                v-html="blog.hovertitle"
              ></span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- ---------------------DESKTOP UI END------------------ -->
  </section>
</template>

<script>
import { Swiper, SwiperSlide } from "swiper/vue";
import AspectRatio from "@/components/utils/AspectRatio.vue";
import { openBlog } from "@/utils.js";
import jacoco_desktop_400w from "@/assets/images/andriod-app-development/blog/desk-800w.webp";
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
      openBlog,
      blogs: [
        {
          id: 1,
          image: [jacoco_400w, jacoco_800w, jacoco_desktop_400w],
          publishDate: "Sep 27, 2021",
          title: "Android code coverage using JaCoCo",
          hovertitle: `<span class="header-3 tw-text-black-core/[0.87]">Android<span><br><span class="sub-h1-regular tw-text-black-core/[0.87]">Code coverage using JaCoCo</span>`,
          link: "https://blog.canopas.com/android-code-coverage-using-jacoco-6639a1fc4293",
          bgColor: "tw-bg-gradient-to-b tw-from-[#E7E7E7] tw-to-[#DADADA]",
        },
        {
          id: 2,
          image: [mvvm_400w, mvvm_800w],
          publishDate: "Dec 3, 2021",
          title: "Jetpack Compose: MVVM State management in a simple way",
          hovertitle: `<span class="header-3 tw-text-white">Jetpack Compose</span><br><span class="tw-text-white sub-h1-regular">MVVM State management in a simple way</span>`,
          link: "https://blog.canopas.com/jetpack-compose-mvvm-state-management-in-a-simple-way-4c632fa6f554",
          bgColor: "tw-bg-gradient-to-b tw-from-[#282828] tw-to-[#282828]",
        },
        {
          id: 3,
          image: [keyboard_400w, keyboard_800w],
          publishDate: "Apr 26, 2022",
          title: "Keyboard Handling In Jetpack Compose — All You Need To Know",
          hovertitle: `<span class="tw-text-white sub-h1-regular">Keyboard Handling In<br> <span class="header-3 tw-text-white">Jetpack Compose</span><br><span class="tw-text-white sub-h1-regular">All You Need To Know</span></span>`,
          link: "https://blog.canopas.com/keyboard-handling-in-jetpack-compose-all-you-need-to-know-3e6fddd30d9a",
          bgColor: "tw-bg-gradient-to-b tw-from-[#070710] tw-to-[#222241]",
        },
        {
          id: 4,
          image: [coroutine_400w, coroutine_800w],
          publishDate: "Feb 17, 2022",
          title:
            "Retrofit — Effective error handling with Kotlin Coroutine and Result API",
          hovertitle: `<span class="header-3 tw-text-white">Retrofit</span> <br><span class="tw-text-white sub-h1-regular">Effective error handling with Kotlin Coroutine and Result API</span>`,
          link: "https://blog.canopas.com/retrofit-effective-error-handling-with-kotlin-coroutine-and-result-api-405217e9a73d",
          className: "tw-basis-[54%]",
          bgColor: "tw-bg-gradient-to-b tw-from-[#282828] tw-to-[#282828]",
        },
        {
          id: 5,
          image: [websocket_400w, websocket_800w],
          publishDate: "Mar 10, 2022",
          title:
            "Android — Send live audio stream from client to server using WebSocket and OkHttp client",
          hovertitle: `<span class="header-3 tw-text-black-core/[0.87]">Android</span><br><span class=" tw-text-black-core/[0.87] sub-h1-regular">Send live audio stream from client to server using WebSocket and OkHttp client</span>`,
          link: "https://blog.canopas.com/android-send-live-audio-stream-from-client-to-server-using-websocket-and-okhttp-client-ecc9f28118d9",
          className: "tw-basis-[60%] ",
          bgColor: "tw-bg-gradient-to-b tw-from-[#FEE4DF] tw-to-[#FDCCBC]",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
    AspectRatio,
  },
  inject: ["mixpanel"],
};
</script>
<style lang="postcss">
@import "swiper/css";
.swiper-slide-active .title {
  @apply tw-block tw-animate-fadeIn tw-delay-500;
}
.swiper-slide-active div {
  @apply tw-scale-[1];
}
.section .blog:not(.blog) {
  @apply tw-container;
}
</style>
