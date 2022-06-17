<template>
  <div class="nav-container" :style="{ height: navContainerHeight + 'px' }">
    <nav
      class="navbar navbar-expand-md navbar-light main-header"
      :class="{
        'navbar-sticky': navbarSticky,
        'navbar-animation': navbarAnimation,
      }"
      ref="mainHeader"
    >
      <div class="container">
        <router-link to="/" replace>
          <div class="navbar-brand">
            <img
              :src="headerLogoImage"
              class="header-logo-image mt-1"
              alt="canopas-logo"
            />
          </div>
        </router-link>

        <div class="navbar-collapse">
          <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
            <li
              v-for="navbar in navbars"
              :key="navbar"
              class="my-2 nav-item-margin"
            >
              <router-link
                v-if="!navbar.target"
                class="v2-normal-3-text"
                :class="[
                  navbar.className,
                  currentRoutePath == '/portfolio/' + id ? 'hide-links' : '',
                  currentRoutePath == navbar.url
                    ? navbar.showContactBtn
                      ? 'v2-button v2-button-active'
                      : 'active-tab'
                    : '',
                ]"
                :to="navbar.url"
                >{{ navbar.name }}</router-link
              >

              <a
                v-else
                class="v2-normal-3-text"
                :class="[
                  navbar.className,
                  currentRoutePath == '/portfolio/' + id ? 'hide-links' : '',
                  currentRoutePath == navbar.url
                    ? navbar.showContactBtn
                      ? 'v2-button v2-button-active'
                      : 'active-tab'
                    : '',
                ]"
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
import headerLogoImage from "@/assets/images/logo/logo-header.svg";
import Config from "@/config.js";

export default {
  data() {
    return {
      id: this.$route.params.id,
      headerLogoImage: headerLogoImage,
      navbarSticky: false,
      navbarAnimation: false,
      navContainerHeight: 133,
      lastScrollY: 0,
      currentRoutePath: this.$router.currentRoute._value.path,
      navbars: [
        {
          name: "Home",
          url: "/",
          className: "nav-link is-animation-tab",
          target: false,
        },
        {
          name: "Career",
          url: "/jobs",
          className: "nav-link is-animation-tab",
          target: false,
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          className: "nav-link is-animation-tab",
          target: true,
        },
        // {
        //   name: "Portfolio",
        //   url: "",
        //   className: "nav-link is-animation-tab",
        //   target: false,
        //   isActive: false,
        // },

        {
          name: "Let's talk",
          url: "/contact",
          className: "me-0 nav-link lets-button-link is-animation-tab",
          target: false,
        },
        {
          name: "Let's talk",
          url: "/contact",
          className: "me-0 v2-button lets-button",
          target: false,
          showContactBtn: true,
        },
      ],
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
  },
};
</script>

<style lang="scss" scoped>
.nav-container {
  background: #fff;
  position: relative;
}

.hide-links {
  visibility: hidden;
}

.navbar {
  padding: 10px 2%;
  width: 100%;
  background: #fff;
  position: absolute;
  left: 0;
  bottom: 0;
  z-index: 1;
  transition: all 0.6s ease-in-out;
}

.navbar-sticky {
  position: fixed;
  left: unset;
  bottom: unset;
  box-shadow: 0 13px 35px -12px rgba(35, 35, 35, 0.15);
}

.navbar-animation {
  animation: menu_sticky 0.6s ease-in-out;
}

@keyframes menu_sticky {
  0% {
    top: -120px;
    opacity: 0;
  }

  100% {
    top: 0;
    opacity: 1;
  }
}

.navbar-nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  flex-wrap: wrap;
}

.navbar-nav .nav-item-margin .nav-link {
  margin: 0 30px 0 0;
  color: rgba(61, 61, 61, 0.8);
  padding: 0;
}

.header-logo-image {
  width: 205px;
  height: 38.5px;
}

.is-animation-tab,
.active-tab {
  display: inline-block;
  position: relative;
}

.active-tab:after,
.is-animation-tab:after {
  content: "";
  position: absolute;
  width: 50%;
  transform: scaleX(0);
  height: 2px;
  bottom: 0;
  left: 0;
  background-color: #3d3d3d;
}

.active-tab:after {
  transform: scaleX(1);
}

.is-animation-tab:after {
  transform-origin: bottom left;
  transition: transform 0.2s ease-out;
}

.is-animation-tab:hover:after {
  transform: scaleX(1);
  transform-origin: bottom left;
}

.v2-button {
  padding: 0.4rem 1.5rem;
  background-color: #3d3d3d;
  color: #fff !important;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.5);
}

.v2-button:hover {
  background-color: #fff;
  color: #3d3d3d !important;
}

.v2-button-active {
  background-color: #fff;
  color: #3d3d3d !important;
}

.v2-button-active:hover {
  background-color: #3d3d3d;
  color: #fff !important;
}

.lets-button-link {
  display: block;
}

.lets-button {
  display: none;
}

@include media-breakpoint-up(md) {
  .navbar {
    padding: 20px 0;
  }
}

@include media-breakpoint-up(lg) {
  .lets-button {
    display: block;
  }

  .lets-button-link {
    display: none;
  }
}

@include media-breakpoint-up(xl) {
  .navbar-nav .nav-item-margin .nav-link {
    margin: 0 40px 0 0;
  }
}
</style>
