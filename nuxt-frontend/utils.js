import { useElementVisibility } from "@vueuse/core";
import mixpanel from "mixpanel-browser";
import config from "./config";
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

function setPostFields(post, slug) {
  const publishedDate = post.attributes.published_on;
  const [date] = formateDate(publishedDate);
  post.attributes.published_on = date || "Draft";
  post.attributes.readingTime = getReadingTime(post.attributes.content);

  post.attributes.image_url =
    post.attributes.image.data?.attributes.url || icon;
  post.attributes.alternativeText =
    post.attributes.image.data?.attributes.alternativeText ||
    post.attributes.title;

  const author = post.attributes.author.data?.attributes;
  post.attributes.authorName = author?.name || "author";
  post.attributes.authorImage = author?.image.data
    ? author.image.data?.attributes.url
    : Avatar;
  post.attributes.authorAltText = author
    ? author.username + " image"
    : "author";
  post.attributes.authorBio = author?.bio || "";
  post.attributes.authorRole = author?.role || "Editor for Canopas";

  if (slug && post.attributes.tags[0]) {
    post.attributes.tags.map((tag) => {
      if (tag.slug == slug) {
        post.tagName = tag.name;
      }
    });
  }

  let newPost = post.attributes;
  newPost.id = post.id;
  newPost.tagName = post.tagName;
  return newPost;
}
// Calculate reading time
function getReadingTime(content) {
  if (!content) return;
  const numberOfWords = content
    .replace(/<\/?[^>]+(>|$)/g, "")
    .split(/\s/g).length;
  return Math.ceil(numberOfWords / config.WORDS_PER_MINUTE);
}

// Formate date and time from millis
function formateDate(date) {
  if (!date) return [null, null];
  const newDate = new Date(date);
  const formattedDate = newDate.toLocaleDateString("en-US", {
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

function filterPostsByCategoryAndTag(post, posts) {
  const relatedPosts = [];
  const postTags = post?.tags.map((tag) => tag.name);
  const postCategoryName = post?.category.data?.attributes.name ?? null;

  if (posts) {
    if (postCategoryName || postTags?.length > 0) {
      posts.forEach((post) => {
        let index = 0;

        const relatedTags = post.tags.map((tag) => tag.name);
        const relatedCategoryName =
          post.category?.data?.attributes.name ?? null;

        index +=
          postCategoryName &&
          relatedCategoryName &&
          postCategoryName === relatedCategoryName
            ? 3
            : 0;

        index += postTags.filter((tag) => relatedTags.includes(tag)).length;

        if (index > 0) {
          post.index = index;
          relatedPosts.push(post);
        }
      });
    }
  }
  return relatedPosts.filter(
    (p) =>
      !post.is_resource || !p.is_resource || p.is_resource == post.is_resource,
  );
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
  filterPostsByCategoryAndTag,
};
