<template>
  <div class="tw-container tw-my-[50px] lg:tw-mt-0 md:tw-w-full">
    <form>
      <div
        class="tw-rounded-2xl tw-border tw-border-slate-200 tw-bg-white tw-shadow-[0px_0px_45px_rgba(0,0,0,0.1)]"
      >
        <div class="tw-py-5 tw-px-8 lg:tw-px-12">
          <div>
            <div class="tw-relative tw-mb-5 md:tw-mb-5 tw-pt-8 tw-text-left">
              <input
                type="text"
                id="username"
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
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
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
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
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
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
                class="tw-block tw-peer tw-mx-0 tw-w-full tw-rounded-none tw-border-b tw-border-slate-400 tw-bg-transparent tw-px-0 tw-transition tw-ease-in-out tw-appearance-none; tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-text-black-core tw-placeholder-black-core/[.6] focus:tw-outline-none active:tw-outline-none"
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
            <div class="tw-relative tw-pt-7">
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
                        ? 'tw-text-white'
                        : 'tw-text-white/[.6]',
                    ]"
                    @click="floatable = !floatable"
                    >I'll invest</label
                  >
                  <span
                    class="md:tw-mt-[0.5rem] lg:tw-mt-[1rem] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem]"
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
                      class="tw-block tw-w-full tw-py-2 tw-px-4 tw-whitespace-nowrap tw-text-black-core/[.6] hover:tw-bg-[#000]/[.2] hover:tw-text-[#000] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem]"
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
              class="tw-relative tw-items-center tw-justify-between tw-mb-3 tw-mx-0 tw-w-full tw-border-b tw-border-slate-400 tw-bg-none tw-pt-7 tw-px-0 tw-text-left tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem] tw-whitespace-nowrap tw-transition tw-duration-150 tw-ease-in-out focus:tw-outline-0 active:tw-outline-0 focus:tw-shadow-none active:tw-shadow-none focus:tw-ring-0 active:tw-ring-0 focus:tw-bg-transparent active:tw-bg-transparent active:tw-text-black-core/[0.6]"
            >
              <span
                class="tw-mr-6 tw-text-black-core/[0.6] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[2.21rem] md:tw-leading-[3.0625rem]"
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
                  aria-label="send NDA"
                />
                <div
                  class="tw-w-9 tw-h-4 tw-rounded-full tw-bg-slate-400 tw-peer after:tw-content-[''] after:tw-absolute after:tw-top-[-2px] after:tw-left-[-2px] after:tw-h-5 after:tw-w-5 after:tw-border-slate-400 after:tw-border after:tw-rounded-full after:tw-bg-white after:tw-transition-all peer-focus:tw-outline-none peer-checked:after:tw-border-slate-400 peer-checked:tw-bg-gradient-to-r peer-checked:tw-from-[#F2709C] peer-checked:tw-to-[#FF9472] peer-checked:after:tw-translate-x-full"
                ></div>
              </label>
            </div>
          </div>

          <div class="tw-pt-6">
            <div
              class="tw-text-black-core/[0.6] tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.21rem] md:tw-leading-[2.0625rem]"
            >
              What's your preferred mode of communication?
            </div>
            <div
              class="tw-flex tw-flex-row tw-items-center tw-justify-center md:tw-justify-start tw-gap-4 tw-pt-4"
            >
              <label class="contactLabel1"
                ><input
                  class="tw-hidden"
                  type="radio"
                  name="contact_type"
                  value="2"
                  v-model.number="contactType"
                />
                <div
                  :class="
                    contactType == 2
                      ? 'tw-text-white tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-text-white tw-shadow-none'
                      : 'gradient-border-btn tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg]'
                  "
                  class="tw-flex tw-items-center tw-mt-1.5 tw-mb-6 tw-cursor-pointer tw-rounded-full tw-px-[2.5rem] lg:tw-px-[1.8rem] tw-py-2 md:tw-py-3 active:tw-scale-[0.98] tw-font-inter-semibold tw-text-[1rem] md:tw-text-[1.1875rem] tw-leading-[1.21rem] md:tw-leading-[1.386875rem] tw-text-center"
                >
                  <font-awesome-icon
                    :class="
                      contactType == 2 ? 'tw-text-white' : 'tw-text-[#f2709c]'
                    "
                    class="fa md:tw-mr-2 tw-w-[22.99px] md:tw-w-[28.18px] tw-h-[15.83px] md:tw-h-[24.18px]"
                    :icon="phone"
                    aria-hidden="true"
                  />
                  <span
                    class="v2-canopas-gradient-text"
                    :class="
                      contactType == 2
                        ? 'tw-text-white'
                        : 'v2-canopas-gradient-text'
                    "
                    >Call</span
                  >
                </div></label
              >

              <label class="contactLabel2"
                ><input
                  class="tw-hidden"
                  type="radio"
                  name="contact_type"
                  value="1"
                  v-model.number="contactType"
                />
                <div
                  :class="
                    contactType == 1
                      ? 'tw-text-white tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] '
                      : 'gradient-border-btn tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg]'
                  "
                  class="tw-flex tw-cursor-pointer tw-items-center tw-mt-1.5 tw-mb-6 tw-rounded-full tw-px-[2.5rem] lg:tw-px-[1.8rem] tw-py-2 md:tw-py-3 active:tw-scale-[0.98] tw-text-center tw-font-inter-semibold tw-text-[1rem] md:tw-text-[1.1875rem] tw-leading-[1.21rem] md:tw-leading-[1.386875rem]"
                >
                  <font-awesome-icon
                    :class="contactType == 1 ? '' : 'tw-text-[#f2709c]'"
                    :icon="mail"
                    class="fa md:tw-mr-2 tw-w-[22.99px] md:tw-w-[28.18px] tw-h-[15.83px] md:tw-h-[24.18px] tw-text-white"
                    aria-hidden="true"
                  />
                  <span
                    :class="contactType == 1 ? '' : 'v2-canopas-gradient-text'"
                    >Mail</span
                  >
                </div></label
              >
            </div>
          </div>
          <div class="tw-flex tw-justify-center tw-py-5 lg:tw-py-8">
            <img
              v-if="showLoader"
              :src="loaderImage"
              class="tw-w-[64px] tw-h-[64px]"
              alt="loader-image"
            />
            <div v-else>
              <div
                class="tw-relative tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.436875rem] md:tw-leading-[1.386875rem] tw-text-center"
              >
                <div
                  class="tw-absolute -tw-top-[2.5rem] sm:-tw-top-[1.875rem] tw-text-center tw-right-[1rem] sm:-tw-right-[5rem] md:-tw-right-[8rem] lg:-tw-right-[4rem] xl:-tw-right-[7rem] 2xl:-tw-right-[7.5rem] sm:tw-w-max"
                >
                  <span
                    v-if="showErrorMessage"
                    class="tw-flex tw-text-center -tw-mr-[0.875rem] sm:tw-mr-[2rem] md:tw-mr-[6.563rem] lg:tw-mr-[3.438rem] xl:tw-mr-[6rem] 2xl:tw-mr-[7rem] tw-text-red-600 tw-text-[1.3rem] sm:tw-text-[1.3rem]"
                    :class="
                      errorMessage == 'Invalid Recaptcha score'
                        ? '!-tw-mr-[1rem] sm:!tw-mr-[5rem] md:!tw-mr-[10rem] lg:!tw-mr-[6.5rem] xl:!tw-mr-[10rem] '
                        : ''
                    "
                    >{{ errorMessage }}</span
                  >

                  <span
                    v-if="contactType == 1 && showSuccessMessage"
                    class="canopas-gradient-text tw-text-[0.8rem] md:tw-text-[1.1rem] lg:tw-text-[0.9rem] xl:tw-text-[1.1rem] tw-leading-[1rem]"
                    >Thank you for choosing us to make a difference in your
                    business.</span
                  >
                </div>
                <button
                  id="submit"
                  v-if="contactType == 1"
                  ref="recaptcha"
                  class="tw-mt-[1rem] sm:tw-mt-0 tw-rounded-full tw-py-2.5 md:tw-py-5 tw-px-20 md:tw-px-28 tw-text-center gradient-btn consultation-btn"
                  @click.prevent="submitForm()"
                >
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6"
                    aria-hidden="true"
                  />
                  <span>Submit</span>
                </button>
                <button
                  v-if="contactType == 2"
                  class="tw-flex tw-items-center tw-m-0 tw-w-max tw-rounded-full tw-py-2.5 md:tw-py-5 sm:tw-px-12 tw-text-center gradient-btn consultation-btn"
                  @click.prevent="submitForm()"
                >
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6"
                    :icon="calendarIcon"
                    aria-hidden="true"
                  />
                  <span>Schedule Meeting</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </form>
    <!-- Show Calendly Iframe -->
    <div v-if="openCalendlyIframeModal">
      <transition name="modal">
        <div
          class="modal-mask tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-table mask tw-z-[1] tw-bg-[#00000080] tw-z-[5]"
        >
          <div
            class="tw-mx-auto tw-left-auto sm:tw-my-7 sm:tw-mx-auto sm:tw-max-w-lg tw-h-full login-modal modal-xl tw-flex tw-items-center"
            role="document"
          >
            <div
              class="tw-relative tw-flex tw-flex-col tw-w-full tw-bg-white tw-bg-clip-padding tw-rounded-md tw-outline-0 tw-border-1 tw-border-gray tw-border-solid tw-p-2.5 tw-rounded-3xl"
            >
              <div class="tw-relative tw-p-4 tw-flex-auto">
                <CalendlyIframe />

                <button
                  type="button"
                  class="close modal-close-btn tw-border-none tw-text-pink-300 tw-text-4xl tw-font-light tw-bg-transparent tw-absolute tw-right-2.5 tw-top-0 focus:tw-outline-none"
                  data-dismiss="modal"
                  aria-label="Close"
                >
                  <span aria-hidden="true" @click="closeCalendlyIframeModal"
                    >&times;</span
                  >
                </button>
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
import CalendlyIframe from "./CalendlyIframe.vue";
import loaderImage from "@/assets/images/theme/small-loader.svg";
import { faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faPhoneVolume,
  faEnvelope,
  faCalendarAlt,
} from "@fortawesome/free-solid-svg-icons";

