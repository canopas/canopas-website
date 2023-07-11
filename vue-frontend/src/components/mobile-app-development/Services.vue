<template>
  <section class="tw-py-[100px]">
    <p
      class="tw-container tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.813rem] tw-text-white md:tw-text-[4.125rem] md:tw-leading-[4.95rem]"
    >
      Services
    </p>
    <!-- MOBILE UI START  -->
    <div
      class="tw-container tw-mt-12 tw-flex tw-flex-col md:tw-hidden"
      v-for="(service, index) in services"
      :key="index"
    >
      <div
        class="tw-flex tw-flex-row tw-items-center tw-justify-center tw-gap-[1rem]"
      >
        <img
          :src="service.image"
          alt="mobile-service"
          class="tw-h-[33px] tw-w-[33px]"
        />
        <p
          class="v2-canopas-gradient-text tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.5rem]"
        >
          {{ service.title }}
        </p>
      </div>
      <div
        class="tw-mt-4 tw-flex tw-h-full tw-w-full tw-flex-col tw-rounded-[20px] tw-bg-[#3F3F3F] tw-px-[1rem] tw-py-[1.56rem]"
      >
        <div
          class="tw-flex tw-flex-row tw-items-center tw-justify-center tw-gap-4"
        >
          <div v-for="(technology, index) in service.technologies" :key="index">
            <div class="tw-flex tw-flex-col tw-items-center tw-gap-1">
              <img
                :src="technology.image"
                alt="technology"
                class="tw-h-[50px] tw-w-[50px] tw-object-contain"
              />
              <p
                class="tw-font-inter-semibold tw-text-[0.75rem] tw-leading-[1.125rem] tw-text-white/[0.76]"
              >
                {{ technology.name }}
              </p>
            </div>
          </div>
        </div>
        <p
          class="tw-mt-6 tw-text-center tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-text-[#F5F5F5]/[0.85]"
        >
          {{ service.description }}
        </p>
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START  -->
    <div
      class="tw-container tw-mt-16 tw-hidden tw-flex-col md:tw-flex"
      v-for="(service, index) in services"
      :key="index"
    >
      <div class="tw-flex tw-flex-row">
        <div
          class="tw-mb-auto tw-mt-[1.5rem] tw-flex tw-flex-row tw-items-center tw-gap-4"
        >
          <img
            :src="service.image"
            alt="mobile-service"
            class="tw-h-[33px] tw-w-[33px]"
          />
          <p
            class="v2-canopas-gradient-text tw-font-inter-semibold tw-text-[1.5rem] tw-leading-[1.75rem] xl:tw-text-[1.75rem] xl:tw-leading-[2.1rem]"
          >
            {{ service.title }}
          </p>
        </div>
        <div class="gradient-border tw-ml-16 tw-h-auto tw-border-l-[1px]"></div>
        <div class="tw-ml-4 tw-flex tw-flex-col tw-gap-6 xl:tw-ml-8">
          <div class="tw-ml-4 tw-flex tw-flex-row tw-justify-start tw-gap-8">
            <div
              v-for="(technology, index) in service.technologies"
              :key="index"
            >
              <div class="tw-flex tw-flex-col tw-items-center tw-gap-1">
                <img
                  :src="technology.image"
                  alt="technology"
                  class="tw-h-[60px] tw-w-[60px] tw-object-contain"
                />
                <p
                  class="tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.5rem] tw-text-white/[0.50]"
                >
                  {{ technology.name }}
                </p>
              </div>
            </div>
          </div>
          <div
            class="tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-text-white/[0.85] xl:tw-text-[1.25rem] xl:tw-leading-[1.875rem]"
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
import java from "@/assets/images/mobile-app-development/services/android/java.webp";
import clang from "@/assets/images/mobile-app-development/services/android/C++.webp";
import ios from "@/assets/images/mobile-app-development/services/ios/iOS.webp";
import swift from "@/assets/images/mobile-app-development/services/ios/swift.webp";
import javaScript from "@/assets/images/mobile-app-development/services/ios/javaScript.webp";
import objectiveC from "@/assets/images/mobile-app-development/services/ios/objectiveC.webp";
import cSharp from "@/assets/images/mobile-app-development/services/ios/cSharp.webp";
import flutter from "@/assets/images/mobile-app-development/services/flutter/flutter.webp";
import dart from "@/assets/images/mobile-app-development/services/flutter/dart.webp";

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
              image: gradle,
              name: "GRADLE",
            },
            {
              image: kotlin,
              name: "KOTLIN",
            },
            {
              image: java,
              name: "JAVA",
            },
            {
              image: clang,
              name: "C++",
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
              image: objectiveC,
              name: "OBJECTIVE C",
            },
            {
              image: javaScript,
              name: "JAVASCRIPT",
            },
            {
              image: cSharp,
              name: "C#",
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
