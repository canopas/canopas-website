<template>
  <section class="xll:container flex flex-col lg:flex-row mt-16 lg:mt-60">
    <div
      class="flex flex-col lg:my-auto lg:w-1/2 2xl:w-[40%] lg:pl-[8%] 2xl:pl-[10%] xll:pl-0"
    >
      <p
        class="lg:w-[87%] xl:w-3/4 2xl:w-4/5 text-center lg:text-left mobile-header-2 lg:desk-header-2 text-black-87"
      >
        Our Success Stories
      </p>
      <p
        class="mt-4 text-center lg:text-left sub-h1-regular lg:mobile-header-2-regular text-black-60"
      >
        Check out the best from our vault of<br />
        50+ projects.
      </p>
      <nuxt-link
        class="hidden lg:flex mt-8 ml-0 gradient-btn primary-btn"
        to="/contact"
        @click.native="$mixpanel.track('tap_landing_cta')"
      >
        <span class="sub-h3-semibold lg:sub-h1-semibold">View Work </span>
      </nuxt-link>
    </div>
    <div class="swiper-content mt-6 lg:w-1/2 2xl:w-3/5 pl-4 md:pl-0">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="false"
        :autoplay="{
          delay: 2000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :spaceBetween="20"
        :pagination="{ type: 'progressbar' }"
        :navigation="true"
        :modules="modules"
        :breakpoints="{
          '576': {
            slidesPerView: 1.2,
          },
          '768': {
            centeredSlides: true,
            slidesPerView: 1.525,
          },
          '992': {
            centeredSlides: false,
            slidesPerView: 2.2,
          },
        }"
        @swiper="onSwiper"
        class="swiper-container !flex flex-col gap-2"
      >
        <swiper-slide
          v-for="(story, index) in stories"
          :key="index"
          class="cursor-pointer"
        >
          <nuxt-link
            :to="story.url"
            @click.native="$mixpanel.track(story.event)"
          >
            <img
              :src="story.image[0]"
              :srcset="`${story.image[0]} 400w, ${story.image[1]} 800w`"
              class="mb-2 w-full h-[600px] lg:h-full object-cover"
              alt="story-image"
              loading="lazy"
            />
          </nuxt-link>
        </swiper-slide>
        <div class="swiper-pagination"></div>
      </swiper>
    </div>
  </section>
</template>

<script>
import { Pagination, Navigation, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import justly400w from "@/assets/images/services/success-stories/justly-400w.webp";
import justly800w from "@/assets/images/services/success-stories/justly-800w.webp";
import luxeradio400w from "@/assets/images/services/success-stories/luxeradio-400w.webp";
import luxeradio800w from "@/assets/images/services/success-stories/luxeradio-800w.webp";
import togness400w from "@/assets/images/services/success-stories/togness-400w.webp";
import togness800w from "@/assets/images/services/success-stories/togness-800w.webp";
import smile400w from "@/assets/images/services/success-stories/smile+400w.webp";
import smile800w from "@/assets/images/services/success-stories/smile+800w.webp";
import config from "@/config.js";

export default {
  data() {
    return {
      modules: [Pagination, Navigation, Autoplay],
      swiper: null,
      reading: false,
      portfolioURL: "/portfolio",
      stories: [
        {
          image: [justly400w, justly800w],
          url: "/portfolio/justly",
          event: "tap_services_justly_portfolio",
        },
        {
          image: [luxeradio400w, luxeradio800w],
          url: "/portfolio/luxeradio",
          event: "tap_services_luxe_radio_portfolio",
        },
        {
          image: [togness400w, togness800w],
          url: "/portfolio/togness",
          event: "tap_services_togness_portfolio",
        },
        {
          image: [smile400w, smile800w],
          url: config.SMILEPLUS_URL,
          target: "_blank",
          event: "tap_services_smile_portfolio",
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
    };
  },
  mounted() {
    if (window.innerWidth > 991) {
      this.stories = Array(50).fill(this.stories).flat();
    }
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
  methods: {
    openPortfolio(story) {
      window.open(story.url, story.target ? story.target : "_self");
      this.$mixpanel.track(story.event);
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-wrapper {
  @apply !items-center;
}

.swiper-pagination {
  @apply !w-full !h-2 !rounded-full;
}

.swiper-horizontal > .swiper-pagination-progressbar,
.swiper-pagination-horizontal,
.swiper-pagination-progressbar-fill {
  @apply lg:hidden !bottom-0 !top-auto !left-0 w-full !bg-neutral-100;
}

.swiper-pagination-progressbar-fill {
  @apply !h-2 !bg-gradient-to-r !from-[#FF835B] !to-[#F2709C] !rounded-full;
}
</style>
