<template>
  <div
    class="tw-my-16 md:tw-my-24 lg:tw-my-60 3xl:tw-my-[264px] tw-flex tw-flex-col xll:tw-container 3xl:tw-outer-container"
  >
    <div
      class="tw-container tw-mb-2.5 md:tw-mb-7 tw-flex tw-flex-col tw-text-center"
    >
      <span
        class="tw-mb-2.5 lg:tw-mb-6 tw-text-black-core/[0.87] md:!tw-mx-0 md:!tw-w-full xs:tw-mx-auto header-2"
      >
        iOS app <br class="lg:tw-hidden" />
        development services</span
      >
      <span
        class="sub-h1-regular tw-text-black-core/[0.60] tw-w-11/12 2xl:tw-w-9/12 tw-mx-auto"
        >With a blend of robust technology, intuitive designs, and innovative
        strategies, our team can help you craft high-quality, standout
        user-friendly apps that drive business growth and customer
        engagement.</span
      >
    </div>
    <div class="swiper-content tw-mt-4 lg:tw-hidden">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :spaceBetween="10"
        :modules="modules"
        :autoplay="{
          delay: 4000,
          disableOnInteraction: false,
        }"
        :breakpoints="{
          '540': {
            slidesPerView: 1.2,
            spaceBetween: 0,
          },
          '620': {
            slidesPerView: 1.4,
            spaceBetween: 0,
          },
          '700': {
            slidesPerView: 1.5,
          },
        }"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(item, index) in items"
          :key="index"
          class="tw-cursor-pointer tw-px-4 tw-pb-4"
        >
          <div
            class="tw-text-center tw-rounded-tr-[50px] tw-h-[700px] xs:tw-h-[600px] tw-p-4 tw-rounded-bl-[50px] tw-bg-[#F8F8F8]"
          >
            <div
              class="tw-justify-center tw-px-5 tw-h-16 tw-mb-5 tw-flex tw-flex-row tw-items-center"
            >
              <span
                class="header-3 tw-text-[#FF9472]"
                v-html="item.title"
              ></span>
            </div>
            <div class="sub-h3-regular tw-text-black-core/[0.87]">
              <p v-for="list in item.description" :key="list" class="tw-mb-4">
                {{ list }}
              </p>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>

    <div
      class="tw-rounded-[20px] tw-hidden lg:tw-flex tw-flex-row tw-container tw-mt-16 tw-bg-[#FFF] tw-drop-shadow-[0px_20px_20px_rgba(0,0,0,0.10)]"
    >
      <div class="tw--ml-4">
        <div
          v-for="(item, index) in item"
          :key="index"
          :class="[
            activeIndex == index ? 'tw-border-r-[#FF9472] tw-bg-[#FFF]' : '',
            index == 4 || index == 0 ? item.className : '',
          ]"
          class="tw-flex tw-flex-col tw-justify-center tw-items-center tw-h-[7.5rem] tw-h-[5.8rem] xl:tw-h-[6.8rem] tw-w-40 xl:tw-w-60 xl:tw-h-[6.2rem] tw-bg-[#F5F5F5] tw-border-b tw-border-[#E3E3E3] tw-border-r-[5px]"
          @mouseover="slideTo(index), (activeIndex = index)"
          @touchstart.passive="slideTo(index), (activeIndex = index)"
          @touchmove.passive="slideTo(index), (activeIndex = index)"
        >
          <img
            v-if="activeIndex != index"
            :src="item.image[0]"
            alt="development-service-image"
            class="tw-h-12 lg:tw-h-[45.6px] xl:tw-h-12 tw-w-12 lg:tw-w-[45.6px] xl:tw-w-12 tw-object-cover"
          />
          <img
            v-else
            :src="item.image[1]"
            alt="development-service-image"
            class="tw-h-12 lg:tw-h-[45.6px] xl:tw-h-12 tw-w-12 lg:tw-w-[45.6px] xl:tw-w-12 tw-object-cover"
          />
        </div>
      </div>
      <div class="swiper-content tw-flex tw-flex-col tw-flex-1">
        <swiper
          :slidesPerView="1"
          :spaceBetween="20"
          :direction="'vertical'"
          :speed="1500"
          :autoplay="{
            delay: 3000,
            disableOnInteraction: false,
          }"
          :loop="true"
          :modules="modules"
          class="swiper-container lg:tw-h-[463px] xl:tw-h-[495px] tw-flex"
          @swiper="setSwiperRef"
          @slideChange="onSlideChange"
        >
          <swiper-slide
            v-for="(item, index) in item"
            :key="index"
            class="lg:tw-px-8 lg:tw-pt-2 2xl:tw-px-11"
          >
            <div
              class="tw-justify-start tw-px-5 tw-mt-3 lg:tw-my-3 tw-flex tw-flex-row"
            >
              <span class="header-3 tw-text-[#FF9472] tw-mr-1.5"
                >{{ index + 1 }}.
              </span>
              <span class="header-3 tw-text-[#FF9472] tw-mr-1.5">{{
                item.desktoptitle
              }}</span>
            </div>
            <div class="tw-px-5 sub-h3-medium tw-text-black-core/[0.60]">
              <p v-for="list in item.description" :key="list" class="tw-mb-6">
                {{ list }}
              </p>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import consultation from "@/assets/images/ios-app-development/development/consultation.webp";
