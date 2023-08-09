<template>
  <section
    class="tw-py-[50px] md:-tw-mt-1px md:tw-bg-multi-gradient-background md:tw-pb-[200px] md:tw-pt-[150px]"
  >
    <!-- MOBILE UI START -->
    <div class="tw-container md:tw-hidden">
      <p
        class="tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-9"
      >
        CASE STUDIES
      </p>
      <p
        class="v2-canopas-gradient-text tw-mt-2 tw-text-center tw-font-inter-medium tw-text-[1rem] tw-leading-[1.2rem]"
      >
        1M+ users use our apps everyday
      </p>
      <div
        class="tw-mt-10 tw-flex tw-flex-col"
        v-for="(portfolio, index) in portfolios"
        :key="index"
      >
        <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
          <img
            :src="portfolio.logo"
            alt="app-logo"
            class="tw-h-[30px] tw-w-[30px] tw-object-contain"
            :class="index == 3 ? 'tw-w-[75px]' : ''"
          />
          <p
            class="tw-font-inter-bold tw-text-[1.25rem] tw-leading-6 tw-text-black-core/[0.87]"
          >
            {{ portfolio.name }}
          </p>
        </div>
        <div
          class="tw-mt-3 tw-font-inter-medium tw-text-[1.2rem] tw-leading-[1.6rem] tw-text-black-core/[0.60]"
        >
          {{ portfolio.description }}
        </div>
        <img
          :src="portfolio.image[0]"
          :srcset="`${portfolio.image[0]} 400w, ${portfolio.image[1]} 600w,  ${portfolio.image[2]} 800w}`"
          alt="app-image"
          class="tw-mx-auto tw-mt-8 tw-h-full tw-max-w-[95%]"
          @click="openPortfolio(portfolio)"
        />
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="tw-container tw-hidden md:tw-block">
      <div class="tw-flex tw-flex-col tw-items-center">
        <p
          class="tw-text-center tw-font-inter-bold tw-text-[3.125rem] tw-leading-[3.438rem] tw-text-white xl:tw-text-[4.125rem] xl:tw-leading-[4.95rem]"
        >
          Our portfolios
        </p>
        <p
          class="v2-canopas-gradient-text tw-mt-1 tw-text-center tw-font-inter-medium tw-text-[1.25rem] tw-leading-6 xl:tw-text-[1.5rem] xl:tw-leading-[1.8rem]"
        >
          1M+ users use our apps everyday
        </p>
        <p
          class="tw-mt-6 tw-w-[73%] tw-text-center tw-font-inter-regular tw-text-[1.25rem] tw-leading-[1.5rem] tw-text-white/[0.85] lg:tw-w-[58%] xl:tw-text-[1.5rem] xl:tw-leading-[1.8rem] 2xl:tw-w-[50%]"
        >
          Discover how we transformed the vision of our clients into successful
          mobile applications.
        </p>
      </div>
      <div
        class="tw-mt-32 tw-flex tw-flex-col"
        v-for="(portfolio, index) in portfolios"
        :key="index"
      >
        <div
          class="tw-flex tw-flex-row tw-gap-6 lg:tw-gap-14"
          :class="
            portfolio.id == 2 || portfolio.id == 4 ? 'tw-flex-row-reverse' : ''
          "
        >
          <div
            class="tw-flex tw-flex-[61%] tw-flex-col tw-justify-center 2xl:tw-w-[70%]"
            :class="
              portfolio.id == 2 || portfolio.id == 4
                ? 'lg:tw-pl-[4.5rem] 2xl:tw-pl-[6.5rem]'
                : ''
            "
          >
            <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
              <img
                :src="
                  portfolio.id == 2 || portfolio.id == 3
                    ? portfolio.desktopLogo
                    : portfolio.logo
                "
                alt="app-logo"
                class="tw-h-[30px] tw-w-[30px] tw-object-contain"
                :class="index == 3 ? 'tw-w-[75px]' : ''"
              />
              <p
                class="tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.513rem]"
                :class="
                  portfolio.id == 1
                    ? 'tw-from-[#26AEF9] tw-to-[#2671FA] tw-bg-clip-text tw-text-transparent tw-bg-gradient-[180deg]'
                    : 'tw-text-white '
                "
              >
                {{ portfolio.name }}
              </p>
            </div>
            <p
              class="tw-mt-2 tw-w-[90%] tw-cursor-pointer tw-font-inter-semibold tw-text-[2.063rem] tw-leading-10 tw-text-white xl:tw-text-[3rem] xl:tw-leading-[3.631rem]"
              :class="
                portfolio.id == 1 || portfolio.id == 3
                  ? 'lg:tw-w-[67%] xl:tw-w-[81%] 2xl:tw-w-[83%]'
                  : 'lg:tw-w-[78%] xl:tw-w-[89%] 2xl:tw-w-[82%]'
              "
              @click="openPortfolio(portfolio)"
            >
              <span
                class="tw-bg-multiple-underline-out tw-box-decoration-clone tw-bg-[length:0%] tw-bg-no-repeat tw-pb-[5px] tw-transition-all tw-duration-500 hover:tw-bg-[length:100%]"
                >{{ portfolio.title }}</span
              >
            </p>
            <p
              class="tw-mt-6 tw-font-inter-medium tw-text-[1.125rem] tw-leading-[1.438rem] tw-text-white/[0.85] lg:tw-w-[94%] xl:tw-text-[1.25rem] xl:tw-leading-[1.513rem]"
              :class="
                portfolio.id == 1 || portfolio.id == 3
                  ? 'xl:tw-w-[85%] 2xl:tw-w-[100%]'
                  : ''
              "
            >
              {{ portfolio.description }}
            </p>
          </div>
          <div
            class="tw-items-enter tw-flex-[39%] tw-cursor-pointer tw-justify-center 2xl:tw-w-[30%]"
          >
            <img
              :src="portfolio.desktopImage[0]"
              :srcset="`${portfolio.desktopImage[0]} 400w, ${portfolio.desktopImage[1]} 600w, ${portfolio.desktopImage[2]} 800w}`"
              alt="app-image"
              class="tw-h-full tw-w-full tw-object-contain"
              @click="openPortfolio(portfolio)"
            />
          </div>
        </div>
      </div>
    </div>
    <!-- DESKTOP UI END -->
  </section>
