<template>
  <div class="tw-container tw-my-[100px] tw-mx-auto">
    <div class="tw-text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Find t</span>he career of your dreams
      </div>
      <div
        class="tw-my-[30px] tw-mx-[2%] md:tw-my-[50px] md:tw-mx-[80px] normal-text"
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

    <div v-if="careers == null" class="tw-flex tw-justify-center">
      <img :src="loader" />
    </div>
    <div
      v-else
      class="tw-rounded-[14px] tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-overflow-hidden tw-my-0 tw-mx-[2%] tw-mt-12 md:tw-mx-[6%] normal-text"
    >
      <div
        v-if="careers.length == 0"
        class="tw-text-center tw-py-[14px] tw-px-0 normal-text"
      >
        No matching jobs found
      </div>
      <div v-else class="tw-flex tw-flex-col">
        <div
          v-for="(career, index) in careers"
          :key="career"
          class="tw-flex tw-flex-col"
        >
          <div
            class="tw-flex tw-flex-row tw-w-full tw-items-center tw-p-[30px] tw-cursor-pointer md:tw-py-[32px] md:tw-px-[48px] lg:tw-py-[40px] lg:tw-px-[48px] normal-text"
            @click="expandListItem(career.id, index)"
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
              :icon="openList && career.id == currentIndex ? 'minus' : 'plus'"
            />
          </div>
          <collapse-transition>
            <div
              ref="careerDetails"
              class="tw-overflow-hidden tw-ease-out tw-duration-300"
              :key="career.summary"
              :style="[
                openList && career.id == currentIndex
                  ? {
                      maxHeight: `1000px`,
                    }
                  : {
                      maxHeight: `0`,
                    },
              ]"
            >
              <div
                class="tw-p-[30px] tw-text-[#3d3d3dcc] md:tw-py-[32px] md:tw-px-[48px] lg:tw-py-[40px] lg:tw-px-[48px] lg:tw-leading-8"
              >
                {{ career.summary }}
              </div>

              <div
                class="tw-flex tw-flex-row tw-justify-center tw-pt-0 tw-pr-0 tw-pb-[30px] tw-pl-0 sm:tw-justify-end sm:tw-pt-0 sm:tw-pr-[20px] sm:tw-pb-[30px] sm:tw-pl-[20px]"
              >
                <router-link
                  @click.native="$gtag.event('tap_read_more_job')"
                  class="gradient-border-btn tw-p-[11px] sm:tw-p-[16px]"
                  :to="'/jobs/' + career.unique_id"
                  aria-label="read more"
                >
                  <font-awesome-icon
                    class="fa tw-text-pink-300"
                    icon="align-left"
                    flip="vertical"
                    aria-hidden="true"
                  />
                  <span
                    class="tw-text-[.9rem] tw-leading-[1.364rem] tw-font-bold tw-tracking-[0.06rem] md:tw-text-[1rem] lg:tw-text-[1.125rem]"
                    >Read More</span
                  >
                </router-link>

                <router-link
                  @click.native="$gtag.event('tap_apply_job')"
                  class="gradient-btn tw-m-[5px] tw-p-[11px] sm:tw-p-[16px]"
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
          <div
            class="tw-bg-[#e2e2e2] tw-h-px"
            v-if="index <= careers.length"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import loader from "@/assets/images/theme/loader.svg";
import { useJobListStore } from "@/stores/jobs";
import { mapState } from "pinia";
import { mapActions } from "pinia";

import {
  faApple,
  faAndroid,
  faNodeJs,
  faGolang,
  faVuejs,
} from "@fortawesome/free-brands-svg-icons";
import { library } from "@fortawesome/fontawesome-svg-core";
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
  faVuejs
);

export default {
  data() {
    return {
      currentIndex: 0,
      openList: true,
      previousIndex: 0,
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
  methods: {
    ...mapActions(useJobListStore, ["loadJobs"]),
    expandListItem(id, index) {
      if (this.previousIndex == id && this.openList) {
        this.openList = false;
        this.$gtag.event("tap_job_collapse");
      } else {
        this.openList = true;
        this.$gtag.event("tap_job_expand");
      }

      this.currentIndex = id;
      this.previousIndex = this.currentIndex;
    },
  },
};
</script>
