import { defineStore } from "pinia";
import axios from "axios";
import config from "@/config.js";

export const useJobListStore = defineStore("jobList", {
  state: () => {
    return {
      items: null,
      error: null,
      isLoading: false,
    };
  },
  actions: {
    async loadJobs() {
      return new Promise((resolve) => {
        this.isLoading = true;
        this.error = null;

        axios
          .get(config.API_BASE + "/api/careers")
          .then((response) => {
            this.items = this.mapJobIcons(response.data);
            this.isLoading = false;
            resolve();
          })
          .catch((error) => {
            this.error = error;
            this.isLoading = false;
            resolve();
          });
      });
    },

    mapJobIcons(jobs) {
      for (let i = 0; i < jobs.length; i++) {
        let prefix = jobs[i].icon_name.split(" ")[0];
        let iconName = jobs[i].icon_name.split(/-(.*)/);
        let icon = {
          prefix: prefix,
          icon: iconName[1],
        };
        jobs[i].icon_name = icon;
      }
      return jobs;
    },
  },
});
