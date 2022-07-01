<template>
  <div class="blogs-div">
    <div class="background">
      <div class="container">
        <div class="row">
          <div class="flex-div">
            <div class="title-div v2-header-3-text">
              Do we have a team with the right skills?
              <div class="v2-title-3-text mt-4">
                Well, at least the community says Hell Yeah.
              </div>
              <div class="v2-title-3-text mt-2">
                Our blogs hosted on medium have
                <span class="v2-canopas-gradient-text">100k+</span> minutes
                monthly reading time and it’s only rising.
              </div>
            </div>
            <img
              :src="backgroundImage"
              class="background-image pt-5"
              loading="lazy"
            />
          </div>
        </div>
      </div>
    </div>
    <div class="container blog-container">
      <div class="blog-div">
        <a
          class="blog-list"
          v-for="blog in blogs"
          :key="blog"
          :href="blog.link"
          target="_blank"
        >
          <aspect-ratio height="50%" class="blog-image">
            <img :src="blog.thumbnail" class="image" loading="lazy" />
          </aspect-ratio>
          <div class="blog-details">
            <div class="v2-title-3-text title-color">{{ blog.title }}</div>
            <div class="blog-info mt-4">
              <div class="v2-normal-3-text title-color blog-date">
                {{ blog.pubDate }}
              </div>
              <div class="v2-normal-3-text title-color blog-author">
                {{ blog.author }}
              </div>
            </div>
          </div>
        </a>
      </div>
      <div class="view-button">
        <a class="v2-button v2-normal-2-text" :href="blogsURL" target="_blank">
          <span>VIEW ALL</span>
          <font-awesome-icon
            class="arrow fa"
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
  },
  mounted() {
    this.getBlogs();
  },
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
            blogs[i].pubDate = String(blogs[i].pubDate);
          }

          blogs = blogs.slice(0, 3);
          this.blogs = blogs;
        })
        .catch(() => {});
    },
  },
};
</script>

<style lang="scss" scoped>
.background {
  background: linear-gradient(
    357.22deg,
    rgb(255, 131, 91, 0.1) -40.56%,
    rgb(242, 112, 156, 0.1) 62.45%,
    rgb(255, 255, 255, 0.1) 98.58%
  );
  border-radius: 0px 0px 35px 35px;
  position: relative;
  z-index: -1;
}

.row {
  padding-bottom: 320px;
}

.flex-div {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.blog-container {
  margin-top: -281px;
  z-index: 1;
}

.v2-header-3-text {
  text-align: center;
}

.background-image {
  display: none;
}

.blog-div {
  display: flex;
  flex-direction: column;
}

.blog-list {
  display: flex;
  flex-direction: column;
  margin: 10px;
  flex: 1 0 0%;
}

.blog-list:active {
  transform: scale(0.98);
}

.blog-image {
  background: -webkit-linear-gradient(#ffffff, #ffffff) padding-box,
    linear-gradient(to bottom, #ff9472, #f2709c) border-box;
  border: 1px solid transparent !important;
  border-radius: 20px 20px 0 0;
  filter: drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25));
}

.image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 20px 20px 0 0;
}

.blog-details {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex: 1;
  padding: 20px 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.5);
  background: #3d3d3d;
  border-radius: 0 0 20px 20px;
}

.blog-info {
  width: 100%;
  overflow: hidden;
}

.blog-date {
  float: left;
}

.blog-author {
  float: right;
}

.title-color {
  color: #ffffff;
}

.view-button {
  display: flex;
  justify-content: flex-end;
  margin: 20px;
}

.v2-button {
  display: flex;
  align-items: center;
}

@include media-breakpoint-up(sm) {
  .blog-list {
    margin: 20px;
  }

  .blog-details {
    padding: 20px;
  }
}

@include media-breakpoint-up(md) {
  .flex-div {
    flex-direction: row;
  }

  .title-div {
    width: 70%;
  }

  .v2-header-3-text {
    text-align: start;
  }

  .background-image {
    display: block;
  }

  .row {
    margin-right: -30px;
    padding-bottom: 150px;
  }
}

@include media-breakpoint-up(lg) {
  .title-div {
    width: 60%;
  }

  .blog-div {
    flex-flow: row;
  }
  .blog-container {
    margin-top: -201px;
  }
}

@include media-breakpoint-up(xl) {
  .blog-container {
    margin-top: -281px;
  }
}
@supports (-webkit-touch-callout: none) {
  .v2-header-3-text {
    letter-spacing: -2px;
  }
}
</style>
