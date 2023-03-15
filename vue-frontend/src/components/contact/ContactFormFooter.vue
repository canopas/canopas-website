<template>
  <div class="tw-mb-9 md:tw-mb-24 tw-font-inter-regular">
    <form>
      <div class="tw-py-5 tw-px-8 lg:tw-px-24">
        <div class="tw-grid tw-grid-cols-1 md:tw-grid-cols-2 tw-gap-4">
          <div class="tw-relative md:tw-mb-5 tw-pt-3 lg:tw-pt-9 tw-text-left">
            <input
              type="text"
              id="username"
              class="tw-block tw-peer tw-my-2 tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-text-white tw-placeholder-white/[.6] floating-input focus:tw-outline-none active:tw-outline-none"
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
              class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-z-[2] tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
              >Your name</label
            >
            <span
              v-if="showNameValidationError"
              class="error canopas-gradient-text"
              >Name is required</span
            >
          </div>
          <div class="tw-relative md:tw-mb-5 tw-pt-3 lg:tw-pt-9 tw-text-left">
            <input
              class="tw-block tw-peer tw-my-2 tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-text-white tw-placeholder-white/[.6] floating-input focus:tw-outline-none active:tw-outline-none"
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
              class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-z-[2] tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
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
            class="tw-relative md:tw-col-span-2 md:tw-mb-5 tw-pt-3 lg:tw-pt-10 tw-text-left"
          >
            <textarea
              class="tw-block tw-peer tw-my-2 tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-text-white tw-placeholder-white/[.6] floating-input focus:tw-outline-none active:tw-outline-none"
              id="project"
              name="project"
              rows="1"
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
              class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-z-[2] tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
              >Tell us about your project</label
            >
            <span
              v-if="showProjectInfoValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
            >
          </div>
          <div
            class="tw-relative md:tw-col-span-2 tw-pt-3 lg:tw-pt-10 tw-text-left"
          >
            <input
              class="tw-block tw-peer tw-my-2 tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none; tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-text-white tw-placeholder-white/[.6] floating-input focus:tw-outline-none active:tw-outline-none"
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
              class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-z-[2] tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
              >How did you find us?</label
            >
            <span
              v-if="showReferenceValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
            >
          </div>
          <div
            class="tw-relative md:tw-col-span-2 md:tw-mt-5 lg:tw-mt-0 md:tw-mb-7 md:tw-pt-3 lg:tw-pt-10 tw-text-left"
          >
            <div ref="invest-list" class="tw-flex">
              <button
                class="tw-flex tw-items-center tw-justify-between tw-w-full tw-pt-3 lg:tw-pb-1 tw-px-0 tw-my-2 md:tw-mt-0 lg:tw-mt-2 tw-mx-0 tw-border-b tw-border-white/[.6] tw-bg-none tw-font-medium !tw-text-white tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-whitespace-nowrap tw-transition tw-duration-150 tw-ease-in-out focus:tw-outline-0 active:tw-outline-0 focus:tw-shadow-none active:tw-shadow-none focus:tw-ring-0 active:tw-ring-0 focus:tw-bg-transparent active:tw-bg-transparent active:tw-text-white"
                type="button"
                id="invest"
                name="invest"
                required
                @click="toggleList"
              >
                <label
                  for="invest"
                  class="tw-absolute tw-top-4 tw-left-0 lg:tw-mt-7 tw-z-[2] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw-duration-300"
                  :class="[
                    floatable
                      ? 'tw--translate-y-4 tw-origin-[0] tw-scale-75 lg:tw-mt-[1rem]'
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
                  class="fab tw-mr-0 tw-mb-2 tw-w-[15px] lg:tw-w-[25px] tw-h-[15px] lg:tw-h-[25px]"
                  :icon="showList ? faCaretUp : faCaretDown"
                />
              </button>
              <ul
                v-if="showList"
                class="tw-absolute tw-mt-12 tw-m-0 tw-w-full tw-min-w-max tw-z-[5] tw-border-none tw-rounded-[10px] tw-shadow-[0px_5px_15px_rgba(0,0,0,0.5)] tw-bg-[#3D3D3D] tw-py-1 tw-list-none tw-text-left tw-text-base tw-transition-all tw-delay-150 tw-duration-300"
              >
                <li
                  v-for="option in investOptions"
                  :key="option"
                  @click="setOption(option)"
                >
                  <span
                    class="tw-block tw-w-full tw-py-2 tw-px-4 tw-font-normal tw-whitespace-nowrap tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.1875rem] md:tw-leading-[1.4375rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.6875rem] hover:tw-bg-[#000]/[.2] hover:tw-text-white"
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
            class="tw-relative tw-inline-flex tw-items-center md:tw-col-span-2 tw-pt-3 lg:tw-pt-0 tw-cursor-pointer tw-text-left"
          >
            <span
              class="tw-mr-6 tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-text-white/[.6]"
              @click="toggleNDA"
              >Send me NDA</span
            >
            <label
              for="nda"
              class="tw-relative tw-inline-flex tw-items-center tw-cursor-pointer"
            >
              <input
                name="nda"
                id="nda"
                type="checkbox"
                v-model="NDA"
                class="tw-sr-only tw-peer"
                aria-label="Send NDA"
              />
              <div
                class="tw-w-9 tw-h-4 tw-rounded-full tw-bg-white/[.6] tw-peer after:tw-content-[''] after:tw-absolute after:tw-top-[-2px] after:tw-left-[-2px] after:tw-h-5 after:tw-w-5 after:tw-border-white/[.6] after:tw-border after:tw-rounded-full after:tw-bg-white after:tw-transition-all peer-focus:tw-outline-none peer-checked:after:tw-border-white/[.6] peer-checked:tw-bg-gradient-to-r peer-checked:tw-from-[#F2709C] peer-checked:tw-to-[#FF9472] peer-checked:after:tw-translate-x-full"
              ></div>
            </label>
          </div>
        </div>

        <div class="tw-flex tw-justify-center tw-mt-9 md:tw-mt-24">
          <img
            v-if="showLoader"
            :src="loaderImage"
            class="tw-w-[64px] tw-h-[64px]"
            alt="loader-image"
          />
          <div class="tw-relative" v-else>
            <div
              class="tw-absolute -tw-top-[2rem] sm:-tw-top-[1.875rem] md:-tw-top-[2.875rem] tw-text-center -tw-right-[4rem] sm:-tw-right-[11rem] md:-tw-right-[15rem] lg:-tw-right-[18rem] xl:-tw-right-[18rem] 2xl:-tw-right-[20.5rem] tw-w-[190%] sm:tw-w-max"
            >
              <span
                v-if="showErrorMessage"
                class="tw-flex tw-text-center canopas-gradient-text sm:tw-mr-[4rem] md:tw-mr-[8rem] lg:tw-mr-[12rem] tw-text-[1.2rem] sm:tw-text-[1.5rem] tw-leading-[1.5rem] sm:tw-leading-[1.8rem]"
                :class="
                  errorMessage == 'Invalid Recaptcha score'
                    ? 'sm:!tw-mr-[7rem] md:!tw-mr-[12rem] lg:!tw-mr-[15rem] xl:!tw-mr-[16rem] 2xl:!tw-mr-[17rem] !tw-text-[1.5rem] '
                    : ''
                "
                >{{ errorMessage }}</span
              >

              <span
                v-if="showSuccessMessage"
                class="canopas-gradient-text tw-text-[1rem] md:tw-text-[1.2rem] lg:tw-text-[1.5rem] 2xl:tw-text-[1.75rem]"
                >Thank you for choosing us to make a difference in your
                business.</span
              >
            </div>
            <button
              id="submit"
              ref="recaptcha"
              class="tw-mt-[1rem] sm:tw-mt-0 tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
              @click.prevent="submitForm()"
            >
              <span
                class="!tw-font-semibold tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0]"
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
      NDA: false,
      showNameValidationError: false,
      showEmailValidationError: false,
      showValidEmailError: false,
      showProjectInfoValidationError: false,
      showReferenceValidationError: false,
      showInvestValidationError: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      showSuccessMessage: false,
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
          project_info: this.projectInfo ? this.projectInfo : "NA",
          reference: this.reference,
          invest: this.invest,
          nda: this.NDA,
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
                    this.showSuccessMessage = true;
                    this.name = "";
                    this.email = "";
                    this.projectInfo = "";
                    this.reference = "";
                    this.invest = "";
                    this.NDA = false;
                    this.floatable = false;
                    this.showLoader = false;
                    setTimeout(() => {
                      this.showSuccessMessage = false;
                    }, 3000);
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
    toggleNDA() {
      this.NDA = !this.NDA;
      this.mixpanel.track("tap_footer_NDA_input");
    },
  },
};
</script>
