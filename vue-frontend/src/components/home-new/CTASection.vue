<template>
  <div class="tw-bg-white">
    <div
      class="tw-bg-black-900 tw-py-[100px] tw-px-0 tw-overflow-hidden tw-rounded-tl-lg tw-rounded-br-lg tw-relative tw-z-0 tw-rounded-tl-[100px] tw-rounded-br-[100px] sm:tw-rounded-tl-[150px] sm:tw-rounded-br-[150px] md:tw-rounded-tl-[200px] md:tw-rounded-br-[200px] lg:tw-rounded-tl-[250px] lg:tw-rounded-br-[250px]"
    >
      <img
        :src="backgroundImage"
        class="tw-w-full tw-h-full tw-top-0 tw-left-0 tw-absolute tw-object-cover tw--z-[1]"
        alt="social-media-background"
      />
      <div
        class="tw-container sm:tw-w-4/5 lg:tw-w-[70%] tw-text-center tw-p-y-0 tw-px-[25px] tw-m-auto v2-header-3-text"
      >
        <div class="tw-m-auto lg:tw-w-4/5 canopas-gradient-text">
          How do you know we’re a good fit for you?
        </div>
        <div class="v2-title-3-text tw-text-white tw-mt-6">
          Let’s have a short chat.
        </div>
        <div class="tw-flex tw-flex-wrap tw-justify-center tw-m-auto">
          <div
            class="tw-m-5 tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-text-center tw-px-6 tw-py-5 tw-rounded-[15px]"
          >
            <div class="v2-normal-3-text tw-text-white">I love talking.</div>
            <router-link
              class="gradient-btn tw-rounded-[2rem] tw-flex tw-items-center hover:tw-shadow-4xl"
              :to="contactURL"
              @click.native="$gtag.event('tap_home_page_end_cta_call')"
            >
              <font-awesome-icon
                class="fa v2-title-3-text tw-text-[1.35rem] tw-leading-[1.5rem] tw-text-white"
                icon="phone"
                aria-hidden="true"
              />
              <span
                class="v2-normal-3-text tw-text-white tw-my-0 tw-mr-0 tw-ml-[10px] tw-my-0 tw-mx-[6px] !tw-font-light"
                >Schedule a call</span
              >
            </router-link>
          </div>
          <div
            class="tw-m-5 tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-text-center tw-px-6 tw-py-5 tw-rounded-[15px]"
          >
            <div class="v2-normal-3-text tw-text-white tw-font-light">
              I prefer mail.
            </div>
            <router-link
              class="gradient-btn tw-rounded-[2rem] tw-flex tw-items-center hover:tw-shadow-4xl"
              :to="contactURL"
              @click.native="$gtag.event('tap_home_page_end_cta_mail')"
            >
              <font-awesome-icon
                class="fa v2-title-3-text tw-text-[1.35rem] tw-leading-[1.5rem] tw-text-white"
                icon="paper-plane"
                aria-hidden="true"
              />
              <span
                class="v2-normal-3-text tw-text-white tw-my-0 tw-mr-0 tw-ml-[10px] !tw-font-light"
                >Drop a mail</span
              >
            </router-link>
          </div>
        </div>
        <div
          class="v2-title-3-text tw-text-[1.35rem] tw-leading-[1.5rem] tw-text-white tw-text-left tw-m-auto lg:tw-w-4/5 tw-mt-12"
        >
          <div class="v2-canopas-gradient-text">
            Four reasons why do we do this call?
          </div>
        </div>
        <ul
          class="v2-normal-2-text tw-text-white tw-list-none tw-m-auto tw-p-0 lg:tw-w-4/5 tw-text-left tw-mt-6"
        >
          <li class="tw-py-2.5 tw-px-0 tw-flex">
            <span
              class="tw-content-none tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-w-[16px] tw-h-[16px] tw-rounded-[30%] tw-inline-block tw-mt-[5px] tw-mr-[24px] tw-mb-0 tw-ml-0"
            ></span>
            <div class="tw-w-[90%]">
              It gives an opportunity to learn more about each other.
            </div>
          </li>
          <li class="tw-py-2.5 tw-px-0 tw-flex">
            <span
              class="tw-content-none tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-w-[16px] tw-h-[16px] tw-rounded-[30%] tw-inline-block tw-mt-[5px] tw-mr-[24px] tw-mb-0 tw-ml-0"
            ></span>
            <div class="tw-w-[90%]">
              We want to give you a chance to put your idea into action. You
              might have faced this situation before where you wish you had
              someone to help you.
            </div>
          </li>
          <li class="tw-py-2.5 tw-px-0 tw-flex">
            <span
              class="tw-content-none tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-w-[16px] tw-h-[16px] tw-rounded-[30%] tw-inline-block tw-mt-[5px] tw-mr-[24px] tw-mb-0 tw-ml-0"
            ></span>
            <div class="tw-w-[90%]">
              Sometimes, you want to check the feasibility of your idea. We will
              give you honest feedback based on our industry experience.
            </div>
          </li>
          <li class="tw-py-2.5 tw-px-0 tw-flex">
            <span
              class="tw-content-none tw-bg-gradient-to-b tw-from-pink-300 tw-to-orange-300 tw-w-[16px] tw-h-[16px] tw-rounded-[30%] tw-inline-block tw-mt-[5px] tw-mr-[24px] tw-mb-0 tw-ml-0"
            ></span>
            <div class="tw-w-[90%]">
              We offer 100% Money-back guarantee. so, there’s no risk to you.
            </div>
          </li>
        </ul>
      </div>
      <div
        class="tw-text-center ..."
        :class="[
          currentRoutePath.includes('portfolio')
            ? 'tw-pt-16 xl:tw-pb-1'
            : 'tw-hidden',
        ]"
      >
        <span class="v2-normal-3-text tw-text-white">
          <font-awesome-icon :icon="copyrightIcon" /> &nbsp;{{
            new Date().getFullYear()
          }}
          Canopas Software LLP. All rights reserved.
        </span>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import { faCopyright } from "@fortawesome/free-regular-svg-icons";
import backgroundImage from "@/assets/images/theme/cta-background.webp";

export default {
  data() {
    return {
      copyrightIcon: faCopyright,
      id: this.$route.params.id,
      currentRoutePath: this.$router.currentRoute._value.path,
      backgroundImage: backgroundImage,
      contactURL: "/contact",
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>