const CONTACT_BY_CHAT_OR_MAIL = 1;
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
      NDA: false,
      showNameValidationError: false,
      showEmailValidationError: false,
      showValidEmailError: false,
      showProjectInfoValidationError: false,
      showReferenceValidationError: false,
      showInvestValidationError: false,
      openCalendlyIframeModal: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      showSuccessMessage: false,
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
          project_info: this.projectInfo ? this.projectInfo : "NA",
          reference: this.reference,
          invest: this.invest,
          nda: this.NDA,
          contact_type:
            this.contactType == CONTACT_BY_CHAT_OR_MAIL
              ? "Chat or Email"
              : "Call",
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
                  this.resetForm();
                  if (this.contactType == CONTACT_BY_CHAT_OR_MAIL) {
                    this.showSuccessMessage = true;
                    setTimeout(() => {
                      this.showSuccessMessage = false;
                    }, 3000);
                  } else {
                    setTimeout(() => {
                      this.mixpanel.track("tap_schedule_meeting_click");
                      this.openCalendlyIframe();
                      if (this.openCalendlyIframeModal) {
                        this.showLoader = false;
                      }
                    }, 1000);
                  }
                  this.contactType = CONTACT_BY_CHAT_OR_MAIL;
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
      this.NDA = false;
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
    toggleNDA() {
      this.NDA = !this.NDA;
      this.mixpanel.track("tap_contact_NDA_input");
    },
    openCalendlyIframe() {
      this.openCalendlyIframeModal = true;
      this.showLoader = false;
    },
    closeCalendlyIframeModal() {
      this.openCalendlyIframeModal = false;
      this.mixpanel.track("close_calendly_dialog_error");
    },
  },
};
</script>
