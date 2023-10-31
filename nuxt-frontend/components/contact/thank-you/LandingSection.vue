<template>
  <section>
    <div class="container flex flex-col mt-4 lg:mt-12">
      <p
        class="text-center font-inter-regular text-[0.875rem] md:text-[1.125rem] xl:text-[1.375rem] leading-[1.313rem] md:leading-[1.688rem] xl:leading-[2.063rem] text-black-core/[0.6]"
      >
        Message from our CEO - <span class="font-inter-bold">{{ CEO }}</span>
      </p>
      <h1
        class="mt-4 md:mt-8 xl:mt-16 text-center text-[2.5rem] md:text-[3.75rem] xl:text-[5rem] leading-[3.25rem] md:leading-[4.688rem] xl:leading-[7.5rem] font-inter-bold text-black-core/[0.87]"
      >
        Thank You
        <span
          :class="{
            'text-[1.5rem] sm:text-[2rem] lg:text-[3.75rem] px-8 sm:pl-16 lg:pl-24':
              isLongText,
          }"
          class="text-white brush-wrap relative inline-block text-center px-12 py-4"
          ><PaintStroke />{{ truncatedName }}
        </span>

        for your Interest in my company.
      </h1>
      <p
        class="mt-8 xl:mt-12 text-center font-inter-medium text-[1rem] md:text-[1.563rem] xl:text-[1.875rem] leading-[1.5rem] md:leading-[2.188rem] xl:leading-[2.813rem] text-black-core/[0.6]"
      >
        I realize that you had a number of other choices, but I really
        appreciate the fact that you chose my company.
      </p>
    </div>
    <div
      class="mt-10 xl:mt-[4.5rem] flex flex-col bg-gradient-to-r from-[#FF835B]/[0.08] to-[#F2709C]/[0.08] py-12 xl:py-16"
    >
      <div class="container">
        <p
          class="v2-canopas-gradient-text font-inter-bold text-center text-[1.5rem] md:text-[1.75rem] xl:text-[2rem] leading-9 md:leading-[2.625rem] xl:leading-[3rem]"
        >
          At Canopas, Our Vision is
        </p>
        <p
          class="mt-2 text-center text-[1rem] md:text-[1.563rem] xl:text-[2.5rem] leading-6 md:leading-9 xl:leading-[3.75rem] font-inter-medium text-black-core/[0.87]"
        >
          "To put humanity and the world into a better position compared to what
          it was prior to our existence in the world."
        </p>
      </div>
    </div>
  </section>
</template>
<script>
import PaintStroke from "@/components/utils/PaintBrushStroke.vue";

export default {
  props: ["name"],
  data() {
    return {
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
