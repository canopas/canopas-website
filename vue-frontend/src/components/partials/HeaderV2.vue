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
      <div class="container col-flow">
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
            <li class="nav-item-margin list-flow">
              <a
                class="nav-link v2-title-3-text"
                v-for="item in items"
                :key="item.id"
                :style="[
                  currentRoutePath == item.url
                    ? {
                        textDecoration: `underline !important`,
                        'text-underline-offset': '0.3rem',
                      }
                    : { textDecoration: `none` },

                  currentRoutePath == contactPath
                    ? isRemoveActiveUnderline == true
                    : isRemoveActiveUnderline == false
                    ? {
                        textDecoration: `none`,
                      }
                    : {
                        textDecoration: `underline !important`,
                        'text-underline-offset': '0.3rem',
                      },
                ]"
                :class="[currentRoutePath == contactPath ? item.btn : item.btn]"
                :href="item.url"
                v-html="item.name"
                :target="item.target ? '_blank' : ''"
              >
              </a>
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
      navbarSticky: false,
      navbarAnimation: false,
      navContainerHeight: 133,
      lastScrollY: 0,
      currentRoutePath: this.$router.currentRoute._value.path,
      contactPath: "/contact",
      isRemoveActiveUnderline: false,
      items: [
        {
          name: `<span class="v2-underline-text">Ho</span>me`,
          url: "/",
          btn: "",
          target: false,
        },
        {
          name: "Career",
          url: "/jobs",
          btn: "",
          target: false,
        },
        {
          name: "Blog",
          url: Config.BLOG_URL,
          btn: "",
          target: true,
        },
        {
          name: "PortFolio",
          url: "",
          btn: "",
          target: false,
        },
        {
          name: "Let's Talk",
          url: "/contact",
          btn: "v2-button",
          target: false,
        },
      ],
    };
  },
  components: {},
  mounted() {
    if (window.innerWidth >= 992 && currentRoutePath == "/contact") {
      isRemoveActiveUnderline == true;
    }

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
.v2-btn {
  border: none;
}
.nav-link {
  margin: 0 3px 0 0;
}
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

.v2-button:hover {
  background-color: #fff;
}
.list-flow {
  display: flex;
  flex-shrink: 0;
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

.header-logo-image {
  width: 205px;
  height: 38.5px;
}

@include media-breakpoint-up(md) {
  .navbar {
    padding: 20px 0;
  }
  .col-flow {
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
  }
}
@include media-breakpoint-up(lg) {
  .col-flow {
    flex-direction: row;
  }
  .v2-button {
    display: block;
    padding: 10px 30px;
  }

  .v2-button {
    border: 1px solid transparent;
    background-color: #3d3d3d;
    color: #fff !important;
    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.5);
  }
  .v2-button:hover {
    background-color: #fff;
    color: #3d3d3d !important ;
    border: 1px solid #3d3d3d;
  }
}
</style>
