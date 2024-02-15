import { defineStore } from "pinia";
import axios from "axios";
import config from "~/config";
import { setPostFields } from "~/utils";

export const useBlogListStore = defineStore("blog-list", {
  state: () => {
    return {
      totalPosts: 0,
      items: null,
      featuredItems: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadResources(showDrafts, resources, start, limit) {
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts
          ? "&publicationState=preview"
          : "&publicationState=live";

        const limitQuery = limit
          ? "&pagination[start]=" + start + "&pagination[limit]=" + limit
          : "";

        let url =
          config.STRAPI_URL +
          "/v1/posts?populate=deep" +
          published +
          "&filters[is_resource]=" +
          resources +
          "&fields[0]=title&fields[1]=slug&fields[2]=published_on&fields[3]=summary&fields[4]=reading_time" +
          limitQuery;

        axios
          .get(config.STRAPI_URL + "/favicon.ico")
          .then(() => {
            axios
              .get(url)
              .then((response) => {
                const resp = response.data.data.attributes;
                let posts = [];
                let featuredPosts = [];
                resp.posts.forEach((post) => {
                  posts.push(setPostFields(post));
                });
                resp.featuredPosts.forEach((post) => {
                  featuredPosts.push(setPostFields(post));
                });
                this.items = posts;
                this.featuredItems = featuredPosts;
                this.totalPosts = resp.count;
                this.isLoading = false;
                this.status = response.status;
                resolve();
              })
              .catch((error) => {
                this.error = error;
                this.isLoading = false;
                this.status = config.NOT_FOUND;
                reject(error);
              });
          })
          .catch((error) => {
            this.status = config.NOT_FOUND;
            reject(error);
          });
      });
    },

    async loadPaginateResources(showDrafts, resources, start, limit) {
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts
          ? "&publicationState=preview"
          : "&publicationState=live";

        const limitQuery = limit
          ? "&pagination[start]=" + start + "&pagination[limit]=" + limit
          : "";

        let url =
          config.STRAPI_URL +
          "/v1/paginate?populate=deep" +
          published +
          "&filters[is_resource]=" +
          resources +
          "&fields[0]=title&fields[1]=slug&fields[2]=published_on&fields[3]=summary&fields[4]=reading_time" +
          limitQuery;

        axios
          .get(config.STRAPI_URL + "/favicon.ico")
          .then(() => {
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
                reject(error);
              });
          })
          .catch((error) => {
            this.status = config.NOT_FOUND;
            reject(error);
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
