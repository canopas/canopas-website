<template>
  <div class="container">
    <form>
      <div class="information-detail">
        <div class="information-pd">
          <div class="row">
            <div class="col-lg-6 col-md-6 col-sm-12 tw-pt-5">
              <input
                class="custom-text-input v2-normal-text"
                type="text"
                name="username"
                required
                autocomplete="given-username"
                v-model="name"
                :disabled="disableInput"
                placeholder="Your name"
              />
              <span v-if="showValidationError" class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div class="col-lg-6 col-md-6 col-sm-12 tw-pt-5">
              <input
                class="custom-text-input v2-normal-text"
                type="text"
                name="email"
                required
                autocomplete="given-email"
                v-model="email"
                :disabled="disableInput"
                placeholder="Your email"
              />
              <span v-if="showValidationError" class="error tw-text-red-600"
                >This field is required</span
              >
            </div>
            <div class="col-lg-12 col-md-12 col-sm-12 tw-pt-5 lg:tw-pt-14">
              <textarea
                class="custom-text-input v2-normal-text"
                name="project"
                rows="1"
                autocomplete="given-project-info"
                v-model="projectInfo"
                :disabled="disableInput"
                placeholder="Tell us about your project"
              ></textarea>
            </div>
            <div class="col-lg-12 col-md-12 col-sm-12 tw-pt-5 lg:tw-pt-10">
              <input
                class="custom-text-input v2-normal-text"
                type="text"
                name="reference"
                required
                autocomplete="given-reference"
                v-model="reference"
                :disabled="disableInput"
                placeholder="How did you find us?"
              />
              <span v-if="showValidationError" class="error tw-text-red-600"
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
                />
                <div class="contact-button tw-mb-6 sm:tw-mr-9 sm:tw-mb-0">
                  <font-awesome-icon
                    class="fa -tw-rotate-45 tw-w-6 tw-h-6 tw-text-black-900"
                    :icon="phone"
                    aria-hidden="true"
                  />
                  <span class="tw-ml-5 v2-normal-3-text">Call</span>
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
                />
                <div class="v2-normal-3-text contact-button">
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6 tw-text-black-900"
                    :icon="mail"
                    aria-hidden="true"
                  />
                  <span class="tw-ml-5">Mail</span>
                </div></label
              >
            </div>
          </div>

          <div v-if="showLoader" class="tw-flex tw-py-5 lg:tw-py-16">
            <img :src="loaderImage" />
          </div>
          <div v-else>
            <div
              class="tw-py-5 lg:tw-py-8 v2-normal-3-text tw-flex tw-justify-center"
            >
              <button
                v-if="contactType == 1"
                class="gradient-btn tw-py-4 tw-px-8 tw-m-0"
                @click.prevent="submitApplication()"
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
                @click.prevent="submitApplication()"
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
    </form>
    <!-- Show Calendly Iframe -->
    <div v-if="openCalendlyIframeModal">
      <transition name="modal">
        <div class="modal-mask mask">
          <div
            class="modal-dialog login-modal modal-xl modal-dialog-centered"
            role="document"
          >
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-body">
                <CalendlyIframe />

                <button
                  type="button"
                  class="close modal-close-btn"
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
        <div class="modal-mask mask">
          <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
              <div class="modal-header">
                <div class="normal-text canopas-gradient-text text-center">
                  Thank you for choosing us to make a difference in your
                  business.
                </div>
              </div>
              <div class="modal-body">
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
        <div class="modal-mask mask">
          <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-header tw-justify-center">
                <div class="header-2-text canopas-gradient-text">Error</div>
              </div>
              <div class="modal-body">
                <div class="normal-text text-center">
                  Something went wrong on our side
                </div>
                <div class="close-btn-div">
                  <button
                    class="gradient-btn tw-float-right"
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
      showValidationError: false,
      openCalendlyIframeModal: false,
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      showLoader: false,
      contactType: 1,
    };
  },
  components: {
    FontAwesomeIcon,
    CalendlyIframe,
  },
  methods: {
    submitApplication() {
      if (this.name === "" || this.email === "" || this.reference === "") {
        this.showValidationError = true;
      } else {
        this.disableInput = true;
        this.showValidationError = false;

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

        axios
          .post(config.API_BASE + "/api/send-contact-mail", formData)
          .then(() => {
            this.showLoader = true;
            if (this.contactType == CONTACT_BY_CHAT_OR_MAIL) {
              setTimeout(() => {
                this.showSuccessMessage();
                if (this.showSuccessMessagePopup) {
                  this.showLoader = false;
                }
              }, 1000);
            } else {
              setTimeout(() => {
                this.openCalendlyIframe();
                if (this.openCalendlyIframeModal) {
                  this.showLoader = false;
                }
              }, 1000);
            }
          })
          .catch(() => {
            this.showErrorMessagePopup = true;
          });
      }
    },
    openCalendlyIframe() {
      this.openCalendlyIframeModal = true;
      this.disableInput = false;
      this.showLoader = false;
    },
    closeCalendlyIframeModal() {
      this.openCalendlyIframeModal = false;
    },
    showSuccessMessage() {
      this.showSuccessMessagePopup = true;
      setTimeout(() => {
        this.$router.push("/");
      }, 2000);
    },
  },
};
</script>

<style lang="postcss" scoped>
.container {
  @apply -tw-mt-16 tw-mb-24;
}

.information-detail {
  @apply tw-rounded-3xl tw-border tw-border-slate-200 tw-bg-white tw-shadow-3xl;
}

.information-pd {
  @apply tw-py-5 tw-px-8 lg:tw-px-24;
}

.custom-text-input {
  @apply tw-w-full tw-py-3 tw-px-0 tw-my-2 tw-mx-0 tw-box-border tw-border-b tw-border-slate-400 tw-rounded-none;
}

.custom-text-input:focus {
  @apply tw-box-border  tw-border-b tw-border-slate-400 tw-outline-none;
}

.custom-file-input {
  @apply tw-w-0 tw-p-0 tw-opacity-0;
}

.v2-title-2-text {
  @apply tw-leading-9 lg:tw-leading-11;
}

.contact-button {
  @apply tw-flex tw-items-center tw-border tw-border-black-900 tw-rounded-lg tw-px-6 tw-py-2  hover:tw-bg-black-900;
}

.contact-button:hover > span,
.contact-button:hover > .fa {
  @apply tw-text-white;
}

.contact-button-active {
  @apply tw-bg-black-900  hover:tw-bg-white;
}

.contact-button:active {
  transform: scale(0.98);
}

.gradient-btn {
  @apply tw-w-80 tw-px-0;
}

input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
  transition: background-color 5000s !important;
}

.modal-content {
  @apply tw-p-2.5 tw-rounded-3xl;
}

.modal-close-btn {
  @apply tw-border-none tw-text-pink-300 tw-text-4xl tw-font-light tw-bg-transparent tw-absolute tw-right-2.5 tw-top-0;
}

.modal-close-btn:focus {
  @apply tw-outline-none;
}

.modal-mask {
  @apply tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-table;
}

.mask {
  z-index: 1;
  background-color: rgba(0, 0, 0, 0.5);
  transition: opacity 0.3s ease;
}

::placeholder {
  @apply tw-text-slate-400;
}
</style>
