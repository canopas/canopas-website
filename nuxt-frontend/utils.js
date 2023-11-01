import { useElementVisibility } from "@vueuse/core";
import mixpanel from "mixpanel-browser";

function elementInViewPort(refs) {
  let element;
  for (const key of Object.keys(refs)) {
    if (refs[key] && refs[key].length > 0) {
      for (const ele of refs[key]) {
        if (useElementVisibility(ele).value) {
          element = key;
          break;
        }
      }
    } else if (useElementVisibility(refs[key]).value) {
      element = key;
      break;
    }
  }
  return element;
}

function handleAnimationOnScroll(data) {
  let n, t;
  data &&
    (({ top: n, bottom: t } = data.name.getBoundingClientRect()),
    n < document.documentElement.clientHeight) &&
    0 < t &&
    data.name.classList.add(data.animation);
}

function setGithubStars(contributions, githubRepos) {
  return contributions.forEach((contribution) => {
    contribution.stars = githubRepos
      .filter(
        (repo) =>
          repo.name ==
          contribution.link.slice(contribution.link.lastIndexOf("/") + 1),
      )
      .map((repo) => repo.stargazers_count.toString())[0];
  });
}

function openBlog(link, event) {
  window.open(link, "_blank");
  mixpanel.track(event);
}

function unescapeHTML(escapedHTML) {
  if (escapedHTML.indexOf("&lt;") !== -1) {
    escapedHTML = escapedHTML.replace(/&lt;/g, "<");
  }
  if (escapedHTML.indexOf("&gt;") !== -1) {
    escapedHTML = escapedHTML.replace(/&gt;/g, ">");
  }
  if (escapedHTML.indexOf("&amp;") !== -1) {
    escapedHTML = escapedHTML.replace(/&amp;/g, "&");
  }
  if (escapedHTML.indexOf("&#34;") !== -1) {
    escapedHTML = escapedHTML.replace(/&#34;/g, '"');
  }
  return escapedHTML;
}

function getJobDates() {
  let maxDays = 15;

  // current month day
  const currentDay = new Date().getDate();

  // get start day of month and subtract two days
  const cDate = new Date();
  const startDateOfMonth = new Date(cDate.getFullYear(), cDate.getMonth(), 1);
  startDateOfMonth.setDate(startDateOfMonth.getDate() - 2);

  // if today date is <= 15 then jobPosted is startDateOfMonth, otherwise it's 15
  let jobPosted =
    currentDay <= maxDays
      ? startDateOfMonth
      : new Date(
          startDateOfMonth.setDate(startDateOfMonth.getDate() + maxDays),
        );

  const datePosted = jobPosted.toISOString().split("T")[0];

  // YYYY-MM-DDT00:00
  const validThrough =
    new Date(jobPosted.setDate(jobPosted.getDate() + maxDays + 5))
      .toISOString()
      .split("T")[0] + "T00:00";

  return { datePosted, validThrough };
}

export {
  elementInViewPort,
  handleAnimationOnScroll,
  setGithubStars,
  openBlog,
  unescapeHTML,
  getJobDates,
};