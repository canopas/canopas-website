<template>
  <div class="my-8 lg:my-[7.813rem] text-black-core/[.87] overflow-hidden">
    <div class="container flex flex-col items-center m-auto w-full">
      <div class="container">
        <p
          class="text-black-core/[.03] text-[2.5rem] md:text-[3.75rem] xl:text-[5rem] leading-[1.875rem] md:leading-[4.0625rem] xl:leading-[5rem] text-center"
        >
          Blogs
        </p>
        <h2
          class="-mt-5 md:mt-[-50px] mx-auto xl:w-[65%] pb-9 md:pb-12 text-[1.875rem] leading-[2.438rem] md:text-[2.65rem] md:leading-[3.218rem] lg:text-[3.438rem] lg:leading-[4rem] text-center text-black-core/[0.87] font-roboto-bold"
        >
          Do we have a team with the right skills?
        </h2>
      </div>
      <div class="mt-[1.125rem] lg:w-[65%]">
        <div
          class="w-full text-center font-inter-medium text-black-core/[0.6] text-[1rem] md:text-[1.25rem] xl:text-[1.5rem] leading-[1.375rem] md:leading-8 xl:leading-[2.625rem]"
        >
          Well, at least the community says Hell Yeah. Our blogs hosted on
          medium have
          <span class="v2-canopas-gradient-text font-inter-bold">100k+</span>
          minutes monthly reading time and it’s only rising.
        </div>
      </div>
    </div>
    <div class="relative mt-12">
      <img
        v-if="width < 680"
        :src="bgMobile"
        loading="lazy"
        class="md:hidden absolute top-0 -right-10 sm:-right-20 h-full"
        alt="blog-background-image"
      />
    </div>
    <div
      class="flex flex-col md:flex-row md:mt-16 z-[1] lg:min-h-[580px] md:pb-[15%] lg:py-0"
      :class="width < 768 || width > 2000 ? 'container' : ''"
    >
      <div
        class="relative flex flex-col items-center md:items-start md:justify-center md:flex-[60%] md:ml-[5%] lg:ml-[10%] xl:ml-[15%]"
        :class="width > 2000 ? 'ml-0' : ''"
      >
        <div
          v-for="(blog, index) in blogs"
          :key="blog"
          @mouseover="activeIndex = index"
          @mouseleave="activeIndex = index"
          @touchstart.passive="activeIndex = index"
          @click.native="$mixpanel.track('tap_blog_post')"
          class="w-full"
        >
          <div class="flex flex-col 2xl:pr-16">
            <hr
              v-if="index == 0"
              class="hidden md:block md:ml-[82px] 2xl:ml-[95px] h-[0.063rem]"
            />
            <div class="flex gap-8 mt-[18px] md:mt-[45px]">
              <div class="flex flex-col items-center">
                <span
                  class="font-inter-semibold md:font-inter-regular text-[1.5rem] md:text-[2.188rem] 2xl:text-[3.125rem] leading-[1.95rem] md:leading-[2.183rem] 2xl:leading-[4.063rem]"
                >
                  {{ blog.pubDate[1] }}
                </span>
                <span
                  class="text-center font-inter-medium md:font-inter-regular text-[0.75rem] md:text-[1rem] leading-4 2xl:leading-[1.125rem]"
                >
                  {{ blog.pubDate[0] }}
                </span>
              </div>

              <div
                class="flex flex-col w-full"
                @mouseover="
                  activeBlog = blog;
                  animate = true;
                "
                @mouseleave="animate = false"
                @touchstart.passive="
                  activeBlog = blog;
                  animate = true;
                "
                @touchend="animate = false"
              >
                <div>
                  <span
                    @click="
                      width > 768 ? openBlog(blog.link, 'tap_blog_post') : ''
                    "
                    :class="
                      activeIndex == index
                        ? 'bg-gradient-underline-out box-decoration-clone bg-no-repeat pb-[5px] transition-all duration-500 bg-[length:100%] hover:bg-gradient-underline-out'
                        : 'bg-[length:0%]'
                    "
                    class="font-inter-semibold text-[1.125rem] md:text-[1.688rem] 2xl:text-[2.125rem] leading-[1.575rem] md:leading-[2.313rem] 2xl:leading-[3.188rem] cursor-pointer"
                  >
                    {{ blog.title }}</span
                  >
                </div>
                <div
                  class="flex flex-row justify-between mt-[0.9375rem] md:mt-[1.9375rem] font-inter-medium text-[0.875rem] lg:text-[1.125rem] leading-[0.8125rem] md:leading-[1.125rem] lg:leading-[1.463rem]"
                >
                  <span class="w-auto">
                    {{ blog.author }}
                  </span>
                  <span
                    @click="openBlog(blog.link, 'tap_blog_post')"
                    class="lg:w-[7.938rem] cursor-pointer v2-canopas-gradient-text"
                  >
                    Read more
                    <Icon
                      class="arrow fa w-4 h-4 text-pink-300"
                      name="fa6-solid:arrow-right-long"
                      id="leftArrow"
                    />
                  </span>
                </div>
                <hr class="h-[0.063rem] mt-[18px] md:mt-[45px] bg-[#C6C6C6]" />
                <div
                  v-if="activeIndex == index"
                  @click="openBlog(blog.link, 'tap_blog_post')"
                  class="block md:!hidden mt-4 w-full animate-fadeInRight"
                >
                  <img
                    :src="blog.thumbnail"
                    class="object-cover"
                    loading="lazy"
                    :alt="blog.title"
                  />
                  <hr
                    v-if="index != blogs.length - 1"
                    class="md:hidden mt-4 h-[0.063rem] bg-[#C6C6C6]"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="hidden md:block relative flex-[40%]">
        <div>
          <img
            :src="bg"
            loading="lazy"
            class="absolute top-[50%] -right-4 -translate-y-1/2"
            alt="blog-background-image"
          />
          <img
            v-if="activeBlog !== null"
            @click="openBlog(activeBlog.link, 'tap_blog_post')"
            :src="activeBlog.thumbnail"
            class="inline-block absolute inset-0 m-auto w-[90%] cursor-pointer"
            :class="animate ? 'animate-fadeInRight' : ''"
            loading="lazy"
            :alt="activeBlog.title"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import bgMobile from "@/assets/images/blog/bg/bg400.svg";
import bg from "@/assets/images/blog/bg/bg2400.svg";
import axios from "axios";
import Config from "@/config.js";
import { openBlog } from "@/utils.js";

export default {
  data() {
    return {
      openBlog,
      width: 680,
      bgMobile,
      bg,
      blogsURL: Config.BLOG_URL,
      blogs: [],
      activeIndex: 0,
      activeBlog: null,
      animate: false,
    };
  },
  mounted() {
    this.getBlogs();
    this.width = window.innerWidth;
  },
  inject: ["mixpanel"],
  methods: {
    getBlogs() {
      axios
        .get(Config.API_BASE + "/api/blogs")
        .then((blogs) => {
          blogs = blogs.data;
          for (const element of blogs) {
            /** replace - to / in date as safari is not supporting this date format.
             * get date in MMM DD, YYYY format
             */
            element.pubDate = new Date(
              element.pubDate.replace(/-/g, "/"),
            ).toLocaleDateString("en-US", {
              month: "long",
              day: "2-digit",
            });
            element.pubDate = element.pubDate.split(" ", 2);
          }
          this.activeBlog = blogs[0];
          this.blogs = blogs;
        })
        .catch(() => {});
    },
  },
  components: {},
};
</script>
