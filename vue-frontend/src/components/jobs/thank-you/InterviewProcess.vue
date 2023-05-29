<template>
  <section class="tw-container tw-mb-28">
    <div class="tw-pt-20 tw-pb-8 xl:tw-pb-20 tw-text-center">
      <span
        class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.4375rem] md:tw-leading-[5.15625rem]"
        >Interview Process
      </span>
    </div>
    <div
      v-for="interview in interviews.splice(3, 5)"
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
        class="tw-absolute tw-w-44 tw-text-center tw-text-black-core/[0.87] tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem] tw-left-1/2 tw-transform tw--translate-x-1/2"
        >{{ interview.content }}</span
      >
    </div>
    <div
      class="tw-hidden md:tw-flex md:tw-flex-wrap md:tw-items-center md:tw-justify-center md:tw-mb-28"
    >
      <div v-for="interview in interviews.splice(0, 3)" :key="interview">
        <img
          alt="numeric-block-image"
          :src="interview.image"
          class="tw-relative tw-mx-auto tw-w-[232px] lg:tw-w-[225px] xl:tw-w-[320px] tw-h-[200px] xl:tw-h-[237px] tw-object-contain"
        />
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
          image: DesktopOne,
        },
        {
          image: Desktoptwo,
        },
        {
          image: Desktopthree,
        },

        {
          image: MobileOne,
          content: "Registration and Aptitude test",
          class: "tw--translate-y-[260%]",
        },
        {
          image: MobileTwo,
          content: "Technical or Non-technical Interview",
          class: "tw--translate-y-[185%]",
        },
        {
          image: MobileThree,
          content: "If selected,You will be called for an HR Interview",
          class: "tw-w-36 tw--translate-y-[180%]",
        },
      ],
    };
  },
};
</script>
