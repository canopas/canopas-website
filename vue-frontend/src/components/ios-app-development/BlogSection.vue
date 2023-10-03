<template>
  <section class="tw-my-20 md:tw-mb-32 lg:tw-my-32 xll:tw-container">
    <div
      class="tw-pb-4 md:tw-py-12 tw-text-center tw-font-inter-bold tw-text-3xl tw-leading-[2.4375rem] md:tw-text-[2.65625rem] md:tw-leading-[3.796875rem] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem] tw-tracking-[-4%] xll:tw-py-24 tw-text-black-core/[0.87]"
    >
      <span>Our Blogs</span>
    </div>
    <!-- Mobile UI start -->
    <div class="tw-mt-8 tw-block md:tw-hidden">
      <swiper
        :grabCursor="true"
        :effect="'creative'"
        :centeredSlides="true"
        :loop="false"
        :creativeEffect="{
          limitProgress: 5,
          loop: false,
          next: {
            translate: [400, 0, 0],
          },
          prev: {
            translate: [0, 0, -10],
          },
        }"
        :slidesPerView="1"
        :spaceBetween="10"
        :modules="modules"
      >
        <swiper-slide
          v-for="(blog, index) in blogs"
          :key="index"
          class="tw-mt-8"
        >
          <div
            :class="blog.className"
            class="tw-w-80 tw-h-[350px] tw-drop-shadow-md tw-text-left tw-cursor-pointer tw-mx-auto"
          >
            <img
              @click="openBlog(blog.link, 'tap_ios_blog_section')"
              :src="blog.image"
              :srcset="`${blog.image} 400w`"
              :alt="blog.title"
              class="tw-h-44 tw-w-80 tw-object-cover"
              loading="lazy"
            />
            <div
              @click="openBlog(blog.link, 'tap_ios_blog_section')"
              class="=tw-w-80 tw-h-28 tw-bg-white tw-flex tw-flex-col tw-justify-center tw-px-8"
            >
              <h2
                class="tw-font-inter-medium tw-text-basetw-leading-6 tw-tracking-[-4%] tw-text-black-core/[0.87]"
              >
                {{ blog.title }}
              </h2>

              <span
                class="tw-mt-2 tw-font-inter-regular tw-text-sm tw-tracking-[-0.72px] lg:tw-leading-[1.3125rem] v2-canopas-gradient-text"
              >
                {{ blog.readTime }}
              </span>
            </div>
          </div>
        </swiper-slide></swiper
      >
    </div>

    <!-- Mobile UI End -->
    <!-- Desktop UI start -->
    <div class="tw-mt-8 md:tw-block tw-hidden">
      <swiper
        :slidesPerView="4"
        :centeredSlides="true"
        :loop="true"
        :autoplay="{
          delay: 2000,
          disableOnInteraction: false,
        }"
        :speed="2000"
        :spaceBetween="10"
        :modules="modules"
        :navigation="true"
        class="swiper-container"
        :breakpoints="{
          '768': {
            slidesPerView: 2,
          },
          '1200': {
            slidesPerView: 2.5,
          },
          '1400': {
            slidesPerView: 3.5,
          },

          '2400': {
            slidesPerView: 2.5,
          },
        }"
      >
        <swiper-slide
          v-for="(blog, index) in iosblog"
          :key="index"
          class="tw-cursor-pointer tw-relative tw-flex tw-flex-col tw-pl-4 sm:tw-pl-8"
        >
          <div class="tw-h-full tw-w-full">
            <img
              @click="openBlog(blog.link, 'tap_ios_blog_section')"
              :src="blog.deskImage[0]"
              :srcset="`${blog.deskImage[0]} 800w,${blog.deskImage[1]} 1200w`"
              :alt="blog.title"
              class="tw-h-full tw-w-full tw-object-cover tw-rounded-[12px]"
              loading="lazy"
            />
          </div>
          <div>
            <p
              @click="openBlog(blog.link, 'tap_ios_blog_section')"
              class="tw-mt-2 tw-font-inter-semibold tw-text-2xl lg:tw-text-[2.125rem] tw-w-full lg:tw-leading-[3.1875rem] !tw-tracking-[-4%] tw-text-black-core/[0.87]"
            >
              {{ blog.title }}
            </p>
            <p
              class="tw-mt-2 tw-font-inter-medium tw-text-base lg:tw-text-lg tw-tracking-[-0.72px] lg:tw-leading-[1.125rem] tw-text-black-core/[0.60]"
            >
              {{ blog.readTime }}
            </p>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>

