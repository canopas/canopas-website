<template>
  <div
    class="tw-fixed tw-top-[10%] tw-w-full tw-z-[500] tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-px-4 tw-py-3 tw-font-inter-semibold tw-text-center tw-text-white tw-text-[1rem] md:tw-text-[1.5rem] xl:tw-text-[1.875rem] xl:tw-leading-[2.813rem]"
    role="alert"
  >
    <span class="tw-block sm:tw-inline">{{ message }}</span>
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
    if (this.currentRoutePath.includes("/jobs/apply")) {
      this.message =
        "Thank you for applying at Canopas, we will get in touch with you soon!";
    }
  },
};
</script>
