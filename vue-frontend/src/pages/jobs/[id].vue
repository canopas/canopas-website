<template>
  <div class="container-fluid">
    <link
      rel="stylesheet"
      href="https://use.fontawesome.com/releases/v5.2.0/css/all.css"
    />
    <ScreenHeader />
    <ScreenLoader v-if="isLoading" />
    <div v-else-if="showErrorMessagePopup">
      <transition name="modal">
        <div class="modal-mask">
          <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-body">
                <div class="error-message-div">
                  <div class="error-message-text text-center">
                    Something went wrong on our side
                  </div>
                  <div class="close-btn-div">
                    <button
                      type="submit"
                      class="gradient-btn close-btn"
                      @click.prevent="closeErrorMessageModal()"
                    >
                      <span>Close</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
    <div v-else>
      <div class="container">
        <div class="canopas-gradient-text text-center">
          {{ job.title }}
        </div>
        <hr class="title-hr mt-4" />
        <div class="normal-text summary-text mt-5">
          {{ job.summary }}
        </div>
        <div class="mt-5">
          <div
            id="description"
            class="normal-text"
            v-html="job.description"
          ></div>
        </div>
        <div class="application-submit-btns mt-5">
          <a class="gradient-btn" :href="jobLink">
            <font-awesome-icon
              class="fa icon"
              :icon="checkCircle"
              aria-hidden="true"
            />
            <span>Apply Now</span>
          </a>
        </div>
      </div>
      <ScreenFooter2 />
    </div>
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeader.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { mapGetters } from "vuex";
import store from "@/store";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import moment from "moment";

import {
  faCheckCircle,
  faChevronRight,
} from "@fortawesome/free-solid-svg-icons";

library.add(faCheckCircle, faChevronRight);

