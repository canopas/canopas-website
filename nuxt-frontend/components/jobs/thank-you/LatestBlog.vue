<template>
  <section class="container mt-16 lg:mb-60">
    <div class="flex flex-col items-center m-auto w-full">
      <h2 class="text-center mobile-header-2 lg:desk-header-2 text-black-87">
        Latest Blogs
      </h2>
    </div>

    <!-- MOBILE UI START  -->

    <div class="mt-6 block lg:hidden flex flex-col">
      <div
        class="relative flex flex-col items-center md:items-start md:justify-center"
      >
        <div
          v-for="(blog, index) in blogs"
          :key="blog"
          @mouseover="(activeIndex = index)"
          @mouseleave="(activeIndex = index)"
          @touchstart.passive="(activeIndex = index)"
          class="block mb-4 w-full border border-black-8 rounded-2xl box-border p-4"
        >
          <div
            class="flex flex-col w-full"
            @click="openBlog(blog.link, 'tap_blog_post')"
          >
            <div class="block w-full">
              <img
                :src="blog.thumbnail"
                class="object-cover rounded-xl cursor-pointer"
                loading="lazy"
                :alt="blog.title"
              />
            </div>
            <div class="flex justify-between mt-3 sub-h3-regular text-black-60">
              {{ blog.author }}
            </div>
            <div class="mt-2 sub-h1-semibold text-black-87 cursor-pointer">
              {{ blog.title }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START  -->
    <div
      class="hidden lg:block flex flex-col lg:min-h-[36.25rem] lg:py-0 lg:mt-16"
    >
      <div class="relative flex space-y-1 space-x-5">
        <div
          v-for="(blog, index) in blogs.slice(0, 1)"
          :key="blog"
          @mouseover="(activeIndex = index)"
          @mouseleave="(activeIndex = index)"
          @touchstart.passive="(activeIndex = index)"
          @click="openBlog(blog.link, 'tap_blog_post')"
          class="block grow-[2] basis-3/5 content-stretch mb-5 w-full border border-black-8 rounded-2xl box-border p-4 cursor-pointer"
        >
          <div class="flex flex-col w-full gap-2">
            <div class="block mt-4 w-full">
              <img
                :src="blog.thumbnail"
                class="object-cover rounded-xl cursor-pointer"
                loading="lazy"
                :alt="blog.title"
              />
            </div>
            <div class="flex justify-between sub-h1-regular text-black-4">
              {{ blog.author }}
            </div>
            <div class="desk-header-3 text-black-87 cursor-pointer">
              {{ blog.title }}
            </div>
            <div
              class="flex justify-between mobile-header-2-regular text-black-60 cursor-pointer"
            >
              <p class="line-clamp-3">
                {{ blog.description }}
              </p>
            </div>
          </div>
        </div>
        <div class="flex flex-col basis-2/5">
          <div
            v-for="(blog, index) in blogs.slice(1, 3)"
            :key="blog"
            @mouseover="(activeIndex = index)"
            @mouseleave="(activeIndex = index)"
            @touchstart.passive="(activeIndex = index)"
            @click="openBlog(blog.link, 'tap_blog_post')"
            class="block mb-5 w-full border border-black-8 rounded-xl box-border p-4 cursor-pointer"
          >
            <div class="flex flex-col w-full gap-2">
              <div class="block mt-4 w-full">
                <img
                  :src="blog.thumbnail"
                  class="object-cover rounded-lg cursor-pointer"
                  loading="lazy"
                  :alt="blog.title"
                />
              </div>
              <div class="flex justify-between sub-h4-regular text-black-4">
                {{ blog.author }}
              </div>
              <div class="sub-h1-semibold text-black-87 cursor-pointer">
                {{ blog.title }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- DESKTOP UI END  -->
    <nuxt-link
      class="white-gradient-border-btn primary-btn mt-2 lg:mt-8"
      :to="blogsURL"
      target="_blank"
      @click.native="$mixpanel.track('tap_flutter_blog_section')"
    >
      <span class="sub-h3-semibold lg:mobile-header-3-semibold"
        >Read More Blogs
      </span>
      <Icon name="fa6-solid:arrow-right" class="fa ml-2" />
    </nuxt-link>
  </section>
</template>

<script type="module">
import axios from "axios";
import Config from "@/config.js";
import { openBlog } from "@/utils.js";

export default {
  data() {
    return {
      openBlog,
      activeIndex: 0,
      blogsURL: Config.BLOG_URL,
      blogs: [],
    };
  },
  mounted() {
    this.getBlogs();
  },
  methods: {
    getBlogs() {
      axios
        .get(Config.API_BASE + "/api/blogs")
        .then((blogs) => {
          this.blogs = blogs.data;
        })
        .catch(() => {});
    },
  },
};
</script>
