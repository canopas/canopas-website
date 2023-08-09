<template>
  <div class="tw-container tw-my-[50px] lg:tw-mt-0 md:tw-w-full">
    <form>
      <div
        class="tw-rounded-2xl tw-border tw-border-slate-200 tw-bg-white tw-shadow-[0px_0px_45px_rgba(0,0,0,0.1)]"
      >
        <div class="tw-pt-5 tw-pb-20 tw-px-8 lg:tw-px-12">
          <div>
            <div class="tw-relative tw-mb-5 md:tw-mb-5 tw-pt-8 tw-text-left">
              <input
                type="text"
                id="username"
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-font-inter-regular tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
                name="username"
                required
                autocomplete="given-username"
                v-model="name"
                placeholder=" "
                @input="
                  showNameValidationError =
                    $event.target.value.trim().length === 0
                "
                @click.native="mixpanel.track('tap_contact_name_input')"
              />
              <label
                for="username"
                class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-text-black-core/[.6] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-black-core peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
                >Your name</label
              >
              <span v-if="showNameValidationError" class="error tw-text-red-600"
                >Name is required</span
              >
            </div>
            <div class="tw-relative tw-mb-5 md:tw-mb-5 tw-pt-8 tw-text-left">
              <input
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-font-inter-regular tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
                type="text"
                name="email"
                id="email"
                required
                autocomplete="given-email"
                v-model="email"
                placeholder=" "
                @input="
                  showEmailValidationError =
                    $event.target.value.trim().length === 0
                "
                @blur="showValidEmailError = isValidEmail()"
                @click.native="mixpanel.track('tap_contact_email_input')"
              />
              <label
                for="email"
                class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-text-black-core/[.6] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-black-core/[0.6] peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
                >Your email</label
              >
              <span
                v-if="showEmailValidationError"
                class="error tw-text-red-600"
                >Email is required</span
              >
              <span
                v-if="email.trim().length != 0 && showValidEmailError"
                class="error tw-text-red-600"
                >Please enter valid email address</span
              >
            </div>
            <div
              class="tw-relative md:tw-col-span-2 tw-mb-5 md:tw-mb-5 tw-pt-8 tw-text-left"
            >
              <textarea
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-min-h-[50px] md:tw-min-h-[90px] tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-font-inter-regular tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
                id="project"
                name="project"
                rows="3"
                required
                autocomplete="given-project-info"
                v-model="projectInfo"
                placeholder=" "
                @input="
                  showProjectInfoValidationError =
                    $event.target.value.trim().length === 0
                "
                @click.native="mixpanel.track('tap_contact_project_info_input')"
              ></textarea>
              <label
                for="project"
                class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-text-black-core/[.6] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-black-core peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
                >Tell us about your project</label
              >
              <span
                v-if="showProjectInfoValidationError"
                class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div
              class="tw-relative md:tw-col-span-2 tw-pt-6 md:tw-pt-8 tw-text-left"
            >
              <input
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-font-inter-regular tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
                type="text"
                name="reference"
                id="reference"
                required
                autocomplete="given-reference"
                v-model="reference"
                placeholder=" "
                @input="
                  showReferenceValidationError =
                    $event.target.value.trim().length === 0
                "
                @click.native="mixpanel.track('tap_contact_reference_input')"
              />
              <label
                for="reference"
                class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-text-black-core/[.6] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-black-core peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
                >How did you find us?</label
              >
              <span
                v-if="showReferenceValidationError"
                class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div class="tw-relative tw-pt-8">
              <div ref="invest-list" class="tw-flex">
                <button
                  class="tw-flex tw-items-center tw-justify-between tw-mt-2.5 md:tw-mt-3 tw-mx-0 tw-w-full tw-border-b tw-border-slate-400 tw-bg-none tw-px-0 tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-whitespace-nowrap tw-transition tw-duration-150 tw-ease-in-out focus:tw-outline-0 active:tw-outline-0 focus:tw-shadow-none active:tw-shadow-none focus:tw-ring-0 active:tw-ring-0 focus:tw-bg-transparent active:tw-bg-transparent active:tw-text-black-core/[0.6]"
                  type="button"
                  id="invest"
                  name="invest"
                  required
                  @click="toggleList"
                >
                  <label
                    for="invest"
                    class="tw-absolute tw-top-8 md:tw-top-8 tw-left-0 tw-text-black-core/[0.6] tw-transform tw-duration-300 placeholder:tw-!text-black-core/[0.6]"
                    :class="[
                      floatable
                        ? 'tw--translate-y-4 xl:tw--translate-y-[0.5rem] tw-origin-[0] tw-scale-75'
                        : 'tw-translate-y-0 tw-scale-100',
                      floatable && showList
                        ? 'tw-text-black-core'
                        : 'tw-text-black-core/[0.6]',
                    ]"
                    @click="floatable = !floatable"
                    >I'll invest</label
                  >
                  <span
                    class="tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem]"
                    >{{ invest }}</span
                  >
                  <font-awesome-icon
                    class="fab tw-mb-3 tw-w-[15px] md:tw-w-[25px] tw-h-[15px] md:tw-h-[25px] tw-text-black-core/[0.6]"
                    :icon="showList ? faCaretUp : faCaretDown"
                  />
                </button>
                <ul
                  v-if="showList"
                  class="tw-absolute tw-mt-9 md:tw-mt-12 tw-w-full tw-min-w-max tw-z-[5] tw-border-none tw-rounded-[10px] tw-shadow-[0px_5px_15px_rgba(0,0,0,0.5)] tw-bg-white tw-py-1 tw-list-none tw-text-left tw-transition-all tw-delay-150 tw-duration-300"
                >
                  <li
                    v-for="option in investOptions"
                    :key="option"
                    @click="setOption(option)"
                  >
                    <span
                      class="tw-block tw-w-full tw-py-2 tw-px-4 tw-whitespace-nowrap tw-text-black-core/[.6] hover:tw-bg-[#000]/[.2] hover:tw-text-[#000] tw-font-inter-regular tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-leading-[1.21rem] md:tw-leading-[2.0625rem]"
                      >{{ option }}</span
                    >
                  </li>
                </ul>
              </div>
              <span
                v-if="showInvestValidationError"
                class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div
              class="tw-mt-8 tw-text-base tw-font-inter-regular tw-text-black-core/[0.87]"
            >
              We sign NDA for all of our projects.
            </div>
          </div>
          <div class="tw-flex tw-justify-center tw-py-5 lg:tw-py-8">
            <img
              v-if="showLoader"
              :src="loaderImage"
              class="tw-w-16 tw-h-16"
              alt="loader-image"
            />
            <div v-else>
              <div
                class="tw-relative tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.436875rem] md:tw-leading-[1.386875rem] tw-text-center"
              >
                <div
                  class="tw-absolute -tw-top-10 sm:-tw-top-8 xl:-tw-top-[1.875rem] 2xl:-tw-top-[1.875rem] tw-text-center -tw-right-32 sm:-tw-right-[10.5rem] md:-tw-right-44 lg:-tw-right-40 2xl:-tw-right-[10.5rem] tw-w-[250px] sm:tw-w-max"
                >
                  <span
                    v-if="showErrorMessage"
                    class="tw-flex tw-text-center tw-text-red-600"
                    :class="
                      errorMessage == 'Invalid Recaptcha score'
                        ? '!tw-mt-2 md:!-tw-mt-2 !tw-mr-4 sm:!tw-mr-16 md:!tw-mr-12 lg:!tw-mr-10 xl:!tw-mr-[2.8rem] 2xl:!tw-mr-14 '
                        : ''
                    "
                    >{{ errorMessage }}</span
                  >
                </div>
                <button
                  id="submit"
                  ref="recaptcha"
                  class="tw-absolute tw-top-[-13px] sm:-tw-top-5 tw-right-[-122px] md:tw-right-[-170px] tw-w-max tw-rounded-full tw-py-2.5 md:tw-py-5 tw-px-20 md:tw-px-28 tw-text-center gradient-btn consultation-btn"
                  @click.prevent="submitForm()"
                >
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6"
                    aria-hidden="true"
                  />
                  <span>Submit</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import axios from "axios";
