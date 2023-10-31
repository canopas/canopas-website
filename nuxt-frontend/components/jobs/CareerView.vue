<template>
  <div class="container my-[14.063rem] mx-auto">
    <div class="text-center">
      <p class="header-text canopas-gradient-text">
        <span class="underline-text">Find t</span>he career of your dreams
      </p>
      <div class="my-[30px] mx-[2%] md:my-[50px] md:mx-20 normal-text">
        <div class="font-bold">
          When was the last time you did something for the first time?
        </div>
        <div class="mt-12">
          At Canopas, we frequently do new things that help you leave your
          comfort zone. And when you leave your comfort zone, you will know what
          you are capable of and you will never be the same again. If you
          believe your work should have meaning and want to make an impact in
          the world with your work, come join us.
        </div>
      </div>
    </div>

    <div v-if="careers == null" class="flex justify-center">
      <img :src="loader" alt="loader-image" />
    </div>
    <div
      v-else
      class="rounded-[14px] border border-solid border-[#e2e2e2] overflow-hidden my-0 mx-[2%] mt-12 md:mx-[6%] normal-text"
    >
      <div
        v-if="careers.length == 0"
        class="text-center py-3.5 px-0 normal-text"
      >
        No matching jobs found
      </div>
      <div v-else class="flex flex-col">
        <div
          v-for="(career, index) in careers"
          :key="career"
          class="flex flex-col"
        >
          <div
            class="flex flex-row w-full items-center p-[30px] cursor-pointer md:py-8 md:px-12 lg:py-10 lg:px-12 normal-text"
            @click="expandListItem(career.id, index)"
          >
            <div class="career-icon w-[50px] sm:w-[60px]">
              <Icon
                class="text-pink-300 icon text-[1.575rem] sm:text-[1.875rem]"
                :name="career.icon_name"
              />
            </div>
            <div class="not-italic font-bold text-black-900">
              {{ career.title }}
            </div>
            <Icon
              class="career-plus-icon ml-auto order-2 text-[#3d3d3dcc]"
              :name="
                openList && career.id == currentIndex
                  ? 'fa6-solid:minus'
                  : 'fa6-solid:plus'
              "
            />
          </div>
          <collapse-transition>
            <div
              ref="careerDetails"
              class="overflow-hidden ease-out duration-300"
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
                class="p-[30px] text-[#3d3d3dcc] md:py-8 md:px-12 lg:py-10 lg:px-12 lg:leading-8"
              >
                {{ career.summary }}
              </div>

              <div
                class="flex flex-row justify-center pt-0 pr-0 pb-[30px] pl-0 sm:justify-end sm:pt-0 sm:pr-5 sm:pb-[30px] sm:pl-5"
              >
                <nuxt-link
                  @click.native="$mixpanel.track('tap_read_more_job')"
                  class="flex items-center gradient-border-btn p-3"
                  :to="'/jobs/' + career.unique_id"
                  :aria-label="'Get more details about ' + career.title"
                >
                  <Icon
                    class="fa text-pink-300 mr-[5px]"
                    name="fa6-solid:align-left"
                    flip="vertical"
                  />
                  <span
                    class="text-[.9rem] leading-[1.364rem] font-bold tracking-[0.06rem] md:text-[1rem] lg:text-[1.125rem]"
                    >Read More</span
                  >
                </nuxt-link>

                <nuxt-link
                  @click.native="$mixpanel.track('tap_apply_job')"
                  class="flex items-center m-[5px] gradient-btn p-3"
                  :to="'/jobs/apply/' + career.unique_id"
                >
                  <Icon
                    class="fa"
                    name="fa6-solid:circle-check"
                    aria-hidden="true"
                  />
                  <span
                    class="text-[.9rem] leading-[1.364rem] font-bold tracking-[0.06rem] md:text-[1rem] lg:text-[1.125rem]"
                    >Apply Now</span
                  >
                </nuxt-link>
              </div>
            </div>
          </collapse-transition>
          <div class="bg-[#e2e2e2] h-px" v-if="index <= careers.length"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import loader from "@/assets/images/theme/loader.svg";
import { useJobListStore } from "@/stores/jobs";

const { $mixpanel } = useNuxtApp();

const currentIndex = ref(0);
const previousIndex = ref(0);
const openList = ref(true);

const store = useJobListStore();
const careers = computed(() => store.items);
const jobsError = computed(() => store.error);

await useAsyncData("jobs", () => store.loadJobs());

function expandListItem(id, index) {
  if (previousIndex.value == id && openList.value) {
    openList.value = false;
    $mixpanel.track("tap_job_collapse_comp");
  } else {
    openList.value = true;
    $mixpanel.track("tap_job_expand_comp");
  }

  currentIndex.value = id;
  previousIndex.value = currentIndex.value;
}
</script>
