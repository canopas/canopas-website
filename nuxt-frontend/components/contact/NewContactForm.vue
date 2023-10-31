<template>
  <div class="container my-[50px] lg:mt-0 md:w-full">
    <form>
      <div
        class="rounded-2xl border border-slate-200 bg-white shadow-[0px_0px_45px_rgba(0,0,0,0.1)]"
      >
        <div class="pt-5 pb-20 px-8 lg:px-12">
          <div>
            <div class="relative mb-5 md:mb-5 pt-8 text-left">
              <input
                type="text"
                id="username"
                class="block peer mx-0 w-full rounded-none border-b border-slate-400 bg-transparent px-0 transition ease-in-out appearance-none text-[1rem] md:text-[1.375rem] font-inter-regular leading-[1.21rem] md:leading-[2.0625rem] text-black-core placeholder-black-core/[.6] focus:outline-none active:outline-none"
                name="username"
                required
                autocomplete="given-username"
                v-model="name"
                placeholder=" "
                @input="
                  showNameValidationError =
                    $event.target.value.trim().length === 0
                "
                @click.native="$mixpanel.track('tap_contact_name_input')"
              />
              <label
                for="username"
                class="absolute top-4 left-0 mb-5 text-black-core/[.6] font-inter-regular text-[1rem] md:text-[1.375rem] leading-[1.21rem] md:leading-[2.0625rem] transform -translate-y-4 origin-[0] scale-75 duration-300 peer-focus:text-black-core peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-4"
                >Your name</label
              >
              <span v-if="showNameValidationError" class="error text-red-600"
                >Name is required</span
              >
            </div>
            <div class="relative mb-5 md:mb-5 pt-8 text-left">
              <input
                class="block peer mx-0 w-full rounded-none border-b border-slate-400 bg-transparent px-0 transition ease-in-out appearance-none text-[1rem] md:text-[1.375rem] font-inter-regular leading-[1.21rem] md:leading-[2.0625rem] text-black-core placeholder-black-core/[.6] focus:outline-none active:outline-none"
                type="text"
                name="email"
                id="email"
                required
                autocomplete="given-email"
                v-model="email"
                placeholder=" "
                @input="
                  showEmailValidationError =
                    $event.target.value.trim().length === 0
                "
                @blur="showValidEmailError = isValidEmail()"
                @click.native="$mixpanel.track('tap_contact_email_input')"
              />
              <label
                for="email"
                class="absolute top-4 left-0 mb-5 text-black-core/[.6] font-inter-regular text-[1rem] md:text-[1.375rem] leading-[1.21rem] md:leading-[2.0625rem] transform -translate-y-4 origin-[0] scale-75 duration-300 peer-focus:text-black-core/[0.6] peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-4"
                >Your email</label
              >
              <span v-if="showEmailValidationError" class="error text-red-600"
                >Email is required</span
              >
              <span
                v-if="email.trim().length != 0 && showValidEmailError"
                class="error text-red-600"
                >Please enter valid email address</span
              >
            </div>
            <div class="relative md:col-span-2 mb-5 md:mb-5 pt-8 text-left">
              <textarea
                class="block peer mx-0 w-full min-h-[50px] md:min-h-[90px] rounded-none border-b border-slate-400 bg-transparent px-0 transition ease-in-out appearance-none text-[1rem] md:text-[1.375rem] font-inter-regular leading-[1.21rem] md:leading-[2.0625rem] text-black-core placeholder-black-core/[.6] focus:outline-none active:outline-none"
                id="project"
                name="project"
                rows="3"
                required
                autocomplete="given-project-info"
                v-model="projectInfo"
                placeholder=" "
                @input="
                  showProjectInfoValidationError =
                    $event.target.value.trim().length === 0
                "
                @click.native="
                  $mixpanel.track('tap_contact_project_info_input')
                "
              ></textarea>
              <label
                for="project"
                class="absolute top-4 left-0 mb-5 text-black-core/[.6] font-inter-regular text-[1rem] md:text-[1.375rem] leading-[1.21rem] md:leading-[2.0625rem] transform -translate-y-4 origin-[0] scale-75 duration-300 peer-focus:text-black-core peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-4"
                >Tell us about your project</label
              >
              <span
                v-if="showProjectInfoValidationError"
                class="error text-red-600"
                >This field is required</span
              >
            </div>
            <div class="relative md:col-span-2 pt-6 md:pt-8 text-left">
              <input
                class="block peer mx-0 w-full rounded-none border-b border-slate-400 bg-transparent px-0 transition ease-in-out appearance-none text-[1rem] md:text-[1.375rem] font-inter-regular text-black-core placeholder-black-core/[.6] focus:outline-none active:outline-none"
                type="text"
                name="reference"
                id="reference"
                required
                autocomplete="given-reference"
                v-model="reference"
                placeholder=" "
                @input="
                  showReferenceValidationError =
                    $event.target.value.trim().length === 0
                "
                @click.native="$mixpanel.track('tap_contact_reference_input')"
              />
              <label
                for="reference"
                class="absolute top-4 left-0 mb-5 text-black-core/[.6] font-inter-regular text-[1rem] md:text-[1.375rem] leading-[1.21rem] md:leading-[2.0625rem] transform -translate-y-4 origin-[0] scale-75 duration-300 peer-focus:text-black-core peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-4"
                >How did you find us?</label
              >
              <span
                v-if="showReferenceValidationError"
                class="error text-red-600"
                >This field is required</span
              >
            </div>
            <div class="relative pt-8">
              <div ref="invest-list" class="flex">
                <button
                  class="flex items-center justify-between mt-2.5 md:mt-3 mx-0 w-full border-b border-slate-400 bg-none px-0 font-inter-regular text-[1rem] md:text-[1.375rem] leading-[1.21rem] md:leading-[2.0625rem] whitespace-nowrap transition duration-150 ease-in-out focus:outline-0 active:outline-0 focus:shadow-none active:shadow-none focus:ring-0 active:ring-0 focus:bg-transparent active:bg-transparent active:text-black-core/[0.6]"
                  type="button"
                  id="invest"
                  name="invest"
                  required
                  @click="toggleList"
                >
                  <label
                    for="invest"
                    class="absolute top-8 md:top-8 left-0 text-black-core/[0.6] transform duration-300 placeholder:!text-black-core/[0.6]"
                    :class="[
                      floatable
                        ? '-translate-y-4 xl:-translate-y-[0.5rem] origin-[0] scale-75'
                        : 'translate-y-0 scale-100',
                      floatable && showList
                        ? 'text-black-core'
                        : 'text-black-core/[0.6]',
                    ]"
                    @click="floatable = !floatable"
                    >I'll invest</label
                  >
                  <span
                    class="font-inter-regular text-[1rem] md:text-[1.375rem] leading-[1.21rem] md:leading-[2.0625rem]"
                    >{{ invest }}</span
                  >
                  <Icon
                    class="fab mb-3 w-[15px] md:w-[25px] h-[15px] md:h-[25px] text-black-core/[0.6]"
                    :name="
                      showList ? 'fa6-solid:caret-up' : 'fa6-solid:caret-down'
                    "
                  />
                </button>
                <ul
                  v-if="showList"
                  class="absolute mt-9 md:mt-12 w-full min-w-max z-[5] border-none rounded-[10px] shadow-[0px_5px_15px_rgba(0,0,0,0.5)] bg-white py-1 list-none text-left transition-all delay-150 duration-300"
                >
                  <li
                    v-for="option in investOptions"
                    :key="option"
                    @click="setOption(option)"
                  >
                    <span
                      class="block w-full py-2 px-4 whitespace-nowrap text-black-core/[.6] hover:bg-[#000]/[.2] hover:text-[#000] font-inter-regular text-lg md:text-xl lg:text-2xl leading-[1.21rem] md:leading-[2.0625rem]"
                      >{{ option }}</span
                    >
                  </li>
                </ul>
              </div>
              <span v-if="showInvestValidationError" class="error text-red-600"
                >This field is required</span
              >
            </div>
            <div
              class="mt-8 text-base font-inter-regular text-black-core/[0.87]"
            >
              We sign NDA for all of our projects.
            </div>
          </div>
          <div class="flex justify-center py-5 lg:py-8">
            <img
              v-if="showLoader"
              :src="loaderImage"
              class="w-16 h-16"
              alt="loader-image"
            />
            <div v-else>
              <div
                class="relative font-inter-semibold text-[1.1875rem] leading-[1.436875rem] md:leading-[1.386875rem] text-center"
              >
                <div
                  class="absolute -top-10 sm:-top-8 xl:-top-[1.875rem] 2xl:-top-[1.875rem] text-center -right-32 sm:-right-[10.5rem] md:-right-44 lg:-right-40 2xl:-right-[10.5rem] w-[250px] sm:w-max"
                >
                  <span
                    v-if="showErrorMessage"
                    class="flex text-center text-red-600"
                    :class="
                      errorMessage == 'Invalid Recaptcha score'
                        ? '!mt-2 md:!-mt-2 !mr-4 sm:!mr-16 md:!mr-12 lg:!mr-10 xl:!mr-[2.8rem] 2xl:!mr-14 '
                        : ''
                    "
                    >{{ errorMessage }}</span
                  >
                </div>
                <button
                  id="submit"
                  ref="recaptcha"
                  class="absolute top-[-13px] sm:-top-5 right-[-122px] md:right-[-170px] w-max rounded-full py-2.5 md:py-5 px-20 md:px-28 text-center gradient-btn consultation-btn"
                  @click.prevent="submitForm()"
                >
                  <span>Submit</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import axios from "axios";
