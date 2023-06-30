<template>
  <div class="tw-mb-9 tw-font-inter-regular md:tw-mb-24">
    <form>
      <div class="tw-px-8 tw-py-5 lg:tw-px-24">
        <div class="tw-grid tw-grid-cols-1 tw-gap-4 md:tw-grid-cols-2">
          <div class="tw-relative tw-pt-3 tw-text-left md:tw-mb-5 lg:tw-pt-9">
            <input
              type="text"
              id="username"
              class="floating-input tw-peer tw-mx-0 tw-my-2 tw-block tw-w-full tw-appearance-none tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-text-lg tw-text-white tw-placeholder-white/[.6] tw-transition tw-ease-in-out focus:tw-outline-none active:tw-outline-none md:tw-text-xl lg:tw-text-2xl"
              name="username"
              required
              autocomplete="given-username"
              v-model="name"
              placeholder=" "
              @input="
                showNameValidationError =
                  $event.target.value.trim().length === 0
              "
              @click.native="mixpanel.track('tap_footer_username_input')"
            />
            <label
              for="username"
              class="tw-absolute tw-left-0 tw-top-4 tw-z-[2] tw-mb-5 tw-origin-[0] tw--translate-y-4 tw-scale-75 tw-transform tw-text-[1rem] tw-leading-[1.1875rem] tw-text-white/[.6] tw-duration-300 peer-placeholder-shown:tw-translate-y-0 peer-placeholder-shown:tw-scale-100 peer-focus:tw--translate-y-4 peer-focus:tw-scale-75 peer-focus:tw-text-white md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem]"
              >Your name</label
            >
            <span
              v-if="showNameValidationError"
              class="error canopas-gradient-text"
              >Name is required</span
            >
          </div>
          <div class="tw-relative tw-pt-3 tw-text-left md:tw-mb-5 lg:tw-pt-9">
            <input
              class="floating-input tw-peer tw-mx-0 tw-my-2 tw-block tw-w-full tw-appearance-none tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-text-lg tw-text-white tw-placeholder-white/[.6] tw-transition tw-ease-in-out focus:tw-outline-none active:tw-outline-none md:tw-text-xl lg:tw-text-2xl"
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
              @click.native="mixpanel.track('tap_footer_email_input')"
            />
            <label
              for="email"
              class="tw-absolute tw-left-0 tw-top-4 tw-z-[2] tw-mb-5 tw-origin-[0] tw--translate-y-4 tw-scale-75 tw-transform tw-text-[1rem] tw-leading-[1.1875rem] tw-text-white/[.6] tw-duration-300 peer-placeholder-shown:tw-translate-y-0 peer-placeholder-shown:tw-scale-100 peer-focus:tw--translate-y-4 peer-focus:tw-scale-75 peer-focus:tw-text-white md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem]"
              >Your email</label
            >
            <span
              v-if="showEmailValidationError"
              class="error canopas-gradient-text"
              >Email is required</span
            >
            <span
              v-if="email.trim().length != 0 && showValidEmailError"
              class="error canopas-gradient-text"
              >Please enter valid email address</span
            >
          </div>
          <div
            class="tw-relative tw-pt-3 tw-text-left md:tw-col-span-2 md:tw-mb-5 lg:tw-pt-10"
          >
            <textarea
              class="floating-input tw-peer tw-mx-0 tw-my-2 tw-block tw-min-h-[50px] tw-w-full tw-appearance-none tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-text-lg tw-text-white tw-placeholder-white/[.6] tw-transition tw-ease-in-out focus:tw-outline-none active:tw-outline-none md:tw-min-h-[90px] md:tw-text-xl lg:tw-text-2xl"
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
              @click.native="mixpanel.track('tap_footer_project_info_input')"
            ></textarea>
            <label
              for="project"
              class="tw-absolute tw-left-0 tw-top-4 tw-z-[2] tw-mb-5 tw-origin-[0] tw--translate-y-4 tw-scale-75 tw-transform tw-text-[1rem] tw-leading-[1.1875rem] tw-text-white/[.6] tw-duration-300 peer-placeholder-shown:tw-translate-y-0 peer-placeholder-shown:tw-scale-100 peer-focus:tw--translate-y-4 peer-focus:tw-scale-75 peer-focus:tw-text-white md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem]"
              >Tell us about your project</label
            >
            <span
              v-if="showProjectInfoValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
            >
          </div>
          <div
            class="tw-relative tw-pt-3 tw-text-left md:tw-col-span-2 lg:tw-pt-10"
          >
            <input
              class="tw-appearance-none; floating-input tw-peer tw-mx-0 tw-my-2 tw-block tw-w-full tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-text-lg tw-text-white tw-placeholder-white/[.6] tw-transition tw-ease-in-out focus:tw-outline-none active:tw-outline-none md:tw-text-xl lg:tw-text-2xl"
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
              @click.native="mixpanel.track('tap_footer_reference_input')"
            />
            <label
              for="reference"
              class="tw-absolute tw-left-0 tw-top-4 tw-z-[2] tw-mb-5 tw-origin-[0] tw--translate-y-4 tw-scale-75 tw-transform tw-text-[1rem] tw-leading-[1.1875rem] tw-text-white/[.6] tw-duration-300 peer-placeholder-shown:tw-translate-y-0 peer-placeholder-shown:tw-scale-100 peer-focus:tw--translate-y-4 peer-focus:tw-scale-75 peer-focus:tw-text-white md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem]"
              >How did you find us?</label
            >
            <span
              v-if="showReferenceValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
            >
          </div>
          <div
            class="tw-relative tw-text-left md:tw-col-span-2 md:tw-mb-7 md:tw-mt-5 md:tw-pt-3 lg:tw-mt-0 lg:tw-pt-10"
          >
            <div ref="invest-list" class="tw-flex">
              <button
                class="tw-mx-0 tw-my-2 tw-flex tw-w-full tw-items-center tw-justify-between tw-whitespace-nowrap tw-border-b tw-border-white/[.6] tw-bg-none tw-px-0 tw-pt-3 tw-text-lg tw-font-medium !tw-text-white tw-transition tw-duration-150 tw-ease-in-out focus:tw-bg-transparent focus:tw-shadow-none focus:tw-outline-0 focus:tw-ring-0 active:tw-bg-transparent active:tw-text-white active:tw-shadow-none active:tw-outline-0 active:tw-ring-0 md:tw-mt-0 md:tw-text-xl lg:tw-mt-2 lg:tw-pb-1 lg:tw-text-2xl"
                type="button"
                id="invest"
                name="invest"
                required
                @click="toggleList"
              >
                <label
                  for="invest"
                  class="tw-absolute tw-left-0 tw-top-4 tw-z-[2] tw-transform tw-text-[1rem] tw-leading-[1.1875rem] tw-duration-300 md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-mt-7 lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem]"
                  :class="[
                    floatable
                      ? 'tw-origin-[0] tw--translate-y-4 tw-scale-75 lg:tw-mt-[1rem]'
                      : 'tw-translate-y-0 tw-scale-100 ',
                    floatable && showList
                      ? 'tw-text-white'
                      : 'tw-text-white/[.6] ',
                  ]"
                  @click="floatable = !floatable"
                  >I'll invest</label
                >
                <span class="tw-cursor-text">{{ invest }}</span>
                <font-awesome-icon
                  class="fab tw-mb-2 tw-mr-0 tw-h-[15px] tw-w-[15px] lg:tw-h-[25px] lg:tw-w-[25px]"
                  :icon="showList ? faCaretUp : faCaretDown"
                />
              </button>
              <ul
                v-if="showList"
                class="tw-absolute tw-z-[5] tw-m-0 tw-mt-12 tw-w-full tw-min-w-max tw-list-none tw-rounded-[10px] tw-border-none tw-bg-[#3D3D3D] tw-py-1 tw-text-left tw-text-base tw-shadow-[0px_5px_15px_rgba(0,0,0,0.5)] tw-transition-all tw-delay-150 tw-duration-300"
              >
                <li
                  v-for="option in investOptions"
                  :key="option"
                  @click="setOption(option)"
                >
                  <span
                    class="tw-block tw-w-full tw-whitespace-nowrap tw-px-4 tw-py-2 tw-text-[1rem] tw-font-normal tw-leading-[1.1875rem] tw-text-white/[.6] hover:tw-bg-[#000]/[.2] hover:tw-text-white md:tw-text-[1.1875rem] md:tw-leading-[1.4375rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.6875rem]"
                    >{{ option }}</span
                  >
                </li>
              </ul>
            </div>
            <span
              v-if="showInvestValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
            >
          </div>
          <div
            class="tw-w-full tw-text-left tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-text-white/[.6] xl:tw-text-[1.5rem] xl:tw-leading-[1.875rem]"
          >
            We sign NDA for all of our projects.
          </div>
        </div>

        <div class="tw-mt-9 tw-flex tw-justify-center md:tw-mt-24">
          <img
            v-if="showLoader"
            :src="loaderImage"
            class="tw-h-[64px] tw-w-[64px]"
            alt="loader-image"
          />
          <div class="tw-relative" v-else>
            <div
              class="tw-absolute -tw-right-[4rem] -tw-top-[2rem] tw-w-[190%] tw-text-center sm:-tw-right-[11rem] sm:-tw-top-[1.875rem] sm:tw-w-max md:-tw-right-[15rem] md:-tw-top-[2.875rem] lg:-tw-right-[18rem] xl:-tw-right-[18rem] 2xl:-tw-right-[20.5rem]"
            >
              <span
                v-if="showErrorMessage"
                class="canopas-gradient-text tw-flex tw-text-center tw-text-[1.2rem] tw-leading-[1.5rem] sm:tw-mr-[4rem] sm:tw-text-[1.5rem] sm:tw-leading-[1.8rem] md:tw-mr-[8rem] lg:tw-mr-[12rem]"
                :class="
                  errorMessage == 'Invalid Recaptcha score'
                    ? '!tw-text-[1.5rem] sm:!tw-mr-[7rem] md:!tw-mr-[12rem] lg:!tw-mr-[15rem] xl:!tw-mr-[16rem] 2xl:!tw-mr-[17rem] '
                    : ''
                "
                >{{ errorMessage }}</span
              >
            </div>
            <button
              id="submit"
              ref="recaptcha"
              class="gradient-btn consultation-btn tw-mt-[1rem] tw-rounded-full tw-p-3 tw-text-center sm:tw-mt-0"
              @click.prevent="submitForm()"
            >
              <span
                class="tw-mr-2.5 tw-font-inter-medium tw-text-[1rem] !tw-font-semibold tw-font-normal tw-leading-[1.1875rem] !tw-tracking-[0] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
                >Send Message</span
              >
            </button>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import axios from "axios";
