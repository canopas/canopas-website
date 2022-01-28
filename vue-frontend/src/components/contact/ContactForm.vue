<template>
  <div class="apply-information">
    <form>
      <div class="required-field-msg">
        All fields marked with * are required.
      </div>
      <div class="information-detail">
        <div class="information-pd">
          <div class="form-group form-inline">
            <label for="username" class="contact-form-text required"
              >Hi, Canopas team. My name is</label
            >
            <label>
              <input
                class="form-control custom"
                type="text"
                name="username"
                required
                autocomplete="given-name"
                v-model="name"
              />
              <span v-if="showValidationError" class="error"
                >This field is required</span
              >
            </label>
          </div>
        </div>

        <hr class="hr-margin" />

        <div class="information-pd">
          <div class="radio-buttons mb-3">
            <label class="label1">
              <input
                type="radio"
                name="designation"
                value="0"
                v-model.number="designationType"
              />
              <div class="box">
                <span class="contact-form-text"
                  >I am individual entrepreneur running my own business.</span
                >
              </div>
            </label>
            <label class="label2">
              <input
                type="radio"
                name="designation"
                value="1"
                v-model.number="designationType"
              />
              <div class="box">
                <span class="contact-form-text"
                  >I work for the company and I am approaching on behalf of the
                  company.</span
                >
              </div>
            </label>
          </div>

          <div class="mb-4">
            <div
              v-if="designationType == 0"
              class="apply-font-color label1-info"
            >
              <span class="contact-form-text">
                I am an owner of the business and I run
                <label>
                  <input
                    class="form-control custom input-text-width"
                    type="text"
                    name="company"
                    v-model="business"
                  />
                </label>
                business. For more information, you can visit my website /
                facebook page / Instagram etc. (Please attach relevant links
                below.)
              </span>
            </div>
            <div
              v-if="designationType == 1"
              class="apply-font-color label2-info"
            >
              <span class="contact-form-text"
                >I am representing
                <label>
                  <input
                    class="form-control custom input-text-width"
                    type="text"
                    name="company"
                    v-model="business"
                  />
                </label>
                company. You can find more information about our business by
                visiting our website that I mentioned below.(Please attach all
                relevant links below.)</span
              >
            </div>
          </div>

          <div class="social-media">
            <div
              v-for="(website, i) in websites"
              class="btn btn-social-media mb-3"
              :key="website"
              :class="{
                'item-selected': i == currentWebsiteIndex || website.url != '',
              }"
              @click="onWebsiteClick(i)"
            >
              <span>
                <font-awesome-icon
                  :icon="website.url == '' ? plusIcon : website.icon"
                  class="fas me-1"
                  :class="{
                    'facebook-twitter-icon':
                      (i == 1 || i == 2) && website.url != '',
                    'website-instagram-icon':
                      (i == 0 || i == 3) && website.url != '',
                  }"
                />
                <div class="contact-form-text">
                  <div id="span-web-text" class="ellipsis-social-media-url">
                    {{ website.url == "" ? website.title : website.url }}
                  </div>
                </div>
                <font-awesome-icon
                  v-if="website.url != ''"
                  :icon="removeIcon"
                  class="remove-icon ms-1"
                  @click="onRemoveWebsiteClick(i)"
                />
              </span>
            </div>
          </div>

          <div v-if="currentWebsiteIndex != -1" class="mb-3 add-url url-text">
            <input
              class="form-control custom social-media-text url-width"
              type="url"
              name="contactURL"
              :placeholder="websites[currentWebsiteIndex].hint"
              v-model="currentWebsiteUrl"
            />
            <font-awesome-icon
              class="fas set-url-icon"
              :icon="arrowRightIcon"
              @click="onAddWebsiteClick"
            />
          </div>
          <div class="url-validation-msg">Please enter valid URL</div>
        </div>

        <hr class="hr-margin" />

        <div class="information-pd">
          <div class="contact-form-text contact-text-black">
            I am here because I need help with...
          </div>
          <br />
          <div class="radio-buttons">
            <label
              v-for="(help, i) in helps"
              :key="help"
              class="reason-radio-div"
            >
              <input
                type="radio"
                name="idea"
                :value="i"
                v-model.number="helpType"
              />
              <div class="box">
                <span class="contact-form-text">{{ help.title }}</span>
              </div>
            </label>
          </div>
        </div>

        <div v-if="helpType != -1" class="information-pd reason-info">
          <div class="contact-form-text contact-text-black">
            Please select reason...
          </div>
          <br />
          <div class="radio-buttons reason-label for-reason-label-1">
            <label v-for="item in helps[helpType].reasons" :key="item">
              <input
                type="radio"
                name="reason"
                :value="item"
                v-model="reason"
              />
              <div class="box">
                <span class="contact-form-text">{{ item }}</span>
              </div>
            </label>
          </div>
        </div>

        <hr class="hr-margin" />

        <div class="information-pd">
          <div>
            <label class="contact-form-text required">
              You can reach out to me at (email address)
            </label>
            <input
              class="form-control custom url-width"
              type="text"
              name="email"
              required
              v-model="email"
            />
            <span v-if="showValidationError" class="error"
              >This field is required</span
            >
          </div>
          <br />
          <div class="my-3">
            <textarea
              class="input-textarea"
              name="message"
              rows="10"
              :v-bind="message"
              v-model="message"
              placeholder="I have a message or information for Canopas team."
            ></textarea>
          </div>
          <div>
            <div v-if="contactType < 2" class="radio-buttons mb-3">
              <label class="contactLabel1">
                <input
                  type="radio"
                  name="contact_type"
                  value="0"
                  v-model.number="contactType"
                />
                <div class="box">
                  <span class="contact-form-text"
                    >If you love to talk, You can schedule a meeting with our
                    team member according to your convenience from here.</span
                  >
                </div>
              </label>
              <label class="contactLabel2">
                <input
                  type="radio"
                  name="contact_type"
                  value="1"
                  v-model.number="contactType"
                />
                <div class="box">
                  <span class="contact-form-text"
                    >If you prefer chat or email, please click here.</span
                  >
                </div>
              </label>
            </div>
            <button
              type="submit"
              v-if="contactType == 1"
              class="gradient-btn gradient-bg-btn chat-email-btn"
              @click.prevent="submitApplication()"
            >
              <font-awesome-icon
                class="fas"
                :icon="planeIcon"
                aria-hidden="true"
              />
              <span>Submit</span>
            </button>
            <button
              type="submit"
              v-if="contactType == 0"
              class="gradient-btn gradient-bg-btn call-now-btn"
              @click.prevent="submitApplication()"
            >
              <font-awesome-icon
                :icon="calendarIcon"
                class="fas"
                aria-hidden="true"
              />
              <span>Schedule Meeting</span>
            </button>
          </div>
        </div>
      </div>
    </form>
    <!-- Show Calendly Iframe -->
    <div v-if="openCalendlyIframeModal">
      <transition name="modal">
        <div class="modal-mask">
          <div
            class="modal-dialog login-modal modal-xl modal-dialog-centered"
            role="document"
          >
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-body">
                <CalendlyIframe />

                <button
                  type="button"
                  class="close modal-close-btn"
                  data-dismiss="modal"
                  aria-label="Close"
                >
                  <span aria-hidden="true" @click="closeCalendlyIframeModal"
                    >&times;</span
                  >
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
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-body">
                <div class="success-message-div">
                  <div
                    class="success-message-text canopas-gradient-text text-center"
                  >
                    Thank you for choosing us to make a difference in your
                    business.
                  </div>
                  <div class="">
                    If you prefer to chat or email, sit back and relax our team
                    will get back to you within 24 hours.
                  </div>
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
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faPlusCircle,
  faArrowRight,
  faGlobe,
  faTimes,
  faPaperPlane,
  faCalendarAlt,
} from "@fortawesome/free-solid-svg-icons";
import {
  faFacebook,
  faTwitter,
  faInstagram,
} from "@fortawesome/free-brands-svg-icons";
import axios from "axios";
import CalendlyIframe from "./CalendlyIframe.vue";
import config from "@/config.js";

