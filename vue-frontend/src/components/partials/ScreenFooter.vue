<template>
  <footer class="container bottom-footer-div">
    <div class="logo-div">
      <img width="180" :src="footerLogoImage" alt="logo" />
      <p class="text-left">
        <font-awesome-icon :icon="copyrightIcon" />
        {{new Date().getFullYear()}} Canopas Software LLP. All rights
        reserved.
      </p>
    </div>
    <div class="social-media-icons">
      <li v-for="(socialMedia, index) in socialMediaIcons" :key="socialMedia">
        <a :href="socialMedia.url" target="_blank">
          <font-awesome-icon
            class="fab"
            :icon="socialMedia.icon"
            @mouseover="hoverIndex = index"
            @mouseleave="hoverIndex = null"
            :class="{ 'fa-icon': hoverIndex == index }"
          />
        </a>
      </li>
    </div>
  </footer>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faMedium,
  faFacebook,
  faTwitter,
  faInstagram,
  faGithub,
  faLinkedin,
} from "@fortawesome/free-brands-svg-icons";
import { faCopyright } from "@fortawesome/free-regular-svg-icons";

import footerLogoImage from "@/assets/images/logo/logo-footer.svg";
import Config from "@/config.js";

export default {
  data() {
    return {
      copyrightIcon: faCopyright,
      footerLogoImage: footerLogoImage,
      hoverIndex: null,
      socialMediaIcons: [
        {
          url: Config.MEDIUM_URL,
          icon: faMedium,
        },
        {
          url: Config.FACEBOOK_URL,
          icon: faFacebook,
        },
        {
          url: Config.TWITTER_URL,
          icon: faTwitter,
        },
        {
          url: Config.INSTAGRAM_URL,
          icon: faInstagram,
        },
        {
          url: Config.GITHUB_URL,
          icon: faGithub,
        },
        {
          url: Config.LINKEDIN_URL,
          icon: faLinkedin,
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
.fa-icon {
  color: #f2709c !important;
}

.bottom-footer-div {
  margin: 0 auto;
  padding: 3% 0;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  justify-content: space-between;
  color: rgba(61, 61, 61, 0.3);
}

.social-media-icons li {
  margin: 0.6rem 0.5rem 0.6rem 0;
  display: inline-flex;
}

.social-media-icons li:hover {
  transform: scale(1.3);
}

.logo-div > p {
  font-size: 1rem;
  line-height: 1.3rem;
  margin-top: 0.5rem;
}

.fab {
  color: rgba(61, 61, 61, 0.25);
  font-size: 1.2rem;
}

@include media-breakpoint-up(md) {
  .bottom-footer-div {
    align-items: center;
    flex-direction: row;
  }

  .social-media-icons li {
    margin: 0.6rem 0.5rem 0.6rem;
  }
}

@include media-breakpoint-up(lg) {
  .fab {
    font-size: 1.5rem;
  }
}
</style>
