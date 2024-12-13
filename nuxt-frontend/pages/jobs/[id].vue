<template>
  <div>
    <Header class="header" />
    <ScreenLoader v-if="isLoading || job == null" />
    <div v-else-if="showErrorMessagePopup">
      <transition name="modal">
        <div
          class="modal-mask fixed left-0 top-0 z-[1] table h-full w-full bg-[#00000080]"
        >
          <div
            class="left-auto mx-auto flex h-full items-center sm:mx-auto sm:my-7 sm:max-w-lg"
            role="document"
          >
            <div
              class="relative flex w-full flex-col rounded-md border-1 border-solid border-gray bg-white bg-clip-padding outline-0"
            >
              <div class="relative flex-auto p-4">
                <div class="p-5">
                  <div class="mb-[30px] text-center text-2xl">
                    Something went wrong on our side
                  </div>
                  <div class="w-full">
                    <button
                      class="gradient-btn float-right flex items-center px-16 py-4 md:px-20 md:py-4"
                      @click.prevent="closeErrorMessageModal()"
                    >
                      <span class="tracking-[0.06rem]">Close</span>
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
      <div class="px-3 py-0">
        <div
          class="container max-w-full px-4 py-12 sm:max-w-[540px] md:max-w-[720px] md:px-24 md:pb-[150px] md:pt-24 lg:max-w-[960px] xl:max-w-[1140px] 2xl:max-w-[1320px]"
        >
          <h1
            class="canopas-gradient-text text-center text-[1.75rem] font-bold leading-8 tracking-[0.1rem] md:text-[2.25rem] md:leading-[2.729rem]"
          >
            {{ job.title }}
          </h1>
          <hr class="m-0 mt-6 border border-solid border-[#e2e2e2]" />
          <div class="summary-text mt-12">
            {{ job.summary }}
          </div>
          <div class="mt-12">
            <div
              id="description"
              class="normal-text"
              v-html="job.description"
            ></div>
          </div>
          <div class="mt-12 flex justify-center text-[1rem] md:text-[1.125rem]">
            <nuxt-link
              @click.native="$mixpanel.track('tap_job_read_more_apply_now')"
              class="gradient-btn flex items-center px-16 py-4 md:px-20 md:py-4"
              :to="jobLink"
            >
              <Icon
                class="fa text-[1.1rem] sm:text-[1.3rem]"
                name="fa6-regular:circle-check"
                aria-hidden="true"
              />
              <span class="tracking-[0.06rem]">Apply Now</span>
            </nuxt-link>
          </div>
        </div>
      </div>
      <NewFooter />
    </div>
  </div>
</template>

<script setup>
import { useRoute } from "vue-router";
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import { useJobDetailStore } from "@/stores/jobs";
import config from "@/config.js";
import { unescapeHTML, getJobDates } from "@/utils.js";

const { $mixpanel } = useNuxtApp();
const route = useRoute();
const id = route.params.id;
const href = route.href;

const store = useJobDetailStore();
const job = computed(() => store.item);
const jobsError = computed(() => store.error);
const isLoading = computed(() => store.isLoading);

const showErrorMessagePopup = ref(false);
const jobLink = ref("");

await setCareerDetails();

useHead({
  script: [
    {
      innerHTML: JSON.stringify(getJsonLdSchema()),
      type: "application/ld+json",
    },
  ],
});

useSeoMeta({
  title: job.value?.seo_title,
  description: job.value?.seo_description,
  ogTitle: job.value?.seo_title,
  ogType: "Jobs Posting Website",
  ogImage: config.OG_IMAGE_URL,
  ogUrl: config.BASE_URL + route.href,
});

async function setCareerDetails() {
  await useAsyncData("jobs", () => store.loadJob(id, href));
  if (jobsError.value) {
    let err = jobsError;
    if (err && err.response && err.response.status == 404) {
      navigateTo({
        name: "Error404Page",
        params: { pathMatch: ["jobs", id] },
      });
    } else {
      showErrorMessagePopup.value = true;
    }
  } else {
    jobLink.value = "/jobs/apply/" + job.value.unique_id;
    setBulletsAndIconsInDescription();
  }
}

