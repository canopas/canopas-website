<template>
  <div class="tw-container -tw-mt-16 tw-mb-24">
    <form>
      <div
        class="tw-rounded-3xl tw-border tw-border-slate-200 tw-bg-white tw-shadow-3xl"
      >
        <div class="tw-py-5 tw-px-8 lg:tw-px-24">
          <div class="row tw-grid tw-grid-cols-1 md:tw-grid-cols-2 tw-gap-4">
            <div class="tw-pt-4">
              <input
                class="tw-w-full tw-py-3 tw-px-0 tw-my-2 tw-mx-0 tw-box-border tw-border-b tw-border-slate-400 tw-rounded-none focus:tw-box-border focus:tw-border-b focus:tw-border-slate-400 focus:tw-outline-none placeholder:tw-text-slate-400 v2-normal-text"
                type="text"
                name="username"
                required
                autocomplete="given-username"
                v-model="name"
                :disabled="disableInput"
                placeholder="Your name"
                @input="showNameValidationError = name.trim().length === 0"
                @click.native="mixpanel.track('tap_name_input')"
              />
              <span v-if="showNameValidationError" class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div class="tw-pt-4">
              <input
                class="tw-w-full tw-py-3 tw-px-0 tw-my-2 tw-mx-0 tw-box-border tw-border-b tw-border-slate-400 tw-rounded-none focus:tw-box-border focus:tw-border-b focus:tw-border-slate-400 focus:tw-outline-none placeholder:tw-text-slate-400 v2-normal-text"
                type="text"
                name="email"
                required
                autocomplete="given-email"
                v-model="email"
                :disabled="disableInput"
                placeholder="Your email"
                @input="
                  showEmailValidationError = email.trim().length === 0;
                  showValidEmailError = isValidEmail();
                "
                @click.native="mixpanel.track('tap_email_input')"
              />
              <span
                v-if="showEmailValidationError"
                class="error tw-text-red-600"
                >This field is required</span
              >
              <span v-if="showValidEmailError" class="error tw-text-red-600"
                >Please enter valid email address</span
              >
            </div>
            <div class="md:tw-col-span-2 tw-pt-1 lg:tw-pt-8">
              <textarea
                class="tw-w-full tw-py-3 tw-px-0 tw-my-2 tw-mx-0 tw-box-border tw-border-b tw-border-slate-400 tw-rounded-none focus:tw-box-border focus:tw-border-b focus:tw-border-slate-400 focus:tw-outline-none placeholder:tw-text-slate-400 v2-normal-text"
                name="project"
                rows="1"
                required
                autocomplete="given-project-info"
                v-model="projectInfo"
                :disabled="disableInput"
                placeholder="Tell us about your project"
                @input="
                  showProjectInfoValidationError =
                    projectInfo.trim().length === 0
                "
                @click.native="mixpanel.track('tap_tell_us_about_project')"
              ></textarea>
              <span
                v-if="showProjectInfoValidationError"
                class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div class="md:tw-col-span-2 tw-pt-1 lg:tw-pt-8">
              <input
                class="tw-w-full tw-py-3 tw-px-0 tw-my-2 tw-mx-0 tw-box-border tw-border-b tw-border-slate-400 tw-rounded-none focus:tw-box-border focus:tw-border-b focus:tw-border-slate-400 focus:tw-outline-none placeholder:tw-text-slate-400 v2-normal-text"
                type="text"
                name="reference"
                required
                autocomplete="given-reference"
                v-model="reference"
                :disabled="disableInput"
                placeholder="How did you find us?"
                @input="
                  showReferenceValidationError = reference.trim().length === 0
                "
                @click.native="mixpanel.track('tap_how_did_you_find_us')"
              />
              <span
                v-if="showReferenceValidationError"
                class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
          </div>

          <div class="tw-py-10 lg:tw-pt-16">
            <div class="tw-text-slate-400 v2-normal-text">
              What's your preferred mode of communication?
            </div>
            <div class="tw-flex tw-flex-col tw-w-fit sm:tw-flex-row tw-pt-6">
              <label class="contactLabel1"
                ><input
                  class="tw-hidden"
                  type="radio"
                  name="contact_type"
                  value="2"
                  v-model.number="contactType"
                  :disabled="disableInput"
                  @click.native="mixpanel.track('tap_preferred_mode_call')"
                />
                <div
                  :class="
                    contactType == 2
                      ? 'tw-bg-black-900 tw-text-white'
                      : 'hover:tw-text-white hover:tw-bg-black-900 tw-bg-white tw-text-black-900'
                  "
                  class="tw-cursor-pointer v2-normal-3-text contact-button tw-flex tw-items-center tw-border tw-border-black-900 tw-rounded-lg tw-px-6 tw-py-2 tw-mb-6 sm:tw-mr-9 sm:tw-mb-0 active:tw-scale-[0.98]"
                >
                  <font-awesome-icon
                    class="-tw-rotate-45 tw-w-6 tw-h-6"
                    :icon="phone"
                    aria-hidden="true"
                  />
                  <span class="tw-ml-5">Call</span>
                </div></label
              >

              <label class="contactLabel2"
                ><input
                  class="tw-hidden"
                  type="radio"
                  name="contact_type"
                  value="1"
                  v-model.number="contactType"
                  :disabled="disableInput"
                  @click.native="mixpanel.track('tap_preferred_mode_mail')"
                />
                <div
                  :class="
                    contactType == 1
                      ? 'tw-bg-black-900 tw-text-white'
                      : 'hover:tw-text-white hover:tw-bg-black-900 tw-text-black-900 tw-bg-white'
                  "
                  class="tw-cursor-pointer v2-normal-3-text contact-button tw-flex tw-items-center tw-border tw-border-black-900 tw-rounded-lg tw-px-6 tw-py-2 active:tw-scale-[0.98]"
                >
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6"
                    :icon="mail"
                    aria-hidden="true"
                  />
                  <span class="tw-ml-5">Mail</span>
                </div></label
              >
            </div>
          </div>
          <div class="tw-py-5 lg:tw-py-8 tw-flex tw-justify-center">
            <img
              v-if="showLoader"
              :src="loaderImage"
              class="tw-w-[64px] tw-h-[64px]"
            />
            <div v-else>
              <div class="v2-normal-3-text">
                <button
                  id="submit"
                  v-if="contactType == 1"
                  ref="recaptcha"
                  class="gradient-btn tw-py-4 tw-px-8 tw-m-0"
                  @click.prevent="submitForm()"
                >
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6"
                    :icon="planeIcon"
                    aria-hidden="true"
                  />
                  <span>Submit</span>
                </button>
                <button
                  v-if="contactType == 2"
                  class="gradient-btn tw-py-4 tw-px-8 tw-m-0"
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
          class="modal-mask tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-table mask tw-z-[1] tw-bg-[#00000080]"
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
    <!-- Show Thank you message -->
    <div v-if="showSuccessMessagePopup">
      <transition name="modal">
        <div
          class="modal-mask tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-table mask tw-z-[1] tw-bg-[#00000080]"
        >
          <div
            class="tw-mx-auto tw-left-auto sm:tw-my-7 sm:tw-mx-auto sm:tw-max-w-lg tw-h-full tw-flex tw-items-center"
            role="document"
          >
            <div
              class="tw-relative tw-flex tw-flex-col tw-w-full tw-bg-white tw-bg-clip-padding tw-rounded-md tw-outline-0 tw-border-1 tw-border-gray tw-border-solid tw-p-2.5 tw-rounded-3xl"
            >
              <div
                class="tw-flex tw-items-start tw-justify-between tw-p-4 tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300"
              >
                <div class="canopas-gradient-text normal-text text-center">
                  Thank you for choosing us to make a difference in your
                  business.
                </div>
              </div>
              <div class="tw-relative tw-p-4 tw-flex-auto">
                <div class="normal-2-text">
                  If you prefer to chat or email, sit back and relax our team
                  will get back to you within 24 hours.
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
          class="modal-mask tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-table mask tw-z-[1] tw-bg-[#00000080]"
        >
          <div
            class="tw-mx-auto tw-left-auto sm:tw-my-7 sm:tw-mx-auto sm:tw-max-w-lg tw-h-full tw-flex tw-items-center"
            role="document"
          >
            <div
              class="tw-relative tw-flex tw-flex-col tw-w-full tw-bg-white tw-bg-clip-padding tw-rounded-md tw-outline-0 tw-border-1 tw-border-gray tw-border-solid tw-p-2.5 tw-rounded-3xl"
            >
              <div
                class="tw-flex tw-items-start tw-justify-between tw-p-4 tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300 header-2-text"
              >
                <div class="canopas-gradient-text">Error</div>
              </div>
              <div class="tw-relative tw-p-4 tw-flex-auto">
                <div class="normal-text text-center">
                  {{ errorMessage }}
                </div>
                <div class="close-btn-div">
                  <button
                    class="gradient-btn tw-w-40 tw-px-0 tw-float-right"
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
import CalendlyIframe from "./CalendlyIframe.vue";
import config from "@/config.js";
import loaderImage from "@/assets/images/theme/small-loader.svg";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faPhoneVolume,
  faEnvelope,
  faCalendarAlt,
  faPaperPlane,
} from "@fortawesome/free-solid-svg-icons";

