<template>
  <div
    class="bg-[#F8F8F8] mt-14 text-center lg:text-left lg:h-[552px] 3xl:outer-container"
  >
    <div
      class="container py-14 lg:py-20 xl:py-12 2xl:py-0 flex flex-col gap-4 lg:flex-row lg:gap-[3.12rem]"
    >
      <div
        class="xs:mx-4 lg:!mx-0 flex flex-col gap-4 lg:gap-6 flex-[52.69%] my-auto"
      >
        <p class="header-1 xl:header-2 text-black-core/[0.87]">
          Crafting tomorrow's technology starts with one click.
          <span class="border-text"> Will you take the leap? </span>
        </p>
        <span class="text-black-core/[0.6] sub-h1-regular"
          >Discover how we can elevate your project.</span
        >
        <nuxt-link
          class="gradient-btn primary-btn hidden lg:flex mx-0"
          to="/contact"
          @click.native="$mixpanel.track('tap_flutter_development_cta')"
        >
          <span class="sub-h3-semibold border-none">Drop Us a Line </span>
          <Icon name="fa6-solid:arrow-right" class="fa ml-3" />
        </nuxt-link>
      </div>
      <div class="mx-auto lg:mx-0 flex-[40%] lg:flex justify-end">
        <img
          :src="leap"
          :srcset="`${leap400} 400w, ${leap800} 800w`"
          loading="lazy"
          alt="Drop-us-a-line-image"
          class="h-full w-[20.42463rem] lg:w-full object-contain"
        />
        <nuxt-link
          class="gradient-btn primary-btn lg:hidden"
          to="/contact"
          @click.native="$mixpanel.track('tap_flutter_development_cta')"
        >
          <span class="sub-h3-semibold border-none">Drop Us a Line </span>
          <Icon name="fa6-solid:arrow-right" class="fa ml-3" />
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import leap from "@/assets/images/flutter-app-development/leap.webp";
import leap400 from "@/assets/images/flutter-app-development/leap400.webp";
import leap800 from "@/assets/images/flutter-app-development/leap800.webp";
</script>
<style lang="postcss" scoped>
.border-text {
  -webkit-text-stroke: 1px rgba(0, 0, 0, 0.87);
}
@media (min-width: 992px) {
  .border-text {
    -webkit-text-stroke: 1.5px rgba(0, 0, 0, 0.87);
  }
}
</style>
