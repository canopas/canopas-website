<template>
  <section>
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center tw-gap-[1.5rem] md:tw-gap-[2.5rem] tw-pb-11 md:tw-pb-20"
    >
      <div
        class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.25rem] md:tw-leading-[5.15625rem] tw-tracking-[-4%]"
      >
        Our Why?
      </div>
      <div
        class="tw-w-[95%] lg:tw-w-[62%] tw-text-center tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.875rem] tw-leading-[1.5rem] md:tw-leading-[2.8125rem] tw-tracking-[-4%]"
      >
        To put humanity into the better position compared to what it was before
        our presence in the world.
      </div>
    </div>
    <div>
      <VirtuesViewSection />
    </div>
  </section>
</template>

<script type="module">
import VirtuesViewSection from "@/components/jobs/VirtuesView.vue";
export default {
  components: {
    VirtuesViewSection,
  },
  inject: ["mixpanel"],
};
</script>