export default {
  setup() {
    const { meta } = useMeta({
      og: {
        type: "Jobs Posting Website",
      },
      script: [{ innerHTML: "", type: "application/ld+json" }],
    });
    return {
      meta,
    };
  },
  data() {
    return {
      id: this.$route.params.id,
      checkCircle: faCheckCircle,
      isLoading: true,
      showErrorMessagePopup: false,
      jobLink: "",
    };
  },
  components: {
    ScreenHeader,
    ScreenFooter2,
    FontAwesomeIcon,
    ScreenLoader,
  },
  computed: {
    ...mapGetters({
      job: "jobById",
      jobsError: "jobsError",
    }),
  },
  async serverPrefetch() {
    await this.getCareerDetails();
  },
  mounted() {
    if (this.job == null) {
      this.getCareerDetails();
    }
    this.setCareerDetails();
    this.$gtag.event("view_page_job_detail");
  },
  methods: {
    async getCareerDetails() {
      var req = {
        jobId: this.id,
        href: this.$route.href,
      };
      await store.dispatch("getJobsById", req);
      if (this.jobsError == null) {
        this.setMetaProperties();
      }
    },
    setCareerDetails() {
      this.isLoading = false;
      if (this.jobsError != null) {
        var err = this.jobsError;
        if (err && err.status == 404) {
          this.$router.push({
            name: "Error404Page",
            params: { catchAll: "jobs/" + this.id },
          });
        } else {
          this.showErrorMessagePopup = true;
        }
      } else {
        this.jobLink = "jobs/apply/" + this.job.unique_id;
      }
    },
    setMetaProperties() {
      var seoMeta = this.job.seoData;
      this.meta.title = seoMeta.title;
      this.meta.description = seoMeta.description;
      this.meta.og.title = seoMeta.title;
      this.meta.og.url = config.BASE_URL + this.$route.href;
      this.meta.script[0].innerHTML = JSON.stringify(this.JSONLDSchema());
    },
    JSONLDSchema() {
      var career = this.job;
      var dates = this.setJobDates();
      return {
        "@context": "http://schema.org",
        "@type": "JobPosting",
        title: career.title,
        hiringOrganization: {
          "@type": "Organization",
          name: "Canopas Software LLP",
          sameAs: "https://www.canopas.com/",
          logo: "https://canopas.com/images/logo/canopas-logo.png",
        },
        employmentType: (career.employment_type
          ? career.employment_type
          : "FULL_TIME"
        ).toUpperCase(),
        baseSalary: {
          "@type": "MonetaryAmount",
          currency: "INR",
          value: {
            "@type": "QuantitativeValue",
            minValue: parseInt(
              career.base_salary ? career.base_salary : "10000"
            ),
            maxValue: 50000,
            unitText: "MONTH",
          },
        },
        datePosted: dates.jobPosted,
        description: this.setDescriptionForGoogleSchema(),
        educationRequirements: {
          "@type": "EducationalOccupationalCredential",
          credentialCategory: "bachelor degree",
        },
        jobBenefits:
          "Stipend in training. Paid leaves. Flexible working hours. Friendly environment. Work-Life balance",
        industry: "Computer Software",
        jobImmediateStart: true,
        jobLocation: {
          "@type": "Place",
          address: {
            "@type": "PostalAddress",
            streetAddress:
              "552-554, Laxmi enclave-2, near laxmi circle, opp. gajera school, katargam",
            addressLocality: "Surat",
            addressRegion: "GJ",
            postalCode: "395004",
            addressCountry: "IN",
          },
        },
        qualifications:
          "A deep desire to learn new technology. Analytical thinking, Decision-making, and problem-solving skills,Ability to work in a team environment with members of varying skill levels. Highly motivated. Learns quickly.",
        salaryCurrency: "INR",
        workHours: "9am-6pm",
        directApply: true,
        validThrough: dates.validThrough,
      };
    },
    setDescriptionForGoogleSchema() {
      var html = this.unescapeHTML(this.job.description);

      //convert string to array
      let descContent = html.split("\n");

      //remove empty string
      descContent.forEach(function (element, i) {
        if (descContent[i].length == 0) {
          descContent.splice(i, 1);
        }
      });

      // append p and li tags to desc
      descContent.forEach(function (element, i) {
        if (i == 0) {
          descContent[i] = element.replace(
            element,
            "<p>" + element + "</p></br><ul>"
          );
        } else if (
          element.includes("Responsibilities and Duties") ||
          element.includes("Benefits") ||
          element.includes("Qualifications") ||
          element.includes("Venue")
        ) {
          descContent[i] = element.replace(
            element,
            "</ul></br><p>" + element + "</p></br><ul>"
          );
        } else if (element.length != 0) {
          descContent[i] = element.replace(
            element,
            "<li>" + element + "</li></br>"
          );
        }
        if (i == descContent.length - 1) {
          descContent[i] = element.replace(
            element,
            "<li>" + element + "</li></br></ul>"
          );
        }
      });

      return descContent.join("");
    },
    unescapeHTML(escapedHTML) {
      if (escapedHTML.indexOf("&lt;") !== -1) {
        escapedHTML = escapedHTML.replace(/&lt;/g, "<");
      }
      if (escapedHTML.indexOf("&gt;") !== -1) {
        escapedHTML = escapedHTML.replace(/&gt;/g, ">");
      }
      if (escapedHTML.indexOf("&amp;") !== -1) {
        escapedHTML = escapedHTML.replace(/&amp;/g, "&");
      }
      if (escapedHTML.indexOf("&#34;") !== -1) {
        escapedHTML = escapedHTML.replace(/&#34;/g, '"');
      }
      return escapedHTML;
    },
    setJobDates() {
      var maxDays = 15;

      // current month day
      var currentDay = moment().format("DD");

      // start date of current month
      var startDateOfMonth = moment().startOf("month").add(-2, "days");

      var postedDate =
        currentDay <= maxDays
          ? startDateOfMonth
          : startDateOfMonth.add(maxDays, "days");

      var jobPosted = postedDate.format("YYYY-MM-DD");

      var validThrough = postedDate
        .add(maxDays + 5, "days")
        .format("YYYY-MM-DDT00:00");

      return { jobPosted, validThrough };
    },
    closeErrorMessageModal() {
      this.$router.push({
        path: `/jobs`,
      });
    },
  },
  updated() {
    let descriptionTitles = document.querySelectorAll("#description > h2");
    descriptionTitles.forEach((descriptionTitle) => {
      let bullet = document.createElement("span");
      bullet.className = "bullet";
      descriptionTitle.prepend(bullet);
    });
    let descriptionLists = document.querySelectorAll(
      "#description > ul > li > div"
    );
    descriptionLists.forEach((descriptionList) => {
      let fa = document.createElement("span");
      fa.className = "fas fa-chevron-right";
      fa.innerHTML = "&nbsp;";
      descriptionList.prepend(fa);
    });
  },
};
</script>

<style lang="scss" scoped>
.container {
  padding: 48px 16px;
}

.title-hr {
  border: 1px solid #e2e2e2;
  margin: 0;
}

.canopas-gradient-text {
  font-weight: 700;
  font-size: 1.75rem !important;
  line-height: 2rem !important;
  letter-spacing: 0.1rem;
}

.application-submit-btns {
  display: flex;
  justify-content: center;
  font-size: 1rem;
}

.gradient-btn {
  padding: 16px 64px;
  display: flex;
  align-items: center;
}

.gradient-btn > span {
  letter-spacing: 0.06rem;
}

.icon {
  font-size: 1.1rem;
}

.summary-text,
:deep(div > span *),
:deep(ul *),
:deep(div) {
  font-size: 1.1rem !important;
  line-height: 2rem;
  text-align: justify;
}

:deep(h2) {
  display: flex;
  flex-direction: row;
  margin: 48px 0 16px;
}

:deep(h2 *) {
  font-weight: bolder;
  font-size: 1.5rem !important;
  line-height: 2rem !important;
  letter-spacing: 0.05rem;
  padding: 12px 0;
  color: #3d3d3d;
}

:deep(h2 > .bullet) {
  content: "";
  background: linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  display: inline-block;
  line-height: 1;
  margin: 5px 20px 5px 0;
  height: auto;
  width: 8px;
  border-radius: 0;
}

:deep(ul) {
  list-style-type: none !important;
  padding-left: 0 !important;
}

// error message
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

.error-message-div {
  padding: 20px;
}

.error-message-text {
  line-height: 2rem;
  font-size: 1.5rem;
  margin-bottom: 30px;
}

.close-btn-div {
  width: 100%;
}

.close-btn {
  float: right;
}

@include media-breakpoint-up(sm) {
  .icon {
    font-size: 1.3rem;
  }
}

@include media-breakpoint-up(md) {
  .container {
    padding: 96px 96px 150px;
  }

  .canopas-gradient-text {
    font-size: 2.25rem !important;
    line-height: 2.729rem !important;
  }

  .gradient-btn {
    padding: 16px 80px;
  }

  .application-submit-btns {
    font-size: 1.125rem;
  }

  .summary-text,
  :deep(div > span *),
  :deep(ul *),
  :deep(div) {
    font-size: 1.125rem !important;
    line-height: 2.5rem !important;
  }

  :deep(h2 *) {
    font-size: 2rem !important;
    line-height: 2.5rem !important;
  }
}

@include media-breakpoint-up(lg) {
  .summary-text,
  :deep(div > span *),
  :deep(ul *),
  :deep(div) {
    font-size: 1.4rem !important;
  }
}
</style>
