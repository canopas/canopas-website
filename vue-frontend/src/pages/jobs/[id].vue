<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <link
      rel="stylesheet"
      href="https://use.fontawesome.com/releases/v5.2.0/css/all.css"
    />
    <ScreenHeaderV2 />
    <ScreenLoader v-if="isLoading" />
    <div v-else-if="showErrorMessagePopup">
      <transition name="modal">
        <div
          class="modal-mask tw-fixed tw-z-[1] tw-top-0 tw-left-0 tw-w-full tw-h-full tw-bg-[#00000080] tw-table"
        >
          <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content calendly-iframe-modal-content">
              <div class="modal-body">
                <div class="tw-p-[20px]">
                  <div
                    class="tw-leading-[2rem] tw-text-[1.5rem] tw-mb-[30px] tw-text-center"
                  >
                    Something went wrong on our side
                  </div>
                  <div class="tw-w-full">
                    <button
                      class="gradient-btn tw-py-[16px] tw-px-[64px] tw-flex tw-items-center md:tw-py-[16px] md:tw-px-[80px] tw-float-right"
                      @click.prevent="closeErrorMessageModal()"
                    >
                      <span class="tw-tracking-[0.06rem]">Close</span>
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
      <div class="tw-py-0 tw-px-[0.75rem]">
        <div
          class="tw-container tw-py-[48px] tw-px-[16px] tw-max-w-full sm:tw-max-w-[540px] md:tw-max-w-[720px] lg:tw-max-w-[960px] xl:tw-max-w-[1140px] 2xl:tw-max-w-[1320px] md:tw-pt-[96px] md:tw-px-[96px] md:tw-pb-[150px]"
        >
          <div
            class="canopas-gradient-text tw-font-bold tw-text-[1.75rem] tw-leading-8 tw-tracking-[0.1rem] md:tw-text-[2.25rem] md:tw-leading-[2.729rem] tw-text-center"
          >
            {{ job.title }}
          </div>
          <hr
            class="tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-m-0 tw-mt-6"
          />
          <div class="summary-text tw-mt-12">
            {{ job.summary }}
          </div>
          <div class="tw-mt-12">
            <div
              id="description"
              class="normal-text"
              v-html="job.description"
            ></div>
          </div>
          <div
            class="tw-flex tw-justify-center tw-text-[1rem] md:tw-text-[1.125rem] tw-mt-12"
          >
            <router-link
              class="gradient-btn tw-py-[16px] tw-px-[64px] tw-flex tw-items-center md:tw-py-[16px] md:tw-px-[80px]"
              :to="jobLink"
            >
              <font-awesome-icon
                class="fa tw-text-[1.1rem] sm:tw-text-[1.3rem]"
                :icon="checkCircle"
                aria-hidden="true"
              />
              <span class="tw-tracking-[0.06rem]">Apply Now</span>
            </router-link>
          </div>
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
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { useJobDetailStore } from "@/stores/jobs";
import { mapState } from "pinia";
import { mapActions } from "pinia";
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
      showErrorMessagePopup: false,
      jobLink: "",
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
  computed: {
    ...mapState(useJobDetailStore, {
      job: "item",
      jobsError: "error",
      isLoading: "isLoading",
    }),
  },
  mounted() {
    this.setCareerDetails();
    this.$gtag.event("view_page_job_detail");
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
            params: { pathMatch: ["jobs", this.id] },
          });
        } else {
          this.showErrorMessagePopup = true;
        }
      } else {
        this.jobLink = "/jobs/apply/" + this.job.unique_id;
        this.setMetaProperties();
        this.setBulletsAndIconsInDescription();
      }
    },
    setMetaProperties() {
      this.meta.title = this.job.seo_title;
      this.meta.description = this.job.seo_description;
      this.meta.og.title = this.job.seo_title;
      this.meta.og.url = config.BASE_URL + this.$route.href;
      this.meta.script[0].innerHTML = JSON.stringify(this.getJsonLdSchema());
    },
    getJsonLdSchema() {
      var career = this.job;
      var dates = this.getJobDates();
      return {
        "@context": "http://schema.org",
        "@type": "JobPosting",
        title: career.title,
        hiringOrganization: {
          "@type": "Organization",
          name: "Canopas",
          sameAs: "https://www.canopas.com/",
          logo: "https://canopas.com/favicon.svg",
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
        description: this.getDescriptionForGoogleSchema(),
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
    getDescriptionForGoogleSchema() {
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
          descContent[i] = element.replace(element, "<p>" + element + "</p>");
        } else if (
          element.includes("Responsibilities and Duties") ||
          element.includes("Benefits") ||
          element.includes("Qualifications") ||
          element.includes("Venue")
        ) {
          descContent[i] = element.replace(element, "<p>" + element + "</p>");
        } else if (element.length != 0) {
          descContent[i] = element.replace(element, element);
        }
        if (i == descContent.length - 1) {
          descContent[i] = element.replace(
            element,
            "<ul><li>" + element + "</li></ul>"
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
    getJobDates() {
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
    setBulletsAndIconsInDescription() {
      let descriptionTitles = this.job.description.match(/<h2>(.*?)<\/h2>/g);

      for (let i = 0; i < descriptionTitles.length; i++) {
        if (descriptionTitles[i]) {
          let title = descriptionTitles[i].match(/<strong>(.*?)<\/strong>/g)[0];
          this.job.description = this.job.description.replace(
            descriptionTitles[i],
            '<h2><span class="bullet"></span>' + title + "</h2>"
          );
        }
      }

      let descriptionLists =
        this.job.description.match(/<li>([\s\S]*?)<\/li>/g);

      for (let i = 0; i < descriptionLists.length; i++) {
        if (descriptionLists[i]) {
          let list = descriptionLists[i].match(/<div>([\s\S]*?)<\/div>/g)[0];
          this.job.description = this.job.description.replace(
            descriptionLists[i],
            '<li><span class="fas fa-chevron-right"></span>&nbsp;' +
              list +
              "</li>"
          );
        }
      }
    },
  },
};
</script>

<style lang="postcss" scoped>
.summary-text,
:deep(div > span *),
:deep(ul) {
  @apply tw-text-[1.1rem] tw-leading-[2rem] tw-text-justify tw-text-[#3d3d3dcc] md:tw-text-[1.125rem] md:tw-leading-[2.5rem] lg:tw-text-[1.4rem];
}

:deep(ul > li) {
  @apply tw-flex tw-items-baseline;
}

:deep(h2) {
  @apply tw-flex tw-flex-row tw-mt-[48px] tw-mx-0 tw-mb-[16px];
}

:deep(h2 *) {
  @apply tw-text-[1.5rem] tw-leading-[2rem] tw-tracking-[0.05rem] tw-py-[12px] tw-px-0 tw-text-[#3d3d3d] md:tw-text-[2rem] md:tw-leading-[2.5rem];
}

:deep(h2 > .bullet) {
  @apply tw-content-[''] tw-from-[#f2709c] tw-to-[#ff9472] tw-bg-gradient-180 tw-inline-block tw-leading-none tw-mt-[5px] tw-mr-[20px] tw-mb-[5px] tw-ml-0 tw-h-auto tw-w-[8px] tw-rounded-none;
}

:deep(ul) {
  @apply tw-list-none tw-pl-0;
}
</style>
