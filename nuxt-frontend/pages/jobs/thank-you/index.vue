<template>
  <div>
    <Header />
    <LandingSection :name="applicantName" />
    <InterviewProcess />
    <VirtueView />
    <LifeAtCanopasVue ref="lifeatcanopas" />
    <LatestBlog />
    <ContributionSection />
    <NewFooter ref="footer" />
  </div>
</template>
<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import InterviewProcess from "@/components/jobs/thank-you/InterviewProcess.vue";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const VirtueView = defineAsyncComponent(() =>
  import("@/components/jobs/VirtuesView.vue"),
);
const LifeAtCanopasVue = defineAsyncComponent(() =>
  import("@/components/jobs/PerksAndBenefits.vue"),
);
const LatestBlog = defineAsyncComponent(() =>
  import("@/components/jobs/thank-you/LatestBlog.vue"),
);
const ContributionSection = defineAsyncComponent(() =>
  import("@/components/jobs/thank-you/ContributionSection.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

import config from "@/config.js";

export default {
  setup() {
    const seoData = config.JOBS_THANKYOU_SEO_META_DATA;
    useSeoMeta({
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });
  },
  components: {
    Header,
    LandingSection,
    InterviewProcess,
    VirtueView,
    LifeAtCanopasVue,
    LatestBlog,
    ContributionSection,
    NewFooter,
  },
  data() {
    return {
      applicantName: "",
      event: "",
      events: {
        lifeatcanopas: "view_jobs_thankyou_lifeatcanopas_section",
        footer: "view_mobileapp_development_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_jobs_thankyou_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  beforeMount() {
    if (localStorage.getItem("applicant-name")) {
      this.applicantName = JSON.parse(localStorage.getItem("applicant-name"));
    } else {
      this.$router.push({
        name: "Error404Page",
        params: { pathMatch: ["jobs", "thank-you"] },
      });
    }
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.$mixpanel.track(event);
      }
    },
  },
};
</script>
