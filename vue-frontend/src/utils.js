import { useElementVisibility } from "@vueuse/core";

export function elementInViewPort(refs) {
  var element;
  Object.keys(refs).forEach((key) => {
    if (refs[key] && refs[key].length > 0) {
      refs[key].forEach((ref, index) => {
        if (useElementVisibility(refs[key][index]).value) {
          element = key;
          return;
        }
      });
    } else {
      if (useElementVisibility(refs[key]).value) {
        element = key;
        return;
      }
    }
  });
  return element;
}

export function handleAnimationOnScroll(data) {
  if (data) {
    let { top, bottom } = data.name.getBoundingClientRect();
    let height = document.documentElement.clientHeight;

    if (top < height && bottom > 0) {
      data.name.classList.add(data.animation);
    }
  }
}

export function setGithubStars(contributions, githubRepos) {
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
