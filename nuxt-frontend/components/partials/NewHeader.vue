<template>
  <header class="relative z-[100] font-inter-medium">
    <div class="sticky top-0" :style="{ height: navContainerHeight + 'px' }">
      <nav
        class="ease w-full bg-white py-5 tracking-[0] text-black-core/[.87] transition-all duration-500 md:px-0 md:py-4"
        :class="[
          showNavbar
            ? 'fixed shadow-[0_13px_35px_-12px_rgba(35,35,35,0.15)]'
            : '',
        ]"
        :style="{
          transform: animateNavbar
            ? 'translateY(0)'
            : 'translateY(-' + navContainerHeight + 'px)',
        }"
        ref="mainHeader"
      >
        <div class="container flex flex-row flex-wrap items-center">
          <nuxt-link to="/" replace>
            <div
              class="mr-4 mt-2 whitespace-nowrap text-[1.25rem] no-underline"
            >
              <img
                src="@/assets/images/logo/logo-header.svg"
                class="h-[38.5px] w-[205px]"
                alt="canopas-logo"
              />
            </div>
          </nuxt-link>
          <div class="relative flex grow basis-[auto] items-center">
            <ul
              class="hidden w-full flex-wrap justify-evenly text-[1rem] leading-[1.125rem] md:text-[1.09375rem] md:leading-[1.28125rem] lg:flex lg:flex-row lg:items-center xl:justify-end xl:text-[1.1875rem] xl:leading-[1.4375rem]"
              @mouseleave="showContributionMenu = false"
            >
              <li v-for="navbar in navbars" :key="navbar" class="my-2 py-2">
                <nuxt-link
                  @mouseover="
                    showContributionMenu = navbar.name == 'Contribution'
                  "
                  :to="navbar.url"
                  @click.native="$mixpanel.track(navbar.event)"
                  class="flex gap-2 group relative mr-5 sm:mr-[30px] lg:mr-2 xl:mr-[30px]"
                  :class="[
                    isActiveRoute(navbar.name)
                      ? 'router-link-exact-active'
                      : '',
                    navbar.className
                      ? navbar.className
                      : ' main-menu after:absolute after:bottom-0 after:left-0 after:top-[27px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100 ',
                  ]"
                >
                  <span
                    :class="[
                      navbar.id == 7 ? 'hover:v2-canopas-gradient-text' : '',
                    ]"
                    >{{ navbar.name }}</span
                  >
                  <span
                    v-if="navbar.subMenus"
                    class="flex items-center transition ease-in-out duration-500"
                    :class="showContributionMenu ? 'rotate-180' : ''"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="13"
                      height="7"
                      viewBox="0 0 13 7"
                      fill="none"
                    >
                      <path
                        d="M6.5 6.5L0.5 0.5H12.5L6.5 6.5Z"
                        fill="url(#paint0_linear_8576_22013)"
                      />
                      <defs>
                        <linearGradient
                          id="paint0_linear_8576_22013"
                          x1="13.76"
                          y1="3.49999"
                          x2="-1.17743"
                          y2="6.50015"
                          gradientUnits="userSpaceOnUse"
                        >
                          <stop stop-color="#F2709C" />
                          <stop offset="1" stop-color="#FF835B" />
                        </linearGradient>
                      </defs>
                    </svg>
                  </span>
                </nuxt-link>

                <ul
                  v-show="showContributionMenu && navbar.subMenus"
                  @mouseleave="showContributionMenu = false"
                  class="absolute w-max flex-col space-y-6 rounded-[5px] border bg-white !px-6 py-[1.5rem] shadow-md lg:top-[4.5rem] lg:left-[169px] xl:left-[296px] 2xl:left-[475px]"
                >
                  <li v-for="navbar in navbar.subMenus" :key="navbar">
                    <nuxt-link
                      class="sub-menu relative py-4 font-inter-medium text-[1rem] leading-[1.125rem] text-black-core/[0.6] md:text-[1.09375rem] md:leading-[1.28125rem] lg:text-[1.188rem] lg:leading-[1.437rem] after:absolute after:bottom-0 after:left-0 after:top-[42px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100"
                      :href="navbar.url"
                      target="_self"
                      >{{ navbar.name }}
                    </nuxt-link>
                  </li>
                </ul>
              </li>
            </ul>
          </div>
          <div class="flex h-6 w-6 items-end lg:hidden cursor-pointer">
            <Icon
              name="fa6-solid:bars"
              class="ml-0 mt-[0.20rem] h-full w-full"
              @click="showMenu = true"
            />
          </div>
        </div>
      </nav>
    </div>
    <div
      :class="showMenu ? 'translate-x-0' : 'translate-x-full'"
      class="fixed right-0 top-0 h-screen w-full overflow-y-hidden bg-white py-4 duration-500 ease-in-out md:py-2 lg:hidden"
    >
      <div class="container">
        <div class="sticky top-0 w-full md:pt-3">
          <div class="flex items-center justify-between">
            <div class="whitespace-nowrap text-[1.25rem]">
              <nuxt-link to="/" replace>
                <div
                  class="mr-4 mt-2 whitespace-nowrap text-[1.25rem] no-underline"
                >
                  <img
                    src="@/assets/images/logo/logo-header.svg"
                    class="h-[38.5px] w-[205px]"
                    alt="canopas-logo"
                  />
                </div>
              </nuxt-link>
            </div>
            <div class="mt-[-3px] h-6 w-6 cursor-pointer">
              <Icon
                name="fa6-solid:xmark"
                class="h-full w-full"
                @click="showMenu = false"
              />
            </div>
          </div>
        </div>
        <ul
          class="mt-8 flex h-[50%] flex-col justify-start overflow-y-scroll scrollbar-hidden text-[1rem] leading-[1.125rem] sm:h-[45%] md:text-[1.09375rem] md:leading-[1.28125rem] lg:ml-auto lg:items-center lg:text-[1.1875rem] lg:leading-[1.4375rem]"
        >
          <li
            v-for="navbar in navbars.slice(0, navbars.length - 1)"
            :key="navbar"
            class="my-2 ml-1 py-2 lg:ml-0"
          >
            <nuxt-link
              :to="navbar.url == '/contributions' ? '' : navbar.url"
              @click.native="$mixpanel.track(navbar.event)"
              class="flex gap-2 group relative mr-5 w-fit sm:mr-[30px] lg:mr-5 cursor-pointer"
              :class="[
                isActiveRoute(navbar.name) ? 'router-link-exact-active' : '',
                navbar.className
                  ? navbar.className
                  : ' main-menu after:absolute after:bottom-0 after:left-0 after:top-[27px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100',
              ]"
              @click="
                navbar.name == 'Contribution'
                  ? (showContributionMenu = !showContributionMenu)
                  : ''
              "
            >
              <span>{{ navbar.name }}</span>
              <span
                v-if="navbar.subMenus"
                class="flex items-center transition ease-in-out duration-500"
                :class="showContributionMenu ? 'rotate-180' : ''"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="13"
                  height="7"
                  viewBox="0 0 13 7"
                  fill="none"
                >
                  <path
                    d="M6.5 6.5L0.5 0.5H12.5L6.5 6.5Z"
                    fill="url(#paint1_linear_8576_22013)"
                  />
                  <defs>
                    <linearGradient
                      id="paint1_linear_8576_22013"
                      x1="13.76"
                      y1="3.49999"
                      x2="-1.17743"
                      y2="6.50015"
                      gradientUnits="userSpaceOnUse"
                    >
                      <stop stop-color="#F2709C" />
                      <stop offset="1" stop-color="#FF835B" />
                    </linearGradient>
                  </defs>
                </svg>
              </span>
            </nuxt-link>
            <ul
              :class="
                showContributionMenu && navbar.name == 'Contribution'
                  ? 'mt-8 w-full bg-white px-6'
                  : 'hidden overflow-hidden'
              "
              class="flex-col space-y-6"
            >
              <li v-for="navbar in navbar.subMenus" :key="navbar">
                <nuxt-link
                  class="sub-menu relative py-2 font-inter-medium text-[1rem] leading-[1.125rem] text-black-core/[0.6] group-hover:z-[1] group-hover:text-white md:text-[1.09375rem] md:leading-[1.28125rem] lg:text-[1.188rem] lg:leading-[1.437rem] after:absolute after:bottom-0 after:left-0 after:top-[30px] md:after:top-[35px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100"
                  :href="navbar.url"
                  target="_self"
                  >{{ navbar.name }}
                </nuxt-link>
              </li>
            </ul>
          </li>
        </ul>
        <div v-if="showMenu" class="h-auto w-full bg-white p-4">
          <div className="grid p-3 pb-5">
            <nuxt-link
              class="gradient-btn primary-btn"
              to="/contact"
              @click.native="$mixpanel.track('tap_landing_cta')"
            >
              <span class="sub-h3-semibold lg:sub-h1-semibold"
                >Get Free Consultation
              </span>
            </nuxt-link>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script type="module">
