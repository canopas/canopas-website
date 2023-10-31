<template>
  <section class="container mb-28">
    <div class="pt-20 pb-8 xl:pb-20 text-center">
      <span
        class="font-inter-bold text-[1.875rem] md:text-[3.4375rem] leading-[2.4375rem] md:leading-[5.15625rem]"
        >Interview Process
      </span>
    </div>
    <div
      v-for="interview in interviews"
      :key="interview"
      class="block md:hidden -mb-7"
    >
      <img
        alt="numeric-block-image"
        :src="interview.image"
        class="relative mx-auto w-[226px] h-[260px] object-contain"
      />

      <span
        :class="interview.class"
        v-html="interview.content"
        class="absolute w-44 text-center text-black-core/[0.87] font-inter-semibold text-[1.25rem] leading-[1.875rem] left-1/2 transform -translate-x-1/2"
      ></span>
    </div>
    <div
      class="hidden md:flex md:flex-wrap md:items-center md:justify-center md:mb-28"
    >
      <div v-for="interview in interviews" :key="interview">
        <img
          alt="numeric-block-image"
          :src="interview.desktopImage"
          class="relative mx-auto w-[232px] lg:w-[225px] xl:w-[320px] h-[200px] xl:h-[237px] object-contain"
        />
        <span
          v-html="interview.content"
          :class="interview.desktopClass"
          class="absolute w-[10rem] xl:w-[12.5rem] text-center text-black-core/[0.87] font-inter-semibold text-[1.3rem] xl:text-[1.625rem] leading-[1.75rem] xl:leading-[2.438rem] left-1/2 transform"
        ></span>
      </div>
    </div>
  </section>
</template>
<script>
import DesktopOne from "@/assets/images/jobsthankyou/desktop1.webp";
import Desktoptwo from "@/assets/images/jobsthankyou/desktop2.webp";
import Desktopthree from "@/assets/images/jobsthankyou/desktop3.webp";
import MobileOne from "@/assets/images/jobsthankyou/mobile1.webp";
import MobileTwo from "@/assets/images/jobsthankyou/mobile2.webp";
import MobileThree from "@/assets/images/jobsthankyou/mobile3.webp";
export default {
  data() {
    return {
      width: 0,
      interviews: [
        {
          image: MobileOne,
          content: `Registration and <br>Aptitude test`,
          class: "-translate-y-[260%]",
          desktopClass:
            "-translate-y-[170%] xl:-translate-y-[153%] -translate-x-[195%] lg:-translate-x-[191%] xl:-translate-x-[211%]",
          desktopImage: DesktopOne,
        },
        {
          image: MobileTwo,
          content: `Technical or Non-technical Interview`,
          class: "-translate-y-[185%]",
          desktopClass:
            "-translate-y-[168%] xl:-translate-y-[153%] -translate-x-[47%]",
          desktopImage: Desktoptwo,
        },
        {
          image: MobileThree,
          content: `If selected, <br>You will be called for an HR Interview`,
          class: "!w-40 -translate-y-[142%]",
          desktopClass:
            "-translate-y-[139%] xl:-translate-y-[127%] translate-x-[107%]  lg:translate-x-[102%] xl:translate-x-[123%]",
          desktopImage: Desktopthree,
        },
      ],
    };
  },
};
</script>
