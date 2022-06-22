<template>
  <div>
    <ScreenHeaderV2 />
    <ScreenLoader v-if="isLoading || job == null" />
    <div v-else>
      <div class="container form-container">
        <div class="job-application">
          <div class="header-2-text text-center canopas-gradient-text pb-3">
            Applying For {{ job.title }}
          </div>
          <form class="contact-form-text pt-5 pb-5">
            <div class="row">
              <div class="col-lg-6 col-md-6 col-sm-12 mb-5">
                <label class="required">Full Name</label>
                <input
                  class="form-control custom"
                  type="text"
                  name="fullname"
                  required
                  autocomplete="given-name"
                  v-model="fullName"
                  :disabled="disableInput"
                />
                <span v-if="showValidationError" class="error"
                  >This field is required</span
                >
              </div>

              <div class="col-lg-6 col-md-6 col-sm-12 mb-5">
                <label class="required">Email</label>
                <input
                  class="form-control custom"
                  type="email"
                  name="email"
                  required
                  autocomplete="email"
                  v-model="email"
                  :disabled="disableInput"
                />
                <span v-if="showValidationError" class="error"
                  >This field is required</span
                >
                <span v-if="showEmailValidationError" class="error"
                  >Please enter valid email address</span
                >
              </div>

              <div class="col-lg-6 col-md-6 col-sm-12 mb-5">
                <label class="required">Phone Number</label>
                <input
                  class="form-control custom"
                  type="tel"
                  name="phonenumber"
                  autocomplete="tel"
                  v-model="phoneNumber"
                  :disabled="disableInput"
                />
                <span id="phoneError"></span>
                <span v-if="showValidationError" class="error"
                  >This field is required</span
                >
                <span v-if="showPhoneValidationError" class="error"
                  >Please enter valid phone number</span
                >
              </div>

              <div class="col-lg-6 col-md-6 col-sm-12 mb-5">
                <label class="">City</label>
                <input
                  class="form-control custom"
                  type="text"
                  name="city"
                  autocomplete="address-level2"
                  v-model="city"
                  :disabled="disableInput"
                />
              </div>

              <div class="col-lg-12 col-md-12 col-sm-12 mb-5">
                <label class="">How did you find Canopas?</label>
                <div class="reference-option-list" ref="referenceList">
                  <div>
                    <input
                      name="howdidyoufindcanopas"
                      type="text"
                      placeholder="Choose anyone"
                      id="reference-option-text"
                      readonly
                      v-model="reference"
                      :disabled="disableInput"
                    />
                  </div>
                  <div class="options">
                    <div
                      class="option"
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
                class="col-lg-12 col-md-12 col-sm-12 mb-5 reference-input"
              >
                <input
                  class="form-control custom reference-by-input"
                  type="text"
                  name="referenceby"
                  autocomplete="given-reference-name"
                  :placeholder="references[currentReferenceIndex].hint"
                  v-model="referenceBy"
                  :disabled="disableInput"
                />
              </div>

              <div class="col-lg-12 col-md-12 col-sm-12 mb-5">
                <label class="">Message</label>
                <textarea
                  class="form-control custom"
                  name="message"
                  rows="1"
                  v-model="message"
                  :disabled="disableInput"
                ></textarea>
              </div>

              <div class="col-lg-5 col-md-5 col-sm-12">
                <label class="required">Resume </label>
                <br />
                <label>
                  <i>Supported formats:</i
                  ><span class="black pdf-text">PDF and Docs </span>
                  <i>only</i>
                </label>
              </div>
              <div class="col-lg-7 col-md-7 col-sm-12 resume-upload-div">
                <button
                  type="button"
                  class="btn resume-upload-btn"
                  @click="chooseFiles()"
                >
                  {{ fileButtonName }}
                </button>
                <input
                  id="fileUpload"
                  type="file"
                  class="form-control custom-file-input"
                  name="resume"
                  accept="application/pdf,.doc,.docx"
                  @change="previewFiles"
                  required
                  :disabled="disableInput"
                />

                <span v-if="showValidationError" class="error"
                  >This field is required</span
                >
              </div>
            </div>
            <div class="application-submit-btns mt-5">
              <img class="loader-image" :src="loaderImage" v-if="showLoader" />
              <button
                v-else
                class="gradient-btn"
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
            <div class="required-field-msg">
              All fields marked with * are required.
            </div>
          </form>
        </div>

        <!-- Show Thank you message -->
        <div v-if="showSuccessMessagePopup">
          <transition name="modal">
            <div class="modal-mask">
              <div class="modal-dialog modal-dialog-centered" role="document">
                <div class="modal-content">
                  <div class="modal-header">
                    <div class="normal-text canopas-gradient-text text-center">
                      Congratulations !!
                    </div>
                  </div>
                  <div class="modal-body">
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
            <div class="modal-mask">
              <div class="modal-dialog modal-dialog-centered" role="document">
                <div class="modal-content">
                  <div class="modal-header justify-center">
                    <div class="header-2-text canopas-gradient-text">Error</div>
                  </div>
                  <div class="modal-body">
                    <div class="normal-text text-center">
                      Something went wrong on our side
                    </div>
                    <div class="close-btn-div">
                      <button
                        class="gradient-btn close-btn"
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
      disableInput: false,
      showLoader: false,
      loaderImage: loaderImage,
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
    this.setCareerDetails();
    this.$gtag.event("view_page_job_apply");
    document.addEventListener("click", this.referenceList);
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
    referenceList(e) {
      var target = e.target;
      if (target.id === "reference-option-text") {
        this.$refs.referenceList.classList.toggle("active");
      } else {
        this.$refs.referenceList.classList.remove("active");
      }
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

      axios
        .post(config.API_BASE + "/api/send-career-mail", formData)
        .then(() => {
          this.isLoad = false;
          this.showLoader = false;
          this.showSuccessMessagePopup = true;
          setTimeout(() => {
            this.$router.push("/jobs");
          }, 2000);
        })
        .catch(() => {
          this.isLoad = false;
          this.showLoader = false;
          this.showErrorMessagePopup = true;
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.gradient-btn {
  padding: 16px 64px;
}

.loader-image {
  margin: 1.25rem auto 0 auto;
  padding: 0 64px;
}

.form-container {
  margin: 80px auto 150px;
}

.job-application {
  border: 1px solid #e2e2e2;
  border-radius: 15px;
  padding: 32px 24px;
}

.form-control.custom {
  border: 1px solid #e2e2e2;
  border-radius: 10px;
  box-shadow: none;
  color: #3d3d3d;
  font-size: 1.125rem;
  margin-top: 15px;
  padding: 10px 16px !important;
}

.form-control.custom:focus {
  border: 1px solid #e2e2e2;
  outline: none;
}

.form-control:disabled {
  opacity: 0.8;
  cursor: not-allowed;
}

.contact-form-text,
.contact-form-text:hover {
  color: rgba(61, 61, 61, 0.8);
  font-size: 1rem;
  font-style: normal;
  font-weight: normal;
}

input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
  transition: background-color 5000s !important;
}

.error {
  margin-top: 8px;
  display: table;
  font-size: 1rem;
  color: #ff0000;
}

.required:after {
  content: "*";
  color: #ff0000;
}

.required-field-msg {
  margin-top: 48px;
  color: #ff0000;
}

.resume-upload-div {
  display: flex;
}

.custom-file-input {
  width: 0;
  opacity: 0;
  padding: 0;
}

.pdf-text {
  padding-left: 8px;
}

.resume-upload-btn,
.resume-upload-btn:focus,
.resume-upload-btn:hover {
  background: #fff;
  color: #3d3d3d;
  border: 1px solid #e2e2e2;
  box-shadow: none !important;
  outline: 0 !important;
  padding: 10px 48px;
  border-radius: 10px;
  height: 50px;
}

.application-submit-btns {
  display: flex;
  justify-content: center;
}

/* for dropdown list */

.reference-option-list {
  position: relative;
  margin-top: 15px;
  width: 100%;
  height: 50px;
  outline: 0 !important;
}

.reference-option-list::before {
  content: "";
  position: absolute;
  top: 20px;
  right: 20px;
  z-index: 1;
  width: 8px;
  height: 8px;
  border: 2px solid rgba(0, 0, 0, 0.5);
  border-top: 2px solid #fff;
  border-right: 2px solid #fff;
  transform: rotate(-45deg);
  transition: 0.5s;
  pointer-events: none;
}

.reference-option-list.active::before {
  top: 22px;
  transform: rotate(-225deg);
}

.reference-option-list input {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  cursor: pointer;
  background: #fff;
  border: 1px solid #e2e2e2;
  outline: none;
  padding: 12px 16px;
  border-radius: 10px;
  box-shadow: none;
}

.reference-option-list .options {
  position: absolute;
  top: 50px;
  width: 100%;
  background: #fff;
  overflow: hidden;
  display: none;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.reference-option-list.active .options {
  display: block;
}

.reference-option-list .options div {
  padding: 12px 12px;
  cursor: pointer;
}

.reference-option-list .options div:hover {
  background: rgba(226, 226, 226, 0.5);
}

.reference-by-input {
  padding-left: 0 !important;
}

.modal-close-btn {
  border: none !important;
  color: #f2709c;
  font-size: 40px;
  font-weight: lighter;
  background-color: transparent;
  position: absolute;
  right: 20px;
  top: -5px;
  font-weight: 300;
  font-size: 40px;
  margin-right: -10px;
}

.modal-close-btn:focus {
  outline: none !important;
}

.modal-mask {
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-content {
  padding: 10px;
  border-radius: 25px;
}

.close-btn {
  float: right;
  padding: 8px 16px !important;
}

.justify-center {
  justify-content: center;
}

@include media-breakpoint-up(md) {
  .gradient-btn {
    padding: 16px 80px;
  }

  .loader-image {
    padding: 0 80px;
  }

  .contact-form-text,
  .contact-form-text:hover {
    font-size: 1.125rem;
  }

  .job-application {
    padding: 48px 48px 0;
  }
}

@include media-breakpoint-up(lg) {
  .form-container {
    padding: 0 160px;
  }

  .loader-image {
    margin: 1.25rem 0.5rem 0 0.5rem;
  }
}
</style>
