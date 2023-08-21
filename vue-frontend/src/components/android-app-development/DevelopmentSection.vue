<template>
  <div class="tw-my-16 tw-flex tw-flex-col xll:tw-container xll:tw-mt-16">
    <div class="tw-container tw-mb-2.5 tw-flex tw-flex-col tw-text-center">
      <span
        class="tw-mb-2.5 tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.4375rem] tw-text-black-core/[0.87] md:!tw-mx-0 md:!tw-w-[100%] md:tw-text-[3.4375rem] md:tw-leading-[5.15625rem] xs:tw-mx-auto xs:tw-w-[71%]"
      >
        Android App Development Services</span
      >
      <span
        class="tw-font-inter-regular tw-text-base tw-text-black-core/[0.60] md:tw-font-inter-medium md:tw-text-[1.5rem] md:tw-leading-9 lg:tw-px-20 2xl:tw-px-56"
        >With over a decade of expertise in Android app development, our team
        can help you craft high- quality, user-friendly apps that drive business
        growth and customer engagement.</span
      >
    </div>
    <div class="swiper-content tw-mt-4">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="true"
        :spaceBetween="10"
        :loop="true"
        :modules="modules"
        class="swiper-container"
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
          '2200': {
            slidesPerView: 2.8,
          },
          '2450': {
            slidesPerView: 1.6,
          },
        }"
      >
        <swiper-slide
          v-for="(item, index) in items"
          :key="index"
          class="tw-cursor-pointer"
        >
          <div
            class="tw-relative !-tw-mr-2 tw-mb-8 tw-mt-12 tw-hidden md:tw-block"
          >
            <div
              class="tw-border-b-2 tw-border-dotted tw-border-[#000000]/[0.87] tw-text-center"
            ></div>
          </div>
          <div
            class="tw-mr-4 tw-rounded-[1.25rem] tw-px-2 tw-py-8 md:tw-w-[34.7rem] lg:tw-w-[40.75rem] xl:!tw-w-[45.75rem]"
            :class="[item.bgColor]"
          >
            <div
              class="tw-absolute tw-left-[38%] tw-top-0 tw-hidden tw-h-20 tw-bg-white tw-text-center md:tw-block lg:tw-left-[40%]"
            >
              <span
                class="tw-text-center tw-font-inter-bold tw-text-[3.75rem] tw-leading-[5.625rem]"
              >
                0{{ index + 1 }}</span
              >
            </div>

            <div
              class="tw-justify-left tw-mb-5 tw-flex tw-flex-row tw-items-center tw-space-x-4"
            >
              <img
                alt="development-icon"
                :src="item.icon"
                class="tw-h-10 tw-w-10 tw-object-contain md:tw-hidden"
                loading="lazy"
              />
              <span
                class="tw-font-inter-bold tw-text-[1.375rem] tw-leading-[1.7875rem] tw-text-black-core/[0.87] md:tw-text-[2rem] md:tw-leading-[3rem]"
                >{{ item.title }}</span
              >
            </div>
            <div
              class="tw-font-inter-regular tw-text-base tw-text-black-core/[0.87] md:tw-p-4 md:tw-text-[1.25rem] md:tw-leading-[1.875rem]"
            >
              <ul class="tw-pl-8 tw-pr-4 md:tw-hidden">
                <li
                  v-for="list in item.description"
                  :key="list"
                  class="tw-mb-3.5 tw-list-disc"
                >
                  {{ list }}
                </li>
              </ul>
              <p
                v-for="list in item.description"
                :key="list"
                class="tw-mb-4 tw-hidden md:tw-block"
              >
                {{ list }}
              </p>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
  </div>
</template>

<script type="module">
import { Autoplay } from "swiper/modules";
import consultation from "@/assets/images/andriod-app-development/development/1.webp";
import design from "@/assets/images/andriod-app-development/development/2.webp";
import custom from "@/assets/images/andriod-app-development/development/3.webp";
import app from "@/assets/images/andriod-app-development/development/4.webp";
import maintenance from "@/assets/images/andriod-app-development/development/5.webp";

import { Swiper, SwiperSlide } from "swiper/vue";
export default {
  data() {
    return {
      modules: [Autoplay],
      items: [
        {
          title: "Android App Development Consultation",
          bgColor: "tw-bg-[#FFEFEF]",
          icon: consultation,
          description: [
            "At Canopas, we believe that a successful project begins with a thorough understanding of your vision, goals, and market dynamics.",

            "Our team of seasoned professionals collabortively works with you during the consultation phase,using their deep expertise to advise on optimal strategies, potential challenges, and innovative solutions.",

            "We meticulously assess your project requirements, suggest suitable tech stacks, provide a realistic timeline, and help you understand the project's potential.",
          ],
        },
        {
          title: "Android UI/UX Design",
          bgColor: "tw-bg-[#EFF7FF]",
          icon: design,
          description: [
            "At Canopas, our design philosophy is centered on creating intuitive, user-friendly, and visually stunning interfaces that resonate with your target audience.",

            "We prioritize the user journey, ensuring seamless navigation, quick loading times, and attractive design elements that captivate and retain users. Our team stays updated with the latest design trends, including its varied device screen sizes and resolutions.",

            "From wireframing to prototyping to the final design, we work closely with you at each step,translating your vision into a compelling app design that looks great and delivers an unmatched user experience.",
          ],
        },
        {
          title: "Custom Android Application Development",
          bgColor: "tw-bg-[#EEFAF0]",
          icon: custom,
          description: [
            "Our approach to development is rooted in agility and focused collaboration.  We build in two-week sprints where our team devotes their complete energy and attention to sculpting the foundation of your app.",

            "With check-ins every two weeks, we ensure you're part of the development journey, giving you the opportunity to review our progress and provide valuable feedback. This leads to a development process that's streamlined and dynamic.",

            "Our aim is to create apps that leave a lasting impression, spur conversations, and consistently draw users back. With our team of professionals, we bring your ideas to life, crafting elegant, all-encompassing designs that transform your vision into a winning app in the competitive marketplace.",
          ],
        },
        {
          title: "Android App Deployment",
          bgColor: "tw-bg-[#F4F1FC]",
          icon: app,
          description: [
            "Launching your Android app into the marketplace involves more than just hitting 'publish'. At Canopas, we'll guide you through this crucial deployment phase with careful planning, meticulous testing, and comprehensive support.",

            "Our team ensures your app complies with all of Google Play Store's guidelines, optimizing everything from the app description and graphics to backend configurations for a smooth launch.",

            "We also devise and implement a robust pre and post-launch marketing strategy to build hype and drive initial downloads. Your success is our end goal, and we stand by you as your Android app takes its first steps into the digital world.",
          ],
        },
        {
          title: "Maintenance and Support",
          bgColor: "tw-bg-[#FDF9ED]",
          icon: maintenance,
          description: [
            "Post-deployment, we monitor the app's performance and user feedback closely, ready to make quick updates or fixes if needed. Our maintenance and support services ensure your app stays relevant, competitive, and effective in the face of new devices, OS updates, and emerging user expectations.",

            "We swiftly handle any emerging issues, from minor bug fixes to major updates, all aimed at enhancing user experience and engagement.",

            "Moreover, we believe in proactive communication, providing you with regular reports and insights into your app's performance and user feedback, while being readily available to answer any queries or concerns. With our ongoing maintenance and support, you can focus on your business, knowing your app's functionality and relevance are in expert hands.",
          ],
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

<style lang="postcss">
@import "swiper/css";
</style>
