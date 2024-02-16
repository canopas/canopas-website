<template>
  <section class="container my-16 lg:my-60">
    <div class="text-center mobile-header-2 lg:desk-header-2 text-black-87">
      We are here to contribute
    </div>
    <div
      class="mt-4 mb-8 lg:mb-[4.5rem] text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
    >
      and that starts with the site itself,
      <span
        class="sub-h3-semibold lg:mobile-header-2-semibold v2-canopas-gradient-text"
        >Canopas Is Open Source!</span
      >
    </div>
    <div
      class="flex flex-col border mt-2 lg:mt-0"
      v-for="(contribution, index) in contributions"
      @click="
        openBlog(contribution.link, 'tap_home_github_contribution_section')
      "
      :key="index"
    >
      <div class="flex-col p-4 lg:p-6">
        <div class="mobile-header-3-semibold lg:desk-header-3 text-black-87">
          {{ contribution.title }}
        </div>
        <div
          class="mt-2 sub-h3-regular lg:mobile-header-2-regular text-black-60"
        >
          {{ contribution.description }}
        </div>
        <div class="mt-6" v-if="width > 991">
          <div
            class="mt-1 inline-flex pr-1"
            v-for="(tag, index) in contribution.tags"
            :key="index"
          >
            <div
              class="bg-white-smoke rounded-[1.875rem] py-1 px-4 lg:sub-h1-regular text-black-4"
              v-html="tag"
            ></div>
          </div>
        </div>
        <swiper
          v-else
          :slidesPerView="2"
          :grid="{
            rows: 2,
            fill: 'row',
          }"
          :modules="modules"
          :breakpoints="{
            '768': {
              slidesPerView: 4,
            },
          }"
          class="swiper-container mt-4 pl-0"
        >
          <swiper-slide
            v-for="(tag, index) in contribution.tags"
            :key="index"
            class="cursor-pointer !mr-1 !w-auto pb-1"
          >
            <div
              class="bg-white-smoke inline-flex rounded-[1.875rem] py-1 px-4 sub-h3-regular text-black-4 whitespace-nowrap"
              v-html="tag"
            ></div></swiper-slide
        ></swiper>
        <div class="flex gap-4 mt-4 lg:mt-6">
          <div class="flex items-center">
            <Icon
              class="fab footer-icon h-[1.625rem] w-[1.625rem] pr-[0.313rem]"
              name="fa6-solid:star"
            />
            <span
              class="v2-canopas-gradient-text mt-0.5 sub-h3-semibold lg:mobile-header-3-semibold"
              >{{ contribution.stars }}</span
            >
          </div>
          <div class="flex items-center gap-0.5 text-black-87">
            <span> <Icon class="fa h-5 w-5" name="fa6-solid:code-fork" /></span
            ><span class="sub-h3-semibold lg:mobile-header-3-semibold">{{
              contribution.forks
            }}</span>
          </div>
          <div class="flex items-center gap-2">
            <span
              class="h-4 w-4 rounded-full"
              :class="contribution.color"
            ></span
            ><span
              class="sub-h3-semibold lg:mobile-header-3-semibold text-black-87"
              >{{ contribution.language }}</span
            >
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script type="module">
import config from "@/config.js";
import { Swiper, SwiperSlide } from "swiper/vue";
import { Grid } from "swiper/modules";
import { openBlog } from "@/utils.js";