const CONTACT_BY_CHAT_OR_MAIL = 1;

export default {
  data() {
    return {
      phone: faPhoneVolume,
      mail: faEnvelope,
      planeIcon: faPaperPlane,
      calendarIcon: faCalendarAlt,
      loaderImage: loaderImage,
      name: "",
      email: "",
      projectInfo: "",
      reference: "",
      disableInput: false,
      showNameValidationError: false,
      showEmailValidationError: false,
      showValidEmailError: false,
      showProjectInfoValidationError: false,
      showReferenceValidationError: false,
      openCalendlyIframeModal: false,
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      contactType: 1,
    };
  },
  components: {
    FontAwesomeIcon,
    CalendlyIframe,
  },
  inject: ["mixpanel"],
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
      return (
        this.showNameValidationError ||
        this.showEmailValidationError ||
        this.showValidEmailError ||
        this.showProjectInfoValidationError ||
        this.showReferenceValidationError
      );
    },
    submitForm() {
      if (!this.validateForm()) {
        this.mixpanel.track("tap_submit_button_click");
        this.disableInput = true;

        let formData = {
          name: this.name,
          email: this.email,
          project_info: this.projectInfo ? this.projectInfo : "NA",
          reference: this.reference,
          contact_type:
            this.contactType == CONTACT_BY_CHAT_OR_MAIL
              ? "Chat or Email"
              : "Call",
        };

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
                  this.showLoader = true;
                  fbq("track", "contact_form_submit_success");
                  if (this.contactType == CONTACT_BY_CHAT_OR_MAIL) {
                    setTimeout(() => {
                      this.showSuccessMessage();
                      if (this.showSuccessMessagePopup) {
                        this.showLoader = false;
                      }
                    }, 1000);
                  } else {
                    setTimeout(() => {
                      this.mixpanel.track("tap_schedule_meeting_click");
                      this.openCalendlyIframe();
                      if (this.openCalendlyIframeModal) {
                        this.showLoader = false;
                      }
                    }, 1000);
                  }
                })
                .catch((err) => {
                  if (err.response.status == 401) {
                    this.errorMessage = "Invalid recaptcha score";
                  }
                  this.showErrorMessagePopup = true;
                });
            })
            .catch(() => {
              this.errorMessage = "Invalid recaptcha score";
              this.showErrorMessagePopup = true;
            });
        });
      }
    },
    openCalendlyIframe() {
      this.openCalendlyIframeModal = true;
      this.disableInput = false;
      this.showLoader = false;
      this.mixpanel.track("open_calendly_dialog_success");
    },
    closeCalendlyIframeModal() {
      this.openCalendlyIframeModal = false;
      this.mixpanel.track("close_calendly_dialog_error");
    },
    showSuccessMessage() {
      this.showSuccessMessagePopup = true;
      this.mixpanel.track("view_contact_success_dialog");
      setTimeout(() => {
        this.$router.push("/");
      }, 3500);
    },
  },
};
</script>
