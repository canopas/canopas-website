<template>
  <div class="apply-information">
    <form action="" method="post" id="contactForm">
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

          <div class="social-media mb-4">
            <div
              v-for="(website, i) in websites"
              class="btn btn-social-media"
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
                />
                <div class="contact-form-text">
                  <span id="span-web-text">{{
                    website.url == "" ? website.title : website.url
                  }}</span>
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
              <input type="radio" name="reason" v-bind="reason" :value="item" />
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
            />
          </div>
          <br />
          <div class="my-3">
            <textarea
              class="input-textarea"
              name="message"
              rows="10"
              :v-bind="message"
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
              onclick="submitApplication()"
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
              onclick="submitApplication()"
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
    };
  },
  components: {
    FontAwesomeIcon,
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
  },
};
</script>

<style scoped>
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
  padding: 25px 40px;
}

.contact-form-text,
.contact-form-text:hover {
  color: rgba(0, 0, 0, 0.5);
  font-size: 1.125rem;
  font-style: normal;
  font-weight: normal;
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
  width: 15rem;
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

.twitter-custom {
  color: #5da9dd !important;
  background: transparent;
  border-radius: 25px !important;
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
  display: flex;
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

@media (min-width: 992px) {
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
    height: 7rem;
    padding: 0 1.5rem;
  }

  .contact-form-text,
  .contact-form-text:hover {
    line-height: 1.875rem;
  }
}

@media (min-width: 1280px) {
  .apply-information {
    margin: 48px 22%;
  }
}
</style>
