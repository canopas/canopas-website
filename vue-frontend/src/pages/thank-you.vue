<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div>
      <LandingSection :name="clientName" />
      <BenefitSection
        class="tw-overflow-y-hidden 2xl:tw-overflow-y-visible 3xl:tw-overflow-y-hidden"
      />
      <ScheduleMeeting />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import BenefitSection from "@/components/contact/thank-you/BenefitSection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";

const ScheduleMeeting = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/ScheduleMeeting.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);
export default {
  setup() {
    var seoData = config.CLIENT_THANKYOU_SEO_META_DATA;
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
  data() {
    return {
      clientName: "Dear Client",
    };
  },
  mounted() {
    this.clientName =
      localStorage.getItem("client-name") != ""
        ? JSON.parse(localStorage.getItem("client-name"))
        : "Dear Client";
  },
  components: {
    Header,
    LandingSection,
    BenefitSection,
    ScheduleMeeting,
    NewFooter,
  },
  beforeRouteEnter(to, from, next) {
    if (config.SHOW_CLIENT_THANKYOU_PAGE === false) {
      next({ name: "Error404Page", params: { pathMatch: "/thank-you" } });
    } else {
      next();
    }
  },
};
</script>
