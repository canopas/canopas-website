<template>
  <div
    class="tw-sticky tw-top-0 tw-z-[2] tw-font-inter-medium"
    :style="{ height: navContainerHeight + 'px' }"
  >
    <nav
      class="tw-absolute tw-left-0 tw-bottom-0 tw-w-full tw-bg-white tw-z-[1] tw-py-2.5 tw-px-[2%] md:tw-py-5 md:tw-px-0 tw-text-black-core/[.87] tw-tracking-[0] tw-transition-all tw-ease-in-out tw-duration-600"
      :class="{
        'tw-fixed tw-left-[unset] tw-bottom-[unset] tw-shadow-[0_13px_35px_-12px_rgba(35,35,35,0.15)]':
          navbarSticky,
        'tw-animated-menuSticky': navbarAnimation,
      }"
      ref="mainHeader"
    >
      <div
        class="tw-container tw-flex tw-flex-col lg:tw-flex-row tw-flex-wrap lg:tw-flex-nowrap tw-items-start lg:tw-items-center"
      >
        <router-link to="/" replace>
          <div
            class="tw-mr-4 tw-text-[1.25rem] tw-whitespace-nowrap tw-no-underline"
          >
            <img
              src="@/assets/images/logo/logo-header.svg"
              class="tw-w-[205px] tw-h-[38.5px] tw-mt-1"
              alt="canopas-logo"
            />
          </div>
        </router-link>

        <div
          class="tw-flex tw-basis-[auto] tw-grow tw-items-center tw-my-[5px]"
        >
          <ul
            class="tw-flex tw-flex-row tw-flex-wrap tw-items-center tw-justify-start tw-mb-[10px] lg:tw-ml-auto tw-p-0 tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
          >
            <li
              v-for="navbar in navbars"
              :key="navbar"
              class="tw-my-2 sm:tw-my-0 tw-ml-0 tw-p-0"
            >
              <router-link
                v-if="!navbar.target"
                :to="navbar.url"
                :class="[
                  navbar.className
                    ? navbar.className
                    : 'tw-inline-block tw-relative tw-my-0 tw-ml-0 tw-mr-[20px] sm:tw-mr-[30px] tw-p-0 after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent',
                  currentRoutePath == navbar.url ||
                  (navbar.name == 'Portfolio' &&
                    currentRoutePath.includes('portfolio')) ||
                  (navbar.name == 'Career' && currentRoutePath.includes('jobs'))
                    ? navbar.showContactBtn
                      ? ''
                      : 'after:tw-w-full after:tw-scale-x-0 after:tw-scale-x-100'
                    : '',
                ]"
                ><span
                  :class="[
                    currentRoutePath == navbar.url && !navbar.showContactBtn
                      ? 'v2-canopas-gradient-text'
                      : '',
                  ]"
                  >{{ navbar.name }}</span
                >
              </router-link>

              <a
                v-else
                class="tw-inline-block tw-relative tw-my-0 tw-ml-0 tw-mr-[20px] sm:tw-mr-[30px] tw-p-0 after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent"
                :href="navbar.url"
                target="_blank"
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
        },
        {
          name: "Career",
          url: "/jobs",
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          target: true,
        },
        {
          name: "Portfolio",
          url: "/portfolio",
        },
        {
          name: "Let's talk",
          url: "/contact",
          className: "lg:tw-hidden",
        },
        {
          name: "Get Free Consultation",
          url: "/contact",
          className:
            " tw-hidden lg:tw-block tw-m-0 tw-rounded-full tw-py-3 tw-px-3 tw-tracking-[0] gradient-btn consultation-btn",
          showContactBtn: true,
        },
      ],
      isJobsUrl: this.$router.currentRoute.value.fullPath.indexOf("/jobs") > -1,
    };
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
    this.navContainerHeight = this.$refs.mainHeader.clientHeight;
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
  },
};
</script>
