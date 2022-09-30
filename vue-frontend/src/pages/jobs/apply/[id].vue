<template>
  <div @click="isShowReferenceOption = false">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <ScreenHeaderV2 />
    <ScreenLoader v-if="isLoading || job == null" />
    <div v-else>
      <div
        class="tw-container tw-mt-[80px] tw-mx-auto tw-mb-[150px] lg:tw-py-0 lg:tw-px-[160px]"
      >
        <div
          class="tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[15px] tw-py-[32px] tw-px-[24px] md:tw-pt-[48px] md:tw-px-[48px] md:tw-pb-0"
        >
          <div class="header-2-text tw-text-center tw-pb-4">
            <div class="canopas-gradient-text">
              Applying For {{ job.title }}
            </div>
          </div>

          <form
            class="tw-text-[#3d3d3dcc] tw-text-[1rem] md:tw-text-[1.125rem] tw-pt-12 tw-pb-12"
          >
            <div class="tw-grid tw-grid-cols-1 md:tw-grid-cols-2 tw-gap-4">
              <div class="tw-mb-8">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Full Name</label
                >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="text"
                  name="fullname"
                  required
                  autocomplete="given-name"
                  v-model="fullName"
                  :disabled="disableInput"
                />
                <span
                  v-if="showValidationError"
                  class="tw-mt-[8px] tw-table tw-text-[1rem] tw-text-[#ff0000]"
                  >This field is required</span
                >
              </div>

              <div class="tw-mb-8">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Email</label
                >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="email"
                  name="email"
                  required
                  autocomplete="email"
                  v-model="email"
                  :disabled="disableInput"
                />
                <span
                  v-if="showValidationError"
                  class="tw-mt-[8px] tw-table tw-text-[1rem] tw-text-[#ff0000]"
                  >This field is required</span
                >
                <span
                  v-if="showEmailValidationError"
                  class="tw-mt-[8px] tw-table tw-text-[1rem] tw-text-[#ff0000]"
                  >Please enter valid email address</span
                >
              </div>

              <div class="tw-mb-8">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Phone Number</label
                >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="tel"
                  name="phonenumber"
                  autocomplete="tel"
                  v-model="phoneNumber"
                  :disabled="disableInput"
                />
                <span id="phoneError"></span>
                <span
                  v-if="showValidationError"
                  class="tw-mt-[8px] tw-table tw-text-[1rem] tw-text-[#ff0000]"
                  >This field is required</span
                >
                <span
                  v-if="showPhoneValidationError"
                  class="tw-mt-[8px] tw-table tw-text-[1rem] tw-text-[#ff0000]"
                  >Please enter valid phone number</span
                >
              </div>

              <div class="tw-mb-8">
                <label class="">City</label>
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="text"
                  name="city"
                  autocomplete="address-level2"
                  v-model="city"
                  :disabled="disableInput"
                />
              </div>

              <div class="md:tw-col-span-2 tw-mb-8">
                <label class="">How did you find Canopas?</label>
                <div
                  class="tw-relative tw-mt-[15px] tw-w-full tw-h-[50px] tw-outline-hidden tw-outline-0 before:tw-content-[''] before:tw-absolute before:tw-top-[20px] before:tw-right-[20px] before:tw-z-[1] before:tw-w-[8px] before:tw-h-[8px] before:tw-border-2 before:tw-border-solid before:tw-border-[#00000080] before:tw-border-t-2 before:tw-border-t-[#fff] before:tw-border-r-2 before:tw-border-r-[#fff] before:tw-rotate-[-45deg] before:tw-duration-[0.5s] before:tw-pointer-events-none"
                  v-on:click.stop="
                    isShowReferenceOption = !isShowReferenceOption
                  "
                  :class="isShowReferenceOption ? 'active' : ''"
                >
                  <div>
                    <input
                      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-full tw-cursor-pointer tw-bg-[#fff] tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-outline-hidden tw-outline-0 tw-py-[12px] tw-px-[16px] tw-rounded-[10px]"
                      name="howdidyoufindcanopas"
                      type="text"
                      placeholder="Choose anyone"
                      id="reference-option-text"
                      readonly
                      v-model="reference"
                      :disabled="disableInput"
                    />
                  </div>
                  <div
                    :style="{
                      display: isShowReferenceOption ? 'block' : 'none',
                    }"
                    class="tw-absolute tw-top-[50px] tw-w-full tw-bg-[#fff] tw-overflow-hidden tw-hidden tw-shadow-[0_5px_20px_rgba(0,0,0,0.1)] tw-rounded-[10px] tw-border-[1px] tw-border-solid tw-border-[#0000000d]"
                  >
                    <div
                      class="tw-py-[12px] tw-px-[12px] tw-cursor-pointer hover:tw-bg-[#e2e2e280]"
                      v-for="reference in references"
                      :key="reference"
                      @click="showOptions(reference)"
                    >
                      {{ reference.option }}
                    </div>
                  </div>
                </div>
              </div>

              <div
                v-if="isShowingReferenceInput"
                class="tw-mb-8 reference-input"
              >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed reference-by-input tw-pl-0"
                  type="text"
                  name="referenceby"
                  autocomplete="given-reference-name"
                  :placeholder="references[currentReferenceIndex].hint"
                  v-model="referenceBy"
                  :disabled="disableInput"
                />
              </div>

              <div class="md:tw-col-span-2 tw-mb-8">
                <label class="">Message</label>
                <textarea
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  name="message"
                  rows="1"
                  v-model="message"
                  :disabled="disableInput"
                ></textarea>
              </div>

              <div class="">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Resume
                </label>
                <br />
                <label>
                  <i>Supported formats:</i
                  ><span class="black tw-pl-[8px]">PDF and Docs </span>
                  <i>only</i>
                </label>
              </div>
              <div class="tw-flex">
                <button
                  type="button"
                  class="resume-upload-btns tw-bg-[#fff] tw-text-[#3d3d3d] tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-shadow-none tw-outline-hidden tw-outline-0 tw-py-[10px] tw-px-[48px] tw-rounded-[10px] tw-h-[50px] focus:tw-outline-hidden focus:tw-outline-0"
                  @click="chooseFiles()"
                >
                  {{ fileButtonName }}
                </button>
                <input
                  id="fileUpload"
                  type="file"
                  class="tw-block tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed custom-file-input tw-w-0 tw-opacity-0 tw-p-0"
                  name="resume"
                  accept="application/pdf,.doc,.docx"
                  @change="previewFiles"
                  required
                  :disabled="disableInput"
                />

                <span
                  v-if="showValidationError"
                  class="tw-mt-[8px] tw-table tw-text-[1rem] tw-text-[#ff0000]"
                  >This field is required</span
                >
              </div>
            </div>
            <div class="tw-flex tw-justify-center tw-mt-12">
              <img
                class="tw-mt-[1.25rem] tw-mx-auto tw-mb-0 tw-py-0 tw-px-[64px] md:tw-px-[80px] lg:tw-mt-[1.25rem] lg:tw-mx-[0.5rem] lg:tw-mb-0"
                :src="loaderImage"
                v-if="showLoader"
              />
              <button
                v-else
                class="gradient-btn tw-py-[16px] tw-px-[64px] md:tw-py-[16px] md:tw-px-[80px]"
                @click.prevent="validateForm()"
              >
                <font-awesome-icon
                  class="fa"
                  :icon="checkCircle"
                  aria-hidden="true"
                />
                <span>Apply Now</span>
              </button>
            </div>
            <div class="tw-mt-[48px] tw-text-[#ff0000]">
              All fields marked with * are required.
            </div>
          </form>
        </div>

        <!-- Show Thank you message -->
        <div v-if="showSuccessMessagePopup">
          <transition name="modal">
            <div
              class="modal-mask tw-fixed tw-z-[1] tw-top-0 tw-left-0 tw-w-full tw-h-full tw-bg-[#00000080] tw-table"
            >
              <div
                class="tw-mx-auto tw-left-auto sm:tw-my-7 sm:tw-mx-auto sm:tw-max-w-lg tw-h-full tw-flex tw-items-center"
                role="document"
              >
                <div
                  class="tw-relative tw-flex tw-flex-col tw-w-full tw-bg-white tw-bg-clip-padding tw-rounded-md tw-outline-0 tw-border-1 tw-border-gray tw-border-solid tw-p-[10px] tw-rounded-[25px]"
                >
                  <div
                    class="tw-flex tw-items-start tw-justify-between tw-p-4 tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300"
                  >
                    <div class="normal-text canopas-gradient-text text-center">
                      Congratulations !!
                    </div>
                  </div>
                  <div class="tw-relative tw-p-4 tw-flex-auto">
                    <div class="normal-2-text">
                      We have received your job application, sit back and relax!
                      Our team will contact you within 24 hours.
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
              class="modal-mask tw-fixed tw-z-[1] tw-top-0 tw-left-0 tw-w-full tw-h-full tw-bg-[#00000080] tw-table"
            >
              <div
                class="tw-mx-auto tw-left-auto sm:tw-my-7 sm:tw-mx-auto sm:tw-max-w-lg tw-h-full tw-flex tw-items-center"
                role="document"
              >
                <div
                  class="tw-relative tw-flex tw-flex-col tw-w-full tw-bg-white tw-bg-clip-padding tw-rounded-md tw-outline-0 tw-border-1 tw-border-gray tw-border-solid tw-p-[10px] tw-rounded-[25px]"
                >
                  <div
                    class="tw-flex tw-items-start tw-justify-between tw-p-4 tw-rounded-t-md tw-border-b tw-border-solid tw-border-slate-300"
                  >
                    <div class="header-2-text canopas-gradient-text">Error</div>
                  </div>
                  <div class="tw-relative tw-p-4 tw-flex-auto">
                    <div class="normal-text text-center">
                      {{ errorMessage }}
                    </div>
                    <div class="close-btn-div">
                      <button
                        class="gradient-btn tw-float-right tw-py-[8px] tw-px-[16px]"
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
      <ScreenFooter2 />
    </div>
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import axios from "axios";
import config from "@/config.js";
import { useJobDetailStore } from "@/stores/jobs";
import { mapState } from "pinia";
import { mapActions } from "pinia";
import { useMeta } from "vue-meta";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import loaderImage from "@/assets/images/theme/small-loader.svg";

