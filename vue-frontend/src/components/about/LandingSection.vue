<template>
  <section
    class="tw-opacity-80 tw-bg-gradient-to-b tw-from-[#FFFFFF]/[0.1] tw-via-[#FF835B]/[0.1] tw-to-[#F2709C]/[0.1] tw-pb-[2.5rem] lg:tw-pb-[5.5rem]"
  >
    <div
      class="tw-container tw-flex tw-flex-col lg:tw-flex-row tw-items-center tw-mt-[2.313rem] 2xl:tw-mt-[5.313rem]"
    >
      <h1
        class="2xl:tw-mt-[2rem] lg:tw-w-[55%] xl:tw-w-[50%] 2xl:tw-w-[58%] tw-text-center lg:tw-text-left tw-text-[2.5rem] md:tw-text-[3.75rem] lg:tw-text-[4rem] 2xl:tw-text-[5rem] tw-leading-[3.25rem] md:tw-leading-[4.5rem] lg:tw-leading-[5rem] 2xl:tw-leading-[6rem] tw-font-inter-semibold tw-text-black-core/[0.87]"
      >
        Expertise and experience you can count on
      </h1>

      <ul
        class="tw-mt-6 tw-w-full sm:tw-w-[85%] md:tw-w-[70%] lg:tw-w-[45%] 2xl:tw-w-[42%] tw-text-left tw-text-[1rem] md:tw-text-[1.125rem] 2xl:tw-text-[1.188rem] tw-leading-[1.5rem] md:tw-leading-[1.625rem] 2xl:tw-leading-[1.781rem] tw-text-black-core/[.87] tw-font-inter-medium"
      >
        <li
          v-for="experience in experiences"
          :key="experience"
          class="tw-flex tw-py-[7.5px] md:tw-py-[17.5px] lg:tw-py-[6.5px] tw-px-0"
        >
          <div>
            {{ experience }}
          </div>
        </li>
        <li class="tw-py-[7.5px] lg:tw-py-[6.5px]">
          <router-link
            class="tw-relative tw-bg-gradient-[1deg] v2-canopas-gradient-text gradient-underline tw-no-underline after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-[2px] after:tw-bg-gradient-to-l after:tw-to-pink-300 after:tw-from-orange-300"
            to="/contact"
            @click.native="mixpanel.track('tap_about_landing_contact')"
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
