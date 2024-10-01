<template>
  <div @click="isShowReferenceOption = false">
    <Header />
    <ScreenLoader v-if="isLoading || job == null" />

    <div v-else>
      <div
        class="container mt-[30px] lg:mt-20 mx-auto mb-[150px] lg:py-0 lg:px-40 overflow-hidden"
      >
        <div
          class="border border-solid border-[#e2e2e2] rounded-[15px] py-8 px-6 md:pt-12 md:px-12 md:pb-0"
        >
          <div class="header-2-text text-center pb-4">
            <h1 class="canopas-gradient-text">Applying For {{ job.title }}</h1>
          </div>

          <form
            class="text-[#3d3d3dcc] text-[1rem] md:text-[1.125rem] pt-12 pb-12"
          >
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="mb-2">
                <label class="after:content-['*'] after:text-[#ff0000]"
                  >Full Name</label
                >
                <input
                  class="block w-full border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[5px] py-2.5 px-4 focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-80 disabled:cursor-not-allowed"
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
                  class="error text-red-600 text-[1rem]"
                  >Name is required</span
                >
              </div>

              <div class="mb-2">
                <label class="after:content-['*'] after:text-[#ff0000]"
                  >Email</label
                >
                <input
                  class="block w-full border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[5px] py-[10px] px-4 focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-80 disabled:cursor-not-allowed"
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
                  class="error text-red-600 text-[1rem]"
                  >Email is required</span
                >
                <span
                  v-if="email.trim().length != 0 && showValidEmailError"
                  class="error text-red-600 text-[1rem]"
                  >Please enter valid email address</span
                >
              </div>

              <div class="mb-[0.5rem]">
                <label class="after:content-['*'] after:text-[#ff0000]"
                  >Phone Number</label
                >
                <input
                  class="block w-full border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[5px] py-2.5 px-4 focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-80 disabled:cursor-not-allowed"
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
                  class="error text-red-600 text-[1rem]"
                  >Phone number is required</span
                >
                <span
                  v-if="phoneNumber.trim().length != 0 && showValidPhoneError"
                  class="error text-red-600 text-[1rem]"
                  >Please enter valid phone number</span
                >
              </div>

              <div class="mb-2">
                <label>City</label>
                <input
                  class="block w-full border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[5px] py-2.5 px-4 focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-[0.8] disabled:cursor-not-allowed"
                  type="text"
                  name="city"
                  autocomplete="address-level2"
                  v-model="city"
                  aria-label="City name"
                />
              </div>

              <div class="mb-2">
                <label>How did you find Canopas?</label>
                <div
                  class="relative mt-[5px] w-full h-[50px] outline-hidden outline-0 before:content-[''] before:absolute before:top-5 before:right-5 before:z-[1] before:w-2 before:h-2 before:border-2 before:border-solid before:border-[#00000080] before:border-t-2 before:border-t-[#fff] before:border-r-2 before:border-r-[#fff] before:rotate-[-45deg] before:duration-[0.5s] before:pointer-events-none"
                  v-on:click.stop="
                    isShowReferenceOption = !isShowReferenceOption
                  "
                  :class="isShowReferenceOption ? 'active' : ''"
                >
                  <div>
                    <input
                      class="absolute top-0 left-0 w-full h-full cursor-pointer bg-[#fff] border border-solid border-[#e2e2e2] outline-hidden outline-0 py-3 px-4 rounded-[10px]"
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
                    class="absolute top-[50px] z-[1] w-full bg-[#fff] overflow-hidden hidden shadow-[0_5px_20px_rgba(0,0,0,0.1)] rounded-[10px] border border-solid border-[#0000000d]"
                  >
                    <div
                      class="py-3 px-3 cursor-pointer hover:bg-[#e2e2e280]"
                      v-for="refer in references"
                      :key="refer"
                      @click="showOptions(refer)"
                    >
                      {{ refer.option }}
                    </div>
                  </div>
                </div>
              </div>

              <div v-if="isShowingReferenceInput" class="mb-8 reference-input">
                <label>Reference</label>
                <input
                  class="block w-full border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[5px] py-2.5 px-4 focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-80 disabled:cursor-not-allowed reference-by-input"
                  type="text"
                  name="referenceby"
                  autocomplete="given-reference-name"
                  :placeholder="references[currentReferenceIndex].hint"
                  v-model="referenceBy"
                  aria-label="reference by"
                />
              </div>

              <div class="md:col-span-2 mb-2">
                <label
                  >Message
                  <textarea
                    class="block w-full min-h-[100px] border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[5px] py-2.5 px-4 focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-80 disabled:cursor-not-allowed hidden-scrollbar whitespace-pre-wrap"
                    name="message"
                    rows="4"
                    v-model="message"
                    @input="resizeTextarea"
                    placeholder="Tell us about yourself..."
                  ></textarea>
                </label>
              </div>

              <div>
                <label class="after:content-['*'] after:text-[#ff0000]"
                  >Resume
                </label>
                <br />
                <label class="text-[#999]">
                  <i>Supported formats:</i
                  ><span class="black pl-2">PDF and Docs </span>
                  <i>only</i>
                </label>
              </div>
              <div class="flex">
                <button
                  type="button"
                  class="resume-upload-btns bg-[#fff] text-[#3d3d3d] truncate border break-normal border-solid border-[#e2e2e2] shadow-none outline-hidden outline-0 py-2.5 px-12 rounded-[10px] h-[50px] focus:outline-hidden focus:outline-0"
                  @click="chooseFiles()"
                >
                  {{ fileButtonName }}
                </button>
                <input
                  id="fileUpload"
                  ref="fileUpload"
                  type="file"
                  class="block border border-solid border-[#e2e2e2] rounded-[10px] text-[#3d3d3d] text-[1.125rem] mt-[15px] focus:border focus:border-solid focus:border-[#e2e2e2] focus:outline-hidden focus:outline-0 disabled:opacity-80 disabled:cursor-not-allowed custom-file-input w-0 opacity-0 p-0 pr-2.5"
                  name="resume"
                  accept="application/pdf,.doc,.docx"
                  @change="previewFiles"
                  required
                  aria-label="file upload"
                />
                <span
                  v-if="showFileValidationError"
                  class="mt-2.5 error text-red-600 text-[1rem]"
                  >Resume is required</span
                >
                <span
                  v-if="showFileSizeValidationError"
                  class="mt-2.5 error text-red-600 text-[1rem]"
                  >File should have size less then 1MB</span
                >
              </div>
            </div>
            <div class="flex items-center mt-8">
              <label
                class="relative block mb-[15px] ml-2 pl-[35px] text-[1.125rem] leading-[1.688rem] font-inter-medium text-black-core/[0.87] box cursor-pointer"
                for="checkbox"
                >I acknowledge that Canopas does not conduct remote interviews
                and does not offer remote job positions.
                <input type="checkbox" id="checkbox" checked />
                <span
                  class="absolute top-0 left-0 mt-[5px] h-6 w-6 border-2 border-solid border-[#999999] rounded after:absolute after:hidden after:content-[''] mark after:left-[9px] after:bottom-[7px] after:w-1.5 after:h-3.5 after:border-solid after:border-[#fff] after:border-t-0 after:border-r-[3px] after:border-b-[3px] after:border-l-0 after:rotate-[45deg]"
                ></span>
              </label>
            </div>
            <div class="flex justify-center mt-4">
              <img
                class="mt-5 mx-auto mb-0 py-0 px-16 md:px-20 lg:mt-5 lg:mx-2 lg:mb-0"
                :src="loaderImage"
                v-if="showLoader"
                alt="loader-image"
              />
              <div class="relative flex flex-col items-center" v-else>
                <div
                  class="absolute text-center -top-10 sm:-top-[1.875rem] right-0 sm:-right-28 lg:-right-24 xl:-right-40 2xl:-right-[13.5rem] sm:w-max"
                >
                  <span
                    v-if="showErrorMessage"
                    class="flex mr-0 sm:mr-10 md:mr-14 xl:mr-[8.5rem] 2xl:mr-[11.5rem] text-center text-red-600 text-[1.3rem] sm:text-[1.5rem]"
                    :class="
                      errorMessage == 'Invalid Recaptcha score'
                        ? '!text-[1.5rem] leading-8 sm:!mr-[6.5rem] md:!mr-32 lg:!mr-[7.5rem] xl:!mr-48 2xl:!mr-60'
                        : ''
                    "
                    >{{ errorMessage }}</span
                  >
                </div>
                <button
                  class="gradient-btn !rounded-full px-16 md:px-20"
                  @click.prevent="submitApplication()"
                >
                  <Icon
                    class="fa"
                    name="fa6-regular:circle-check"
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

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import axios from "axios";
import config from "@/config.js";
import { useJobDetailStore } from "@/stores/jobs";

