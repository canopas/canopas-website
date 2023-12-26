<template>
  <section
    class="bg-opacity-[0.08] bg-gradient-to-b from-white-core-1 to-pink-1 pb-14 lg:pb-[7.5rem] 3xl:outer-container"
  >
    <div
      class="container flex flex-col lg:flex-row items-center mt-[2.313rem] 2xl:mt-[5.313rem]"
    >
      <div class="flex flex-col 2xl:mt-8 lg:w-full xl:w-2/4 2xl:w-[42%]">
        <h1 class="mobile-header-1 lg:desk-header-1 text-black-87">
          Expertise and experience you can count on
        </h1>
        <nuxt-link
          class="mt-9 ml-0 hidden lg:flex gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>

      <ul
        class="mt-6 w-full md:w-[90%] lg:w-full xl:w-[55%] 2xl:w-[55%] sub-h1-regular lg:mobile-header-2-regular pl-0 text-black-60"
      >
        <li
          v-for="experience in experiences"
          :key="experience"
          class="flex py-[7.5px] md:py-[17.5px] lg:py-[6.5px] px-0"
        >
          <div v-html="experience"></div>
        </li>
        <li class="py-[7.5px] lg:py-[6.5px]">
          <router-link
            class="relative sub-h3-semibold lg:mobile-header-2-regular bg-gradient-[1deg] v2-canopas-gradient-text lg:gradient-underline lg:no-underline lg:after:absolute lg:after:bottom-0 lg:after:left-0 lg:after:w-full lg:after:h-0.5 lg:after:bg-gradient-to-l lg:after:to-pink-300 lg:after:from-orange-300"
            to="/contact"
            @click.native="$mixpanel.track('tap_about_landing_contact')"
          >
            Let's do it again</router-link
          >
          and make it even better this time with you on board.
        </li>
      </ul>
      <nuxt-link
        class="ml-0 md:mx-auto lg:hidden gradient-btn primary-btn"
        to="/contact"
        @click.native="$mixpanel.track('tap_landing_cta')"
      >
        <span class="sub-h3-semibold lg:sub-h1-semibold"
          >Get Free Consultation
        </span>
      </nuxt-link>
    </div>
  </section>
</template>
<script type="module">
export default {
  data() {
    return {
      experiences: [
        `<span class="sub-h1-semibold lg:mobile-header-2"> From IDEA to 10M Monthly Active Users.</span> We know how it works and what it takes.`,
        `<span class="sub-h1-semibold lg:mobile-header-2">From 3.6 App ratings to 5 star App ratings</span> in each store with over 100k installations, we know how you get there.`,
        `Reducing server bills from thousands of dollars to a few hundred dollars, we know how you get there.`,
        `<span class="sub-h1-semibold lg:mobile-header-2">Want to load your website under 2 seconds?</span> We’ve got you covered.`,
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
