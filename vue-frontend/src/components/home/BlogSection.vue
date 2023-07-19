<template>
  <div>
    <div
      class="tw-bg-gradient-to-t tw-from-orange-300/10 tw-via-pink-300/10 tw-to-white tw-rounded-[35px] tw-relative tw--z-[1]"
    >
      <div class="tw-container">
        <div class="tw-pb-80 md:tw--mr-[30px] md:tw-pb-[150px]">
          <div class="tw-flex tw-flex-col tw-justify-between md:tw-flex-row">
            <div
              class="md:tw-w-[70%] lg:tw-w-[60%] v2-header-3-text tw-text-center md:tw-text-start"
            >
              Do we have a team with the right skills?
              <div class="v2-title-3-text tw-mt-6">
                Well, at least the community says Hell Yeah.
              </div>
              <div class="v2-title-3-text tw-mt-2">
                Our blogs hosted on medium have
                <span class="v2-canopas-gradient-text">100k+</span> minutes
                monthly reading time and it’s only rising.
              </div>
            </div>
            <img
              :src="backgroundImage"
              class="tw-hidden md:tw-block tw-pt-12 tw-h-full tw-w-[260px]"
              loading="lazy"
              alt="blog-background-image"
            />
          </div>
        </div>
      </div>
    </div>
    <div
      class="tw-container tw--mt-[281px] tw-z-[1] lg:tw--mt-[201px] xl:tw--mt-[281px]"
    >
      <div class="tw-flex tw-flex-col lg:tw-flex-row">
        <a
          class="tw-flex tw-flex-col tw-m-2.5 sm:tw-m-5 tw-flex-[1_1_0%] active:tw-scale-[0.98]"
          v-for="blog in blogs"
          :key="blog"
          :href="blog.link"
          target="_blank"
          @click.native="mixpanel.track('tap_home_blog_post')"
        >
          <aspect-ratio
            height="50%"
            class="tw-drop-shadow-[0_4px_4px_rgba(0,0,0,0.25)] tw-rounded-t-x-lg tw-border-[1px] tw-border-solid tw-border-transparent tw-bg-gradient-to-t tw-from-orange-300 tw-to-pink-300"
          >
            <img
              :src="blog.thumbnail"
              class="tw-w-full tw-h-full tw-object-cover tw-rounded-t-x-lg"
              loading="lazy"
              :alt="blog.title"
            />
          </aspect-ratio>
          <div
            class="tw-flex tw-flex-col tw-justify-between tw-flex-1 tw-px-2.5 tw-py-5 sm:tw-p-5 tw-shadow-[0_4px_4px_rgba(0,0,0,0.5)] tw-bg-black-900 tw-rounded-b-x-lg"
          >
            <div class="v2-title-3-text tw-text-white">{{ blog.title }}</div>
            <div class="tw-w-full tw-overflow-hidden tw-mt-6">
              <div class="v2-normal-3-text tw-text-white tw-float-left">
                {{ blog.pubDate }}
              </div>
              <div class="v2-normal-3-text tw-text-white tw-float-right">
                {{ blog.author }}
              </div>
            </div>
          </div>
        </a>
      </div>
      <div class="tw-flex tw-justify-end tw-m-5">
        <a
          class="v2-button tw-flex tw-item-center v2-normal-2-text"
          :href="blogsURL"
          target="_blank"
          @click.native="mixpanel.track('tap_home_view_all_blogs')"
        >
          <span>VIEW ALL</span>
          <font-awesome-icon
            class="arrow fa md:tw-mt-1"
            icon="arrow-right"
            id="leftArrow"
          />
        </a>
      </div>
    </div>
  </div>
</template>

<script type="module">
import backgroundImage from "@/assets/images/theme/blog-bckground.svg";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import AspectRatio from "@/components/utils/AspectRatio.vue";
import Config from "@/config.js";
import axios from "axios";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      blogsURL: Config.BLOG_URL,
      blogs: "",
    };
  },
  components: {
    FontAwesomeIcon,
    AspectRatio,
  },
  mounted() {
    this.getBlogs();
  },
  inject: ["mixpanel"],
  methods: {
    getBlogs() {
      axios
        .get("https://api.rss2json.com/v1/api.json", {
          params: {
            rss_url: "https://medium.com/feed/canopas",
          },
        })
        .then((response) => {
          var blogs = response.data.items.filter(function (item) {
            return !item.title.includes("Weekly");
          });

          for (let i = 0; i < blogs.length; i++) {
            /** date in MMM DD, YYYY format */
            blogs[i].pubDate = new Date(blogs[i].pubDate).toLocaleDateString(
              "en-US",
              {
                month: "short",
                day: "numeric",
                year: "numeric",
              },
            );
          }

          blogs = blogs.slice(0, 3);
          this.blogs = blogs;
        })
        .catch(() => {});
    },
  },
};
</script>