import design from "@/assets/images/ios-app-development/development/design.webp";
import development from "@/assets/images/ios-app-development/development/development.webp";
import deployment from "@/assets/images/ios-app-development/development/deployment.webp";
import maintenance from "@/assets/images/ios-app-development/development/maintenance.webp";
import activeconsultation from "@/assets/images/ios-app-development/development/activeconsultation.webp";
import activedesign from "@/assets/images/ios-app-development/development/activedesign.webp";
import activedevelopment from "@/assets/images/ios-app-development/development/activedevelopment.webp";
import activedeployment from "@/assets/images/ios-app-development/development/activedeployment.webp";
import activemaintenance from "@/assets/images/ios-app-development/development/activemaintenance.webp";
export default {
  data() {
    return {
      modules: [Autoplay],
      activeIndex: 0,
      swiperRef: 0,
      items: [
        {
          title: `iOS App Development <br/>Consultation`,
          desktoptitle: "iOS App Development Consultation",
          image: [consultation, activeconsultation],
          description: [
            "At Canopas, we understand that every great app starts with a compelling concept.",

            "Our team of experts is here to provide personalized guidance, helping you refine your idea, set clear objectives, and create a robust development strategy.",

            "We offer insights into market trends, technical feasibility, user behavior, and monetization strategies, aiming to ensure your app idea's transformation into a successful application.",
          ],
          className: "tw-rounded-tl-[20px]",
        },
        {
          title: "iOS UI/UX Design",
          desktoptitle: "iOS UI/UX Design",
          image: [design, activedesign],
          description: [
            "We believe that exceptional user experience is the cornerstone of every successful app.",

            "Our dedicated team of creative designers works meticulously to craft visually striking and user-friendly designs that resonate with your target audience.",

            "Following Apple's Human Interface Guidelines, we ensure seamless navigation, logical flow, and an aesthetically pleasing design that elevates user engagement and satisfaction.",
          ],
        },
        {
          title: `Custom iOS Application <br/> Development`,
          desktoptitle: "Custom iOS Application  Development",
          image: [development, activedevelopment],
          description: [
            "We understand that your business is unique, and so are its needs. Our approach to development is rooted in agility and focused collaboration.",

            "We build in two-week sprints where our team devotes their complete energy and attention to sculpting the foundation of your app.",

            "Our team of skilled developers leverages the power of Swift along with the robustness of XCode and other cutting-edge technologies to build tailor-made applications that reflect your business ethos.",

            "Whether you need an app for iPhone, iPad, Apple Watch, or Apple TV, we deliver high-performing, secure, and scalable applications that perfectly align with your business objectives and offer your users an unmatched experience.",
          ],
        },
        {
          title: "iOS App Deployment",
          desktoptitle: "iOS App Deployment",
          image: [deployment, activedeployment],
          description: [
            "At Canopas, we don't just create your iOS app, we also guide it successfully to the Apple App Store.",

            "Our team ensures that your application is optimized, bug-free, and adheres to Apple's stringent App Store Review Guidelines before submission.",

            "We handle every detail, from setting up the necessary configurations, managing app metadata, to executing the launch process. We also monitor the app’s performance and user feedback post- launch to make any necessary adjustments.",
          ],
        },
        {
          title: "Maintenance and Support",
          desktoptitle: "Maintenance and Support",
          image: [maintenance, activemaintenance],
          description: [
            "Post-deployment, we monitor the app's performance and user feedback closely,ready to make quick updates or fixes if needed. Our maintenance and support services ensure your app stays relevant, competitive, and effective in the face of new devices, OS updates, and emerging user expectations.",

            "We swiftly handle any emerging issues, from minor bug fixes to major updates, all aimed at enhancing user experience and engagement.",

            "Moreover, we believe in proactive communication, providing you with regular reports and insights into your app's performance and user feedback, while being readily available to answer any queries or concerns. With our ongoing maintenance and support, you can focus on your business, knowing your app's functionality and relevance are in expert hands.",
          ],
          className: "tw-rounded-bl-[20px]",
        },
      ],
    };
  },
  computed: {
    item() {
      return Array(1).fill(this.items).flat();
    },
  },
  methods: {
    onSlideChange() {
      this.activeIndex = this.swiperRef.realIndex;
    },
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
      this.activeIndex = this.swiperRef.realIndex;
    },
    slideTo(index) {
      this.activeIndex = index;
      this.swiperRef.slideToLoop(index);
    },
  },
  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
</style>
