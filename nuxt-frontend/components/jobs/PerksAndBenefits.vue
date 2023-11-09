<template>
  <div
    class="container my-[14.063rem] mx-auto py-0 px-[5%] sm:px-[10%] only:md:max-w-[760px] md:p-0"
  >
    <div class="text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Perks</span> and Benefits
      </div>
      <div
        class="description my-[30px] mx-2.5 md:my-[50px] md:mx-20 normal-text"
      >
        Whether it's about your learning, well being, or competitive
        compensation, we've got you covered. Check out all perks and benefits
        offered at Canopas! Many more to come as we grow!
      </div>
    </div>
    <div
      class="flex-div flex flex-col md:flex-row md:flex-wrap md:items-center"
    >
      <div
        v-for="perk in perks"
        :key="perk.id"
        class="flex-elements p-5 even:px-[30px] sm:even:px-10 md:flex-[30%] md:p-2.5 md:even:py-2.5 md:even:px-2.5 lg:p-5 lg:even:p-5"
        :class="isMobile ? 'even:py-2.5 even:px-10' : ''"
      >
        <aspect-ratio
          height="120%"
          v-if="perk.image"
          class="shadow-[0_4px_4px_rgba(0,0,0,0.25)] rounded-[10px]"
        >
          <img
            class="image shadow-[0_4px_4px_rgba(0,0,0,0.25)] rounded-2.5 w-full h-full object-cover ease-in-out duration-300 hover:scale-110"
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
          class="rounded-[10px]"
        >
          <div
            class="details shadow-[0_4px_4px_rgba(0,0,0,0.25)] rounded-[10px] flex flex-col justify-center items-center h-full p-4 sm:p-[30px] md:p-0.5 lg:p-[30px] xl:p-10"
            :class="isMobile ? 'p-10' : ''"
          >
            <div
              class="normal-2-text flex-none title inline-flex items-center text-black-900 font-bold"
            >
              <img
                :src="perk.icon"
                loading="lazy"
                class="w-[1.875rem] h-[1.875rem]"
                :alt="perk.alt"
              />
              <span class="ml-2.5 text-[1.3rem] lg:text-[1.5rem]">{{
                perk.title
              }}</span>
            </div>
            <div class="normal-2-text text-center flex-none mt-2.5">
              <div
                v-html="perk.description"
                @click="scrollToCareer"
                @click.native="$mixpanel.track('tap_perks_hiring')"
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
import axios from "axios";
import config from "@/config.js";
import { getDiffrentWidthImages } from "@/utils.js";
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
          image: [],
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
          image: [],
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
          image: [],
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
          image: [],
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
          image: [],
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
          image: [],
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
          image: [],
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
          image: [],
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
    async fetchImages() {
      try {
        const response = await axios.get(config.API_BASE + "/api/perksimages");
        var slides = getDiffrentWidthImages(response);
        this.perks.forEach((perk, index) => {
          if (index % 2 === 0 && index / 2 < slides.length) {
            const slideIndex = Math.floor(index / 2);
            perk.image = [
              slides[slideIndex].image[0],
              slides[slideIndex].image[1],
            ];
          }
        });
      } catch (error) {
        console.error("Error fetching data from the API", error);
      }
    },
  },
  created() {
    this.fetchImages();
  },
};
</script>
