<template>
  <section class="tw-mt-[100px] tw-font-inter-regular">
    <h1
      class="tw-text-black-core/[.03] tw-text-[2.5rem] md:tw-text-[3.75rem] xl:tw-text-[5rem] tw-leading-[1.875rem] md:tw-leading-[4.0625rem] xl:tw-leading-[5rem] tw-text-center"
    >
      Our Portfolio
    </h1>
    <h1
      class="tw-container tw-mt-[-20px] md:tw-mt-[-50px] tw-pb-9 md:tw-pb-12 tw-text-[1.875rem] tw-leading-[2.438rem] md:tw-text-[2.65rem] md:tw-leading-[3.218rem] lg:tw-text-[3.438rem] lg:tw-leading-[4rem] tw-text-center tw-text-black-core/[0.87] tw-font-roboto-bold"
    >
      Case Studies
    </h1>
    <div
      class="tw-container tw-flex tw-justify-around tw-m-auto tw-w-[100%] sm:tw-w-[95%] 2xl:tw-w-[45%] tw-rounded-[30px] tw-bg-white tw-py-[5px] tw-drop-shadow-[0_2px_10px_rgba(61,61,61,0.1)]"
    >
      <a
        v-for="(navbar, index) in navbars"
        :key="index"
        class="tw-flex-[16.66%] tw-m-0 tw-w-full tw-rounded-[25px] tw-px-[5px] tw-py-2 md:tw-py-[13px] tw-text-center tw-tracking-[1px] tw-text-[0.75rem] tw-leading-[0.938rem] md:tw-text-[1rem] md:tw-leading-[1.1875rem] lg:tw-text-[1.188rem] lg:tw-leading-[1.438rem] tw-font-inter-semibold active:tw-scale-[0.98] portfolio-nav"
        :class="[
          currentRoutePath == navbar.url
            ? 'gradient-btn hover:tw-text-[#f2709c]'
            : 'tw-border-[1px] tw-border-solid tw-border-transparent hover:tw-from-[#ff835b] hover:tw-to-[#f2709c] hover:tw-bg-gradient-[270.11deg] tw-text-black-core/[0.6] hover:tw-text-white',

          index == navbars.length - 1 ? '' : 'sm:tw-mr-2.5',
        ]"
        :href="navbar.url"
        @click.native="mixpanel.track(navbar.event)"
        :target="navbar.target ? '_blank' : '_self'"
      >
        {{ navbar.name }}
      </a>
    </div>

    <div
      class="tw-container tw-py-12 md:tw-py-20 lg:tw-py-28 tw-text-[1.375rem] tw-leading-[1.625rem] md:tw-text-[1.9375rem] md:tw-leading-[1.812rem] lg:tw-text-[2.5rem] lg:tw-leading-[2rem] tw-text-black-core/[0.87] tw-text-center tw-font-inter-medium"
    >
      <span
        class="tw-relative after:tw-absolute after:-tw-bottom-[2px] after:tw-right-0 after:tw-content-[''] after:tw-w-[116px] md:after:tw-w-[152px] lg:after:tw-w-[205px] after:tw-h-[10px] md:after:tw-h-[14px] after:tw-bg-[#FFF3F7] after:-tw-z-[1] v2-canopas-gradient-text"
        >1M+ users</span
      >
      use our apps every day
    </div>
    <PortfolioSection ref="newPortfolio" />
    <PortfolioCTASection ref="newPortfolioCTA" />
  </section>
</template>

<script type="module">
import config from "@/config.js";
import PortfolioSection from "@/components/home-new/PortfolioSection.vue";
import PortfolioCTASection from "@/components/home-new/PortfolioCTASection.vue";
import { elementInViewPort } from "@/utils.js";

export default {
  data() {
    return {
      currentRoutePath: this.$router.currentRoute._value.path,
      event: "",
      navbars: [
        {
          name: "All",
          url: "/",
          target: false,
        },
        {
          name: "Justly",
          url: "/portfolio/justly",
          target: false,
          event: "tap_home_justly_portfolio",
        },
        {
          name: "Luxeradio",
          url: "/portfolio/luxeradio",
          target: false,
          event: "tap_home_luxe_radio_portfolio",
        },
        {
          name: "Togness",
          url: "/portfolio/togness",
          target: false,
          event: "tap_home_togness_portfolio",
        },
        {
          name: "Smile+",
          url: config.SMILEPLUS_URL,
          target: true,
          event: "tap_home_smile_portfolio",
        },
      ],
      events: {
        newPortfolio: "view_portfolio_section",
        newPortfolioCTA: "view_home_portfolio_cta",
      },
    };
  },
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  inject: ["mixpanel"],
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
  components: {
    PortfolioSection,
    PortfolioCTASection,
  },
};
</script>
