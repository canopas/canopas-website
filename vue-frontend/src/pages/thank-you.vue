<template>
  <div v-if="show">
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
      <HappyClientSection />

      <ScheduleMeeting />
    </div>
    <NewFooter />
  </div>
  <Error404Page v-else />
</template>

<script>
import { defineAsyncComponent } from "vue";
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import BenefitSection from "@/components/contact/thank-you/BenefitSection.vue";
const HappyClientSection = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/HappyClient.vue")
);
import config from "@/config.js";
import { useMeta } from "vue-meta";

const ScheduleMeeting = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/ScheduleMeeting.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);
const Error404Page = defineAsyncComponent(() =>
  import("@/components/error404/index.vue")
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
      clientName: "",
      show: false,
    };
  },
  components: {
    Header,
    LandingSection,
    BenefitSection,
    HappyClientSection,
    ScheduleMeeting,
    NewFooter,
    Error404Page,
  },
  beforeMount() {
    if (localStorage.getItem("client-name")) {
      this.clientName = JSON.parse(localStorage.getItem("client-name"));
      this.show = true;
    }
  },
};
</script>
