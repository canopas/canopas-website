<template>
  <section>
    <div class="div tw-relative tw-mt-10 tw-mb-20">
      <h2
        class="tw-pt-5 sm:tw-py-10 tw-text-center tw-text-[1.875rem] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem] md:tw-text-[2.813rem] tw-leading-[2.4375rem] md:tw-leadig-[3.875rem] tw-font-inter-bold tw-text-black-core/[0.87]"
      >
        Success Stories
      </h2>
      <!-- Mobile UI -->
      <div
        class="md:tw-hidden tw-container tw-my-8 tw-h-[700px] tw-relative tw-overflow-hidden"
      >
        <div
          class="tw-w-full tw--mt-4 tw-overflow-y-scroll hidden-scrollbar before:tw-z-[9] before:tw-blur-lg before:tw-absolute before:tw-left-0 before:tw-h-[8%] before:tw-w-full before:tw-bg-white-gradient"
        ></div>
        <div class="swiper-content">
          <swiper
            :slidesPerView="1.2"
            :spaceBetween="20"
            :direction="'vertical'"
            :modules="modules"
            :autoplay="{
              delay: 1000,
              disableOnInteraction: false,
            }"
            :loopedSlides="100"
            :speed="8000"
            :loop="true"
            class="swiper-container tw-h-[700px]"
          >
            <swiper-slide
              v-for="(client, index) in clients"
              :key="index"
              class="tw-pointer-events-none !tw-mb-0 !tw-h-[auto]"
            >
              <div
                v-for="(client, index) in clients"
                :key="index"
                class="tw-mb-8 tw-mx-auto tw-rounded-[5px] sm:tw-px-10 md:tw-px-6 lg:tw-px-4 xl:tw-px-24 tw-text-black-core/[0.87] tw-p-5 tw-bg-gradient-to-l tw-from-[#FF835B]/[0.10] tw-to-[#F2709C]/[0.10]"
              >
                <div
                  class="tw-font-inter-medium tw-text-[1.125rem] md:tw-text-[2.25rem] lg:tw-text-[3.125rem] tw-leading-[1.6875rem] md:tw-leading-[2.75rem] lg:tw-leading-[3.75rem] tw-transition-all tw-ease tw-duration-500"
                  v-html="client.review"
                ></div>
                <div class="tw-flex tw-flex-row tw-justify-between tw-mt-10">
                  <div class="tw-flex tw-flex-col">
                    <span class="tw-flex tw-flex-row">
                      <font-awesome-icon
                        v-for="i in 5"
                        :key="i"
                        class="fa-star tw-w-5 tw-h-5 tw-text-[#FF3D2E]"
                        icon="star"
                    /></span>
                    <span
                      class="tw-mt-3 tw-text-black-core/[0.87] tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.2rem] md:tw-text-[1.625rem]"
                      >{{ client.name }}</span
                    >
                    <span
                      class="tw-mt-2 tw-font-inter-regular tw-text-[0.75rem] md:tw-text-[1.25rem] tw-leading-[0.9rem] tw-text-black-core/[0.60]"
                      >{{ client.designation }}</span
                    >
                  </div>
                </div>
              </div>
            </swiper-slide>
          </swiper>
        </div>
        <div
          class="tw--mt-4 tw-w-full tw-h-12 after:tw-absolute after:tw-z-[9] after:tw-left-0 after:tw-h-[5%] after:tw-w-full after:tw-blur-lg after:tw-bg-white-gradient-bottom"
        ></div>
      </div>
      <!-- Mobile UI End-->
      <div class="tw-hidden md:tw-block">
        <h2
          class="testimonial tw-opacity-[5%] tw-top-1/2 tw-left-1/2 tw-transform tw--translate-x-1/2 tw--translate-y-1/2 tw-absolute tw-tracking-wide tw-font-black tw-font-inter-bold tw-text-[7.5rem] lg:tw-text-[9.5rem] xl:tw-text-[11.75rem] 2xl:tw-text-[13.75rem] xl:tw-leading-[20.625rem] tw-text-black-core/[0.4] tw-text-center"
        >
          Testimonials
        </h2>
        <div
          class="tw-mx-auto tw-hidden tw-container md:tw-flex tw-flex-row sm:tw-text-black-core/[0.87] tw-text-black-900 md:tw-px-6 lg:tw-px-4 xl:tw-px-24"
        >
          <div
            class="tw-basis-[190%] xl:tw-basis-[150%] tw-pointer-events-none lg:tw-mt-[7%]"
          >
            <swiper
              :slidesPerView="5"
              :centeredSlides="true"
              :spaceBetween="10"
              :allowTouchMove="false"
              :slideToClickedSlide="true"
              :direction="'vertical'"
              :speed="950"
              :autoplay="{
                delay: 4000,
                disableOnInteraction: false,
              }"
              :loop="true"
              :modules="modules"
              class="swiper-container tw-h-[600px] tw-px-2"
              @swiper="setSwiperRef"
              @slideChange="onSlideChange"
            >
              <swiper-slide v-for="(client, index) in client" :key="index">
                <div
                  class="tw-my-10 tw-flex tw-flex-col tw-items-center tw-justify-center"
                  :class="
                    activeIndex == index || activeindex == 0
                      ? 'tw-animate-zoomEffect tw-mt-8 '
                      : ''
                  "
                >
                  <img
                    @click="slideTo(index), (activeIndex = index)"
                    :class="
                      activeIndex == index || activeindex == 0
                        ? 'lg:tw-h-[50px] lg:tw-w-[50px]'
                        : 'tw-grayscale'
                    "
                    :src="getImageUrl(index)"
                    :alt="client.name"
                    class="tw-pointer-events-auto tw-cursor-pointer tw-h-10 tw-w-10 tw-mb-2"
                  />
                  <span
                    :class="[
                      client.className && activeIndex == index
                        ? 'tw--ml-6'
                        : '',
                      activeIndex == index
                        ? 'tw--ml-2 tw-text-black-core/[0.87] tw-text-[0.5rem] lg:tw-text-[1rem] tw-leading-[1.2rem] tw-font-inter-semibold tw-animate-zoomEffect'
                        : '',
                    ]"
                    class="tw-text-black-core/[0.60] tw-text-[0.875rem] tw-text-center tw-leading-[1.05rem] tw-font-inter-regular"
                  >
                    {{ getFirstWord(client.name) }}</span
                  >
                </div>
              </swiper-slide></swiper
            >
          </div>
          <div class="swiper-content">
            <swiper
              :slidesPerView="1"
              :centeredSlides="true"
              :spaceBetween="10"
              :speed="950"
              :direction="'vertical'"
              :allowTouchMove="false"
              :autoplay="{
                delay: 4000,
                disableOnInteraction: false,
              }"
              :loop="true"
              :modules="modules"
              class="swiper-container tw-h-[600px] lg:tw-h-[700px] tw-mt-[4%] tw-flex"
              @swiper="setSwiperRef"
              @slideChange="onSlideChange"
            >
              <swiper-slide
                v-for="(client, index) in client"
                :key="index"
                class="tw-pl-8 xl:tw-pl-0 !tw-flex tw-justify-center tw-flex-col"
              >
                <div class="-tw-ml-4 md:tw-pr-6 lg:tw-px-4 xl:tw-px-24">
                  <div class="tw-flex tw-gap-x-4 tw-my-5">
                    <span class="tw-flex tw-flex-row">
                      <font-awesome-icon
                        v-for="i in 5"
                        :key="i"
                        class="fa-star tw-w-6 tw-h-6 footer-icon"
                        icon="star"
                    /></span>
                    <span
                      class="tw-font-inter-regular tw-text-[1.125rem] tw-leading-[1.6875rem] tw-text-black-core/[0.87]"
                      >(5.0)</span
                    >
                  </div>
                  <div
                    class="tw-font-inter-semibold tw-text-[1.5rem] lg:tw-text-[2.125rem] tw-leading-[2.5rem] lg:tw-leading-[3.1875rem]"
                    v-html="client.review"
                  ></div>
                  <div class="tw-flex tw-flex-row tw-justify-between tw-my-5">
                    <div class="tw-flex tw-flex-col">
                      <span
                        class="tw-mt-3 v2-canopas-gradient-text tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.5rem]"
                        >{{ client.name }}</span
                      >
                      <span
                        class="tw-mt-4 tw-font-inter-regular tw-text-[1.125rem] tw-leading-[1.35rem] tw-text-black-core/[0.60]"
                        v-html="client.desktop_designation"
                      ></span>
                    </div>
                  </div>
                </div>
              </swiper-slide>
            </swiper>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import usericon from "@/assets/images/andriod-app-development/icon.webp";