</template>
<script>
import justly400w from "@/assets/images/mobile-app-development/portfolio/mobile/justly-400w.webp";
import justly600w from "@/assets/images/mobile-app-development/portfolio/mobile/justly-600w.webp";
import justly800w from "@/assets/images/mobile-app-development/portfolio/mobile/justly-800w.webp";
import luxeradio400w from "@/assets/images/mobile-app-development/portfolio/mobile/luxeradio-400w.webp";
import luxeradio600w from "@/assets/images/mobile-app-development/portfolio/mobile/luxeradio-600w.webp";
import luxeradio800w from "@/assets/images/mobile-app-development/portfolio/mobile/luxeradio-800w.webp";
import togness400w from "@/assets/images/mobile-app-development/portfolio/mobile/togness-400w.webp";
import togness600w from "@/assets/images/mobile-app-development/portfolio/mobile/togness-600w.webp";
import togness800w from "@/assets/images/mobile-app-development/portfolio/mobile/togness-800w.webp";
import smile400w from "@/assets/images/mobile-app-development/portfolio/mobile/smile-400w.webp";
import smile600w from "@/assets/images/mobile-app-development/portfolio/mobile/smile-600w.webp";
import smile800w from "@/assets/images/mobile-app-development/portfolio/mobile/smile-800w.webp";
import justlyLogo from "@/assets/images/mobile-app-development/portfolio/mobile/justly-logo.webp";
import luxeradioLogo from "@/assets/images/mobile-app-development/portfolio/mobile/luxeradio-logo.webp";
import tognessLogo from "@/assets/images/mobile-app-development/portfolio/togness-logo.webp";
import smileLogo from "@/assets/images/mobile-app-development/portfolio/smile-logo.webp";
import desktop_justly400w from "@/assets/images/mobile-app-development/portfolio/desktop/justly-400w.webp";
import desktop_justly600w from "@/assets/images/mobile-app-development/portfolio/desktop/justly-600w.webp";
import desktop_justly800w from "@/assets/images/mobile-app-development/portfolio/desktop/justly-800w.webp";
import desktop_luxeradio400w from "@/assets/images/mobile-app-development/portfolio/desktop/luxeradio-400w.webp";
import desktop_luxeradio600w from "@/assets/images/mobile-app-development/portfolio/desktop/luxeradio-600w.webp";
import desktop_luxeradio800w from "@/assets/images/mobile-app-development/portfolio/desktop/luxeradio-800w.webp";
import desktop_togness400w from "@/assets/images/mobile-app-development/portfolio/desktop/togness-400w.webp";
import desktop_togness600w from "@/assets/images/mobile-app-development/portfolio/desktop/togness-600w.webp";
import desktop_togness800w from "@/assets/images/mobile-app-development/portfolio/desktop/togness-800w.webp";
import desktop_smile400w from "@/assets/images/mobile-app-development/portfolio/desktop/smile-400w.webp";
import desktop_smile600w from "@/assets/images/mobile-app-development/portfolio/desktop/smile-600w.webp";
import desktop_smile800w from "@/assets/images/mobile-app-development/portfolio/desktop/smile-800w.webp";
import desktop_justlyLogo from "@/assets/images/mobile-app-development/portfolio/desktop/justly-logo.webp";
import desktop_luxeradioLogo from "@/assets/images/mobile-app-development/portfolio/desktop/luxeradio-logo.webp";
import config from "@/config.js";

