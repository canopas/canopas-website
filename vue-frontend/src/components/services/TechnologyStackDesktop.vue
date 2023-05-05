<template>
  <section
    class="lg:tw-mt-[18rem] xl:tw-mt-[21rem] tw-mb-20 md:tw-mb-40 tw-flex"
  >
    <div
      class="tw-container tw--mt-48 tw-mb-28 tw-mx-auto tw-w-8/12 tw-text-center"
    >
      <span
        class="tw-font-inter-bold tw-text-black-core/[0.87] tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.25rem] md:tw-leading-[5.15625rem]"
        >Technology Stack</span
      >
      <p
        class="tw-mt-5 tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1.5rem] tw-leading-[1.5rem] lg:tw-leading-[2.25rem]"
      >
        The products we build are typically deployed to mobile and web. Web and
        mobile are our bread and butter. From native to hybrid, frontend to
        backend, designing to deployment, we've got you covered. At Canopas,
        we're not bound by technologies. We learn new technology if it has
        chances to improve your app performance significantly.
      </p>
    </div>

    <div class="tw-relative 2xll:tw-container">
      <aspect-ratio height="40%">
        <img
          class="tw-w-ful tw-h-full tw-object-cover"
          :src="techicon"
          alt="techicon-view-image"
        />
      </aspect-ratio>
      <div
        :class="{
          'tw-opacity-50 tw-transition-opacity tw-duration-1000 tw-ease-in':
            isHovered,
        }"
      >
        <img
          :class="
            width > 1300 || width < 1600
              ? 'tw-w-[425px]   xl:tw-w-[476px]  2xl:tw-w-[590px] 2xll:tw-w-[500px]    tw-h-[385px] xl:tw-h-[413px] 2xl:tw-h-[515px]   2xll:tw-h-[440px] xll:tw-h-[445px] '
              : ''
          "
          class="tw-absolute tw-inset-0 tw-m-auto"
          :src="ellipse"
          alt="ellipse-view-image"
        />
        <img
          class="tw-absolute tw-w-[220px] tw-h-[315px] xl:tw-w-[285px] xl:tw-h-[454px] tw-inset-0 tw-m-auto tw-top-[75%]"
          :src="graphic"
          alt="graphic-view-image"
        />
      </div>
      <div
        v-for="service in service"
        :key="service.id"
        class="tw-w-fit tw-h-fit tw-flex tw-justify-center tw-items-center"
      >
        <img
          class="tw-absolute tw-z-[1] tw-w-[100px] tw-h-[100x]"
          :class="{
            'tw-opacity-10 tw-transition-opacity tw-duration-700 tw-ease-in':
              isHovered,
          }"
          :src="service.image"
          :style="service.style"
          alt="service-view-image"
          @mouseover="(activeIndex = service.id), (isHovered = true)"
          @mouseleave="(activeIndex = service.id), (isHovered = false)"
          @touchstart.passive="activeIndex = service.id"
        />
        <img
          class="tw-absolute tw-w-[417px] tw-h-[460px] tw-opacity-0"
          :src="service.hoverimg"
          alt="service-hover-image"
          :style="service.style"
          :class="
            activeIndex == service.id
              ? 'tw-opacity-100  tw-transition tw-duration-700 tw-ease-in'
              : ''
          "
        />
      </div>
    </div>
  </section>
</template>

<script>
import backend from "@/assets/images/services/technology/backend.webp";
import backendHover from "@/assets/images/services/technology/backendHover.webp";
import frontend from "@/assets/images/services/technology/frontend.webp";
import frontendHover from "@/assets/images/services/technology/frontendHover.webp";
import graphic from "@/assets/images/services/technology/graphic.webp";
import mobileapp from "@/assets/images/services/technology/mobileapp.webp";
import mobileappHover from "@/assets/images/services/technology/mobileappHover.webp";
import seo from "@/assets/images/services/technology/seo.webp";
import seoHover from "@/assets/images/services/technology/seoHover.webp";
import techicon from "@/assets/images/services/technology/techicon.webp";
import ux from "@/assets/images/services/technology/ux.webp";
import uxHover from "@/assets/images/services/technology/uiHover.webp";
import ellipse from "@/assets/images/services/technology/ellipse.webp";
import AspectRatio from "@/components/utils/AspectRatio.vue";
export default {
  data() {
    return {
      isHovered: false,
      graphic,
      techicon,
      ellipse,
      frontendHover,
      width: 0,
      activeIndex: 0,
      service: [
        {
          id: 1,
          image: frontend,
          style: "left:31%;top: 63%; transform: translate(-50%, -50%) ;",
          hoverimg: frontendHover,
        },
        {
          id: 2,
          image: seo,
          style: "left:69%; top: 63%; transform: translate(-50%, -50%) ;",
          hoverimg: seoHover,
        },

        {
          id: 3,
          image: backend,
          style: "left:34%; top:26%;  transform: translate(-50%, -50%);",
          hoverimg: backendHover,
        },
        {
          id: 4,
          image: mobileapp,
          style: "left: 50%; top:8.5%; transform: translate(-50%, -50%) ;",
          hoverimg: mobileappHover,
        },
        {
          id: 5,
          image: ux,
          style: "left:65%;   top:28%;  transform: translate(-50%, -50%) ;",
          hoverimg: uxHover,
        },
      ],
    };
  },
  components: {
    AspectRatio,
  },
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>
