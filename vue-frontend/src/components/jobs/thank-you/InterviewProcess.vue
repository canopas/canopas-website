<template>
  <section class="tw-container tw-mb-28">
    <div class="tw-pt-20 tw-pb-8 xl:tw-pb-20 tw-text-center">
      <span
        class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.4375rem] md:tw-leading-[5.15625rem]"
        >Interview Process
      </span>
    </div>
    <div
      v-for="interview in interviews"
      :key="interview"
      class="tw-block md:tw-hidden tw--mb-7"
    >
      <img
        alt="numeric-block-image"
        :src="interview.image"
        class="tw-relative tw-mx-auto tw-w-[226px] tw-h-[260px] tw-object-contain"
      />

      <span
        :class="interview.class"
        v-html="interview.content"
        class="tw-absolute tw-w-44 tw-text-center tw-text-black-core/[0.87] tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem] tw-left-1/2 tw-transform tw--translate-x-1/2"
      ></span>
    </div>
    <div
      class="tw-hidden md:tw-flex md:tw-flex-wrap md:tw-items-center md:tw-justify-center md:tw-mb-28"
    >
      <div v-for="interview in interviews" :key="interview">
        <img
          alt="numeric-block-image"
          :src="interview.desktopImage"
          class="tw-relative tw-mx-auto tw-w-[232px] lg:tw-w-[225px] xl:tw-w-[320px] tw-h-[200px] xl:tw-h-[237px] tw-object-contain"
        />
        <span
          v-html="interview.content"
          :class="interview.desktopClass"
          class="tw-absolute tw-w-[10rem] xl:tw-w-[12.5rem] tw-text-center tw-text-black-core/[0.87] tw-font-inter-semibold tw-text-[1.3rem] xl:tw-text-[1.625rem] tw-leading-[1.75rem] xl:tw-leading-[2.438rem] tw-left-1/2 tw-transform"
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
          class: "tw--translate-y-[260%]",
          desktopClass:
            "tw--translate-y-[170%] xl:tw--translate-y-[153%] tw--translate-x-[195%] lg:tw--translate-x-[191%] xl:tw--translate-x-[211%]",
          desktopImage: DesktopOne,
        },
        {
          image: MobileTwo,
          content: `Technical or Non-technical Interview`,
          class: "tw--translate-y-[185%]",
          desktopClass:
            "tw--translate-y-[168%] xl:tw--translate-y-[153%] tw--translate-x-[47%]",
          desktopImage: Desktoptwo,
        },
        {
          image: MobileThree,
          content: `If selected, <br>You will be called for an HR Interview`,
          class: "!tw-w-40 tw--translate-y-[142%]",
          desktopClass:
            "tw--translate-y-[139%] xl:tw--translate-y-[127%] tw-translate-x-[107%]  lg:tw-translate-x-[102%] xl:tw-translate-x-[123%]",
          desktopImage: Desktopthree,
        },
      ],
    };
  },
};
</script>
