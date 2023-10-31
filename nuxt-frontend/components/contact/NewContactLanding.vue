<template>
  <div class="pr-2">
    <div
      class="flex items-center py-7 md:py-12 lg:pb-5 lg:pt-0 font-product-sans font-bold text-[1.875rem] md:text-[2.5rem] leading-[2.274375rem] md:leading-[3.0325rem]"
    >
      <h1 class="v2-canopas-gradient-text pr-2">Hey Smarty</h1>
      <img
        :src="HandWing"
        alt="HandWing"
        class="w-[31.26px] md:w-[49.38px] h-[31.79px] md:h-[50.29px]"
      />
    </div>

    <div
      class="lg:my-10 font-inter-medium text-[1rem] md:text-[1.375rem] leading-6 md:leading-[2.0625rem] text-black-core/[0.87]"
    >
      Thanks for reaching out to us. This form requires
      <span
        class="font-inter-bold text-[1rem] md:text-[1.375rem] leading-6 md:leading-[2.0625rem] canopas-gradient-text"
        >less than 60 seconds</span
      >
      to fill up.
    </div>
    <div
      class="mt-7 font-inter-medium text-[1rem] md:text-[1.375rem] leading-6 md:leading-[2.0625rem] text-black-core/[0.87]"
    >
      Here's what you will get by filling out this form.
    </div>
    <div>
      <ul class="list-none m-auto mt-4 p-0 text-left">
        <li
          v-for="reasons in reasons"
          :key="reasons"
          class="flex items-start px-0 py-3"
        >
          <span
            class="mt-2 md:mt-3 mr-[15px] w-1.5 md:w-[8.81px] h-1.5 md:h-[9px] rounded-full bg-gradient-to-b from-pink-300 to-orange-300"
          ></span>
          <div
            class="w-11/12 lg:w-10/12 font-inter-medium text-[1rem] md:text-[1.375rem] leading-6 md:leading-[2.0625rem] text-black-core/[0.87]"
          >
            {{ reasons.title }}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import HandWing from "@/assets/images/theme/HandWing.png";
export default {
  data() {
    return {
      HandWing: HandWing,
      reasons: [
        {
          title:
            "Free consultation and QA session to know more about each other.",
        },
        {
          title:
            "  The call is only 15 minutes and we promise not to give you a hard\
              sell. We will talk more ONLY if you find the conversation\
              interesting.",
        },
        {
          title:
            " 100% MONEY BACK GUARANTEE for the first TWO WEEKS of work if\
              you're not satisfied. No questions asked.",
        },
      ],
    };
  },
};
</script>
