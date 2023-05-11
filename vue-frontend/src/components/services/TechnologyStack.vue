<template>
  <section
    class="tw-flex tw-flex-col tw-items-center tw-mb-20 md:tw-mb-28 md:tw-mx-auto md:tw-w-8/12"
  >
    <div class="tw-container tw-mt-20 lg:tw-mt-0 tw-text-center">
      <span
        class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.25rem] md:tw-leading-[5.15625rem] tw-text-black-core/[0.87]"
        >Technology Stack</span
      >
      <p
        class="tw-mt-5 tw-font-inter-medium tw-text-black-core/[0.87] md:tw-text-black-core/[0.6] tw-text-[1rem] md:tw-text-[1.5rem] tw-leading-[1.5rem] md:tw-leading-[2.25rem]"
      >
        The products we build are typically deployed to mobile and web. Web and
        mobile are our bread and butter. From native to hybrid, frontend to
        backend, designing to deployment, we've got you covered. At Canopas,
        we're not bound by technologies. We learn new technology if it has
        chances to improve your app performance significantly.
      </p>
    </div>
    <div
      class="xl:tw-hidden tw-mt-16 tw-h-auto tw-w-[380px] tw-rounded-2xl tw-bg-[#F6F6F6]"
      v-for="stack in stacks"
      :key="stack"
    >
      <span
        class="tw-flex tw-items-center tw--mt-6 tw-mx-auto tw-w-max tw-rounded-full tw-py-2.5 md:tw-py-3 tw-px-10 tw-text-center gradient-btn tw-pointer-events-none"
      >
        <span
          class="tw-mr-2.5 tw-text-[1.25rem] md:tw-text-[1.09375rem] lg:tw-text-[1.1875rem] tw-leading-[1.875rem] md:tw-leading-[1.3125rem] lg:tw-leading-[1.4375rem] tw-font-inter-semibold"
          >{{ stack.title }}
        </span>
      </span>
      <div
        class="tw-flex tw-flex-row tw-flex-wrap tw-items-center tw-my-7 tw-pl-7"
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
    <div class="tw-hidden xl:tw-block">
      <div
        class="tw-grid tw-grid-rows-2 tw-grid-flow-col tw-gap-x-4 xl:tw-gap-x-12"
      >
        <div class="tw-row-start-2 tw-row-span-1 tw--mt-40">
          <div v-for="stack in stack1" :key="stack" class="tw-flex">
            <div
              class="tw-mt-[3.5rem] tw-h-auto tw-w-[368px] tw-rounded-2xl tw-bg-white tw-drop-shadow-md"
            >
              <span
                class="tw-flex tw-items-center tw-mx-auto tw-mt-[-23px] tw-w-max tw-rounded-full tw-py-3 tw-px-5 tw-text-center gradient-btn tw-pointer-events-none"
              >
                <span
                  class="tw-mr-2.5 tw-text-[1.1875rem] tw-leading-[1.4375rem] tw-font-inter-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div
                class="tw-flex tw-flex-wrap tw-my-7 tw-px-[30px] tw-pb-[30px]"
              >
                <div
                  v-for="techs in stack.icons"
                  :key="techs"
                  class="tw-flex tw-flex-col tw-items-center tw-my-5 tw-ml-[0.5rem] tw-w-[30%]"
                >
                  <img
                    class="tw-mb-4 tw-h-[30px] tw-w-[30px]"
                    :src="techs.image"
                    alt="technology-stacks-image"
                  />
                  <span
                    class="tw-font-inter-regular tw-text-[1.125rem] tw-leading-[1.6875rem] tw-text-center"
                    >{{ techs.content }}</span
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="tw-row-start-1 tw-row-span-2 tw-mt-10">
          <div
            v-for="stack in stack2"
            :key="stack"
            class="tw-flex tw-flex-col tw-nowrap"
          >
            <div
              class="tw-mt-16 tw-h-auto tw-w-[368px] tw-rounded-2xl tw-bg-white tw-drop-shadow-md"
            >
              <span
                class="tw-flex tw-items-center tw-mx-auto tw-mt-[-23px] tw-w-max tw-rounded-full tw-py-3 tw-px-5 tw-text-center gradient-btn tw-pointer-events-none"
              >
                <span
                  class="tw-mr-2.5 tw-text-[1.1875rem] tw-leading-[1.4375rem] tw-font-inter-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div
                class="tw-flex tw-flex-wrap tw-my-7 tw-px-[30px] tw-pb-[30px]"
              >
                <div
                  v-for="techs in stack.icons"
                  :key="techs"
                  class="tw-flex tw-flex-col tw-items-center tw-my-5 tw-ml-[0.5rem] tw-w-[30%]"
                >
                  <img
                    class="tw-mb-4 tw-h-[30px] tw-w-[30px]"
                    :src="techs.image"
                    alt="technology-stacks-image"
                  />

                  <span
                    class="tw-font-inter-regular tw-text-[1.125rem] tw-leading-[1.6875rem] tw-text-center"
                    >{{ techs.content }}</span
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="tw-row-start-1 tw-row-span-2 tw-mt-64">
          <div
            v-for="stack in stack3"
            :key="stack"
            class="tw-flex tw-flex-col tw-nowrap"
          >
            <div
              class="tw-mt-16 tw-h-auto tw-w-[368px] tw-rounded-2xl tw-bg-white tw-drop-shadow-md"
            >
              <span
                class="tw-flex tw-items-center tw-mx-auto tw-mt-[-23px] tw-w-max tw-rounded-full tw-py-3 tw-px-5 tw-text-center gradient-btn tw-pointer-events-none"
              >
                <span
                  class="tw-mr-2.5 tw-text-[1.1875rem] tw-leading-[1.4375rem] tw-font-inter-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div
                class="tw-flex tw-flex-wrap tw-my-7 tw-px-[30px] tw-pb-[30px]"
              >
                <div
                  v-for="techs in stack.icons"
                  :key="techs"
                  class="tw-flex tw-flex-col tw-items-center tw-my-5 tw-ml-[0.5rem] tw-w-[30%]"
                >
                  <img
                    class="tw-mb-4 tw-h-[30px] tw-w-[30px]"
                    :src="techs.image"
                    alt="technology-stacks-image"
                  />

                  <span
                    class="tw-font-inter-regular tw-text-[1.125rem] tw-leading-[1.6875rem] tw-text-center"
                    >{{ techs.content }}</span
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
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

