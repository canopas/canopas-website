<template>
  <section>
    <div class="tw-container">
      <div class="tw-flex tw-flex-col tw-text-center">
        <span
          class="tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.4375rem] md:tw-text-[3.4375rem] md:tw-leading-[5.15625rem] tw-text-black-core/[0.87]"
          >Case Studies</span
        >
        <span
          class="tw-py-4 tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-mx-auto md:tw-text-[1.5rem] md:tw-leading-[2.25rem] tw-text-black-core/[0.60] lg:tw-w-[80%] xl:tw-w-[61%]"
          >Explore our case studies to witness how we've transformed our
          client's ideas into successful Android apps.</span
        >
      </div>
      <!-- ---------------------Mobile UI------------------ -->
      <div
        v-for="item in cases"
        :key="item"
        class="tw-mt-7 tw-relative md:tw-hidden"
        :class="item.className"
      >
        <img
          :src="item.images[0]"
          :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
          alt="casestudy-image"
          class="tw-w-full tw-h-full tw-object-cover tw-mt-7"
          loading="lazy"
        />
        <div
          :class="item.textalign"
          class="tw-absolute tw-bottom-[10%] tw-flex tw-flex-col tw-text-black-core/[0.87]"
        >
          <span
            class="tw-font-inter-semibold tw-text-[1.375rem] tw-leading-[1.7875rem]"
            >{{ item.title }}</span
          >
          <span
            class="tw-font-inter-regular tw-text-[1rem] tw-leading-[1.4rem]"
            v-html="item.content"
          ></span>
        </div>
      </div>
      <div
        v-for="item in cta"
        :key="item"
        class="tw-mt-7 tw-relative md:tw-hidden"
        :class="item.className"
      >
        <img
          :src="item.images[0]"
          :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
          alt="casestudy-image"
          class="tw-w-full tw-h-full tw-object-cover tw-mt-7"
          loading="lazy"
        />
        <div
          :class="item.textalign"
          class="tw-absolute tw-bottom-[10%] tw-flex tw-flex-col tw-text-black-core/[0.87]"
        >
          <span
            class="tw-font-inter-semibold tw-text-[1.375rem] tw-leading-[1.7875rem]"
            >{{ item.title }}</span
          >
          <span
            class="tw-font-inter-regular tw-text-[1rem] tw-leading-[1.4rem]"
            v-html="item.content"
          ></span>
          <div class="tw-mt-[1rem] tw-flex tw-flex-wrap tw-justify-center">
            <router-link
              class="white-btn gradient-border-btn tw-flex tw-h-[2.8125rem] tw-w-max tw-items-center tw-rounded-full tw-border-none"
              to="/contact"
              @click.native="mixpanel.track('tap_android_development_cta4')"
            >
              <span
                class="tw-font-inter-semibold tw-text-[1rem] tw-font-normal tw-leading-[1.21rem] !tw-tracking-[0] md:tw-text-[1.09375rem] md:tw-text-[1.1875rem] md:tw-leading-[1.3125rem] md:tw-leading-[1.4375rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
                >Inquire Now</span
              >
            </router-link>
          </div>
        </div>
      </div>
    </div>
    <!-- ------------------------ MobileUI END---------------------- -->
    <div class="tw-hidden md:tw-block tw-bg-[#F6F6F6] tw-mt-14">
      <div class="swiper-content tw-relative tw--mt-6 xll:tw-container">
        <swiper
          :slidesPerView="1"
          :centeredSlides="true"
          :loop="true"
          :speed="4000"
          :autoplay="{
            delay: 4000,
            disableOnInteraction: false,
          }"
          :loopedSlides="50"
          :spaceBetween="20"
          :navigation="true"
          class="swiper-container tw-absolute tw-left-0 tw-top-[10px] tw-z-30 tw-mx-auto tw-h-full tw-w-full"
        >
          <swiper-slide v-for="item in cases" :key="item">
            <div
              class="tw-flex tw-flex-row xl:tw-px-[4rem] tw-py-[6rem] tw-w-[91%] 2xl:tw-w-[77%] 2xll:tw-w-[72%] xll:tw-w-[91%] tw-mx-auto"
            >
              <div
                class="tw-flex tw-flex-col tw-w-[77%] xl:tw-w-[48%] 2xll:tw-w-[50%] lg:tw-mr-auto xll:tw-ml-auto"
              >
                <div class="tw-mb-4">
                  <span
                    class="v2-canopas-gradient-text tw-text-left tw-font-inter-medium tw-text-[1.5rem] tw-leading-[1.816rem]"
                  >
                    {{ item.title }}
                  </span>
                </div>
                <span
                  class="tw-mb-[1rem] tw-font-inter-semibold tw-text-[2rem] tw-leading-[2.7375rem] lg:tw-text-[2.875rem] lg:tw-leading-[3.7375rem]"
                  v-html="item.deskcontent"
                ></span>
                <span
                  class="tw-text-black-core/[0.60] tw-font-inter-medium tw-text-[1rem] tw-leading-[1.575rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.875rem]"
                  v-html="item.description"
                ></span>
                <img
                  :src="resultsimage"
                  class="tw-h-[3rem] tw-w-[3rem] lg:tw-h-[5rem] lg:tw-w-[5rem] tw-object-cover tw-ml-[60%] tw-mt-[2rem] tw-mb-[1rem] 2xll:tw-ml-[50%]"
                  alt="case_study_result_image"
                />
                <div
                  class="tw-bg-white tw-drop-shadow-md tw-rounded md:tw-w-[390px] lg:tw-w-[457px]"
                >
                  <div class="tw-inline-flex tw-p-[0.5rem] lg:tw-p-4">
                    <div class="tw-p-[0.8rem]">
                      <span
                        class="tw-font-inter-semibold tw-text-[1rem] lg:tw-text-[2.125rem] tw-leading-[1.7625rem] lg:tw-leading-[2.7625rem] v2-canopas-gradient-text"
                        >{{ item.resultcount[0] }}</span
                      >
                      <div
                        class="tw-flex-1 tw-font-inter-regular tw-text-[1rem] tw-leading-[1rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.625rem]"
                      >
                        <span>Downloads</span>
                      </div>
                    </div>
                    <div
                      class="tw-border-l-[0.5px] tw-border-black-core/[0.6] tw-p-[0.8rem]"
                    >
                      <span
                        class="tw-font-inter-semibold tw-text-[1rem] lg:tw-text-[2.125rem] tw-leading-[1.7625rem] lg:tw-leading-[2.7625rem] v2-canopas-gradient-text"
                        >{{ item.resultcount[1] }}</span
                      >
                      <div
                        class="tw-flex-1 tw-font-inter-regular tw-text-[1rem] tw-leading-[1rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.625rem]"
                      >
                        <span>App Rating</span>
                      </div>
                    </div>
                    <div
                      class="tw-border-l-[0.5px] tw-border-black-core/[0.6] tw-p-[0.8rem]"
                    >
                      <span
                        class="tw-font-inter-semibold lg:tw-text-[2.125rem] tw-leading-[1.7625rem] lg:tw-leading-[2.7625rem] v2-canopas-gradient-text"
                        >{{ item.resultcount[2] }}</span
                      >
                      <div
                        class="tw-flex-1 tw-font-inter-regular tw-text-[1rem] tw-leading-[1rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.625rem]"
                      >
                        <span>Monthly Users</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div>
                <img
                  :src="item.images[2]"
                  :srcset="`${item.images[2]} 400w,${item.images[3]} 800w`"
                  class="tw-mt-[3rem] tw-h-[20rem] lg:tw-h-[24.75006rem] xl:tw-h-[29.75006rem] tw-w-[25.87525rem] lg:tw-w-[30.87525rem] xl:tw-w-[29.87525rem] tw-object-fit"
                  alt="CaseStudyImage"
                />
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </section>
</template>