const CONTACT_BY_CHAT_OR_MAIL = 1;

export default {
  data() {
    return {
      plusIcon: faPlusCircle,
      removeIcon: faTimes,
      arrowRightIcon: faArrowRight,
      planeIcon: faPaperPlane,
      calendarIcon: faCalendarAlt,
      name: "",
      designationType: 0,
      websites: [
        {
          title: "website",
          url: "",
          hint: "Business website (e.g) - canopas.com",
          icon: faGlobe,
        },
        {
          title: "facebook",
          url: "",
          hint: "Facebook page of your business",
          icon: faFacebook,
        },
        {
          title: "twitter",
          url: "",
          hint: "Twitter page of your business",
          icon: faTwitter,
        },
        {
          title: "instagram",
          url: "",
          hint: "Instagram page of your business",
          icon: faInstagram,
        },
        {
          title: "other",
          url: "",
          hint: "Additional connection to your business",
          icon: faGlobe,
        },
      ],
      helps: [
        {
          title:
            "I have an idea for my business that I want to implement with your help.",
          reasons: [
            "I have a rough design for the product I want to develop.",
            "I have design or UI/UX ready in figma, sketch, photoshop, or illustrator.",
            "I have a complete understanding of what I want. But, I don’t know anything about these tools.",
          ],
        },
        {
          title:
            "My existing product is poor or I want to add new features or functionalities to it.",
          reasons: [
            "I want to fix bugs in my app reported by users.",
            "My app is broken and I need someone who can make my app work again.",
            "I want to add additional functionalities to the app.",
          ],
        },
        {
          title:
            "I want to validate my IDEA and thus I need MVP(minimum viable product) or PoC(Proof of Concept) for the same.",
          reasons: [
            "I have a great idea that I want to test out by building MVP(Minimum viable product).",
            "I want to develop a PoC(Proof of concept) app so that I can pitch investors my idea.",
            "I have a unique idea and I want to develop a fully fledged product based on it.",
          ],
        },
        {
          title:
            "I want to check the feasibility of my idea and need your consultation.",
          reasons: [
            "I have an idea. But, I’m not sure about the feasibility of my idea.",
            "I’m not a techie and I need consultation and guidance from your team validating my idea.",
          ],
        },
      ],
      helpType: -1,
      reason: "",
      message: "",
      currentWebsiteUrl: "",
      currentWebsiteIndex: -1,
      contactType: -1,
      email: "",
      openCalendlyIframeModal: false,
      showValidationError: false,
      showSuccessMessagePopup: false,
    };
  },
  components: {
    FontAwesomeIcon,
    CalendlyIframe,
  },
  methods: {
    onWebsiteClick(i) {
      if (this.currentWebsiteIndex == i) {
        this.currentWebsiteIndex = -1;
      } else {
        this.currentWebsiteIndex = i;
      }
    },
    onAddWebsiteClick() {
      this.websites[this.currentWebsiteIndex].url = this.currentWebsiteUrl;
      this.currentWebsiteIndex = -1;
      this.currentWebsiteUrl = "";
    },
    onRemoveWebsiteClick(i) {
      this.websites[i].url = "";
      this.currentWebsiteIndex = i;
    },
    submitApplication() {
      let designationValue, designationInfo;

      if (this.name === "" || this.email === "") {
        this.showValidationError = true;
      } else {
        this.showValidationError = false;
        if (this.designationType == 0) {
          designationValue =
            "I am individual entrepreneur running my own business.";
          designationInfo =
            "I am an owner of the business and I run " +
            this.business +
            " business. For more information about my business, you can visit the following links.";
        } else {
          designationValue =
            "I work for the company and I am approaching on behalf of the company.";
          designationInfo =
            "I am representing " +
            this.business +
            " company. You can find more information about our business by visiting the following links.";
        }

        var socialMediaUrls = {};
        this.websites.forEach((website) => {
          socialMediaUrls[website.title] = website.url ? website.url : "NA";
        });

        let formData = {
          name: this.name,
          designation: designationValue ? designationValue : "NA",
          designation_info: designationInfo ? designationInfo : "NA",
          social_media_links: socialMediaUrls,
          idea: this.helpType != -1 ? this.helps[this.helpType].title : "NA",
          reason: this.reason ? this.reason : "NA",
          email: this.email,
          message: this.message ? this.message : "NA",
          contact_type:
            this.contactType == CONTACT_BY_CHAT_OR_MAIL
              ? "Chat or Email"
              : "Call",
        };

        axios
          .post(config.API_BASE + "/api/send-contact-mail", formData)
          .then(() => {
            if (this.contactType == CONTACT_BY_CHAT_OR_MAIL) {
              this.showSuccessMessage();
            } else {
              this.openCalendlyIframe();
            }
          })
          .catch(() => {
            // If error is there show model with below message.
            alert("Something went wrong on our side");
          });
      }
    },
    openCalendlyIframe() {
      this.openCalendlyIframeModal = true;
    },
    closeCalendlyIframeModal() {
      this.openCalendlyIframeModal = false;
    },
    showSuccessMessage() {
      this.showSuccessMessagePopup = true;
      setTimeout(() => {
        this.$router.push("/");
      }, 2000);
    },
  },
};
</script>

