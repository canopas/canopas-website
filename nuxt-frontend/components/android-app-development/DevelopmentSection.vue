<template>
  <div class="mt-16 lg:my-60">
    <div class="container flex flex-col text-center">
      <span
        class="mobile-header-2 lg:desk-header-2 text-black-87 lg:!mx-0 lg:!w-full xs:mx-auto xs:w-[71%]"
      >
        Android app development services</span
      >
      <span
        class="mt-4 md:mt-6 sub-h1-regular lg:mobile-header-2-regular text-black-60"
        >With over a decade of expertise in Android app development, our team
        can help you craft high- quality, user-friendly apps that drive business
        growth and customer engagement.</span
      >
    </div>
    <!-- ---------------------MOBILE UI START------------------ -->
    <div class="swiper-content !mt-6 md:!mt-12 lg:hidden">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="true"
        :spaceBetween="10"
        :modules="modules"
        class="swiper-container"
        @swiper="setSwiperRef"
        :breakpoints="{
          '768': {
            slidesPerView: 1.3,
          },
          '992': {
            slidesPerView: 1.5,
          },
          '1200': {
            slidesPerView: 1.6,
          },
          '1400': {
            slidesPerView: 1.9,
          },
          '1600': {
            slidesPerView: 2.1,
          },
          '1800': {
            slidesPerView: 2.4,
          },
          '2000': {
            slidesPerView: 2.6,
          },
          '2200': {
            slidesPerView: 2.8,
          },
          '2500': {
            slidesPerView: 3.2,
          },
          '2800': {
            slidesPerView: 3.8,
          },
          '3500': {
            slidesPerView: 4.5,
          },
        }"
      >
        <swiper-slide
          v-for="(item, index) in items"
          :key="index"
          class="cursor-pointer"
        >
          <div
            class="mr-4 rounded-[1.25rem] px-4 py-6 md:w-[34.7rem] lg:w-[40.75rem] xl:!w-[45.75rem]"
            :class="[item.bgColor]"
          >
            <div class="justify-left mb-5 flex items-center space-x-4">
              <img
                alt="development-icon"
                :src="item.icon"
                class="h-10 w-10 object-contain lg:hidden"
                loading="lazy"
              />
              <span class="mobile-header-3 text-black-87">{{
                item.title
              }}</span>
            </div>
            <div class="sub-h3-regular text-black-60 lg:p-4">
              <ul class="pl-8 pr-4 lg:hidden">
                <li
                  v-for="list in item.description"
                  :key="list"
                  class="mb-3.5 list-disc"
                >
                  {{ list }}
                </li>
              </ul>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- ---------------------MOBILE UI END------------------ -->

    <!-- ---------------------DESKTOP UI START------------------ -->

    <div class="hidden lg:flex flex-col container mt-16">
      <div class="flex gap-10 justify-center items-center -ml-4">
        <div
          v-for="(item, index) in items"
          :key="index"
          :class="[
            activeIndex == index
              ? 'border-primary-1 bg-white animate-zoomEffect'
              : '',
            index == 4 || index == 0 ? item.className : '',
          ]"
          class="flex justify-center items-center w-12 w-12 rounded-full border-2 border-black-8"
          @mouseover="(slideTo(index), (activeIndex = index))"
          @touchstart.passive="(slideTo(index), (activeIndex = index))"
          @touchmove.passive="(slideTo(index), (activeIndex = index))"
        >
          <img
            v-if="activeIndex != index"
            :src="item.image[0]"
            alt="development-service-image"
            class="p-2.5 h-[2.875rem] w-[2.875rem] object-cover"
          />
          <img
            v-else
            :src="item.image[1]"
            alt="development-service-image"
            class="p-2 object-cover"
          />
        </div>
      </div>
      <div class="swiper-content">
        <swiper
          :slidesPerView="1"
          :centeredSlides="true"
          :spaceBetween="20"
          :speed="1500"
          :modules="modules"
          class="swiper-container flex"
          @swiper="setSwiperRef"
          @slideChange="onSlideChange"
        >
          <swiper-slide
            v-for="(item, index) in items"
            :key="index"
            class="lg:px-8 2xl:px-11"
          >
            <div class="desk-header-3 text-primary-1 text-center mt-4">
              {{ item.title }}
            </div>
            <div
              class="p-10 xl:w-4/5 2xl:w-[70%] h-full mx-auto bg-white drop-shadow-[0px_0px_26px_rgba(0,0,0,0.08)] rounded-xl mt-6 mb-5"
            >
              <p
                v-for="list in item.description"
                :key="list"
                class="mb-6 mobile-header-3-medium text-black-60"
              >
                {{ list }}
              </p>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>

    <!-- ---------------------DESKTOP UI END------------------ -->
  </div>
