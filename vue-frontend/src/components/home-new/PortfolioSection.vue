<template>
  <div class="container-fluid">
    <div class="text-center v2-header-text">Portfolio</div>
    <div class="portfolio-section">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        class="portfolio-details mt-4"
      >
        <div class="image-div">
          <img
            :src="portfolio.images[3]"
            :srcset="`${portfolio.images[0]} 400w, ${portfolio.images[1]} 800w, ${portfolio.images[2]} 1200w, ${portfolio.images[3]} 1600w`"
            sizes="(min-width: 992px) 45vw, 100vw"
            class="portfolio-image"
            loading="lazy"
            :alt="portfolio.title + `-image`"
          />
        </div>

        <div class="description">
          <div class="v2-title-text">{{ portfolio.title }}</div>
          <div class="v2-normal-text mt-3">
            {{ portfolio.detail }}
          </div>
          <div class="mt-4 d-flex">
            <a
              class="v2-normal-2-text v2-button"
              :href="portfolio.link"
              :target="portfolio.target"
            >
              <span>VIEW</span>
              <font-awesome-icon
                class="arrow fa"
                icon="arrow-right"
                id="leftArrow"
              />
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import luxeradio400w from "@/assets/images/portfolio/luxeradio-400w.webp";
import luxeradio800w from "@/assets/images/portfolio/luxeradio-800w.webp";
import luxeradio1200w from "@/assets/images/portfolio/luxeradio-1200w.webp";
import luxeradio16000w from "@/assets/images/portfolio/luxeradio-1600w.webp";

import smilep400w from "@/assets/images/portfolio/smilep-400w.webp";
import smilep800w from "@/assets/images/portfolio/smilep-800w.webp";
import smilep1200w from "@/assets/images/portfolio/smilep-1200w.webp";
import smilep16000w from "@/assets/images/portfolio/smilep-1600w.webp";

import nolonely400w from "@/assets/images/portfolio/nolonely-400w.webp";
import nolonely800w from "@/assets/images/portfolio/nolonely-800w.webp";
import nolonely1200w from "@/assets/images/portfolio/nolonely-1200w.webp";
import nolonely16000w from "@/assets/images/portfolio/nolonely-1600w.webp";

import togness400w from "@/assets/images/portfolio/togness-400w.webp";
import togness800w from "@/assets/images/portfolio/togness-800w.webp";
import togness1200w from "@/assets/images/portfolio/togness-1200w.webp";
import togness16000w from "@/assets/images/portfolio/togness-1600w.webp";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import config from "@/config.js";

export default {
  data() {
    return {
      portfolios: [
        {
          images: [
            luxeradio400w,
            luxeradio800w,
            luxeradio1200w,
            luxeradio16000w,
          ],
          title: "Luxeradio",
          detail:
            "Luxe Radio, the radio of taste, elegance, and refinement, intends to be the showcase of excellence and the best of Moroccan and international creation.",
          link: "portfolio/luxeradio",

          target: "",
        },
        {
          images: [togness400w, togness800w, togness1200w, togness16000w],
          title: "Togness",
          aspectRatio: "55%",
          detail:
            "Togness is a photo editor and slideshow maker app for your life's most memorable events like weddings, pets, friends & family, and memorials, etc.",
          link: "portfolio/togness",

          target: "",
        },
        {
          images: [nolonely400w, nolonely800w, nolonely1200w, nolonely16000w],
          title: "Nolonely",
          detail:
            "Nolonely is a start-up with a strong vision for overcoming loneliness, depression, and mental health-related issues for humanity.",
          link: "portfolio/nolonely",

          target: "",
        },
        {
          images: [smilep400w, smilep800w, smilep1200w, smilep16000w],
          title: "Smile+",
          detail:
            "Smile+ app is designed for dentists to create a perfect smile for their patients. Using the Smile+ app, dentists can get their patients the best smile simulation in a minute, automated with AI.",
          link: config.IS_SHOW_PORTFOLIO_PAGES
            ? "portfolio/smileplus"
            : config.SMILEPLUS_URL,
          target: config.IS_SHOW_PORTFOLIO_PAGES ? "" : "_blank",
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {
    if (window.innerWidth < 768) {
      this.portfolios[0].image = mobileLuxeradio;
    }
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  padding: 0;
  margin-bottom: 30px;
  margin-top: 50px;
}

.v2-header-text {
  font-family: "FuturaLT-Bold";
}

.portfolio-section {
  margin: 50px 0;
}

.image-div {
  width: 100%;
}

.portfolio-image {
  height: 100%;
  width: 100%;
  object-fit: cover;
}

.portfolio-details {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.description {
  padding: 5% 20%;
}

.v2-button {
  display: flex;
  align-items: center;
}

.v2-button > span {
  margin-right: 10px;
}

.arrow {
  color: #3d3d3d;
}

@include media-breakpoint-up(md) {
  .container-fluid {
    margin-top: 90px;
  }

  .portfolio-section {
    margin: 80px 0;
  }

  .portfolio-details:nth-child(even) {
    flex-direction: row-reverse;
  }

  .portfolio-details:nth-child(odd) {
    flex-direction: row;
  }

  .description {
    width: 35%;
    margin: auto;
    padding: 0;
  }

  .image-div {
    width: 50%;
  }
}

@include media-breakpoint-up(lg) {
  .container-fluid {
    margin-top: 150px;
  }

  .portfolio-section {
    margin: 150px 0;
  }

  .image-div {
    width: 45%;
  }
}
</style>
