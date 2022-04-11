import getters from "./getters";
import actions from "./actions";
import mutations from "./mutations";

const defaultState = {
  jobs: null,
  jobById: null,
  jobsError: null,
};

export default {
  state: defaultState,
  getters,
  actions,
  mutations,
};
