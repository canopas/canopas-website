<template>
  <section
    :class="
      currentRoutePath == '/portfolio' ? 'mt-16 lg:mt-28' : 'mt-16 lg:mt-60'
    "
  >
    <div>
      <p
        class="text-[#F8F8F8] background-text text-center"
        :class="currentRoutePath == '/portfolio' ? 'hidden' : 'block'"
      >
        Our Portfolio
      </p>
      <h2
        class="container -mt-5 md:mt-[-50px] mobile-header-2 lg:desk-header-2 text-center text-black-87"
      >
        Case studies
      </h2>
    </div>

    <div
      class="container mt-4 mb-6 text-black-60 text-center sub-h1-regular lg:mobile-header-2-regular"
    >
      <span
        class="sub-h1-semibold lg:mobile-header-2-semibold relative after:absolute after:-bottom-0.5 after:right-0 after:content-[''] after:w-[89px] md:after:w-[90px] lg:after:w-[108px] xl:after:w-[120px] after:h-2.5 md:after:h-3.5 after:bg-[#FFF3F7] after:-z-[1] v2-canopas-gradient-text"
        >1M+ users</span
      >
      use our apps every day
    </div>
    <PortfolioSection ref="newPortfolio" />
    <CTA4 ref="portfolioPageCTA" v-if="currentRoutePath == '/portfolio'" />
    <PortfolioCTASection ref="newPortfolioCTA" v-else />
  </section>
</template>

<script type="module">
import config from "@/config.js";
import PortfolioSection from "@/components/home-new/PortfolioSection.vue";
import PortfolioCTASection from "@/components/home-new/PortfolioCTASection.vue";
import CTA4 from "@/components/CTA/CTA4.vue";

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
    CTA4,
  },
};
</script>
