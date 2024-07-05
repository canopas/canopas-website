<template>
  <section class="flex flex-col gap-40">
    <div class="container flex flex-col mt-4 lg:mt-12">
      <p
        class="text-center sub-h4-regular lg:mobile-header-3-regular text-black-60"
      >
        Message from our CEO -
        <span class="sub-h4-medium lg:mobile-header-3-semibold">{{ CEO }}</span>
      </p>
      <h1
        class="mt-4 lg:mt-6 text-center mobile-header-1 lg:desk-header-1 text-black-87 xl:w-4/5 2xl:w-[70%] xl:mx-auto"
      >
        Thank You
        <span
          :class="{
            'px-8 sm:pl-16 lg:pl-24': isLongText,
          }"
          class="text-white brush-wrap relative inline-block text-center px-12 py-4"
          ><PaintStroke />{{ truncatedName }}
        </span>

        for your Interest in my company.
      </h1>
      <p
        class="mt-4 lg:mt-6 text-center sub-h1-regular lg:mobile-header-2-regular text-black-60 xl:w-[70%] 2xl:w-[62%] xl:mx-auto"
      >
        I realize that you had a number of other choices, but I really
        appreciate the fact that you chose my company.
      </p>
    </div>
    <!-- <div class="bg-primary-1 py-5">
        <p class="container text-center text-white text-xl font-semibold tracking-wider">
          Thank you for booking a meeting with Canopas. Here's the details of your
          meeting with Canopas.
          <br />
          Time - 9:00am - 9:30am, Wednesday, June 26, 2024 (India Standard Time)
        </p>
    </div> -->
    <div class="container">
      <p class="text-center mobile-header-2 lg:desk-header-2">Our Vision</p>
      <p
        class="mt-3 text-center sub-h3-medium lg:desk-header-3 text-black-87 w-[68%] mx-auto"
      >
        "To put humanity and the world into a better position compared to what
        it was prior to our existence in the world."
      </p>
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
