<template>
  <div>
    <hr class="md:hidden" />
    <div class="relative top-5 h-auto w-full md:top-0.5">
      <div
        class="blog-container flex flex-col md:-mt-[60px] md:flex-row md:space-x-6 lg:space-x-0 xl:px-20 2x:px-32"
      >
        <div class="mt-10 basis-7/12 md:mt-40 lg:basis-3/5">
          <div
            class="flex flex-col items-center space-y-4 md:w-[16.2rem] md:items-stretch lg:w-80 2xl:w-[22.23rem]"
          >
            <div
              class="font-poppins-regular text-[1.2rem] leading-7 lg:text-2xl 2xl:text-[1.75rem] 2xl:leading-[2.3rem]"
            >
              Talk to an expert
            </div>
            <div
              class="grid-rows-2 justify-items-stretch font-poppins-medium text-[2.3rem] uppercase sm:text-5xl md:grid lg:text-[3.4rem] lg:text-[3.7rem] 2xl:text-[4.125rem]"
            >
              <span class="mr-3 md:mr-0">get in</span>
              <span class="md:justify-self-end">touch</span>
            </div>
            <div
              class="text-center font-poppins-regular text-[0.9rem] leading-5 md:text-left lg:text-base 2xl:text-xl"
            >
              Our team is happy to answer your questions. Fill out the form and
              we’ll get back to you as soon as possible
            </div>
          </div>
        </div>
        <div
          class="cta-box-shadow relative top-14 h-fit w-full bg-white text-center font-poppins-regular lg:top-20 xl:top-[5.5rem]"
        >
          <form method="POST" @submit="submitForm">
            <div class="px-8 pb-10 pt-[3rem] md:pt-20 lg:pb-14 lg:pt-24">
              <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
                <div
                  class="relative pt-3 text-left md:col-span-2 md:mb-5 lg:pt-10"
                >
                  <input
                    id="username"
                    v-model="name"
                    type="text"
                    class="text-md peer mx-0 mb-2 block w-full appearance-none rounded-none border-b border-black-900 bg-transparent px-0 text-black-900 placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-lg lg:text-xl"
                    name="username"
                    required
                    autoComplete="given-username"
                    placeholder=" "
                    @input="
                      showNameValidationError =
                        $event.target.value.trim().length === 0
                    "
                  />
                  <label
                    htmlFor="username"
                    class="absolute left-0 top-1 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-[1.07rem] leading-[1.1875rem] text-black-900/[.8] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-black-900 md:leading-[1.6875rem] lg:top-[1.6rem] lg:text-[1.35rem]"
                  >
                    Name*
                  </label>

                  <span
                    v-if="showNameValidationError"
                    class="error cta-gradient-text"
                  >
                    Name is required
                  </span>
                </div>

                <div class="relative pt-3 text-left md:mb-5 lg:pt-9">
                  <input
                    id="email"
                    v-model="email"
                    class="text-md peer mx-0 mb-2 block w-full appearance-none rounded-none border-b border-black-900 bg-transparent px-0 text-black-900 placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-lg lg:text-xl"
                    type="text"
                    name="email"
                    required
                    autoComplete="given-email"
                    placeholder=" "
                    @input="
                      showEmailValidationError =
                        $event.target.value.trim().length === 0
                    "
                    @blur="showValidEmailError = isValidEmail(email)"
                  />
                  <label
                    htmlFor="email"
                    class="absolute left-0 top-1 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-[1.07rem] leading-[1.1875rem] text-black-900/[.8] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-black-900 md:leading-[1.6875rem] lg:top-[1.6rem] lg:text-[1.35rem]"
                  >
                    Email*
                  </label>

                  <span
                    v-if="showEmailValidationError"
                    class="error cta-gradient-text"
                  >
                    Email is required
                  </span>

                  <span
                    v-if="email.trim().length != 0 && showValidEmailError"
                    class="error cta-gradient-text"
                  >
                    Please enter valid email address
                  </span>
                </div>
                <div class="relative pt-3 text-left md:mb-5 lg:pt-9">
                  <input
                    id="phonenumber"
                    v-model="phoneNumber"
                    type="text"
                    class="text-md peer mx-0 mb-2 block w-full appearance-none rounded-none border-b border-black-900 bg-transparent px-0 text-black-900 placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-lg lg:text-xl"
                    name="phonenumber"
                    required
                    autoComplete="given-phonenumber"
                    placeholder=" "
                    @input="
                      showPhoneNumberValidationError =
                        $event.target.value.trim().length === 0
                    "
                    @blur="
                      showValidPhoneNumberError = isValidPhone(phoneNumber)
                    "
                  />
                  <label
                    htmlFor="phonenumber"
                    class="absolute left-0 top-1 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-[1.07rem] leading-[1.1875rem] text-black-900/[.8] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-black-900 md:leading-[1.6875rem] lg:top-[1.6rem] lg:text-[1.35rem]"
                  >
                    Phone number*
                  </label>

                  <span
                    v-if="showPhoneNumberValidationError"
                    class="error cta-gradient-text"
                  >
                    Phone number is required
                  </span>

                  <span
                    v-if="
                      phoneNumber.trim().length != 0 &&
                      showValidPhoneNumberError
                    "
                    class="error cta-gradient-text"
                  >
                    Please enter valid Phone number
                  </span>
                </div>
                <div
                  class="relative pt-3 text-left md:col-span-2 md:mb-5 lg:pt-10"
                >
                  <textarea
                    id="project"
                    v-model="projectInfo"
                    class="text-md peer mx-0 mb-2 block w-full appearance-none rounded-none border-b border-black-900 bg-transparent px-0 text-black-900 placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-lg lg:text-xl"
                    name="project"
                    rows="3"
                    required
                    autoComplete="given-project-info"
                    placeholder=" "
                  />
                  <label
                    htmlFor="project"
                    class="absolute left-0 top-1 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-[1.07rem] leading-[1.1875rem] text-black-900/[.8] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-black-900 md:leading-[1.6875rem] lg:top-[1.6rem] lg:text-[1.35rem]"
                  >
                    Project detail*
                  </label>

                  <span
                    v-if="showProjectInfoValidationError"
                    class="error cta-gradient-text"
                  >
                    This field is required
                  </span>
                </div>
              </div>

              <div class="mt-6 flex lg:mt-10">
                <img
                  v-if="showLoader"
                  :src="loaderImage"
                  class="h-16 w-16"
                  alt="loader-image"
                />

                <div v-else class="relative">
                  <div
                    v-if="showErrorMessage"
                    class="sm:-vright-44 absolute -right-16 -top-8 w-[190%] text-center sm:-top-[1.875rem] sm:w-max md:-right-60 md:-top-[2.875rem] lg:-right-72 xl:-right-72 2xl:-right-[20.5rem]"
                  >
                    <span
                      class="cta-gradient-text flex text-center"
                      :class="
                        errorMessage === 'Invalid Recaptcha score'
                          ? '!text-2xl sm:!mr-[7rem] md:!mr-[12rem] lg:!mr-[15rem] xl:!mr-[16rem] 2xl:!mr-[17rem]'
                          : ''
                      "
                    >
                      {{ errorMessage }}
                    </span>
                  </div>

                  <button
                    id="submit"
                    class="border-2 border-solid border-transparent bg-black-900 font-inter-bold text-white hover:border-black-900 hover:bg-white hover:text-black-900"
                  >
                    <span
                      class="inline-block px-4 py-[0.6rem] text-base lg:px-[1.8rem] lg:py-[0.8rem] lg:text-[1.1875rem]"
                    >
                      Get Free Consultation
                    </span>
                  </button>
                </div>
              </div>
            </div>
          </form>
        </div>
      </div>
      <div
        class="h-[84px] w-full bg-black-900 md:h-[70px] lg:h-[83px] xl:h-[94px]"
      />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, toRefs } from "vue";
