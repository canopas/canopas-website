<template>
  <section>
    <div
      class="tw-flex tw-flex-col lg:tw-flex-row tw-mt-[2.313rem] tw-items-center"
    >
      <div
        class="md:tw-w-[76%] lg:tw-w-[55%] tw-text-center tw-text-[2.5rem] md:tw-text-[3.75rem] lg:tw-text-[5rem] tw-leading-[3.25rem] md:tw-leading-[4.5rem] lg:tw-leading-[6rem] tw-font-inter-semibold tw-text-black-core/[0.87]"
      >
        Expertise and experience you can count on
      </div>

      <ul
        class="tw-w-full md:tw-w-[63%] lg:tw-w-[45%] tw-px-4 md:tw-px-0 tw-mt-6 tw-text-[1rem] md:tw-text-[1.125rem] tw-leading-[1.5rem] md:tw-leading-[1.625rem] tw-text-left tw-text-black-core/[.87] tw-font-inter-medium"
      >
        <li
          v-for="experience in experiences"
          :key="experience"
          class="tw-flex tw-py-[7.5px] md:tw-py-[17.5px] tw-px-0"
        >
          <div>
            {{ experience }}
          </div>
        </li>
        <li class="tw-py-[7.5px]">
          <span
            class="tw-relative tw-bg-gradient-[1deg] v2-canopas-gradient-text gradient-underline tw-no-underline after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-[2px] after:tw-bg-gradient-to-l after:tw-to-pink-300 after:tw-from-orange-300"
          >
            Let's do it again</span
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
};
</script>
