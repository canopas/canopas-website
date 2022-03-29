<template>
  <div class="container-fluid">
    <div class="container" ref="howDoWeDiffer">
      <div class="description text-left animate__animated" ref="description">
        <span class="header-text canopas-gradient-text text-center"
          >How do we differ from others?</span
        >
        <div class="steve-box gradient-border">
          <span class="normal-text" style="color: #3d3d3d"
            >You gotta start with the customer experience and work backward to
            the technology.</span
          >
          <div class="normal-text orange-gradient-text mt-3 font-weight-bold">
            - Steve Jobs
          </div>
        </div>
      </div>
      <div class="normal-text differ-description">
        By following the advice of him, we focus on user experience first when
        we start working on the project.
        <br /><br />
        We are not bound by technologies. We will learn new technology if it
        significantly improves the performance of your app.
        <br /><br />
        We will solve your tech problems even though we are not THE EXPERT in
        it. And we've done it multiple times in the last seven years.
        <br /><br />
        We design and develop products that make people's lives easier,
        smoother, and better. We help quality-conscious people(or companies)
        design, develop, and expand their business through digital products.
      </div>
      <div class="flex-grid mt-3 pb-5 animate__animated" ref="flexGridIcon">
        <div class="flex-grid-1 normal-text" ref="flexGridOne">
          <div class="text-center text-black">
            <object
              type="image/svg+xml"
              :data="mobileIcon"
              class="pb-4"
            ></object>
            <b
              >Mobile app<br />
              development</b
            >
          </div>
          <div class="text-center text-black">
            <object
              type="image/svg+xml"
              :data="watchIcon"
              class="pb-4"
            ></object>
            <b
              >Watch and TV apps<br />
              development</b
            >
          </div>
          <div class="text-center text-black">
            <object type="image/svg+xml" :data="webIcon" class="pb-4"></object>
            <b>Web app <br />development</b>
          </div>
          <div class="text-center text-black" ref="uiUxDiv">
            <object type="image/svg+xml" :data="uiuxIcon" class="pb-4"></object>
            <b
              >UI/UX<br />
              design</b
            >
          </div>
        </div>
        <div class="flex-grid-2 normal-text" ref="flexGridTwo">
          <div class="text-center text-black">
            <object
              type="image/svg+xml"
              :data="desktopIcon"
              class="pb-4"
            ></object>
            <b>
              Desktop app <br />
              development</b
            >
          </div>
          <div class="text-center text-black">
            <object
              type="image/svg+xml"
              :data="customIcon"
              class="pb-4"
            ></object>
            <b>
              Custom software <br />
              development</b
            >
          </div>
          <div class="text-center text-black">
            <object
              type="image/svg+xml"
              :data="cloudIcon"
              class="pb-4"
            ></object>
            <b
              >Cloud and <br />
              DevOps services</b
            >
          </div>
          <div class="text-center text-black">
            <object type="image/svg+xml" :data="qaIcon" class="pb-4"></object>
            <b
              >QA and <br />
              Testing</b
            >
          </div>
          <div class="text-center text-black">
            <object type="image/svg+xml" :data="seoIcon" class="pb-4"></object>
            <b>
              SEO and <br />
              Marketing</b
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import cloudIcon from "@/assets/images/icon/cloud.svg";
import customIcon from "@/assets/images/icon/custom.svg";
import desktopIcon from "@/assets/images/icon/desktop.svg";
import mobileIcon from "@/assets/images/icon/mobile.svg";
import qaIcon from "@/assets/images/icon/qa.svg";
import scrollArrowToUpIcon from "@/assets/images/icon/scroll-arrow-to-up.svg";
import seoIcon from "@/assets/images/icon/seo.svg";
import uiuxIcon from "@/assets/images/icon/uiux.svg";
import watchIcon from "@/assets/images/icon/watch.svg";
import webIcon from "@/assets/images/icon/web.svg";

export default {
  data() {
    return {
      cloudIcon: cloudIcon,
      customIcon: customIcon,
      desktopIcon: desktopIcon,
      mobileIcon: mobileIcon,
      qaIcon: qaIcon,
      scrollArrowToUpIcon: scrollArrowToUpIcon,
      seoIcon: seoIcon,
      uiuxIcon: uiuxIcon,
      watchIcon: watchIcon,
      webIcon: webIcon,
    };
  },
  mounted() {
    if (window.innerWidth >= 768 && window.innerWidth <= 992) {
      this.$refs.flexGridTwo.prepend(this.$refs.uiUxDiv);
    } else {
      this.$refs.flexGridOne.append(this.$refs.uiUxDiv);
    }
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      let data = {
        name: this.$refs.howDoWeDiffer,
        animation: "",
        childRef: [
          {
            name: this.$refs.description,
            animation: "animate__fadeIn",
          },
          {
            name: this.$refs.flexGridIcon,
            animation: "animate__fadeIn",
          },
        ],
      };
      if (data) {
        this.$emit("add-animation", data);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin: auto;
}

.container-fluid {
  background: rgba(255, 148, 114, 0.05);
  z-index: 100;
}

.description {
  width: auto;
  padding-top: 75px;
  margin: auto;
  display: flex;
  flex-direction: column;
}

.steve-box {
  margin: 32px 8px;
  line-height: 2;
  padding: 24px;
  border-radius: 14px;
  text-align: left;
  border: 1px solid transparent !important;
  position: relative;
}

.gradient-border {
  background: -webkit-linear-gradient(
        rgba(255, 250, 248, 255),
        rgba(255, 250, 248, 255)
      )
      padding-box,
    linear-gradient(to bottom, #ff9472, #f2709c) border-box;
  border: 2px solid transparent;
  border-radius: 15px;
  display: inline-block;
}

.orange-gradient-text {
  background: linear-gradient(180deg, #ff9472 0%, #f2709c 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.flex-grid {
  display: flex;
  flex-direction: column;
}

.flex-grid-1,
.flex-grid-2 {
  display: flex;
  flex-wrap: wrap;
  justify-content: unset;
}

.flex-grid-1 > div,
.flex-grid-2 > div {
  padding: 20px;
  width: 50%;
  display: flex;
  flex-direction: column;
}

.flex-grid-1 > div > object,
.flex-grid-2 > div > object {
  height: 120px;
}

.flex-grid-1 > div > b,
.flex-grid-2 > div > b {
  font-size: 1.125rem;
}

@include media-breakpoint-up(md) {
  .container-fluid {
    margin-top: -150px;
  }

  .description {
    width: 50%;
  }

  .steve-box {
    padding: 32px;
  }

  .flex-grid-1 > div,
  .flex-grid-2 > div {
    width: 33%;
  }

  .description {
    padding-top: 126px;
  }
}

@include media-breakpoint-up(lg) {
  .container {
    padding: 0 10%;
  }

  .steve-box {
    padding: 48px;
  }

  .flex-grid-1 > div,
  .flex-grid-2 > div {
    width: 20%;
  }

  .flex-grid-1 > div > b,
  .flex-grid-2 > div > b {
    font-size: 1.25rem;
  }

  .flex-grid-1,
  .flex-grid-2 {
    justify-content: center;
  }

  .description {
    padding-top: 112px;
  }
}

@include media-breakpoint-up(xl) {
  .container {
    padding: 0;
  }
  .differ-description {
    padding: 0 10%;
  }

  .description {
    padding-top: 116px;
  }
}
</style>
