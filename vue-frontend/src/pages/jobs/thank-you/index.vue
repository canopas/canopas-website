<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <LandingSection ref="landing" :name="applicantName" />
    <InterviewProcess ref="process" />
    <VirtueView ref="virtues" />
    <LifeAtCanopasVue ref="lifeatcanopas" />
    <LatestBlog ref="blogs" />
    <ContributionSection ref="contributions" />
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
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = config.JOBS_THANKYOU_SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
        image: seoData.image,
      },
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
        landing: "view_jobs_thankyou_landing_section",
        process: "view_jobs_thankyou_process_section",
        virtues: "view_jobs_thankyou_virtues_section",
        lifeatcanopas: "view_jobs_thankyou_lifeatcanopas_section",
        blogs: "view_jobs_thankyou_blogs_section",
        contributions: "view_jobs_thankyou_contributions_section",
        footer: "view_mobileapp_development_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_jobs_thankyou_page");
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
        this.mixpanel.track(event);
      }
    },
  },
};
</script>

<style scoped></style>
