<template>
  <div>
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <div v-if="isLoading" class="loader-div">
      <img :src="loader" />
    </div>
    <div
      class="container form-container"
      :style="{
        'pointer-events': isLoading ? 'none' : '',
        filter: isLoading ? 'blur(1px)' : '',
      }"
    >
      <div class="job-application">
        <div class="header-text text-center canopas-gradient-text">
          Applying For {{ title }}
        </div>
        <form class="contact-form-text apply-form">
          <div class="row">
            <span class="required-field-msg"
              >All fields marked with * are required.</span
            >

            <div class="col-lg-12 col-md-12 col-sm-12 mb-5">
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

            <div class="col-lg-12 col-md-12 col-sm-12 mb-5">
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
                    class="reference-option-text"
                    type="text"
                    placeholder="choose any one"
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
                rows="5"
                v-model="message"
              ></textarea>
            </div>

            <div class="col-lg-5 col-md-5 col-sm-12">
              <label class="required"
                >Resume <br />
                <i>Supported formats:</i
                ><span class="black pdf-text">.PDF</span>
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
              <div class="file-name-wrapper" hidden>
                <span class="file-name"></span>
              </div>
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
        </form>
      </div>

      <!-- Show applicant details review popup -->
      <div v-if="showReviewFormPopup">
        <transition name="modal">
          <div class="modal-mask">
            <div class="modal-dialog modal-dialog-centered" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <div class="jobs-normal-text canopas-gradient-text text-left">
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
                        <label for="email" class="col-form-label">Email:</label>
                        <input
                          type="email"
                          class="form-control"
                          v-model="email"
                          disabled
                        />
                      </div>
                      <div class="mb-3">
                        <label for="howdidyoufindcanopas" class="col-form-label"
                          >How did you find Canopas?</label
                        >
                        <input
                          type="text"
                          class="form-control"
                          v-model="reference"
                          disabled
                        />
                      </div>
                      <div class="mb-3 referenceBy">
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
                  <div
                    class="jobs-normal-text canopas-gradient-text text-center"
                  >
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
          <div class="modal-mask">
            <div class="modal-dialog modal-dialog-centered" role="document">
              <div class="modal-content">
                <div class="modal-header justify-center">
                  <div class="header-2-text canopas-gradient-text">Error</div>
                </div>
                <div class="modal-body">
                  <div class="jobs-normal-text text-center">
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
    <ScreenFooter v-if="!showJobs" />
    <ScreenFooter2 v-if="showJobs" />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import ScreenMeta from "./partials/ScreenMeta.vue";
import axios from "axios";
import config from "@/config.js";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import loader from "@/assets/images/theme/loader.svg";

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
      loader: loader,
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
      isShowingReferenceInput: false,
      showValidationError: false,
      showEmailValidationError: false,
      showPhoneValidationError: false,
      fileButtonName: "Upload",
      showSuccessMessagePopup: false,
      showErrorMessagePopup: false,
      showReviewFormPopup: false,
      showJobs: config.IS_SHOW_JOBS,
      seoData: {
        type: "Jobs Posting Website",
        url: location.toString(),
      },
    };
  },
  components: {
    ScreenHeader,
    ScreenFooter,
    ScreenFooter2,
    ScreenMeta,
    FontAwesomeIcon,
  },
  mounted() {
    this.getCareerDetails();
    document.addEventListener("click", this.referenceList);
  },
  methods: {
    getCareerDetails() {
      axios
        .get(config.API_BASE + "/api/careers/" + this.$route.params.id)
        .then((res) => {
          this.isLoading = false;
          this.title = res.data.title;
          this.prepareSEOdata(res.data);
        })
        .catch(() => {
          this.isLoading = false;
          this.showErrorMessagePopup = true;
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
      this.showReviewFormPopup = false;
      this.isLoading = true;
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
          this.isLoading = false;
          this.showSuccessMessagePopup = true;
          setTimeout(() => {
            this.$router.push("/");
          }, 2000);
        })
        .catch(() => {
          this.isLoading = false;
          this.showErrorMessagePopup = true;
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.gradient-btn {
  padding: 1rem 4rem;
}

.form-container {
  margin: 5rem auto;
}

.job-application {
  border: 1px solid #e2e2e2;
  border-radius: 15px;
  padding: 2rem 1.5rem;
}

.apply-form {
  padding: 3rem 0;
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
  color: rgba(0, 0, 0, 0.5);
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
  margin-top: 0.5rem;
  display: table;
  font-size: 1rem;
}

.required:after {
  content: "*";
  color: #ff0000;
}

.required-field-msg {
  color: #ff0000;
  margin-bottom: 3rem;
}

.error {
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
  padding: 1rem 3rem;
  border-radius: 10px;
}

.file-name-wrapper {
  border: 1px solid #e2e2e2;
  border-radius: 10px;
  width: auto;
  padding: 1rem 3rem;
  height: 3rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.application-submit-btns {
  display: flex;
  justify-content: center;
}

.loader-div {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
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
  font-size: 1.2em;
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
  padding: 2rem;
  border-radius: 25px;
}

.close-btn {
  float: right;
  padding: 0.5rem 1rem !important;
}

.justify-center {
  justify-content: center;
}

@include media-breakpoint-up(md) {
  .gradient-btn {
    padding: 1rem 5rem;
  }

  .header-text {
    font-size: 2.5rem;
    line-height: 2.729rem;
  }

  .contact-form-text,
  .contact-form-text:hover {
    font-size: 1.125rem;
  }

  .job-application {
    padding: 3rem;
  }
}

@include media-breakpoint-up(lg) {
  .form-container {
    padding: 0 10rem;
  }
}
</style>
