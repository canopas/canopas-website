import { defineStore } from "pinia";
import axios from "axios";
import config from "~/config";
import { setPostFields } from "~/utils";

export const useBlogListStore = defineStore("blog-list", {
  state: () => {
    return {
      items: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadResources(showDrafts, resources) {
      return new Promise((resolve) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts
          ? "&publicationState=preview"
          : "&publicationState=live";

        let url =
          config.STRAPI_URL +
          "/v1/posts?populate=deep" +
          published +
          "&filters[is_resource]=" +
          resources;

        axios
          .get(url)
          .then((response) => {
            let posts = [];
            response.data.data.forEach((post) => {
              posts.push(setPostFields(post));
            });
            this.items = posts;
            this.isLoading = false;
            this.status = response.status;
            resolve();
          })
          .catch((error) => {
            this.error = error;
            this.isLoading = false;
            this.status = config.NOT_FOUND;
            resolve();
          });
      });
    },
  },
});

export const useBlogDetailStore = defineStore("resources-detail", {
  state: () => {
    return {
      item: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadResource(slug) {
      return new Promise((resolve) => {
        this.isLoading = true;
        this.error = null;
        let url = config.STRAPI_URL + "/v1/posts/" + slug + "?populate=deep";

        axios
          .get(url)
          .then((response) => {
            this.item = setPostFields(response.data.data);
            this.isLoading = false;
            this.status = response.status;
            resolve();
          })
          .catch((error) => {
            console.error(error);
            this.error = error;
            this.isLoading = false;
            this.status = config.NOT_FOUND;
            resolve();
          });
      });
    },
  },
});

// fetch posts other than this post by category name
export const useRecommandedBlogStore = defineStore("recommanded-blog", {
  state: () => {
    return {
      items: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadRecommandedBlog(slug, showDrafts) {
      return new Promise((resolve) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts
          ? "&publicationState=preview"
          : "&publicationState=live";
        let url =
          config.STRAPI_URL +
          "/v1/posts?filters[slug][$ne]=" +
          slug +
          published;

        axios
          .get(url)
          .then((response) => {
            let posts = [];
            response.data.data.forEach((post) => {
              posts.push(setPostFields(post));
            });
            this.items = posts;
            this.isLoading = false;
            this.status = response.status;
            resolve();
          })
          .catch((error) => {
            this.error = error;
            this.isLoading = false;
            this.status = config.NOT_FOUND;
            resolve();
          });
      });
    },
  },
});

// fetch featured posts
export const useFeaturedBlogStore = defineStore("featured-blog", {
  state: () => {
    return {
      items: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadFeaturedBlog(showDrafts) {
      return new Promise((resolve) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts
          ? "&publicationState=preview"
          : "&publicationState=live";
        let url =
          config.STRAPI_URL +
          "/v1/posts?filters[is_featured][$eq]=true" +
          published;

        axios
          .get(url)
          .then((response) => {
            let posts = [];
            response.data.data.forEach((post) => {
              posts.push(setPostFields(post));
            });
            this.items = posts;
            this.isLoading = false;
            this.status = response.status;
            resolve();
          })
          .catch((error) => {
            this.error = error;
            this.status = config.NOT_FOUND;
            this.isLoading = false;
            resolve();
          });
      });
    },
  },
});