</template>

<script type="module">
import { Zoom, Autoplay, Pagination } from "swiper/modules";
import mobile_consultation from "@/assets/images/andriod-app-development/development/1.webp";
import mobile_design from "@/assets/images/andriod-app-development/development/2.webp";
import mobile_custom from "@/assets/images/andriod-app-development/development/3.webp";
import mobile_app from "@/assets/images/andriod-app-development/development/4.webp";
import mobile_maintenance from "@/assets/images/andriod-app-development/development/5.webp";
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
import { Swiper, SwiperSlide } from "swiper/vue";
export default {
  data() {
    return {
      modules: [Zoom, Autoplay, Pagination],
      activeIndex: 0,
      swiper: null,
      items: [
        {
          title: "Android App Development Consultation",
          bgColor: "bg-pink-40",
          icon: mobile_consultation,
          image: [consultation, activeconsultation],
          description: [
            "At Canopas, we believe that a successful project begins with a thorough understanding of your vision, goals, and market dynamics.",

            "Our team of seasoned professionals collabortively works with you during the consultation phase,using their deep expertise to advise on optimal strategies, potential challenges, and innovative solutions.",

            "We meticulously assess your project requirements, suggest suitable tech stacks, provide a realistic timeline, and help you understand the project's potential.",
          ],
        },
        {
          title: "Android UI/UX Design",
          bgColor: "bg-[#F8F8F8]",
          icon: mobile_design,
          image: [design, activedesign],
          description: [
            "At Canopas, our design philosophy is centered on creating intuitive, user-friendly, and visually stunning interfaces that resonate with your target audience.",

            "We prioritize the user journey, ensuring seamless navigation, quick loading times, and attractive design elements that captivate and retain users. Our team stays updated with the latest design trends, including its varied device screen sizes and resolutions.",

            "From wireframing to prototyping to the final design, we work closely with you at each step,translating your vision into a compelling app design that looks great and delivers an unmatched user experience.",
          ],
        },
        {
          title: "Custom Android Application Development",
          bgColor: "bg-pink-40",
          icon: mobile_custom,
          image: [development, activedevelopment],
          description: [
            "Our approach to development is rooted in agility and focused collaboration.  We build in two-week sprints where our team devotes their complete energy and attention to sculpting the foundation of your app.",

            "With check-ins every two weeks, we ensure you're part of the development journey, giving you the opportunity to review our progress and provide valuable feedback. This leads to a development process that's streamlined and dynamic.",

            "Our aim is to create apps that leave a lasting impression, spur conversations, and consistently draw users back. With our team of professionals, we bring your ideas to life, crafting elegant, all-encompassing designs that transform your vision into a winning app in the competitive marketplace.",
          ],
        },
        {
          title: "Android App Deployment",
          bgColor: "bg-[#F8F8F8]",
          icon: mobile_app,
          image: [deployment, activedeployment],
          description: [
            "Launching your Android app into the marketplace involves more than just hitting 'publish'. At Canopas, we'll guide you through this crucial deployment phase with careful planning, meticulous testing, and comprehensive support.",

            "Our team ensures your app complies with all of Google Play Store's guidelines, optimizing everything from the app description and graphics to backend configurations for a smooth launch.",

            "We also devise and implement a robust pre and post-launch marketing strategy to build hype and drive initial downloads. Your success is our end goal, and we stand by you as your Android app takes its first steps into the digital world.",
          ],
        },
        {
          title: "Maintenance and Support",
          bgColor: "bg-pink-40",
          icon: mobile_maintenance,
          image: [maintenance, activemaintenance],
          description: [
            "Post-deployment, we monitor the app's performance and user feedback closely, ready to make quick updates or fixes if needed. Our maintenance and support services ensure your app stays relevant, competitive, and effective in the face of new devices, OS updates, and emerging user expectations.",

            "We swiftly handle any emerging issues, from minor bug fixes to major updates, all aimed at enhancing user experience and engagement.",

            "Moreover, we believe in proactive communication, providing you with regular reports and insights into your app's performance and user feedback, while being readily available to answer any queries or concerns. With our ongoing maintenance and support, you can focus on your business, knowing your app's functionality and relevance are in expert hands.",
          ],
        },
      ],
    };
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
      this.swiperRef.slideTo(index);
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
</style>
