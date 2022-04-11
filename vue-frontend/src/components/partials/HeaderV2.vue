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
      <div class="container col-md">
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
            <li class="nav-item-margin">
              <a class="nav-link v2-title-3-text bottom-border" target="_blank"
                >Home</a
              >
            </li>
            <li class="nav-item-margin v2-title-3-text">
              <a class="nav-link" :href="careerURL">Career</a>
            </li>
            <li class="nav-item-margin v2-title-3-text">
              <a class="nav-link" :href="blogsURL" target="_blank">Blogs</a>
            </li>
            <li class="nav-item-margin v2-title-3-text">
              <a class="nav-link" target="_blank">Portfolio</a>
            </li>
            <router-link :to="contactURL">
              <li class="nav-item-margin round-btn v2-title-3-text">
                <a class="anchor btn-border" target="_blank">Let's Talk</a>
              </li></router-link
            >
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
      headerLogoImage: headerLogoImage,
      careerURL: "/jobs",
      contactURL: "/contact",
      blogsURL: Config.BLOG_URL,
      navbarSticky: false,
      navbarAnimation: false,
      navContainerHeight: 0,
      lastScrollY: 0,
      activeCareerPath: false,
      activeContactPath: false,
    };
  },
  components: {},
  mounted() {
    let currentRoutePath = this.$router.currentRoute._value.path;

    this.activeCareerPath = currentRoutePath == this.careerURL;

    this.activeContactPath = currentRoutePath == this.contactURL;

    window.addEventListener("scroll", this.handleScroll);
    this.navContainerHeight = this.$refs.mainHeader.clientHeight + 30;
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  computed: {},
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

a {
  color: #000;
}

.navbar-animation {
  animation: menu_sticky 0.6s ease-in-out;
}

.bottom-border::after {
  content: "";
  display: block;
  position: relative;
  width: 50.4%;
  border-top: 3px solid black;
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
}

.navbar-nav .nav-item-margin .nav-link {
  margin: 0 20px 0 0;
  color: rgb(0, 0, 0);
}

.header-logo-image {
  width: 205px;
  height: 38.5px;
}

@include media-breakpoint-up(md) {
  .navbar-nav .nav-item-margin .nav-link {
    margin: 0 20px 0 0;
  }

  .navbar {
    padding: 20px 0;
  }
  .col-md {
    flex-direction: column;
  }
}

@include media-breakpoint-up(lg) {
  .round-btn {
    background-color: #000;
    border-radius: 25px;
    padding: 10px 25px;
    color: #fff;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  }
  .round-btn:hover > a {
    color: black;
  }
  .round-btn:hover {
    background-color: white;
    color: black;
  }
  .anchor {
    color: white;
  }
  .col-md {
    flex-direction: row;
  }
}
@include media-breakpoint-up(xl) {
  .round-btn {
    padding: 10px 40px;
  }
  .col-md {
    flex-direction: row;
  }
}
</style>
