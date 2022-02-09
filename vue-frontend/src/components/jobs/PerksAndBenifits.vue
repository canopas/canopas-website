<template>
  <div class="container">
    <div class="text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Perks</span> and Benefits section
      </div>
      <div class="description normal-text">
        We wanted to create an environment that helps you get out of bed in the
        morning and an office that creates excellent work and encourages fun. As
        we play the infinite game, it is not about winning or losing, it's about
        striving to be better at whatever we do.
      </div>
    </div>
    <div class="flex-div">
      <div v-for="perk in perks" :key="perk.id" class="padding flex-elements">
        <aspect-ratio height="120%" v-if="perk.image">
          <img
            class="image"
            :src="perk.image"
            loading="lazy"
            alt="perks-and-benefits-image"
          />
        </aspect-ratio>
        <aspect-ratio
          height="90%"
          v-else
          class="padding"
          :style="{ backgroundColor: perk.bgColor }"
        >
          <div class="details padding">
            <div class="normal-text title">
              <img :src="perk.icon" loading="lazy" />
              <span>{{ perk.title }}</span>
            </div>
            <div class="normal-text desc-text">
              {{ perk.description }}
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
          image: require("@/assets/images/perks/jobs_canopas_perks_1.webp"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_2.webp"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_3.webp"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_4.webp"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_5.webp"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_6.webp"),
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
          image: require("@/assets/images/perks/jobs_canopas_perks_7.webp"),
          description: "",
        },
        {
          id: 14,
          title: "We are hiring",
          icon: require("@/assets/images/benifits/jobs_canopas_hiring.svg"),
          bgColor: "rgba(249, 129, 136, 0.08)",
          image: "",
          description:
            "Want to become part of a Canopas team? Check out our openings and apply today.",
        },
        {
          id: 15,
          title: "",
          icon: "",
          bgColor: "transparent",
          image: require("@/assets/images/perks/jobs_canopas_perks_8.webp"),
          description: "",
        },
      ],
    };
  },
  components: {
    AspectRatio,
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin: 100px auto;
}

.description {
  margin: 30px 10px;
  color: #3d3d3d;
}

.padding {
  padding: 10px;
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

.flex-elements:nth-child(even) {
  padding: 10px 30px;
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
  height: 100%;
}

.title {
  display: flex;
  align-items: center;
  color: #3d3d3d;
  font-weight: 700;

  span {
    margin-left: 10px;
  }
}

.desc-text {
  text-align: left;
  padding-top: 10px;
}

@include media-breakpoint-up(sm) {
  .flex-elements:nth-child(even) {
    padding: 10px 35px;
  }
  .details {
    padding: 40px;
  }
}

@include media-breakpoint-up(md) {
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

  .flex-elements:nth-child(even),
  .details {
    padding: 9px;
  }

  .details {
    height: auto;
  }

  .title,
  .desc-text {
    font-size: 1rem;
    line-height: 1.2rem;
  }
}

@include media-breakpoint-up(lg) {
  .title,
  .desc-text {
    font-size: 1.2rem;
    line-height: 1.5rem;
  }

  .padding {
    padding: 20px;
  }
}

@include media-breakpoint-up(xl) {
  .title,
  .desc-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
    padding-top: 20px;
  }
}

@include media-breakpoint-up(xxl) {
  .details {
    padding: 40px;
  }
}
</style>