<script>
import { EffectCreative, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";

import { openBlog } from "@/utils.js";

import mobile_blog1_400w from "@/assets/images/ios-app-development/blog/blog1-400w.webp";
import mobile_blog2_400w from "@/assets/images/ios-app-development/blog/blog2-400w.webp";
import mobile_blog3_400w from "@/assets/images/ios-app-development/blog/blog3-400w.webp";
import mobile_blog4_400w from "@/assets/images/ios-app-development/blog/blog4-400w.webp";
import mobile_blog5_400w from "@/assets/images/ios-app-development/blog/blog5-400w.webp";

import dektop_blog1_800w from "@/assets/images/ios-app-development/blog/desk-blog1-800w.webp";
import dektop_blog1_1200w from "@/assets/images/ios-app-development/blog/desk-blog1-1200w.webp";
import dektop_blog2_800w from "@/assets/images/ios-app-development/blog/desk-blog2-800w.webp";
import dektop_blog2_1200w from "@/assets/images/ios-app-development/blog/desk-blog2-1200w.webp";
import dektop_blog3_800w from "@/assets/images/ios-app-development/blog/desk-blog3-800w.webp";
import dektop_blog3_1200w from "@/assets/images/ios-app-development/blog/desk-blog3-1200w.webp";
import dektop_blog4_800w from "@/assets/images/ios-app-development/blog/desk-blog4-800w.webp";
import dektop_blog4_1200w from "@/assets/images/ios-app-development/blog/desk-blog4-1200w.webp";
import dektop_blog5_800w from "@/assets/images/ios-app-development/blog/desk-blog5-800w.webp";
import dektop_blog5_1200w from "@/assets/images/ios-app-development/blog/desk-blog5-1200w.webp";

export default {
  data() {
    return {
      openBlog,
      modules: [EffectCreative, Autoplay],

      blogs: [
        {
          title: "The Ultimate Guide to iOS unit testing with best practices",
          image: mobile_blog1_400w,
          readTime: "5 min read",
          deskImage: [dektop_blog1_800w, dektop_blog1_1200w],
          link: "https://blog.canopas.com/the-ultimate-guide-to-ios-unit-testing-with-best-practices-part-2-f2eacbf3d8d4",
        },
        {
          title: "SwiftUI Performance Tuning: Tips and Tricks",
          image: mobile_blog2_400w,
          readTime: "8 min read",
          deskImage: [dektop_blog2_800w, dektop_blog2_1200w],
          link: "https://blog.canopas.com/swiftui-performance-tuning-tips-and-tricks-a8f9eeb23ec4",
          className: "-tw-rotate-6",
        },
        {
          title: "Designing Data Layer For Offline-first Apps",
          image: mobile_blog3_400w,
          readTime: "7 min read",
          deskImage: [dektop_blog3_800w, dektop_blog3_1200w],
          link: "https://blog.canopas.com/ios-designing-data-layer-for-offline-first-apps-f929c83bb195",
          className: "tw-rotate-6",
        },
        {
          title: "SwiftUI — Complex navigation made easier with UIPilot",
          image: mobile_blog4_400w,
          readTime: "5 min read",
          deskImage: [dektop_blog4_800w, dektop_blog4_1200w],
          link: "https://blog.canopas.com/swiftui-complex-navigation-made-easier-with-uipilot-5b33279f3476",
          className: "-tw-rotate-6",
        },
        {
          title: "A Complete Guide to iOS App Auto Deployment with CI-CD",
          image: mobile_blog5_400w,
          readTime: "7 min read",
          deskImage: [dektop_blog5_800w, dektop_blog5_1200w],
          link: "https://blog.canopas.com/a-complete-guide-to-ios-app-auto-deployment-with-ci-cd-b5dc516ba41d",
          className: "tw-rotate-6",
        },
      ],
    };
  },

  computed: {
    iosblog() {
      return Array(50).fill(this.blogs).flat();
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
@import "swiper/css/effect-creative";
.swiper-slide {
  @apply !tw-overflow-visible;
}
.swiper-wrapper {
  @apply !tw-items-start;
}
</style>
