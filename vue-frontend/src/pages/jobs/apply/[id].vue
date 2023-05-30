<template>
  <div @click="isShowReferenceOption = false">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <SuccessMessage v-if="showSuccessMessage" />
    <ScreenLoader v-if="isLoading || job == null" />

    <div v-else>
      <div
        class="tw-container tw-mt-[80px] tw-mx-auto tw-mb-[150px] lg:tw-py-0 lg:tw-px-[160px]"
      >
        <div
          class="tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[15px] tw-py-[32px] tw-px-[24px] md:tw-pt-[48px] md:tw-px-[48px] md:tw-pb-0"
        >
          <div class="header-2-text tw-text-center tw-pb-4">
            <h1 class="canopas-gradient-text">Applying For {{ job.title }}</h1>
          </div>

          <form
            class="tw-text-[#3d3d3dcc] tw-text-[1rem] md:tw-text-[1.125rem] tw-pt-12 tw-pb-12"
          >
            <div class="tw-grid tw-grid-cols-1 md:tw-grid-cols-2 tw-gap-4">
              <div class="tw-mb-[0.5rem]">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Full Name</label
                >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[5px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="text"
                  name="fullname"
                  autocomplete="given-name"
                  v-model="fullName"
                  required
                  @input="
                    showNameValidationError =
                      $event.target.value.trim().length === 0
                  "
                  aria-label="Full Name"
                />
                <span
                  v-if="showNameValidationError"
                  class="error tw-text-red-600 tw-text-[1rem]"
                  >Name is required</span
                >
              </div>

              <div class="tw-mb-[0.5rem]">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Email</label
                >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[5px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="text"
                  name="email"
                  id="email"
                  required
                  autocomplete="email"
                  v-model="email"
                  @input="
                    showEmailValidationError =
                      $event.target.value.trim().length === 0
                  "
                  @blur="showValidEmailError = isValidEmail()"
                  aria-label="User Email "
                />
                <span
                  v-if="showEmailValidationError"
                  class="error tw-text-red-600 tw-text-[1rem]"
                  >Email is required</span
                >
                <span
                  v-if="email.trim().length != 0 && showValidEmailError"
                  class="error tw-text-red-600 tw-text-[1rem]"
                  >Please enter valid email address</span
                >
              </div>

              <div class="tw-mb-[0.5rem]">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Phone Number</label
                >
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[5px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="tel"
                  name="phonenumber"
                  autocomplete="tel"
                  v-model="phoneNumber"
                  @blur="showValidPhoneError = isValidPhone()"
                  @input="
                    showPhoneValidationError =
                      $event.target.value.trim().length === 0
                  "
                  required
                  aria-label="Phone Number"
                />
                <span id="phoneError"></span>
                <span
                  v-if="showPhoneValidationError"
                  class="error tw-text-red-600 tw-text-[1rem]"
                  >Phone number is required</span
                >
                <span
                  v-if="phoneNumber.trim().length != 0 && showValidPhoneError"
                  class="error tw-text-red-600 tw-text-[1rem]"
                  >Please enter valid phone number</span
                >
              </div>

              <div class="tw-mb-[0.5rem]">
                <label>City</label>
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[5px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed"
                  type="text"
                  name="city"
                  autocomplete="address-level2"
                  v-model="city"
                  aria-label="City name"
                />
              </div>

              <div class="tw-mb-[0.5rem]">
                <label class="">How did you find Canopas?</label>
                <div
                  class="tw-relative tw-mt-[5px] tw-w-full tw-h-[50px] tw-outline-hidden tw-outline-0 before:tw-content-[''] before:tw-absolute before:tw-top-[20px] before:tw-right-[20px] before:tw-z-[1] before:tw-w-[8px] before:tw-h-[8px] before:tw-border-2 before:tw-border-solid before:tw-border-[#00000080] before:tw-border-t-2 before:tw-border-t-[#fff] before:tw-border-r-2 before:tw-border-r-[#fff] before:tw-rotate-[-45deg] before:tw-duration-[0.5s] before:tw-pointer-events-none"
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
                      placeholder="Choose any option"
                      id="reference-option-text"
                      readonly
                      v-model="reference"
                      aria-label="How did you find canopas"
                    />
                  </div>
                  <div
                    :style="{
                      display: isShowReferenceOption ? 'block' : 'none',
                    }"
                    class="tw-absolute tw-top-[50px] tw-z-[1] tw-w-full tw-bg-[#fff] tw-overflow-hidden tw-hidden tw-shadow-[0_5px_20px_rgba(0,0,0,0.1)] tw-rounded-[10px] tw-border-[1px] tw-border-solid tw-border-[#0000000d]"
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
                <label class="">Reference</label>
                <input
                  class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[5px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed reference-by-input"
                  type="text"
                  name="referenceby"
                  autocomplete="given-reference-name"
                  :placeholder="references[currentReferenceIndex].hint"
                  v-model="referenceBy"
                  aria-label="reference by"
                />
              </div>

              <div class="md:tw-col-span-2 tw-mb-[0.5rem]">
                <label
                  >Message
                  <textarea
                    class="tw-block tw-w-full tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[5px] tw-py-[10px] tw-px-[16px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed hidden-scrollbar tw-whitespace-pre-wrap"
                    name="message"
                    rows="3"
                    v-model="message"
                    @input="resizeTextarea"
                    placeholder="Tell us about yourself..."
                  ></textarea>
                </label>
              </div>

              <div class="">
                <label class="after:tw-content-['*'] after:tw-text-[#ff0000]"
                  >Resume
                </label>
                <br />
                <label class="tw-text-[#999]">
                  <i>Supported formats:</i
                  ><span class="black tw-pl-[8px]">PDF and Docs </span>
                  <i>only</i>
                </label>
              </div>
              <div class="tw-flex">
                <button
                  type="button"
                  class="resume-upload-btns tw-bg-[#fff] tw-text-[#3d3d3d] tw-truncate tw-border-[1px] tw-break-normal tw-border-solid tw-border-[#e2e2e2] tw-shadow-none tw-outline-hidden tw-outline-0 tw-py-[10px] tw-px-[48px] tw-rounded-[10px] tw-h-[50px] focus:tw-outline-hidden focus:tw-outline-0"
                  @click="chooseFiles()"
                >
                  {{ fileButtonName }}
                </button>
                <input
                  id="fileUpload"
                  ref="fileUpload"
                  type="file"
                  class="tw-block tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-rounded-[10px] tw-text-[#3d3d3d] tw-text-[1.125rem] tw-mt-[15px] focus:tw-border-[1px] focus:tw-border-solid focus:tw-border-[#e2e2e2] focus:tw-outline-hidden focus:tw-outline-0 disabled:tw-opacity-[0.8] disabled:tw-cursor-not-allowed custom-file-input tw-w-0 tw-opacity-0 tw-p-0 tw-pr-[10px]"
                  name="resume"
                  accept="application/pdf,.doc,.docx"
                  @change="previewFiles"
                  required
                  @input="showFileValidationError = fileUpload === ''"
                  aria-label="file upload"
                />

                <span
                  v-if="showFileValidationError"
                  class="tw-mt-2.5 error tw-text-red-600 tw-text-[1rem]"
                  >Resume is required</span
                >
              </div>
            </div>
            <div class="tw-flex tw-items-center tw-mt-8">
              <input
                type="checkbox"
                id="checkbox"
                class="md:tw-mt-[-25px] xl:tw-mt-[-23px] 2xl:tw-mt-[-17px] tw-w-[1.5rem] tw-h-[1.5rem] tw-rounded-[5px] tw-accent-pink-500"
                checked
              />
              <label
                class="tw-ml-2 tw-text-[1.125rem] tw-leading-[1.688rem] tw-font-inter-medium tw-text-black-core/[0.87]"
                for="checkbox"
                >I acknowledge that Canopas does not conduct remote interviews
                and does not offer remote job positions.</label
              >
            </div>
            <div class="tw-flex tw-justify-center tw-mt-12">
              <img
                class="tw-mt-[1.25rem] tw-mx-auto tw-mb-0 tw-py-0 tw-px-[64px] md:tw-px-[80px] lg:tw-mt-[1.25rem] lg:tw-mx-[0.5rem] lg:tw-mb-0"
                :src="loaderImage"
                v-if="showLoader"
                alt="loader-image"
              />
              <div
                class="tw-relative tw-flex tw-flex-col tw-items-center"
                v-else
              >
                <div
                  class="tw-absolute tw-text-center -tw-top-[2.5rem] sm:-tw-top-[1.875rem] tw-right-0 sm:-tw-right-[7rem] lg:-tw-right-[6rem] xl:-tw-right-[10rem] 2xl:-tw-right-[13.5rem] sm:tw-w-max"
                >
                  <span
                    v-if="showErrorMessage"
                    class="tw-flex tw-mr-0 sm:tw-mr-[2.5rem] md:tw-mr-[3.5rem] xl:tw-mr-[8.5rem] 2xl:tw-mr-[11.5rem] tw-text-center tw-text-red-600 tw-text-[1.3rem] sm:tw-text-[1.5rem]"
                    :class="
                      errorMessage == 'Invalid Recaptcha score'
                        ? '!tw-text-[1.5rem] tw-leading-[2rem] sm:!tw-mr-[6.5rem] md:!tw-mr-[8rem] lg:!tw-mr-[7.5rem] xl:!tw-mr-[12rem] 2xl:!tw-mr-[15rem]'
                        : ''
                    "
                    >{{ errorMessage }}</span
                  >
                </div>
                <button
                  class="gradient-btn tw-rounded-full tw-px-[64px] md:tw-px-[80px]"
                  @click.prevent="submitApplication()"
                >
                  <font-awesome-icon
                    class="fa"
                    :icon="checkCircle"
                    aria-hidden="true"
                  />
                  <span>Apply Now</span>
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
      <NewFooter />
    </div>
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import SuccessMessage from "@/components/contact/SuccessMessage.vue";
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
        image: config.OG_IMAGE_URL,
      },
    });
    return {
      meta,
    };
  },
  data() {
    return {
      event: "",
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
          hint: "Other referral source",
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
      showNameValidationError: false,
      showValidEmailError: false,
      showEmailValidationError: false,
      showPhoneValidationError: false,
      showValidPhoneError: false,
      showFileValidationError: false,
      fileButtonName: "Upload",
      showSuccessMessage: false,
      showErrorMessage: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      loaderImage: loaderImage,
      isShowReferenceOption: false,
    };
  },
  components: {
    Header,
    NewFooter,
    FontAwesomeIcon,
    ScreenLoader,
    SuccessMessage,
  },
  async serverPrefetch() {
    await this.setCareerDetails();
  },
  inject: ["mixpanel"],
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
    this.mixpanel.track("view_apply_job_page");
  },
  computed: {
    ...mapState(useJobDetailStore, {
      job: "item",
      jobsError: "error",
      isLoading: "isLoading",
    }),
  },
  methods: {
    resizeTextarea(event) {
      event.target.style.height = "auto";
      event.target.style.height = event.target.scrollHeight + "px";
    },
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
          this.mixpanel.track("job_apply_failed");
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
    isValidPhone() {
      const phoneRegex = /^[0-9]{10}$/;
      return !phoneRegex.test(this.phoneNumber);
    },
    isValidEmail() {
      var emailRegx =
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return !emailRegx.test(this.email);
    },
    validateForm() {
      this.showNameValidationError = this.fullName.trim().length === 0;
      this.showEmailValidationError = this.email.trim().length === 0;
      this.showPhoneValidationError = this.phoneNumber.trim().length === 0;
      this.showFileValidationError = this.fileButtonName === "Upload";

      return (
        this.showNameValidationError ||
        this.showPhoneValidationError ||
        this.showFileValidationError ||
        this.showEmailValidationError ||
        this.showValidPhoneError ||
        this.showValidEmailError
      );
    },
    submitApplication() {
      this.mixpanel.track("job_submit");
      if (!this.validateForm()) {
        this.showLoader = true;
        this.isLoad = true;

        //resume file name with the current date
        var year = new Date().getFullYear();
        var mon = new Date().toLocaleString("default", { month: "short" });
        var date = new Date().getDate();

        const currentDate = date + "-" + mon + "-" + year;

        const splitFileName = this.file.name.split(".");

        const fileName =
          splitFileName[0] +
          "-" +
          currentDate +
          "." +
          splitFileName[splitFileName.length - 1];

        //prepare form data
        const formData = new FormData();
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

        formData.append("file", this.file, fileName);
        formData.append("job_title", this.job.title);
        formData.append(
          "message",
          this.message
            ? this.message.replace(/\n/g, "<br>\n").replace(/\./g, ".\n")
            : "NA"
        );
        formData.append("save_record_to_spreadsheet", config.IS_PROD);
        formData.append("save_data_to_database", config.IS_PROD);
        localStorage.setItem("applicant-name", JSON.stringify(this.fullName));

        // verify recpatcha
        grecaptcha.enterprise.ready(() => {
          grecaptcha.enterprise
            .execute(import.meta.env.VITE_RECAPTCHA_SITE_KEY, {
              action: "verify",
            })
            .then((token) => {
              formData.append("token", token);
              axios
                .post(config.API_BASE + "/api/send-jobs-applications", formData)
                .then(() => {
                  this.isLoad = false;
                  this.showLoader = false;
                  this.fullName = "";
                  this.email = "";
                  this.phoneNumber = "";
                  this.city = "";
                  this.reference = "";
                  this.referenceBy = "";
                  this.message = "";
                  this.$refs.fileUpload.value = null;
                  this.fileButtonName = "Upload";
                  this.mixpanel.track("job_apply_success");
                  if (config.SHOW_JOBSTHANKYOU_PAGE) {
                    this.$router.push({
                      path: `/jobs/thank-you`,
                    });
                  } else {
                    this.showSuccessMessage = true;
                    setTimeout(() => {
                      this.showSuccessMessage = false;
                      this.$router.push({
                        path: `/jobs`,
                      });
                    }, 3000);
                  }
                })
                .catch((err) => {
                  this.isLoad = false;
                  this.showLoader = false;
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
              this.isLoad = false;
              this.showLoader = false;
              this.showErrorMessage = true;
              setTimeout(() => {
                this.showErrorMessage = false;
              }, 3000);
            });
        });
      }
    },
  },
};
</script>

<style lang="postcss" scoped>
.active::before {
  @apply tw-top-[22px] -tw-rotate-[225deg];
}
</style>
