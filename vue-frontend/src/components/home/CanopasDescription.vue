<template>
  <div class="tw-relative tw-p-0 tw-m-0 tw-py-[60px] tw-px-0">
    <div class="tw-container tw-text-center">
      <div class="tw-m-auto tw-w-[74%] sm:tw-w-6/12" ref="description">
        <div
          ref="firstHeaderText"
          class="header-text canopas-gradient-text tw-pb-3"
        >
          Do you have a massive ad budget for your business?
        </div>

        <div ref="secondHeaderText" class="tw-mt-12 normal-text">
          Unless you’re a Multimillion or a Billion-dollar company, you probably
          don’t have a multimillion-dollar ad budget or professional Spinners.
          <br /><br /><br />
          Your product needs to stand out on its own merits like App Quality,
          Performance, UI design, and User Experience.
        </div>
        <div ref="arrowImage" class="tw-text-center down-arrow">
          <img
            class="tw-h-48 tw-w-full tw-my-[5%] tw-mx-auto"
            :src="arrowImage"
            alt=""
          />
        </div>
        <div ref="thirdHeaderText" class="normal-text second-div">
          Most companies don't care about you, your product, and your vision or
          dreams. They don't give a damn about whether their work helped you to
          get more business, revenue, users, or solving a problem.
          <br /><br /><br />
          That's where CANOPAS comes into the picture.
          <br /><br /><br />
          We help Entrepreneurs, startups, and small companies to bring their
          IDEA to LIFE by developing digital products for their business.
          <br /><br /><br />
          We offer a 100% MONEY BACK GUARANTEE FOR 14 DAYS if you don't like our
          work. No questions asked.
        </div>

        <div class="tw-flex tw-flex-col tw-items-center tw-pt-12">
          <div class="tw-text-center">
            <router-link
              class="v2-normal-3-text active:tw-scale-[0.98] tw-flex tw-items-center tw-py-4 tw-px-[1.1rem] tw-bg-black-900 tw-text-white hover:tw-bg-transparent hover:tw-text-black-900 tw-rounded-[3rem] tw-text-center tw-border-[1px] tw-border-solid tw-border-[#3d3d3d]"
              :to="contactURL"
              @click.native="mixpanel.track('tap_ad_budget_cta')"
            >
              <span class="tw-mr-2.5">Start your free trial</span>
              <font-awesome-icon
                class="arrow fa tw-h-4 tw-w-4"
                icon="arrow-right"
                id="leftArrow"
              />
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import arrowImage from "@/assets/images/theme/arrow.svg";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      arrowImage: arrowImage,
      contactURL: "/contact",
    };
  },
  components: { FontAwesomeIcon },
  inject: ["mixpanel"],
};
</script>
