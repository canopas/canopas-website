<template>
  <section class="section mt-16 md:mt-40 xl:mt-60">
    <!-- ---------------------MOBILE UI START------------------ -->
    <div
      class="relative z-[1] h-[27rem] sm:h-[30rem] overflow-hidden lg:hidden"
    >
      <p class="text-center mobile-header-2 text-black-87">Our blogs</p>
      <div class="blog mt-6 block lg:hidden">
        <swiper
          :slidesPerView="1.425"
          :spaceBetween="0"
          :centeredSlides="true"
          class="p-4"
        >
          <swiper-slide
            v-for="(blog, index) in blogs"
            :key="index"
            class="relative flex flex-col"
          >
            <div class="overflow-hidden">
              <div class="scale-75 transition-all duration-200 ease-out">
                <img
                  @click="openBlog(blog.link, 'tap_android_app_blog_section')"
                  :src="blog.image[0]"
                  :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
                  :alt="blog.title"
                  class="h-full w-full rounded-lg object-cover"
                  loading="lazy"
                />
              </div>
              <div class="title hidden">
                <p class="mt-3 text-black-60">
                  {{ blog.publishDate }}
                </p>
                <a
                  @click.native="$mixpanel.track('tap_android_app_blog_title')"
                  :href="blog.link"
                  target="_blank"
                >
                  <p class="mt-2 cursor-pointer sub-h1-semibold text-black-87">
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
      class="container flex hidden flex-col lg:block lg:min-h-[580px] lg:py-0"
    >
      <p class="mb-8 text-center desk-header-2 text-black-87">Our blogs</p>
      <div class="flex flex-row space-x-5 space-y-1">
        <div
          @click="openBlog(blog.link, 'tap_android_app_blog_section')"
          v-for="blog in blogs.slice(0, 1)"
          :key="blog"
          class="content group relative mt-[1.3rem] flex basis-[54.3%] xl:basis-[53.6%] cursor-pointer flex-col rounded-[20px]"
        >
          <aspect-ratio height="96%">
            <img
              :src="blog.image[2]"
              :srcset="`${blog.image[2]} 800w`"
              class="object-fill h-full w-full cursor-pointer rounded-[20px]"
              loading="lazy"
              :alt="blog.title"
            />
          </aspect-ratio>

          <div
            class="ease absolute bottom-[0] left-[0] right-[0] rounded-[18px] transition-all duration-700 opacity-0 overflow-hidden w-full h-0 group-hover:opacity-100 group-hover:h-full"
            :class="blog.bgColor"
          >
            <div
              class="absolute top-1/2 left-1/2 text-center -translate-x-1/2 -translate-y-1/2"
            >
              <span
                v-html="blog.hovertitle"
                @click="openBlog(blog.link, 'tap_android_app_blog_section')"
              ></span>
            </div>
          </div>
        </div>
        <div class="mt-[0.3rem] flex basis-[45%] flex-col">
          <div
            class="cursor-pointer group content relative mt-4"
            v-for="blog in blogs.slice(1, 3)"
            :key="blog"
            @click="openBlog(blog.link, 'tap_android_app_blog_section')"
            :class="blog.className"
          >
            <img
              :src="blog.image[1]"
              :srcset="`${blog.image[1]} 800w`"
              class="object-fit h-full w-full cursor-pointer rounded-[20px]"
              loading="lazy"
              :alt="blog.title"
            />
            <div
              class="ease absolute bottom-[0] left-[0] right-[0] rounded-[18px] transition-all duration-700 opacity-0 overflow-hidden w-full h-0 group-hover:opacity-100 group-hover:h-full"
              :class="blog.bgColor"
            >
              <div
                class="absolute top-1/2 left-1/2 text-center -translate-x-1/2 -translate-y-1/2 w-[75%]"
              >
                <span
                  @click="openBlog(blog.link, 'tap_android_app_blog_section')"
                  class="text-black-87 sub-h3-regular lg:mobile-header-2-regular text-white"
                  v-html="blog.hovertitle"
                ></span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-[0.3rem] flex gap-x-[1.3rem]">
        <div
          class="cursor-pointer group content relative mt-4 flex-col"
          v-for="blog in blogs.slice(3, 5)"
          :key="blog"
          @click="openBlog(blog.link, 'tap_android_app_blog_section')"
          :class="blog.className"
        >
          <img
            :src="blog.image[1]"
            :srcset="`${blog.image[1]} 800w`"
            class="h-full w-full cursor-pointer rounded-[20px]"
            :class="blog.subclass"
            loading="lazy"
            :alt="blog.title"
          />
          <div
            class="ease absolute bottom-[0] left-[0] right-[0] rounded-[18px] transition-all duration-700 opacity-0 overflow-hidden w-full h-0 group-hover:opacity-100 group-hover:h-full"
            :class="blog.bgColor"
          >
            <div
              class="absolute top-1/2 left-1/2 text-center -translate-x-1/2 -translate-y-1/2"
            >
              <span
                @click="openBlog(blog.link, 'tap_android_app_blog_section')"
                class="sub-h3-regular lg:mobile-header-2-regular text-black-87"
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
          hovertitle: `<span class="desk-header-3 text-black-87">Android<span><br><span class="mobile-header-2-regular text-black-87">Code coverage using JaCoCo</span>`,
          link: "https://canopas.com/android-code-coverage-using-jacoco-6639a1fc4293",
          bgColor: "bg-gradient-to-b from-[#E7E7E7] to-[#DADADA]",
          subclass: "object-cover",
        },
        {
          id: 2,
          image: [mvvm_400w, mvvm_800w],
          publishDate: "Dec 3, 2021",
          title: "Jetpack Compose: MVVM State management in a simple way",
          hovertitle: `<span class="desk-header-3 text-black-87">Jetpack Compose</span><br><span class="text-black-87 mobile-header-2-regular">MVVM State management in a simple way</span>`,
          link: "https://canopas.com/jetpack-compose-mvvm-state-management-in-a-simple-way-4c632fa6f554",
          bgColor: "bg-gradient-to-b from-[#fffbfd] to-[#fffbfd]",
          subclass: "object-cover",
        },
        {
          id: 3,
          image: [keyboard_400w, keyboard_800w],
          publishDate: "Apr 26, 2022",
          title: "Keyboard Handling In Jetpack Compose — All You Need To Know",
          hovertitle: `<span class="text-black-87 mobile-header-2-regular">Keyboard Handling In<br> <span class="desk-header-3 text-black-87">Jetpack Compose</span><br><span class="text-black-87 mobile-header-2-regular">All You Need To Know</span></span>`,
          link: "https://canopas.com/keyboard-handling-in-jetpack-compose-all-you-need-to-know-3e6fddd30d9a",
          bgColor: "bg-gradient-to-b from-[#f9e1e9] to-[#fef2f6]",
          subclass: "object-cover",
        },
        {
          id: 4,
          image: [coroutine_400w, coroutine_800w],
          publishDate: "Feb 17, 2022",
          title:
            "Retrofit — Effective error handling with Kotlin Coroutine and Result API",
          hovertitle: `<span class="desk-header-3 text-white">Retrofit</span> <br><span class="text-white mobile-header-2-regular">Effective error handling with Kotlin Coroutine and Result API</span>`,
          link: "https://canopas.com/retrofit-effective-error-handling-with-kotlin-coroutine-and-result-api-405217e9a73d",
          className: "basis-[54%]",
          bgColor: "bg-gradient-to-b from-[#1f1f1f] to-[#1f1f1f]",
          subclass: "",
        },
        {
          id: 5,
          image: [websocket_400w, websocket_800w],
          publishDate: "Mar 10, 2022",
          title:
            "Android — Send live audio stream from client to server using WebSocket and OkHttp client",
          hovertitle: `<span class="desk-header-3 text-white">Android</span><br><span class=" text-white mobile-header-2-regular">Send live audio stream from client to server using WebSocket and OkHttp client</span>`,
          link: "https://canopas.com/android-send-live-audio-stream-from-client-to-server-using-websocket-and-okhttp-client-ecc9f28118d9",
          className: "basis-[60%] ",
          bgColor: "bg-gradient-to-b from-[#1f1f1f] to-[#1f1f1f]",
          subclass: "object-cover",
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
  @apply block animate-fadeIn delay-500;
}
.swiper-slide-active div {
  @apply scale-[1];
}
.section .blog:not(.blog) {
  @apply container;
}
</style>
