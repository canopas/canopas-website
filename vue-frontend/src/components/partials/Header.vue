<template>
  <nav class="navbar navbar-expand-md navbar-light main-header" id="mainHeader">
    <div class="container">
      <router-link to="/" replace>
        <div class="navbar-brand">
          <img
            :src="headerLogoImage"
            class="header-logo-image"
            alt="canopas-logo"
          />
        </div>
      </router-link>
      <div class="navbar-collapse">
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
          <li class="nav-item-margin">
            <a class="nav-link normal-text" :href="careerURL" target="_blank"
              >Career</a
            >
          </li>
          <li class="nav-item-margin">
            <a
              class="nav-link blogs normal-text"
              :href="blogsURL"
              target="_blank"
            >
              Blogs</a
            >
          </li>
          <li>
            <router-link
              to="/contact"
              class="nav-link start-btn normal-text gradient-border-btn"
            >
              <span
                ><span class="canopas-gradient-text">Get started &nbsp;</span>
                <font-awesome-icon class="fa-arrow-right" icon="arrow-right" />
              </span>
            </router-link>
          </li>
          <li>
            <a
              class="nav-link start-btn-link normal-text"
              :href="getStartedURL"
              target="_self"
            >
              Get started</a
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import headerLogoImage from "@/assets/images/logo/logo-header.svg";

export default {
  data() {
    return {
      headerLogoImage: headerLogoImage,
      careerURL: process.env.JOBS_URL,
      blogsURL: process.env.BLOGS_URL,
      getStartedURL: "/content",
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style scoped>
.navbar {
  padding: 22px 3%;
}

.navbar-nav {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
}

.navbar-nav .nav-item-margin .nav-link {
  margin: 0 40px 0 0;
}

.gradient-border-btn {
  background: linear-gradient(rgba(255, 255, 255, 0), rgba(255, 255, 255, 0)),
    linear-gradient(to top, #ff9472, #f2709c);
  box-shadow: 2px 1000px 1px #fff inset !important;
}

.header-logo-image {
  width: 205px;
  height: 40px;
}

.start-btn-link {
  display: block;
}

.start-btn {
  display: none;
}

@media (min-width: 768px) {
  .navbar-nav .nav-item-margin .nav-link {
    margin: 0 40px 0 0;
  }

  .navbar {
    padding: 45px 10%;
    transition: all 0.6s ease-in-out;
    background: #fff;
    z-index: 10000;
  }
}

@media (min-width: 992px) {
  .navbar {
    padding: 45px 14%;
  }

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

  .fa-arrow-right {
    color: #f2709c;
  }
}
</style>
