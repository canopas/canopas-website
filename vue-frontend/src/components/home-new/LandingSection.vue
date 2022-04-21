<template>
  <div class="container">
    <img :src="backgroundImage" class="background" />
    <div class="con-display">
      <div class="v2-header-2-text text-mg">
        We develop amazing products to help
        <span class="canopas-gradient-text text-weight">Entrepreneurs </span>
        bring their ideas to life.
      </div>
      <div class="img-center"><img :src="boyImage" /></div>
    </div>
    <div class="btn-mg">
      <router-link to="">
        <a class="v2-normal-3-text v2-button btn-center"
          >Let's talk
          <font-awesome-icon class="arrow fa" icon="arrow-right" id="leftArrow"
        /></a>
      </router-link>
    </div>
  </div>
</template>
<script type="module">
import backgroundImage from "@/assets/images/Landing/landingbackground.webp";
import boyImage from "@/assets/images/Landing/Image.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      boyImage: boyImage,
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style lang="scss" scoped>
.img-center {
  margin-left: 5rem;
}

.text-weight {
  font-weight: bold;
}
.background {
  width: 100%;
  height: 100%;
  object-fit: cover;
  position: absolute;
  top: 0;
  left: 0;
  z-index: -5;
  justify-content: flex-start;
}
.v2-header-2-text {
  font-weight: 400;
}

.v2-black-btn {
  box-shadow: none;
  padding: 20px 40px;
}
.btn-mg {
  margin-top: 5rem;
}
.arrow {
  color: white;
}
.v2-black-btn:hover > .arrow {
  color: #3d3d3d;
}

@include media-breakpoint-down(lg) {
  .con-display {
    display: block;
    margin-top: 20px;
  }

  .btn-center {
    display: block;
    margin: 10% auto;
    width: 11rem;
  }
  .v2-header-2-text {
    font-size: 2.125rem;
    line-height: 3rem;
  }

  .img-center {
    width: 25rem;
    display: block;
    margin: 5rem auto;
  }
}

@include media-breakpoint-up(lg) {
  .con-display {
    display: flex;
  }
  .v2-header-2-text {
    font-size: 3.625rem;
    line-height: 4.6875rem;
  }
  .text-mg {
    margin-top: 14%;
    width: 49rem;
  }
  .btn-mg {
    margin-top: 0;
  }
}
</style>
