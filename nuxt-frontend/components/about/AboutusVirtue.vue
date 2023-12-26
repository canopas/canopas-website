<template>
  <section>
    <div
      class="container flex flex-col items-center gap-6 md:gap-10 pb-6 lg:pb-16"
    >
      <div class="mobile-header-2 lg:desk-header-2">Our Why?</div>
      <div
        class="w-[95%] lg:w-[62%] text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
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
