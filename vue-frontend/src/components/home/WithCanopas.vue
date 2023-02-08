<template>
  <div
    class="tw-bg-gradient-to-t tw-from-orange-300/[0.15] tw-to-white lg:tw-mb-[100px]"
  >
    <div
      class="tw-container tw-pb-4 tw-pt-12 md:tw-pb-12 md:tw-pt-16 lg:tw-pb-20 lg:tw-pt-32"
    >
      <div class="v2-header-3-text tw-m-auto sm:tw-w-[90%] tw-text-center">
        With <span class="canopas-gradient-text">canopas</span>, you will have
        everything you need under one roof.
      </div>
      <div class="tw-mt-8 md:tw-mt-20 lg:tw-mt-40">
        <div
          ref="service1"
          class="tw-flex tw-flex-wrap tw-m-auto md:tw-mt-20 md:tw-mx-auto md:tw-mb-auto lg:even:tw-w-[70%]"
        >
          <div
            v-for="(phase, i) in phases.slice(0, 3)"
            :key="i"
            class="tw-text-center tw-py-4 tw-px-0 md:tw-flex-[1_1_30%] md:tw-flex md:tw-flex-col md:tw-p-0 md:tw-text-left last-of-type:md:tw-py-0 last-of-type:md:tw-px-40 last-of-type:lg:tw-p-0"
          >
            <aspect-ratio height="56.25%">
              <img
                :src="phase.image"
                class="tw-w-full tw-h-full tw-object-cover"
                loading="lazy"
                :alt="phase.alt"
              />
            </aspect-ratio>
            <div
              class="tw-p-4 sm:tw-py-4 sm:tw-px-12 md:tw-py-4 md:tw-pr-12 md:tw-pl-6 lg:tw-py-4 lg:tw-px-6"
            >
              <div
                class="tw-p-2 canopas-gradient-text v2-title-2-text tw-bg-gradient-to-r tw-from-pink-300 tw-to-orange-300 tw-bg-clip-text tw-text-transparent"
              >
                {{ phase.title }}
              </div>
              <div class="v2-normal-2-text md:tw-p-2">
                {{ phase.description }}
              </div>
            </div>
          </div>
        </div>
        <div
          ref="service2"
          class="tw-flex tw-flex-wrap tw-m-auto md:tw-mt-20 md:tw-mx-auto md:tw-mb-auto lg:even:tw-w-[70%]"
        >
          <div
            v-for="(phase, i) in phases.slice(3, 5)"
            :key="i"
            class="tw-text-center tw-py-4 tw-px-0 md:tw-flex-[1_1_30%] md:tw-flex md:tw-flex-col md:tw-p-0 md:tw-text-left"
          >
            <aspect-ratio height="56.25%">
              <img
                :src="phase.image"
                class="tw-w-full tw-h-full tw-object-cover"
                :class="
                  phase.title == 'Marketing' ? 'tw-h-full md:tw-h-[84%]' : ''
                "
                loading="lazy"
                :alt="phase.alt"
              />
            </aspect-ratio>
            <div
              class="tw-p-4 sm:tw-py-4 sm:tw-px-12 md:tw-py-4 md:tw-pr-12 md:tw-pl-6 lg:tw-py-4 lg:tw-px-6"
            >
              <div
                class="tw-p-2 canopas-gradient-text v2-title-2-text tw-bg-gradient-to-r tw-from-pink-300 tw-to-orange-300 tw-bg-clip-text tw-text-transparent"
              >
                {{ phase.title }}
              </div>
              <div class="v2-normal-2-text md:tw-p-2">
                {{ phase.description }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import AspectRatio from "@/components/utils/AspectRatio.vue";
import planning from "@/assets/images/phases/canopas_phases_planning.gif";
import designing from "@/assets/images/phases/canopas_phases_designing.gif";
import development from "@/assets/images/phases/canopas_phases_development.gif";
import marketing from "@/assets/images/phases/canopas_phases_marketing.gif";
import support from "@/assets/images/phases/canopas_phases_support.gif";
import { elementInViewPort } from "@/utils.js";

export default {
  data() {
    return {
      phases: [
        {
          image: planning,
          alt: "planning",
          title: "Planning",
          description:
            "We research to understand the vision. The audience. The product. The goals. We collaborate, we re-examine, we ask and conclude.",
        },
        {
          image: designing,
          alt: "designing",
          title: "Design",
          description:
            "Working passionately with you, our designers create aesthetically pleasing  and user friendly digital experiences.",
        },
        {
          image: development,
          alt: "development",
          title: "Development",
          description:
            "Backend and Frontend. From the beginning, we have performance and maintainability in mind. Always insisting on best practices and keeping up with the latest trends.",
        },
        {
          image: marketing,
          alt: "marketing",
          title: "Marketing",
          description:
            "From organic growth to writing irresistible ads for your marketing campaign, we're here to help you. We'll help you build a brand and community of prospects who can give you a boost on your product launch.",
        },
        {
          image: support,
          alt: "support",
          title: "Support",
          description:
            "From bug fixing to feature updates, you can count on us. For the last 7 years, our clients had a 99.87% uptime for their products.",
        },
      ],
      event: "",
      events: {
        service1: "view_service_section_first_row",
        service2: "view_service_section_second_row",
      },
    };
  },
  components: {
    AspectRatio,
  },
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  inject: ["mixpanel"],
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (this.mixpanel.__loaded && event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
};
</script>
