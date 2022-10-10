<template>
  <div
    class="tw-sticky tw-top-0 tw-z-[1]"
    :style="{ height: navContainerHeight + 'px' }"
  >
    <nav
      class="tw-py-2.5 tw-px-[2%] tw-w-full tw-bg-white tw-absolute tw-left-0 tw-bottom-0 tw-z-[1] tw-transition-all tw-ease-in-out tw-duration-600 md:tw-py-5 md:tw-px-0"
      :class="{
        'tw-fixed tw-left-[unset] tw-bottom-[unset] tw-shadow-[0_13px_35px_-12px_rgba(35,35,35,0.15)]':
          navbarSticky,
        'tw-animated-menuSticky': navbarAnimation,
      }"
      ref="mainHeader"
    >
      <div
        class="tw-container tw-flex tw-flex-col md:tw-flex-row tw-flex-wrap md:tw-flex-nowrap tw-justify-start"
      >
        <router-link to="/" replace>
          <div
            class="tw-py-[0.3125rem] tw-mr-4 tw-text-[1.25rem] tw-whitespace-nowrap tw-no-underline tw-text-black"
          >
            <img
              src="@/assets/images/logo/logo-header.svg"
              class="tw-w-[205px] tw-h-[38.5px] tw-mt-1"
              alt="canopas-logo"
            />
          </div>
        </router-link>

        <div class="tw-flex tw-basis-[auto] tw-grow tw-items-center">
          <ul
            class="tw-flex tw-flex-row tw-flex-wrap tw-items-center tw-justify-start tw-mb-[10px] tw-my-0 tw-text-black-900/80 tw-p-0 md:tw-ml-auto"
          >
            <li
              v-for="navbar in navbars"
              :key="navbar"
              class="tw-my-2 sm:tw-my-0 tw-ml-0 tw-text-black-900/80 tw-p-0"
            >
              <router-link
                v-if="!navbar.target"
                class="v2-normal-3-text"
                :class="[
                  navbar.className,
                  currentRoutePath == navbar.url ||
                  (navbar.name == 'Portfolio' &&
                    currentRoutePath.includes('portfolio')) ||
                  (navbar.name == 'Career' && currentRoutePath.includes('jobs'))
                    ? navbar.showContactBtn
                      ? 'v2-button'
                      : 'tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-h-0.5 after:tw-scale-x-0 after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-scale-x-100'
                    : '',
                ]"
                :to="navbar.url"
                @click.native="sendEvent(navbar.event)"
                >{{ navbar.name }}</router-link
              >

              <a
                v-else
                class="v2-normal-3-text"
                :class="[
                  navbar.className,
                  currentRoutePath == navbar.url ||
                  (navbar.name == 'Portfolio' &&
                    currentRoutePath.includes('portfolio')) ||
                  (navbar.name == 'Career' && currentRoutePath.includes('jobs'))
                    ? navbar.showContactBtn
                      ? 'v2-button '
                      : 'tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-h-0.5 after:tw-scale-x-0 after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-scale-x-100'
                    : '',
                ]"
                :href="navbar.url"
                target="_blank"
                @click.native="sendEvent(navbar.event)"
                >{{ navbar.name }}</a
              >
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </div>
</template>

<script type="module">
import Config from "@/config.js";
export default {
  data() {
    return {
      id: this.$route.params.id,
      navbarSticky: false,
      navbarAnimation: false,
      navContainerHeight: 133,
      lastScrollY: 0,
      currentRoutePath: this.$router.currentRoute._value.path,
      navbars: [
        {
          name: "Home",
          url: "/",
          className:
            " tw-mt-0 tw-mb-0 tw-ml-0 tw-p-0 tw-text-black-900/80 hover:tw-text-black-900/80 tw-mr-[20px] sm:tw-mr-[40px] tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left",
          target: false,
          event: this.isJobsUrl ? "tap_jobs_header_home" : " ",
        },
        {
          name: "Career",
          url: "/jobs",
          className:
            "tw-mt-0  tw-mb-0 tw-ml-0 tw-p-0 tw-text-black-900/80 hover:tw-text-black-900/80 tw-mr-[20px] sm:tw-mr-[40px] tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left",
          target: false,
          event: this.isJobsUrl ? " " : "tap_header_view_career_page",
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          className:
            "tw-mt-0  tw-mb-0 tw-ml-0 tw-p-0 tw-text-black-900/80 hover:tw-text-black-900/80 tw-mr-[20px] sm:tw-mr-[40px] tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left ",
          target: true,
          event: this.isJobsUrl
            ? "tap_jobs_header_blog"
            : "tap_header_blog_page",
        },
        {
          name: "Portfolio",
          url: "/portfolio",
          className:
            " tw-mt-0  tw-mb-0 tw-ml-0 tw-p-0 tw-text-black-900/80 hover:tw-text-black-900/80 tw-mr-[20px] sm:tw-mr-[40px] tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left",
          target: false,
          event: this.isJobsUrl
            ? "tap_jobs_header_portfolio"
            : "tap_header_portfolio_page",
        },
        {
          name: "Let's talk",
          url: "/contact",
          className:
            " lg:tw-mr-0 tw-block lg:tw-hidden tw-mt-0 tw-mb-0 tw-ml-0 tw-p-0 tw-text-black-900/80 hover:tw-text-black-900/80 tw-inline-block tw-relative after:tw-absolute after:tw-w-2/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left",
          target: false,
          event: this.isJobsUrl
            ? "tap_jobs_header_contact_cta"
            : "tap_header_contact_cta",
        },
        {
          name: "Let's talk",
          url: "/contact",
          className:
            "tw-mr-0 tw-hidden lg:tw-block tw-border tw-border-black-900 v2-button tw-py-[0.4rem] tw-px-6 tw-bg-black-900 tw-rounded-full tw-text-white tw-shadow-[0_4px_4px_rgba(0,0,0,0.5)] hover:tw-bg-[#fff] hover:tw-text-[#3d3d3d]",
          target: false,
          showContactBtn: true,
          event: this.isJobsUrl
            ? "tap_jobs_header_contact_cta"
            : "tap_header_contact_cta",
        },
      ],
      isJobsUrl: this.$router.currentRoute.value.fullPath.indexOf("/jobs") > -1,
    };
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
    this.navContainerHeight = this.$refs.mainHeader.clientHeight + 30;
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      let wasSticky = this.navbarSticky;
      this.navbarSticky = window.scrollY > 30;
      // If scroll diff is large, we show navbar with animation
      let diff = window.scrollY - this.lastScrollY;
      if (this.navbarSticky != wasSticky && diff > 15) {
        this.navbarAnimation = true;
      }
      if (!this.navbarSticky) {
        this.navbarAnimation = false;
      }
      this.lastScrollY = window.scrollY;
    },
    sendEvent(event) {
      if (event && event != " ") {
        this.$gtag.event(event);
      }
    },
  },
};
</script>
