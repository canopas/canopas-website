<template>
  <section class="my-16 lg:mt-60">
    <!-- --------------------Mobile UI Start----------------------- -->
    <div class="text-center lg:hidden">
      <p class="mobile-header-2 text-black-87">Our blogs</p>
      <div class="mt-2">
        <swiper
          :slidesPerView="1.425"
          :spaceBetween="0"
          :centeredSlides="true"
          class="!p-4"
        >
          <swiper-slide v-for="(blog, index) in blogs" :key="index">
            <div
              class="scale-75 text-left transition-all duration-200 ease-out"
              @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
            >
              <img
                :src="blog.image[0]"
                :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
                :alt="blog.title"
                class="h-full w-full rounded-xl object-cover"
                loading="lazy"
              />
              <div class="p-4 flex flex-col gap-3">
                <p
                  class="sub-h1-semibold lg:mobile-header-2-semibold text-black-87"
                >
                  {{ blog.title }}
                </p>
                <p class="sub-h4-regular text-black-60">
                  {{ blog.readTime }}
                </p>
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
    <!-- --------------------Desktop UI Start----------------------- -->
    <div class="container hidden lg:block lg:min-h-[36.25rem]">
      <p class="mb-14 text-center lg:desk-header-2 text-black-87">Our blogs</p>
      <div class="flex space-x-5">
        <div
          @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
          v-for="blog in blogs.slice(0, 1)"
          :key="blog"
          class="content group relative flex basis-[65%]"
        >
          <img
            :src="blog.image[1]"
            :srcset="`${blog.image[1]} 800w`"
            :alt="blog.title"
            class="object-fill h-full w-full cursor-pointer rounded-xl"
            loading="lazy"
          />
          <div
            class="bg-deep-charcoal ease absolute top-1/2 left-1/2 right-1/2 rounded-xl transition-all duration-700 opacity-0 overflow-hidden w-0 h-0 hover:opacity-100 group-hover:h-full group-hover:w-full group-hover:top-0 group-hover:left-0 group-hover:right-0"
          >
            <div
              class="absolute w-full -top-72 left-1/2 text-center -translate-x-1/2 -translate-y-1/2 transition-all duration-700 transform group-hover:top-1/2"
            >
              <span
                v-html="blog.hoverTitle"
                @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
              ></span>
            </div>
          </div>
        </div>
        <div class="flex flex-col space-y-5 basis-[35%]">
          <div
            class="cursor-pointer group content relative"
            v-for="blog in blogs.slice(1, 3)"
            :key="blog"
            @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
            :class="blog.className"
          >
            <img
              :src="blog.image[1]"
              :srcset="`${blog.image[1]} 800w`"
              class="object-fill h-full w-full cursor-pointer rounded-xl"
              :alt="blog.title"
              loading="lazy"
            />
            <div
              class="bg-deep-charcoal ease absolute top-1/2 left-1/2 right-1/2 rounded-xl transition-all duration-700 opacity-0 overflow-hidden w-0 h-0 hover:opacity-100 group-hover:h-full group-hover:w-full group-hover:top-0 group-hover:left-0 group-hover:right-0"
            >
              <div
                class="absolute w-72 xl:w-80 -top-72 left-1/2 text-center -translate-x-1/2 -translate-y-1/2 transition-all duration-700 transform group-hover:top-1/2"
              >
                <span
                  v-html="blog.hoverTitle"
                  @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
                ></span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="flex space-x-5 mt-5">
        <div
          class="cursor-pointer group content relative"
          v-for="blog in blogs.slice(3, 5)"
          :key="blog"
          @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
        >
          <img
            :src="blog.image[1]"
            :srcset="`${blog.image[1]} 800w`"
            class="object-fit h-full w-full cursor-pointer rounded-xl"
            loading="lazy"
            :alt="blog.title"
          />
          <div
            class="bg-deep-charcoal ease absolute top-1/2 left-1/2 right-1/2 rounded-xl transition-all duration-700 opacity-0 overflow-hidden w-0 h-0 hover:opacity-100 group-hover:h-full group-hover:w-full group-hover:top-0 group-hover:left-0 group-hover:right-0"
          >
            <div
              class="absolute w-96 -top-72 left-1/2 right-0 text-center -translate-x-1/2 -translate-y-1/2 transition-all duration-700 transform group-hover:top-1/2"
            >
              <span
                v-html="blog.hoverTitle"
                @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
              ></span>
            </div>
          </div>
        </div>
      </div>
      <a
        class="primary-btn gradient-border-btn white-btn mt-[4.5rem]"
        :href="canopas"
        target="_blank"
        @click.native="$mixpanel.track('tap_flutter_blog_section')"
      >
        <span class="mobile-header-3-semibold">Read More Blogs </span>
        <Icon name="fa6-solid:arrow-right" class="fa ml-2" />
      </a>
    </div>
  </section>
