<template>
  <div class="background">
    <div class="container tw-pb-24 tw-pt-10 sm:tw-pt-20">
      <div class="tw-flex tw-items-center">
        <div
          class="v2-title-2-text tw-font-bold v2-canopas-gradient-text tw-mr-4"
        >
          Hey Smarty
        </div>
        <img :src="HandWing" class="tw-w-8 tw-h-8 lg:tw-w-12 lg:tw-h-12" />
      </div>
      <div class="v2-normal-text tw-pt-8">
        Thanks for reaching out to us. We're happy to help you achieve your
        goals and dreams.
      </div>
      <div class="v2-normal-text tw-pt-6">
        This form requires
        <span class="canopas-gradient-text">less then 60 seconds</span>
        to fill up. Here's what you will get by filling out this form.
      </div>
      <div>
        <ul class="tw-list-none tw-m-auto tw-p-0 text-left mt-3">
          <li class="tw-flex tw-items-start tw-px-0 tw-py-2">
            <span
              class="tw-h-3.5 tw-w-3.5 tw-mr-6 tw-rounded tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-mt-0.5 md:tw-mt-2 md:tw-h-4 md:tw-w-4"
            ></span>
            <div class="tw-w-11/12 lg:tw-w-10/12 v2-normal-2-text tw-font-bold">
              Free consultation for 30 Minutes.
            </div>
          </li>
          <li class="tw-flex tw-items-start tw-px-0 tw-py-2">
            <span
              class="tw-h-3.5 tw-w-3.5 tw-mr-6 tw-rounded tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-mt-0.5 md:tw-mt-2 md:tw-h-4 md:tw-w-4"
            ></span>
            <div class="tw-w-11/12 lg:tw-w-10/12 v2-normal-2-text tw-font-bold">
              100% MONEY BACK GUARANTEE for the first TWO WEEKS of work if
              you're not satisfied. No questions asked.
            </div>
          </li>
          <li class="tw-flex tw-items-start tw-px-0 tw-py-2">
            <span
              class="tw-h-3.5 tw-w-3.5 tw-mr-6 tw-rounded tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-mt-0.5 md:tw-mt-2 md:tw-h-4 md:tw-w-4"
            ></span>
            <div class="tw-w-11/12 lg:tw-w-10/12 v2-normal-2-text tw-font-bold">
              An opportunity to grow your business exponentially.
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import HandWing from "@/assets/images/theme/HandWing.png";
export default {
  data() {
    return {
      HandWing: HandWing,
    };
  },
};
</script>

<style lang="postcss" scoped>
.background {
  background: linear-gradient(
    270.11deg,
    rgba(249, 131, 134, 0.05) -24.42%,
    rgba(249, 131, 134, 0.05) 101.76%
  );
  box-shadow: 0px 4px 2px 1px rgba(0, 0, 0, 5%);
  z-index: -1;
}
</style>
