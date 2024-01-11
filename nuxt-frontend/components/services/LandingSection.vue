<template>
  <section
    class="relative mt-8 lg:mt-[4.5rem] bg-gradient-to-t from-orange-1 via-pink-1 to-white pb-14 3xl:outer-container"
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
    <div class="container flex flex-col items-center lg:flex-row-reverse">
      <aspect-ratio class="w-[45%] 2xl:w-1/2 lg:block hidden" height="60%">
        <img
          class="absolute lg:top-[-2.188rem] right-0 xll:right-auto h-full object-cover"
          :src="landingImages[0]"
          :srcset="`${landingImages[0]} 400w, ${landingImages[1]} 800w`"
          alt="landing-view-image"
        />
      </aspect-ratio>
      <div
        class="container w-full lg:w-[55%] relative flex flex-col 2xl:w-1/2 items-center"
      >
        <h1
          class="mobile-header-1 lg:desk-header-2 xl:desk-header-1 text-black-87"
        >
          We Build Products that Customers can't Help but Love.
        </h1>
        <div
          class="mt-6 w-full sub-h1-regular lg:mobile-header-3-regular xl:mobile-header-2-regular text-black-60"
        >
          <p>
            At Canopas, we combine expert engineering with a touch of magic.
            What do you get in results? Exquisite, high-quality apps our clients
            (and their customers) love.
          </p>
          <p class="mt-5">Experience 100% peace of mind with us.</p>
          <p class="mt-5">
            No more wondering what we're doing or when your app will be done.
            With us, you can expect weekly or bi-weekly check-ins, progress
            updates, and plenty of hand-holding.
          </p>
        </div>
        <nuxt-link
          class="mt-8 lg:ml-0 gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
    </div>
  </section>
</template>

<script>
import landing400w from "@/assets/images/services/landing/landing-400w.webp";
import landing800w from "@/assets/images/services/landing/landing-800w.webp";
import bg400w from "@/assets/images/services/bg/bg-400w.webp";
import bg800w from "@/assets/images/services/bg/bg-800w.webp";
import bg1200w from "@/assets/images/services/bg/bg-1200w.webp";
import bg1600w from "@/assets/images/services/bg/bg-1600w.webp";
import AspectRatio from "@/components/utils/AspectRatio.vue";

export default {
  data() {
    return {
      width: 680,
      landingImages: [landing400w, landing800w],
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
