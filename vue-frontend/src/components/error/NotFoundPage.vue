<template>
  <div class="not-found-body-wrapper">
    <div class="middle">
      <div class="center">
        <div class="image-404">
          <img :src="firstErrorLetter" alt="404" />
          <img :src="middleErrorLetter" class="middle-logo" alt="404" />
          <img :src="lastErrorLetter" alt="404" />
        </div>
        <div class="middle-text">
          The page you’re looking for was moved, renamed or might never existed.
        </div>
        <a
          type="button"
          class="back-btn gradient-border-btn"
          onclick="window.location.replace('/')"
        >
          <font-awesome-icon class="arrow" icon="arrow-left" id="leftArrow" />
          <span class="black">Back to Home page</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script type="module">
import firstErrorLetter from "@/assets/images/logo/404page_4_1.svg";
import middleErrorLetter from "@/assets/images/logo/canopas-icon.svg";
import lastErrorLetter from "@/assets/images/logo/404page_4_2.svg";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      firstErrorLetter: firstErrorLetter,
      middleErrorLetter: middleErrorLetter,
      lastErrorLetter: lastErrorLetter,
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style scoped>
.not-found-body-wrapper {
  display: table;
  top: 0;
  left: 0;
  height: 100vh;
  width: 100%;
  position: absolute;
}

.middle {
  display: table-cell;
  vertical-align: middle;
}

.center {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 50%;
  margin-left: auto;
  margin-right: auto;
}

.image-404 {
  margin: 0 auto;
  text-align: center;
  display: flex;
  flex-direction: row;
}

.middle-text {
  margin-top: 2rem;
  text-align: center;
  font-size: 1.4rem !important;
  line-height: 2rem !important;
}

.middle-logo {
  margin: 1rem;
  width: 6.5rem;
}

.back-btn,
.back-btn:hover,
.back-btn:focus {
  width: auto;
  padding: 1rem 1.5rem !important;
  margin: 7% auto;
  border: 1px solid transparent;
  border-radius: 10px;
  display: flex;
  justify-content: space-around !important;
  align-items: center;
}

.back-btn > span {
  text-align: center !important;
  font-weight: 700;
  font-size: 1.1rem !important;
  line-height: 0.5rem;
}

.black {
  color: rgba(61, 61, 61, 1);
}

.back-btn > i {
  margin-right: 1rem;
}

.not-found-footer {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
}

.arrow {
  color: #f2709c;
  font-size: 1.5rem;
  margin-right: 16px;
}

.gradient-border-btn {
  background: linear-gradient(rgba(255, 255, 255, 0), rgba(255, 255, 255, 0)),
    linear-gradient(to top, #ff9472, #f2709c);
  box-shadow: 2px 1000px 1px #fff inset !important;
}
</style>
