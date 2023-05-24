<template>
  <section>
    <div class="tw-container tw-flex tw-flex-col tw-mt-10 lg:tw-mt-12">
      <p
        class="tw-text-center tw-font-inter-regular tw-text-[0.875rem] md:tw-text-[1.125rem] xl:tw-text-[1.375rem] tw-leading-[1.313rem] md:tw-leading-[1.688rem] xl:tw-leading-[2.063rem] tw-text-black-core/[0.6]"
      >
        Message from our CEO - <span class="tw-font-inter-bold">{{ CEO }}</span>
      </p>
      <h1
        class="tw-mt-4 tw-text-center tw-text-[2.5rem] md:tw-text-[3.75rem] xl:tw-text-[5rem] tw-leading-[3.25rem] md:tw-leading-[4.688rem] xl:tw-leading-[7.5rem] tw-font-inter-bold tw-text-black-core/[0.87]"
      >
        Thank You
        <span
          :class="{
            'tw-text-[1.5rem] sm:tw-text-[2rem] lg:tw-text-[3.75rem] tw-px-[2rem] sm:tw-pl-[4rem] lg:tw-pl-[6rem]':
              isLongText,
          }"
          class="tw-text-white brush-wrap tw-relative tw-inline-block tw-text-center tw-px-[3rem] tw-py-[1rem]"
          ><PaintStroke />{{ truncatedName }}
        </span>

        for your Interest in my company.
      </h1>
      <p
        class="tw-mt-4 tw-text-center tw-font-inter-medium tw-text-[1rem] md:tw-text-[1.563rem] xl:tw-text-[1.875rem] tw-leading-[1.5rem] md:tw-leading-[2.188rem] xl:tw-leading-[2.813rem] tw-text-black-core/[0.6]"
      >
        I realize that you had a number of other choices, but I really
        appreciate the fact that you chose my company.
      </p>
    </div>
    <div
      class="tw-mt-6 xl:tw-mt-10 tw-flex tw-flex-col tw-bg-gradient-to-r tw-from-[#FF835B]/[0.08] tw-to-[#F2709C]/[0.08] tw-py-12"
    >
      <div class="tw-container">
        <p
          class="v2-canopas-gradient-text tw-font-inter-bold tw-text-center tw-text-[1.5rem] md:tw-text-[1.75rem] xl:tw-text-[2rem] tw-leading-[2.25rem] md:tw-leading-[2.625rem] xl:tw-leading-[3rem]"
        >
          At Canopas, Our Vision is
        </p>
        <p
          class="tw-mt-2 tw-text-center tw-text-[1rem] md:tw-text-[1.563rem] xl:tw-text-[2.5rem] tw-leading-[1.5rem] md:tw-leading-[2.5rem] xl:tw-leading-[3.75rem] tw-font-inter-medium tw-text-black-core/[0.87]"
        >
          "To put humanity and the world into a better position compared to what
          it was prior to our existence in the world."
        </p>
      </div>
    </div>
  </section>
</template>
<script>
import landingbg400 from "@/assets/images/thankyou/landingbg-400w.webp";
import PaintStroke from "@/components/utils/PaintBrushStroke.vue";

export default {
  props: ["name"],
  data() {
    return {
      landingbg400,
      CEO: "Jimmy",
    };
  },
  components: {
    PaintStroke,
  },
  computed: {
    truncatedName() {
      if (this.name.length > 15) {
        return this.name.substring(0, 15);
      }
      return this.name;
    },
    isLongText() {
      return this.name.length > 7;
    },
  },
};
</script>
