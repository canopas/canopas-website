<template>
  <div class="tw-container tw-pt-[50px] md:tw-pt-[100px]">
    <div class="tw-pb-12 tw-px-2">
      <div
        class="border-text v2-header-2-text tw-leading-[4rem] tw-font-futura-bold md:tw-leading-[5.0875rem] lg:tw-leading-[6.6875rem]"
      >
        We are
      </div>
      <div
        class="border-text v2-header-2-text tw-leading-[4rem] tw-font-futura-bold md:tw-leading-[5.0875rem] lg:tw-leading-[6.6875rem]"
      >
        here to
      </div>
      <div class="v2-title-3-text">
        <span
          class="tw-leading-[4rem] v2-header-2-text tw-font-futura-bold md:tw-leading-[5.0875rem] lg:tw-leading-[6.6875rem] tw-pr-2"
          >contribute</span
        >
        and that starts with the site itself,
      </div>
      <div class="v2-title-3-text tw-mt-1">
        <a
          class="v2-canopas-gradient-text tw-font-semibold gradient-underline tw-pb-[1.5px] tw-relative tw-no-underline after:tw-bg-gradient-to-r tw-from-orange-300 tw-to-pink-300 after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-[1.5px]"
          :href="websiteOpenSourceUrl"
          target="_blank"
          @click.native="$gtag.event('tap_canopas_open_source_link')"
          >Canopas is open source!</a
        >
        We also have several projects that are widely recognized in the open
        source community.
      </div>
    </div>
    <div
      class="tw-flex tw-flex-col tw-w-full tw-m-auto md:tw-flex-row md:tw-w-[95%]"
    >
      <div
        v-for="contribution in contributions"
        :key="contribution"
        class="tw-rounded-[10px] tw-m-5 tw-flex tw-flex-col tw-justify-center tw-cursor-pointer tw-shadow-[0_4px_4px_rgba(0,0,0,0.25)]"
        :style="{ background: contribution.bgColor }"
      >
        <img
          :src="contribution.image"
          class="tw-w-full tw-h-full tw-object-cover tw-rounded-[10px]"
          loading="lazy"
        />

        <div class="tw-w-full tw-p-[20px] v2-title-3-text">
          <a
            class="tw-float-right tw-flex tw-items-end tw-duration-300 tw-ease-in-out hover:tw-scale-[1.1] tw-hover-hover"
            :href="contribution.link"
            target="_blank"
            @click.native="$gtag.event('tap_contribution_library')"
          >
            <font-awesome-icon
              class="tw-h-[30px] tw-width-[30px] tw-text-pink-300 tw-mr-2.5 lg:tw-h-10 lg:tw-w-10"
              :icon="icon"
            />
            <div class="v2-canopas-gradient-text tw-font-semibold">GitHub</div>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import introShowcase from "@/assets/images/contribution/intro-showcase.webp";
import uiPilot from "@/assets/images/contribution/ui-pilot.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
import Config from "@/config.js";

library.add(faGithub);

export default {
  data() {
    return {
      icon: faGithub,
      websiteOpenSourceUrl: Config.WEBSITE_OPEN_SOURCE_URL,
      contributions: [
        {
          image: introShowcase,
          bgColor: "#FAF5F2",
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          image: uiPilot,
          bgColor: "#3D3D3D",
          link: "https://github.com/canopas/UIPilot",
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>
