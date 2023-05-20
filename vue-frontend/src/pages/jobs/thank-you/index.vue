<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <InterviewProcess />
    <VirtueView />
    <LifeAtCanopasVue />
    <ContributionSection />
    <NewFooter />
  </div>
</template>
<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import InterviewProcess from "@/components/jobs-thankyou/InterviewProcess.vue";
import VirtueView from "@/components/jobs/VirtuesView.vue";
import LifeAtCanopasVue from "@/components/jobs/PerksAndBenefits.vue";
import ContributionSection from "@/components/jobs-thankyou/ContributionSection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = config.JOBS_SEO_META_DATA;
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
    InterviewProcess,
    VirtueView,
    LifeAtCanopasVue,
    ContributionSection,
    NewFooter,
  },
  beforeRouteEnter(to, from, next) {
    if (config.SHOW_JOBSTHANKYOU_PAGE === false) {
      next({ name: "Error404Page", params: { pathMatch: "/jobs/thank-you" } });
    } else {
      next();
    }
  },
  data() {
    return {};
  },
};
</script>

<style scoped></style>
