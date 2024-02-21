import { defineStore } from "pinia";
import axios from "axios";
import config from "~/config";
import { setPostFields } from "~/utils";

export const useAuthorListStore = defineStore("authors", {
  state: () => {
    return {
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
          "&filters[author][username]=" +
          slug +
          "&fields[0]=title&fields[1]=slug&fields[2]=published_on&fields[3]=summary&fields[4]=reading_time&fields[5]=tags" +
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
                this.status =
                  posts.length > 0 ? response.status : config.NOT_FOUND;
                resolve();
              })
              .catch((error) => {
                this.error = error;
                this.isLoading = false;
                this.status = config.NOT_FOUND;
                resolve();
              });
          })
          .catch((error) => {
            this.error = error;
            this.status = config.NOT_FOUND;
            reject(error);
          });
      });
    },
  },
});
