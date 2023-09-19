<template>
  <div
    class="tw-my-8 lg:tw-my-[7.813rem] tw-text-black-core/[.87] tw-overflow-hidden"
  >
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center tw-m-auto tw-w-full"
    >
      <div class="tw-container">
        <p
          class="tw-text-black-core/[.03] tw-text-[2.5rem] md:tw-text-[3.75rem] xl:tw-text-[5rem] tw-leading-[1.875rem] md:tw-leading-[4.0625rem] xl:tw-leading-[5rem] tw-text-center"
        >
          Blogs
        </p>
        <h2
          class="-tw-mt-5 md:tw-mt-[-50px] tw-mx-auto xl:tw-w-[65%] tw-pb-9 md:tw-pb-12 tw-text-[1.875rem] tw-leading-[2.438rem] md:tw-text-[2.65rem] md:tw-leading-[3.218rem] lg:tw-text-[3.438rem] lg:tw-leading-[4rem] tw-text-center tw-text-black-core/[0.87] tw-font-roboto-bold"
        >
          Do we have a team with the right skills?
        </h2>
      </div>
      <div class="tw-mt-[1.125rem] lg:tw-w-[65%]">
        <div
          class="tw-w-full tw-text-center tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1rem] md:tw-text-[1.25rem] xl:tw-text-[1.5rem] tw-leading-[1.375rem] md:tw-leading-8 xl:tw-leading-[2.625rem]"
        >
          Well, at least the community says Hell Yeah. Our blogs hosted on
          medium have
          <span class="v2-canopas-gradient-text tw-font-inter-bold">100k+</span>
          minutes monthly reading time and it’s only rising.
        </div>
      </div>
    </div>
    <div class="tw-relative tw-mt-12">
      <img
        v-if="width < 680"
        :src="bgMobile"
        loading="lazy"
        class="md:tw-hidden tw-absolute tw-top-0 -tw-right-10 sm:-tw-right-20 tw-h-full"
        alt="blog-background-image"
      />
    </div>
    <div
      class="tw-flex tw-flex-col md:tw-flex-row md:tw-mt-16 tw-z-[1] lg:tw-min-h-[580px] md:tw-pb-[15%] lg:tw-py-0"
      :class="width < 768 || width > 2000 ? 'tw-container' : ''"
    >
      <div
        class="tw-relative tw-flex tw-flex-col tw-items-center md:tw-items-start md:tw-justify-center md:tw-flex-[60%] md:tw-ml-[5%] lg:tw-ml-[10%] xl:tw-ml-[15%]"
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
              class="tw-hidden md:tw-block md:tw-ml-[82px] 2xl:tw-ml-[95px] tw-h-[0.063rem]"
            />
            <div class="tw-flex tw-gap-8 tw-mt-[18px] md:tw-mt-[45px]">
              <div class="tw-flex tw-flex-col tw-items-center">
                <span
                  class="tw-font-inter-semibold md:tw-font-inter-regular tw-text-[1.5rem] md:tw-text-[2.188rem] 2xl:tw-text-[3.125rem] tw-leading-[1.95rem] md:tw-leading-[2.183rem] 2xl:tw-leading-[4.063rem]"
                >
                  {{ blog.pubDate[1] }}
                </span>
                <span
                  class="tw-text-center font-inter-medium md:tw-font-inter-regular tw-text-[0.75rem] md:tw-text-[1rem] tw-leading-4 2xl:tw-leading-[1.125rem]"
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
                    @click="
                      width > 768 ? openBlog(blog.link, 'tap_blog_post') : ''
                    "
                    :class="
                      activeIndex == index
                        ? 'tw-bg-gradient-underline-out tw-box-decoration-clone tw-bg-no-repeat tw-pb-[5px] tw-transition-all tw-duration-500 tw-bg-[length:100%] hover:tw-bg-gradient-underline-out'
                        : 'tw-bg-[length:0%]'
                    "
                    class="tw-font-inter-semibold tw-text-[1.125rem] md:tw-text-[1.688rem] 2xl:tw-text-[2.125rem] tw-leading-[1.575rem] md:tw-leading-[2.313rem] 2xl:tw-leading-[3.188rem] tw-cursor-pointer"
                  >
                    {{ blog.title }}</span
                  >
                </div>
                <div
                  class="tw-flex tw-flex-row tw-justify-between tw-mt-[0.9375rem] md:tw-mt-[1.9375rem] tw-font-inter-medium tw-text-[0.875rem] lg:tw-text-[1.125rem] tw-leading-[0.8125rem] md:tw-leading-[1.125rem] lg:tw-leading-[1.463rem]"
                >
                  <span class="tw-w-auto">
                    {{ blog.author }}
                  </span>
                  <span
                    @click="openBlog(blog.link, 'tap_blog_post')"
                    class="lg:tw-w-[7.938rem] tw-cursor-pointer v2-canopas-gradient-text"
                  >
                    Read more
                    <font-awesome-icon
                      class="arrow fa tw-w-4 tw-h-4 tw-text-pink-300"
                      :icon="rightArrow"
                      id="leftArrow"
                    />
                  </span>
                </div>
                <hr
                  class="tw-h-[0.063rem] tw-mt-[18px] md:tw-mt-[45px] tw-bg-[#C6C6C6]"
                />
                <div
                  v-if="activeIndex == index"
                  @click="openBlog(blog.link, 'tap_blog_post')"
                  class="tw-block md:!tw-hidden tw-mt-4 tw-w-full tw-animate-fadeInRight"
                >
                  <img
                    :src="blog.thumbnail"
                    class="tw-object-cover"
                    loading="lazy"
                    :alt="blog.title"
                  />
                  <hr
                    v-if="index != blogs.length - 1"
                    class="md:tw-hidden tw-mt-4 tw-h-[0.063rem] tw-bg-[#C6C6C6]"
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
            class="tw-absolute tw-top-[50%] -tw-right-4 -tw-translate-y-1/2"
            alt="blog-background-image"
          />
          <img
            v-if="activeBlog !== null"
            @click="openBlog(activeBlog.link, 'tap_blog_post')"
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
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faArrowRightLong } from "@fortawesome/free-solid-svg-icons";
import bgMobile from "@/assets/images/blog/bg/bg400.svg";
import bg from "@/assets/images/blog/bg/bg2400.svg";
import axios from "axios";
import Config from "@/config.js";
import { openBlog } from "@/utils.js";

export default {
  data() {
    return {
      openBlog,
      rightArrow: faArrowRightLong,
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
  components: {
    FontAwesomeIcon,
  },
};
</script>
