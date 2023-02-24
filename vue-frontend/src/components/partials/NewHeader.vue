<template>
  <div class="tw-relative tw-z-[5] tw-font-inter-medium">
    <div
      class="tw-sticky tw-top-0"
      :style="{ height: navContainerHeight + 'px' }"
    >
      <nav
        class="tw-w-full tw-bg-white tw-z-[1] tw-py-5 lg:tw-py-[1.75rem] md:tw-px-0 tw-text-black-core/[.87] tw-tracking-[0] tw-transition-all tw-ease tw-duration-500"
        :class="[
          showNavbar
            ? 'tw-fixed tw-shadow-[0_13px_35px_-12px_rgba(35,35,35,0.15)]'
            : '',
        ]"
        :style="{
          transform: animateNavbar
            ? 'translateY(0)'
            : 'translateY(-' + navContainerHeight + 'px)',
        }"
        ref="mainHeader"
      >
        <div
          class="tw-container tw-flex tw-flex-col lg:tw-flex-row tw-flex-wrap lg:tw-flex-nowrap tw-items-start lg:tw-items-center"
        >
          <router-link to="/" replace>
            <div
              class="lg:tw-mt-2 tw-mr-4 tw-text-[1.25rem] tw-whitespace-nowrap tw-no-underline"
            >
              <img
                src="@/assets/images/logo/logo-header.svg"
                class="tw-w-[205px] tw-h-[38.5px]"
                alt="canopas-logo"
              />
            </div>
          </router-link>

          <div
            class="tw-flex tw-basis-[auto] tw-grow tw-items-center tw-mt-3 lg:tw-mt-0"
          >
            <ul
              class="tw-flex tw-flex-row tw-flex-wrap tw-items-center tw-justify-start lg:tw-ml-auto tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
            >
              <li
                v-for="navbar in navbars"
                :key="navbar"
                class="tw-my-2 sm:tw-my-0 tw-ml-0"
              >
                <router-link
                  v-if="!navbar.target"
                  :to="navbar.url"
                  @click.native="mixpanel.track(navbar.event)"
                  :class="[
                    navbar.className
                      ? navbar.className
                      : 'tw-inline-block tw-relative tw-mr-[20px] sm:tw-mr-[30px] after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent',
                    currentRoutePath == navbar.url ||
                    (navbar.name == 'Portfolio' &&
                      currentRoutePath.includes('portfolio')) ||
                    (navbar.name == 'Career' &&
                      currentRoutePath.includes('jobs'))
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
                  ></router-link
                >
                <a
                  v-else
                  class="tw-inline-block tw-relative tw-mr-[20px] sm:tw-mr-[30px] after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent"
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
  </div>
</template>

<script type="module">
import Config from "@/config.js";
export default {
  data() {
    return {
      id: this.$route.params.id,
      animateNavbar: true,
      showNavbar: false,
      navContainerHeight: 100,
      lastScrollY: 0,
      currentRoutePath: this.$router.currentRoute._value.path,
      podcastRef: null,
      navbars: [
        {
          name: "Portfolio",
          url: "/portfolio",
          event: "tap_header_portfolio",
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          target: true,
          event: "tap_header_blog",
        },
        {
          name: "Career",
          url: "/jobs",
          event: "tap_header_career",
        },
        {
          name: "About",
          url: "/about",
          className: Config.SHOW_ABOUT_US_PAGE ? "" : "tw-hidden",
          event: "tap_header_about",
        },
        {
          name: "Let's talk",
          url: "/contact",
          className: "lg:tw-hidden",
          event: "tap_header_cta",
        },
        {
          name: "Get Free Consultation",
          url: "/contact",
          className:
            "tw-hidden lg:tw-block tw-m-0 tw-ml-8 tw-rounded-full tw-py-3 tw-px-3 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0] tw-tracking-[0] gradient-btn consultation-btn",
          showContactBtn: true,
          event: "tap_header_cta",
        },
      ],
    };
  },
  mounted() {
    if (
      this.currentRoutePath.includes("portfolio/") &&
      window.innerWidth > 992 &&
      !["/portfolio", "/portfolio/"].includes(this.currentRoutePath)
    ) {
      this.podcastRef = document.getElementById("response");
      this.podcastRef.style.marginTop = -this.navContainerHeight + "px";
      this.podcastRef.addEventListener("scroll", this.handleScroll);
    }
    window.addEventListener("scroll", this.handleScroll);
    this.navContainerHeight = this.$refs.mainHeader.clientHeight;
  },
  unmounted() {
    if (this.podcastRef) {
      this.podcastRef.removeEventListener("scroll", this.handleScroll);
    }
    window.removeEventListener("scroll", this.handleScroll);
  },
  inject: ["mixpanel"],
  methods: {
    handleScroll() {
      let scrollY = this.podcastRef
        ? this.podcastRef.scrollTop
        : window.scrollY;
      let pageYOffset = this.podcastRef ? scrollY : window.pageYOffset;
      let scrollUp = this.lastScrollY > scrollY && pageYOffset > 100;
      this.showNavbar = scrollUp
        ? scrollUp
        : this.showNavbar && pageYOffset > 0;
      this.animateNavbar = scrollUp || pageYOffset < 100;
      this.lastScrollY = scrollY;
    },
  },
};
</script>