import userinactiveicon from "@/assets/images/andriod-app-development/inactive-icon.webp";
export default {
  data() {
    return {
      modules: [Autoplay],
      usericon,
      userinactiveicon,
      swiper: null,
      activeIndex: 0,
      swiperRef: 0,
      imageUrl: "",

      clients: [
        {
          id: 1,
          name: "Rebecca Kimura",
          designation: "Founder at Togness, Australia",
          desktop_designation: `Founder at Togness,<span class="tw-font-inter-medium tw-text-black-core/[0.87]"> Australia</span>`,
          review:
            "There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they understood the first time.",
          className: "tw--ml-8",
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          desktop_designation: `Product Manager at Luxe, <span class="tw-font-inter-medium tw-text-black-core/[0.87]"> Morocco</span>`,
          review:
            "The Play Store is the hardest app store to get good reviews on, and we've just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas’ work.</span>",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          desktop_designation: `Founder at ActivScout,<span class="tw-font-inter-medium tw-text-black-core/[0.87]"> Canada</span>`,
          review:
            "I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a high performing iOS app.",
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          designation: "Founder at Brickandbatten, USA",
          desktop_designation: `Founder at Brickandbatten, <span class="tw-font-inter-medium tw-text-black-core/[0.87]"> USA</span>`,
          review:
            "There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the iOS and Android stores.",
        },
        {
          id: 5,
          name: "Cyril Trosset",
          designation: "CTO at Udini, France",
          desktop_designation: `CTO at Udini, <span class="tw-font-inter-medium tw-text-black-core/[0.87]"> France</span> `,
          review:
            "Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and high quality apps.",
        },
      ],
    };
  },
  computed: {
    client() {
      return Array(50).fill(this.clients).flat();
    },
  },
  mounted() {
    this.imageUrl = this.getImageUrl();
  },
  methods: {
    getFirstWord(str) {
      const words = str.trim().split(" ");
      return words[0];
    },
    onSlideChange() {
      this.activeIndex = this.swiperRef.realIndex;
    },
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
      this.activeIndex = this.swiperRef.realIndex;
    },
    slideTo(index) {
      this.swiperRef.slideToLoop(index);
    },
    getImageUrl(index) {
      if (this.activeIndex == index) {
        return usericon;
      } else {
        return userinactiveicon;
      }
    },
  },
  components: {
    Swiper,
    SwiperSlide,
    FontAwesomeIcon,
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
.swiper-wrapper {
  @apply tw-ease-linear;
}
.div .testimonial:not(.testimonial) {
  @apply tw-container;
}
</style>
