<template>
  <section class="mt-16 lg:mt-60 pb-14 lg:pb-60">
    <p
      class="container text-center text-white mobile-header-2 lg:desk-header-2"
    >
      Services
    </p>
    <!-- MOBILE UI START  -->
    <div
      class="container mt-8 flex flex-col lg:hidden"
      v-for="(service, index) in services"
      :key="index"
    >
      <div class="flex items-start justify-center gap-4">
        <img :src="service.image" alt="mobile-service" class="h-8 w-8" />
        <p class="v2-canopas-gradient-text mobile-header-3-semibold">
          {{ service.title }}
        </p>
      </div>
      <div
        class="mt-4 flex h-full w-full flex-col rounded-xl bg-white-core-8 px-4 py-[1.56rem]"
      >
        <div class="flex items-start justify-center gap-4">
          <div v-for="(technology, index) in service.technologies" :key="index">
            <div class="flex flex-col items-center gap-1">
              <img
                :src="technology.image"
                alt="technology"
                class="h-[3.75rem] w-[3.75rem] object-contain"
                loading="lazy"
              />
              <p
                class="text-center sub-h4-semibold text-white-core-60"
                v-html="technology.name"
              ></p>
            </div>
          </div>
        </div>
        <p class="mt-6 text-center sub-h3-regular text-white-core-80">
          {{ service.description }}
        </p>
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START  -->
    <div
      class="container mt-16 hidden flex-col lg:flex"
      v-for="(service, index) in services"
      :key="index"
    >
      <div class="flex">
        <div class="mb-auto mt-6 flex items-center gap-4">
          <img
            :src="service.image"
            alt="mobile-service"
            class="h-[33px] w-[33px]"
          />
          <p class="v2-canopas-gradient-text mobile-header-2-semibold">
            {{ service.title }}
          </p>
        </div>
        <div class="gradient-border ml-16 h-auto border-l"></div>
        <div class="ml-4 flex flex-col gap-6 xl:ml-8">
          <div class="ml-4 flex justify-start gap-8">
            <div
              v-for="(technology, index) in service.technologies"
              :key="index"
            >
              <div class="flex flex-col items-center gap-1">
                <img
                  :src="technology.image"
                  alt="technology"
                  class="h-[4.375rem] w-[4.375rem] object-contain"
                />
                <p
                  class="text-center sub-h1-regular text-white-core-60"
                  v-html="technology.name"
                ></p>
              </div>
            </div>
          </div>
          <div class="mobile-header-3-regular text-white-core-80">
            {{ service.description }}
          </div>
        </div>
      </div>
    </div>
    <!-- DESKTOP UI END  -->
  </section>
</template>
<script>
import android from "@/assets/images/mobile-app-development/services/android/android.webp";
import gradle from "@/assets/images/mobile-app-development/services/android/gradle.webp";
import kotlin from "@/assets/images/mobile-app-development/services/android/kotlin.webp";
import androidstudio from "@/assets/images/mobile-app-development/services/android/androidstudio.webp";
import jetpackcompose from "@/assets/images/mobile-app-development/services/android/jetpackcompose.webp";
import ios from "@/assets/images/mobile-app-development/services/ios/iOS.webp";
import swift from "@/assets/images/mobile-app-development/services/ios/swift.webp";
import xcode from "@/assets/images/mobile-app-development/services/ios/xcode.webp";
import cocoapods from "@/assets/images/mobile-app-development/services/ios/cocoapods.webp";
import flutter from "@/assets/images/mobile-app-development/services/flutter/flutter.webp";
import dart from "@/assets/images/mobile-app-development/services/flutter/dart.webp";
import firebase from "@/assets/images/mobile-app-development/services/flutter/firebase.webp";

export default {
  data() {
    return {
      services: [
        {
          id: 1,
          title: "Android App Development",
          image: android,
          technologies: [
            {
              image: kotlin,
              name: "KOTLIN",
            },
            {
              image: gradle,
              name: "GRADLE",
            },
            {
              image: androidstudio,
              name: `ANDROID <br>
                     STUDIO`,
            },
            {
              image: jetpackcompose,
              name: `JETPACK <br>
               COMPOSE`,
            },
          ],
          description:
            "Take your business to new heights with an Android app! Tap into a global user base, boost customer engagement, communicate directly, and gather insightful data. Plus, enhance your visibility, add revenue streams, and provide top-notch customer service. With over a decade of expertise in Android app development, our team can help you craft high-quality, user-friendly apps that drive business growth and customer engagement.",
        },
        {
          id: 2,
          title: "iOS App Development",
          image: ios,
          technologies: [
            {
              image: swift,
              name: "SWIFT",
            },
            {
              image: xcode,
              name: "XCODE",
            },
            {
              image: cocoapods,
              name: "COCOAPODS",
            },
          ],
          description:
            "Our seasoned iOS developers specialize in creating aesthetically pleasing, intuitive, and user-friendly apps that seamlessly function across all Apple devices. We leverage the latest iOS technologies to deliver solutions that give you an edge in the competitive Apple app market. Plus, our team follows Human Interface Guidelines defined by Apple to deliver exceptional user experience.",
        },
        {
          id: 3,
          title: "Flutter App Development",
          image: flutter,
          technologies: [
            {
              image: dart,
              name: "DART",
            },
            {
              image: firebase,
              name: "FIREBASE",
            },
          ],
          description:
            "Bring your multi-platform app vision to life with our Flutter app development services. We excel in crafting high-quality, stunning applications that work flawlessly on Android, iOS, and the web, all from a single codebase, facilitating faster time to market and cost savings.",
        },
      ],
    };
  },
};
</script>
<style scoped>
.gradient-border {
  border-image: linear-gradient(to bottom, #f2709c, #ff835b) 1;
}
</style>
