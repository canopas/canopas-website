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

        let published = showDrafts ? "is_published=false" : "is_published=true";

        const limitQuery = limit ? "&skip=" + start + "&limit=" + limit : "";

        let url =
          config.API_BASE +
          "/api/posts?" +
          published +
          "&is_resource=" +
          resources +
          limitQuery;

        axios
          .get(url)
          .then((response) => {
            const resp = response.data;
            let posts = [];
            let featuredPosts = [];

            resp.posts.forEach((post) => {
              posts.push(setPostFields(post));
            });

            if (resources) {
              resp.featuredPosts.forEach((post) => {
                featuredPosts.push(setPostFields(post));
              });
            }

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
      });
    },

    async loadPaginateResources(showDrafts, resources, start, limit) {
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts ? "is_published=false" : "is_published=true";

        const limitQuery = limit ? "&skip=" + start + "&limit=" + limit : "";

        let url =
          config.API_BASE +
          "/api/posts?" +
          published +
          "&is_resource=" +
          resources +
          limitQuery;

        axios
          .get(url)
          .then((response) => {
            let posts = [];
            response.data.posts.forEach((post) => {
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
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;
        let url = config.API_BASE + "/api/posts/" + slug;

        axios
          .get(url)
          .then((response) => {
            this.item = setPostFields(response.data);
            this.isLoading = false;
            this.status = response.status;
            resolve();
          })
          .catch((error) => {
            console.error(error);
            this.error = error;
            this.isLoading = false;
            this.status = config.NOT_FOUND;
            reject(error);
          });
      });
    },
  },
});
