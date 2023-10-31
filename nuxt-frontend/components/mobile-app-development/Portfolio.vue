<template>
  <section
    class="py-[50px] md:-mt-1px md:bg-multi-gradient-background md:pb-[200px] md:pt-[150px]"
  >
    <!-- MOBILE UI START -->
    <div class="container md:hidden">
      <p class="text-center font-inter-bold text-[1.875rem] leading-9">
        CASE STUDIES
      </p>
      <p
        class="v2-canopas-gradient-text mt-2 text-center font-inter-medium text-[1rem] leading-[1.2rem]"
      >
        1M+ users use our apps everyday
      </p>
      <div
        class="mt-10 flex flex-col"
        v-for="(portfolio, index) in portfolios"
        :key="index"
      >
        <div class="flex flex-row items-center gap-2">
          <img
            :src="portfolio.logo"
            alt="app-logo"
            class="h-[30px] w-[30px] object-contain"
            :class="index == 3 ? 'w-[75px]' : ''"
          />
          <p
            class="font-inter-bold text-[1.25rem] leading-6 text-black-core/[0.87]"
          >
            {{ portfolio.name }}
          </p>
        </div>
        <div
          class="mt-3 font-inter-medium text-[1.2rem] leading-[1.6rem] text-black-core/[0.60]"
        >
          {{ portfolio.description }}
        </div>
        <img
          :src="portfolio.image[0]"
          :srcset="`${portfolio.image[0]} 400w, ${portfolio.image[1]} 600w,  ${portfolio.image[2]} 800w}`"
          alt="app-image"
          class="mx-auto mt-8 h-full max-w-[95%]"
          @click="openPortfolio(portfolio)"
        />
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="container hidden md:block">
      <div class="flex flex-col items-center">
        <p
          class="text-center font-inter-bold text-[3.125rem] leading-[3.438rem] text-white xl:text-[4.125rem] xl:leading-[4.95rem]"
        >
          Our portfolios
        </p>
        <p
          class="v2-canopas-gradient-text mt-1 text-center font-inter-medium text-[1.25rem] leading-6 xl:text-[1.5rem] xl:leading-[1.8rem]"
        >
          1M+ users use our apps everyday
        </p>
        <p
          class="mt-6 w-[73%] text-center font-inter-regular text-[1.25rem] leading-[1.5rem] text-white/[0.85] lg:w-[58%] xl:text-[1.5rem] xl:leading-[1.8rem] 2xl:w-[50%]"
        >
          Discover how we transformed the vision of our clients into successful
          mobile applications.
        </p>
      </div>
      <div
        class="mt-32 flex flex-col"
        v-for="(portfolio, index) in portfolios"
        :key="index"
      >
        <div
          class="flex flex-row gap-6 lg:gap-14"
          :class="
            portfolio.id == 2 || portfolio.id == 4 ? 'flex-row-reverse' : ''
          "
        >
          <div
            class="flex flex-[61%] flex-col justify-center 2xl:w-[70%]"
            :class="
              portfolio.id == 2 || portfolio.id == 4
                ? 'lg:pl-[4.5rem] 2xl:pl-[6.5rem]'
                : ''
            "
          >
            <div class="flex flex-row items-center gap-2">
              <img
                :src="
                  portfolio.id == 2 || portfolio.id == 3
                    ? portfolio.desktopLogo
                    : portfolio.logo
                "
                alt="app-logo"
                class="h-[30px] w-[30px] object-contain"
                :class="index == 3 ? 'w-[75px]' : ''"
              />
              <p
                class="font-inter-semibold text-[1.25rem] leading-[1.513rem]"
                :class="
                  portfolio.id == 1
                    ? 'from-[#26AEF9] to-[#2671FA] bg-clip-text text-transparent bg-gradient-[180deg]'
                    : 'text-white '
                "
              >
                {{ portfolio.name }}
              </p>
            </div>
            <p
              class="mt-2 w-[90%] cursor-pointer font-inter-semibold text-[2.063rem] leading-10 text-white xl:text-[3rem] xl:leading-[3.631rem]"
              :class="
                portfolio.id == 1 || portfolio.id == 3
                  ? 'lg:w-[67%] xl:w-[81%] 2xl:w-[83%]'
                  : 'lg:w-[78%] xl:w-[89%] 2xl:w-[82%]'
              "
              @click="openPortfolio(portfolio)"
            >
              <span
                class="bg-multiple-underline-out box-decoration-clone bg-[length:0%] bg-no-repeat pb-[5px] transition-all duration-500 hover:bg-[length:100%]"
                >{{ portfolio.title }}</span
              >
            </p>
            <p
              class="mt-6 font-inter-medium text-[1.125rem] leading-[1.438rem] text-white/[0.85] lg:w-[94%] xl:text-[1.25rem] xl:leading-[1.513rem]"
              :class="
                portfolio.id == 1 || portfolio.id == 3
                  ? 'xl:w-[85%] 2xl:w-[100%]'
                  : ''
              "
            >
              {{ portfolio.description }}
            </p>
          </div>
          <div
            class="items-enter flex-[39%] cursor-pointer justify-center 2xl:w-[30%]"
          >
            <img
              :src="portfolio.desktopImage[0]"
              :srcset="`${portfolio.desktopImage[0]} 400w, ${portfolio.desktopImage[1]} 600w, ${portfolio.desktopImage[2]} 800w}`"
              alt="app-image"
              class="h-full w-full object-contain"
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
      this.$mixpanel.track(portfolio.event);
    },
  },
};
</script>