import desktopfrontend1 from "@/assets/images/services/technology/desktop-icon/frontend-1-400w.webp";
import desktopfrontend2 from "@/assets/images/services/technology/desktop-icon/frontend-2-400w.webp";
import desktopfrontend3 from "@/assets/images/services/technology/desktop-icon/frontend-3-400w.webp";
import desktopfrontend4 from "@/assets/images/services/technology/desktop-icon/frontend-4-400w.webp";
import desktopfrontend5 from "@/assets/images/services/technology/desktop-icon/frontend-5-400w.webp";
import desktopfrontend6 from "@/assets/images/services/technology/desktop-icon/frontend-6-400w.webp";
import desktopfrontend7 from "@/assets/images/services/technology/desktop-icon/frontend-7-400w.webp";

import backend1 from "@/assets/images/services/technology/backend-1-400w.webp";
import backend2 from "@/assets/images/services/technology/backend-2-400w.webp";
import backend3 from "@/assets/images/services/technology/backend-3-400w.webp";
import backend4 from "@/assets/images/services/technology/backend-4-400w.webp";
import backend5 from "@/assets/images/services/technology/backend-5-400w.webp";

import desktopbackend1 from "@/assets/images/services/technology/desktop-icon/backend-1-400w.webp";
import desktopbackend2 from "@/assets/images/services/technology/desktop-icon/backend-2-400w.webp";
import desktopbackend3 from "@/assets/images/services/technology/desktop-icon/backend-3-400w.webp";
import desktopbackend4 from "@/assets/images/services/technology/desktop-icon/backend-4-400w.webp";
import desktopbackend5 from "@/assets/images/services/technology/desktop-icon/backend-5-400w.webp";

import mobileapp1 from "@/assets/images/services/technology/mobileapp-1-400w.webp";
import mobileapp2 from "@/assets/images/services/technology/mobileapp-2-400w.webp";
import mobileapp3 from "@/assets/images/services/technology/mobileapp-3-400w.webp";
import mobileapp4 from "@/assets/images/services/technology/mobileapp-4-400w.webp";

