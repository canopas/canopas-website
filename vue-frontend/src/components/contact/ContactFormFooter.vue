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
              :disabled="disableInput"
              placeholder=" "
              @input="showNameValidationError = name === ''"
              @click.native="$gtag.event('tap_footer_username_input')"
            />
            <label
              for="username"
              class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-z-[2] tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
              >Your name</label
            >
            <span
              v-if="showNameValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
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
              :disabled="disableInput"
              placeholder=" "
              @input="showEmailValidationError = email === ''"
              @click.native="$gtag.event('tap_footer_email_input')"
            />
            <label
              for="email"
              class="tw-absolute tw-top-4 tw-left-0 tw-mb-5 tw-z-[2] tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4"
              >Your email</label
            >
            <span
              v-if="showEmailValidationError"
              class="error canopas-gradient-text"
              >This field is required</span
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
              :disabled="disableInput"
              placeholder=" "
              @input="showProjectInfoValidationError = projectInfo === ''"
              @click.native="$gtag.event('tap_footer_project_info_input')"
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
            class="tw-relative md:tw-col-span-2 md:tw-mb-5 tw-pt-3 lg:tw-pt-10 tw-text-left"
          >
            <input
              class="tw-block tw-peer tw-my-2 tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-white/[.6] tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none; tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-text-white tw-placeholder-white/[.6] floating-input focus:tw-outline-none active:tw-outline-none"
              type="text"
              name="reference"
              id="reference"
              required
              autocomplete="given-reference"
              v-model="reference"
              :disabled="disableInput"
              placeholder=" "
              @input="showReferenceValidationError = reference === ''"
              @click.native="$gtag.event('tap_footer_reference_input')"
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
          <div class="tw-relative md:tw-col-span-2 md:tw-mb-5 tw-text-left">
            <div ref="invest-list" class="tw-flex">
              <button
                class="tw-flex tw-items-center tw-justify-between tw-w-full tw-pb-0 tw-px-0 tw-my-[2rem] tw-mx-0 tw-border-b tw-border-white/[.6] tw-bg-none tw-font-medium tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-whitespace-nowrap tw-transition tw-duration-150 tw-ease-in-out focus:tw-outline-0 active:tw-outline-0 focus:tw-shadow-none active:tw-shadow-none focus:tw-ring-0 active:tw-ring-0 focus:tw-bg-transparent active:tw-bg-transparent active:tw-text-white"
                type="button"
                id="invest"
                name="invest"
                required
                :disabled="disableInput"
                @click="
                  toggleList();
                  makefloatActive();
                "
                @blur="
                  makeFloatInactive();
                  deselectOption();
                "
              >
                {{ invest }}
                <label
                  class="tw-pb-8"
                  :class="
                    floatable
                      ? 'tw-absolute tw-mt-[5rem] tw-mb-[7rem] tw-left-0 tw-z-[2] tw-text-white tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4'
                      : '' || investSelected
                      ? 'tw-absolute tw-bottom-10 tw-left-0 tw-z-[2] tw-text-white tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-transform tw--translate-y-4 tw-origin-[0] tw-scale-75 tw-duration-300 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4'
                      : ''
                  "
                  >I'll invest</label
                >
                <font-awesome-icon
                  class="tw-absolute fab tw-right-0 tw-w-[15px] tw-h-[15px]"
                  :icon="showList ? faCaretUp : faCaretDown"
                  :class="showList ? 'tw-mb-[2rem]' : ''"
                />
              </button>
              <ul
                v-if="showList"
                class="tw-absolute tw-float-left tw-mt-12 tw-m-0 tw-w-full tw-min-w-max tw-z-[5] tw-border-none tw-rounded-[10px] tw-shadow-[0px_5px_15px_rgba(0,0,0,0.5)] tw-bg-[#3D3D3D] tw-py-1 tw-list-none tw-text-left tw-text-base tw-transition-all tw-delay-150 tw-duration-300"
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
            class="tw-relative tw-inline-flex tw-items-center md:tw-col-span-2 tw-pt-3 lg:tw-pt-10 tw-cursor-pointer tw-text-left"
          >
            <span
              class="tw-mr-6 tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-text-white/[.6]"
              @click="toggleNDA"
              >Send me NDA</span
            >
            <label
              class="tw-relative tw-inline-flex tw-items-center tw-cursor-pointer"
            >
              <input type="checkbox" v-model="NDA" class="tw-sr-only tw-peer" />
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
          />
          <div v-else>
            <button
              id="submit"
              ref="recaptcha"
              class="tw-flex tw-items-center tw-w-max tw-m-0 tw-rounded-full tw-py-3 tw-px-3 tw-text-center gradient-btn consultation-btn"
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
    <!-- Show Thank you message -->
    <div v-if="showSuccessMessagePopup">
      <transition name="modal">
        <div
          class="modal-mask tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-z-[5] tw-bg-[#00000080]"
        >
          <div
            class="tw-flex tw-items-center sm:tw-mx-auto sm:tw-max-w-lg tw-h-full"
          >
            <div class="tw-w-full tw-rounded-md tw-bg-white">
              <div
                class="tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300 tw-p-4"
              >
                <div
                  class="canopas-gradient-text tw-text-[1.75rem] tw-leading-[2.25rem] tw-text-center"
                >
                  Thank you for choosing us to make a difference in your
                  business.
                </div>
              </div>
              <div class="tw-relative tw-flex-auto tw-p-4">
                <div class="normal-2-text">
                  If you prefer to chat or email, sit back and relax our team
                  will get back to you within 24 hours.
                </div>
                <div class="close-btn-div">
                  <button
                    class="gradient-btn tw-w-40 tw-mb-2 tw-px-0"
                    @click.prevent="showSuccessMessagePopup = false"
                  >
                    <span>Close</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
    <!-- Show error message -->
    <div v-if="showErrorMessagePopup">
      <transition name="modal">
        <div
          class="tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-z-[5] tw-bg-[#00000080]"
        >
          <div
            class="tw-flex tw-items-center sm:tw-mx-auto sm:tw-max-w-lg tw-h-full"
          >
            <div
              class="tw-w-full tw-rounded-md tw-bg-white tw-bg-clip-padding tw-p-2"
            >
              <div
                class="tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300 tw-p-3"
              >
                <div
                  class="canopas-gradient-text tw-text-[1.75rem] tw-leading-[2.25rem] tw-font-semibold"
                >
                  Error
                </div>
              </div>
              <div class="tw-relative tw-flex-auto tw-p-4">
                <div class="normal-text text-center">
                  {{ errorMessage }}
                </div>
                <div class="close-btn-div">
                  <button
                    class="gradient-btn tw-float-right tw-w-40 tw-px-0"
                    @click.prevent="showErrorMessagePopup = false"
                  >
                    <span>Close</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
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
      disableInput: false,
      showNameValidationError: false,
      showEmailValidationError: false,
      showProjectInfoValidationError: false,
      showReferenceValidationError: false,
      showInvestValidationError: false,
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      floatable: false,
      investSelected: false,
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {
    document.addEventListener("click", this.closePopUps);
  },
  unmounted() {
    document.removeEventListener("click", this.closePopUps);
  },
  methods: {
    makefloatActive() {
      this.floatable = true;
    },
    makeFloatInactive() {
      this.floatable = false;
    },
    validateForm() {
      this.showNameValidationError = this.name === "";
      this.showEmailValidationError = this.email === "";
      this.showProjectInfoValidationError = this.projectInfo === "";
      this.showReferenceValidationError = this.reference === "";
      this.showInvestValidationError = this.invest === "";
      return (
        this.showNameValidationError ||
        this.showEmailValidationError ||
        this.showProjectInfoValidationError ||
        this.showReferenceValidationError ||
        this.showInvestValidationError
      );
    },
    submitForm() {
      this.$gtag.event("tap_home_form_submission");

      if (!this.validateForm()) {
        this.disableInput = true;
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
        script.onload = function () {
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
                    setTimeout(() => {
                      this.showSuccessMessagePopup = true;
                      this.showLoader = false;
                    }, 1000);
                  })
                  .catch((err) => {
                    if (err.response.status == 401) {
                      this.errorMessage = "Invalid recaptcha score";
                    }
                    this.showErrorMessagePopup = true;
                    this.showLoader = false;
                  });
              })
              .catch(() => {
                this.errorMessage = "Invalid recaptcha score";
                this.showErrorMessagePopup = true;
                this.showLoader = false;
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
      }
    },
    setOption(option) {
      this.showInvestValidationError = invest === "";
      this.invest = option;
      this.showList = false;
      this.investSelected = true;
    },
    deselectOption() {
      this.investSelected = false;
    },
    toggleList() {
      this.$gtag.event("tap_footer_invest_input");
      this.showList = !this.showList;
    },
    toggleNDA() {
      this.$gtag.event("tap_footer_NDA_input");
      this.NDA = !this.NDA;
    },
  },
};
</script>
