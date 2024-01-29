<template>
  <section
    class="mt-16 md:mt-40 lg:-mt-48 xl:-mt-[32rem] 2xl:mt-[-30rem] 3xl:outer-container"
  >
    <div class="container">
      <div class="flex flex-col text-center">
        <span class="mobile-header-2 lg:desk-header-2 text-black-87"
          >Case studies</span
        >
        <span
          class="mt-4 md:mt-6 sub-h1-regular lg:mobile-header-2-regular mx-auto text-black-60 lg:w-[80%] xl:w-[61%]"
          >Explore our case studies to witness how we've transformed our
          client's ideas into successful Android apps.</span
        >
      </div>
      <!-- ---------------------MOBILE UI START------------------ -->
      <div
        v-for="item in cases"
        :key="item"
        class="mt-6 relative lg:hidden"
        :class="item.className"
      >
        <a
          class="cursor-pointer"
          :href="item.url"
          @click.native="$mixpanel.track('tap_android_app_portfolio')"
          :target="item.target"
        >
          <img
            :src="item.images[0]"
            :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
            alt="casestudy-image"
            class="w-full h-full object-cover"
            loading="lazy"
          />

          <div :class="item.textalign" class="absolute bottom-4 flex flex-col">
            <span class="mobile-header-2 text-black-87">{{ item.title }}</span>
            <span
              class="mt-2 sub-h3-regular text-black-60"
              v-html="item.content"
            ></span>
          </div>
        </a>
      </div>
      <div
        v-for="item in cta"
        :key="item"
        class="mt-16 relative lg:hidden"
        :class="item.className"
      >
        <img
          :src="item.images[0]"
          :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w,`"
          alt="casestudy-image"
          class="w-full h-full object-cover"
          loading="lazy"
        />
        <div
          :class="item.textalign"
          class="absolute -mt-2 xs:mt-0 bottom-[1.53rem] flex flex-col text-black-87"
        >
          <span class="sub-h3-regular" v-html="item.content"></span>
          <div class="mt-2 flex flex-wrap justify-center">
            <router-link
              class="white-btn gradient-border-btn primary-btn border-none"
              to="/contact"
              @click.native="$mixpanel.track('tap_android_development_cta4')"
            >
              <span class="sub-h3-semibold">Inquire Now</span>
            </router-link>
          </div>
        </div>
      </div>
    </div>
    <!-- ------------------------ MOBILE UI END---------------------- -->

    <!-- ---------------------DESKTOP UI START------------------ -->
    <div class="hidden lg:block bg-white-smoke-2 mt-[4.5rem]">
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
        class="swiper-container xll:container"
      >
        <swiper-slide v-for="item in cases" :key="item">
          <div
            class="flex flex-row xl:px-16 py-24 w-[91%] 2xl:w-[77%] 2xll:w-[72%] xll:w-[91%] mx-auto"
          >
            <div
              class="flex flex-col w-[77%] xl:w-[48%] 2xll:w-[50%] lg:mr-auto xll:ml-auto"
            >
              <div>
                <span class="text-primary-1 text-left mobile-header-3-semibold">
                  {{ item.title }}
                </span>
              </div>
              <a
                class="cursor-pointer"
                :href="item.url"
                @click.native="$mixpanel.track('tap_android_app_portfolio')"
                :target="item.target"
              >
                <h2
                  class="mt-4 font-opensans-semibold desk-header-3 text-black-87"
                  v-html="item.deskcontent"
                ></h2
              ></a>
              <span
                class="mt-4 text-black-60 sub-h3-medium md:mobile-header-3-medium"
                v-html="item.description"
              ></span>

              <img
                :src="resultsimage"
                class="h-12 w-12 lg:h-20 lg:w-20 object-cover ml-[60%] mt-8 mb-4 2xll:ml-[50%]"
                alt="case_study_result_image"
              />

              <div class="bg-white drop-shadow-md rounded lg:w-[457px]">
                <div class="inline-flex p-2 lg:p-4">
                  <div class="p-[0.8rem]">
                    <span
                      class="font-opensans-semibold desk-header-3 text-primary-1"
                      >{{ item.resultcount[0] }}</span
                    >
                    <div class="flex-1 text-black-87 mobile-header-3-regular">
                      <span>Downloads</span>
                    </div>
                  </div>
                  <div class="border-l-[0.5px] p-[0.8rem]">
                    <span
                      class="font-opensans-semibold desk-header-3 text-primary-1"
                      >{{ item.resultcount[1] }}</span
                    >
                    <div class="flex-1 text-black-87 mobile-header-3-regular">
                      <span>App Rating</span>
                    </div>
                  </div>
                  <div class="border-l-[0.5px] p-[0.8rem]">
                    <span
                      class="font-opensans-semibold desk-header-3 text-primary-1"
                      >{{ item.resultcount[2] }}</span
                    >
                    <div class="flex-1 text-black-87 mobile-header-3-regular">
                      <span>Monthly Users</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div>
              <a
                class="cursor-pointer"
                :href="item.url"
                @click.native="$mixpanel.track('tap_android_app_portfolio')"
                :target="item.target"
              >
                <img
                  :src="item.images[2]"
                  :srcset="`${item.images[2]} 400w,${item.images[3]} 800w`"
                  class="mt-12 h-80 lg:h-[24.75006rem] xl:h-[29.75006rem] w-[25.87525rem] lg:w-[30.87525rem] xl:w-[29.87525rem] object-fit"
                  alt="CaseStudyImage"
              /></a>
            </div>
          </div>
        </swiper-slide>
        <div class="absolute top-0 left-0 w-full h-1 bg-black-8">
          <div
            class="progressbar h-1 bg-gradient-to-r from-orange-300 to-primary-1 rounded-xl"
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
          className: "text-left items-start",
          textalign: "w-[62%] ml-[1rem]",
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
          className: "text-left items-start",
          textalign: "right-[10%] ml-[1rem] ",
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
          className: "text-left items-start ",
          textalign: "w-[65%] ml-[1rem] ",
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
          className: "text-left items-start ",
          textalign: "right-[5%] w-[52%] ml-[1rem]",
        },
      ],
      cta: [
        {
          title: "",
          content: `Let’s Start a <span class="mobile-header-2">New Project</span> Together`,
          images: [cta_400w, cta_800w],
          className: "text-center ",
          textalign:
            "top-[65%] sm:top-[68%]  w-full left-1/2 transform -translate-x-1/2 -translate-y-1/2 ",
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
