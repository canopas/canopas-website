<template>
  <div class="text-black-core-900/[0.87] text-center sm:container">
    <div
      class="py-12 text-center font-inter-bold text-[1.875rem] leading-9 sm:py-16"
    >
      <span>What We Offer</span>
    </div>

    <div class="flex flex-col">
      <div v-for="service in services" :key="service" class="m-auto md:w-[70%]">
        <div>
          <img
            :src="service.images[0]"
            :srcset="`${service.images[0]} 400w, ${service.images[1]} 800w`"
            alt="Service-images"
            class="h-full w-full object-cover"
            loading="lazy"
          />
        </div>

        <div :class="service.class" class="px-2.5 sm:px-20">
          <div
            class="py-5 sm:py-6"
            @click="service.showServices ? openUrl(service.url) : ''"
            :class="service.showServices ? 'cursor-pointer' : ''"
          >
            <span
              class="text-center font-inter-semibold text-[1.5rem] leading-9"
              >{{ service.title }}</span
            >
          </div>
          <div class="mb-10 text-center font-inter-regular text-base sm:mb-14">
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
      services: [
        {
          images: [androidApp400w, androidApp800w],
          title: "Android App Development",
          description:
            "Take your business to new heights with an Android app! Tap into a global user base, boost customer engagement, communicate directly, and gather insightful data. Plus, enhance your visibility, add revenue streams, and provide top-notch customer service. With over a decade of expertise in Android app development, our team can help you craft high-quality, user-friendly apps that drive business growth and customer engagement. ",
          class:
            "bg-gradient-to-b from-[#EBF2FC]  via-[#EBF2FC]  to-white bg-gradient-[180deg] ",
          url: "/android-app-development",
          showServices: config.SHOW_ANDROID_APP_DEVELOPMENT_PAGE,
        },
        {
          images: [iosApp400w, iosApp800w],
          title: "iPhone/iOS App Development",
          description:
            "Our seasoned iOS developers specialize in creating aesthetically pleasing, intuitive, and user-friendly apps that seamlessly function across all Apple devices. We leverage the latest iOS technologies to deliver solutions that give you an edge in the competitive Apple app market. Plus, our team follows Human Interface Guidelines defined by Apple to deliver exceptional user experience.",
          class:
            "bg-gradient-to-b from-[#FADFB7]  via-[#FADFB7]  to-white bg-gradient-[180deg] ",
          url: "/ios-app-development",
          showServices: config.SHOW_IOS_APP_DEVELOPMENT_PAGE,
        },
        {
          images: [flutterApp400w, flutterApp800w],
          title: "Flutter App Development ",
          description:
            "Bring your multi-platform app vision to life with our Flutter app development services. We excel in crafting high-quality, stunning applications that work flawlessly on Android, iOS, and the web, all from a single codebase, facilitating faster time to market and cost savings.",
          class:
            "bg-gradient-to-b from-[#E0EDFF]  via-[#E0EDFF]  to-white bg-gradient-[180deg] ",
          url: "/flutter-app-development",
          showServices: config.SHOW_FLUTTER_APP_DEVELOPMENT_PAGE,
        },
        {
          images: [uiuxdesign400w, uiuxdesign800w],
          title: "UI/UX Design",
          description:
            "We help our clients see themselves through the eyes of the people whom they are trying to reach. Then we design something amazing.\
            Product design is about how it works, but we make sure it look nice too. Then, we'll design a mobile or web experience that feels great.",
          class:
            "bg-gradient-to-b from-[#E3D9EE]  via-[#E3D9EE]  to-white bg-gradient-[180deg] ",
        },
        {
          images: [backend400w, backend800w],
          title: "Backend Development",
          description:
            "The 'behind-the-scenes' backbone of any app, the backend, handles the critical tasks of data management, server configuration, and application logic. Our experts ensure your app is robust, secure, and efficient from its core, enabling it to handle high traffic loads, process data swiftly, and interact flawlessly with front-end requests.",
          class:
            "bg-gradient-to-b from-[#D2D8E3]/[0.95] via-[#D2D8E3]/[0.95] to-white bg-gradient-[180deg] ",
        },
        {
          images: [frontend400w, frontend800w],
          title: "Frontend Development",
          description:
            "The front end is the first point of interaction between your users and your app, making its design and functionality crucial. We focus on creating interactive, visually stunning, and easy-to-navigate interfaces that make your digital presence stand out, promoting enhanced user satisfaction and brand loyalty. ",
          class:
            "bg-gradient-to-b from-[#BBDEFF] via-[#BBDEFF] to-white bg-gradient-[180deg] ",
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
