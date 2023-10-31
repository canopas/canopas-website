<template>
  <section>
    <div
      class="container flex flex-col items-center gap-6 md:gap-10 pb-11 md:pb-20"
    >
      <div
        class="font-inter-bold text-[1.875rem] md:text-[3.4375rem] leading-9 md:leading-[5.15625rem]"
      >
        Our Why?
      </div>
      <div
        class="w-[95%] lg:w-[62%] text-center font-inter-regular text-[1rem] md:text-[1.875rem] leading-6 md:leading-[2.8125rem]"
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
};
</script>
