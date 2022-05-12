<template>
  <div class="container-fluid">
    <div class="container">
      <div class="v2-header-3-text text-center">
        With <span class="pink-text">canopas</span>, you will have everything
        you need under one roof.
      </div>
      <div class="phases-container">
        <div class="flex-container">
          <div v-for="(phase, i) in phases.slice(0, 3)" :key="i" class="phases">
            <aspect-ratio height="60%">
              <LottieAnimation :jsonData="phase.jsonData" />
            </aspect-ratio>
            <div class="details">
              <div class="v2-title-2-text canopas-gradient-text">
                {{ phase.title }}
              </div>
              <div class="v2-normal-2-text description">
                {{ phase.description }}
              </div>
            </div>
          </div>
        </div>
        <div class="flex-container">
          <div v-for="(phase, i) in phases.slice(3, 5)" :key="i" class="phases">
            <aspect-ratio
              height="60%"
              :class="phase.title == 'Marketing' ? 'lottie-animation' : ''"
            >
              <LottieAnimation :jsonData="phase.jsonData" />
            </aspect-ratio>
            <div class="details">
              <div class="v2-title-2-text canopas-gradient-text">
                {{ phase.title }}
              </div>
              <div class="v2-normal-2-text description">
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
import LottieAnimation from "@/components/utils/LottieAnimation.vue";
import AspectRatio from "@/components/utils/AspectRatio.vue";
import planning from "@/assets/lottie/phases/canopas_phases_planning.json";
import designing from "@/assets/lottie/phases/canopas_phases_designing.json";
import development from "@/assets/lottie/phases/canopas_phases_development.json";
import marketing from "@/assets/lottie/phases/canopas_phases_marketing.json";
import support from "@/assets/lottie/phases/canopas_phases_support.json";

export default {
  data() {
    return {
      phases: [
        {
          jsonData: planning,
          title: "Planning",
          description:
            "We research to understand the vision. The audience. The product. The goals. We collaborate, we re-examine, we ask and conclude.",
        },
        {
          jsonData: designing,
          title: "Design",
          description:
            "Working passionately with you, our designers create aesthetically pleasing  and user friendly digital experiences. ",
        },
        {
          jsonData: development,
          title: "Development",
          description:
            "Backend and Frontend. From the beginning, we have performance and maintainability in mind. Always insisting on best practices and keeping up with the latest trends.",
        },
        {
          jsonData: marketing,
          title: "Marketing",
          description:
            "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
        },
        {
          jsonData: support,
          title: "Support",
          description:
            "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. ",
        },
      ],
    };
  },
  components: {
    LottieAnimation,
    AspectRatio,
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  background: linear-gradient(
    174.39deg,
    #ffffff 9.06%,
    rgba(255, 148, 114, 0.15) 90%
  );
  margin-bottom: 200px;
}

.container {
  padding: 5rem 0rem 8rem;
}

.phases-container {
  margin-top: 5rem;
}

.phases {
  padding: 1rem 0;
  text-align: center;
}

.pink-text {
  color: #f2709c;
}

.v2-header-3-text {
  margin: auto;
}

.flex-container {
  display: flex;
  flex-flow: wrap;
  margin: auto;
}

.lottie-animation {
  width: 70%;
  margin: auto;
}

.details {
  padding: 1rem;

  div {
    padding: 0.5rem;
  }
}

@media (min-width: 440px) {
  .details > .description {
    padding: 0 3rem;
  }
}

@include media-breakpoint-up(sm) {
  .v2-header-3-text {
    width: 80%;
  }

  .details {
    padding: 1rem 3rem;
  }
}

@include media-breakpoint-up(md) {
  .phases-container {
    margin-top: 10rem;
  }

  .phases {
    flex: 33%;
    display: flex;
    flex-direction: column;
    padding: 0;
    text-align: left;
  }

  .flex-container {
    margin: 5rem auto auto;
    :nth-child(3) {
      padding: 0 10rem;
    }
  }

  .details {
    padding: 1rem 3rem 1rem 1.5rem;
    .description {
      padding: 0.5rem;
    }
  }
}

@include media-breakpoint-up(lg) {
  .flex-container:nth-of-type(even) {
    width: 70%;
  }

  .flex-container {
    :nth-child(3) {
      padding: 0;
    }
  }

  .details {
    padding: 1rem 1.5rem;
  }
}
</style>
