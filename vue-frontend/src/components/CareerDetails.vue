<template>
  <div class="container-fluid">
    <link
      rel="stylesheet"
      href="https://use.fontawesome.com/releases/v5.2.0/css/all.css"
    />
    <ScreenHeader />
    <div v-if="isLoading" class="loader-div">
      <img :src="loader" />
    </div>
    <div v-else-if="showErrorMessagePopup">
      <transition name="modal">
        <div class="modal-mask">
          <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-body">
                <div class="error-message-div">
                  <div class="error-message-text text-center">
                    Something went wrong on our side
                  </div>
                  <div class="close-btn-div">
                    <button
                      type="submit"
                      class="gradient-btn close-btn"
                      @click.prevent="closeErrorMessageModal()"
                    >
                      <span>Close</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
    <div v-else>
      <div class="container">
        <div class="canopas-gradient-text text-center">
          {{ this.details.title }}
        </div>
        <hr class="title-hr mt-4" />
        <div class="normal-text summary-text mt-5">
          {{ this.details.summary }}
        </div>
        <div class="mt-5">
          <div id="description" v-html="description"></div>
        </div>
        <div class="text-center mt-5">
          <button class="gradient-btn">
            <font-awesome-icon
              class="fa icon"
              :icon="checkCircle"
              aria-hidden="true"
            />
            <span>Apply Now</span>
          </button>
        </div>
      </div>
      <ScreenFooter v-if="!showJobs" />
      <ScreenFooter2 v-if="showJobs" />
    </div>
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import axios from "axios";
import config from "@/config.js";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import loader from "@/assets/images/theme/loader.svg";
import router from "@/router";
import { library } from "@fortawesome/fontawesome-svg-core";
import Config from "@/config.js";

import {
  faCheckCircle,
  faChevronRight,
} from "@fortawesome/free-solid-svg-icons";

library.add(faCheckCircle, faChevronRight);

export default {
  data() {
    return {
      details: "",
      description: "",
      checkCircle: faCheckCircle,
      loader: loader,
      isLoading: true,
      showErrorMessagePopup: false,
      showJobs: Config.IS_SHOW_JOBS,
    };
  },
  components: {
    ScreenHeader,
    ScreenFooter,
    ScreenFooter2,
    FontAwesomeIcon,
  },
  mounted() {
    this.getCareerDetails();
  },
  methods: {
    getCareerDetails() {
      var id = this.$route.params.id;
      axios
        .get(config.API_BASE + "/api/careers/" + id)
        .then((res) => {
          this.isLoading = false;
          this.details = res.data;
          this.description = this.details.description;
        })
        .catch((err) => {
          this.isLoading = false;
          if (err.response && err.response.status == 404) {
            router.push({
              name: "Error404Page",
              params: { catchAll: "jobs/" + id },
            });
          } else {
            this.showErrorMessagePopup = true;
          }
        });
    },
    closeErrorMessageModal() {
      router.push({
        path: `/jobs`,
      });
    },
  },
  updated() {
    let descriptionTitles = document.querySelectorAll("#description > h2");
    descriptionTitles.forEach((descriptionTitle) => {
      let bullet = document.createElement("span");
      bullet.className = "bullet";
      descriptionTitle.prepend(bullet);
    });
  },
};
</script>

<style lang="scss" scoped>
.loader-div {
  display: flex;
  justify-content: center;
}

.container {
  padding: 3rem 1rem;
}

.title-hr {
  border: 1px solid #e2e2e2;
  margin: 0;
}

.canopas-gradient-text {
  font-weight: 700;
  font-size: 1.75rem !important;
  line-height: 2rem !important;
  letter-spacing: 0.1rem;
}

.gradient-btn {
  padding: 1rem 3rem;
}

.gradient-btn > span {
  letter-spacing: 0.06rem;
}

.icon {
  font-size: 1.1rem;
}

.summary-text,
:deep(div > span *) {
  font-size: 1.1rem !important;
  line-height: 2rem;
  text-align: justify;
}

:deep(h2) {
  display: flex;
  flex-direction: row;
  margin: 3rem 0 1rem;
}

:deep(h2 *) {
  font-weight: bolder;
  font-size: 1.5rem !important;
  line-height: 2rem !important;
  letter-spacing: 0.05rem;
}

:deep(h2 > .bullet) {
  content: "";
  background: linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  display: inline-block;
  line-height: 1;
  margin: -0.3rem 1.2rem -0.3rem 0;
  height: auto;
  width: 0.5rem;
  border-radius: 0;
}

// error message
.modal-mask {
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.error-message-div {
  padding: 20px;
}

.error-message-text {
  line-height: 2rem;
  font-size: 1.5rem;
  margin-bottom: 30px;
}

.close-btn-div {
  width: 100%;
}

.close-btn {
  float: right;
}

@include media-breakpoint-up(sm) {
  .icon {
    font-size: 1.3rem;
  }
}

@include media-breakpoint-up(md) {
  .container {
    padding: 6rem;
  }

  .canopas-gradient-text {
    font-size: 2.25rem !important;
    line-height: 2.729rem !important;
  }
  .gradient-btn {
    padding: 1rem 5rem;
  }

  .gradient-btn > span {
    font-size: 1.125rem;
  }

  .summary-text,
  :deep(div > span *) {
    font-size: 1.125rem !important;
    line-height: 2.5rem !important;
  }

  :deep(h2 *) {
    font-size: 2rem !important;
    line-height: 2.5rem !important;
  }
}

@include media-breakpoint-up(lg) {
  .summary-text,
  :deep(div > span *) {
    font-size: 1.4rem !important;
  }
}
</style>