import config from "@/config.js";
import CalendlyIframe from "./CalendlyIframe.vue";
import loaderImage from "@/assets/images/theme/small-loader.svg";
export default {
  data() {
    return {
      loaderImage,
      showList: false,
      floatable: false,
      name: "",
      email: "",
      projectInfo: "",
      reference: "",
      invest: "",
      investOptions: [
        "< USD 50000",
        "USD 50000 - USD 100000",
        "USD 100000 - USD 250000",
        "> USD 250000",
      ],
      showNameValidationError: false,
      showEmailValidationError: false,
      showValidEmailError: false,
      showProjectInfoValidationError: false,
      showReferenceValidationError: false,
      showInvestValidationError: false,
      openCalendlyIframeModal: false,
      errorMessage: "Something went wrong on our side",
      showLoader: false,
      showErrorMessage: false,
      contactType: 1,
    };
  },
  components: {
    CalendlyIframe,
  },
  inject: ["mixpanel"],
  mounted() {
    document.addEventListener("click", this.closePopUps);
  },
  unmounted() {
    document.removeEventListener("click", this.closePopUps);
  },
  methods: {
    isValidEmail() {
      var emailRegx =
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return !emailRegx.test(this.email);
    },
    validateForm() {
      this.showNameValidationError = this.name.trim().length === 0;
      this.showEmailValidationError = this.email.trim().length === 0;
      this.showProjectInfoValidationError =
        this.projectInfo.trim().length === 0;
      this.showReferenceValidationError = this.reference.trim().length === 0;
      this.showInvestValidationError = this.invest.trim().length === 0;
      return (
        this.showNameValidationError ||
        this.showEmailValidationError ||
        this.showValidEmailError ||
        this.showProjectInfoValidationError ||
        this.showReferenceValidationError ||
        this.showInvestValidationError
      );
    },
    submitForm() {
      if (!this.validateForm()) {
        this.$mixpanel.track("tap_contact_submit_button");
        let formData = {
          name: this.name,
          email: this.email,
          project_info: this.projectInfo
            ? this.projectInfo.replace(/\./g, ".\n")
            : "NA",
          reference: this.reference,
          invest: this.invest,
          send_mail_to_client: config.IS_PROD,
        };
        grecaptcha.enterprise.ready(() => {
          grecaptcha.enterprise
            .execute(config.VITE_RECAPTCHA_SITE_KEY, {
              action: "verify",
            })

            .then((token) => {
              formData.token = token;
              axios
                .post(config.API_BASE + "/api/send-contact-mail", formData)
                .then(() => {
                  this.$router.push({
                    path: "/thank-you",
                  });
                  localStorage.setItem(
                    "client-name",
                    JSON.stringify(formData.name),
                  );
                  this.resetForm();
                })
                .catch((err) => {
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
              this.showErrorMessage = true;
              setTimeout(() => {
                this.showErrorMessage = false;
              }, 3000);
            });
        });
      }
    },
    resetForm() {
      this.name = "";
      this.email = "";
      this.projectInfo = "";
      this.reference = "";
      this.invest = "";
      this.floatable = false;
    },
    closePopUps(e) {
      const showList = this.$refs["invest-list"];
      if (showList && !showList.contains(e.target)) {
        this.showList = false;
        if (this.invest == "") {
          this.floatable = false;
        }
      }
    },
    setOption(option) {
      this.invest = option;
      this.showInvestValidationError = this.invest.trim().length == 0;
      this.showList = false;
    },
    toggleList() {
      this.showList = !this.showList;
      this.floatable = this.showList || this.invest.trim().length > 0;
      this.$mixpanel.track("tap_contact_invest_input");
    },
  },
};
</script>
