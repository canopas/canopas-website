<template>
  <section
    class="tw-mt-16 md:tw-mt-40 lg:tw--mt-48 xl:tw--mt-[32rem] 2xl:tw-mt-[-30rem] 3xl:tw-outer-container"
  >
    <div class="tw-container">
      <div class="tw-flex tw-flex-col tw-text-center">
        <span class="header-2 tw-text-black-core-[0.87]">Case studies</span>
        <span
          class="tw-mt-4 md:tw-mt-6 sub-h1-regular tw-mx-auto xl:tw-text-[1.5rem] xl:tw-leading-9 tw-text-black-core/[0.60] lg:tw-w-[80%] xl:tw-w-[61%]"
          >Explore our case studies to witness how we've transformed our
          client's ideas into successful Android apps.</span
        >
      </div>
      <!-- ---------------------MOBILE UI START------------------ -->
      <div
        v-for="item in cases"
        :key="item"
        class="tw-mt-6 tw-relative lg:tw-hidden"
        :class="item.className"
      >
        <a
          class="tw-cursor-pointer"
          :href="item.url"
          @click.native="mixpanel.track('tap_android_app_portfolio')"
          :target="item.target"
        >
          <img
            :src="item.images[0]"
            :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
            alt="casestudy-image"
            class="tw-w-full tw-h-full tw-object-cover"
            loading="lazy"
          />

          <div
            :class="item.textalign"
            class="tw-absolute tw-bottom-4 tw-flex tw-flex-col"
          >
            <span class="header-2 tw-text-black-core/[0.87]">{{
              item.title
            }}</span>
            <span
              class="tw-mt-2 sub-h3-regular tw-text-black-core/[0.6]"
              v-html="item.content"
            ></span>
          </div>
        </a>
      </div>
      <div
        v-for="item in cta"
        :key="item"
        class="tw-mt-16 tw-relative lg:tw-hidden"
        :class="item.className"
      >
        <img
          :src="item.images[0]"
          :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
          alt="casestudy-image"
          class="tw-w-full tw-h-full tw-object-cover"
          loading="lazy"
        />
        <div
          :class="item.textalign"
          class="tw-absolute tw--mt-2 xs:tw-mt-0 tw-bottom-[1.53rem] tw-flex tw-flex-col tw-text-black-core/[0.87]"
        >
          <span class="sub-h3-regular" v-html="item.content"></span>
          <div class="tw-mt-2 tw-flex tw-flex-wrap tw-justify-center">
            <router-link
              class="white-btn gradient-border-btn primary-btn tw-border-none"
              to="/contact"
              @click.native="mixpanel.track('tap_android_development_cta4')"
            >
              <span class="sub-h3-semibold">Inquire Now</span>
            </router-link>
          </div>
        </div>
      </div>
    </div>
    <!-- ------------------------ MOBILE UI END---------------------- -->

    <!-- ---------------------DESKTOP UI START------------------ -->
    <div class="tw-hidden lg:tw-block tw-bg-[#F6F6F6] tw-mt-[4.5rem]">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :loop="true"
        :speed="4000"
        :autoplay="{
          delay: 4000,
          disableOnInteraction: false,
        }"
        :spaceBetween="20"
        :navigation="true"
        :modules="modules"
        @autoplayTimeLeft="onAutoplayTimeLeft"
        class="swiper-container xll:tw-container"
      >
        <swiper-slide v-for="item in cases" :key="item">
          <div
            class="tw-flex tw-flex-row xl:tw-px-16 tw-py-24 tw-w-[91%] 2xl:tw-w-[77%] 2xll:tw-w-[72%] xll:tw-w-[91%] tw-mx-auto"
          >
            <div
              class="tw-flex tw-flex-col tw-w-[77%] xl:tw-w-[48%] 2xll:tw-w-[50%] lg:tw-mr-auto xll:tw-ml-auto"
            >
              <div>
                <span class="tw-text-[#FF9472] tw-text-left sub-h3-semibold">
                  {{ item.title }}
                </span>
              </div>
              <a
                class="tw-cursor-pointer"
                :href="item.url"
                @click.native="mixpanel.track('tap_android_app_portfolio')"
                :target="item.target"
              >
                <h2
                  class="tw-mt-4 tw-font-opensans-semibold header-3 tw-text-black-core/[0.87]"
                  v-html="item.deskcontent"
                ></h2
              ></a>
              <span
                class="tw-mt-4 tw-text-black-core/[0.6] sub-h3-medium"
                v-html="item.description"
              ></span>

              <img
                :src="resultsimage"
                class="tw-h-12 tw-w-12 lg:tw-h-20 lg:tw-w-20 tw-object-cover tw-ml-[60%] tw-mt-8 tw-mb-4 2xll:tw-ml-[50%]"
                alt="case_study_result_image"
              />

              <div
                class="tw-bg-white tw-drop-shadow-md tw-rounded lg:tw-w-[457px]"
              >
                <div class="tw-inline-flex tw-p-2 lg:tw-p-4">
                  <div class="tw-p-[0.8rem]">
                    <span
                      class="tw-font-opensans-semibold header-3 tw-text-[#FF9472]"
                      >{{ item.resultcount[0] }}</span
                    >
                    <div
                      class="tw-flex-1 tw-text-black-core-[0.87] sub-h3-regular"
                    >
                      <span>Downloads</span>
                    </div>
                  </div>
                  <div class="tw-border-l-[0.5px] tw-p-[0.8rem]">
                    <span
                      class="tw-font-opensans-semibold header-3 tw-text-[#FF9472]"
                      >{{ item.resultcount[1] }}</span
                    >
                    <div
                      class="tw-flex-1 tw-text-black-core-[0.87] sub-h3-regular"
                    >
                      <span>App Rating</span>
                    </div>
                  </div>
                  <div class="tw-border-l-[0.5px] tw-p-[0.8rem]">
                    <span
                      class="tw-font-opensans-semibold header-3 tw-text-[#FF9472]"
                      >{{ item.resultcount[2] }}</span
                    >
                    <div
                      class="tw-flex-1 tw-text-black-core-[0.87] sub-h3-regular"
                    >
                      <span>Monthly Users</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div>
              <a
                class="tw-cursor-pointer"
                :href="item.url"
                @click.native="mixpanel.track('tap_android_app_portfolio')"
                :target="item.target"
              >
                <img
                  :src="item.images[2]"
                  :srcset="`${item.images[2]} 400w,${item.images[3]} 800w`"
                  class="tw-mt-12 tw-h-80 lg:tw-h-[24.75006rem] xl:tw-h-[29.75006rem] tw-w-[25.87525rem] lg:tw-w-[30.87525rem] xl:tw-w-[29.87525rem] tw-object-fit"
                  alt="CaseStudyImage"
              /></a>
            </div>
          </div>
        </swiper-slide>
        <div
          class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-1 tw-bg-black-core/[0.08]"
        >
          <div
            class="progressbar tw-h-1 tw-bg-gradient-to-r tw-from-[#FF835B] tw-to-[#F2709C] tw-rounded-xl"
            ref="progressBar"
          ></div>
        </div>
      </swiper>
    </div>
    <!-- ---------------------DESKTOP UI END------------------ -->
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
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import config from "@/config.js";
import { ref } from "vue";