export default {
  data() {
    return {
      openBlog,
      width: 0,
      modules: [Grid],
      contributions: [
        {
          title: "compose-intro-showcase",
          description:
            "Highlight different features of the app using Jetpack Compose",
          tags: [
            `android`,
            `kotlin`,
            `material-design`,
            `kotlin-library`,
            `showcase`,
            `android-library`,
            `android-development`,
            `jetpack`,
            `intro`,
            `feature`,
            `jetpack-android`,
            `app-intro`,
            `jetpack-compose`,
            `jetpack-compose-tutorial`,
          ],
          link: "https://github.com/canopas/compose-intro-showcase",
          stars: config.INTRO_SHOWCASE_STARS,
          forks: "29",
          color: "bg-royal-purple",
          language: "Kotlin",
        },
        {
          title: "UIPilot",
          description: "The missing typesafe SwiftUI navigation library",
          tags: [
            `macos`,
            `swift`,
            `ios`,
            `router`,
            `xcode`,
            `navigation`,
            `routing`,
            `ios-ui`,
            `navigation-view`,
            `swiftui`,
            `swift5`,
          ],
          link: "https://github.com/canopas/UIPilot",
          stars: config.UIPILOT_STARS,
          forks: "22",
          color: "bg-flamingo",
          language: "Swift",
        },
        {
          title: "Jetpack Compose animations",
          description: "Cool animations implemented with Jetpack compose",
          tags: [],
          link: "https://github.com/canopas/compose-animations-examples",
          stars: config.JETPACK_COMPOSE_ANIMATION_STARS,
          forks: "29",
          color: "bg-royal-purple",
          language: "Kotlin",
        },
        {
          title: "Swiftui-animations-examples",
          description: "Cool animations implemented with SwiftUI.",
          tags: [],
          link: "https://github.com/canopas/swiftui-animations-examples",
          stars: config.SWIFT_UI_ANIMATION_STARS,
          forks: "2",
          color: "bg-flamingo",
          language: "Swift",
        },
        {
          title: "Tailwind Animations Examples",
          description: "Cool animations implemented with tailwindcss.",
          tags: [
            `css`,
            `vuejs`,
            `animations`,
            `animation-css`,
            `tailwind`,
            `tailwindcss`,
            `tailwind-elements`,
            `tailwind-animations`,
            `animation-example`,
            `tailwindcss-animate`,
            `tailwind-animations-examples`,
            `animation-examples`,
          ],
          link: "https://github.com/canopas/tailwind-animations-examples",
          stars: config.TAILWIND_ANIMATIONS_STARS,
          forks: "6",
          color: "bg-soft-yellow",
          language: "JavaScript",
        },
        {
          title: "ComposeRecyclerView",
          description:
            "Seamlessly integrate Jetpack Compose composables in RecyclerView with ComposeRecyclerView🔥. This library enhances performance⚡, tackles LazyList issues🔨, and offers built-in drag-and-drop👨🏽‍💻 support for dynamic UIs.",
          tags: [
            `kotlin`,
            `android`,
            `drag-and-drop`,
            `recyclerview`,
            `kotlin-library`,
            `android-library`,
            `recyclerview-adapter`,
            `draggable`,
            `composer-library`,
            `lazylist`,
            `jetpack-compose`,
            `jetpack-compose-tutorial`,
            `compose-ui`,
          ],
          link: "https://github.com/canopas/compose-recyclerview",
          stars: config.COMPOSE_RECYCLEVIEW_STARS,
          forks: "6",
          color: "bg-royal-purple",
          language: "Kotlin",
        },
        {
          title: "Canopas - Blog",
          description:
            "Feature-Rich blogs platform built with strapi and next.js",
          tags: [
            `website`,
            `development`,
            `web-development`,
            `web`,
            `responsive`,
            `reactjs`,
            `nextjs`,
          ],
          link: "https://github.com/canopas/canopas-blog",
          stars: config.CANOPAS_BLOG_WEBSITE_STARS,
          forks: "5",
          color: "bg-soft-yellow",
          language: "JavaScript",
        },
        {
          title: "Canopas - website",
          description:
            "Responsive website built with Vue.js and vite by following best practices.",
          tags: [
            `golang`,
            `aws`,
            `website`,
            `boilerplate`,
            `development`,
            `vuejs`,
            `web-development`,
          ],
          link: "https://github.com/canopas/canopas-website",
          stars: config.CANOPAS_WEBSITE_STARS,
          forks: "4",
          color: "bg-sea-green",
          language: "Vue",
        },
        {
          title: "RichEditorSwiftUI",
          description:
            "RichEditorSwiftUI is swift based library, it is made to make rich text editing easy with SwiftUI.",
          tags: [
            `editor`,
            `swift`,
            `string`,
            `uikit`,
            `rich-text-editor`,
            `richtext`,
            `swiftui`,
          ],
          link: "https://github.com/canopas/rich-editor-swiftui",
          stars: config.RICHEDITOR_SWIFTUI_STARS,
          forks: "1",
          color: "bg-flamingo",
          language: "Swift",
        },
        {
          title: "RichEditor",
          description: "Android WYSIWYG Rich editor for Jetpack compose.",
          tags: [],
          link: "https://github.com/canopas/rich-editor-compose",
          stars: config.RICHEDITOR_STARS,
          forks: "1",
          color: "bg-royal-purple",
          language: "Kotlin",
        },
      ],
    };
  },
  mounted() {
    this.width = window.innerWidth;
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
};
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/grid";
</style>
