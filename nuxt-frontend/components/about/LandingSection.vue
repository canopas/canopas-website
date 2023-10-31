<template>
  <section
    class="opacity-80 bg-gradient-to-b from-[#FFFFFF]/[0.1] via-[#FF835B]/[0.1] to-[#F2709C]/[0.1] pb-10 lg:pb-[5.5rem]"
  >
    <div
      class="container flex flex-col lg:flex-row items-center mt-[2.313rem] 2xl:mt-[5.313rem]"
    >
      <h1
        class="2xl:mt-8 lg:w-[55%] xl:w-[50%] 2xl:w-[58%] text-center lg:text-left text-[2.5rem] md:text-[3.75rem] lg:text-[4rem] 2xl:text-[5rem] leading-[3.25rem] md:leading-[4.5rem] lg:leading-[5rem] 2xl:leading-[6rem] font-inter-semibold text-black-core/[0.87]"
      >
        Expertise and experience you can count on
      </h1>

      <ul
        class="mt-6 w-full sm:w-[85%] md:w-[70%] lg:w-[45%] 2xl:w-[42%] text-left text-[1rem] md:text-[1.125rem] 2xl:text-[1.188rem] leading-6 md:leading-[1.625rem] 2xl:leading-[1.781rem] text-black-core/[.87] font-inter-medium"
      >
        <li
          v-for="experience in experiences"
          :key="experience"
          class="flex py-[7.5px] md:py-[17.5px] lg:py-[6.5px] px-0"
        >
          <div>
            {{ experience }}
          </div>
        </li>
        <li class="py-[7.5px] lg:py-[6.5px]">
          <router-link
            class="relative bg-gradient-[1deg] v2-canopas-gradient-text gradient-underline no-underline after:absolute after:bottom-0 after:left-0 after:w-full after:h-[2px] after:bg-gradient-to-l after:to-pink-300 after:from-orange-300"
            to="/contact"
            @click.native="$mixpanel.track('tap_about_landing_contact')"
          >
            Let's do it again</router-link
          >
          and make it even better this time with you on board.
        </li>
      </ul>
    </div>
  </section>
</template>
<script type="module">
export default {
  data() {
    return {
      experiences: [
        "From IDEA to 10M Monthly Active Users. We know how it works and what it takes.",
        "From 3.6 App ratings to 5 star App ratings in each store with over 100k installations, we know how you get there.",
        "Reducing server bills from thousands of dollars to a few hundred dollars, we know how you get there.",
        "Want to load your website under 2 seconds? We’ve got you covered.",
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