<style lang="scss" scoped>
.apply-information {
  margin: 48px 10%;
}

.information-detail {
  border-radius: 10px;
  border: 1px solid #e2e2e2;
}

.hr-margin {
  margin: 0 !important;
}

.information-pd {
  padding: 25px 15px;
}

.contact-form-text,
.contact-form-text:hover {
  color: rgba(0, 0, 0, 0.5);
  font-size: 1.125rem;
  font-style: normal;
  font-weight: normal;
}

.contact-form-text > a,
.contact-form-text > a:hover {
  color: rgba(0, 0, 0, 0.5) !important;
  text-decoration: none !important;
}

.contact-text-black {
  color: rgba(61, 61, 61, 1) !important;
}

.form-control.custom {
  border: none;
  border-radius: 0;
  border-bottom: 1px dashed rgba(61, 61, 61, 0.5);
  box-shadow: none;
  padding: 0.5rem 0 0 !important;
  margin-left: 5px;
  margin-right: 5px;
}

.form-control.custom:focus {
  border-bottom: 1px dashed rgba(61, 61, 61, 0.5);
}

.url-width {
  width: 100%;
}

.send-url-arrow {
  margin-left: 10px;
}

.info-block_container {
  text-align: left;
}

.info {
  display: inline-block;
}

.radio-buttons {
  display: flex;
  flex-direction: column;
}

