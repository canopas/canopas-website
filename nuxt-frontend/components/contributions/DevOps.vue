<template>
  <section
    class="mb-[100px] mt-[12%] overflow-hidden sm:mt-[7%] 2xll:mt-[2%] 3xl:mt-0"
  >
    <p
      class="container text-center font-inter-bold text-[1.875rem] leading-[2.438rem] text-black-core/[0.87] lg:text-[3.438rem] lg:leading-[5.156rem]"
    >
      DevOps - CI/CD and Auto Deployment
    </p>
    <!-- Mobile UI start -->
    <div class="mt-12 block md:hidden">
      <swiper :slidesPerView="1.325" :centeredSlides="true" :spaceBetween="0">
        <swiper-slide v-for="(blog, index) in devOps" :key="index">
          <div class="scale-75 transition-all duration-200 ease-out">
            <img
              @click="openBlog(blog.link, 'tap_contribution_Devops_section')"
              :src="blog.image[0]"
              :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
              class="relative h-full w-full rounded-[10px] object-cover"
              alt="devops"
              loading="lazy"
            />
            <div
              class="absolute right-0 top-0 rounded-bl-[5px] bg-white px-3 py-1 font-inter-medium text-[0.875rem] leading-[1.313rem]"
            >
              {{ blog.readTime }}
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div class="container mt-16 hidden flex-col gap-10 md:flex">
      <div class="flex flex-row gap-10 overflow-hidden">
        <div v-for="blog in devOps.slice(0, 2)" :key="blog.id">
          <div class="group relative h-full w-full cursor-pointer">
            <img
              :src="blog.image[0]"
              :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
              alt="devops"
              class="h-full w-full rounded-[20px] object-cover"
              loading="lazy"
            />
            <div
              @click="openBlog(blog.link, 'tap_contribution_Devops_section')"
            >
              <div
                class="gradient-class opacity-1 ease absolute bottom-[0] left-[0] h-[21%] w-full rounded-b-[20px] transition-all duration-1000 group-hover:h-full lg:group-hover:h-[85%]"
              ></div>
              <div
                class="opacity-1 absolute -bottom-16 translate-y-[50%] p-4 font-inter-medium text-white transition-all duration-1000 group-hover:bottom-[-1rem] group-hover:translate-y-[0%] group-hover:transition-all group-hover:duration-1000 xl:group-hover:bottom-0 2xl:group-hover:bottom-8"
              >
                <p
                  class="text-[1.1rem] leading-[1.1875rem] xl:text-[1.125rem] xl:leading-[1.688rem]"
                >
                  {{ blog.readTime }}
                </p>
                <p
                  class="mt-4 w-full text-[1.1rem] leading-[1.875rem] xl:text-[1.5rem] xl:leading-9 2xl:w-[75%]"
                >
                  {{ blog.title }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div
        v-for="blog in devOps.slice(2, 3)"
        :key="blog.id"
        class="mx-auto w-[50%] overflow-hidden"
      >
        <div class="group relative h-full w-full cursor-pointer">
          <img
            :src="blog.image[0]"
            :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
            alt="devops"
            class="h-full w-full rounded-[20px] object-cover"
            loading="lazy"
          />
          <div @click="openBlog(blog.link, 'tap_contribution_Devops_section')">
            <div
              class="gradient-class opacity-1 ease absolute bottom-[0] left-[0] h-[21%] w-full rounded-b-[20px] transition-all duration-1000 group-hover:h-full lg:group-hover:h-[85%]"
            ></div>
            <div
              class="opacity-1 absolute -bottom-16 translate-y-[50%] p-4 font-inter-medium text-white transition-all duration-1000 group-hover:-bottom-4 group-hover:translate-y-[0%] group-hover:transition-all group-hover:duration-1000 xl:group-hover:bottom-0 2xl:group-hover:bottom-8"
            >
              <p
                class="text-[1.1rem] leading-[1.1875rem] xl:text-[1.125rem] xl:leading-[1.688rem]"
              >
                {{ blog.readTime }}
              </p>
              <p
                class="mt-4 w-full text-[1.1rem] leading-[1.875rem] xl:text-[1.5rem] xl:leading-9 2xl:w-[75%]"
              >
                {{ blog.title }}
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-row gap-10 overflow-hidden">
        <div v-for="blog in devOps.slice(3, 5)" :key="blog.id">
          <div class="group relative h-full w-full cursor-pointer">
            <img
              :src="blog.image[0]"
              :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
              alt="devops"
              class="h-full w-full cursor-pointer rounded-[20px] object-cover"
              loading="lazy"
            />
            <div
              @click="openBlog(blog.link, 'tap_contribution_Devops_section')"
            >
              <div
                class="gradient-class opacity-1 ease absolute bottom-[0] left-[0] h-[21%] w-full rounded-b-[20px] transition-all duration-1000 group-hover:h-full lg:group-hover:h-[85%]"
              ></div>
              <div
                class="absolute -bottom-16 translate-y-[50%] p-4 font-inter-medium text-white transition-all duration-1000 group-hover:-bottom-4 group-hover:translate-y-[0%] group-hover:transition-all group-hover:duration-1000 xl:group-hover:bottom-0 2xl:group-hover:bottom-8"
              >
                <p
                  class="text-[1.1rem] leading-[1.1875rem] xl:text-[1.125rem] xl:leading-[1.688rem]"
                >
                  {{ blog.readTime }}
                </p>
                <p
                  class="mt-4 w-full text-[1.1rem] leading-[1.875rem] xl:text-[1.5rem] xl:leading-9 2xl:w-[75%]"
                >
                  {{ blog.title }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>
<script>
import { openBlog } from "@/utils.js";

import { Swiper, SwiperSlide } from "swiper/vue";
import devops1_400w from "@/assets/images/contributions/devops/1-400w.webp";
import devops1_800w from "@/assets/images/contributions/devops/1-800w.webp";
import devops2_400w from "@/assets/images/contributions/devops/2-400w.webp";
import devops2_800w from "@/assets/images/contributions/devops/2-800w.webp";
import devops3_400w from "@/assets/images/contributions/devops/3-400w.webp";
import devops3_800w from "@/assets/images/contributions/devops/3-800w.webp";
import devops4_400w from "@/assets/images/contributions/devops/4-400w.webp";
import devops4_800w from "@/assets/images/contributions/devops/4-800w.webp";
import devops5_400w from "@/assets/images/contributions/devops/5-400w.webp";
import devops5_800w from "@/assets/images/contributions/devops/5-800w.webp";

export default {
  data() {
    return {
      openBlog,

      devOps: [
        {
          id: 1,
          image: [devops1_400w, devops1_800w],
          title:
            "Deploy a website with AWS S3 + CloudFront using CloudFormation",
          readTime: "8 min read",
          link: "https://blog.canopas.com/deploy-a-website-with-aws-s3-cloudfront-using-cloudformation-c2199dc6c435",
        },
        {
          id: 2,
          image: [devops2_400w, devops2_800w],
          title: "Android Continuous integration using Gitlab CI and Fastlane",
          readTime: "4 min read",
          link: "https://blog.canopas.com/android-continuous-integration-using-gitlab-ci-and-fastlane-1a5bbb8b5b38",
        },
        {
          id: 3,
          image: [devops3_400w, devops3_800w],
          title: "A Complete Guide to iOS App Auto Deployment with CI-CD",
          readTime: "7 min read",
          link: "https://blog.canopas.com/a-complete-guide-to-ios-app-auto-deployment-with-ci-cd-b5dc516ba41d",
        },
        {
          id: 4,
          image: [devops4_400w, devops4_800w],
          title: "Golang: Serverless deployment using AWS Cloudformation",
          readTime: "9 min read",
          link: "https://blog.canopas.com/golang-serverless-deployment-using-aws-cloudformation-4dee745cbf28",
        },
        {
          id: 4,
          image: [devops5_400w, devops5_800w],
          title:
            "Complete guide to deploying SSR Vite apps on AWS with automation",
          readTime: "5 min read",
          link: "https://blog.canopas.com/complete-guide-to-deploying-ssr-vite-apps-on-aws-with-automation-27676113d6ac",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>

<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-slide-active div {
  @apply scale-[1];
}

.gradient-class {
  background-image: linear-gradient(
    7deg,
    #f77783 49%,
    rgba(255, 131, 91, 0) 50%
  );
}
</style>