export default {
  data() {
    return {
      id: this.$route.params.id,
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
          subMenus: [
            {
              name: "Open Source",
              url: "/contributions",
              event: "tap_header_contributions",
            },
            {
              name: "Blog",
              url: "/blog",
              event: "tap_header_blog",
            },
            {
              name: "Resources",
              url: "/resources",
              event: "tap_header_resources",
            },
          ],
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
            "m-0 px-3 pt-2.5  gradient-btn primary-btn sub-h3-semibold lg:sub-h1-semibold",
          showContactBtn: true,
          event: "tap_header_cta",
        },
      ],
    };
  },
  mounted() {
    this.navContainerHeight = this.$refs.mainHeader.clientHeight;
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
      this.showNavbar = scrollUp || (this.showNavbar && pageYOffset > 0);
      this.animateNavbar = scrollUp || pageYOffset < 100;
      this.lastScrollY = scrollY;
    },
    isActiveRoute(name) {
      const navbar = this.navbars.find((item) => item.name === name);
      if (!navbar) {
        return false;
      }

      // Check if any child URL matches
      if (navbar.subMenus) {
        return navbar.subMenus.some(
          (subMenu) => this.$route.path === subMenu.url,
        );
      }

      if (name == "Portfolio" && this.$route.path.startsWith("/portfolio")) {
        return true;
      }

      return false;
    },
  },
};
</script>

<style lang="postcss" scoped>
.main-menu.router-link-exact-active {
  @apply from-[#ff9472] to-[#f2709c] bg-clip-text text-transparent bg-gradient-[270.11deg];
}

.sub-menu.router-link-exact-active {
  @apply from-[#ff9472] to-[#f2709c] bg-clip-text text-transparent bg-gradient-[270.11deg];
}

.scrollbar-hidden {
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}

/* Hide scrollbar for Chrome, Safari and Opera */
.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
</style>