import loaderImage from "@/assets/images/theme/small-loader.svg";
import { useRoute } from "vue-router";

const { $mixpanel } = useNuxtApp();

const route = useRoute();
const id = route.params.id;
const href = route.href;

const references = [
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
];

const ONE_MB = 1048576
const event = ref("");
const currentReferenceIndex = ref(-1);
const fullName = ref("");
const email = ref("");
const phoneNumber = ref("");
const city = ref("");
const reference = ref("");
const referenceBy = ref("");
const message = ref("");
const file = ref("");
const isLoad = ref(false);
const isShowingReferenceInput = ref(false);
const showNameValidationError = ref(false);
const showValidEmailError = ref(false);
const showEmailValidationError = ref(false);
const showPhoneValidationError = ref(false);
const showValidPhoneError = ref(false);
const showFileValidationError = ref(false);
const showFileSizeValidationError = ref(false);
const showErrorMessage = ref(false);
const showLoader = ref(false);
const isShowReferenceOption = ref(false);
const fileButtonName = ref("Upload");
const errorMessage = ref("Something went wrong on our side");
const fileUpload = ref(null);

const store = useJobDetailStore();
const job = computed(() => store.item);
const jobsError = computed(() => store.error);
const isLoading = computed(() => store.isLoading);

await setCareerDetails();

