<template>
  <section>
    <div class="div relative mt-16 md:mt-40 xl:mt-60">
      <h2 class="sm:pb-10 text-center header-2 text-black-core/[0.87]">
        Success stories
      </h2>
      <!-- ---------------------MOBILE UI START------------------ -->
      <div class="lg:hidden container my-8 h-[700px] relative overflow-hidden">
        <div
          class="w-full -mt-4 overflow-y-scroll hidden-scrollbar before:z-[9] before:blur-lg before:absolute before:left-0 before:h-[8%] before:w-full before:bg-white-gradient"
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
            class="swiper-container h-[700px]"
          >
            <swiper-slide
              v-for="(client, index) in clients"
              :key="index"
              class="pointer-events-none !mb-0 !h-[auto]"
            >
              <div
                v-for="(client, index) in clients"
                :key="index"
                class="mb-8 mx-auto rounded-[5px] sm:px-10 md:px-6 lg:px-4 xl:px-24 p-5 bg-[#F2709C]/[0.08]"
              >
                <div
                  class="text-black-core/[0.87] sub-h1-semibold transition-all ease duration-500"
                  v-html="client.review"
                ></div>
                <div class="flex flex-row justify-between mt-6">
                  <div class="flex flex-col">
                    <span class="flex flex-row">
                      <Icon
                        v-for="i in 5"
                        :key="i"
                        class="fa-star w-5 h-5 text-[#FFC700]"
                        name="fa6-solid:star"
                      />
                    </span>
                    <span class="mt-2 text-black-core/[0.87] sub-h3-regular">{{
                      client.name
                    }}</span>
                    <span class="sub-h4-regular text-black-core/[0.6]">{{
                      client.designation
                    }}</span>
                  </div>
                </div>
              </div>
            </swiper-slide>
          </swiper>
        </div>
        <div
          class="-mt-4 w-full h-12 after:absolute after:z-[9] after:left-0 after:h-[5%] after:w-full after:blur-lg after:bg-white-gradient-bottom"
        ></div>
      </div>
      <!-- ---------------------MOBILE UI END------------------ -->

      <!-- ---------------------DESKTOP UI START------------------ -->
      <div class="hidden lg:block">
        <h2
          class="testimonial opacity-[5%] top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 absolute tracking-wide font-black font-inter-bold text-[9.5rem] xl:text-[11.75rem] 2xl:text-[13.75rem] xl:leading-[20.625rem] text-black-core/[0.4] text-center"
        >
          Testimonials
        </h2>

        <div
          class="container w-[90%] 2xl:w-[50%] mx-auto flex items-center justify-center swiper-content"
        >
          <swiper
            :slidesPerView="1"
            :centeredSlides="true"
            :spaceBetween="0"
            :speed="900"
            :direction="'vertical'"
            :autoplay="{
              delay: 4000,
              disableOnInteraction: false,
            }"
            :loop="true"
            :modules="modules"
            class="swiper-container h-[700px] flex items-center justify-center"
          >
            <swiper-slide
              v-for="(client, index) in client"
              :key="index"
              class="!flex justify-center flex-col"
            >
              <div class="flex gap-x-4 my-5">
                <span class="flex flex-row">
                  <Icon
                    v-for="i in 5"
                    :key="i"
                    class="fa-star w-5 h-5 text-[#FFC700]"
                    name="fa6-solid:star"
                  />
                </span>
                <span
                  class="font-inter-regular text-[1.125rem] leading-[1.688rem] text-black-core/[0.87]"
                  >(5.0)</span
                >
              </div>
              <div
                class="font-opensans-semibold header-3 text-black-core/[0.87]"
                v-html="client.review"
              ></div>
              <div class="flex flex-row justify-between my-5">
                <div class="flex flex-col">
                  <span class="mt-3 text-black-core/[0.6] sub-h2-medium">{{
                    client.name
                  }}</span>
                  <span
                    class="mt-1 sub-h4-regular text-black-core/[0.40]"
                    v-html="client.desktop_designation"
                  ></span>
                </div>
              </div>
            </swiper-slide>
          </swiper>
        </div>
      </div>
    </div>
    <!-- ---------------------DESKTOP UI END------------------ -->
  </section>
</template>

<script>
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
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
          desktop_designation: `Founder at Togness, Australia`,
          review:
            "There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they understood the first time.",
          className: "-ml-8",
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          desktop_designation: `Product Manager at Luxe, Morocco`,
          review:
            "The Play Store is the hardest app store to get good reviews on, and we've just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas’ work.</span>",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          desktop_designation: `Founder at ActivScout, Canada`,
          review:
            "I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a high performing iOS app.",
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          designation: "Founder at Brickandbatten, USA",
          desktop_designation: `Founder at Brickandbatten, USA`,
          review:
            "There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the iOS and Android stores.",
        },
        {
          id: 5,
          name: "Cyril Trosset",
          designation: "CTO at Udini, France",
          desktop_designation: `CTO at Udini, France`,
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
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
.swiper-wrapper {
  @apply ease-linear;
}
.div .testimonial:not(.testimonial) {
  @apply container;
}
</style>
