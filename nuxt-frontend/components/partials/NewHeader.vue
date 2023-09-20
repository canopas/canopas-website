<template>
  <div class="tw-relative tw-z-10 tw-font-inter-medium">
    <div
      class="tw-sticky tw-top-0"
      :style="{ height: navContainerHeight + 'px' }"
    >
      <nav
        class="tw-ease tw-z-[1] tw-w-full tw-bg-white tw-py-5 tw-tracking-[0] tw-text-black-core/[.87] tw-transition-all tw-duration-500 md:tw-px-0 md:tw-py-4"
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
          <nuxt-link to="/" replace>
            <div
              class="tw-mr-4 tw-mt-2 tw-whitespace-nowrap tw-text-[1.25rem] tw-no-underline"
            >
              <img
                src="@/assets/images/logo/logo-header.svg"
                class="tw-h-[38.5px] tw-w-[205px]"
                alt="canopas-logo"
              />
            </div>
          </nuxt-link>
          <div
            class="tw-relative tw-flex tw-grow tw-basis-[auto] tw-items-center"
          >
            <ul
              class="tw-hidden tw-w-full tw-flex-wrap tw-justify-evenly tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-flex lg:tw-flex-row lg:tw-items-center xl:tw-justify-end xl:tw-text-[1.1875rem] xl:tw-leading-[1.4375rem]"
              @mouseleave="showContributionMenu = false"
            >
              <li
                v-for="navbar in navbars"
                :key="navbar"
                class="tw-my-2 tw-py-2"
              >
                <nuxt-link
                  @mouseover="
                    showContributionMenu = navbar.name == 'Contribution'
                  "
                  v-if="!navbar.target"
                  :to="navbar.url"
                  @click.native="mixpanel.track(navbar.event)"
                  class="tw-group tw-relative tw-mr-5 tw-inline-block sm:tw-mr-[30px] lg:tw-mr-2 xl:tw-mr-[30px]"
                  :class="[
                    navbar.className
                      ? navbar.className
                      : navbar.class +
                        ' after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-top-[27px] after:tw-h-[3px] after:tw-w-full  after:tw-origin-bottom-left after:tw-scale-x-0 after:tw-from-[#f2709c] after:tw-to-[#ff9472] after:tw-duration-300 after:tw-bg-gradient-[90deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-bg-gradient-[270.11deg] hover:after:tw-origin-bottom-left hover:after:tw-scale-x-100 ',
                  ]"
                >
                  <span
                    :class="[
                      navbar.id == 7 ? 'hover:v2-canopas-gradient-text' : '',
                    ]"
                    >{{ navbar.name }}</span
                  >
                </nuxt-link>
                <a
                  v-else
                  class="tw-relative tw-mr-5 tw-inline-block after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-top-[27px] after:tw-h-[3px] after:tw-w-full after:tw-origin-bottom-left after:tw-scale-x-0 after:tw-from-[#f2709c] after:tw-to-[#ff9472] after:tw-duration-300 after:tw-bg-gradient-[90deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-bg-gradient-[270.11deg] hover:after:tw-origin-bottom-left hover:after:tw-scale-x-100 sm:tw-mr-[30px] lg:tw-mr-2 xl:tw-mr-[30px]"
                  :href="navbar.url"
                  :target="navbar.target"
                  >{{ navbar.name }}</a
                >

                <ul
                  v-show="showContributionMenu"
                  @mouseleave="showContributionMenu = false"
                  class="tw-absolute tw-w-max tw-flex-col tw-space-y-6 tw-rounded-[5px] tw-border tw-bg-white !tw-px-6 tw-py-[1.5rem] tw-shadow-md lg:tw-top-[4.5rem] lg:tw-left-[169px] xl:tw-left-[296px] 2xl:tw-left-[475px]"
                >
                  <li v-for="navbar in subMenus" :key="navbar">
                    <a
                      class="tw-relative tw-py-4 tw-font-inter-medium tw-text-[1rem] tw-leading-[1.125rem] tw-text-black-core/[0.6] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-text-[1.188rem] lg:tw-leading-[1.437rem] after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-top-[42px] after:tw-h-[3px] after:tw-w-full after:tw-origin-bottom-left after:tw-scale-x-0 after:tw-from-[#f2709c] after:tw-to-[#ff9472] after:tw-duration-300 after:tw-bg-gradient-[90deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-bg-gradient-[270.11deg] hover:after:tw-origin-bottom-left hover:after:tw-scale-x-100"
                      :href="navbar.url"
                      :target="navbar.target"
                      >{{ navbar.name }}
                    </a>
                  </li>
                </ul>
              </li>
            </ul>
          </div>
          <div class="tw-flex tw-h-6 tw-w-6 tw-items-end lg:tw-hidden">
            <font-awesome-icon
              :icon="faBars"
              class="tw-ml-0 tw-mt-[0.20rem] tw-h-full tw-w-full"
              @click="showMenu = true"
            />
          </div>
        </div>
      </nav>
    </div>
    <div
      :class="showMenu ? 'tw-translate-x-0' : 'tw-translate-x-full'"
      class="tw-fixed tw-right-0 tw-top-0 tw-h-screen tw-w-full tw-overflow-y-hidden tw-bg-white tw-py-4 tw-duration-500 tw-ease-in-out md:tw-py-2 lg:tw-hidden"
    >
      <div class="tw-container">
        <div class="tw-sticky tw-top-0 tw-w-full md:tw-pt-3">
          <div class="tw-flex tw-items-center tw-justify-between">
            <div class="tw-whitespace-nowrap tw-text-[1.25rem]">
              <nuxt-link to="/" replace>
                <div
                  class="tw-mr-4 tw-mt-2 tw-whitespace-nowrap tw-text-[1.25rem] tw-no-underline"
                >
                  <img
                    src="@/assets/images/logo/logo-header.svg"
                    class="tw-h-[38.5px] tw-w-[205px]"
                    alt="canopas-logo"
                  />
                </div>
              </nuxt-link>
            </div>
            <div class="tw-mt-[-3px] tw-h-6 tw-w-6">
              <font-awesome-icon
                :icon="faXmark"
                class="tw-h-full tw-w-full"
                @click="showMenu = false"
              />
            </div>
          </div>
        </div>
        <ul
          class="tw-mt-8 tw-flex tw-h-[50%] tw-flex-col tw-justify-start tw-overflow-y-scroll tw-text-[1rem] tw-leading-[1.125rem] sm:tw-h-[45%] md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-ml-auto lg:tw-items-center lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
        >
          <li
            v-for="navbar in navbars.slice(0, navbars.length - 1)"
            :key="navbar"
            class="tw-my-2 tw-ml-1 tw-py-2 lg:tw-ml-0"
          >
            <nuxt-link
              v-if="!navbar.target"
              :to="navbar.url == '/contributions' ? '' : navbar.url"
              @click.native="mixpanel.track(navbar.event)"
              class="tw-group tw-relative tw-mr-5 tw-inline-block sm:tw-mr-[30px] lg:tw-mr-5"
              :class="[
                navbar.className
                  ? navbar.className
                  : navbar.class +
                    ' after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-top-[27px] after:tw-h-[3px] after:tw-w-full  after:tw-origin-bottom-left after:tw-scale-x-0 after:tw-from-[#f2709c] after:tw-to-[#ff9472] after:tw-duration-300 after:tw-bg-gradient-[90deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-bg-gradient-[270.11deg] hover:after:tw-origin-bottom-left hover:after:tw-scale-x-100',
              ]"
              @click="
                navbar.name == 'Contribution'
                  ? (showContributionMenu = !showContributionMenu)
                  : ''
              "
            >
              {{ navbar.name }}</nuxt-link
            >

            <a
              v-else
              class="tw-relative tw-mr-5 tw-inline-block after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-top-[27px] after:tw-h-[3px] after:tw-w-full after:tw-origin-bottom-left after:tw-scale-x-0 after:tw-from-[#f2709c] after:tw-to-[#ff9472] after:tw-duration-300 after:tw-bg-gradient-[90deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-bg-gradient-[270.11deg] hover:after:tw-origin-bottom-left hover:after:tw-scale-x-100 sm:tw-mr-[30px]"
              :href="navbar.url"
              :target="navbar.target"
              >{{ navbar.name }}</a
            >
            <ul
              :class="
                showContributionMenu && navbar.name == 'Contribution'
                  ? 'tw-mt-8 tw-w-full tw-bg-white tw-px-6'
                  : 'tw-hidden tw-overflow-hidden'
              "
              class="tw-flex-col tw-space-y-6"
            >
              <li v-for="navbar in subMenus" :key="navbar">
                <a
                  class="tw-relative tw-py-2 tw-font-inter-medium tw-text-[1rem] tw-leading-[1.125rem] tw-text-black-core/[0.6] group-hover:tw-z-[1] group-hover:tw-text-white md:tw-text-[1.09375rem] md:tw-leading-[1.28125rem] lg:tw-text-[1.188rem] lg:tw-leading-[1.437rem] after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-top-[30px] md:after:tw-top-[35px] after:tw-h-[3px] after:tw-w-full after:tw-origin-bottom-left after:tw-scale-x-0 after:tw-from-[#f2709c] after:tw-to-[#ff9472] after:tw-duration-300 after:tw-bg-gradient-[90deg] hover:tw-from-[#ff9472] hover:tw-to-[#f2709c] hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-bg-gradient-[270.11deg] hover:after:tw-origin-bottom-left hover:after:tw-scale-x-100"
                  :href="navbar.url"
                  :target="navbar.target"
                  >{{ navbar.name }}
                </a>
              </li>
            </ul>
          </li>
        </ul>
        <div v-if="showMenu" class="tw-h-auto tw-w-full tw-bg-white tw-p-4">
          <div className="tw-grid tw-p-3 tw-pb-5">
            <a
              href="/contact"
              className="tw-relative tw-justify-self-center tw-rounded-full tw-border-[1px] tw-border-solid tw-border-transparent tw-bg-gradient-to-r tw-from-[#f2709c] tw-to-[#ff9472] hover:tw-shadow-[inset_2px_1000px_1px_#fff] tw-font-bold tw-text-white "
            >
              <div
                className="tw-py-[0.8rem] tw-px-10 tw-text-lg hover:tw-bg-clip-text hover:tw-bg-gradient-to-r hover:tw-from-[#f2709c] hover:tw-via-[#ff909c] hover:tw-to-[#ff9472] hover:tw-text-transparent tw-inline-block"
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
      showContributionMenu: false,
      showMenu: false,
      animateNavbar: true,
      showNavbar: false,
      navContainerHeight: 100,
      lastScrollY: 0,
      currentRoutePath: this.$router.currentRoute._value.path,
      podcastRef: null,
      navbars: [
        {
          name: "Services",
          url: "/services",
          event: "tap_header_services",
        },
        {
          name: "Portfolio",
          url: "/portfolio",
          event: "tap_header_portfolio",
        },
        {
          name: "Contribution",
          url: "/contributions",
          event: "tap_header_contribution",
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
          id: 7,
          name: "Get Free Consultation",
          url: "/contact",
          className:
            "!tw-m-0 tw-rounded-full tw-py-3 tw-px-3 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0] gradient-btn consultation-btn",
          showContactBtn: true,
          event: "tap_header_cta",
        },
      ],
      subMenus: [
        {
          name: "Open Source",
          target: "_self",
          url: "/contributions",
          event: "tap_header_contributions",
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          target: "_blank",
          event: "tap_header_blog",
        },
        {
          name: "Resources",
          url: "/resources",
          target: "_self",
          event: "tap_header_resources",
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
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
    this.width = window.innerWidth;
  },
  unmounted() {
    if (this.podcastRef) {
      this.podcastRef.removeEventListener("scroll", this.handleScroll);
    }
    window.removeEventListener("scroll", this.handleScroll);
  },
  // inject: ["mixpanel"],
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