export default {
  setup() {
    const progressBar = ref(null);
    const onAutoplayTimeLeft = (s, time, progress) => {
      progressBar.value.style.setProperty("--progress", `${progress * 100}%`);
    };

    return {
      onAutoplayTimeLeft,
      progressBar,
      resultsimage,
      modules: [Autoplay],
      cases: [
        {
          title: "Justly",
          url: "/portfolio/justly",
          target: "_self",
          content: "Create systems to build your dream life",
          deskcontent: `Overcome depression <br>by building habits`,
          resultcount: ["10k+", "4.9", "1100+"],
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
          url: "/portfolio/luxeradio",
          target: "_self",
          content: "Radio, music and podcast",
          deskcontent: `A radio, music and <br> podcast platform`,
          resultcount: ["300k+", "5", "25k+"],
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
          url: "/portfolio/togness",
          target: "_self",
          content: "Create your life’s most memorable events",
          deskcontent: `Photo editor and <br> video maker app`,
          resultcount: ["25k+", "4.7 ", "2700+"],
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
          url: config.SMILEPLUS_URL,
          target: "_blank",
          content: "Create systems to build your dream life",
          deskcontent: `For Dentists and <br>Dental labs`,
          resultcount: ["25k+", "4.9", "5k+"],
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
          content: `Let’s Start a <span class="header-2">New Project</span> Together`,
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
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";
.progressbar {
  width: calc(100% - var(--progress, 0%));
  transition: width 0.04s ease;
}
</style>
