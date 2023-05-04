<template>
  <section
    class="tw-container tw-mb-20 md:tw-mb-40 tw-flex tw-flex-col tw-items-center"
  >
    <div class="tw-mt-20 tw-text-center">
      <span
        class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.25rem] md:tw-leading-[5.15625rem]"
        >Technology Stack</span
      >
      <p
        class="tw-mt-5 tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem]"
      >
        The products we build are typically deployed to mobile and web. Web and
        mobile are our bread and butter. From native to hybrid, frontend to
        backend, designing to deployment, we've got you covered. At Canopas,
        we're not bound by technologies. We learn new technology if it has
        chances to improve your app performance significantly.
      </p>
    </div>
    <div
      class="tw-mt-16 tw-h-auto tw-w-[380px] tw-rounded-2xl tw-bg-[#F6F6F6]"
      v-for="stack in stacks"
      :key="stack"
    >
      <span
        class="tw--mt-6 tw-flex tw-items-center tw-w-max tw-rounded-full tw-p-3 tw-text-center gradient-btn tw-pointer-events-none"
      >
        <span
          class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0]"
          >{{ stack.title }}
        </span>
      </span>
      <div
        class="tw-my-7 tw-flex tw-flex-row tw-flex-wrap tw-items-center tw-pl-7"
      >
        <img
          v-for="techs in stack.image"
          :key="techs"
          class="tw-flex tw-w-28 tw-h-28 tw-object-cover tw-mt-4 tw-rounded-[5px] tw-pr-[10px]"
          :src="techs"
          alt="technology-stacks-image"
        />
      </div>
    </div>
  </section>
</template>

<script>
import frontend1 from "@/assets/images/services/technology/frontend-1-400w.webp";
import frontend2 from "@/assets/images/services/technology/frontend-2-400w.webp";
import frontend3 from "@/assets/images/services/technology/frontend-3-400w.webp";
import frontend4 from "@/assets/images/services/technology/frontend-4-400w.webp";
import frontend5 from "@/assets/images/services/technology/frontend-5-400w.webp";
import frontend6 from "@/assets/images/services/technology/frontend-6-400w.webp";
import frontend7 from "@/assets/images/services/technology/frontend-7-400w.webp";

import backend1 from "@/assets/images/services/technology/backend-1-400w.webp";
import backend2 from "@/assets/images/services/technology/backend-2-400w.webp";
import backend3 from "@/assets/images/services/technology/backend-3-400w.webp";
import backend4 from "@/assets/images/services/technology/backend-4-400w.webp";
import backend5 from "@/assets/images/services/technology/backend-5-400w.webp";

import mobileapp1 from "@/assets/images/services/technology/mobileapp-1-400w.webp";
import mobileapp2 from "@/assets/images/services/technology/mobileapp-2-400w.webp";
import mobileapp3 from "@/assets/images/services/technology/mobileapp-3-400w.webp";
import mobileapp4 from "@/assets/images/services/technology/mobileapp-4-400w.webp";

import ui1 from "@/assets/images/services/technology/ui-1-400w.webp";
import ui2 from "@/assets/images/services/technology/ui-2-400w.webp";
import ui3 from "@/assets/images/services/technology/ui-3-400w.webp";
import ui4 from "@/assets/images/services/technology/ui-4-400w.webp";
import ui5 from "@/assets/images/services/technology/ui-5-400w.webp";

import seo1 from "@/assets/images/services/technology/seo-1-400w.webp";
import seo2 from "@/assets/images/services/technology/seo-2-400w.webp";
import seo3 from "@/assets/images/services/technology/seo-3-400w.webp";
import seo4 from "@/assets/images/services/technology/seo-4-400w.webp";
import seo5 from "@/assets/images/services/technology/seo-5-400w.webp";

export default {
  data() {
    return {
      stacks: [
        {
          title: "Frontend",
          image: [
            frontend1,
            frontend2,
            frontend3,
            frontend4,
            frontend5,
            frontend6,
            frontend7,
          ],
        },
        {
          title: "Backend ",
          image: [backend1, backend2, backend3, backend4, backend5],
        },
        {
          title: "Mobile App",
          image: [mobileapp1, mobileapp2, mobileapp3, mobileapp4],
        },
        {
          title: "UI/UX Design",
          image: [ui1, ui2, ui3, ui4, ui5],
        },
        {
          title: "SEO",
          image: [seo1, seo2, seo3, seo4, seo5],
        },
      ],
    };
  },
};
</script>
