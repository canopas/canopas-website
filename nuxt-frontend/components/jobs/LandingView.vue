<template>
  <div class="text-center container mt-8 lg:mt-[7.5rem] mb-16 lg:mb-60 mx-auto">
    <div
      class="xl:w-[89%] 2xl:w-4/5 xl:mx-auto mobile-header-1 lg:desk-header-2 text-black-87 animate__animated animate__fadeInDown"
    >
      <h1>
        Are you looking for a job where you can leverage your knowledge and
        experience to build digital products that will be used by millions of
        customers?
      </h1>
    </div>
    <div
      class="xl:w-[71%] 2xl:w-[65%] xl:mx-auto sub-h1-regular lg:mobile-header-2-regular text-black-60 mt-4 lg:mt-6 animate__animated animate__fadeInLeft"
    >
      You've come to the right place. We're looking for an AVENGER who can join
      our existing team of AVENGERS.
    </div>
    <button
      class="mt-6 lg:mt-10 gradient-btn primary-btn"
      @click="$emit('scroll-to-career')"
      @click.native="$mixpanel.track('tap_see_all_positions')"
    >
      <span class="sub-h3-semibold lg:sub-h1-semibold">SEE ALL POSITIONS</span>
    </button>
  </div>
</template>

<script>
export default {
  inject: ["mixpanel"],
};
</script>
