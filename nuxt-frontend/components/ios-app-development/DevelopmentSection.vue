<template>
  <div class="my-16 md:my-24 flex flex-col xll:container">
    <div class="container mb-2.5 flex flex-col text-center">
      <span
        class="tracking-[-1.2px] mb-2.5 font-inter-bold text-3xl leading-[2.4375rem] text-black-core/[0.87] md:!mx-0 md:!w-full md:text-[2.65625rem] lg:text-[3.4375rem] md:leading-[3.796875rem] lg:leading-[5.15625rem] xs:mx-auto xs:w-[71%]"
      >
        iOS App Development Services</span
      >
      <span
        class="font-inter-regular text-base md:text-xl leading-6 md:leading-[1.875rem] text-black-core/[0.87] md:font-inter-medium lg:text-2xl lg:leading-9 w-11/12 2xl:w-9/12 mx-auto"
        >With a blend of robust technology, intuitive designs, and innovative
        strategies, our team can help you craft high-quality, standout
        user-friendly apps that drive business growth and customer
        engagement.</span
      >
    </div>
    <div class="swiper-content mt-4 md:hidden">
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
          class="cursor-pointer px-4 pb-4"
        >
          <div class="text-center">
            <span
              class="text-center font-inter-bold text-6xl leading-[5.625rem] border-text"
            >
              0{{ index + 1 }}</span
            >
          </div>
          <div
            class="text-center border rounded-tr-[50px] h-[700px] xs:h-[600px] p-4 rounded-bl-[50px] border-solid border-black-core/[0.87] bg-[#F8F8F8]"
          >
            <div
              class="justify-center px-5 h-16 mb-5 flex flex-row items-center"
            >
              <span
                class="tracking-[-0.96px] font-inter-semibold text-2xl leading-[1.95rem] v2-canopas-gradient-text"
                v-html="item.title"
              ></span>
            </div>
            <div
              class="font-inter-regular text-base leading-6 text-black-core/[0.87]"
            >
              <p v-for="list in item.description" :key="list" class="mb-4">
                {{ list }}
              </p>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>

    <div
      class="rounded-[20px] hidden md:flex flex-row container mt-16 bg-[#FFF] drop-shadow-[0px_20px_20px_rgba(0,0,0,0.10)]"
    >
      <div class="-ml-4">
        <div
          v-for="(item, index) in item"
          :key="index"
          :class="[
            activeIndex == index ? 'border-r-[#ff835b] bg-[#FFF]' : '',
            index == 4 || index == 0 ? item.className : '',
          ]"
          class="flex flex-col justify-center items-center h-[7.5rem] lg:h-[6.8rem] w-40 xl:w-60 xl:h-[6.2rem] bg-[#F5F5F5] border-b border-[#E3E3E3] border-r-[5px]"
          @mouseover="slideTo(index), (activeIndex = index)"
          @touchstart.passive="slideTo(index), (activeIndex = index)"
          @touchmove.passive="slideTo(index), (activeIndex = index)"
        >
          <img
            v-if="activeIndex != index"
            :src="item.image[0]"
            alt="development-service-image"
            class="h-12 lg:h-[45.6px] xl:h-12 w-12 lg:w-[45.6px] xl:w-12 object-cover"
          />
          <img
            v-else
            :src="item.image[1]"
            alt="development-service-image"
            class="h-12 lg:h-[45.6px] xl:h-12 w-12 lg:w-[45.6px] xl:w-12 object-cover"
          />
        </div>
      </div>
      <div class="swiper-content flex flex-col flex-1">
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
          class="swiper-container h-[600px] lg:h-[540px] xl:h-[495px] flex"
          @swiper="setSwiperRef"
          @slideChange="onSlideChange"
        >
          <swiper-slide
            v-for="(item, index) in item"
            :key="index"
            class="lg:px-8 lg:pt-2 2xl:px-11"
          >
            <div class="justify-start px-5 mt-3 lg:my-3 flex flex-row">
              <span
                class="font-inter-semibold text-[1.5625rem] leading-[4.03125rem] lg:text-[1.625rem] lg:leading-[2.4375rem] v2-canopas-gradient-text mr-1.5"
                >{{ index + 1 }}.
              </span>
              <span
                class="font-inter-semibold text-[1.5625rem] leading-[4.03125rem] lg:text-[1.625rem] lg:leading-[2.4375rem] v2-canopas-gradient-text"
                >{{ item.desktoptitle }}</span
              >
            </div>
            <div
              class="px-5 font-inter-regular text-lg lg:text-xl leading-[1.6875rem] lg:leading-[1.875rem] text-black-core/[0.87]"
            >
              <p v-for="list in item.description" :key="list" class="mb-8">
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
          className: "rounded-tl-[20px]",
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
          className: "rounded-bl-[20px]",
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
