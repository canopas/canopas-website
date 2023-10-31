<template>
  <header class="relative z-10 font-inter-medium">
    <div class="sticky top-0" :style="{ height: navContainerHeight + 'px' }">
      <nav
        class="ease z-[1] w-full bg-white py-5 tracking-[0] text-black-core/[.87] transition-all duration-500 md:px-0 md:py-4"
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
                  v-if="!navbar.target"
                  :to="navbar.url"
                  @click.native="$mixpanel.track(navbar.event)"
                  class="group relative mr-5 inline-block sm:mr-[30px] lg:mr-2 xl:mr-[30px]"
                  :class="[
                    navbar.className
                      ? navbar.className
                      : navbar.class +
                        ' after:absolute after:bottom-0 after:left-0 after:top-[27px] after:h-[3px] after:w-full  after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100 ',
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
                  class="relative mr-5 inline-block after:absolute after:bottom-0 after:left-0 after:top-[27px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100 sm:mr-[30px] lg:mr-2 xl:mr-[30px]"
                  :href="navbar.url"
                  :target="navbar.target"
                  >{{ navbar.name }}</a
                >

                <ul
                  v-show="showContributionMenu"
                  @mouseleave="showContributionMenu = false"
                  class="absolute w-max flex-col space-y-6 rounded-[5px] border bg-white !px-6 py-[1.5rem] shadow-md lg:top-[4.5rem] lg:left-[169px] xl:left-[296px] 2xl:left-[475px]"
                >
                  <li v-for="navbar in subMenus" :key="navbar">
                    <a
                      class="relative py-4 font-inter-medium text-[1rem] leading-[1.125rem] text-black-core/[0.6] md:text-[1.09375rem] md:leading-[1.28125rem] lg:text-[1.188rem] lg:leading-[1.437rem] after:absolute after:bottom-0 after:left-0 after:top-[42px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100"
                      :href="navbar.url"
                      :target="navbar.target"
                      >{{ navbar.name }}
                    </a>
                  </li>
                </ul>
              </li>
            </ul>
          </div>
          <div class="flex h-6 w-6 items-end lg:hidden">
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
            <div class="mt-[-3px] h-6 w-6">
              <Icon
                name="fa6-solid:xmark"
                class="h-full w-full"
                @click="showMenu = false"
              />
            </div>
          </div>
        </div>
        <ul
          class="mt-8 flex h-[50%] flex-col justify-start overflow-y-scroll text-[1rem] leading-[1.125rem] sm:h-[45%] md:text-[1.09375rem] md:leading-[1.28125rem] lg:ml-auto lg:items-center lg:text-[1.1875rem] lg:leading-[1.4375rem]"
        >
          <li
            v-for="navbar in navbars.slice(0, navbars.length - 1)"
            :key="navbar"
            class="my-2 ml-1 py-2 lg:ml-0"
          >
            <nuxt-link
              v-if="!navbar.target"
              :to="navbar.url == '/contributions' ? '' : navbar.url"
              @click.native="$mixpanel.track(navbar.event)"
              class="group relative mr-5 inline-block sm:mr-[30px] lg:mr-5"
              :class="[
                navbar.className
                  ? navbar.className
                  : navbar.class +
                    ' after:absolute after:bottom-0 after:left-0 after:top-[27px] after:h-[3px] after:w-full  after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100',
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
              class="relative mr-5 inline-block after:absolute after:bottom-0 after:left-0 after:top-[27px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100 sm:mr-[30px]"
              :href="navbar.url"
              :target="navbar.target"
              >{{ navbar.name }}</a
            >
            <ul
              :class="
                showContributionMenu && navbar.name == 'Contribution'
                  ? 'mt-8 w-full bg-white px-6'
                  : 'hidden overflow-hidden'
              "
              class="flex-col space-y-6"
            >
              <li v-for="navbar in subMenus" :key="navbar">
                <a
                  class="relative py-2 font-inter-medium text-[1rem] leading-[1.125rem] text-black-core/[0.6] group-hover:z-[1] group-hover:text-white md:text-[1.09375rem] md:leading-[1.28125rem] lg:text-[1.188rem] lg:leading-[1.437rem] after:absolute after:bottom-0 after:left-0 after:top-[30px] md:after:top-[35px] after:h-[3px] after:w-full after:origin-bottom-left after:scale-x-0 after:from-[#f2709c] after:to-[#ff9472] after:duration-300 after:bg-gradient-[90deg] hover:from-[#ff9472] hover:to-[#f2709c] hover:bg-clip-text hover:text-transparent hover:bg-gradient-[270.11deg] hover:after:origin-bottom-left hover:after:scale-x-100"
                  :href="navbar.url"
                  :target="navbar.target"
                  >{{ navbar.name }}
                </a>
              </li>
            </ul>
          </li>
        </ul>
        <div v-if="showMenu" class="h-auto w-full bg-white p-4">
          <div className="grid p-3 pb-5">
            <a
              href="/contact"
              className="relative justify-self-center rounded-full border-[1px] border-solid border-transparent bg-gradient-to-r from-[#f2709c] to-[#ff9472] hover:shadow-[inset_2px_1000px_1px_#fff] font-bold text-white "
            >
              <div
                className="py-[0.8rem] px-10 text-lg hover:bg-clip-text hover:bg-gradient-to-r hover:from-[#f2709c] hover:via-[#ff909c] hover:to-[#ff9472] hover:text-transparent inline-block"
              >
                Get Free Consultation
              </div>
            </a>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script type="module">
import Config from "@/config.js";
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
            "m-0 rounded-full py-3 px-3 font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-medium !tracking-[0] gradient-btn consultation-btn",
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
  },
};
</script>
