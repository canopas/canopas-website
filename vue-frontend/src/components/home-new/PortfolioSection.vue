<template>
  <div class="container-fluid mt-5">
    <div class="text-center v2-header-text">Portfolio</div>
    <div class="portfolio-section mt-5">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        class="portfolio-details mt-4"
      >
        <img
          :src="portfolio.image"
          class="portfolio-image"
          loading="lazy"
          :alt="portfolio.title + `-image`"
        />
        <div class="description">
          <div class="v2-title-text">{{ portfolio.title }}</div>
          <div class="v2-normal-text mt-3">
            {{ portfolio.detail }}
          </div>
          <div class="mt-4">
            <router-link class="v2-button v2-normal-2-text" to="">
              <span>VIEW</span>
              <font-awesome-icon
                class="arrow fa"
                icon="arrow-right"
                id="leftArrow"
              />
            </router-link>
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

export default {
  data() {
    return {
      luxeradio: luxeradio,
      noLonely: noLonely,
      portfolios: [
        {
          image: luxeradio,
          title: "Luxeradio",
          detail:
            "Luxe Radio is a music streaming and podcast app similar to Spotify. The android app has 150K+ downloads and 5 STAR APP RATINGS.",
          link: "",
        },
        {
          image: smilePlue,
          title: "Smile+",
          detail:
            "Smile+ app is designed for dentists to create a perfect smile for their patients. Using the Smile+ app, dentists can get their patients the best smile simulation in a minute, automated with AI.",
          link: "",
        },
        {
          image: noLonely,
          title: "NoLonely",
          detail:
            "NoLonely is a start-up with a strong vision for overcoming loneliness, depression, and mental health-related issues for humanity.",
          link: "",
        },
        {
          image: togness,
          title: "Togness",
          detail:
            "Togness is a photo editor and slideshow maker app for your life's most memorable events like weddings, pets, friends & family, and memorials, etc.",
          link: "",
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

  .portfolio-image {
    height: 50%;
    width: 50%;
  }
}

@include media-breakpoint-up(lg) {
  .portfolio-image {
    height: 40%;
    width: 40%;
  }
}
</style>
