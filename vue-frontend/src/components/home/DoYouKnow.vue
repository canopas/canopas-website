<template>
  <div>
    <div class="m-lr-15 flex-row que-section">
      <div class="image que-image"></div>
      <div class="text-div do-you-know-text">
        <h4 class="header-text pb-30">Do you know ?</h4>

        <span class="normal-text pink-text mt-20">
          We’ve worked with just <b>10 clients</b> in the last <b>5 years</b>.
          WHY?
        </span>
        <br /><br />
        <span class="normal-text mt-20">
          Because 8 of them have been offering(committed to) us repeat business
          since we started working for them. Think for a moment, Isn't it the
          main reason why you're here.</span
        >
      </div>
    </div>
    <div class="light-orange-bg exclude-parent-margin">
      <div class="m-lr-15">
        <div class="description text-left">
          <div
            class="header-text canopas-gradient-text line-header-text text-center"
          >
            How do we differ from others?
          </div>
          <div class="steve-box gradient-border wow">
            <span class="normal-text" style="color: #3d3d3d"
              >You gotta start with the customer experience and work backward to
              the technology.</span
            >
            <p class="normal-text orange-gradient-text mt-20 font-weight-bold">
              - Steve Jobs
            </p>
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
        <div class="flex-grid">
          <div class="flex-grid-1 wow">
            <div class="text-center normal-text text-black">
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
            <div class="text-center normal-text text-black">
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
            <div class="text-center normal-text text-black">
              <object
                type="image/svg+xml"
                :data="webIcon"
                class="pb-4"
              ></object>
              <b>Web app <br />development</b>
            </div>
            <div class="text-center normal-text text-black" id="ui-ux-div">
              <object
                type="image/svg+xml"
                :data="uiuxIcon"
                class="pb-4"
              ></object>
              <b
                >UI/UX<br />
                design</b
              >
            </div>
          </div>
          <div class="flex-grid-2 wow">
            <div class="text-center normal-text text-black">
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
            <div class="text-center normal-text text-black">
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
            <div class="text-center normal-text text-black">
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
            <div class="text-center normal-text text-black">
              <object type="image/svg+xml" :data="qaIcon" class="pb-4"></object>
              <b
                >QA and <br />
                Testing</b
              >
            </div>
            <div class="text-center normal-text text-black">
              <object
                type="image/svg+xml"
                :data="seoIcon"
                class="pb-4"
              ></object>
              <b>
                SEO and <br />
                Marketing</b
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import doYouKnowImage from "@/assets/images/theme/question.svg";
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
      doYouKnowImage: doYouKnowImage,
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
};
</script>

<style scoped>
/* Do you know */
.que-section {
  margin: 15% 10%;
}

.flex-row {
  display: flex;
  align-items: center;
}

.flex-row {
  position: relative;
}

.pink-text {
  color: #f2709c;
}

.do-you-know-text {
  width: 100%;
  flex: none;
  margin: 0;
}

.do-you-know-text > .header-text {
  color: #3d3d3d;
}

/* How do we differ from others */

.light-orange-bg {
  background: rgba(255, 148, 114, 0.05);
  z-index: 100;
  margin-top: 0;
}

.light-orange-bg > .m-lr-15 > .description {
  width: 45% !important;
  padding-top: 173px;
  margin: auto;
}

.horizontal-gradient-line {
  margin-left: -0.5rem;
}

.steve-box {
  margin: 2rem 0.5rem;
  line-height: 2;
  padding: 3rem;
  border-radius: 0.9rem;
  text-align: left;
  border: 1px solid transparent;
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

.differ-description {
  margin: 3% 10%;
  color: rgba(0, 0, 0, 0.5);
}

.flex-grid {
  display: flex;
  flex-direction: column;
  padding-bottom: 7%;
}

.flex-grid-1,
.flex-grid-2 {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.flex-grid-1 > div,
.flex-grid-2 > div {
  padding: 1.25rem;
  width: 20%;
  display: flex;
  flex-direction: column;
}

.flex-grid-1 > div > object,
.flex-grid-2 > div > object {
  height: 7.5rem;
}

.flex-grid-1 > div > b,
.flex-grid-2 > div > b {
  font-size: 1.25rem;
}

@media (min-width: 768px) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }

  .m-lr-15 {
    margin: 0 15%;
  }

  .flex-row > .image {
    padding: 40% 0;
    background: url("./../../assets/images/theme/question.svg") no-repeat center;
    flex: 0 0 25%;
    background-size: cover;
  }

  .do-you-know-text {
    width: 49%;
    margin: -10rem auto 0;
  }

  .light-orange-bg {
    margin-top: -13rem;
  }
}

@media (min-width: 992px) {
  .header-text {
    font-size: 3rem;
    line-height: 3.6rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }
}
</style>