</template>
<script setup>
import { Swiper, SwiperSlide } from "swiper/vue";
import AspectRatio from "@/components/utils/AspectRatio.vue";
import { openBlog } from "@/utils.js";
import blog1_400w from "@/assets/images/frontend-development/blog/1-400w.webp";
import blog1_800w from "@/assets/images/frontend-development/blog/1-800w.webp";
import blog2_400w from "@/assets/images/frontend-development/blog/2-400w.webp";
import blog2_800w from "@/assets/images/frontend-development/blog/2-800w.webp";
import blog3_400w from "@/assets/images/frontend-development/blog/3-400w.webp";
import blog3_800w from "@/assets/images/frontend-development/blog/3-800w.webp";
import blog4_400w from "@/assets/images/frontend-development/blog/4-400w.webp";
import blog4_800w from "@/assets/images/frontend-development/blog/4-800w.webp";
import blog5_400w from "@/assets/images/frontend-development/blog/5-400w.webp";
import blog5_800w from "@/assets/images/frontend-development/blog/5-800w.webp";

const { $mixpanel } = useNuxtApp();
const blogs = [
  {
    id: 1,
    image: [blog1_400w, blog1_800w],
    title:
      "How to Migrate Large Applications Efficiently from Vue.js To Nuxt.js",
    readTime: "7 min",
    hoverTitle: `<span class="mobile-header-2-semibold text-center text-white">How to Migrate Large Applications Efficiently from Vue.js To Nuxt.js</span>`,
    link: "https://blog.canopas.com/how-to-migrate-large-applications-efficiently-from-vue-js-to-nuxt-js-bdd9e41f82a1",
  },
  {
    id: 2,
    image: [blog2_400w, blog2_800w],
    title:
      "How to configure SSR in Nuxt3 and deploy it with Serverless Architecture",
    readTime: "10 min",
    hoverTitle: `<span class="mobile-header-2-semibold text-center text-white">How to configure SSR in Nuxt3 and deploy it with Serverless Architecture</span>`,
    link: "https://blog.canopas.com/how-to-configure-ssr-in-nuxt3-and-deploy-it-with-serverless-architecture-f86e63ae8b8d",
  },
  {
    id: 3,
    image: [blog3_400w, blog3_800w],
    title: "Guidelines for Designing a Responsive Website",
    readTime: "8 min",
    hoverTitle: `<span class="mobile-header-2-semibold text-center text-white">Guidelines for Designing a Responsive Website</span>`,
    link: "https://blog.canopas.com/guidelines-to-consider-while-designing-a-responsive-website-bbee716d0389",
  },
  {
    id: 4,
    image: [blog4_400w, blog4_800w],
    title: "How to improve the page speed index of your website? — Part 2",
    readTime: "6 min",
    hoverTitle: `<span class="mobile-header-2-semibold text-center text-white">How to improve the page speed index of your website? — Part 2</span>`,
    link: "https://blog.canopas.com/how-to-improve-the-page-speed-index-of-your-website-part-2-8e747f595f68",
  },
  {
    id: 5,
    image: [blog5_400w, blog5_800w],
    title: "Vue 3 component testing with Jest",
    readTime: "10 min",
    hoverTitle: `<span class="mobile-header-2-semibold text-center text-white">Vue 3 component testing with Jest</span>`,
    link: "https://blog.canopas.com/vue-3-component-testing-with-jest-8b80a8a8946b",
  },
];
const canopas = "https://blog.canopas.com";
</script>
<style lang="postcss" scoped>
.swiper-slide-active div {
  @apply scale-[1] rounded-xl shadow-lg;
}
.gradient-border-btn > .fa {
  @apply text-[#f2709c];
}
.gradient-border-btn:hover > .fa {
  @apply text-white;
}
</style>
