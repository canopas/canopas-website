<template>
  <section class="tw-my-[7.813rem]">
    <div
      class="tw-py-12 tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.25rem] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem] xl:tw-py-24"
    >
      <span>What We Offer</span>
    </div>
    <div class="tw-container"></div>
    <div class="tw-text-black-core-900/[0.87] tw-container tw-text-center">
      <div class="tw-mb-16 tw-flex tw-justify-center">
        <div class="tw-flex tw-w-[50%] tw-flex-col tw-justify-end tw-text-left">
          <div
            v-for="(service, index) in services"
            :key="service"
            class="tw-relative tw-h-[100vh]"
            :ref="'image-' + index"
          >
            <div class="tw-absolute tw-top-[50%] tw-translate-y-[-50%]">
              <div
                class="tw-py-5"
                @click="service.showServices ? openUrl(service.url) : ''"
                :class="service.showServices ? 'tw-cursor-pointer' : ''"
              >
                <span
                  class="tw-font-inter-semibold tw-text-[1.5rem] tw-leading-9 md:tw-text-[2.25rem] md:tw-leading-[3.43rem] xl:tw-text-[3.125rem] xl:tw-leading-[4.6875rem]"
                  >{{ service.title }}</span
                >
              </div>
              <div
                class="tw-font-inter-regular tw-text-base md:tw-text-[1.094rem] md:tw-leading-[1.625rem] xl:tw-text-[1.188rem] xl:tw-leading-[1.781rem]"
              >
                <p>{{ service.description }}</p>
              </div>
            </div>
          </div>
        </div>
        <div
          class="-tw-translate-[50%] tw-sticky tw-top-0 tw-flex tw-w-[50%] tw-h-[100vh] tw-items-center"
        >
          <div
            v-for="service in services"
            :key="service"
            class="tw-absolute tw-w-[100%] tw-overflow-hidden tw-pb-[100%]"
            :style="{
              clipPath: `polygon(0% ${service.path[0]}%, 100% ${service.path[1]}%, 100% 100%, 0% 100%)`,
            }"
          >
            <img
              :src="service.image"
              alt="Service-images"
              class="tw-absolute tw-right-0 tw-top-0 tw-z-[1] tw-ml-auto tw-h-full tw-w-[80%] tw-object-cover"
              loading="lazy"
            />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script type="module">
import androidApp800w from "@/assets/images/services/service/androidapp-800w.webp";
import flutterApp800w from "@/assets/images/services/service/flutterapp-800w.webp";
import frontend800w from "@/assets/images/services/service/frontend-800w.webp";
import backend800w from "@/assets/images/services/service/backend-800w.webp";
import uiuxdesign800w from "@/assets/images/services/service/uiuxdesign-800w.webp";
import iosApp800w from "@/assets/images/services/service/iosapp-800w.webp";
import { elementInViewPort } from "@/utils";
import config from "@/config.js";
export default {
  data() {
    return {
      services: [
        {
          path: [0, 0],
          image: androidApp800w,
          title: "Android App Development",
          description:
            "Take your business to new heights with an Android app! Tap into a global user base, boost customer engagement, communicate directly, and gather insightful data. Plus, enhance your visibility, add revenue streams, and provide top-notch customer service. With over a decade of expertise in Android app development, our team can help you craft high-quality, user-friendly apps that drive business growth and customer engagement. ",
          url: "/android-app-development",
          showServices: config.SHOW_ANDROID_APP_DEVELOPMENT_PAGE,
        },
        {
          path: [100, 100],
          image: iosApp800w,
          title: "iPhone/iOS App Development",
          description:
            "Our seasoned iOS developers specialize in creating aesthetically pleasing, intuitive, and user-friendly apps that seamlessly function across all Apple devices. We leverage the latest iOS technologies to deliver solutions that give you an edge in the competitive Apple app market. Plus, our team follows Human Interface Guidelines defined by Apple to deliver exceptional user experience.",
          url: "/ios-app-development",
          showServices: config.SHOW_IOS_APP_DEVELOPMENT_PAGE,
        },
        {
          path: [100, 100],
          image: flutterApp800w,
          title: "Flutter App Development ",
          description:
            "Bring your multi-platform app vision to life with our Flutter app development services. We excel in crafting high-quality, stunning applications that work flawlessly on Android, iOS, and the web, all from a single codebase, facilitating faster time to market and cost savings.",
          url: "/flutter-app-development",
          showServices: config.SHOW_FLUTTER_APP_DEVELOPMENT_PAGE,
        },
        {
          path: [100, 100],
          image: uiuxdesign800w,
          title: "UI/UX Design",
          description:
            "We help our clients see themselves through the eyes of the people whom they are trying to reach. Then we design something amazing.\
            Product design is about how it works, but we make sure it look nice too. Then, we'll design a mobile or web experience that feels great.",
        },
        {
          path: [100, 100],
          image: backend800w,
          title: "Backend Development",
          description:
            "The 'behind-the-scenes' backbone of any app, the backend, handles the critical tasks of data management, server configuration, and application logic. Our experts ensure your app is robust, secure, and efficient from its core, enabling it to handle high traffic loads, process data swiftly, and interact flawlessly with front-end requests.",
        },
        {
          path: [100, 100],
          image: frontend800w,
          title: "Frontend Development",
          description:
            "The front end is the first point of interaction between your users and your app, making its design and functionality crucial. We focus on creating interactive, visually stunning, and easy-to-navigate interfaces that make your digital presence stand out, promoting enhanced user satisfaction and brand loyalty. ",
        },
      ],
      lastScrollY: 0,
      prevIndex: 0,
      height: 0,
    };
  },
  methods: {
    handleScroll() {
      // get scroll direction
      const scrollY = window.scrollY;
      const scrollUp = this.lastScrollY > scrollY && window.pageYOffset > 100;

      // check which element is in viewport
      const elementIdx = elementInViewPort(this.$refs);
      const index = parseInt(
        elementIdx ? elementIdx.charAt(elementIdx.length - 1) : 0,
      );

      if (this.prevIndex != index) {
        // when image will change, prevImage should be at path 0 or 100
        if (this.prevIndex != 0) {
          this.services[this.prevIndex].path = [
            scrollUp ? 100 : 0,
            scrollUp ? 100 : 0,
          ];
        }

        this.prevIndex = index;
      }

      // make previous images clip path 0
      this.services = this.services.map((val, i) => {
        val.path = i < index ? [0, 0] : i > index ? [100, 100] : val.path;
        return val;
      });

      // random offset used when scrolling
      const rect = this.$refs[elementIdx][0].getBoundingClientRect();
      const scrollOffset =
        ((window.innerHeight - rect.top) * 100) / window.innerHeight;

      if (scrollUp) {
        this.services[0].path = [0, 0];
      }

      // if firstImage then path should be 0
      // else set respective path
      if (index == 0) {
        this.services[index].path = [0, 0];
      } else if (
        !scrollUp &&
        index == this.services.length - 1 &&
        rect.top < window.innerHeight / 2
      ) {
        this.services[index].path = [
          100 - scrollOffset * 1.2,
          100 - scrollOffset * 1.2,
        ];
      } else {
        this.services[index].path = [100 - scrollOffset, 100 - scrollOffset];
      }

      this.lastScrollY = scrollY;
    },
    openUrl(url) {
      window.open(url, "_self");
    },
  },
  mounted() {
    this.height = window.innerHeight;
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
};
</script>
