import { defineStore } from "pinia";
import axios from "axios";
import config from "~/config";
import { setPostFields } from "~/utils";

export const useAuthorListStore = defineStore("authors", {
  state: () => {
    return {
      totalPosts: 0,
      items: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadAuthorBlogs(showDrafts, slug, start, limit) {
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts ? "is_published=false" : "is_published=true";

        const limitQuery = limit ? "&skip=" + start + "&limit=" + limit : "";

        let url =
          config.API_BASE +
          "/api/posts/author/" +
          slug +
          "?" +
          published +
          limitQuery;

        axios
          .get(url)
          .then((response) => {
            let posts = [];
            response.data.posts.forEach((post) => {
              posts.push(setPostFields(post));
            });
            this.items = posts;
            this.totalPosts = response.data.count;
            this.isLoading = false;
            this.status = posts.length > 0 ? response.status : config.NOT_FOUND;
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

    async loadPaginateAuthorBlogs(showDrafts, slug, start, limit) {
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;

        let published = showDrafts ? "is_published=false" : "is_published=true";

        const limitQuery = limit ? "&skip=" + start + "&limit=" + limit : "";

        let url =
          config.API_BASE +
          "/api/posts/author/" +
          slug +
          "?" +
          published +
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