import config from "@/config.js";
import loaderImage from "@/assets/images/theme/small-loader.svg";
import { faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
export default {
  data() {
    return {
      loaderImage,
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
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      showErrorMessage: false,
    };
  },
  components: {
    FontAwesomeIcon,
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
      this.mixpanel.track("tap_footer_form_submission");
      if (!this.validateForm()) {
        this.showLoader = true;
        let formData = {
          name: this.name,
          email: this.email,
          project_info: this.projectInfo
            ? this.projectInfo.replace(/\./g, ".\n")
            : "NA",
          reference: this.reference,
          invest: this.invest,
          contact_type: "",
        };
        var head = document.getElementsByTagName("head")[0];
        var script = document.createElement("script");
        script.type = "text/javascript";
        script.onload = () => {
          // verify recaptcha
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
                    this.name = "";
                    this.email = "";
                    this.projectInfo = "";
                    this.reference = "";
                    this.invest = "";
                    this.floatable = false;
                    this.showLoader = false;
                    this.$router.push({
                      path: "/thank-you",
                    });
                  })
                  .catch((err) => {
                    if (err.response.status == 401) {
                      this.errorMessage = "Invalid recaptcha score";
                      this.showErrorMessage = true;
                      this.showLoader = false;
                      setTimeout(() => {
                        this.showErrorMessage = false;
                      }, 3000);
                    }
                  });
              })
              .catch(() => {
                this.errorMessage = "Invalid recaptcha score";
                this.showErrorMessage = true;
                this.showLoader = false;
                setTimeout(() => {
                  this.showErrorMessage = false;
                }, 3000);
              });
          });
        };
      }
      script.src =
        "https://www.google.com/recaptcha/enterprise.js?render=" +
        import.meta.env.VITE_RECAPTCHA_SITE_KEY;
      head.appendChild(script);
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
      this.mixpanel.track("tap_footer_invest_input");
    },
  },
};
</script>
