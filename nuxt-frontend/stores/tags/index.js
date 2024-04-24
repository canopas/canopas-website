import { defineStore } from "pinia";
import axios from "axios";
import config from "~/config";
import { setPostFields } from "~/utils";

export const useTagListStore = defineStore("tag-list", {
  state: () => {
    return {
      items: null,
      status: 0,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadTagBlogs(slug) {
      return new Promise((resolve, reject) => {
        this.isLoading = true;
        this.error = null;

        let url =
          config.API_BASE + "/api/posts/tags/" + slug + "?is_published=true";

        axios
          .get(url)
          .then((response) => {
            let posts = [];
            response.data.forEach((post) => {
              posts.push(setPostFields(post));
            });
            this.items = posts;
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
  },
});
