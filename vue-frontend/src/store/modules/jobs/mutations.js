import {
  faApple,
  faAndroid,
  faNodeJs,
  faGolang,
  faVuejs,
} from "@fortawesome/free-brands-svg-icons";
import { library } from "@fortawesome/fontawesome-svg-core";
import config from "@/config.js";

import {
  faGlobe,
  faBullhorn,
  faPenNib,
  faUser,
  faF,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faApple,
  faAndroid,
  faPenNib,
  faGlobe,
  faBullhorn,
  faUser,
  faF,
  faNodeJs,
  faGolang,
  faVuejs
);

function getSEOData(data) {
  var job = data.job;
  var href = data.href;

  var seo_title = job.seo_title
    ? job.seo_title
    : job.title + " jobs at canopas";
  if (href.includes("apply")) {
    seo_title = job.apply_seo_title
      ? job.apply_seo_title
      : "Apply for " + job.title + " jobs at canopas";
  }
  return {
    type: "Jobs Posting Website",
    url: config.BASE_URL + href,
    title: seo_title,
    description: job.seo_description,
  };
}

export default {
  SET_JOBS: (state, jobs) => {
    for (let i = 0; i < jobs.length; i++) {
      let prefix = jobs[i].icon_name.split(" ")[0];
      let iconName = jobs[i].icon_name.split(/-(.*)/);
      let icon = {
        prefix: prefix,
        icon: iconName[1],
      };
      jobs[i].icon_name = icon;
    }
    state.jobs = jobs;
    state.jobsError = null;
  },
  SET_JOB_BY_ID: (state, data) => {
    data.job.seoData = getSEOData(data);
    state.jobById = data.job;
    state.jobsError = null;
  },
  SET_JOBS_ERROR: (state, jobsError) => {
    state.jobsError = jobsError;
  },
};
