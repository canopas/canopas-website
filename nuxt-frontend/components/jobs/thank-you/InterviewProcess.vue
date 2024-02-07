<template>
  <section class="container my-16 lg:my-60">
    <div class="mobile-header-2 text-black-87 lg:desk-header-2 text-center">
      Interview Process
    </div>
    <div
      v-for="interview in interviews"
      :key="interview"
      class="block lg:hidden"
    >
      <img
        alt="numeric-block-image"
        :src="interview.image"
        class="relative mx-auto w-[14.125rem] h-[16.25rem] object-contain"
      />

      <span
        :class="interview.class"
        v-html="interview.content"
        class="absolute w-44 text-center mobile-header-3-semibold text-black-87 left-1/2 transform -translate-x-1/2"
      ></span>
    </div>
    <div class="hidden lg:flex flex-wrap items-center justify-center mt-16">
      <div v-for="interview in interviews" :key="interview">
        <img
          alt="numeric-block-image"
          :src="interview.desktopImage"
          class="relative mx-auto w-[16.25rem] xl:w-80 h-[12.5rem] xl:h-[14.813rem] object-contain"
        />
        <span
          v-html="interview.content"
          :class="interview.desktopClass"
          class="absolute w-40 xl:w-[12.5rem] text-center mobile-header-2-semibold text-black-87 left-1/2 transform"
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
          content: `Registration<br> and <br>Aptitude test`,
          class: "-translate-y-[192%]",
          desktopClass:
            "-translate-y-[147%] xl:-translate-y-[160%] -translate-x-[213%] xl:-translate-x-[211%]",
          desktopImage: DesktopOne,
        },
        {
          image: MobileTwo,
          content: `Technical or <br>Non-technical Interview`,
          class: "-translate-y-[192%]",
          desktopClass:
            "-translate-y-[119%] xl:-translate-y-[158%] -translate-x-[47%]",
          desktopImage: Desktoptwo,
        },
        {
          image: MobileThree,
          content: `If selected, <br>You will be called for an HR Interview`,
          class: "!w-[8.5rem] -translate-y-[142%]",
          desktopClass:
            "-translate-y-[121%] xl:-translate-y-[131%] translate-x-[125%] xl:translate-x-[123%]",
          desktopImage: Desktopthree,
        },
      ],
    };
  },
};
</script>
