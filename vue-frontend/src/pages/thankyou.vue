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
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/ThankyouLandingSection.vue";
import BenefitSection from "@/components/contact/ThankyouBenefitSection.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
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
      clientName:
        localStorage.getItem("client-name") != ""
          ? JSON.parse(localStorage.getItem("client-name"))
          : "Dear Client",
    };
  },
  components: {
    Header,
    LandingSection,
    BenefitSection,
    NewFooter,
  },
  beforeRouteEnter(to, from, next) {
    if (config.SHOW_CLIENT_THANKYOU_PAGE === false) {
      next({ name: "Error404Page", params: { pathMatch: "/thankyou" } });
    } else {
      next();
    }
  },
};
</script>
