<template>
  <div
    class="text-center container max-w-[687px] mt-[1.875rem] mb-[14.063rem] mx-auto py-0 px-5 sm:px-[6.25rem] lg:px-0"
  >
    <div class="header-2-text mt-12 animate__animated animate__fadeInDown">
      <h1>
        Are you looking for a job where you can leverage your knowledge and
        experience to build digital products that will be used by millions of
        customers?
      </h1>
    </div>
    <div
      class="normal-text my-8 lg:my-12 animate__animated animate__fadeInLeft"
    >
      You've come to the right place. We're looking for an AVENGER who can join
      our existing team of AVENGERS.
    </div>
    <button
      class="gradient-btn rounded-[15px] sm:m-auto p-3 normal-text animate__animated animate__fadeInUp"
      @click="$emit('scroll-to-career')"
      @click.native="$mixpanel.track('tap_see_all_positions')"
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
