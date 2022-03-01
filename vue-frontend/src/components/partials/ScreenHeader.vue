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
            <li class="nav-item-margin">
              <a class="nav-link normal-text" :href="careerURL">Career</a>
            </li>
            <li class="nav-item-margin">
              <a class="nav-link normal-text" :href="blogsURL" target="_blank"
                >Blogs</a
              >
            </li>
            <li>
              <router-link
                to="/contact"
                class="nav-link start-btn normal-text gradient-border-btn"
              >
                <span
                  ><span class="canopas-gradient-text"
                    >Let's talk business</span
                  >
                </span>
              </router-link>
            </li>
            <li>
              <router-link
                to="/contact"
                class="nav-link start-btn-link normal-text"
                >Let's talk business</router-link
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
      headerLogoImage: headerLogoImage,
      careerURL: "/jobs",
      blogsURL: Config.BLOG_URL,
      navbarSticky: false,
      navbarAnimation: false,
      navContainerHeight: 0,
      lastScrollY: 0,
    };
  },
  components: {},
  mounted() {
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
}

.navbar-nav .nav-item-margin .nav-link {
  margin: 0 20px 0 0;
}

.header-logo-image {
  width: 205px;
  height: 38.5px;
}

.start-btn-link {
  display: block;
}

.start-btn {
  display: none;
}

@include media-breakpoint-up(md) {
  .navbar-nav .nav-item-margin .nav-link {
    margin: 0 40px 0 0;
  }

  .navbar {
    padding: 20px 0;
  }
}

@include media-breakpoint-up(lg) {
  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }

  a.start-btn,
  a.start-btn:focus {
    border: 2px solid transparent !important;
    border-radius: 15px;
    padding: 10px !important;
  }

  .start-btn {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .start-btn-link {
    display: none;
  }

  .gradient-border-btn > span {
    margin: 0 4px;
  }
}
</style>
