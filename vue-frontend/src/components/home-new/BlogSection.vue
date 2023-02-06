<template>
  <div class="tw-my-[5rem] tw-text-black-core/[.87] tw-overflow-hidden">
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center tw-m-auto tw-w-full"
    >
      <h1
        class="tw-text-[#F8F8F8] tw-text-center tw-text-[3.75rem] xl:tw-text-[5rem] tw-leading-[4.0625rem] xl:tw-leading-[5rem]"
      >
        Blogs
      </h1>
      <h1
        class="tw-mt-[-30px] xl:tw-mt-[-50px] lg:tw-w-[50%] tw-text-center tw-font-roboto-bold tw-text-[1.875rem] md:tw-text-[2.625rem] xl:tw-text-[3.438rem] tw-leading-[2.5rem] md:tw-leading-[2.75rem] xl:tw-leading-[5rem]"
      >
        Do we have a team with the right skills?
      </h1>
      <div class="tw-mt-[1.125rem] lg:tw-w-[65%]">
        <div
          class="tw-w-full tw-text-center tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1rem] md:tw-text-[1.25rem] xl:tw-text-[1.5rem] tw-leading-[1.375rem] md:tw-leading-[2rem] xl:tw-leading-[2.625rem]"
        >
          Well, at least the community says Hell Yeah. Our blogs hosted on
          medium have
          <span class="v2-canopas-gradient-text tw-font-inter-bold">100k+</span>
          minutes monthly reading time and it’s only rising.
        </div>
      </div>
    </div>
    <div class="tw-relative tw-mt-[5rem]">
      <img
        v-if="width < 680"
        :src="bgMobile"
        loading="lazy"
        class="md:tw-hidden tw-absolute tw-top-0 -tw-right-10 sm:-tw-right-20 tw-h-full"
        alt="blog-background-image"
      />
    </div>
    <div
      class="tw-flex tw-flex-col md:tw-flex-row md:tw-mt-[4rem] tw-z-[1]"
      :class="width < 768 || width > 2000 ? 'tw-container' : ''"
    >
      <div
        class="tw-relative tw-flex tw-flex-col tw-items-center md:tw-items-start md:tw-flex-[60%] md:tw-ml-[5%] lg:tw-ml-[10%] xl:tw-ml-[15%]"
        :class="width > 2000 ? 'tw-ml-0' : ''"
      >
        <div
          v-for="(blog, index) in blogs"
          :key="blog"
          @mouseover="activeIndex = index"
          @mouseleave="activeIndex = index"
          @touchstart.passive="activeIndex = index"
          @click.native="mixpanel.track('tap_blog_post')"
          class="tw-w-full"
        >
          <div class="tw-flex tw-flex-col 2xl:tw-pr-16">
            <hr
              v-if="index == 0"
              class="tw-hidden md:tw-block md:tw-ml-[82px] 2xl:tw-ml-[95px] tw-h-[0.063rem] tw-mt-[18px] md:tw-mt-[45px]"
            />
            <div class="tw-flex tw-gap-8 tw-mt-[18px] md:tw-mt-[45px]">
              <div class="tw-flex tw-flex-col tw-items-center tw-mt-[0.3rem]">
                <span
                  class="tw-font-inter-semibold md:tw-font-inter-regular tw-text-[1.5rem] md:tw-text-[2.188rem] 2xl:tw-text-[3.125rem] tw-leading-[1.95rem] md:tw-leading-[2.183rem] 2xl:tw-leading-[4.063rem]"
                >
                  {{ blog.pubDate[1] }}
                </span>
                <span
                  class="tw-text-center font-inter-medium md:tw-font-inter-regular tw-text-[0.75rem] md:tw-text-[1rem] tw-leading-[1rem] 2xl:tw-leading-[1.125rem]"
                >
                  {{ blog.pubDate[0] }}
                </span>
              </div>

              <div
                class="tw-flex tw-flex-col tw-w-full"
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
                    :class="
                      activeIndex == index
                        ? 'tw-bg-gradient-to-b tw-from-[#ff9472] tw-to-[#f2709c] tw-bg-[length:100%_3px] md:tw-bg-[length:100%_5px] tw-bg-no-repeat tw-bg-bottom tw-pb-[3px]'
                        : ''
                    "
                    class="tw-font-inter-semibold tw-text-[1.125rem] md:tw-text-[1.688rem] 2xl:tw-text-[2.125rem] tw-leading-[1.575rem] md:tw-leading-[2.313rem] 2xl:tw-leading-[3.188rem]"
                  >
                    {{ blog.title }}</span
                  >
                </div>
                <div
                  class="tw-flex tw-flex-row tw-justify-between tw-gap-none lg:tw-gap-[22rem] xl:tw-gap-[24rem] 2xl:tw-gap-[38rem] tw-mt-[0.9375rem] md:tw-mt-[1.9375rem] tw-font-inter-medium tw-text-[0.875rem] lg:tw-text-[1.125rem] tw-leading-[0.8125rem] md:tw-leading-[1.125rem] lg:tw-leading-[1.463rem]"
                >
                  <span class="lg:tw-w-[5rem]">
                    {{ blog.author }}
                  </span>
                  <span
                    @click="openBlog(blog)"
                    class="lg:tw-w-[5.938rem] tw-cursor-pointer"
                  >
                    Read more
                  </span>
                </div>
                <hr
                  class="tw-h-[0.063rem] tw-mt-[18px] md:tw-mt-[45px] tw-bg-[#C6C6C6]"
                />
                <div
                  v-if="activeIndex == index"
                  @click="openBlog(blog)"
                  class="tw-block md:!tw-hidden tw-mt-[1rem] tw-w-full tw-animate-fadeInRight"
                >
                  <img
                    :src="blog.thumbnail"
                    class="tw-object-cover"
                    loading="lazy"
                    :alt="blog.title"
                  />
                  <hr
                    v-if="index != blogs.length - 1"
                    class="md:tw-hidden tw-mt-[1rem] tw-h-[0.063rem] tw-bg-[#C6C6C6]"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="tw-hidden md:tw-block tw-relative tw-flex-[40%]">
        <div>
          <img
            :src="bg"
            loading="lazy"
            class="tw-absolute tw-top-[50%] tw-right-[-1rem] -tw-translate-y-1/2"
            alt="blog-background-image"
          />

          <img
            v-if="activeBlog !== null"
            @click="openBlog(activeBlog)"
            :src="activeBlog.thumbnail"
            class="tw-inline-block tw-absolute tw-inset-0 tw-m-auto tw-w-[90%] tw-cursor-pointer"
            :class="animate ? 'tw-animate-fadeInRight' : ''"
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

export default {
  data() {
    return {
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
    openBlog(blog) {
      this.mixpanel.track("tap_blog_post");
      window.open(blog.link, "_blank");
    },
    getBlogs() {
      axios
        .get("https://api.rss2json.com/v1/api.json", {
          params: {
            rss_url: "https://medium.com/feed/canopas",
          },
        })
        .then((response) => {
          var blogs = response.data.items.filter(function (item) {
            return (
              !item.title.includes("Weekly") &&
              !item.title.includes("Newsletter")
            );
          });
          blogs = blogs.slice(0, 3);

          for (let i = 0; i < blogs.length; i++) {
            /** replace - to / in date as safari is not supporting this date format.
             * get date in MMM DD, YYYY format
             */
            blogs[i].pubDate = new Date(
              blogs[i].pubDate.replace(/-/g, "/")
            ).toLocaleDateString("en-US", {
              month: "long",
              day: "2-digit",
            });
            blogs[i].pubDate = blogs[i].pubDate.split(" ", 2);
          }
          this.activeBlog = blogs[0];
          this.blogs = blogs;
        })
        .catch(() => {});
    },
  },
};
</script>
