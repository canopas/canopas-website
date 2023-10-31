import { defineStore } from "pinia";
import config from "@/config.js";

export const useContributionStore = defineStore("github-list", {
  state: () => {
    return {
      data: null,
    };
  },
  actions: {
    async fetchStars() {
      return new Promise((resolve, reject) => {
        const eventData =
          `
        self.addEventListener("message", (event) => {
          if (event.data !== "fetchContributionStars") {
            return;
          }
          fetch("` +
          config.API_BASE +
          `/api/github/stars")
            .then((response) => response.json())
            .then((data) => {
              self.postMessage(data.items);
            })
            .catch((error) => {
              console.error(error);
            });
        });
        `;

        const worker = new Worker(
          URL.createObjectURL(
            new Blob([eventData], { type: "text/javascript" }),
          ),
        );
        worker.postMessage("fetchContributionStars");
        worker.onmessage = (event) => {
          this.data = event.data;
          resolve();
        };
        worker.onerror = (error) => {
          reject(error);
        };
      });
    },
  },
});
