<template>
  <div class="container">
    <div class="text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Find t</span>he career of your dreams
      </div>
      <div class="description normal-text">
        <div class="text-weight">
          When was the last time you did something for the first time?
        </div>
        <div class="mt-5">
          At Canopas, we frequently do new things that help you leave your
          comfort zone. And when you leave your comfort zone, you will know what
          you are capable of and you will never be the same again. If you
          believe your work should have meaning and want to make an impact in
          the world with your work, come join us.
        </div>
      </div>
    </div>

    <div v-if="isLoading" class="loader-div">
      <img :src="loader" />
    </div>
    <div v-else class="career-container normal-text mt-5">
      <div class="career-list-container">
        <div v-for="career in careers" :key="career" class="career-list">
          <div
            class="career-header normal-text text-left"
            @click="expandListItem(career.id)"
          >
            <div class="career-icon">
              <font-awesome-icon
                class="gradient-icon icon"
                :icon="career.icon_name"
              />
            </div>
            <div class="career-title">{{ career.title }}</div>
            <font-awesome-icon
              class="career-plus-icon"
              :icon="openList && career.id == currentIndex ? 'minus' : 'plus'"
            />
          </div>
          <collapse-transition>
            <div
              v-if="openList && career.id == currentIndex"
              :key="career.summary"
            >
              <div class="career-summary">
                {{ career.summary }}
              </div>

              <div class="read-apply-btns">
                <router-link
                  class="gradient-border-btn"
                  :to="career.detail_link"
                >
                  <font-awesome-icon
                    class="fa gradient-icon"
                    icon="align-left"
                    flip="vertical"
                    aria-hidden="true"
                  />
                  <span>Read More</span>
                </router-link>
                <router-link class="gradient-btn" :to="career.job_link">
                  <font-awesome-icon
                    class="fa"
                    icon="check-circle"
                    aria-hidden="true"
                  />
                  <span>Apply Now</span>
                </router-link>
              </div>
            </div>
          </collapse-transition>
          <div class="career-divider" v-if="career.id != careers.length"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import axios from "axios";
import config from "./../../config.js";
import loader from "./../../assets/images/theme/loader.svg";
import router from "./../../router";

import { faApple, faAndroid } from "@fortawesome/free-brands-svg-icons";

import {
  faGlobe,
  faBullhorn,
  faPenNib,
  faUser,
} from "@fortawesome/free-solid-svg-icons";

export default {
  data() {
    return {
      careers: [],
      icon: [faApple, faAndroid, faGlobe, faBullhorn, faPenNib, faUser],
      currentIndex: 0,
      openList: true,
      previousIndex: 0,
      loader: loader,
      isLoading: true,
    };
  },
  components: {
    CollapseTransition,
    FontAwesomeIcon,
  },
  mounted() {
    this.getCareers();
  },
  methods: {
    expandListItem(index) {
      if (this.previousIndex == index && this.openList) {
        this.openList = false;
      } else {
        this.openList = true;
      }

      this.currentIndex = index;
      this.previousIndex = this.currentIndex;
    },
    getCareers() {
      axios
        .get(config.API_BASE + "/api/careers")
        .then((response) => {
          this.isLoading = false;
          this.careers = response.data;
          for (let i = 0; i < this.careers.length; i++) {
            var unique_id = this.careers[i].unique_id;
            this.careers[i].detail_link = "/jobs/" + unique_id;
            this.careers[i].job_link = "/jobs/apply/" + unique_id;
            this.careers[i].icon_name = this.icon[i];
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getCareerDetails(id) {
      router.push({
        path: `/jobs/${id}`,
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin: 100px auto;
}

.description {
  margin: 30px 2%;
}

.text-weight {
  font-weight: 700;
}

.career-container {
  border-radius: 14px;
  border: 1px solid #e2e2e2;
  overflow: hidden;
  margin: 0 2%;
}

.career-list-container {
  display: flex;
  flex-direction: column;
}

.career-list {
  display: flex;
  flex-direction: column;
}

.career-divider {
  background: #e2e2e2;
  height: 1px;
}

.career-header {
  display: flex;
  flex-direction: row;
  width: 100%;
  align-items: center;
  padding: 30px;
  cursor: pointer;
}

.career-icon {
  width: 50px;
}

.career-summary {
  padding: 30px;
  color: rgba(61, 61, 61, 0.8);
}

.career-title {
  font-style: normal;
  font-weight: 700;
  color: #3d3d3d;
}

.career-plus-icon {
  margin-left: auto;
  order: 2;
  color: rgba(61, 61, 61, 0.8);
}

.read-apply-btns {
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 0 0 30px;
}

.gradient-btn {
  margin: 5px;
}

.gradient-btn > span {
  font-size: 1.125rem;
  line-height: 1.364rem;
}

.gradient-icon {
  color: #f2709c;
}

.gradient-border-btn,
.gradient-btn {
  display: flex;
  align-items: center;
  padding: 11px;
}

.gradient-border-btn > span,
.gradient-btn > span {
  font-size: 0.9rem;
  font-weight: 700;
  letter-spacing: 0.06rem;
}

.loader-div {
  display: flex;
  justify-content: center;
}

.icon {
  font-size: 1.575rem;
}

@include media-breakpoint-up(sm) {
  .read-apply-btns {
    justify-content: flex-end;
    padding: 0 20px 30px;
  }

  .icon {
    font-size: 1.875rem;
  }

  .gradient-border-btn,
  .gradient-btn {
    padding: 16px;
  }

  .career-icon {
    width: 60px;
  }
}

@include media-breakpoint-up(md) {
  .career-header {
    padding: 32px 48px;
  }

  .career-summary {
    padding: 32px 48px;
  }

  .description {
    margin: 50px 80px;
  }

  .career-container {
    margin: 0 6%;
  }

  .gradient-border-btn > span,
  .gradient-btn > span {
    font-size: 1rem;
  }
}

@include media-breakpoint-up(lg) {
  .career-header {
    padding: 40px 48px;
  }

  .career-summary {
    padding: 40px 48px;
    line-height: 2rem;
  }

  .gradient-border-btn > span,
  .gradient-btn > span {
    font-size: 1.125rem;
  }
}
</style>
