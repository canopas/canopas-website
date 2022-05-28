<template>
  <div>
    <div class="tw-container">
      <div class="sm:tw-w-4/6 xl:tw-w-3/6 sm:tw-mt-14 md:tw-mt-44 tw-mb-32">
        <div class="v2-header-3-text">
          The <span class="border-text">design</span> will not leave you cold
        </div>
        <div>
          <p class="v2-normal-text tw-my-5 tw-w-11/12">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum
            is simply dummy text of the printing and typesetting industry.
          </p>
        </div>
      </div>
    </div>

    <div class="tw-my-28">
      <img :src="luxeRadioCoverImage" />
    </div>
    <div class="tw-container sm:tw-my-20 md:tw-my-56">
      <div class="tw-grid tw-grid-flow-col">
        <div>
          <p class="v2-title-2-text tw-row-span-3 ... tw-w-3/5">
            Extra attention to detail development
          </p>
        </div>

        <div>
          <p class="v2-title-3-text">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum
            is simply dummy text of the printing and typesetting industry.
          </p>
        </div>
      </div>
    </div>
    <div>
      <img :src="discogsImage" />
    </div>
  </div>
</template>
<script>
import luxeRadioCoverImage from "@/assets/images/Luxeradio-portfolio/luxeradiopage.webp";
import discogsImage from "@/assets/images/Luxeradio-portfolio/Discogs-image.webp";
export default {
  data() {
    return {
      luxeRadioCoverImage: luxeRadioCoverImage,
      discogsImage: discogsImage,
    };
  },
};
</script>
