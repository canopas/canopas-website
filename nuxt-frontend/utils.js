import { useElementVisibility } from "@vueuse/core";
import mixpanel from "mixpanel-browser";
import Avatar from "./assets/images/user.png";
import icon from "./assets/images/icon.svg";

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

function getDiffrentWidthImages(response) {
  let imageUrls = response.data;
  let slides = [];
  slides = imageUrls.map((set) => ({
    image: parseImageUrls(set.image_urls),
  }));
  return slides;
}

function parseImageUrls(imageUrls) {
  return imageUrls.split(",").map((url) => url.trim());
}

function setPostFields(post) {
  post = post.attributes ? post.attributes : post;
  const publishedDate = post.published_on;
  const [date] = formateDate(publishedDate);
  post.published_on = date || "Draft";

  const postImg = post.image?.data?.attributes
    ? post.image.data?.attributes
    : post.image;
  post.image_url = postImg?.url || icon;
  post.alternative_text = postImg?.alternative_text || post.title;

  const author = post.author?.data?.attributes
    ? post.author.data?.attributes
    : post.author;

  const authorImg = author?.image?.data?.attributes
    ? author?.image.data?.attributes
    : author?.image;

  post.author = {
    username: author?.username || "user",
    name: author?.name || "author",
    image: authorImg?.url || Avatar,
    alt_text: author ? author.username + " image" : "author",
    bio: author?.bio || "",
    role: author?.role || "Editor for Canopas",
  };

  post.recommended_posts = post.recommendedPosts;

  return post;
}

// Formate date and time from millis
function formateDate(date) {
  if (!date) return [null, null];
  const newDate = new Date(date);
  const userTimezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
  const formattedDate = newDate.toLocaleDateString("en-US", {
    timeZone: userTimezone,
    month: "short",
    day: "numeric",
    year: "numeric",
  });

  // time formate
  const formattedTime = newDate.toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
  });
  return [formattedDate, formattedTime];
}

function isValidPhone(phoneNumber) {
  const phoneRegex = /^[0-9]{10}$/;
  return !phoneRegex.test(phoneNumber);
}

function isValidEmail(email) {
  const emailRegx =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return !emailRegx.test(email);
}

export {
  elementInViewPort,
  handleAnimationOnScroll,
  setGithubStars,
  openBlog,
  unescapeHTML,
  getJobDates,
  getDiffrentWidthImages,
  setPostFields,
  isValidPhone,
  isValidEmail,
};
