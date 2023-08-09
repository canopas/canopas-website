<template>
  <section
    class="tw-relative tw-mt-[15%] lg:tw-mt-0 tw-bg-gradient-to-t tw-from-[#F47695]/[0.1] tw-via-[#F98584]/[0.1] tw-to-[#FFFFFF] tw-pb-[6.5rem]"
  >
    <div>
      <img
        v-if="width < 680"
        :src="bgImages[0]"
        :srcset="`${bgImages[0]} 400w, ${bgImages[1]} 800w, ${bgImages[2]} 1200w, ${bgImages[2]} 1600w`"
        class="lg:tw-hidden tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-full tw--z-[1] tw-object-cover"
        alt="background-image"
      />
    </div>
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center lg:tw-flex-row-reverse tw-mt-[30px] tw-pb-[25px]"
    >
      <aspect-ratio
        class="tw-flex-[45%] 2xl:tw-w-[50%] lg:tw-block tw-hidden"
        height="60%"
      >
        <img
          class="tw-absolute lg:tw-top-[-35px] tw-right-0 xll:tw-right-auto tw-h-full tw-object-cover"
          :src="landingImages[0]"
          :srcset="`${landingImages[0]} 400w, ${landingImages[1]} 800w, ${landingImages[2]} 1200w, ${landingImages[2]} 1400w`"
          alt="landing-view-image"
          loading="lazy"
        />
      </aspect-ratio>
      <div
        class="tw-relative tw-flex tw-flex-col tw-flex-[55%] 2xl:tw-w-[50%] tw-items-center"
      >
        <h1
          class="tw-text-[2.5rem] md:tw-text-[3.125rem] xl:tw-text-[3.75rem] tw-leading-[3.25rem] md:tw-leading-[4.5rem] xl:tw-leading-[5.625rem] tw-text-center lg:tw-text-left tw-font-inter-bold tw-text-black-core/[0.87]"
        >
          We Build Products that Customers can't Help but Love.
        </h1>
        <div
          class="tw-mt-5 tw-w-full sm:tw-w-[85%] md:tw-w-[76%] lg:tw-w-full tw-text-left tw-text-[1rem] md:tw-text-[1.094rem] xl:tw-text-[1.188rem] tw-leading-6 md:tw-leading-[1.625rem] xl:tw-leading-[1.781rem] tw-font-inter-medium tw-text-black-core/[0.87]"
        >
          <p>
            At Canopas, we combine expert engineering with a touch of magic.
            What do you get in results? Exquisite, high-quality apps our clients
            (and their customers) love.
          </p>
          <br />
          <p>Experience 100% peace of mind with us.</p>
          <br />
          <p>
            No more wondering what we're doing or when your app will be done.
            With us, you can expect weekly or bi-weekly check-ins, progress
            updates, and plenty of hand-holding.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import landing400w from "@/assets/images/services/landing/landing-400w.webp";
import landing800w from "@/assets/images/services/landing/landing-800w.webp";
import landing1200w from "@/assets/images/services/landing/landing-1200w.webp";
import landing1400w from "@/assets/images/services/landing/landing-1400w.webp";
import bg400w from "@/assets/images/services/bg/bg-400w.webp";
import bg800w from "@/assets/images/services/bg/bg-800w.webp";
import bg1200w from "@/assets/images/services/bg/bg-1200w.webp";
import bg1600w from "@/assets/images/services/bg/bg-1600w.webp";
import AspectRatio from "@/components/utils/AspectRatio.vue";

export default {
  data() {
    return {
      width: 680,
      landingImages: [landing400w, landing800w, landing1200w, landing1400w],
      bgImages: [bg400w, bg800w, bg1200w, bg1600w],
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
