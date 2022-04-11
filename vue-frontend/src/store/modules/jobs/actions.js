import axios from "axios";
import config from "@/config.js";

export default {
  getJobs({ commit }) {
    return new Promise((resolve, reject) => {
      axios
        .get(config.API_BASE + "/api/careers")
        .then((response) => {
          commit("SET_JOBS", response.data);
          resolve(response);
        })
        .catch((error) => {
          commit("SET_JOBS_ERROR", error);
          reject(error);
        });
    });
  },
  getJobsById({ commit }, data) {
    return new Promise((resolve, reject) => {
      axios
        .get(config.API_BASE + "/api/careers/" + data.jobId)
        .then((response) => {
          var res = {
            job: response.data,
            href: data.href,
          };
          commit("SET_JOB_BY_ID", res);
          resolve(response);
        })
        .catch((error) => {
          commit("SET_JOBS_ERROR", error);
          reject(error);
        });
    });
  },
};
