<template>
  <div>
    <div class="tw-container tw-mx-auto">
      <div class="tw-mt-28 tw-mb-36">
        <p class="v2-normal-2-text">Luxeradio</p>
        <p class="v2-title-text tw-mt-10 xl:tw-pr-72">
          Lorem Ipsum is simply dummy text of the printing.
        </p>
      </div>
    </div>
    <div>
      <img :src="landingViewImage" alt="landing-view-image" />
    </div>
  </div>
</template>
<script>
import landingViewImage from "@/assets/images/Luxeradio-portfolio/landingview-image.webp";
export default {
  data() {
    return {
      landingViewImage: landingViewImage,
    };
  },
};
</script>
