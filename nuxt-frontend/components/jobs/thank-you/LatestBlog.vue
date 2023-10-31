<template>
  <section class="container -mt-20 md:-mt-12 md:mt-0 md:mb-36">
    <div class="flex flex-col items-center mb-7 m-auto w-full">
      <h2
        class="text-[1.875rem] md:text-[3.4375rem] leading-[2.8125rem] md:leading-[5.15625rem] text-center text-black-core/[0.87] font-inter-bold"
      >
        Latest Blogs
      </h2>
    </div>

    <!-- MOBILE UI START  -->

    <div class="block lg:hidden flex flex-col md:flex-row md:mt-16">
      <div
        class="relative flex flex-col items-center md:items-start md:justify-center md:flex-[60%] md:ml-[5%]"
      >
        <div
          v-for="(blog, index) in blogs"
          :key="blog"
          @mouseover="activeIndex = index"
          @mouseleave="activeIndex = index"
          @touchstart.passive="activeIndex = index"
          class="block mb-5 w-full border border-black-core/[0.6] rounded-[20px] box-border p-4"
        >
          <div
            class="flex flex-col w-full"
            @click="openBlog(blog.link, 'tap_blog_post')"
          >
            <div class="block mt-4 w-full">
              <img
                :src="blog.thumbnail"
                class="object-cover rounded-lg cursor-pointer"
                loading="lazy"
                :alt="blog.title"
              />
            </div>
            <div
              class="flex flex-row justify-between my-2.5 font-inter-regular text-[0.875rem] leading-[1.3125rem] text-black-core/[0.6]"
            >
              <span class="w-auto">
                {{ blog.author }}
              </span>
            </div>
            <div class="my-2.5">
              <span
                class="font-inter-semibold text-[1.25rem] leading-[1.875rem] text-black-core[0.87] cursor-pointer"
              >
                {{ blog.title }}</span
              >
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START  -->
    <div class="hidden lg:block flex flex-col lg:min-h-[580px] lg:py-0">
      <div class="relative flex flex-row space-y-1 space-x-5">
        <div
          v-for="(blog, index) in blogs.slice(0, 1)"
          :key="blog"
          @mouseover="activeIndex = index"
          @mouseleave="activeIndex = index"
          @touchstart.passive="activeIndex = index"
          @click="openBlog(blog.link, 'tap_blog_post')"
          class="block grow-[2] basis-[70%] xl:basis-[70%] content-stretch mb-5 w-full border border-black-core/[0.6] rounded-[20px] box-border p-4 cursor-pointer"
        >
          <div class="flex flex-col w-full">
            <div class="block mt-4 w-full">
              <img
                :src="blog.thumbnail"
                class="object-cover rounded-lg cursor-pointer"
                loading="lazy"
                :alt="blog.title"
              />
            </div>
            <div
              class="flex flex-row justify-between my-2.5 font-inter-regular text-[0.875rem] leading-[1.3125rem] text-black-core/[0.6]"
            >
              <span class="w-auto">
                {{ blog.author }}
              </span>
            </div>
            <div class="my-2.5">
              <span
                class="font-inter-semibold text-[1.625rem] leading-[2.438rem] text-black-core[0.87] cursor-pointer"
              >
                {{ blog.title }}</span
              >
            </div>
            <div
              class="flex flex-row justify-between my-2.5 font-inter-regular text-[1.25rem] leading-[1.875rem] text-black-core/[0.6] cursor-pointer"
            >
              <p class="line-clamp-3">
                {{ blog.description }}
              </p>
            </div>
          </div>
        </div>
        <div class="flex flex-col basis-[30%] xl:basis-[30%]">
          <div
            v-for="(blog, index) in blogs.slice(1, 3)"
            :key="blog"
            @mouseover="activeIndex = index"
            @mouseleave="activeIndex = index"
            @touchstart.passive="activeIndex = index"
            @click="openBlog(blog.link, 'tap_blog_post')"
            class="block mb-5 w-full border border-black-core/[0.6] rounded-[20px] box-border p-4 cursor-pointer"
          >
            <div class="flex flex-col w-full">
              <div class="block mt-4 w-full">
                <img
                  :src="blog.thumbnail"
                  class="object-cover rounded-lg cursor-pointer"
                  loading="lazy"
                  :alt="blog.title"
                />
              </div>
              <div
                class="flex flex-row justify-between my-2 font-inter-regular text-[0.875rem] leading-[1.3125rem] text-black-core/[0.6]"
              >
                <span class="w-auto">
                  {{ blog.author }}
                </span>
              </div>
              <div class="my-2.5">
                <span
                  class="font-inter-semibold text-[1.25rem] leading-[1.875rem] text-black-core[0.87] cursor-pointer"
                >
                  {{ blog.title }}</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- DESKTOP UI END  -->

    <div class="flex justify-center mx-auto lg:mt-[30px]">
      <a
        target="_blank"
        :href="blogsURL"
        class="flex items-center m-0 mt-6 w-max rounded-full p-3 text-center gradient-btn consultation-btn"
      >
        <span
          class="mr-2.5 font-inter-semibold text-[1.1875rem] leading-[1.425rem]"
          >Read More Blogs</span
        >
      </a>
    </div>
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
