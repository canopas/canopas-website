<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div>
      <LandingSection :name="clientName" ref="landing" />
      <BenefitSection
        class="tw-overflow-y-hidden 2xl:tw-overflow-y-visible 3xl:tw-overflow-y-hidden"
        ref="benefits"
      />
      <HappyClientSection ref="clientreview" />
      <ScheduleMeeting ref="schedulemeeting" />
    </div>
    <NewFooter ref="footer" />
  </div>
</template>

<script>
import { defineAsyncComponent } from "vue";
import config from "@/config.js";
// import { useMeta } from "vue-meta";
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import BenefitSection from "@/components/contact/thank-you/BenefitSection.vue";
const HappyClientSection = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/HappyClient.vue")
);

const ScheduleMeeting = defineAsyncComponent(() =>
  import("@/components/contact/thank-you/ScheduleMeeting.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  setup() {
    var seoData = config.CLIENT_THANKYOU_SEO_META_DATA;
    // useMeta({
    //   title: seoData.title,
    //   description: seoData.description,
    //   og: {
    //     type: seoData.type,
    //     title: seoData.title,
    //     url: seoData.url,
    //     image: seoData.image,
    //   },
    // });
  },
  data() {
    return {
      clientName: "",
      event: "",
      events: {
        landing: "view_thankyou_landing_section",
        benefits: "view_thankyou_benefits_section",
        clientreview: "view_thankyou_clientreview_section",
        schedulemeeting: "view_thankyou_schedulemeeting_section",
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
  // inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    // // this.mixPanel.track("view_thankyou_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  beforeMount() {
    if (localStorage.getItem("client-name")) {
      this.clientName = JSON.parse(localStorage.getItem("client-name"));
    } else {
      this.$router.push({
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
        // // this.mixPanel.track(event);
      }
    },
  },
};
</script>