import axios from "axios";
import { isValidEmail, isValidPhone } from "./../../utils";
import loaderImage from "../../assets/images/small-loader.svg";

const name = ref("");
const email = ref("");
const projectInfo = ref("");
const phoneNumber = ref("");
const showNameValidationError = ref(false);
const showEmailValidationError = ref(false);
const showValidEmailError = ref(false);
const showProjectInfoValidationError = ref(false);
const showPhoneNumberValidationError = ref(false);
const showValidPhoneNumberError = ref(false);
const errorMessage = ref("Something went wrong on our side");
const showLoader = ref(false);
const showErrorMessage = ref(false);

const props = defineProps({
  "recaptcha-key": {
    type: String,
    required: true,
  },
  "contact-api-url": {
    type: String,
    required: true,
  },
});

const { recaptchaKey, contactApiUrl } = toRefs(props);

onMounted(() => {
  const recaptchaScript = document.createElement("script");
  recaptchaScript.setAttribute(
    "src",
    "https://www.google.com/recaptcha/enterprise.js?render=" +
      recaptchaKey.value,
  );
  recaptchaScript.setAttribute("async", "true");
  recaptchaScript.setAttribute("defer", "true");
  document.head.appendChild(recaptchaScript);
});

function validateForm() {
  showNameValidationError.value = name.value.trim().length === 0;
  showEmailValidationError.value = email.value.trim().length === 0;
  showProjectInfoValidationError.value = projectInfo.value.trim().length === 0;
  showPhoneNumberValidationError.value = phoneNumber.value.trim().length === 0;

  return (
    showNameValidationError.value ||
    showEmailValidationError.value ||
    showValidEmailError.value ||
    showProjectInfoValidationError.value ||
    showPhoneNumberValidationError.value ||
    showValidPhoneNumberError.value
  );
}

const submitForm = (event) => {
  event.preventDefault();
  if (!validateForm()) {
    showLoader.value = true;
  }

  grecaptcha.enterprise.ready(() => {
    grecaptcha.enterprise
      .execute(recaptchaKey.value, {
        action: "verify",
      })
      .then((token) => {
        if (!validateForm()) {
          const formData = {
            name: name.value,
            email: email.value,
            project_info: projectInfo.value
              ? projectInfo.value.replace(/\./g, ".\n")
              : "NA",
            phone_number: phoneNumber.value,
            token,
          };
          axios
            .post(contactApiUrl + "/api/send-contact-mail", formData)
            .then(() => {
              localStorage.setItem(
                "client-name",
                JSON.stringify(formData.name),
              );
              window.location.href = "/thank-you";
              resetForm();
            })
            .catch((err) => {
              if (err.response.status === 401) {
                errorMessage.value = "Invalid recaptcha score";
                showErrorMessage.value = true;
                setTimeout(() => {
                  showErrorMessage.value = false;
                }, 3000);
              }
            })
            .finally(() => {
              showLoader.value = false;
            });
        }
      })
      .catch(() => {
        errorMessage.value = "Invalid recaptcha score";
        showErrorMessage.value = true;
        setTimeout(() => {
          showErrorMessage.value = false;
        }, 3000);
      });
  });
};

const resetForm = () => {
  name.value = "";
  email.value = "";
  projectInfo.value = "";
  phoneNumber.value = "";
};
</script>
