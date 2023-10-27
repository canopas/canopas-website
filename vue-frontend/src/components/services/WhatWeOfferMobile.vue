<template>
  <div class="tw-text-black-core-900/[0.87] tw-text-center sm:tw-container">
    <div
      class="tw-py-12 tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-9 sm:tw-py-16"
    >
      <span>What We Offer</span>
    </div>

    <div class="tw-flex tw-flex-col">
      <div
        v-for="service in services"
        :key="service"
        class="tw-m-auto md:tw-w-[70%]"
      >
        <div>
          <img
            :src="service.images[0]"
            :srcset="`${service.images[0]} 400w, ${service.images[1]} 800w`"
            alt="Service-images"
            class="tw-h-full tw-w-full tw-object-cover"
            loading="lazy"
          />
        </div>

        <div :class="service.class" class="tw-px-2.5 sm:tw-px-20">
          <div
            class="tw-py-5 sm:tw-py-6"
            @click="showServices ? openUrl(service.url) : ''"
            :class="showServices ? 'tw-cursor-pointer' : ''"
          >
            <span
              class="tw-text-center tw-font-inter-semibold tw-text-[1.5rem] tw-leading-9"
              >{{ service.title }}</span
            >
          </div>
          <div
            class="tw-mb-10 tw-text-center tw-font-inter-regular tw-text-base sm:tw-mb-14"
          >
            <p>{{ service.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import androidApp400w from "@/assets/images/services/service/androidapp-400w.webp";
import androidApp800w from "@/assets/images/services/service/androidapp-800w.webp";

import flutterApp400w from "@/assets/images/services/service/flutterapp-400w.webp";
import flutterApp800w from "@/assets/images/services/service/flutterapp-800w.webp";

import iosApp400w from "@/assets/images/services/service/iosapp-400w.webp";
import iosApp800w from "@/assets/images/services/service/iosapp-800w.webp";

import backend400w from "@/assets/images/services/service/backend-400w.webp";
import backend800w from "@/assets/images/services/service/backend-800w.webp";

import uiuxdesign400w from "@/assets/images/services/service/uiuxdesign-400w.webp";
import uiuxdesign800w from "@/assets/images/services/service/uiuxdesign-800w.webp";

import frontend400w from "@/assets/images/services/service/frontend-400w.webp";
import frontend800w from "@/assets/images/services/service/frontend-800w.webp";

import config from "@/config.js";

export default {
  data() {
    return {
      showServices: !config.IS_PROD,

      services: [
        {
          images: [androidApp400w, androidApp800w],
          title: "Android App Development",
          description:
            "Take your business to new heights with an Android app! Tap into a global user base, boost customer engagement, communicate directly, and gather insightful data. Plus, enhance your visibility, add revenue streams, and provide top-notch customer service. With over a decade of expertise in Android app development, our team can help you craft high-quality, user-friendly apps that drive business growth and customer engagement. ",
          class:
            "tw-bg-gradient-to-b tw-from-[#EBF2FC]  via-[#EBF2FC]  tw-to-white tw-bg-gradient-[180deg] ",
          url: "/android-app-development",
        },
        {
          images: [iosApp400w, iosApp800w],
          title: "iPhone/iOS App Development",
          description:
            "Our seasoned iOS developers specialize in creating aesthetically pleasing, intuitive, and user-friendly apps that seamlessly function across all Apple devices. We leverage the latest iOS technologies to deliver solutions that give you an edge in the competitive Apple app market. Plus, our team follows Human Interface Guidelines defined by Apple to deliver exceptional user experience.",
          class:
            "tw-bg-gradient-to-b tw-from-[#FADFB7]  via-[#FADFB7]  tw-to-white tw-bg-gradient-[180deg] ",
          url: "/ios-app-development",
        },
        {
          images: [flutterApp400w, flutterApp800w],
          title: "Flutter App Development ",
          description:
            "Bring your multi-platform app vision to life with our Flutter app development services. We excel in crafting high-quality, stunning applications that work flawlessly on Android, iOS, and the web, all from a single codebase, facilitating faster time to market and cost savings.",
          class:
            "tw-bg-gradient-to-b tw-from-[#E0EDFF]  via-[#E0EDFF]  tw-to-white tw-bg-gradient-[180deg] ",
          url: "/flutter-app-development",
        },
        {
          images: [uiuxdesign400w, uiuxdesign800w],
          title: "UI/UX Design",
          description:
            "We help our clients see themselves through the eyes of the people whom they are trying to reach. Then we design something amazing.\
            Product design is about how it works, but we make sure it look nice too. Then, we'll design a mobile or web experience that feels great.",
          class:
            "tw-bg-gradient-to-b tw-from-[#E3D9EE]  via-[#E3D9EE]  tw-to-white tw-bg-gradient-[180deg] ",
        },
        {
          images: [backend400w, backend800w],
          title: "Backend Development",
          description:
            "The 'behind-the-scenes' backbone of any app, the backend, handles the critical tasks of data management, server configuration, and application logic. Our experts ensure your app is robust, secure, and efficient from its core, enabling it to handle high traffic loads, process data swiftly, and interact flawlessly with front-end requests.",
          class:
            "tw-bg-gradient-to-b tw-from-[#D2D8E3]/[0.95] via-[#D2D8E3]/[0.95] tw-to-white tw-bg-gradient-[180deg] ",
        },
        {
          images: [frontend400w, frontend800w],
          title: "Frontend Development",
          description:
            "The front end is the first point of interaction between your users and your app, making its design and functionality crucial. We focus on creating interactive, visually stunning, and easy-to-navigate interfaces that make your digital presence stand out, promoting enhanced user satisfaction and brand loyalty. ",
          class:
            "tw-bg-gradient-to-b tw-from-[#BBDEFF] via-[#BBDEFF] tw-to-white tw-bg-gradient-[180deg] ",
        },
      ],
    };
  },
  methods: {
    openUrl(url) {
      window.open(url, "_self");
    },
  },
};
</script>
