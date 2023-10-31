<template>
  <div>
    <Header />
    <div>
      <LandingSection :name="clientName" />
      <BenefitSection
        class="overflow-y-hidden 2xl:overflow-y-visible 3xl:overflow-y-hidden"
      />
      <HappyClientSection ref="clientreview" />
      <ScheduleMeeting />
    </div>
    <NewFooter ref="footer" />
  </div>
</template>

<script>
import { defineAsyncComponent } from "vue";
import config from "@/config.js";
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import BenefitSection from "@/components/contact/thank-you/BenefitSection.vue";
import { elementInViewPort } from "@/utils.js";
const HappyClientSection = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/HappyClient.vue"),
);

const ScheduleMeeting = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/ScheduleMeeting.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

export default {
  setup() {
    const seoData = config.CLIENT_THANKYOU_SEO_META_DATA;
    useSeoMeta({
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });
  },
  data() {
    return {
      clientName: "",
      event: "",
      events: {
        clientreview: "view_thankyou_clientreview_section",
        footer: "view_thankyou_footer",
      },
    };
  },
  components: {
    Header,
    LandingSection,
    BenefitSection,
    HappyClientSection,
    ScheduleMeeting,
    NewFooter,
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_thankyou_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  beforeMount() {
    let cName = localStorage.getItem("client-name");
    if (cName) {
      this.clientName = JSON.parse(cName);
    } else {
      navigateTo({
        name: "Error404Page",
        params: { pathMatch: ["thank-you"] },
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
