<template>
  <div
    class="tw-container tw-my-[14.063rem] tw-mx-auto tw-py-0 tw-px-[5%] sm:tw-px-[10%] only:md:tw-max-w-[760px] md:tw-p-0"
  >
    <div class="tw-text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Perks</span> and Benefits
      </div>
      <div
        class="description tw-my-[30px] tw-mx-[10px] md:tw-my-[50px] md:tw-mx-[80px] normal-text"
      >
        Whether it's about your learning, well being, or competitive
        compensation, we've got you covered. Check out all perks and benefits
        offered at Canopas! Many more to come as we grow!
      </div>
    </div>
    <div
      class="flex-div tw-flex tw-flex-col md:tw-flex-row md:tw-flex-wrap md:tw-items-center"
    >
      <div
        v-for="perk in perks"
        :key="perk.id"
        class="flex-elements tw-p-[20px] even:tw-px-[30px] sm:even:tw-px-[40px] md:tw-flex-[30%] md:tw-p-[10px] md:even:tw-py-[10px] md:even:tw-px-[10px] lg:tw-p-[20px] lg:even:tw-p-[20px]"
        :class="isMobile ? 'even:tw-py-[10px] even:tw-px-[40px]' : ''"
      >
        <aspect-ratio
          height="120%"
          v-if="perk.image"
          class="tw-shadow-[0_4px_4px_rgba(0,0,0,0.25)] tw-rounded-[10px]"
        >
          <img
            class="image tw-shadow-[0_4px_4px_rgba(0,0,0,0.25)] tw-rounded-[10px] tw-w-full tw-h-full tw-object-cover tw-ease-in-out tw-duration-300 hover:tw-scale-110"
            :src="perk.image[0]"
            :srcset="`${perk.image[0]} 400w, ${perk.image[1]} 800w`"
            loading="lazy"
            alt="perks-and-benefits-image"
          />
        </aspect-ratio>
        <aspect-ratio
          v-else
          :height="isMobile ? '120%' : '100%'"
          :style="{ backgroundColor: perk.bgColor }"
          class="tw-rounded-[10px]"
        >
          <div
            class="details tw-shadow-[0_4px_4px_rgba(0,0,0,0.25)] tw-rounded-[10px] tw-flex tw-flex-col tw-justify-center tw-items-center tw-h-full tw-p-[16px] sm:tw-p-[30px] md:tw-p-[2px] lg:tw-p-[30px] xl:tw-p-[40px]"
            :class="isMobile ? 'tw-p-[40px]' : ''"
          >
            <div
              class="normal-2-text tw-flex-none title tw-inline-flex tw-items-center tw-text-black-900 tw-font-bold"
            >
              <img
                :src="perk.icon"
                loading="lazy"
                class="tw-w-[1.875rem] tw-h-[1.875rem]"
                :alt="perk.alt"
              />
              <span class="tw-ml-[10px] tw-text-[1.3rem] lg:tw-text-[1.5rem]">{{
                perk.title
              }}</span>
            </div>
            <div class="normal-2-text tw-text-center tw-flex-none tw-mt-[10px]">
              <div
                v-html="perk.description"
                @click="scrollToCareer"
                @click.native="mixpanel.track('tap_perks_hiring')"
              ></div>
            </div>
          </div>
        </aspect-ratio>
      </div>
    </div>
  </div>
</template>

<script type="module">
import AspectRatio from "@/components/utils/AspectRatio.vue";

