<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div
      class="tw-text-center v2-header-text tw-font-futura-bold tw-mt-8 lg:tw-mt-36"
    >
      Portfolio
    </div>
    <div
      class="tw-container tw-flex tw-justify-around tw-m-auto tw-mt-12 lg:tw-mt-44 tw-w-[100%] sm:tw-w-[95%] 2xl:tw-w-[45%] tw-rounded-[30px] tw-bg-white tw-py-[5px] tw-drop-shadow-[0_2px_10px_rgba(61,61,61,0.1)]"
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
    <NewPortfolioCTASection />
    <UserReview class="sm:tw-top-[40px]" />
    <CTASection ref="cta" />
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import CTA from "@/components/home/CTASection.vue";
import { useMeta } from "vue-meta";
import config from "@/config.js";
import { elementInViewPort } from "@/utils.js";
import PortfolioSection from "@/components/home-new/PortfolioSection.vue";
import NewPortfolioCTASection from "@/components/home-new/NewPortfolioCTASection.vue";
import UserReview from "@/components/home/UserReview.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import NewFooter from "@/components/partials/NewFooter.vue";

export default {
  data() {
    return {
      currentRoutePath: this.$router.currentRoute._value.path,
      event: "",
      navbars: [
        {
          name: "All",
          url: "/portfolio",
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
  setup() {
    var seoData = config.PORTFOLIO_SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
      },
    });
  },
  components: {
    Header,
    PortfolioSection,
    NewPortfolioCTASection,
    UserReview,
    CTASection,
    NewFooter,
  },
  inject: ["mixpanel"],
  mounted() {
    this.mixpanel.track("view_page_portfolio_list");
  },
};
</script>
