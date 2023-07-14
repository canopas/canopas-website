<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div class="xll:tw-bg-black-core/[90%]">
      <LandingSection ref="landing" />
    </div>
    <DevelopmentSection ref="development" />
    <PinkCtaSection ref="cta1" class="tw-hidden md:tw-block" />
    <BlackCtaSection ref="cta2" />
    <CtaSection ref="cta3" />
    <NewFooter ref="footer" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/android-app-development/LandingSection.vue";
import DevelopmentSection from "@/components/android-app-development/DevelopmentSection.vue";
const PinkCtaSection = defineAsyncComponent(() =>
  import("@/components/android-app-development/PinkCtaSection.vue")
);
const CtaSection = defineAsyncComponent(() =>
  import("@/components/android-app-development/CtaSection.vue")
);
const BlackCtaSection = defineAsyncComponent(() =>
  import("@/components/android-app-development/BlackCtaSection.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  data() {
    return {
      event: "",
      events: {
        landing: "view_android_development_landing_section",
        footer: "view_android_development_footer",
        development: "view_android_development_section",
        cta1: "view_android_development_cta_section",
        cta2: "view_android_development_cta2_section",
        cta3: "view_android_development_cta3_section",
      },
    };
  },
  setup() {
    var seoData = config.ANDRIOD_APP_DEVELOPMENT_SEO_META_DATA;
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
    DevelopmentSection,
    PinkCtaSection,
    BlackCtaSection,
    CtaSection,
    NewFooter,
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
  beforeRouteEnter(to, from, next) {
    if (!config.SHOW_ANDROID_APP_DEVELOPMENT_PAGE) {
      next({
        name: "Error404Page",
        params: { pathMatch: "/android-app-development" },
      });
    } else {
      next();
    }
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_android_development_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
};
</script>
