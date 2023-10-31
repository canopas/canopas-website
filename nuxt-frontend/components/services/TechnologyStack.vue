<template>
  <section
    class="flex flex-col mt-[7.813rem] items-center mb-20 md:mb-[14.063rem] md:mx-auto md:w-8/12"
  >
    <div class="container mt-20 lg:mt-32 text-center">
      <span
        class="font-inter-bold text-[1.875rem] md:text-[3.4375rem] leading-9 md:leading-[5.15625rem] text-black-core/[0.87]"
        >Technology Stack</span
      >
      <p
        class="mt-5 font-inter-medium text-black-core/[0.87] md:text-black-core/[0.6] text-[1rem] md:text-[1.5rem] leading-6 md:leading-9"
      >
        The products we build are typically deployed to mobile and web. Web and
        mobile are our bread and butter. From native to hybrid, frontend to
        backend, designing to deployment, we've got you covered. At Canopas,
        we're not bound by technologies. We learn new technology if it has
        chances to improve your app performance significantly.
      </p>
    </div>
    <div
      class="xl:hidden mt-16 h-auto w-[380px] rounded-2xl bg-[#F6F6F6]"
      v-for="stack in stacks"
      :key="stack"
    >
      <span
        class="flex items-center -mt-6 mx-auto w-max rounded-full py-2.5 md:py-3 px-10 text-center gradient-btn pointer-events-none"
      >
        <span
          class="mr-2.5 text-[1.25rem] md:text-[1.09375rem] lg:text-[1.1875rem] leading-[1.875rem] md:leading-[1.3125rem] lg:leading-[1.4375rem] font-inter-semibold"
          >{{ stack.title }}
        </span>
      </span>
      <div class="flex flex-row flex-wrap items-center my-7 pl-7">
        <img
          v-for="techs in stack.image"
          :key="techs"
          class="flex w-28 h-28 object-cover mt-4 rounded-[5px] pr-[10px]"
          :src="techs"
          alt="technology-stacks-image"
        />
      </div>
    </div>
    <div class="hidden xl:block">
      <div class="grid grid-rows-2 grid-flow-col gap-x-4 xl:gap-x-12">
        <div class="row-start-2 row-span-1 -mt-40">
          <div v-for="stack in stack1" :key="stack" class="flex">
            <div
              class="mt-14 h-auto w-[368px] rounded-2xl bg-white drop-shadow-md"
            >
              <span
                class="flex items-center mx-auto mt-[-23px] w-max rounded-full py-3 px-5 text-center gradient-btn pointer-events-none"
              >
                <span
                  class="mr-2.5 text-[1.1875rem] leading-[1.4375rem] font-inter-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div class="flex flex-wrap my-7 px-[30px] pb-[30px]">
                <div
                  v-for="techs in stack.icons"
                  :key="techs"
                  class="flex flex-col items-center my-5 ml-2 w-[30%]"
                >
                  <img
                    class="mb-4 h-[30px] w-[30px]"
                    :src="techs.image"
                    alt="technology-stacks-image"
                  />
                  <span
                    class="font-inter-regular text-[1.125rem] leading-[1.6875rem] text-center"
                    >{{ techs.content }}</span
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="row-start-1 row-span-2 mt-10">
          <div
            v-for="stack in stack2"
            :key="stack"
            class="flex flex-col nowrap"
          >
            <div
              class="mt-16 h-auto w-[368px] rounded-2xl bg-white drop-shadow-md"
            >
              <span
                class="flex items-center mx-auto mt-[-23px] w-max rounded-full py-3 px-5 text-center gradient-btn pointer-events-none"
              >
                <span
                  class="mr-2.5 text-[1.1875rem] leading-[1.4375rem] font-inter-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div class="flex flex-wrap my-7 px-[30px] pb-[30px]">
                <div
                  v-for="techs in stack.icons"
                  :key="techs"
                  class="flex flex-col items-center my-5 ml-2 w-[30%]"
                >
                  <img
                    class="mb-4 h-[30px] w-[30px]"
                    :src="techs.image"
                    alt="technology-stacks-image"
                  />

                  <span
                    class="font-inter-regular text-[1.125rem] leading-[1.6875rem] text-center"
                    >{{ techs.content }}</span
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="row-start-1 row-span-2 mt-64">
          <div
            v-for="stack in stack3"
            :key="stack"
            class="flex flex-col nowrap"
          >
            <div
              class="mt-16 h-auto w-[368px] rounded-2xl bg-white drop-shadow-md"
            >
              <span
                class="flex items-center mx-auto mt-[-23px] w-max rounded-full py-3 px-5 text-center gradient-btn pointer-events-none"
              >
                <span
                  class="mr-2.5 text-[1.1875rem] leading-[1.4375rem] font-inter-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div class="flex flex-wrap my-7 px-[30px] pb-[30px]">
                <div
                  v-for="techs in stack.icons"
                  :key="techs"
                  class="flex flex-col items-center my-5 ml-2 w-[30%]"
                >
                  <img
                    class="mb-4 h-[30px] w-[30px]"
                    :src="techs.image"
                    alt="technology-stacks-image"
                  />

                  <span
                    class="font-inter-regular text-[1.125rem] leading-[1.6875rem] text-center"
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
import frontend1 from "@/assets/images/services/technology/frontend-1-200w.webp";
import frontend2 from "@/assets/images/services/technology/frontend-2-200w.webp";
import frontend3 from "@/assets/images/services/technology/frontend-3-200w.webp";
import frontend4 from "@/assets/images/services/technology/frontend-4-200w.webp";
import frontend5 from "@/assets/images/services/technology/frontend-5-200w.webp";
import frontend6 from "@/assets/images/services/technology/frontend-6-200w.webp";
import frontend7 from "@/assets/images/services/technology/frontend-7-200w.webp";

