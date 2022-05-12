<template>
  <div class="container">
    <div class="pb-5">
      <div class="border-text v2-header-text">We are</div>
      <div class="border-text v2-header-text">here to</div>
      <div class="v2-title-3-text">
        <span class="v2-header-text pe-2">contribute</span>
        and that starts with the site itself, canopas is open ource! We also
        have several projects that are widely recognized in the open source
        community.
      </div>
    </div>
    <div class="flex-div">
      <div
        v-for="contribution in contributions"
        :key="contribution"
        class="contribute-info"
        :style="{ background: contribution.bgColor }"
      >
        <img :src="contribution.image" class="image" loading="lazy" />

        <div class="details">
          <a class="github-info" :href="contribution.link" target="_blank">
            <font-awesome-icon class="github" :icon="icon" />
            <div class="v2-title-3-text v2-canopas-gradient-text">GitHub</div>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import introShowcase from "@/assets/images/contribution/intro-showcase.webp";
import uiPilot from "@/assets/images/contribution/ui-pilot.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faGithub } from "@fortawesome/free-brands-svg-icons";

library.add(faGithub);

export default {
  data() {
    return {
      icon: faGithub,
      contributions: [
        {
          image: introShowcase,
          bgColor: "#FAF5F2",
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          image: uiPilot,
          bgColor: "#3D3D3D",
          link: "https://github.com/canopas/UIPilot",
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
.container {
  padding: 100px 0 40px;
}

.border-text {
  -webkit-text-stroke: 1px #3d3d3d;
  color: #fff;
}

.flex-div {
  display: flex;
  flex-direction: column;
}

.contribute-info {
  border-radius: 10px;
  margin: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
}

.image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 10px 10px 0 0;
}

.details {
  width: 100%;
  padding: 20px;
}

.github-info {
  float: right;
  display: flex;
  align-items: flex-end;
}

@media (hover: hover) and (pointer: fine) {
  .github-info {
    transition: 0.3s ease-in-out;
  }
  .github-info:hover {
    transform: scale(1.1);
  }
}

.github {
  height: 30px;
  width: 30px;
  color: #f2709c;
  margin-right: 10px;
}

.v2-canopas-gradient-text {
  font-weight: 600;
}

@include media-breakpoint-up(md) {
  .flex-div {
    flex-flow: row;
  }
}

@include media-breakpoint-up(lg) {
  .github {
    height: 40px;
    width: 40px;
  }
}
</style>