import { library } from "@fortawesome/fontawesome-svg-core";
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons";
library.add(faCheckCircle);

export default {
  setup() {
    const { meta } = useMeta({
      og: {
        type: "Jobs Posting Website",
      },
    });
    return {
      meta,
    };
  },
  data() {
    return {
      id: this.$route.params.id,
      references: [
        {
          id: 1,
          option: "Instagram",
          hint: "",
        },
        {
          id: 2,
          option: "LinkedIn",
          hint: "",
        },
        {
          id: 3,
          option: "Google search",
          hint: "",
        },
        {
          id: 4,
          option: "Facebook",
          hint: "",
        },
        {
          id: 5,
          option: "Friends or Family",
          hint: "",
        },
        {
          id: 6,
          option: "Canopas Employee",
          hint: "Employee name",
        },
        {
          id: 7,
          option: "College",
          hint: "",
        },
        {
          id: 8,
          option: "Job posting website",
          hint: "Website name",
        },
        {
          id: 9,
          option: "Other",
          hint: "Please specify referral source",
        },
      ],
      checkCircle: faCheckCircle,
      currentReferenceIndex: -1,
      fullName: "",
      email: "",
      phoneNumber: "",
      city: "",
      reference: "",
      referenceBy: "",
      message: "",
      file: "",
      isLoad: false,
      isShowingReferenceInput: false,
      showValidationError: false,
      showEmailValidationError: false,
      showPhoneValidationError: false,
      fileButtonName: "Upload",
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      errorMessage: "Something went wrong on our side",
      disableInput: false,
      showLoader: false,
      loaderImage: loaderImage,
      isShowReferenceOption: false,
    };
  },
  components: {
    ScreenHeaderV2,
    ScreenFooter2,
    FontAwesomeIcon,
    ScreenLoader,
  },
  async serverPrefetch() {
    await this.setCareerDetails();
  },
  mounted() {
    let recaptchaScript = document.createElement("script");
    recaptchaScript.setAttribute(
      "src",
      "https://www.google.com/recaptcha/enterprise.js?render=" +
        import.meta.env.VITE_RECAPTCHA_SITE_KEY
    );
    recaptchaScript.setAttribute("async", "true");
    recaptchaScript.setAttribute("defer", "true");
    document.head.appendChild(recaptchaScript);
    this.setCareerDetails();
    this.$gtag.event("view_page_job_apply");
  },
  computed: {
    ...mapState(useJobDetailStore, {
      job: "item",
      jobsError: "error",
      isLoading: "isLoading",
    }),
  },
  methods: {
    ...mapActions(useJobDetailStore, ["loadJob"]),
    async setCareerDetails() {
      await this.loadJob(this.id, this.$route.href);

      if (this.jobsError != null) {
        var err = this.jobsError;
        if (err && err.response && err.response.status == 404) {
          this.$router.push({
            name: "Error404Page",
            params: { pathMatch: ["jobs", "apply", this.id] },
          });
        } else {
          this.showErrorMessagePopup = true;
        }
      } else {
        this.setMetaProperties();
      }
    },
    setMetaProperties() {
      this.meta.title = this.job.apply_seo_title;
      this.meta.description = this.job.apply_seo_description;
      this.meta.og.title = this.job.apply_seo_title;
      this.meta.og.url = config.BASE_URL + this.$route.href;
    },
    showOptions(reference) {
      this.reference = reference.option;
      let names = ["Canopas Employee", "Job posting website", "Other"];
      if (names.includes(reference.option)) {
        this.isShowingReferenceInput = true;
        this.currentReferenceIndex = reference.id - 1;
      } else {
        this.isShowingReferenceInput = false;
      }
    },
    previewFiles(event) {
      this.file = event.target.files[0];
      if (this.file.name != null) {
        this.fileButtonName = this.file.name;

        this.fileButtonName.length > 10
          ? this.fileButtonName.substr(0, 20) + "..."
          : this.fileButtonName;
      }
    },
    chooseFiles: function () {
      document.getElementById("fileUpload").click();
    },
    validateForm() {
      var emailRegx =
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

      if (
        this.fullName === "" ||
        this.email === "" ||
        this.phoneNumber === "" ||
        this.fileButtonName == "Upload"
      ) {
        this.showValidationError = true;
      } else if (!this.email.match(emailRegx)) {
        this.showEmailValidationError = true;
      } else if (!this.phoneNumber.match(/[0-9]{10}/)) {
        this.showPhoneValidationError = true;
      } else {
        this.showValidationError = false;
        this.showPhoneValidationError = false;
        this.showEmailValidationError = false;
        this.disableInput = true;
        this.submitApplication();
      }
    },
    submitApplication() {
      this.$gtag.event("job_submit");
      this.showLoader = true;
      this.disableInput = false;
      this.isLoad = true;

      //prepare form data
      const formData = new FormData();
      formData.append("job_title", this.job.title);
      formData.append("name", this.fullName);
      formData.append("email", this.email);
      formData.append("phone", this.phoneNumber);
      formData.append("place", this.city ? this.city : "NA");
      formData.append(
        "references",
        this.reference &&
          this.reference != "" &&
          this.referenceBy &&
          this.referenceBy != ""
          ? this.reference + " - " + this.referenceBy
          : this.reference && this.reference != ""
          ? this.reference
          : "NA"
      );
      formData.append("message", this.message ? this.message : "NA");
      formData.append("file", this.file, this.file.name);
      formData.append("save_record_to_spreadsheet", config.IS_PROD);

      //verify recpatcha
      grecaptcha.enterprise.ready(() => {
        grecaptcha.enterprise
          .execute(import.meta.env.VITE_RECAPTCHA_SITE_KEY, {
            action: "verify",
          })
          .then((token) => {
            formData.append("token", token);
            axios
              .post(config.API_BASE + "/api/send-career-mail", formData)
              .then(() => {
                this.isLoad = false;
                this.showLoader = false;
                this.showSuccessMessagePopup = true;
                setTimeout(() => {
                  this.$router.push("/jobs");
                }, 3500);
              })
              .catch((err) => {
                this.isLoad = false;
                this.showLoader = false;
                if (err.response.status == 401) {
                  this.errorMessage = "Invalid recaptcha score";
                }
                this.showErrorMessagePopup = true;
              });
          })
          .catch(() => {
            this.errorMessage = "Invalid recaptcha score";
            this.isLoad = false;
            this.showLoader = false;
            this.showErrorMessagePopup = true;
          });
      });
    },
  },
};
</script>

<style lang="postcss" scoped>
.active::before {
  @apply tw-top-[22px] -tw-rotate-[225deg];
}
</style>
