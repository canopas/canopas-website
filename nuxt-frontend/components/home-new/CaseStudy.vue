<template>
  <section class="mt-[7.813rem] font-inter-regular">
    <div
      v-if="currentRoutePath == '/portfolio'"
      class="text-center mt-8 lg:mt-36 pb-9 md:pb-20 text-[#3d3d3d] font-futura-bold text-[4.5rem] md:text-[5.625rem] lg:text-[7.5rem] leading-[6rem] md:leading-[7.625rem] lg:leading-[9rem]"
    >
      <h1>Portfolio</h1>
    </div>
    <div v-else>
      <p
        class="text-black-core/[.03] text-[2.5rem] md:text-[3.75rem] xl:text-[5rem] leading-[1.875rem] md:leading-[4.0625rem] xl:leading-[5rem] text-center"
      >
        Our Portfolio
      </p>
      <h2
        class="container -mt-5 md:mt-[-50px] pb-9 md:pb-12 text-[1.875rem] leading-[2.438rem] md:text-[2.65rem] md:leading-[3.218rem] lg:text-[3.438rem] lg:leading-[4rem] text-center text-black-core/[0.87] font-roboto-bold"
      >
        Case Studies
      </h2>
    </div>
    <div
      class="container flex justify-around m-auto w-[100%] sm:w-[95%] 2xl:w-[45%] rounded-[30px] bg-white py-[5px] drop-shadow-[0_2px_10px_rgba(61,61,61,0.1)]"
    >
      <a
        v-for="(navbar, index) in navbars"
        :key="index"
        class="flex-[16.66%] m-0 w-full rounded-[25px] px-[5px] py-2 md:py-[13px] text-center tracking-[1px] text-[0.75rem] leading-[0.938rem] md:text-[1rem] md:leading-[1.1875rem] lg:text-[1.188rem] lg:leading-[1.438rem] font-inter-semibold active:scale-[0.98] portfolio-nav"
        :class="[
          currentRoutePath == navbar.url
            ? 'gradient-btn hover:text-[#f2709c]'
            : 'border border-solid border-transparent hover:from-[#ff835b] hover:to-[#f2709c] hover:bg-gradient-[270.11deg] text-black-core/[0.6] hover:text-white',

          index == navbars.length - 1 ? '' : 'sm:mr-2.5',
        ]"
        :href="navbar.url"
        @click.native="$mixpanel.track(navbar.event)"
        :target="navbar.target ? '_blank' : '_self'"
      >
        {{ navbar.name }}
      </a>
    </div>

    <div
      class="container py-12 md:py-20 lg:py-28 text-[1.375rem] leading-[1.625rem] md:text-[1.9375rem] md:leading-[1.812rem] lg:text-[2.5rem] lg:leading-[2rem] text-black-core/[0.87] text-center font-inter-medium"
    >
      <span
        class="relative after:absolute after:-bottom-0.5 after:right-0 after:content-[''] after:w-[116px] md:after:w-[152px] lg:after:w-[205px] after:h-2.5 md:after:h-3.5 after:bg-[#FFF3F7] after:-z-[1] v2-canopas-gradient-text"
        >1M+ users</span
      >
      use our apps every day
    </div>
    <PortfolioSection ref="newPortfolio" />
    <PortfolioPageCTASection
      ref="portfolioPageCTA"
      v-if="currentRoutePath == '/portfolio'"
    />
    <PortfolioCTASection ref="newPortfolioCTA" v-else />
  </section>
</template>

<script type="module">
import config from "@/config.js";
import PortfolioSection from "@/components/home-new/PortfolioSection.vue";
import PortfolioCTASection from "@/components/home-new/PortfolioCTASection.vue";
import PortfolioPageCTASection from "@/components/home-new/PortfolioPageCTASection.vue";

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
        portfolioPageCTA: "view_portfolio_page_cta",
      },
    };
  },
  mounted() {
    if (this.currentRoutePath == "/portfolio") {
      this.navbars[0].url = "/portfolio";
    }
  },
  inject: ["mixpanel"],
  components: {
    PortfolioSection,
    PortfolioCTASection,
    PortfolioPageCTASection,
  },
};
</script>