import desktopmobileapp1 from "@/assets/images/services/technology/desktop-icon/mobileapp-1-400w.webp";
import desktopmobileapp2 from "@/assets/images/services/technology/desktop-icon/mobileapp-2-400w.webp";
import desktopmobileapp3 from "@/assets/images/services/technology/desktop-icon/mobileapp-3-400w.webp";
import desktopmobileapp4 from "@/assets/images/services/technology/desktop-icon/mobileapp-4-400w.webp";

import ui1 from "@/assets/images/services/technology/ui-1-400w.webp";
import ui2 from "@/assets/images/services/technology/ui-2-400w.webp";
import ui3 from "@/assets/images/services/technology/ui-3-400w.webp";
import ui4 from "@/assets/images/services/technology/ui-4-400w.webp";
import ui5 from "@/assets/images/services/technology/ui-5-400w.webp";

import desktopui1 from "@/assets/images/services/technology/desktop-icon/ui-1-400w.webp";
import desktopui2 from "@/assets/images/services/technology/desktop-icon/ui-2-400w.webp";
import desktopui3 from "@/assets/images/services/technology/desktop-icon/ui-3-400w.webp";
import desktopui4 from "@/assets/images/services/technology/desktop-icon/ui-4-400w.webp";
import desktopui5 from "@/assets/images/services/technology/desktop-icon/ui-5-400w.webp";

import seo1 from "@/assets/images/services/technology/seo-1-400w.webp";
import seo2 from "@/assets/images/services/technology/seo-2-400w.webp";
import seo3 from "@/assets/images/services/technology/seo-3-400w.webp";
import seo4 from "@/assets/images/services/technology/seo-4-400w.webp";
import seo5 from "@/assets/images/services/technology/seo-5-400w.webp";

import desktopseo1 from "@/assets/images/services/technology/desktop-icon/seo-1-400w.webp";
import desktopseo2 from "@/assets/images/services/technology/desktop-icon/seo-2-400w.webp";
import desktopseo3 from "@/assets/images/services/technology/desktop-icon/seo-3-400w.webp";
import desktopseo4 from "@/assets/images/services/technology/desktop-icon/seo-4-400w.webp";
import desktopseo5 from "@/assets/images/services/technology/desktop-icon/seo-5-400w.webp";

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
      stack1: [
        {
          title: "Backend",
          icons: [
            { image: desktopbackend1, content: "Golang" },
            { image: desktopbackend2, content: "Node.js" },
            { image: desktopbackend3, content: "PHP Laravel" },
            { image: desktopbackend4, content: "Ruby on Rails" },
            { image: desktopbackend5, content: "Python" },
          ],
        },
      ],

      stack2: [
        {
          title: "Frontend",
          icons: [
            { image: desktopfrontend1, content: "Vue.js" },
            { image: desktopfrontend2, content: "React.js" },
            { image: desktopfrontend3, content: "Tailwind CSS" },
            { image: desktopfrontend4, content: "Bootstrap" },
            { image: desktopfrontend5, content: "WordPress" },
            { image: desktopfrontend6, content: "Next.js" },
            { image: desktopfrontend7, content: "Nuxt.js" },
          ],
        },
        {
          title: "Mobile App",
          icons: [
            { image: desktopmobileapp1, content: "Kotlin" },
            { image: desktopmobileapp2, content: "Swift" },
            { image: desktopmobileapp3, content: "Dart" },
            { image: desktopmobileapp4, content: "React Native" },
          ],
        },
      ],
      stack3: [
        {
          title: "UI/UX Design",
          icons: [
            { image: desktopui1, content: "Figma" },
            { image: desktopui2, content: "Adobe XD" },
            { image: desktopui3, content: "Sketch" },
            { image: desktopui4, content: "Illustrator" },
            { image: desktopui5, content: "Photoshop" },
          ],
        },
        {
          title: "SEO",
          icons: [
            { image: desktopseo1, content: "Ahrefs" },
            { image: desktopseo2, content: "Google Adwords" },
            { image: desktopseo3, content: "Google Analytics" },
            { image: desktopseo4, content: "SEMrush" },
            { image: desktopseo5, content: "Google Search Console" },
          ],
        },
      ],
    };
  },
};
</script>
