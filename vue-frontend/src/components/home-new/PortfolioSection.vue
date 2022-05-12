<template>
  <div class="container-fluid mt-5">
    <div class="text-center v2-header-text">Portfolio</div>
    <div class="portfolio-section">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        class="portfolio-details mt-4"
      >
        <div class="image-div">
          <img
            :src="portfolio.image"
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
          <div class="mt-4">
            <a
              class="v2-normal-2-text v2-button"
              :href="portfolio.link"
              target="_blank"
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
import luxeradio from "@/assets/images/portfolio/luxeradio.png";
import noLonely from "@/assets/images/portfolio/noLonely.png";
import smilePlue from "@/assets/images/portfolio/smilePlus.png";
import togness from "@/assets/images/portfolio/togness.png";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import config from "@/config.js";

export default {
  data() {
    return {
      luxeradio: luxeradio,
      noLonely: noLonely,
      portfolios: [
        {
          image: luxeradio,
          title: "Luxeradio",
          aspectRatio: "47%",
          detail:
            "Luxe Radio is a music streaming and podcast app similar to Spotify. The android app has 150K+ downloads and 5 STAR APP RATINGS.",
          link: config.LUXERADIO_URL,
        },
        {
          image: smilePlue,
          title: "Smile+",
          aspectRatio: "47%",
          detail:
            "Smile+ app is designed for dentists to create a perfect smile for their patients. Using the Smile+ app, dentists can get their patients the best smile simulation in a minute, automated with AI.",
          link: config.SMILEPLUS_URL,
        },
        {
          image: noLonely,
          title: "NoLonely",
          aspectRatio: "47%",
          detail:
            "NoLonely is a start-up with a strong vision for overcoming loneliness, depression, and mental health-related issues for humanity.",
          link: config.NOLONELY_URL,
        },
        {
          image: togness,
          title: "Togness",
          aspectRatio: "55%",
          detail:
            "Togness is a photo editor and slideshow maker app for your life's most memorable events like weddings, pets, friends & family, and memorials, etc.",
          link: config.TOGNESS_URL,
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  padding: 0;
}

.portfolio-section {
  margin-top: 150px;
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

.arrow {
  color: #3d3d3d;
}

@include media-breakpoint-up(md) {
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
  .image-div {
    width: 45%;
  }
}
</style>
