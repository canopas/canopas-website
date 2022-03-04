<template>
  <div class="container-fluid">
    <img :src="backgroundFooterImage" class="background" />
    <div class="container text-center">
      <div class="container-flex" v-if="isJobsUrl">
        <a :href="glassdoorLink">
          <img :src="glassdoorImage" class="glassdoor-img" />
        </a>
      </div>
      <div class="container-flex">
        <div class="jobs-normal-text mb-24">Follow us on</div>
        <div class="social-media-icons">
          <li
            v-for="socialMedia in socialMediaIcons"
            :key="socialMedia"
            class="gradient-btn"
          >
            <a :href="socialMedia.url" target="_blank">
              <font-awesome-icon class="fab" :icon="socialMedia.icon" />
            </a>
          </li>
        </div>
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

import {
  faFacebookF,
  faInstagram,
  faTwitter,
  faMediumM,
  faLinkedinIn,
} from "@fortawesome/free-brands-svg-icons";

import backgroundFooterImage from "@/assets/images/footer/Canopas-footer-img.svg";
import glassdoorImage from "@/assets/images/footer/glassdoor.webp";
import Config from "@/config.js";
import router from "@/router";

export default {
  data() {
    return {
      backgroundFooterImage: backgroundFooterImage,
      glassdoorImage: glassdoorImage,
      glassdoorLink:
        "https://www.glassdoor.co.in/Overview/Working-at-Canopas-EI_IE3194462.11,18.htm",
      copyrightIcon: faCopyright,
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
      ],
      isJobsUrl: router.currentRoute.value.fullPath.indexOf("/jobs") > -1,
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

.jobs-normal-text {
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
  .jobs-normal-text {
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
