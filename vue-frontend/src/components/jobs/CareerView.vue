<template>
  <div class="tw-container tw-my-[14.063rem] tw-mx-auto">
    <div class="tw-text-center">
      <p class="header-text canopas-gradient-text">
        <span class="underline-text">Find t</span>he career of your dreams
      </p>
      <div
        class="tw-my-[30px] tw-mx-[2%] md:tw-my-[50px] md:tw-mx-20 normal-text"
      >
        <div class="tw-font-bold">
          When was the last time you did something for the first time?
        </div>
        <div class="tw-mt-12">
          At Canopas, we frequently do new things that help you leave your
          comfort zone. And when you leave your comfort zone, you will know what
          you are capable of and you will never be the same again. If you
          believe your work should have meaning and want to make an impact in
          the world with your work, come join us.
        </div>
      </div>
    </div>

    <div v-if="!careers" class="tw-flex tw-justify-center">
      <img :src="loader" alt="loader-image" />
    </div>
    <div v-else>
      <div
        v-if="!careers.length"
        class="tw-text-center tw-py-3.5 tw-px-0 normal-text"
      >
        No matching jobs found
      </div>
      <div
        class="tw-rounded-[14px] tw-border tw-border-solid tw-border-[#e2e2e2] tw-overflow-hidden tw-my-0 tw-mx-[2%] tw-mt-12 md:tw-mx-[6%] normal-text"
      >
        <div v-for="career in careers" :key="career.id">
          <div
            class="tw-flex tw-flex-row tw-w-full tw-items-center tw-p-[30px] tw-cursor-pointer md:tw-py-8 md:tw-px-12 lg:tw-py-10 lg:tw-px-12 normal-text"
            @click="expandListItem(career.id)"
          >
            <div class="career-icon tw-w-[50px] sm:tw-w-[60px]">
              <font-awesome-icon
                class="tw-text-pink-300 icon tw-text-[1.575rem] sm:tw-text-[1.875rem]"
                :icon="[career.icon_name.prefix, career.icon_name.icon]"
              />
            </div>
            <div class="tw-not-italic tw-font-bold tw-text-black-900">
              {{ career.title }}
            </div>
            <font-awesome-icon
              class="career-plus-icon tw-ml-auto tw-order-2 tw-text-[#3d3d3dcc]"
              :icon="openList && currentIndex === career.id ? 'minus' : 'plus'"
            />
          </div>
          <collapse-transition>
            <div
              class="tw-overflow-hidden tw-ease-out tw-duration-300"
              :style="{
                maxHeight:
                  openList && currentIndex === career.id ? '1000px' : '0',
              }"
            >
              <div
                class="tw-p-[30px] tw-text-[#3d3d3dcc] md:tw-py-8 md:tw-px-12 lg:tw-py-10 lg:tw-px-12 lg:tw-leading-8"
              >
                {{ career.summary }}
              </div>
              <div
                class="tw-flex tw-flex-row tw-justify-center tw-pt-0 tw-pr-0 tw-pb-[30px] tw-pl-0 sm:tw-justify-end sm:tw-pt-0 sm:tw-pr-5 sm:tw-pb-[30px] sm:tw-pl-5"
              >
                <router-link
                  @click.native="trackJobAction('tap_read_more_job', career)"
                  class="tw-flex tw-items-center gradient-border-btn tw-p-3"
                  :to="'/jobs/' + career.unique_id"
                  :aria-label="'Get more details about ' + career.title"
                >
                  <font-awesome-icon
                    class="fa tw-text-pink-300 tw-mr-[5px]"
                    icon="align-left"
                    flip="vertical"
                  />
                  <span
                    class="tw-text-[.9rem] tw-leading-[1.364rem] tw-font-bold tw-tracking-[0.06rem] md:tw-text-[1rem] lg:tw-text-[1.125rem]"
                    >Read More</span
                  >
                </router-link>
                <router-link
                  @click.native="trackJobAction('tap_apply_job', career)"
                  class="tw-flex tw-items-center tw-m-[5px] gradient-btn tw-p-3"
                  :to="'/jobs/apply/' + career.unique_id"
                >
                  <font-awesome-icon
                    class="fa"
                    icon="check-circle"
                    aria-hidden="true"
                  />
                  <span
                    class="tw-text-[.9rem] tw-leading-[1.364rem] tw-font-bold tw-tracking-[0.06rem] md:tw-text-[1rem] lg:tw-text-[1.125rem]"
                    >Apply Now</span
                  >
                </router-link>
              </div>
            </div>
          </collapse-transition>
          <div class="tw-bg-[#e2e2e2] tw-h-px" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import loader from "@/assets/images/theme/loader.svg";
import { useJobListStore } from "@/stores/jobs";
import { mapState } from "pinia";
import { mapActions } from "pinia";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faApple,
  faAndroid,
  faNodeJs,
  faGolang,
  faVuejs,
} from "@fortawesome/free-brands-svg-icons";
import {
  faGlobe,
  faBullhorn,
  faPenNib,
  faUser,
  faF,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faApple,
  faAndroid,
  faPenNib,
  faGlobe,
  faBullhorn,
  faUser,
  faF,
  faNodeJs,
  faGolang,
  faVuejs,
);

export default {
  data() {
    return {
      currentIndex: -1, // Initialize to -1 to ensure no item is expanded initially
      openList: false, // Initialize to false
      loader: loader,
    };
  },
  components: {
    CollapseTransition,
    FontAwesomeIcon,
  },
  async serverPrefetch() {
    await this.loadJobs();
  },
  mounted() {
    this.loadJobs();
  },
  computed: {
    ...mapState(useJobListStore, {
      careers: "items",
      jobsError: "error",
    }),
  },
  inject: ["mixpanel"],
  methods: {
    ...mapActions(useJobListStore, ["loadJobs"]),
    expandListItem(id) {
      if (this.currentIndex === id) {
        // Toggle openList if clicking the same item
        this.openList = !this.openList;
      } else {
        // Expand a different item
        this.openList = true;
        this.currentIndex = id;
      }
      this.mixpanel.track(
        this.openList ? "tap_job_expand" : "tap_job_collapse",
      );
    },
    trackJobAction(action, career) {
      this.mixpanel.track(action, { career });
    },
  },
};
</script>
