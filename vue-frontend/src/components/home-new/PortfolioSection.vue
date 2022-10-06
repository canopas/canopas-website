<template>
  <div
    class="tw-p-0 tw-mb-[30px] tw-mt-[50px] md:tw-mt-[90px] lg:tw-mt-[150px]"
  >
    <div class="tw-text-center v2-header-text tw-font-futura-bold">
      Portfolio
    </div>
    <div class="tw-mt-[50px] md:tw-mt-[80px] lg:tw-mt-[150px]">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        ref="portfolios"
        class="tw-flex tw-flex-col tw-items-center md:odd:tw-flex-row md:even:tw-flex-row-reverse tw-mt-6"
      >
        <div class="tw-w-full md:tw-w-6/12 lg:tw-w-[45%]">
          <img
            :src="portfolio.images[3]"
            :srcset="`${portfolio.images[0]} 400w, ${portfolio.images[1]} 800w, ${portfolio.images[2]} 1200w, ${portfolio.images[3]} 1600w`"
            sizes="(min-width: 992px) 45vw, 100vw"
            class="tw-w-full tw-h-full tw-object-cover"
            loading="lazy"
            :alt="portfolio.title + `-image`"
          />
        </div>

        <div
          class="tw-py-[5%] tw-px-[20%] md:tw-m-auto md:tw-p-0 md:tw-w-[35%]"
        >
          <div class="v2-title-text">{{ portfolio.title }}</div>
          <div class="v2-normal-text tw-mt-4">
            {{ portfolio.detail }}
          </div>
          <div class="tw-mt-10">
            <router-link
              v-if="!portfolio.target"
              :to="portfolio.link"
              class="v2-normal-2-text v2-button tw-flex tw-items-center tw-w-fit"
              @click.native="$gtag.event(portfolio.event)"
            >
              <span class="tw-mr-2.5">VIEW</span>
              <font-awesome-icon
                class="tw-text-black-900 fa"
                icon="arrow-right"
                id="leftArrow"
            /></router-link>

            <a
              v-else
              class="v2-normal-2-text v2-button tw-flex tw-items-center tw-w-fit"
              :href="portfolio.link"
              target="_blank"
              @click.native="$gtag.event(portfolio.event)"
            >
              <span class="tw-mr-2.5">VIEW</span>
              <font-awesome-icon
                class="tw-text-black-900 fa"
                icon="arrow-right"
                id="leftArrow"
              />
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import luxeradio400w from "@/assets/images/portfolio/luxeradio-400w.webp";
import luxeradio800w from "@/assets/images/portfolio/luxeradio-800w.webp";
import luxeradio1200w from "@/assets/images/portfolio/luxeradio-1200w.webp";
import luxeradio16000w from "@/assets/images/portfolio/luxeradio-1600w.webp";
import smilep400w from "@/assets/images/portfolio/smilep-400w.webp";
import smilep800w from "@/assets/images/portfolio/smilep-800w.webp";
import smilep1200w from "@/assets/images/portfolio/smilep-1200w.webp";
import smilep16000w from "@/assets/images/portfolio/smilep-1600w.webp";
import justly400w from "@/assets/images/portfolio/justly-400w.webp";
import justly800w from "@/assets/images/portfolio/justly-800w.webp";
import justly1200w from "@/assets/images/portfolio/justly-1200w.webp";
import justly16000w from "@/assets/images/portfolio/justly-1600w.webp";
import togness400w from "@/assets/images/portfolio/togness-400w.webp";
import togness800w from "@/assets/images/portfolio/togness-800w.webp";
import togness1200w from "@/assets/images/portfolio/togness-1200w.webp";
import togness16000w from "@/assets/images/portfolio/togness-1600w.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import config from "@/config.js";
import { analyticsEvent } from "@/utils.js";

export default {
  data() {
    return {
      portfolios: [
        {
          images: [
            luxeradio400w,
            luxeradio800w,
            luxeradio1200w,
            luxeradio16000w,
          ],
          title: "Luxeradio",
          detail:
            "Luxe Radio, the radio of taste, elegance, and refinement, intends to be the showcase of excellence and the best of Moroccan and international creation.",
          link: "/portfolio/luxeradio",
          target: false,
          event: "tap_home_luxe_radio_portfolio",
        },
        {
          images: [togness400w, togness800w, togness1200w, togness16000w],
          title: "Togness",
          aspectRatio: "55%",
          detail:
            "Togness is a photo editor and slideshow maker app for your life's most memorable events like weddings, pets, friends & family, and memorials, etc.",
          link: "/portfolio/togness",
          target: false,
          event: "tap_home_togness_portfolio",
        },
        {
          images: [justly400w, justly800w, justly1200w, justly16000w],
          title: "Justly",
          detail:
            "Justly is a start-up with a strong vision for overcoming loneliness, depression, and mental health-related issues for humanity.",
          link: "/portfolio/justly",
          target: false,
          event: "tap_home_justly_portfolio",
        },
        {
          images: [smilep400w, smilep800w, smilep1200w, smilep16000w],
          title: "Smile+",
          detail:
            "Smile+ app is designed for dentists to create a perfect smile for their patients. Using the Smile+ app, dentists can get their patients the best smile simulation in a minute, automated with AI.",
          link: config.SHOW_SMILEPLUS_PORTFOLIO
            ? "/portfolio/smileplus"
            : config.SMILEPLUS_URL,
          target: config.SHOW_SMILEPLUS_PORTFOLIO ? false : true,
          event: "tap_home_smile_portfolio",
        },
      ],
      event: "",
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    if (window.innerWidth < 768) {
      this.portfolios[0].image = mobileLuxeradio;
    }
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = analyticsEvent(this.$refs);
      if (event && this.event !== event) {
        this.event = event;
        this.$gtag.event(event);
      }
    },
  },
};
</script>