import perks1_400w from "@/assets/images/perks/jobs_canopas_perks_1_400w.webp";
import perks1_800w from "@/assets/images/perks/jobs_canopas_perks_1_800w.webp";
import perks2_400w from "@/assets/images/perks/jobs_canopas_perks_2_400w.webp";
import perks2_800w from "@/assets/images/perks/jobs_canopas_perks_2_800w.webp";
import perks3_400w from "@/assets/images/perks/jobs_canopas_perks_3_400w.webp";
import perks3_800w from "@/assets/images/perks/jobs_canopas_perks_3_800w.webp";
import perks4_400w from "@/assets/images/perks/jobs_canopas_perks_4_400w.webp";
import perks4_800w from "@/assets/images/perks/jobs_canopas_perks_4_800w.webp";
import perks5_400w from "@/assets/images/perks/jobs_canopas_perks_5_400w.webp";
import perks5_800w from "@/assets/images/perks/jobs_canopas_perks_5_800w.webp";
import perks6_400w from "@/assets/images/perks/jobs_canopas_perks_6_400w.webp";
import perks6_800w from "@/assets/images/perks/jobs_canopas_perks_6_800w.webp";
import perks7_400w from "@/assets/images/perks/jobs_canopas_perks_7_400w.webp";
import perks7_800w from "@/assets/images/perks/jobs_canopas_perks_7_800w.webp";
import perks8_400w from "@/assets/images/perks/jobs_canopas_perks_8_400w.webp";
import perks8_800w from "@/assets/images/perks/jobs_canopas_perks_8_800w.webp";
import learning from "@/assets/images/benefits/jobs_canopas_learning.svg";
import health from "@/assets/images/benefits/jobs_canopas_health.svg";
import clock from "@/assets/images/benefits/jobs_canopas_clock.svg";
import cooking from "@/assets/images/benefits/jobs_canopas_cooking.svg";
import working from "@/assets/images/benefits/jobs_canopas_working.svg";
import laptop from "@/assets/images/benefits/jobs_canopas_laptop.svg";
import celebration from "@/assets/images/benefits/jobs_canopas_celebration.svg";

export default {
  data() {
    return {
      perks: [
        {
          id: 1,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks1_400w, perks1_800w],
          description: "",
        },
        {
          id: 2,
          title: "5 Days Work Week",
          icon: working,
          alt: "working-icon",
          bgColor: "rgba(249, 129, 136, 0.08)",
          image: "",
          description: `At Canopas, you will experience a truly 5 days work-week. No extra time daily. No shortened lunch break.`,
        },
        {
          id: 3,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks2_400w, perks2_800w],
          description: "",
        },
        {
          id: 4,
          title: "Lunch and learn sessions",
          icon: cooking,
          alt: "cooking-icon",
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "We organize bi-weekly or monthly lunch and learn sessions where one of our team members shares knowledge about the tech and non-tech topics like AI, blockchain, ML, etc.",
        },
        {
          id: 5,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks3_400w, perks3_800w],
          description: "",
        },
        {
          id: 6,
          title: "Flexible working",
          icon: clock,
          alt: "clock-icon",
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "We offer flexible working hours ranging from 7 AM to 9 PM for early birds and night owls. We also don't put unnecessary restrictions in the workplace. We want everyone to enjoy their work by being themselves. ",
        },
        {
          id: 7,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks4_400w, perks4_800w],
          description: "",
        },
        {
          id: 8,
          title: "Celebrations",
          icon: celebration,
          alt: "celebration-icon",
          bgColor: "rgba(249, 129, 136, 0.08)",
          image: "",
          description:
            "Whether it's a festival, birthdays (and births!), work anniversaries, or national holidays - we celebrate wherever we can, be it at the office or anywhere else.",
        },
        {
          id: 9,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks5_400w, perks5_800w],
          description: "",
        },
        {
          id: 10,
          title: "Learning & Development",
          icon: learning,
          alt: "learning-icon",
          bgColor: "rgba(249, 129, 136, 0.08)",
          image: "",
          description:
            "You will get exposure to new challenges, weekly internal knowledge sharing sessions, book library for career and personal growth.",
        },
        {
          id: 11,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks6_400w, perks6_800w],
          description: "",
        },
        {
          id: 12,
          title: "Equipment And Tools",
          icon: laptop,
          alt: "laptop-icon",
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "Sitting desks to standing desks, up-to-date desktop or laptop, MacBooks, and other tools and accessories based on your requirements.",
        },
        {
          id: 13,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks7_400w, perks7_800w],
          description: "",
        },
        {
          id: 14,
          title: "Health & Wellbeing",
          icon: health,
          alt: "health-icon",
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "We organize weekly cardio sessions and daily meditation in the office. We also organize yoga events or similar activities monthly.",
        },
        {
          id: 15,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: [perks8_400w, perks8_800w],
          description: "",
        },
      ],
      isMobile: false,
    };
  },
  components: {
    AspectRatio,
  },
  inject: ["mixpanel"],
  mounted() {
    if (window.innerWidth >= 768 && window.innerWidth < 992) {
      this.isMobile = true;
    }
  },
  methods: {
    scrollToCareer(e) {
      if (e.target.matches(".career-link")) {
        this.$emit("scroll-to-career");
      }
    },
  },
};
</script>
