<template>
  <div class="container-fluid">
    <img
      src="@/assets/images/footer/Canopas-footer-img.svg"
      class="background"
      alt="footer-background-image"
    />
    <div class="container text-center">
      <div class="container-flex" v-if="isJobsUrl">
        <a :href="glassdoorLink">
          <img
            src="@/assets/images/footer/glassdoor.webp"
            class="glassdoor-img tw-h-full"
            alt="glassdoor-image"
          />
        </a>
      </div>
      <div class="container-flex">
        <div class="normal-text mb-24">Follow us on</div>
        <ul class="social-media-icons">
          <li
            v-for="socialMedia in socialMediaIcons"
            :key="socialMedia"
            :class="socialMedia.image ? 'design-rush-icon' : 'gradient-btn'"
          >
            <a :href="socialMedia.url" target="_blank">
              <img
                v-if="socialMedia.image"
                :src="hover ? designRushHover : designRush"
                @mouseover="hover = true"
                @mouseleave="hover = false"
              />
              <font-awesome-icon v-else class="fab" :icon="socialMedia.icon" />
            </a>
          </li>
        </ul>
      </div>
      <div class="container-flex">
        <p class="normal-2-text">
          <font-awesome-icon :icon="copyrightIcon" /> &nbsp;
          <span
            >{{ new Date().getFullYear() }} Canopas Software LLP. All rights
            reserved.</span
          >
        </p>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { faCopyright } from "@fortawesome/free-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import designRush from "@/assets/images/footer/design-rush.svg";
import designRushHover from "@/assets/images/footer/design-rush-hover.svg";

import {
  faFacebookF,
  faInstagram,
  faTwitter,
  faMediumM,
  faLinkedinIn,
  faYoutube,
} from "@fortawesome/free-brands-svg-icons";

import Config from "@/config.js";

export default {
  data() {
    return {
      glassdoorLink:
        "https://www.glassdoor.co.in/Overview/Working-at-Canopas-EI_IE3194462.11,18.htm",
      copyrightIcon: faCopyright,
      designRush: designRush,
      designRushHover: designRushHover,
      hover: false,
      socialMediaIcons: [
        {
          url: Config.FACEBOOK_URL,
          icon: faFacebookF,
        },
        {
          url: Config.INSTAGRAM_URL,
          icon: faInstagram,
        },
        {
          url: Config.TWITTER_URL,
          icon: faTwitter,
        },
        {
          url: Config.BLOG_URL,
          icon: faMediumM,
        },
        {
          url: Config.LINKEDIN_URL,
          icon: faLinkedinIn,
        },
        {
          url: Config.YOUTUBE_URL,
          icon: faYoutube,
        },
        {
          url: Config.DESIGNRUSH_URL,
          image: true,
        },
      ],
      isJobsUrl: this.$router.currentRoute.value.fullPath.indexOf("/jobs") > -1,
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  margin: 0;
  padding: 0;
  position: relative;
}

.background {
  width: 100%;
  height: 100%;
  object-fit: cover;
  position: absolute;
  top: 0;
  left: 0;
  z-index: -5;
}

.container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
}

.container-flex {
  margin-bottom: 4rem;
}

.glassdoor-img {
  width: 80%;
}

.social-media-icons {
  list-style-type: none;
  display: flex;
  justify-content: space-between;

  .design-rush-icon {
    margin: 0 0.25rem;
    width: 50px;
    height: 50px;
    cursor: pointer;
  }

  .gradient-btn {
    padding: 0.7rem;
    margin: 0 0.25rem;
    border-radius: 50%;
    width: 50px;
    height: 50px;
    cursor: pointer;

    * {
      width: 100%;
      height: 100%;
    }
  }
}

.fab {
  color: white;
}

.normal-text {
  font-size: 1.375rem;
  line-height: 1.75rem;
  font-weight: bold;
}

.mb-24 {
  margin-bottom: 1.5rem;
}

@include media-breakpoint-up(sm) {
  .social-media-icons > li {
    margin: 1.25rem 0.5rem 0;
    width: 60px;
    height: 60px;
  }

  .glassdoor-img {
    width: 100%;
  }
}

@include media-breakpoint-up(lg) {
  .normal-text {
    font-size: 1.75rem;
    line-height: 2.125rem;
  }
}

@include media-breakpoint-up(xl) {
  .container-flex {
    margin-bottom: 6rem;
  }
}
</style>
