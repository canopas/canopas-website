<template>
  <div class="tw-mb-9 md:tw-mb-24 tw-font-inter-regular">
    <form>
      <div class="tw-py-5 tw-px-8 lg:tw-px-24">
        <div class="tw-grid tw-grid-cols-1 md:tw-grid-cols-2 tw-gap-4">
          <div class="tw-relative md:tw-mb-5 tw-pt-3 lg:tw-pt-9 tw-text-left">
            <input
              type="text"
              id="username"
              class="floating-input tw-peer"
              name="username"
              required
              autocomplete="given-username"
              v-model="name"
              :disabled="disableInput"
              placeholder=" "
            />
            <label for="username" class="floating-label">Your name</label>
            <span v-if="showValidationError" class="error tw-text-red-600"
              >This field is required</span
            >
          </div>
          <div class="tw-relative md:tw-mb-5 tw-pt-3 lg:tw-pt-9 tw-text-left">
            <input
              class="floating-input tw-peer"
              type="text"
              name="email"
              id="email"
              required
              autocomplete="given-email"
              v-model="email"
              :disabled="disableInput"
              placeholder=" "
            />
            <label for="email" class="floating-label">Your email</label>
            <span v-if="showValidationError" class="error tw-text-red-600"
              >This field is required</span
            >
          </div>
          <div
            class="tw-relative md:tw-col-span-2 md:tw-mb-5 tw-pt-3 lg:tw-pt-10 tw-text-left"
          >
            <textarea
              class="floating-input tw-peer"
              id="project"
              name="project"
              rows="1"
              required
              autocomplete="given-project-info"
              v-model="projectInfo"
              :disabled="disableInput"
              placeholder=" "
            ></textarea>
            <label for="project" class="floating-label"
              >Tell us about your project</label
            >
            <span v-if="showValidationError" class="error tw-text-red-600"
              >This field is required</span
            >
          </div>
          <div
            class="tw-relative md:tw-col-span-2 md:tw-mb-5 tw-pt-3 lg:tw-pt-10 tw-text-left"
          >
            <input
              class="floating-input tw-peer"
              type="text"
              name="reference"
              id="reference"
              required
              autocomplete="given-reference"
              v-model="reference"
              :disabled="disableInput"
              placeholder=" "
            />
            <label for="reference" class="floating-label"
              >How did you find us?</label
            >
            <span v-if="showValidationError" class="error tw-text-red-600"
              >This field is required</span
            >
          </div>
          <div class="tw-relative md:tw-col-span-2 tw-text-left">
            <div ref="invest-list">
              <button
                class="form-font tw-flex tw-items-center tw-justify-between tw-w-full tw-py-2 tw-px-0 tw-my-2 tw-mx-0 tw-bg-none tw-border-b tw-border-white/[.6] tw-font-medium focus:tw-outline-0 active:tw-outline-0 focus:tw-shadow-none active:tw-shadow-none focus:tw-ring-0 active:tw-ring-0 focus:tw-bg-transparent active:tw-bg-transparent active:tw-text-white tw-transition tw-duration-150 tw-ease-in-out tw-whitespace-nowrap"
                :class="
                  invest == `I'll invest`
                    ? '!tw-text-white/[.6]'
                    : '!tw-text-white'
                "
                type="button"
                id="invest"
                name="invest"
                required
                :disabled="disableInput"
                @click="showList = !showList"
              >
                {{ invest }}
                <font-awesome-icon
                  class="fab tw-w-[15px] tw-h-[15px] tw-mr-0"
                  :icon="showList ? faCaretUp : faCaretDown"
                />
              </button>
              <ul
                v-if="showList"
                class="tw-absolute tw-w-full tw-min-w-max tw-py-2 tw-mt-1 tw-m-0 tw-bg-white tw-bg-clip-padding tw-border-none tw-float-left tw-list-none tw-text-left tw-text-base tw-shadow-lg tw-transition-all tw-delay-150 tw-duration-300 tw-z-[2]"
              >
                <li
                  v-for="option in investOptions"
                  :key="option"
                  @click="setOption(option)"
                >
                  <span
                    class="tw-block tw-w-full tw-py-2 tw-px-4 tw-bg-transparent tw-font-normal tw-whitespace-nowrap tw-text-gray-700 tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.1875rem] md:tw-leading-[1.4375rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.6875rem] hover:tw-bg-[#3D3D3D] hover:tw-text-white/[.6]"
                    >{{ option }}</span
                  >
                </li>
              </ul>
            </div>
            <span v-if="showValidationError" class="error tw-text-red-600"
              >This field is required</span
            >
          </div>
          <div
            class="tw-inline-flex tw-items-center tw-relative md:tw-col-span-2 tw-pt-3 lg:tw-pt-10 tw-cursor-pointer tw-text-left"
          >
            <span
              class="form-font tw-mr-6 tw-text-sm tw-font-medium tw-text-gray-900 dark:tw-text-gray-300 tw-text-white/[.6]"
              @click="NDA = !NDA"
              >Send me NDA</span
            >
            <label
              class="tw-inline-flex tw-relative tw-items-center tw-cursor-pointer"
            >
              <input type="checkbox" v-model="NDA" class="tw-sr-only tw-peer" />
              <div
                class="tw-w-9 tw-h-4 tw-bg-white/[.6] peer-focus:tw-outline-none tw-rounded-full tw-peer peer-checked:after:tw-translate-x-full peer-checked:after:tw-border-white/[.6] after:tw-content-[''] after:tw-absolute after:tw-top-[-2px] after:tw-left-[-2px] after:tw-bg-white after:tw-border-white/[.6] after:tw-border after:tw-rounded-full after:tw-h-5 after:tw-w-5 after:tw-transition-all peer-checked:tw-bg-gradient-to-r peer-checked:tw-from-[#F2709C] peer-checked:tw-to-[#FF9472]"
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
              class="gradient-btn tw-rounded-[45px] tw-px-[1.3125rem] tw-py-[0.1625rem] md:tw-py-[1.1575rem] md:tw-px-[1.8125rem] tw-m-0 tw-text-[1rem] tw-leading-[2.25rem] md:tw-text-[1.125rem] md:tw-leading-[1.34375rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.5rem]"
              @click.prevent="submitForm()"
            >
              <span class="!tw-font-semibold">Send Message</span>
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
            class="sm:tw-mx-auto sm:tw-max-w-lg tw-h-full tw-flex tw-items-center"
          >
            <div class="tw-w-full tw-bg-white tw-rounded-md">
              <div
                class="tw-p-4 tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300"
              >
                <div
                  class="canopas-gradient-text tw-text-[1.75rem] tw-leading-[2.25rem] tw-text-center"
                >
                  Thank you for choosing us to make a difference in your
                  business.
                </div>
              </div>
              <div class="tw-relative tw-p-4 tw-flex-auto">
                <div class="normal-2-text">
                  If you prefer to chat or email, sit back and relax our team
                  will get back to you within 24 hours.
                </div>
                <div class="close-btn-div">
                  <button
                    class="gradient-btn tw-w-40 tw-px-0 tw-mb-2"
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
            class="sm:tw-mx-auto sm:tw-max-w-lg tw-h-full tw-flex tw-items-center"
          >
            <div
              class="tw-w-full tw-bg-white tw-bg-clip-padding tw-rounded-md tw-p-2"
            >
              <div
                class="tw-p-3 tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300"
              >
                <div
                  class="canopas-gradient-text tw-text-[1.75rem] tw-leading-[2.25rem] tw-font-semibold"
                >
                  Error
                </div>
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
import config from "@/config.js";
import loaderImage from "@/assets/images/theme/small-loader.svg";
import { faCaretDown, faCaretUp, faL } from "@fortawesome/free-solid-svg-icons";
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
      invest: "I'll invest",
      investOptions: [
        "< USD 50000",
        "USD 50000 - USD 100000",
        "USD 100000 - USD 250000",
        "> USD 250000",
      ],
      NDA: false,
      disableInput: false,
      showValidationError: false,
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
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
    submitForm() {
      if (
        this.name === "" ||
        this.email === "" ||
        this.projectInfo === "" ||
        this.reference === "" ||
        this.invest === "I'll invest"
      ) {
        this.showValidationError = true;
      } else {
        this.disableInput = true;
        this.showValidationError = false;
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
                });
            })
            .catch(() => {
              this.errorMessage = "Invalid recaptcha score";
              this.showErrorMessagePopup = true;
            });
        });
      }
    },
    closePopUps(e) {
      const showList = this.$refs["invest-list"];
      if (showList && !showList.contains(e.target)) {
        this.showList = false;
      }
    },
    setOption(option) {
      this.invest = option;
      this.showList = false;
    },
  },
};
</script>

<style lang="postcss" scoped>
.floating-input {
  @apply tw-text-lg md:tw-text-xl lg:tw-text-2xl tw-block tw-w-full tw-px-0 tw-my-2 tw-mx-0 tw-text-white tw-bg-transparent tw-rounded-none tw-border-b tw-border-white/[.6] tw-placeholder-white/[.6] focus:tw-outline-0 active:tw-outline-0 focus:tw-ring-0 active:tw-ring-0 tw-transition tw-ease-in-out tw-appearance-none;
}

.floating-label {
  @apply form-font tw-absolute tw-duration-300 tw-mb-5 tw-transform tw--translate-y-4 tw-scale-75 tw-top-4 tw-z-[2] tw-origin-[0] tw-left-0 peer-focus:tw-text-white peer-placeholder-shown:tw-scale-100 peer-placeholder-shown:tw-translate-y-0 peer-focus:tw-scale-75 peer-focus:tw--translate-y-4;
}

.form-font {
  @apply tw-text-white/[.6] tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem];
}
</style>