.radio-buttons > label {
  padding: 8px 0;
  width: 100%;
}

.ellipsis-social-media-url {
  max-width: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.facebook-twitter-icon {
  color: #5da9dd;
}

.website-instagram-icon {
  color: #f2709c;
}

input[type="radio"] {
  display: none;
}

input[type="radio"]:checked ~ .box {
  background: rgba(255, 148, 114, 0.05);
  background: -webkit-linear-gradient(
        rgba(255, 250, 248, 255),
        rgba(255, 250, 248, 255)
      )
      padding-box,
    linear-gradient(to bottom, #ff9472, #f2709c) border-box;
  border: 2px solid transparent;
  border-radius: 15px;
}

.box > span:before {
  opacity: 1;
}

.box {
  display: flex;
  align-items: center;
  text-align: left;
  cursor: pointer;
  position: relative;
  width: 100%;
  padding: 1.5rem;
  background: rgba(226, 226, 226, 0.5);
  border-radius: 15px;
}

.box > span,
.box > span :hover {
  transition: all 300ms ease;
  user-select: none;
  color: rgba(61, 61, 61, 1) !important;
}

.social-media {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}

.input-textarea {
  width: 100%;
  height: 150px !important;
  box-sizing: border-box;
  box-shadow: none;
  color: #3d3d3d;
  font-size: 1rem;
  border-radius: 10px;
  border: 1px solid rgba(61, 61, 61, 0.1);
  padding: 0.5rem 1rem !important;
}

.input-textarea:focus {
  border: 1px solid rgba(61, 61, 61, 0.1);
  outline: none;
}

.btn-social-media > span {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.fas {
  font-size: 1.1rem;
}

.remove-icon {
  font-size: 1rem;
  color: rgba(0, 0, 0, 0.4);
  margin-top: 2px;
  transition: transform 0.2s;
}

.remove-icon:hover {
  transform: scale(1.2);
  color: rgba(0, 0, 0, 0.5);
}

.item-selected {
  border: 1px solid #ff9472 !important;
}

.call-now-btn,
.chat-email-btn {
  width: 100%;
  margin-left: 0 !important;
}

.call-now-btn > i,
.chat-email-btn > i {
  margin-right: 0.5rem;
}

.btn-social-media {
  display: flex;
  align-items: center;
  margin-right: 12px;
  border-radius: 20px;
  color: rgba(0, 0, 0, 0.5) !important;
  border: 1px solid #e2e2e2;
  position: relative;
  transition: 0.5s;
}

.input-text-width {
  width: 10rem;
}

.modal-middle-div {
  display: flex;
  flex-direction: row;
}

.custom-fa-icon {
  background: -moz-linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  background: -webkit-linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  background: linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  -webkit-background-clip: text;
  -moz-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.add-url {
  display: flex;
  flex-direction: row;
  align-items: center;
}

input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
  transition: background-color 5000s !important;
}

.contact-list {
  list-style-type: disc;
}

.contact-list > li {
  display: list-item !important;
}

.contact-btn {
  width: auto !important;
  padding: 0.5rem 1rem !important;
  justify-content: center !important;
  display: inline !important;
  margin-left: 0 !important;
}

.set-url-icon {
  font-size: 1rem;
  color: rgba(0, 0, 0, 0.5);
  cursor: pointer;
}

.ml-10 {
  margin-left: 10px;
  margin-right: 0 !important;
}

.meeting-btn {
  width: auto !important;
  cursor: pointer;
}

.calendly-iframe {
  width: 100% !important;
  height: 80vh !important;
}

.swal-wide {
  width: 80vw !important;
  overflow: hidden;
}

.swal2-container {
  z-index: 100000 !important;
}

.error {
  margin-top: 0.5rem;
  display: table;
  font-size: 1rem;
}

.required:after {
  content: "*";
}

.required-field-msg {
  margin-bottom: 1rem;
}

.required-field-msg,
.required:after,
.error {
  color: #ff0000;
}

.url-validation-msg {
  color: #ff0000;
  display: none;
}

.apply-navbar {
  position: fixed;
}

.contact-container {
  flex-wrap: initial !important;
}

.gradient-btn {
  border-radius: 0.6rem;
  padding: 1rem 1.25rem;
  margin: 1.25rem auto 0 auto;
  text-align: center;
  color: white;
  background: linear-gradient(91.53deg, #f2709c 3.91%, #ff9472 100%);
  border: 1px solid transparent;
}

.gradient-btn > span {
  margin: 0 10px;
  font-weight: 700;
  font-size: 1.1rem;
}

// calendly iframe

.calendly-iframe-modal-content {
  border-radius: 25px !important;
}

.calendly-iframe-modal-content > .modal-header {
  border: none;
}

.calendly-iframe-modal-content > .modal-header > .close {
  font-weight: 300;
  font-size: 40px;
  margin-right: -10px;
}

.calendly-iframe-modal-content > .modal-header > .close:focus {
  outline: none !important;
}

.modal-close-btn {
  border: none;
  color: #f2709c;
  font-size: 40px;
  font-weight: lighter;
  background-color: transparent;
  position: absolute;
  right: 10px;
  bottom: 516px;
}

.modal-close-btn:focus {
  outline: none;
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

// success/thank-you message

.success-message-div {
  padding: 20px;
}

.success-message-text {
  line-height: 2rem;
  font-size: 1.875rem;
  margin-bottom: 30px;
}

@include media-breakpoint-up(md) {
  .information-pd {
    padding: 25px 40px;
  }

  .call-now-btn,
  .chat-email-btn {
    width: 40%;
  }

  .modal-close-btn {
    bottom: 680px;
  }
}

@include media-breakpoint-up(lg) {
  .apply-information {
    margin: 48px 18%;
  }

  .fas {
    font-size: 1.4rem;
  }

  .remove-icon {
    font-size: 1.2rem;
  }

  .url-width {
    width: 50%;
  }

  .radio-buttons {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    justify-content: center;
    flex-flow: wrap;
  }

  .radio-buttons > label {
    padding: 0.5rem 16px 0.5rem 0;
  }

  .box {
    height: 9rem;
    padding: 0 1.5rem;
  }

  .contact-form-text,
  .contact-form-text:hover {
    line-height: 1.875rem;
  }
}

@include media-breakpoint-up(xl) {
  .apply-information {
    margin: 48px 22%;
  }
}
</style>
