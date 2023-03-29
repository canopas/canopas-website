<template>
  <div class="tw-px-2">
    <div
      class="tw-flex tw-items-center tw-py-7 md:tw-py-12 lg:tw-pb-5 lg:tw-pt-0 tw-font-product-sans tw-font-bold tw-text-[1.875rem] md:tw-text-[2.5rem] tw-leading-[2.274375rem] md:tw-leading-[3.0325rem]"
    >
      <h1 class="v2-canopas-gradient-text tw-pr-2">Hey Smarty</h1>
      <img
        :src="HandWing"
        alt="HandWing"
        class="tw-w-[31.26px] md:tw-w-[49.38px] tw-h-[31.79px] md:tw-h-[50.29px]"
      />
    </div>

    <div
      class="lg:tw-my-10 tw-font-inter-medium tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.5rem] md:tw-leading-[2.0625rem] tw-text-black-core/[0.87]"
    >
      Thanks for reaching out to us. This form requires
      <span
        class="tw-font-inter-bold tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.5rem] md:tw-leading-[2.0625rem] canopas-gradient-text"
        >less than 60 seconds</span
      >
      to fill up.
    </div>
    <div
      class="tw-mt-7 tw-font-inter-medium tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.5rem] md:tw-leading-[2.0625rem] tw-text-black-core/[0.87]"
    >
      Here's what you will get by filling out this form.
    </div>
    <div>
      <ul class="tw-list-none tw-m-auto tw-mt-4 tw-p-0 text-left">
        <li
          v-for="reasons in reasons"
          :key="reasons"
          class="tw-flex tw-items-start tw-px-0 tw-py-3"
        >
          <span
            class="tw-mt-2 md:tw-mt-3 tw-mr-[15px] tw-w-[6px] md:tw-w-[8.81px] tw-h-[6px] md:tw-h-[9px] tw-rounded-full tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300"
          ></span>
          <div
            class="tw-w-11/12 lg:tw-w-10/12 tw-font-inter-medium tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.5rem] md:tw-leading-[2.0625rem] tw-text-black-core/[0.87]"
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
