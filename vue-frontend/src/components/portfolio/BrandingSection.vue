<template>
  <div>
    <div class="tw-container">
      <div class="v2-header-3-text tw-w-24">
        <span class="border-text">Careful</span> branding
      </div>
    </div>
    <div class="tw-my-28">
      <img :src="luxeRadioCoverImage" />
    </div>
    <div class="tw-container">
      <div
        class="sm:tw-grid-flow-row sm:tw-gap-x-8 md:tw-grid md:tw-grid-flow-col md:tw-gap-x-20"
      >
        <div class="tw-row-span-3 ...">
          <img :src="designMobile" />
          <p class="v2-normal-text xl:tw-p-24 md:tw-p-8 sm:tw-p-7">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry.
          </p>
          <img :src="equalizerImage" />
          <img :src="nightRideImage" class="tw-mt-14" />
        </div>
        <div class="md:tw-mt-80 sm:tw-mt-20">
          <img :src="essentialImage" />
          <p class="v2-normal-text xl:tw-p-28 md:tw-p-8 sm:tw-p-7">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry.
          </p>
          <img :src="genreImage" />
          <p class="v2-normal-text xl:tw-p-28 md:tw-p-8 sm:tw-p-7">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import luxeRadioCoverImage from "@/assets/images/Luxeradio-portfolio/luxeradiopage.webp";
import designMobile from "@/assets/images/Luxeradio-portfolio/mobile.webp";
import essentialImage from "@/assets/images/Luxeradio-portfolio/essential.webp";
import equalizerImage from "@/assets/images/Luxeradio-portfolio/equalizer.webp";
import nightRideImage from "@/assets/images/Luxeradio-portfolio/nightride.webp";
import genreImage from "@/assets/images/Luxeradio-portfolio/genre.webp";
export default {
  data() {
    return {
      luxeRadioCoverImage: luxeRadioCoverImage,
      designMobile: designMobile,
      essentialImage: essentialImage,
      equalizerImage: equalizerImage,
      genreImage: genreImage,
      nightRideImage: nightRideImage,
    };
  },
};
</script>
