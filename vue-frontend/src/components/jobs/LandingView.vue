<template>
  <div
    class="tw-text-center tw-container tw-max-w-[687px] tw-mt-[30px] tw-mb-[75px] tw-mx-auto tw-py-0 tw-px-[20px] sm:tw-px-[100px] lg:tw-px-0"
  >
    <div class="header-2-text tw-mt-12 animate__animated animate__fadeInDown">
      Are you looking for a job where you can leverage your knowledge and
      experience to build digital products that will be used by millions of
      customers?
    </div>
    <div
      class="normal-text tw-my-8 lg:tw-my-12 animate__animated animate__fadeInLeft"
    >
      You've come to the right place. We're looking for an AVENGER who can join
      our existing team of AVENGERS.
    </div>
    <button
      class="gradient-btn tw-rounded-[15px] sm:tw-m-auto normal-text animate__animated animate__fadeInUp"
      @click="$emit('scroll-to-career')"
      @click.native="mixpanel.track('tap_see_all_positions')"
    >
      <span>SEE ALL POSITIONS</span>
    </button>
  </div>
</template>

<script>
export default {
  inject: ["mixpanel"],
};
</script>