import config from "@/config.js";
import CalendlyIframe from "./CalendlyIframe.vue";
import loaderImage from "@/assets/images/theme/small-loader.svg";
import { faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faPhoneVolume,
  faEnvelope,
  faCalendarAlt,
} from "@fortawesome/free-solid-svg-icons";
export default {
  data() {
    return {
      loaderImage,
      phone: faPhoneVolume,
      mail: faEnvelope,
      calendarIcon: faCalendarAlt,
      faCaretDown,
      faCaretUp,
      showList: false,
      floatable: false,
      name: "",
      email: "",
      projectInfo: "",
      reference: "",
      invest: "",
      investOptions: [
        "< USD 50000",
        "USD 50000 - USD 100000",
        "USD 100000 - USD 250000",
        "> USD 250000",
      ],
      showNameValidationError: false,
      showEmailValidationError: false,
      showValidEmailError: false,
      showProjectInfoValidationError: false,
      showReferenceValidationError: false,
      showInvestValidationError: false,
      openCalendlyIframeModal: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      showErrorMessage: false,
      contactType: 1,
    };
  },
  components: {
    FontAwesomeIcon,
    CalendlyIframe,
  },
  inject: ["mixpanel"],
  mounted() {
    document.addEventListener("click", this.closePopUps);
  },
  unmounted() {
    document.removeEventListener("click", this.closePopUps);
  },

  methods: {
    isValidEmail() {
      var emailRegx =
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return !emailRegx.test(this.email);
    },
    validateForm() {
      this.showNameValidationError = this.name.trim().length === 0;
      this.showEmailValidationError = this.email.trim().length === 0;
      this.showProjectInfoValidationError =
        this.projectInfo.trim().length === 0;
      this.showReferenceValidationError = this.reference.trim().length === 0;
      this.showInvestValidationError = this.invest.trim().length === 0;
      return (
        this.showNameValidationError ||
        this.showEmailValidationError ||
        this.showValidEmailError ||
        this.showProjectInfoValidationError ||
        this.showReferenceValidationError ||
        this.showInvestValidationError
      );
    },
    submitForm() {
      if (!this.validateForm()) {
        this.mixpanel.track("tap_contact_submit_button");
        let formData = {
          name: this.name,
          email: this.email,
          project_info: this.projectInfo
            ? this.projectInfo.replace(/\./g, ".\n")
            : "NA",
          reference: this.reference,
          invest: this.invest,
          send_mail_to_client: config.IS_PROD,
        };
        grecaptcha.enterprise.ready(() => {
          grecaptcha.enterprise
            .execute(import.meta.env.VITE_RECAPTCHA_SITE_KEY, {
              action: "verify",
            })

            .then((token) => {
              formData.token = token;
              axios
                .post(config.API_BASE + "/api/send-contact-mail", formData)
                .then(() => {
                  gtag("event", "conversion", {
                    send_to: "AW-11157168108/OpYlCPjY4poYEOy_k8gp",
                  });
                  this.$router.push({
                    path: "/thank-you",
                  });
                  localStorage.setItem(
                    "client-name",
                    JSON.stringify(formData.name),
                  );
                  this.resetForm();
                })
                .catch((err) => {
                  if (err.response.status == 401) {
                    this.errorMessage = "Invalid recaptcha score";
                    this.showErrorMessage = true;
                    setTimeout(() => {
                      this.showErrorMessage = false;
                    }, 3000);
                  }
                });
            })
            .catch(() => {
              this.errorMessage = "Invalid recaptcha score";
              this.showErrorMessage = true;
              setTimeout(() => {
                this.showErrorMessage = false;
              }, 3000);
            });
        });
      }
    },
    resetForm() {
      this.name = "";
      this.email = "";
      this.projectInfo = "";
      this.reference = "";
      this.invest = "";
      this.floatable = false;
    },
    closePopUps(e) {
      const showList = this.$refs["invest-list"];
      if (showList && !showList.contains(e.target)) {
        this.showList = false;
        if (this.invest == "") {
          this.floatable = false;
        }
      }
    },
    setOption(option) {
      this.invest = option;
      this.showInvestValidationError = this.invest.trim().length == 0;
      this.showList = false;
    },
    toggleList() {
      this.showList = !this.showList;
      this.floatable = this.showList || this.invest.trim().length > 0;
      this.mixpanel.track("tap_contact_invest_input");
    },
  },
};
</script>
