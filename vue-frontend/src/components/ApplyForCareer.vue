<template>
  <div>
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <ScreenLoader v-if="isLoading" />
    <div v-else>
      <ScreenLoader v-if="isLoad" v-bind:loader="true" />
      <div class="container form-container">
        <div class="job-application">
          <div class="header-2-text text-center canopas-gradient-text pb-3">
            Applying For {{ title }}
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
                />
              </div>

              <div class="col-lg-12 col-md-12 col-sm-12 mb-5">
                <label class="">Message</label>
                <textarea
                  class="form-control custom"
                  name="message"
                  rows="1"
                  v-model="message"
                ></textarea>
              </div>

              <div class="col-lg-5 col-md-5 col-sm-12">
                <label class="required">Resume </label>
                <br />
                <label>
                  <i>Supported formats:</i
                  ><span class="black pdf-text">PDF</span>
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
                  accept="application/pdf"
                  @change="previewFiles"
                  required
                />

                <span v-if="showValidationError" class="error"
                  >This field is required</span
                >
              </div>
            </div>
            <div class="application-submit-btns mt-5">
              <button class="gradient-btn" @click.prevent="validateForm()">
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

        <!-- Show applicant details review popup -->
        <div v-if="showReviewFormPopup">
          <transition name="modal">
            <div class="modal-mask">
              <div class="modal-dialog modal-dialog-centered" role="document">
                <div class="modal-content">
                  <div class="modal-header">
                    <div class="normal-text canopas-gradient-text text-left">
                      Please review your details before submitting !
                    </div>
                    <button
                      type="button"
                      class="modal-close-btn"
                      data-dismiss="modal"
                      aria-label="Close"
                    >
                      <span
                        aria-hidden="true"
                        @click="showReviewFormPopup = false"
                        >&times;</span
                      >
                    </button>
                  </div>
                  <div class="modal-body">
                    <div class="">
                      <form>
                        <div class="mb-3">
                          <label for="fullname" class="col-form-label"
                            >Full name:</label
                          >
                          <input
                            type="text"
                            class="form-control"
                            v-model="fullName"
                            disabled
                          />
                        </div>
                        <div class="mb-3">
                          <label for="phonenumber" class="col-form-label"
                            >Phone:</label
                          >
                          <input
                            type="tel"
                            class="form-control"
                            v-model="phoneNumber"
                            disabled
                          />
                        </div>
                        <div class="mb-3">
                          <label for="email" class="col-form-label"
                            >Email:</label
                          >
                          <input
                            type="email"
                            class="form-control"
                            v-model="email"
                            disabled
                          />
                        </div>
                        <div class="mb-3" v-if="reference">
                          <label
                            for="howdidyoufindcanopas"
                            class="col-form-label"
                            >How did you find Canopas?</label
                          >
                          <input
                            type="text"
                            class="form-control"
                            v-model="reference"
                            disabled
                          />
                        </div>
                        <div class="mb-3 referenceBy" v-if="referenceBy">
                          <label for="referenceby" class="col-form-label"
                            >Reference By:</label
                          >
                          <input
                            type="text"
                            class="form-control"
                            v-model="referenceBy"
                            disabled
                          />
                        </div>
                      </form>
                    </div>
                  </div>
                  <div class="modal-footer">
                    <button
                      class="gradient-btn"
                      @click.prevent="submitApplication()"
                    >
                      <span> Submit </span>
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
                        type="submit"
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
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import ScreenMeta from "./partials/ScreenMeta.vue";
import ScreenLoader from "./utils/ScreenLoader.vue";
import axios from "axios";
import config from "./../config.js";
import router from "./../router";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import { library } from "@fortawesome/fontawesome-svg-core";
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons";
library.add(faCheckCircle);

export default {
  data() {
    return {
      title: "",
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
      isLoading: true,
      isLoad: false,
      isShowingReferenceInput: false,
      showValidationError: false,
      showEmailValidationError: false,
      showPhoneValidationError: false,
      fileButtonName: "Upload",
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      showReviewFormPopup: false,
      seoData: {
        type: "Jobs Posting Website",
        url: location.toString(),
      },
    };
  },
  components: {
    ScreenHeader,
    ScreenFooter2,
    ScreenMeta,
    FontAwesomeIcon,
    ScreenLoader,
  },
  mounted() {
    this.$gtag.event("view_page_job_apply");
    this.getCareerDetails();
    document.addEventListener("click", this.referenceList);
  },
  methods: {
    getCareerDetails() {
      var id = this.$route.params.id;
      axios
        .get(config.API_BASE + "/api/careers/" + id)
        .then((res) => {
          this.isLoading = false;
          this.title = res.data.title;
          this.prepareSEOdata(res.data);
        })
        .catch((err) => {
          this.isLoading = false;
          if (err.response && err.response.status == 404) {
            router.push({
              name: "Error404Page",
              params: { catchAll: "jobs/" + id },
            });
          } else {
            this.showErrorMessagePopup = true;
          }
        });
    },
    prepareSEOdata(job) {
      var seo_title = job.apply_seo_title
        ? job.apply_seo_title
        : "Apply for " + job.title + " job at canopas";

      this.seoData.title = seo_title;
      this.seoData.description = job.apply_seo_description;
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
        this.name === "" ||
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
        this.showReviewFormPopup = true;
      }
    },
    submitApplication() {
      this.$gtag.event("job_submit");
      this.showReviewFormPopup = false;
      this.isLoad = true;
      const formData = new FormData();
      formData.append("job_title", this.title);
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
          this.showSuccessMessagePopup = true;
          setTimeout(() => {
            this.$router.push("/");
          }, 2000);
        })
        .catch(() => {
          this.isLoad = false;
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

.form-container {
  margin: 80px auto 150px;
}

.job-application {
  border: 1px solid #e2e2e2;
  border-radius: 15px;
  padding: 32px 24px;
}

.form-control.custom {
  border: none;
  border-radius: 0;
  border-bottom: 1px solid #e2e2e2;
  box-shadow: none;
  color: #3d3d3d;
  font-size: 1.125rem;
  padding-left: 0;
}

.form-control.custom:focus {
  border-bottom: 1px solid #3d3d3d;
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
  padding: 16px 48px;
  border-radius: 10px;
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
  border-bottom: 1px solid #e2e2e2;
  outline: 0 !important;
}

.reference-option-list::before {
  content: "";
  position: absolute;
  top: 15px;
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
  border: none;
  outline: none;
  padding: 12px 0;
  border-radius: 10px;
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
}
</style>
