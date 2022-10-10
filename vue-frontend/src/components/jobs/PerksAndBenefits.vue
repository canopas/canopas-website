<template>
  <div
    class="tw-container tw-my-[100px] tw-mx-auto tw-py-0 tw-px-[5%] sm:tw-px-[10%] only:md:tw-max-w-[760px] md:tw-p-0"
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
            :src="perk.image"
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
              <img :src="perk.icon" loading="lazy" :alt="perk.alt" />
              <span class="tw-ml-[10px] tw-text-[1.3rem] lg:tw-text-[1.5rem]">{{
                perk.title
              }}</span>
            </div>
            <div class="normal-2-text tw-text-center tw-flex-none tw-mt-[10px]">
              <div
                v-html="perk.description"
                @click="scrollToCareer"
                @click.native="$gtag.event('tap_perks_hiring')"
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

import perks1 from "@/assets/images/perks/jobs_canopas_perks_1.webp";
import perks2 from "@/assets/images/perks/jobs_canopas_perks_2.webp";
import perks3 from "@/assets/images/perks/jobs_canopas_perks_3.webp";
import perks4 from "@/assets/images/perks/jobs_canopas_perks_4.webp";
import perks5 from "@/assets/images/perks/jobs_canopas_perks_5.webp";
import perks6 from "@/assets/images/perks/jobs_canopas_perks_6.webp";
import perks7 from "@/assets/images/perks/jobs_canopas_perks_7.webp";
import perks8 from "@/assets/images/perks/jobs_canopas_perks_8.webp";

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
          image: perks1,
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
          image: perks2,
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
          image: perks3,
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
          image: perks4,
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
          image: perks5,
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
          image: perks6,
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
          image: perks7,
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
          image: perks8,
          description: "",
        },
      ],
      isMobile: false,
    };
  },
  components: {
    AspectRatio,
  },

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