<script type="module">
import justlymobile_400w from "@/assets/images/andriod-app-development/casestudy/p1-400w.webp";
import justlymobile_800w from "@/assets/images/andriod-app-development/casestudy/p1-800w.webp";
import luxeradiomobile_400w from "@/assets/images/andriod-app-development/casestudy/p2-400w.webp";
import luxeradiomobile_800w from "@/assets/images/andriod-app-development/casestudy/p2-800w.webp";
import tognessmobile_400w from "@/assets/images/andriod-app-development/casestudy/p3-400w.webp";
import tognessmobile_800w from "@/assets/images/andriod-app-development/casestudy/p3-800w.webp";
import smileplusmobile_400w from "@/assets/images/andriod-app-development/casestudy/p4-400w.webp";
import smileplusmobile_800w from "@/assets/images/andriod-app-development/casestudy/p4-800w.webp";
import cta_400w from "@/assets/images/andriod-app-development/casestudy/cta-400w.webp";
import cta_800w from "@/assets/images/andriod-app-development/casestudy/cta-800w.webp";
import resultsimage from "@/assets/images/andriod-app-development/casestudy/results.webp";
import justly_400w from "@/assets/images/andriod-app-development/casestudy/desktop/p1-400w.webp";
import justly_800w from "@/assets/images/andriod-app-development/casestudy/desktop/p1-800w.webp";
import luxeradio_400w from "@/assets/images/andriod-app-development/casestudy/desktop/p2-400w.webp";
import luxeradio_800w from "@/assets/images/andriod-app-development/casestudy/desktop/p2-800w.webp";
import togness_400w from "@/assets/images/andriod-app-development/casestudy/desktop/p3-400w.webp";
import togness_800w from "@/assets/images/andriod-app-development/casestudy/desktop/p3-800w.webp";
import smileplus_400w from "@/assets/images/andriod-app-development/casestudy/desktop/p4-400w.webp";
import smileplus_800w from "@/assets/images/andriod-app-development/casestudy/desktop/p4-800w.webp";
import SwiperCore, { Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
SwiperCore.use([Autoplay]);
export default {
  data() {
    return {
      resultsimage,
      cases: [
        {
          title: "Justly",
          content: "Create systems to build your dream life",
          deskcontent: `Overcome depression <br>by building habits`,
          resultcount: ["5k+", "4.9", "100"],
          description: `Justly aims to tackle loneliness, depression, and <br>mental health through innovative solutions.`,
          images: [
            justlymobile_400w,
            justlymobile_800w,
            justly_400w,
            justly_800w,
          ],
          className: "tw-text-left tw-items-start",
          textalign: "tw-w-[62%] tw-ml-[1rem]",
        },
        {
          title: "Luxeradio",
          content: "Radio, music and podcast",
          deskcontent: `A radio, music and <br> podcast platform`,
          resultcount: ["100k+", "5", "15k+"],
          description: `Luxe Radio displays the best of Moroccan and <br>international creation, emphasizing taste, elegance, and refinement.`,
          images: [
            luxeradiomobile_400w,
            luxeradiomobile_800w,
            luxeradio_400w,
            luxeradio_800w,
          ],
          className: "tw-text-left tw-items-start",
          textalign: "tw-right-[10%] tw-ml-[1rem] ",
        },
        {
          title: "Togness",
          content: "Create your life’s most memorable events",
          deskcontent: `Photo editor and <br> video maker app`,
          resultcount: ["10k+", "4.7 ", "1200+"],
          description: `Togness are create slideshows, <br>and add quotes and your own music for every special occasion of your life!`,
          images: [
            tognessmobile_400w,
            tognessmobile_800w,
            togness_400w,
            togness_800w,
          ],
          className: "tw-text-left tw-items-start ",
          textalign: "tw-w-[65%] tw-ml-[1rem] ",
        },
        {
          title: "Smile+",
          content: "Create systems to build your dream life",
          deskcontent: `For Dentists and <br>Dental labs`,
          resultcount: ["5k+", "4.9", "350+"],
          description: `Smile+ allows to create the perfect smile of your <br>patient in a few minutes.`,
          images: [
            smileplusmobile_400w,
            smileplusmobile_800w,
            smileplus_400w,
            smileplus_800w,
          ],
          className: "tw-text-left tw-items-start ",
          textalign: "tw-right-[5%] tw-w-[52%] tw-ml-[1rem]",
        },
      ],
      cta: [
        {
          title: "",
          content: `"Let’s Start a <span class="tw-text-black-core/[0.87] tw-font-inter-semibold tw-text-[1.375rem] tw-leading-[2.0625rem]">New Project</span> Together"`,
          images: [cta_400w, cta_800w],
          className: "tw-text-center ",
          textalign:
            "tw-top-[65%] sm:tw-top-[68%]  tw-w-full tw-left-1/2 tw-transform tw--translate-x-1/2 tw--translate-y-1/2 ",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
};
</script>
