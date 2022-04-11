import { createStore } from "vuex";
import jobs from "./modules/jobs";

export default createStore({
  modules: {
    jobs,
  },
});
