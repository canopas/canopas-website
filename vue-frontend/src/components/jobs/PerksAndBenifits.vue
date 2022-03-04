<template>
  <div class="container">
    <div class="text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Perks</span> and Benefits
      </div>
      <div class="description jobs-normal-text">
        Whether it's about your learning, well being, or competitive
        compensation, we've got you covered. Check out all perks and benefits
        offered at Canopas! Many more to come as we grow!
      </div>
    </div>
    <div class="flex-div">
      <div v-for="perk in perks" :key="perk.id" class="flex-elements">
        <aspect-ratio height="120%" v-if="perk.image">
          <img
            class="image"
            :src="perk.image"
            loading="lazy"
            alt="perks-and-benefits-image"
          />
        </aspect-ratio>
        <aspect-ratio v-else :style="{ backgroundColor: perk.bgColor }">
          <div class="details">
            <div class="normal-2-text title">
              <img :src="perk.icon" loading="lazy" />
              <span>{{ perk.title }}</span>
            </div>
            <div class="normal-2-text text-center">
              <div v-html="perk.description" @click="scrollToCareer"></div>
            </div>
          </div>
        </aspect-ratio>
      </div>
    </div>
  </div>
</template>

<script type="module">
import AspectRatio from "@/components/utils/AspectRatio.vue";

export default {
  data() {
    return {
      perks: [
        {
          id: 1,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_1.jpg"),
          description: "",
        },
        {
          id: 2,
          title: "Learning & Development",
          icon: require("@/assets/images/benifits/jobs_canopas_learning.svg"),
          bgColor: "rgba(249, 129, 136, 0.08)",
          image: "",
          description:
            "You will get exposure to new challenges, weekly internal knowledge sharing sessions, book library for career and personal growth.",
        },
        {
          id: 3,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_2.jpg"),
          description: "",
        },
        {
          id: 4,
          title: "Health & Wellbeing",
          icon: require("@/assets/images/benifits/jobs_canopas_health.svg"),
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "We organize weekly cardio sessions and daily meditation in the office. We also organize yoga events or similar activities monthly.",
        },
        {
          id: 5,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_3.jpg"),
          description: "",
        },
        {
          id: 6,
          title: "Equipment And Tools",
          icon: require("@/assets/images/benifits/jobs_canopas_laptop.svg"),
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "Sitting desks to standing desks, up-to-date desktop or laptop, MacBooks, and other tools and accessories based on your requirements.",
        },
        {
          id: 7,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_4.jpg"),
          description: "",
        },
        {
          id: 8,
          title: "Celebrations",
          icon: require("@/assets/images/benifits/jobs_canopas_celebration.svg"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_5.jpg"),
          description: "",
        },
        {
          id: 10,
          title: "Lunch and learn sessions",
          icon: require("@/assets/images/benifits/jobs_canopas_cooking.svg"),
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "We organize bi-weekly or monthly lunch and learn sessions where one of our team members shares knowledge about the tech and non-tech topics like AI, blockchain, ML, etc.",
        },
        {
          id: 11,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_6.jpg"),
          description: "",
        },
        {
          id: 12,
          title: "Flexible working",
          icon: require("@/assets/images/benifits/jobs_canopas_clock.svg"),
          bgColor: "rgba(61, 61, 61, 0.08)",
          image: "",
          description:
            "We offer flexible working hours ranging from 7 AM to 9 PM for early birds and night owls. We also don't put unnecessary restrictions in the workplace. We want everyone to enjoy their work by being themselves. ",
        },
        {
          id: 13,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_7.jpg"),
          description: "",
        },
        {
          id: 14,
          title: "We are hiring",
          icon: require("@/assets/images/benifits/jobs_canopas_hiring.svg"),
          bgColor: "rgba(249, 129, 136, 0.08)",
          image: "",
          description: `Want to become part of a Canopas team? Check out our openings and <div class="career-link" style="color: #fe88a2; font-weight:700; cursor:pointer"> apply today.</div>`,
        },
        {
          id: 15,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_8.jpg"),
          description: "",
        },
      ],
    };
  },
  components: {
    AspectRatio,
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

<style lang="scss" scoped>
.container {
  margin: 100px auto;
  padding: 0 10%;
}

.description {
  margin: 30px 10px;
  color: rgba(61, 61, 61, 0.8);
}

.flex-div {
  display: flex;
  flex-direction: column;
}

.flex-elements > div,
.image {
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
}

.flex-elements {
  padding: 20px;
}

.flex-elements:nth-child(even) {
  .vue-aspect-ratio {
    padding-top: 120% !important;
  }
}

.image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.details {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 16px;

  div {
    flex: none !important;
  }

  .text-center {
    margin-top: 10px;
  }
}

.title {
  display: inline-flex;
  align-items: center;
  color: #3d3d3d;
  font-weight: 700;

  span {
    margin-left: 10px;
    word-wrap: break-word;
    overflow-wrap: break-word;
  }
}

@media (min-width: 440px) {
  .flex-elements:nth-child(even) {
    padding: 10px 40px;
  }

  .details {
    padding: 40px;
  }

  .flex-elements:nth-child(even) {
    .vue-aspect-ratio {
      padding-top: 100% !important;
    }
  }
}

@include media-breakpoint-only(md) {
  .container {
    max-width: 760px;
  }

  .flex-elements:nth-child(even) {
    .vue-aspect-ratio {
      padding-top: 120% !important;
    }
  }
}

@include media-breakpoint-up(md) {
  .container {
    padding: 0;
  }

  .description {
    margin: 50px 80px;
  }

  .flex-div {
    display: flex;
    flex-flow: wrap;
    align-items: center;

    div {
      flex: 30%;
    }
  }

  .flex-elements,
  .flex-elements:nth-child(even) {
    padding: 10px;
  }

  .details {
    padding: 16px;
  }
}

@include media-breakpoint-up(lg) {
  .flex-elements,
  .flex-elements:nth-child(even) {
    padding: 20px;
  }

  .details {
    padding: 30px;
  }
}

@include media-breakpoint-up(xl) {
  .details {
    padding: 40px;
  }
}
</style>
