<template>
  <div class="container mt-16 lg:my-60 mx-auto">
    <div class="text-center">
      <div class="mobile-header-2 lg:desk-header-2 text-primary-1">
        Find the career of your dreams
      </div>
      <div
        class="mt-4 lg:mt-6 sub-h1-semibold lg:mobile-header-2-semibold text-black-60"
      >
        When was the last time you did something for the first time?
      </div>
      <div
        class="mt-4 lg:mt-6 sub-h1-regular lg:mobile-header-2-regular text-black-60"
      >
        At Canopas, we frequently do new things that help you leave your comfort
        zone. And when you leave your comfort zone, you will know what you are
        capable of and you will never be the same again. If you believe your
        work should have meaning and want to make an impact in the world with
        your work, come join us.
      </div>
    </div>

    <div v-if="careers == null" class="flex justify-center">
      <img :src="loader" alt="loader-image" />
    </div>
    <div
      v-else
      class="rounded-2xl border overflow-hidden my-0 mx-[2%] mt-8 lg:mt-[4.5rem] md:mx-[6%]"
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
            class="flex flex-row gap-4 items-center justify-center w-full items-center px-4 lg:px-10 py-6 lg:py-8 cursor-pointer"
            @click="expandListItem(career.id, index)"
          >
            <Icon class="text-pink-300 icon w-6 h-6" :name="career.icon_name" />
            <div
              class="sub-h1-semibold lg:mobile-header-2-semibold text-black-87"
            >
              {{ career.title }}
            </div>
            <Icon
              class="career-plus-icon ml-auto order-2 text-gray-40"
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
                class="sub-h3-regular lg:mobile-header-3-regular text-black-60 px-4 lg:px-10 pb-6 lg:pb-8"
              >
                {{ career.summary }}
              </div>

              <div
                class="flex flex-row justify-center pb-6 px-4 sm:justify-end sm:pt-0 sm:pr-5 sm:pb-[30px] sm:pl-5"
              >
                <nuxt-link
                  @click.native="$mixpanel.track('tap_read_more_job')"
                  class="flex items-center gradient-border-btn primary-btn mx-0"
                  :to="'/jobs/' + career.unique_id"
                  :aria-label="'Get more details about ' + career.title"
                >
                  <Icon
                    class="fa text-black-60 mr-[5px]"
                    name="fa6-solid:align-left"
                    flip="vertical"
                  />
                  <span class="sub-h3-semibold lg:sub-h1-semibold text-black-60"
                    >Read More</span
                  >
                </nuxt-link>

                <nuxt-link
                  @click.native="$mixpanel.track('tap_apply_job')"
                  class="flex items-center m-[5px] gradient-btn primary-btn"
                  :to="'/jobs/apply/' + career.unique_id"
                >
                  <Icon
                    class="fa"
                    name="fa6-solid:circle-check"
                    aria-hidden="true"
                  />
                  <span class="sub-h3-semibold lg:sub-h1-semibold"
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
