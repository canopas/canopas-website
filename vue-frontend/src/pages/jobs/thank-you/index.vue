<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <LandingSection :name="applicantName" />
    <InterviewProcess />
    <VirtueView />
    <LifeAtCanopasVue />
    <LatestBlog />
    <ContributionSection />
    <NewFooter />
  </div>
</template>
<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import InterviewProcess from "@/components/jobs/thank-you/InterviewProcess.vue";
import { defineAsyncComponent } from "vue";

const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);
const VirtueView = defineAsyncComponent(() =>
  import("@/components/jobs/VirtuesView.vue")
);

const LifeAtCanopasVue = defineAsyncComponent(() =>
  import("@/components/jobs/PerksAndBenefits.vue")
);
const LatestBlog = defineAsyncComponent(() =>
  import("@/components/jobs/thank-you/LatestBlog.vue")
);
const ContributionSection = defineAsyncComponent(() =>
  import("@/components/jobs/thank-you/ContributionSection.vue")
);
const Error404Page = () => import("@/components/error404/index.vue");
import config from "@/config.js";
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = config.JOBS_THANKYOU_SEO_META_DATA;
    useMeta({
      meta: [
        {
          name: "robots",
          content: "noindex, nofollow",
          vmid: "robots",
        },
      ],
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
      applicantName: "Dear Applicant",
    };
  },
  mounted() {
    this.applicantName = localStorage.getItem("applicant-name")
      ? JSON.parse(localStorage.getItem("applicant-name"))
      : "Dear Applicant";

    if (!localStorage.getItem("applicant-name")) {
      routes.push({
        path: "/:pathMatch(.*)*",
        name: "Error404Page",
        component: Error404Page,
      });
    }
  },
};
</script>

<style scoped></style>
