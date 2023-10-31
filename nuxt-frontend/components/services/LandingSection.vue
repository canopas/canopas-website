<template>
  <section
    class="relative mt-[15%] lg:mt-0 bg-gradient-to-t from-[#F47695]/[0.1] via-[#F98584]/[0.1] to-[#FFFFFF] pb-[6.5rem]"
  >
    <div>
      <img
        v-if="width < 680"
        :src="bgImages[0]"
        :srcset="`${bgImages[0]} 400w, ${bgImages[1]} 800w, ${bgImages[2]} 1200w, ${bgImages[2]} 1600w`"
        class="lg:hidden absolute top-0 left-0 w-full h-full -z-[1] object-cover"
        alt="background-image"
        loading="eager"
      />
    </div>
    <div
      class="container flex flex-col items-center lg:flex-row-reverse mt-[30px] pb-[25px]"
    >
      <aspect-ratio class="flex-[45%] 2xl:w-[50%] lg:block hidden" height="60%">
        <img
          class="absolute lg:top-[-35px] right-0 xll:right-auto h-full object-cover"
          :src="landingImages[0]"
          :srcset="`${landingImages[0]} 400w, ${landingImages[1]} 800w, ${landingImages[2]} 1200w, ${landingImages[2]} 1400w`"
          alt="landing-view-image"
        />
      </aspect-ratio>
      <div class="relative flex flex-col flex-[55%] 2xl:w-[50%] items-center">
        <h1
          class="text-[2.5rem] md:text-[3.125rem] xl:text-[3.75rem] leading-[3.25rem] md:leading-[4.5rem] xl:leading-[5.625rem] text-center lg:text-left font-inter-bold text-black-core/[0.87]"
        >
          We Build Products that Customers can't Help but Love.
        </h1>
        <div
          class="mt-5 w-full sm:w-[85%] md:w-[76%] lg:w-full text-left text-[1rem] md:text-[1.094rem] xl:text-[1.188rem] leading-6 md:leading-[1.625rem] xl:leading-[1.781rem] font-inter-medium text-black-core/[0.87]"
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
