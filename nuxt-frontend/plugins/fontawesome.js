import { library, config } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
  faThumbsUp,
  faTimes,
  faAngleUp,
  faAngleDown,
  faAngleRight,
  faAngleLeft,
} from "@fortawesome/free-solid-svg-icons";
import { faMediumM, faGithub } from "@fortawesome/free-brands-svg-icons";

config.autoAddCss = false;

library.add(
  faTimes,
  faArrowRight,
  faStar,
  faArrowLeft,
  faThumbsUp,
  faAngleUp,
  faAngleDown,
  faAngleRight,
  faAngleLeft,
  faMediumM,
  faGithub
);

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component("font-awesome-icon", FontAwesomeIcon, {});
});
