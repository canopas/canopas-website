<template>
  <footer class="bottom-footer-div">
    <div class="logo-div">
      <img width="180" :src="footerLogoImage" alt="logo" />
      <p class="text-left">
        <font-awesome-icon :icon="copyrightIcon" />
        <span id="currentYear"></span> Canopas Software LLP. All rights
        reserved.
      </p>
    </div>
    <div class="social-media-icons">
      <li>
        <a href="<%= process.env.BLOGS_URL %>" target="_blank">
          <font-awesome-icon class="fab" :icon="mediumIcon" />
        </a>
      </li>
      <li>
        <a href="<%= process.env.FACEBOOK_URL %>" target="_blank">
          <font-awesome-icon class="fab" :icon="facebookIcon" />
        </a>
      </li>
      <li>
        <a href="<%= process.env.TWITTER_URL %>" target="_blank">
          <font-awesome-icon class="fab" :icon="twitterIcon" />
        </a>
      </li>
      <li>
        <a href="<%= process.env.INSTAGRAM_URL %>" target="_blank">
          <font-awesome-icon class="fab" :icon="instagramIcon" />
        </a>
      </li>
      <li>
        <a href="<%= process.env.GITHUB_URL %>" target="_blank">
          <font-awesome-icon class="fab" :icon="githubIcon" />
        </a>
      </li>
      <li>
        <a href="<%= process.env.LINKEDIN_URL %>" target="_blank">
          <font-awesome-icon class="fab" :icon="linkedInIcon" />
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

export default {
  data() {
    return {
      copyrightIcon: faCopyright,
      mediumIcon: faMedium,
      facebookIcon: faFacebook,
      twitterIcon: faTwitter,
      instagramIcon: faInstagram,
      githubIcon: faGithub,
      linkedInIcon: faLinkedin,
      footerLogoImage: footerLogoImage,
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style scoped>
.bottom-footer-div {
  margin: 0 15%;
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

@media (min-width: 768px) {
  .bottom-footer-div {
    align-items: center;
    flex-direction: row;
  }

  .social-media-icons li {
    margin: 0.6rem 0.5rem 0.6rem;
  }
}

@media (min-width: 992px) {
  .fab {
    font-size: 1.5rem;
  }
}
</style>
