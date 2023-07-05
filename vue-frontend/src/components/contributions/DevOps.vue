<template>
  <section
    class="tw-mb-[100px] tw-mt-[12%] tw-overflow-hidden sm:tw-mt-[7%] 2xll:tw-mt-[2%] 3xl:tw-mt-0"
  >
    <p
      class="tw-container tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.438rem] tw-text-black-core/[0.87] lg:tw-text-[3.438rem] lg:tw-leading-[5.156rem]"
    >
      DevOps: CI/CD and Auto Deployment
    </p>
    <!-- Mobile UI start -->
    <div class="tw-mt-[3rem] tw-block md:tw-hidden">
      <swiper :slidesPerView="1.325" :centeredSlides="true" :spaceBetween="0">
        <swiper-slide v-for="(blog, index) in devOps" :key="index">
          <div
            class="tw-scale-[0.75] tw-transition-all tw-duration-200 tw-ease-out"
          >
            <img
              @click="openBlog(blog.link)"
              :src="blog.image[0]"
              :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
              class="tw-relative tw-h-full tw-w-full tw-rounded-[10px] tw-object-cover"
              alt="devops"
            />
            <div
              class="tw-absolute tw-right-0 tw-top-0 tw-rounded-bl-[5px] tw-bg-white tw-px-3 tw-py-1 tw-font-inter-medium tw-text-[0.875rem] tw-leading-[1.313rem]"
            >
              {{ blog.readTime }}
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div
      class="tw-container tw-mt-16 tw-hidden tw-flex-col tw-gap-10 md:tw-flex"
    >
      <div class="tw-flex tw-flex-row tw-gap-10 tw-overflow-hidden">
        <div v-for="blog in devOps.slice(0, 2)" :key="blog.id">
          <div
            class="tw-group tw-relative tw-h-full tw-w-full tw-cursor-pointer"
          >
            <img
              :src="blog.image[0]"
              :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
              alt="devops"
              class="tw-h-full tw-w-full tw-rounded-[20px] tw-object-cover"
              loading="lazy"
            />
            <div @click="openBlog(blog.link)">
              <div
                class="gradient-class tw-opacity-1 tw-ease tw-absolute tw-bottom-[0] tw-left-[0] tw-h-[21%] tw-w-[100%] tw-rounded-b-[20px] tw-transition-all tw-duration-1000 group-hover:tw-h-full lg:group-hover:tw-h-[85%]"
              ></div>
              <div
                class="tw-opacity-1 tw-absolute tw-bottom-[-4rem] tw-translate-y-[50%] tw-p-[1rem] tw-font-inter-medium tw-text-white tw-transition-all tw-duration-1000 group-hover:tw-bottom-[-1rem] group-hover:tw-translate-y-[0%] group-hover:tw-transition-all group-hover:tw-duration-1000 xl:group-hover:tw-bottom-0 2xl:group-hover:tw-bottom-8"
              >
                <p
                  class="tw-text-[1.1rem] tw-leading-[1.1875rem] xl:tw-text-[1.125rem] xl:tw-leading-[1.688rem]"
                >
                  {{ blog.readTime }}
                </p>
                <p
                  class="tw-mt-4 tw-w-[100%] tw-text-[1.1rem] tw-leading-[1.875rem] xl:tw-text-[1.5rem] xl:tw-leading-[2.25rem] 2xl:tw-w-[75%]"
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
        class="tw-mx-auto tw-w-[50%] tw-overflow-hidden"
      >
        <div class="tw-group tw-relative tw-h-full tw-w-full tw-cursor-pointer">
          <img
            :src="blog.image[0]"
            :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
            alt="devops"
            class="tw-h-full tw-w-full tw-rounded-[20px] tw-object-cover"
            loading="lazy"
          />
          <div @click="openBlog(blog.link)">
            <div
              class="gradient-class tw-opacity-1 tw-ease tw-absolute tw-bottom-[0] tw-left-[0] tw-h-[21%] tw-w-[100%] tw-rounded-b-[20px] tw-transition-all tw-duration-1000 group-hover:tw-h-full lg:group-hover:tw-h-[85%]"
            ></div>
            <div
              class="tw-opacity-1 tw-absolute tw-bottom-[-4rem] tw-translate-y-[50%] tw-p-[1rem] tw-font-inter-medium tw-text-white tw-transition-all tw-duration-1000 group-hover:tw-bottom-[-1rem] group-hover:tw-translate-y-[0%] group-hover:tw-transition-all group-hover:tw-duration-1000 xl:group-hover:tw-bottom-0 2xl:group-hover:tw-bottom-8"
            >
              <p
                class="tw-text-[1.1rem] tw-leading-[1.1875rem] xl:tw-text-[1.125rem] xl:tw-leading-[1.688rem]"
              >
                {{ blog.readTime }}
              </p>
              <p
                class="tw-mt-4 tw-w-[100%] tw-text-[1.1rem] tw-leading-[1.875rem] xl:tw-text-[1.5rem] xl:tw-leading-[2.25rem] 2xl:tw-w-[75%]"
              >
                {{ blog.title }}
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="tw-flex tw-flex-row tw-gap-10 tw-overflow-hidden">
        <div v-for="blog in devOps.slice(3, 5)" :key="blog.id">
          <div
            class="tw-group tw-relative tw-h-full tw-w-full tw-cursor-pointer"
          >
            <img
              :src="blog.image[0]"
              :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
              alt="devops"
              class="tw-h-full tw-w-full tw-cursor-pointer tw-rounded-[20px] tw-object-cover"
              loading="lazy"
            />
            <div @click="openBlog(blog.link)">
              <div
                class="gradient-class tw-opacity-1 tw-ease tw-absolute tw-bottom-[0] tw-left-[0] tw-h-[21%] tw-w-[100%] tw-rounded-b-[20px] tw-transition-all tw-duration-1000 group-hover:tw-h-full lg:group-hover:tw-h-[85%]"
              ></div>
              <div
                class="tw-absolute tw-bottom-[-4rem] tw-translate-y-[50%] tw-p-[1rem] tw-font-inter-medium tw-text-white tw-transition-all tw-duration-1000 group-hover:tw-bottom-[-1rem] group-hover:tw-translate-y-[0%] group-hover:tw-transition-all group-hover:tw-duration-1000 xl:group-hover:tw-bottom-0 2xl:group-hover:tw-bottom-8"
              >
                <p
                  class="tw-text-[1.1rem] tw-leading-[1.1875rem] xl:tw-text-[1.125rem] xl:tw-leading-[1.688rem]"
                >
                  {{ blog.readTime }}
                </p>
                <p
                  class="tw-mt-4 tw-w-[100%] tw-text-[1.1rem] tw-leading-[1.875rem] xl:tw-text-[1.5rem] xl:tw-leading-[2.25rem] 2xl:tw-w-[75%]"
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
  methods: {
    openBlog(link) {
      window.open(link, "_blank");
      this.mixpanel.track("tap_contribution_Devops_section");
    },
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
  @apply tw-scale-[1];
}

.gradient-class {
  background-image: linear-gradient(
    7deg,
    #f77783 49%,
    rgba(255, 131, 91, 0) 50%
  );
}
</style>
