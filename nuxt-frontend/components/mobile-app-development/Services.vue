<template>
  <section class="py-[50px]">
    <p
      class="container text-center font-inter-bold text-[1.875rem] leading-[2.813rem] text-white md:text-[4.125rem] md:leading-[4.95rem]"
    >
      Services
    </p>
    <!-- MOBILE UI START  -->
    <div
      class="container mt-12 flex flex-col md:hidden"
      v-for="(service, index) in services"
      :key="index"
    >
      <div class="flex flex-row items-start justify-center gap-4">
        <img
          :src="service.image"
          alt="mobile-service"
          class="h-[33px] w-[33px]"
        />
        <p
          class="v2-canopas-gradient-text font-inter-semibold text-[1.25rem] leading-[1.5rem]"
        >
          {{ service.title }}
        </p>
      </div>
      <div
        class="mt-4 flex h-full w-full flex-col rounded-[20px] bg-[#3F3F3F] px-4 py-[1.56rem]"
      >
        <div class="flex flex-row items-start justify-center gap-4">
          <div v-for="(technology, index) in service.technologies" :key="index">
            <div class="flex flex-col items-center gap-1">
              <img
                :src="technology.image"
                alt="technology"
                class="h-[60px] w-[60px] object-contain"
              />
              <p
                class="text-center font-inter-semibold text-[0.75rem] leading-[1.125rem] text-white/[0.76]"
                v-html="technology.name"
              ></p>
            </div>
          </div>
        </div>
        <p
          class="mt-6 text-center font-inter-regular text-base text-[#F5F5F5]/[0.85]"
        >
          {{ service.description }}
        </p>
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START  -->
    <div
      class="container mt-[6.25rem] hidden flex-col md:flex"
      v-for="(service, index) in services"
      :key="index"
    >
      <div class="flex flex-row">
        <div class="mb-auto mt-6 flex flex-row items-center gap-4">
          <img
            :src="service.image"
            alt="mobile-service"
            class="h-[33px] w-[33px]"
          />
          <p
            class="v2-canopas-gradient-text font-inter-semibold text-[1.5rem] leading-7 xl:text-[1.75rem] xl:leading-[2.1rem]"
          >
            {{ service.title }}
          </p>
        </div>
        <div class="gradient-border ml-16 h-auto border-l"></div>
        <div class="ml-4 flex flex-col gap-6 xl:ml-8">
          <div class="ml-4 flex flex-row justify-start gap-8">
            <div
              v-for="(technology, index) in service.technologies"
              :key="index"
            >
              <div class="flex flex-col items-center gap-1">
                <img
                  :src="technology.image"
                  alt="technology"
                  class="h-[70px] w-[70px] object-contain"
                />
                <p
                  class="text-center font-inter-semibold text-base text-white/[0.50]"
                  v-html="technology.name"
                ></p>
              </div>
            </div>
          </div>
          <div
            class="font-inter-regular text-base text-white/[0.85] xl:text-[1.25rem] xl:leading-[1.875rem]"
          >
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
