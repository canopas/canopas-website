<template>
  <section>
    <div
      class="tw-relative tw-w-full tw-mx-auto tw-py-0 tw-px-[0.75rem] lg:tw-py-[50px]"
    >
      <img
        :src="background"
        alt="landing-view-image"
        class="tw-w-full tw-h-[115%] lg:tw-h-full tw-absolute tw-left-0 -tw-z-[5]"
      />
      <div class="tw-relative tw-z-0 tw-pt-8 lg:tw-mt-16 lg:tw-pb-24">
        <div
          class="tw-container tw-flex tw-flex-col tw-items-center lg:tw-flex-row-reverse lg:tw-items-start"
        >
          <div class="tw-flex-1 lg:tw-block lg:tw-mt-5">
            <aspect-ratio height="150%">
              <img
                :src="landingImages[2]"
                :srcset="`${landingImages[0]} 400w, ${landingImages[1]} 800w, ${landingImages[2]} 1200w`"
                alt="landing-view-image"
                class="tw-w-full tw-h-full tw-object-cover"
              />
            </aspect-ratio>
          </div>

          <div class="tw-flex-2">
            <div class="description">
              <div
                class="tw-pb-5 lg:tw-w-11/12 v2-header-3-text lg:tw-text-5xl-1 lg:tw-text-7xl-1 lg:tw-leading-18 lg:tw-text-[3.0625rem] lg:tw-leading-[4rem] xl:tw-text-[4rem] xl:tw-leading-[4.7rem] lg:tw-pb-8 xl:tw-pb-10"
              >
                For Over A Decade We Have Helped Businesses Grow Exponentially.
                <div>
                  — <span class="canopas-gradient-text"> Is Yours Next?</span>
                </div>
              </div>
              <div
                class="v2-normal-text sm:tw-w-[85%] xl:tw-w-3/4 tw-pb-8 lg:tw-pb-8 xl:tw-text-[1.68rem] xl:tw-leading-[2.4rem] xl:tw-pb-16"
              >
                We develop Mobile and Web apps that help you
                <span
                  class="tw-bg-no-repeat tw-inline tw-bg-top-bottom-150 md:tw-bg-top-bottom-140 lg:tw-bg-top-bottom-130 tw-bg-[length:100%_10px] tw-bg-gradient-underline"
                  >generate more revenue</span
                >,
                <span
                  class="tw-bg-no-repeat tw-inline tw-bg-top-bottom-150 md:tw-bg-top-bottom-140 lg:tw-bg-top-bottom-130 tw-bg-[length:100%_10px] tw-bg-gradient-underline"
                  >acquire new users</span
                >, and
                <span
                  class="tw-bg-no-repeat tw-inline tw-bg-top-bottom-150 md:tw-bg-top-bottom-140 lg:tw-bg-top-bottom-130 tw-bg-[length:100%_10px] tw-bg-gradient-underline"
                  >deliver an exceptional user experience</span
                >.
              </div>
              <div
                class="tw-flex tw-items-center tw-flex-col-reverse md:tw-flex-row"
              >
                <div class="tw-text-center">
                  <router-link
                    class="tw-rounded-[3rem] tw-m-0 tw-py-[0.7rem] tw-px-5 tw-text-center gradient-btn v2-normal-2-text tw-w-max tw-flex tw-items-center"
                    :to="contactURL"
                    @click.native="$gtag.event('tap_landing_cta')"
                  >
                    <span class="tw-mr-2.5"
                      ><span class="tw-font-medium"
                        >Get Free Consultation
                      </span></span
                    > </router-link
                  ><small class="md:tw-hidden">100% MONEY BACK GUARANTEE</small>
                </div>

                <div
                  class="tw-flex tw-items-center tw-pb-3.5 md:tw-pb-0 md:tw-pl-7"
                >
                  <div v-for="review in reviews" :key="review">
                    <div
                      v-if="review.image"
                      class="tw-w-[40px] tw-h-[40px] md:tw-w-[50px] md:tw-h-[50px] tw-rounded-full"
                      :class="review.className"
                    >
                      <img
                        :src="review.image"
                        class="tw-w-full tw-h-full tw-object-cover"
                        loading="lazy"
                        alt="phase.alt"
                      />
                    </div>
                    <div v-if="review.plusIcon">
                      <div
                        class="tw-w-[40px] tw-h-[40px] md:tw-w-[50px] md:tw-h-[50px] tw-border-[3px] tw-border-white tw-rounded-full tw-bg-[#ffc4c0]"
                        :class="review.className"
                      >
                        <font-awesome-icon
                          class="fa tw-text-white tw-w-[35px] tw-h-[23px] tw-pt-[5px] md:tw-w-[45px] md:tw-pt-[10px]"
                          icon="plus"
                          id="leftArrow"
                        />
                      </div>
                    </div>
                  </div>
                  <div class="tw-pl-4">
                    <font-awesome-icon
                      v-for="i in 5"
                      :key="i"
                      class="fas tw-text-[#FF3D2E]"
                      icon="star"
                    />
                    <div class="tw-font-light">40+ Reviews</div>
                  </div>
                </div>
              </div>
              <small
                class="tw-hidden md:tw-flex v2-normal-3-text tw-pt-5 tw-w-11/12 md:tw-w-2/3 lg:tw-w-9/12 xl:tw-w-3/5 lg:tw-text-[1.0625rem] lg:tw-leading-[1.5rem] lg:tw-pt-8 xl:tw-pt-8"
              >
                We offer a 100% MONEY BACK GUARANTEE FOR 14 DAYS if you don't
                like our work. No questions asked.
              </small>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import landing400w from "@/assets/images/Landing/landing-400w.webp";
import landing800w from "@/assets/images/Landing/landing-800w.webp";
import landing1200w from "@/assets/images/Landing/landing-1200w.webp";
import review1 from "@/assets/images/Landing/review1.webp";
import review2 from "@/assets/images/Landing/review2.webp";
import review3 from "@/assets/images/Landing/review3.webp";
import review4 from "@/assets/images/Landing/review4.webp";
import backgroundImage from "@/assets/images/Landing/background-1200w.webp";

import AspectRatio from "@/components/utils/AspectRatio.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
export default {
  data() {
    return {
      landingImages: [landing400w, landing800w, landing1200w],
      background: backgroundImage,
      contactURL: "/contact",
      reviews: [
        {
          id: 1,
          className: "tw-z-0",
          image: review1,
        },
        {
          id: 2,
          className: "tw-z-10 -tw-ml-[20px]",
          image: review2,
        },
        {
          id: 3,
          className: "tw-z-20 -tw-ml-[20px]",
          image: review3,
        },
        {
          id: 4,
          className: "tw-z-30 -tw-ml-[20px]",
          image: review4,
        },
        {
          id: 5,
          className: "tw-z-40 -tw-ml-[20px]",
          plusIcon: true,
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
    AspectRatio,
  },
};
</script>
