<template>
  <div class="sm:tw-my-12 md:tw-my-36">
    <div class="tw-container">
      <div class="v2-header-3-text tw-w-1/2">
        Every <span class="border-text">Element</span> made with love
      </div>
    </div>

    <div class="tw-container tw-mt-20">
      <div
        class="sm:tw-grid-flow-row sm:tw-gap-x-8 md:tw-grid md:tw-grid-flow-col md:tw-gap-x-20"
      >
        <div class="tw-row-span-3 ...">
          <img :src="helloSummerImage" />
          <p class="v2-normal-text xl:tw-p-24 md:tw-p-8 sm:tw-p-14">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry.
          </p>
          <img :src="mobileCoverImage" />
        </div>
        <div class="md:tw-mt-80 sm:tw-mt-20">
          <img :src="helloSummerImage" />
          <img :src="keepGoingImage" class="tw-mt-20" />
          <p class="v2-normal-text xl:tw-p-28 md:tw-p-8 sm:tw-p-14">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry.
          </p>
        </div>
      </div>
    </div>
    <div class="tw-my-28">
      <img :src="elementCoverImage" />
    </div>
    <div class="sm:tw-my-24 md:tw-my-48 v2-header-3-text tw-text-center">
      Next Project
    </div>
  </div>
</template>
<script>
import elementCoverImage from "@/assets/images/Luxeradio-portfolio/element-cover.webp";
import helloSummerImage from "@/assets/images/Luxeradio-portfolio/hello-summer.webp";
import keepGoingImage from "@/assets/images/Luxeradio-portfolio/keepgoing.webp";
import mobileCoverImage from "@/assets/images/Luxeradio-portfolio/mobile-cover-image.webp";

export default {
  data() {
    return {
      elementCoverImage: elementCoverImage,
      helloSummerImage: helloSummerImage,
      keepGoingImage: keepGoingImage,
      mobileCoverImage: mobileCoverImage,
    };
  },
};
</script>