useSeoMeta({
  title: job.value?.apply_seo_title,
  description: job.value?.apply_seo_description,
  ogTitle: job.value?.apply_seo_title,
  ogType: "Jobs Posting Website",
  ogImage: config.OG_IMAGE_URL,
  ogUrl: config.BASE_URL + route.href,
});

onMounted(() => {
  let recaptchaScript = document.createElement("script");
  recaptchaScript.setAttribute(
    "src",
    "https://www.google.com/recaptcha/enterprise.js?render=" +
      config.VITE_RECAPTCHA_SITE_KEY,
  );
  recaptchaScript.setAttribute("async", "true");
  recaptchaScript.setAttribute("defer", "true");
  document.head.appendChild(recaptchaScript);
  $mixpanel.track("view_apply_job_page");
});

function resizeTextarea(event) {
  event.target.style.height = "auto";
  event.target.style.height = event.target.scrollHeight + "px";
}

async function setCareerDetails() {
  await useAsyncData("jobs", () => store.loadJob(id, href));

  if (jobsError.value != null) {
    let err = jobsError.vale;
    if (err && err.response && err.response.status == 404) {
      navigateTo({
        name: "Error404Page",
        params: { pathMatch: ["jobs", "apply", id] },
      });
    } else {
      showErrorMessage.value = true;
      $mixpanel.track("job_apply_failed");
    }
  }
}

function showFileErrors() {
  showFileValidationError.value = fileUpload.value === '';
  showFileSizeValidationError.value = file.value.size > ONE_MB;
}

function showOptions(refer) {
  reference.value = refer.option;
  let names = ["Canopas Employee", "Job posting website", "Other"];
  if (names.includes(refer.option)) {
    isShowingReferenceInput.value = true;
    currentReferenceIndex.value = refer.id - 1;
  } else {
    isShowingReferenceInput.value = false;
  }
}
function previewFiles(event) {
  file.value = event.target.files[0];
  if (file.value.name != null) {
    fileButtonName.value = file.value.name;

    fileButtonName.value.length > 10
      ? fileButtonName.value.substr(0, 20) + "..."
      : fileButtonName.value;
  }
  showFileErrors();
}
function chooseFiles() {
  document.getElementById("fileUpload").click();
}
function isValidPhone() {
  const phoneRegex = /^[0-9]{10}$/;
  return !phoneRegex.test(phoneNumber.value);
}