function setBulletsAndIconsInDescription() {
  let descriptionTitles = job.value.description?.match(/<h2>(.*?)<\/h2>/g);

  for (let i = 0; i < descriptionTitles?.length; i++) {
    if (descriptionTitles[i]) {
      let title = descriptionTitles[i].replace(/<(\/*).[^>]*>/g, "");
      job.value.description = job.value.description.replace(
        descriptionTitles[i],
        '<h2><span class="bullet"></span><strong>' + title + "</strong></h2>",
      );
    }
  }

  let descriptionLists = job.value.description?.match(/<li>([\s\S]*?)<\/li>/g);

  for (let i = 0; i < descriptionLists?.length; i++) {
    if (descriptionLists[i]) {
      let list = descriptionLists[i].replace(/<(\/*).[^>]*>/g, "");
      job.value.description = job.value.description.replace(
        descriptionLists[i],
        '<li><span class="chevron right pr-2 "></span>' + list + "</li>",
      );
    }
  }
}

function getJsonLdSchema() {
  let career = job.value;
  let dates = getJobDates();
  return {
    "@context": "http://schema.org",
    "@type": "JobPosting",
    title: career?.title,
    hiringOrganization: {
      "@type": "Organization",
      name: "Canopas",
      sameAs: "https://www.canopas.com/",
      logo: "https://canopas.com/favicon.svg",
    },
    employmentType: (career?.employment_type || "FULL_TIME").toUpperCase(),
    baseSalary: {
      "@type": "MonetaryAmount",
      currency: "INR",
      value: {
        "@type": "QuantitativeValue",
        minValue: parseInt(career?.base_salary || "10000"),
        maxValue: 50000,
        unitText: "MONTH",
      },
    },
    datePosted: dates.datePosted,
    description: getDescriptionForGoogleSchema(),
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
}

function getDescriptionForGoogleSchema() {
  let html = unescapeHTML(job.value.description);

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
        "<ul><li>" + element + "</li></ul>",
      );
    }
  });

  return descContent.join("");
}

function closeErrorMessageModal() {
  route.push({
    path: `/jobs`,
  });
}

onMounted(() => {
  $mixpanel.track("view_jobs_by_id_page");
});
</script>

<style lang="postcss" scoped>
.summary-text,
:deep(div > span *),
:deep(ul:not(.header ul)) {
  @apply text-left text-[1.1rem] leading-8 text-[#3d3d3dcc] md:text-[1.125rem] md:leading-10 lg:text-[1.4rem];
}

:deep(ul:not(.header ul) > li) {
  @apply flex items-baseline;
}
:deep(ul:not(.header ul) > li > a > svg) {
  @apply flex items-center;
}

:deep(h2) {
  @apply mx-0 mb-4 mt-12 flex flex-row;
}

:deep(h2 *) {
  @apply px-0 py-3 text-2xl tracking-[0.05rem] text-[#3d3d3d] md:text-[2rem] md:leading-10;
}

:deep(h2 > .bullet) {
  @apply mb-[5px] ml-0 mr-5 mt-[5px] inline-block h-auto w-2 rounded-none from-[#f2709c] to-[#ff9472] leading-none content-[''] bg-gradient-180;
}

:deep(ul:not(.header ul)) {
  @apply list-none pl-0;
}

:deep(ul:not(.header ul) > li > span) {
  @apply relative left-[0.15em] ml-[-5px] mr-5 inline-block h-2.5 w-2.5 rotate-[50deg] border-b-0 border-l-0 border-r-[3px] border-t-[3px] border-solid border-[#3d3d3dcc] align-top  content-none md:h-3 md:w-[13px];
}
</style>
