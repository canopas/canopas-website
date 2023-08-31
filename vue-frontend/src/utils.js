import { useElementVisibility } from "@vueuse/core";
import mixpanel from "mixpanel-browser";
function elementInViewPort(refs) {
  var element;
  return (
    Object.keys(refs).forEach((key) => {
      refs[key] && 0 < refs[key].length
        ? refs[key].forEach((ref, index) => {
            useElementVisibility(refs[key][index]).value && (element = key);
          })
        : useElementVisibility(refs[key]).value && (element = key);
    }),
    element
  );
}
function handleAnimationOnScroll(data) {
  var n, t;
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
  window.open(link, "_blank"), mixpanel.track(event);
}
export { elementInViewPort, handleAnimationOnScroll, setGithubStars, openBlog };