export default {
  data() {
    return {
      portfolios: [
        {
          id: 1,
          logo: tognessLogo,
          name: "Togness",
          title: "Photo editor and slideshow maker app",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          image: [togness400w, togness600w, togness800w],
          desktopImage: [
            desktop_togness400w,
            desktop_togness600w,
            desktop_togness800w,
          ],
          url: "/portfolio/togness",
          event: "tap_mobileapp_development_togness_portfolio",
        },
        {
          id: 2,
          logo: justlyLogo,
          desktopLogo: desktop_justlyLogo,
          name: "Justly",
          title: "Overcome depression by building habits",
          description:
            "Justly aims to tackle loneliness, depression, and mental health through innovative solutions.",
          image: [justly400w, justly600w, justly800w],
          desktopImage: [
            desktop_justly400w,
            desktop_justly600w,
            desktop_justly800w,
          ],
          url: "/portfolio/justly",
          event: "tap_mobileapp_development_justly_portfolio",
        },
        {
          id: 3,
          logo: luxeradioLogo,
          desktopLogo: desktop_luxeradioLogo,
          name: "Luxeradio",
          title: "A radio, music and podcast platform",
          description:
            "Luxe Radio displays the best of Moroccan and international creation, emphasizing taste, elegance, and refinement.",
          image: [luxeradio400w, luxeradio600w, luxeradio800w],
          desktopImage: [
            desktop_luxeradio400w,
            desktop_luxeradio600w,
            desktop_luxeradio800w,
          ],
          url: "/portfolio/luxeradio",
          event: "tap_mobileapp_development_luxeradio_portfolio",
        },
        {
          id: 4,
          logo: smileLogo,
          name: "",
          title: "For Dentists and Dental labs",
          description:
            "Smile+ app is designed for dentists to create a perfect smile for their patients automated with AI.",
          image: [smile400w, smile600w, smile800w],
          desktopImage: [
            desktop_smile400w,
            desktop_smile600w,
            desktop_smile800w,
          ],
          url: config.SMILEPLUS_URL,
          target: "_blank",
          event: "tap_mobileapp_development_smile+_portfolio",
        },
      ],
    };
  },
  inject: ["mixpanel"],
  methods: {
    openPortfolio(portfolio) {
      setTimeout(function () {
        window.open(
          portfolio.url,
          portfolio.target ? portfolio.target : "_self",
        );
      }, 200);
      this.mixpanel.track(portfolio.event);
    },
  },
};
</script>
