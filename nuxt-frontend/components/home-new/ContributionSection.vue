<template>
  <section class="container my-16 lg:my-60">
    <div class="text-center mobile-header-2 lg:desk-header-2 text-black-87">
      We are here to contribute
    </div>
    <div
      class="mt-4 mb-8 lg:mb-[4.5rem] text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
    >
      and that starts with this site itself,
      <nuxt-link
        class="sub-h3-semibold lg:mobile-header-2-semibold v2-canopas-gradient-text border-b border-[#ff835b]"
        :to="websiteOpenSourceUrl"
        target="_blank"
        @click.native="$mixpanel.track('tap_canopas_website_github')"
        >canopas is open source!</nuxt-link
      >
    </div>
    <div
      class="flex flex-col border mt-2 lg:mt-0"
      v-for="(contribution, index) in contributions"
      :key="index"
    >
      <nuxt-link
        :to="contribution.link"
        target="_blank"
        @click.native="$mixpanel.track('tap_home_github_contribution_section')"
        class="flex-col p-4 lg:p-6"
      >
        <div class="mobile-header-3-semibold lg:desk-header-3 text-black-87">
          {{ contribution.title }}
        </div>
        <div
          class="mt-2 sub-h3-regular lg:mobile-header-2-regular text-black-60"
        >
          {{ contribution.description }}
        </div>
        <div
          v-if="contribution.tags.length > 0"
          class="my-6 overflow-hidden overflow-x-scroll scrollbar-hidden"
        >
          <div class="w-full min-w-[1000px]">
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
        </div>
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
          <div
            v-if="contribution.forks > 0"
            class="flex items-center gap-0.5 text-black-87"
          >
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
      </nuxt-link>
    </div>
  </section>
</template>
<script type="module">
import config from "@/config.js";

export default {
  data() {
    return {
      width: 0,
      websiteOpenSourceUrl: "https://github.com/canopas/canopas-website",
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
          forks: "32",
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
          forks: "25",
          color: "bg-flamingo",
          language: "Swift",
        },
        {
          title: "Jetpack Compose animations",
          description: "Cool animations implemented with Jetpack compose",
          tags: [],
          link: "https://github.com/canopas/compose-animations-examples",
          stars: config.JETPACK_COMPOSE_ANIMATION_STARS,
          forks: "34",
          color: "bg-royal-purple",
          language: "Kotlin",
        },
        {
          title: "Swiftui-animations-examples",
          description: "Cool animations implemented with SwiftUI.",
          tags: [],
          link: "https://github.com/canopas/swiftui-animations-examples",
          stars: config.SWIFT_UI_ANIMATION_STARS,
          forks: "3",
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
          forks: "7",
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
          forks: "7",
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
          forks: "0",
          color: "bg-royal-purple",
          language: "Kotlin",
        },
      ],
    };
  },
  mounted() {
    this.width = window.innerWidth;
  },
  inject: ["mixpanel"],
};
</script>

<style scoped>
.scrollbar-hidden {
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}

/* Hide scrollbar for Chrome, Safari and Opera */
.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
</style>