function isValidEmail() {
  const emailRegx =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return !emailRegx.test(email.value);
}

function validateForm() {
  showNameValidationError.value = fullName.value.trim().length === 0;
  showEmailValidationError.value = email.value.trim().length === 0;
  showPhoneValidationError.value = phoneNumber.value.trim().length === 0;
  showFileValidationError.value = fileButtonName.value === "Upload";
  showFileSizeValidationError.value = file.value.size > ONE_MB;
  return (
    showNameValidationError.value ||
    showPhoneValidationError.value ||
    showFileValidationError.value || 
    showFileSizeValidationError.value ||
    showEmailValidationError.value ||
    showValidPhoneError.value ||
    showValidEmailError.value 
  );
}

function submitApplication() {
  $mixpanel.track("job_submit");
  if (!validateForm()) {
    showLoader.value = true;
    isLoad.value = true;

    //resume file name with the current date
    const year = new Date().getFullYear();
    const mon = new Date().toLocaleString("default", { month: "short" });
    const date = new Date().getDate();

    const currentDate = date + "-" + mon + "-" + year;

    const splitFileName = file.value.name.split(".");

    const fileName =
      splitFileName[0] +
      "-" +
      currentDate +
      "." +
      splitFileName[splitFileName.length - 1];

    //prepare form data
    const formData = new FormData();
    formData.append("name", fullName.value);
    formData.append("email", email.value);
    formData.append("phone", phoneNumber.value);
    formData.append("place", city.value ? city.value : "NA");
    formData.append(
      "references",
      reference.value &&
        reference.value != "" &&
        referenceBy.value &&
        referenceBy.value != ""
        ? reference.value + " - " + referenceBy.value
        : reference.value && reference.value != ""
          ? reference.value
          : "NA",
    );

    formData.append("file", file.value, fileName);
    formData.append("job_title", job.value.title);
    formData.append(
      "message",
      message.value
        ? message.value.replace(/\n/g, "<br>\n").replace(/\./g, ".\n")
        : "NA",
    );
    formData.append("save_record_to_spreadsheet", config.IS_PROD);
    formData.append("save_data_to_database", config.IS_PROD);
    localStorage.setItem("applicant-name", JSON.stringify(fullName.value));

    // verify recpatcha
    grecaptcha.enterprise.ready(() => {
      grecaptcha.enterprise
        .execute(config.VITE_RECAPTCHA_SITE_KEY, {
          action: "verify",
        })
        .then((token) => {
          formData.append("token", token);
          axios
            .post(config.API_BASE + "/api/send-jobs-applications", formData)
            .then(() => {
              isLoad.value = false;
              showLoader.value = false;
              navigateTo({
                path: `/jobs/thank-you`,
              });

              localStorage.setItem(
                "applicant-name",
                JSON.stringify(fullName.value),
              );
              fullName.value = "";
              email.value = "";
              phoneNumber.value = "";
              city.value = "";
              reference.value = "";
              referenceBy.value = "";
              message.value = "";
              fileUpload.value = null;
              fileButtonName.value = "Upload";
              $mixpanel.track("job_apply_success");
            })
            .catch((err) => {
              isLoad.value = false;
              showLoader.value = false;
              if (err.response.status == 401) {
                errorMessage.value = "Invalid recaptcha score";
                showErrorMessage.value = true;
                setTimeout(() => {
                  showErrorMessage.value = false;
                }, 3000);
              }
            });
        })
        .catch(() => {
          errorMessage.value = "Invalid recaptcha score";
          isLoad.value = false;
          showLoader.value = false;
          showErrorMessage.value = true;
          setTimeout(() => {
            showErrorMessage.value = false;
          }, 3000);
        });
    });
  }
}
</script>

<style lang="postcss" scoped>
.active::before {
  @apply top-[22px] -rotate-[225deg];
}

input[type="checkbox"] {
  @apply invisible;
}
.box input:checked + .mark {
  @apply bg-[#f2709c] border-0;
}

.box input:checked + .mark:after {
  @apply block;
}
</style>