import desktopfrontend1 from "@/assets/images/services/technology/desktop-icon/frontend-1-200w.webp";
import desktopfrontend2 from "@/assets/images/services/technology/desktop-icon/frontend-2-200w.webp";
import desktopfrontend3 from "@/assets/images/services/technology/desktop-icon/frontend-3-200w.webp";
import desktopfrontend4 from "@/assets/images/services/technology/desktop-icon/frontend-4-200w.webp";
import desktopfrontend5 from "@/assets/images/services/technology/desktop-icon/frontend-5-200w.webp";
import desktopfrontend6 from "@/assets/images/services/technology/desktop-icon/frontend-6-200w.webp";
import desktopfrontend7 from "@/assets/images/services/technology/desktop-icon/frontend-7-200w.webp";

import backend1 from "@/assets/images/services/technology/backend-1-200w.webp";
import backend2 from "@/assets/images/services/technology/backend-2-200w.webp";
import backend3 from "@/assets/images/services/technology/backend-3-200w.webp";
import backend4 from "@/assets/images/services/technology/backend-4-200w.webp";
import backend5 from "@/assets/images/services/technology/backend-5-200w.webp";

import desktopbackend1 from "@/assets/images/services/technology/desktop-icon/backend-1-200w.webp";
import desktopbackend2 from "@/assets/images/services/technology/desktop-icon/backend-2-200w.webp";
import desktopbackend3 from "@/assets/images/services/technology/desktop-icon/backend-3-200w.webp";
import desktopbackend4 from "@/assets/images/services/technology/desktop-icon/backend-4-200w.webp";
import desktopbackend5 from "@/assets/images/services/technology/desktop-icon/backend-5-200w.webp";

import mobileapp1 from "@/assets/images/services/technology/mobileapp-1-200w.webp";
import mobileapp2 from "@/assets/images/services/technology/mobileapp-2-200w.webp";
import mobileapp3 from "@/assets/images/services/technology/mobileapp-3-200w.webp";
import mobileapp4 from "@/assets/images/services/technology/mobileapp-4-200w.webp";

import desktopmobileapp1 from "@/assets/images/services/technology/desktop-icon/mobileapp-1-200w.webp";
import desktopmobileapp2 from "@/assets/images/services/technology/desktop-icon/mobileapp-2-200w.webp";
import desktopmobileapp3 from "@/assets/images/services/technology/desktop-icon/mobileapp-3-200w.webp";
import desktopmobileapp4 from "@/assets/images/services/technology/desktop-icon/mobileapp-4-200w.webp";

import ui1 from "@/assets/images/services/technology/ui-1-200w.webp";
import ui2 from "@/assets/images/services/technology/ui-2-200w.webp";
import ui3 from "@/assets/images/services/technology/ui-3-200w.webp";
import ui4 from "@/assets/images/services/technology/ui-4-200w.webp";
import ui5 from "@/assets/images/services/technology/ui-5-200w.webp";

import desktopui1 from "@/assets/images/services/technology/desktop-icon/ui-1-200w.webp";
import desktopui2 from "@/assets/images/services/technology/desktop-icon/ui-2-200w.webp";
import desktopui3 from "@/assets/images/services/technology/desktop-icon/ui-3-200w.webp";
import desktopui4 from "@/assets/images/services/technology/desktop-icon/ui-4-200w.webp";
import desktopui5 from "@/assets/images/services/technology/desktop-icon/ui-5-200w.webp";

import devOps1 from "@/assets/images/services/technology/devops-1-200w.webp";
import devOps2 from "@/assets/images/services/technology/devops-2-200w.webp";
import devOps3 from "@/assets/images/services/technology/devops-3-200w.webp";
import devOps4 from "@/assets/images/services/technology/devops-4-200w.webp";
import devOps5 from "@/assets/images/services/technology/devops-5-200w.webp";
import devOps6 from "@/assets/images/services/technology/devops-6-200w.webp";

import desktopDevOps1 from "@/assets/images/services/technology/desktop-icon/devops-1-200w.webp";
import desktopDevOps2 from "@/assets/images/services/technology/desktop-icon/devops-2-200w.webp";
import desktopDevOps3 from "@/assets/images/services/technology/desktop-icon/devops-3-200w.webp";
import desktopDevOps4 from "@/assets/images/services/technology/desktop-icon/devops-4-200w.webp";
import desktopDevOps5 from "@/assets/images/services/technology/desktop-icon/devops-5-200w.webp";
import desktopDevOps6 from "@/assets/images/services/technology/desktop-icon/devops-6-200w.webp";

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
          title: "DevOps",
          image: [devOps1, devOps2, devOps3, devOps4, devOps5, devOps6],
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
          title: "DevOps",
          icons: [
            { image: desktopDevOps1, content: "AWS" },
            { image: desktopDevOps2, content: "GCP" },
            { image: desktopDevOps3, content: "Firebase" },
            { image: desktopDevOps4, content: "Docker" },
            { image: desktopDevOps5, content: "Git" },
            { image: desktopDevOps6, content: "CI/CD" },
          ],
        },
      ],
    };
  },
};
</script>
