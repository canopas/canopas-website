<template>
  <div class="container-fluid">
    <div class="container">
      <div class="v2-header-3-text text-center">
        With <span class="v2-canopas-gradient-text">canopas</span>, you will
        have everything you need under one roof.
      </div>
      <div class="phases-container">
        <div class="flex-container">
          <div v-for="(phase, i) in phases.slice(0, 3)" :key="i" class="phases">
            <aspect-ratio height="56.25%">
              <img :src="phase.image" class="image" loading="lazy" />
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
            <aspect-ratio height="56.25%">
              <img
                :src="phase.image"
                class="image"
                :class="phase.title == 'Marketing' ? 'animation-height' : ''"
                loading="lazy"
              />
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
import planning from "@/assets/images/phases/canopas_phases_planning.gif";
import designing from "@/assets/images/phases/canopas_phases_designing.gif";
import development from "@/assets/images/phases/canopas_phases_development.gif";
import marketing from "@/assets/images/phases/canopas_phases_marketing.gif";
import support from "@/assets/images/phases/canopas_phases_support.gif";

export default {
  data() {
    return {
      phases: [
        {
          image: planning,
          title: "Planning",
          description:
            "We research to understand the vision. The audience. The product. The goals. We collaborate, we re-examine, we ask and conclude.",
        },
        {
          image: designing,
          title: "Design",
          description:
            "Working passionately with you, our designers create aesthetically pleasing  and user friendly digital experiences.",
        },
        {
          image: development,
          title: "Development",
          description:
            "Backend and Frontend. From the beginning, we have performance and maintainability in mind. Always insisting on best practices and keeping up with the latest trends.",
        },
        {
          image: marketing,
          title: "Marketing",
          description:
            "From organic growth to writing irresistible ads for your marketing campaign, we're here to help you. We'll help you build a brand and community of prospects who can give you a boost on your product launch.",
        },
        {
          image: support,
          title: "Support",
          description:
            "From bug fixing to feature updates, you can count on us. For the last 7 years, our clients had a 99.87% uptime for their products.",
        },
      ],
    };
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
}

.container {
  padding-bottom: 16px;
  padding-top: 48px;
}

.phases-container {
  margin-top: 32px;
}

.phases {
  padding: 1rem 0;
  text-align: center;
}

.v2-header-3-text {
  margin: auto;
}

.flex-container {
  display: flex;
  flex-flow: wrap;
  margin: auto;
}

.animation-height {
  height: 100% !important;
}

.image {
  width: 100%;
  height: 100%;
  object-fit: cover;
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
    width: 85%;
  }

  .details {
    padding: 1rem 3rem;
  }
}

@include media-breakpoint-up(md) {
  .container {
    padding-bottom: 48px;
    padding-top: 64px;
  }

  .phases-container {
    margin-top: 80px;
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

  .animation-height {
    height: 84% !important;
  }
}

@include media-breakpoint-up(lg) {
  .container {
    padding-bottom: 80px;
    padding-top: 128px;
  }

  .phases-container {
    margin-top: 160px;
  }

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

  .container-fluid {
    margin-bottom: 100px;
  }
}
@supports (-webkit-touch-callout: none) {
  .v2-header-3-text {
    letter-spacing: -2px;
  }
}
</style>
