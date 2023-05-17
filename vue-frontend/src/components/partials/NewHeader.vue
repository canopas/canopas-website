<template>
  <div class="tw-relative tw-z-[5] tw-font-inter-medium">
    <div
      class="tw-sticky tw-top-0"
      :style="{ height: navContainerHeight + 'px' }"
    >
      <nav
        class="tw-w-full tw-bg-white tw-z-[1] tw-py-5 md:tw-py-[1rem] md:tw-px-0 tw-text-black-core/[.87] tw-tracking-[0] tw-transition-all tw-ease tw-duration-500"
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
          class="tw-container tw-flex tw-flex-row tw-flex-wrap tw-items-center"
        >
          <router-link to="/" replace>
            <div
              class="tw-mt-2 tw-mr-4 tw-text-[1.25rem] tw-whitespace-nowrap tw-no-underline"
            >
              <img
                src="@/assets/images/logo/logo-header.svg"
                class="tw-w-[205px] tw-h-[38.5px]"
                alt="canopas-logo"
              />
            </div>
          </router-link>
          <div
            class="tw-flex tw-relative tw-basis-[auto] tw-grow tw-items-center"
          >
            <ul
              class="tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] xl:tw-text-[1.1875rem] xl:tw-leading-[1.4375rem] tw-hidden lg:tw-flex lg:tw-flex-row tw-flex-wrap lg:tw-items-center tw-justify-evenly xl:tw-justify-end tw-w-full"
            >
              <li
                v-for="navbar in navbars"
                :key="navbar"
                class="tw-my-2 tw-py-2"
              >
                <router-link
                  v-if="!navbar.target"
                  :to="navbar.url"
                  @click.native="mixpanel.track(navbar.event)"
                  class="tw-inline-block tw-relative tw-mr-[20px] sm:tw-mr-[30px] lg:tw-mr-[8px] xl:tw-mr-[30px]"
                  :class="[
                    navbar.className
                      ? navbar.className
                      : navbar.class +
                        ' after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent',
                    currentRoutePath.includes(navbar.url)
                      ? navbar.showContactBtn
                        ? ''
                        : 'after:tw-w-full after:tw-scale-x-0 after:tw-scale-x-100'
                      : '',
                  ]"
                >
                  <span
                    :class="[
                      currentRoutePath == navbar.url && !navbar.showContactBtn
                        ? 'v2-canopas-gradient-text'
                        : '',
                    ]"
                    class="tw-flex tw-mt-[0.1rem]"
                    >{{ navbar.name }}</span
                  ></router-link
                >
                <a
                  v-else
                  class="tw-inline-block tw-relative tw-mr-[20px] sm:tw-mr-[30px] lg:tw-mr-[8px] xl:tw-mr-[30px] after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent"
                  :href="navbar.url"
                  :target="navbar.target"
                  >{{ navbar.name }}</a
                >
              </li>
            </ul>
          </div>
          <div class="lg:tw-hidden tw-flex tw-items-end tw-w-6 tw-h-6">
            <font-awesome-icon
              :icon="faBars"
              class="tw-ml-0 tw-mt-[0.20rem] tw-w-full tw-h-full"
              @click="showMenu = true"
            />
          </div>
        </div>
      </nav>
    </div>
    <div
      :class="showMenu ? 'tw-translate-x-0' : 'tw-translate-x-full'"
      class="lg:tw-hidden tw-fixed tw-top-0 tw-right-0 tw-w-full tw-h-screen tw-bg-white tw-py-4 md:tw-py-2 tw-duration-500 tw-ease-in-out tw-overflow-y-hidden"
    >
      <div class="tw-container">
        <div class="tw-sticky tw-top-0 tw-w-full md:tw-pt-[0.75rem]">
          <div class="tw-flex tw-justify-between tw-items-center">
            <div class="tw-text-[1.25rem] tw-whitespace-nowrap">
              <router-link to="/" replace>
                <div
                  class="tw-mt-2 tw-mr-4 tw-text-[1.25rem] tw-whitespace-nowrap tw-no-underline"
                >
                  <img
                    src="@/assets/images/logo/logo-header.svg"
                    class="tw-w-[205px] tw-h-[38.5px]"
                    alt="canopas-logo"
                  />
                </div>
              </router-link>
            </div>
            <div class="tw-w-6 tw-h-6 tw-mt-[-3px]">
              <font-awesome-icon
                :icon="faXmark"
                class="tw-h-full tw-w-full"
                @click="showMenu = false"
              />
            </div>
          </div>
        </div>
        <ul
          class="tw-flex tw-flex-col lg:tw-items-center tw-justify-start tw-h-[50%] sm:tw-h-[45%] tw-mt-8 lg:tw-ml-auto tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-overflow-y-scroll"
        >
          <li
            v-for="navbar in navbars.slice(0, navbars.length - 1)"
            :key="navbar"
            class="tw-my-2 tw-py-2 tw-ml-1 lg:tw-ml-0"
          >
            <router-link
              v-if="!navbar.target"
              :to="navbar.url"
              @click.native="mixpanel.track(navbar.event)"
              class="tw-inline-block tw-relative tw-mr-[20px] sm:tw-mr-[30px] lg:tw-mr-[20px]"
              :class="[
                navbar.className
                  ? navbar.className
                  : navbar.class +
                    ' after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent',
                currentRoutePath.includes(navbar.url)
                  ? ''
                    ? ''
                    : 'after:tw-w-full after:tw-scale-x-0 after:tw-scale-x-100'
                  : '',
              ]"
            >
              <span
                :class="[
                  currentRoutePath == navbar.url
                    ? 'v2-canopas-gradient-text'
                    : '',
                ]"
                class="tw-flex tw-mt-[0.1rem]"
                >{{ navbar.name }}</span
              ></router-link
            >
            <a
              v-else
              class="tw-inline-block tw-relative tw-mr-[20px] sm:tw-mr-[30px] after:tw-absolute after:tw-w-full after:tw-h-[3px] after:tw-top-[27px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-pink-300 after:tw-origin-bottom-left after:tw-duration-300 after:tw-scale-x-0 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left hover:tw-bg-clip-text hover:tw-bg-gradient-[270.11deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-text-transparent"
              :href="navbar.url"
              :target="navbar.target"
              >{{ navbar.name }}</a
            >
          </li>
        </ul>
        <div v-if="showMenu" class="tw-w-full tw-h-auto tw-bg-white tw-p-4">
          <div className="tw-grid tw-p-3 tw-pb-5">
            <a
              href="/contact"
              className="tw-relative tw-justify-self-center tw-rounded-full tw-border-[1px] tw-border-solid tw-border-transparent tw-bg-gradient-to-r tw-from-[#f2709c] tw-to-[#ff9472] hover:tw-shadow-[inset_2px_1000px_1px_#fff] tw-font-bold tw-text-white "
            >
              <div
                className="tw-py-[0.8rem] tw-px-[2.5rem] tw-text-lg hover:tw-bg-clip-text hover:tw-bg-gradient-to-r hover:tw-from-[#f2709c] hover:tw-via-[#ff909c] hover:tw-to-[#ff9472] hover:tw-text-transparent tw-inline-block"
              >
                Get Free Consultation
              </div>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import Config from "@/config.js";
import { faBars, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
export default {
  data() {
    return {
      id: this.$route.params.id,
      faBars,
      faXmark,
      width: 576,
      showMenu: false,
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
          name: "Resources",
          url: "/resources",
          target: "_self",
          event: "tap_header_resources",
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          target: "_blank",
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
          event: "tap_header_about",
        },
        {
          name: "Get Free Consultation",
          url: "/contact",
          className:
            "tw-m-0 tw-rounded-full tw-py-3 tw-px-3 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0] tw-tracking-[0] gradient-btn consultation-btn",
          showContactBtn: true,
          event: "tap_header_cta",
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {
    if (Config.SHOW_SERVICES_PAGE) {
      this.navbars.splice(0, 0, {
        name: "Services",
        url: "/services",
        event: "tap_header_services",
      });
    }

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
    this.width = window.innerWidth;
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
