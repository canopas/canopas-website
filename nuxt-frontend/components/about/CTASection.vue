<template>
  <section>
    <div
      class="container flex flex-col lg:flex-row-reverse items-center mt-[100px] 2xl:gap-[5.5rem]"
    >
      <div
        class="flex flex-col items-center lg:items-start lg:mt-[14.5rem] xl:mt-[15.5rem] 2xl:mt-[15.5rem] lg:w-[52%] xl:w-2/4"
      >
        <div
          class="md:w-[70%] lg:w-full text-center lg:text-left mobile-header-2 lg:desk-header-2 text-black-87"
        >
          Ready to take the next step and work together?
        </div>
        <div
          class="md:w-[70%] lg:w-full pt-4 lg:text-left text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
        >
          We'd love to ask some questions, learn a little more about what you do
          and who you serve and kick around some ideas for you.
        </div>
        <div class="pt-4 pb-60 hidden lg:flex">
          <nuxt-link
            class="gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_landing_cta')"
          >
            <span class="sub-h3-semibold lg:sub-h1-semibold"
              >Talk to our experts
            </span>
          </nuxt-link>
        </div>
      </div>
      <div class="md:w-[80%] lg:w-[48%] xl:w-w-2/4 lg:pr-10 2xl:pr-0">
        <img
          :src="CTAImage[0]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="aboutus-cta-image"
          class="w-full h-full object-contain"
          loading="lazy"
        />
      </div>
      <div class="pb-16 mt-[-43px] lg:hidden">
        <nuxt-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Talk to our experts
          </span>
        </nuxt-link>
      </div>
    </div>
  </section>
</template>
<script type="module">
import cta400 from "@/assets/images/about/cta-400w.webp";
import cta800 from "@/assets/images/about/cta-800w.webp";
import cta1600 from "@/assets/images/about/cta-1600w.webp";
import cta2400 from "@/assets/images/about/cta-2400w.webp";
export default {
  data() {
    return {
      CTAImage: [cta400, cta800, cta1600, cta2400],
    };
  },
  inject: ["mixpanel"],
};
</script>
