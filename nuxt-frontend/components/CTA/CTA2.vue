<template>
  <div
    class="relative -bottom-5 z-[1] overflow-hidden px-0 font-inter-medium sm:-bottom-2 xl:mt-32"
  >
    <div class="absolute h-[8%] w-full bg-white sm:h-[6%] md:hidden" />

    <img
      v-if="width < 600"
      :src="bg400"
      class="absolute -z-[1] w-full"
      alt="contact-footer"
    />

    <img
      v-else
      :src="bg2400"
      class="absolute left-0 top-0 -z-[1] h-full w-full object-cover"
      alt="contact-footer"
    />

    <div
      class="cta-gradient-text mt-20 text-center font-inter-bold text-3xl sm:mt-14 md:text-5xl md:leading-[3.625rem] lg:text-[4.0625rem] lg:leading-[4.9375rem]"
    >
      Say Hello!
    </div>
    <div class="blog-container text-center">
      <div>
        <form method="POST" @submit="submitForm">
          <div class="py-5 lg:px-20 xl:px-44">
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <div
                class="relative pt-3 text-left md:col-span-2 md:mb-5 lg:pt-10"
              >
                <input
                  id="username"
                  v-model="name"
                  type="text"
                  class="floating-input peer mx-0 my-2 block w-full appearance-none rounded-none border-b border-white/[.6] bg-transparent px-0 text-lg text-white placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-xl lg:text-2xl"
                  name="username"
                  required
                  autoComplete="given-username"
                  placeholder=" "
                  @input="
                    (showNameValidationError =
                      $event.target.value.trim().length === 0)
                  "
                />
                <label
                  htmlFor="username"
                  class="absolute left-0 top-4 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-base leading-[1.1875rem] text-white/[.6] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-white md:text-[1.375rem] md:leading-[1.6875rem] lg:text-[1.75rem] lg:leading-[2.125rem]"
                >
                  Your name
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
                  class="floating-input peer mx-0 my-2 block w-full appearance-none rounded-none border-b border-white/[.6] bg-transparent px-0 text-lg text-white placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-xl lg:text-2xl"
                  type="text"
                  name="email"
                  required
                  autoComplete="given-email"
                  placeholder=" "
                  @input="
                    (showEmailValidationError =
                      $event.target.value.trim().length === 0)
                  "
                  @blur="(showValidEmailError = isValidEmail(email))"
                />
                <label
                  htmlFor="email"
                  class="absolute left-0 top-4 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-base leading-[1.1875rem] text-white/[.6] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-white md:text-[1.375rem] md:leading-[1.6875rem] lg:text-[1.75rem] lg:leading-[2.125rem]"
                >
                  Your email
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
                  class="floating-input peer mx-0 my-2 block w-full appearance-none rounded-none border-b border-white/[.6] bg-transparent px-0 text-lg text-white placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-xl lg:text-2xl"
                  name="phonenumber"
                  required
                  autoComplete="given-phonenumber"
                  placeholder=" "
                  @input="
                    (showPhoneNumberValidationError =
                      $event.target.value.trim().length === 0)
                  "
                  @blur="
                    (showValidPhoneNumberError = isValidPhone(phoneNumber))
                  "
                />
                <label
                  htmlFor="phonenumber"
                  class="absolute left-0 top-4 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-base leading-[1.1875rem] text-white/[.6] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-white md:text-[1.375rem] md:leading-[1.6875rem] lg:text-[1.75rem] lg:leading-[2.125rem]"
                >
                  Phone number
                </label>

                <span
                  v-if="showPhoneNumberValidationError"
                  class="error cta-gradient-text"
                >
                  Phone number is required
                </span>

                <span
                  v-if="
                    phoneNumber.trim().length != 0 && showValidPhoneNumberError
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
                  class="floating-input peer mx-0 my-2 block w-full appearance-none rounded-none border-b border-white/[.6] bg-transparent px-0 text-lg text-white placeholder-white/[.6] transition ease-in-out focus:outline-none active:outline-none md:text-xl lg:text-2xl"
                  name="project"
                  rows="3"
                  required
                  autoComplete="given-project-info"
                  placeholder=" "
                />
                <label
                  htmlFor="project"
                  class="absolute left-0 top-4 z-[2] mb-5 origin-[0] -translate-y-4 scale-75 transform text-base leading-[1.1875rem] text-white/[.6] duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-white md:text-[1.375rem] md:leading-[1.6875rem] lg:text-[1.75rem] lg:leading-[2.125rem]"
                >
                  Project detail
                </label>

                <span
                  v-if="showProjectInfoValidationError"
                  class="error cta-gradient-text"
                >
                  This field is required
                </span>
              </div>
            </div>

            <div class="mt-10 flex justify-center md:mt-12">
              <img
                v-if="showLoader"
                :src="loaderImage"
                class="h-16 w-16"
                alt="loader-image"
              />

              <div v-else class="relative">
                <div
                  v-if="showErrorMessage"
                  class="absolute -right-16 -top-8 w-[190%] text-center sm:-right-44 sm:-top-[1.875rem] sm:w-max md:-right-60 md:-top-[2.875rem] lg:-right-72 xl:-right-72 2xl:-right-[20.5rem]"
                >
                  <span
                    class="cta-gradient-text flex text-center"
                    :class="
                      errorMessage === 'Invalid Recaptcha score'
                        ? '!text-[1.5rem] sm:!mr-[7rem] md:!mr-[12rem] lg:!mr-[15rem] xl:!mr-[16rem] 2xl:!mr-[17rem]'
                        : ''
                    "
                  >
                    {{ errorMessage }}
                  </span>
                </div>

                <button
                  v-else
                  id="submit"
                  class="relative justify-self-center rounded-full border border-solid border-transparent bg-gradient-to-r from-[#f2709c] to-[#ff9472] font-inter-bold text-white hover:shadow-[inset_2px_1000px_1px_#fff]"
                >
                  <span
                    class="text-md hoverable-text inline-block px-[1.8rem] py-[0.6rem] md:py-[0.8rem] lg:text-xl"
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
  </div>
</template>

<script setup>
import { onMounted, ref, toRefs } from "vue";
import axios from "axios";
import { isValidEmail, isValidPhone } from "./../../utils";
import bg400 from "../../assets/images/CTA/second-cta-400.svg";
import bg2400 from "../../assets/images/CTA/second-cta-2400.svg";
import loaderImage from "../../assets/images/small-loader.svg";

const width = 680;
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
