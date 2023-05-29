<template>
  <section class="tw-container tw--mt-12 md:tw-mt-0 md:tw-mb-36">
    <div
      class="tw-flex tw-flex-col tw-items-center tw-mb-7 tw-m-auto tw-w-full"
    >
      <h2
        class="tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.8125rem] md:tw-leading-[5.15625rem] tw-text-center tw-text-black-core/[0.87] tw-font-inter-bold"
      >
        Latest Blogs
      </h2>
    </div>
    <div
      class="tw-block lg:tw-hidden tw-flex tw-flex-col md:tw-flex-row md:tw-mt-[4rem]"
    >
      <div
        class="tw-relative tw-flex tw-flex-col tw-items-center md:tw-items-start md:tw-justify-center md:tw-flex-[60%] md:tw-ml-[5%]"
      >
        <div
          v-for="(blog, index) in blogs"
          :key="blog"
          @mouseover="activeIndex = index"
          @mouseleave="activeIndex = index"
          @touchstart.passive="activeIndex = index"
          @click.native="mixpanel.track('tap_blog_post')"
          class="tw-block tw-mb-[20px] tw-w-full tw-border tw-border-black-core/[0.6] tw-rounded-[20px] tw-box-border tw-p-4"
        >
          <div class="tw-flex tw-flex-col tw-w-full">
            <div
              @click="openBlog(blog)"
              class="tw-block tw-mt-[1rem] tw-w-full"
            >
              <img
                :src="blog.thumbnail"
                class="tw-object-cover tw-rounded-lg tw-cursor-pointer"
                loading="lazy"
                :alt="blog.title"
              />
            </div>
            <div
              class="tw-flex tw-flex-row tw-justify-between tw-my-[10px] tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.3125rem] tw-text-black-core/[0.6]"
            >
              <span class="tw-w-auto">
                {{ blog.author }}
              </span>
            </div>
            <div class="tw-my-[10px]">
              <span
                @click="openBlog(blog)"
                class="tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem] tw-text-black-core[0.87] tw-cursor-pointer"
              >
                {{ blog.title }}</span
              >
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="tw-hidden lg:tw-block tw-flex tw-flex-col lg:tw-min-h-[580px] lg:tw-py-0"
    >
      <div class="tw-relative tw-flex tw-flex-row tw-space-y-1 tw-space-x-5">
        <div
          v-for="(blog, index) in blogs.slice(0, 1)"
          :key="blog"
          @mouseover="activeIndex = index"
          @mouseleave="activeIndex = index"
          @touchstart.passive="activeIndex = index"
          @click.native="mixpanel.track('tap_blog_post')"
          class="tw-block tw-grow-[2] tw-basis-[70%] xl:tw-basis-[70%] tw-content-stretch tw-mb-[20px] tw-w-full tw-border tw-border-black-core/[0.6] tw-rounded-[20px] tw-box-border tw-p-4"
        >
          <div class="tw-flex tw-flex-col tw-w-full">
            <div
              @click="openBlog(blog)"
              class="tw-block tw-mt-[1rem] tw-w-full"
            >
              <img
                :src="blog.thumbnail"
                class="tw-object-cover tw-rounded-lg tw-cursor-pointer"
                loading="lazy"
                :alt="blog.title"
              />
            </div>
            <div
              class="tw-flex tw-flex-row tw-justify-between tw-my-[10px] tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.3125rem] tw-text-black-core/[0.6]"
            >
              <span class="tw-w-auto">
                {{ blog.author }}
              </span>
            </div>
            <div class="tw-my-[10px]">
              <span
                @click="openBlog(blog)"
                class="tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem] tw-text-black-core[0.87] tw-cursor-pointer"
              >
                {{ blog.title }}</span
              >
            </div>
            <div
              @click="openBlog(blog)"
              class="tw-flex tw-flex-row tw-justify-between tw-my-[10px] tw-font-inter-regular tw-text-[1.625rem] tw-leading-[2.4375rem] tw-text-black-core/[0.6] tw-cursor-pointer"
            >
              <p class="tw-line-clamp-3">
                {{ blog.description }}
              </p>
            </div>
          </div>
        </div>
        <div class="tw-flex tw-flex-col tw-basis-[30%] xl:tw-basis-[30%]">
          <div
            v-for="(blog, index) in blogs.slice(1, 3)"
            :key="blog"
            @mouseover="activeIndex = index"
            @mouseleave="activeIndex = index"
            @touchstart.passive="activeIndex = index"
            @click.native="mixpanel.track('tap_blog_post')"
            class="tw-block tw-mb-[20px] tw-w-full tw-border tw-border-black-core/[0.6] tw-rounded-[20px] tw-box-border tw-p-4"
          >
            <div class="tw-flex tw-flex-col tw-w-full">
              <div
                @click="openBlog(blog)"
                class="tw-block tw-mt-[1rem] tw-w-full"
              >
                <img
                  :src="blog.thumbnail"
                  class="tw-object-cover tw-rounded-lg tw-cursor-pointer"
                  loading="lazy"
                  :alt="blog.title"
                />
              </div>
              <div
                class="tw-flex tw-flex-row tw-justify-between tw-my-[8px] tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.3125rem] tw-text-black-core/[0.6]"
              >
                <span class="tw-w-auto">
                  {{ blog.author }}
                </span>
              </div>
              <div class="tw-my-[10px]">
                <span
                  @click="openBlog(blog)"
                  class="tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem] tw-text-black-core[0.87] tw-cursor-pointer"
                >
                  {{ blog.title }}</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="tw-flex tw-justify-center tw-mx-auto lg:tw-mt-[30px]">
      <a
        target="_blank"
        :href="blogsURL"
        class="tw-flex tw-w-max tw-items-center tw-rounded-full tw-py-2 sm:tw-py-3 tw-px-12 sm:tw-px-3 gradient-btn consultation-btn"
      >
        <span
          class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.425rem]"
          >Read More Blogs</span
        >
      </a>
    </div>
  </section>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import axios from "axios";
import Config from "@/config.js";

export default {
  data() {
    return {
      activeIndex: 0,
      blogsURL: Config.BLOG_URL,
      blogs: [],
    };
  },
  mounted() {
    this.getBlogs();
  },
  inject: ["mixpanel"],
  methods: {
    openBlog(blog) {
      window.open(blog.link, "_blank");
      this.mixpanel.track("tap_blog_post");
    },
    getBlogs() {
      axios
        .get(Config.API_BASE + "/api/blogs")
        .then((blogs) => {
          this.blogs = blogs.data;
        })
        .catch(() => {});
    },
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>
