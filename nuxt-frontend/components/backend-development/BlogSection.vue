<template>
  <section class="mb-16 lg:mb-60 3xl:outer-container">
    <h2
      class="text-center mobile-header-2 lg:desk-header-2 text-black-87 pb-4 lg:pb-[4.5rem]"
    >
      Our blogs
    </h2>
    <!-- Mobile UI start -->
    <div class="lg:hidden">
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
        <swiper-slide v-for="(blog, index) in blogs" :key="index" class="my-8">
          <div
            :class="blog.className"
            class="w-80 drop-shadow-md bg-white cursor-pointer mx-auto"
          >
            <img
              :src="blog.image"
              :srcset="`${blog.image} 400w`"
              :alt="blog.title"
              class="h-44 w-80 object-cover"
              loading="lazy"
              @click.native="openBlog(blog.link, 'tap_backend_blog_section')"
            />
            <div
              class="w-80 p-4 flex flex-col gap-3"
              @click.native="openBlog(blog.link, 'tap_backend_blog_section')"
            >
              <span class="text-black-87 sub-h3-medium">{{ blog.title }}</span>
              <span class="text-black-60 sub-h4-medium">{{
                blog.readTime
              }}</span>
            </div>
          </div></swiper-slide
        >
      </swiper>
    </div>
    <!-- Mobile UI ends -->
    <!-- Desktop UI start -->
    <div class="hidden lg:block xll:container">
      <swiper
        :slidesPerView="3"
        :centeredSlides="true"
        :loop="true"
        :autoplay="{
          delay: 2000,
          disableOnInteraction: false,
        }"
        :speed="2000"
        :spaceBetween="20"
        :modules="modules"
        :navigation="true"
        class="swiper-container"
        :breakpoints="{
          '992': {
            slidesPerView: 2,
          },
          '1200': {
            slidesPerView: 2.5,
          },
          '1400': {
            slidesPerView: 3.5,
          },
          '2440': {
            slidesPerView: 2.5,
          },
        }"
      >
        <swiper-slide
          v-for="(blog, index) in backendblog"
          :key="index"
          class="cursor-pointer"
        >
          <div class="h-full w-full">
            <img
              @click="openBlog(blog.link, 'tap_backend_blog_section')"
              :src="blog.deskImage[0]"
              :srcset="`${blog.deskImage[0]} 800w,${blog.deskImage[1]} 1200w`"
              :alt="blog.title"
              class="h-full w-full object-cover rounded-xl"
              loading="lazy"
            />
          </div>
          <div class="mt-6 flex flex-col gap-6">
            <span
              @click="openBlog(blog.link, 'tap_backend_blog_section')"
              class="desk-header-3 text-black-87"
            >
              {{ blog.title }}
            </span>
            <span class="sub-h1-regular text-black-60">
              {{ blog.readTime }}
            </span>
          </div></swiper-slide
        >
      </swiper>
    </div>
    <!-- Desktop UI ends -->
  </section>
</template>
<script setup>
import { EffectCreative, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { openBlog } from "@/utils.js";

import mobile_blog1_400w from "@/assets/images/backend-development/blogs/p1-400w.webp";
import mobile_blog2_400w from "@/assets/images/backend-development/blogs/p2-400w.webp";
import mobile_blog3_400w from "@/assets/images/backend-development/blogs/p3-400w.webp";
import mobile_blog4_400w from "@/assets/images/backend-development/blogs/p4-400w.webp";
import mobile_blog5_400w from "@/assets/images/backend-development/blogs/p5-400w.webp";

import dektop_blog1_800w from "@/assets/images/backend-development/blogs/desktop/p1-800w.webp";
import dektop_blog1_1200w from "@/assets/images/backend-development/blogs/desktop/p1-800w.webp";
import dektop_blog2_800w from "@/assets/images/backend-development/blogs/desktop/p2-800w.webp";
import dektop_blog2_1200w from "@/assets/images/backend-development/blogs/desktop/p2-800w.webp";
import dektop_blog3_800w from "@/assets/images/backend-development/blogs/desktop/p3-800w.webp";
import dektop_blog3_1200w from "@/assets/images/backend-development/blogs/desktop/p3-1200w.webp";
import dektop_blog4_800w from "@/assets/images/backend-development/blogs/desktop/p4-800w.webp";
import dektop_blog4_1200w from "@/assets/images/backend-development/blogs/desktop/p4-1200w.webp";
import dektop_blog5_800w from "@/assets/images/backend-development/blogs/desktop/p5-800w.webp";
import dektop_blog5_1200w from "@/assets/images/backend-development/blogs/desktop/p5-1200w.webp";

const modules = [Autoplay, EffectCreative];
const blogs = [
  {
    title:
      "Scaling with Confidence — The Backend Strategies Every Startup Needs",
    image: mobile_blog1_400w,
    readTime: "13 min read",
    deskImage: [dektop_blog1_800w, dektop_blog1_1200w],
    link: "https://canopas.com/scaling-with-confidence-the-backend-strategies-every-startup-needs-8449426557",
  },
  {
    title: "What is Backend Development & Its Role in Mobile App Development?",
    image: mobile_blog2_400w,
    readTime: "7 min read",
    deskImage: [dektop_blog2_800w, dektop_blog2_1200w],
    link: "https://canopas.com/what-is-backend-development-its-role-in-mobile-app-development-6114455770",
    className: "-rotate-6",
  },
  {
    title: "Serverless Microservices using AWS lambda and API Gateway",
    image: mobile_blog3_400w,
    readTime: "4 min read",
    deskImage: [dektop_blog3_800w, dektop_blog3_1200w],
    link: "https://canopas.com/serverless-microservices-aws-lambda-and-api-gateway-d23f9346f0b6",
    className: "rotate-6",
  },
  {
    title:
      "Simplified Your Deployment: GitHub Workflow with EC2, ECR, and Docker",
    image: mobile_blog4_400w,
    readTime: "8 min read",
    deskImage: [dektop_blog4_800w, dektop_blog4_1200w],
    link: "https://canopas.com/simplified-your-deployment-github-workflow-with-ec2-ecr-and-docker-4c9c392d238d",
    className: "-rotate-6",
  },
  {
    title: "Golang: Serverless deployment using AWS Cloudformation",
    image: mobile_blog5_400w,
    readTime: "9 min read",
    deskImage: [dektop_blog5_800w, dektop_blog5_1200w],
    link: "https://canopas.com/golang-serverless-deployment-using-aws-cloudformation-4dee745cbf28",
    className: "rotate-6",
  },
];
const backendblog = computed(() => Array(5).fill(blogs).flat());
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/effect-creative";
.swiper-slide {
  @apply !overflow-visible;
}
.swiper-wrapper {
  @apply !items-start;
}
</style>
