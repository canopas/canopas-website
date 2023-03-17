<template>
  <div class="tw-relative tw-w-full">
    <div
      class="tw-absolute tw-top-[0.3rem] md:tw-top-[0rem] md:-tw-bottom-[4rem] tw-w-full tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-text-center xl:tw-text-right tw-p-[1rem] md:tw-p-[0.5rem] tw-bg-pink-500 tw-font-inter-semibold tw-text-white tw-text-[1rem] md:tw-text-[1.5rem] xl:tw-text-[1.875rem] xl:tw-leading-[2.813rem]"
    >
      <span class="lg:tw-mr-[-13rem] xl:tw-mr-16"> {{ message }}</span>
    </div>
  </div>
</template>

<script type="module">
export default {
  data() {
    return {
      message:
        "Thank you for choosing us to make a difference in your business.",
      currentRoutePath: this.$router.currentRoute._value.path,
    };
  },
  mounted() {
    console.log(this.currentRoutePath);
    if (this.currentRoutePath.includes("/jobs/apply")) {
      this.message =
        "We have received your job application, sit back and relax!";
    }
  },
};
</script>
