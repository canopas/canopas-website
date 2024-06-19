<template>
  <section class="mt-16 lg:mt-60 flex flex-col items-center md:mx-auto">
    <div class="container text-center">
      <span class="mobile-header-2 lg:desk-header-2 text-black-87"
        >Technology Stack</span
      >
      <p
        class="mt-4 sub-h1-regular lg:mobile-header-2-regular text-black-60 w-full xl:w-[91%] 2xl:w-[81%] text-center mx-auto"
      >
        The products we build are typically deployed to mobile and web. Web and
        mobile are our bread and butter. From native to hybrid, frontend to
        backend, designing to deployment, we've got you covered. At Canopas,
        we're not bound by technologies. We learn new technology if it has
        chances to improve your app performance significantly.
      </p>
    </div>
    <div
      class="lg:hidden mt-16 h-auto w-[380px] rounded-2xl bg-white-smoke"
      v-for="stack in stacks"
      :key="stack"
    >
      <span
        class="flex items-center -mt-6 mx-auto w-max rounded-xl py-2.5 md:py-3 px-10 text-center bg-black-core pointer-events-none"
      >
        <span class="mobile-header-3 text-white">{{ stack.title }} </span>
      </span>
      <div class="flex flex-row flex-wrap items-center my-7 pl-7">
        <img
          v-for="techs in stack.image"
          :key="techs"
          class="flex w-28 h-28 object-cover mt-4 rounded-[0.313rem] pr-2.5"
          :src="techs"
          alt="technology-stacks-image"
        />
      </div>
    </div>
    <div class="hidden lg:block">
      <div class="grid grid-rows-2 grid-flow-col gap-x-4 xl:gap-x-6">
        <div class="row-start-2 row-span-1 -mt-40">
          <div v-for="stack in stack1" :key="stack" class="flex">
            <div
              class="mt-14 h-auto w-[313px] xl:w-[345px] 2xl:w-[368px] rounded-2xl bg-white drop-shadow-md"
            >
              <span
                class="flex items-center mx-auto mt-[-1.438rem] w-max rounded-xl py-3 px-5 text-center bg-black-core pointer-events-none"
              >
                <span class="text-white mobile-header-2-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div class="flex flex-wrap mt-7 px-1 xl:px-5 pb-5">
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
                  <span class="sub-h1-regular text-black-87 text-center">{{
                    techs.content
                  }}</span>
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
              class="mt-16 h-auto w-[313px] xl:w-[345px] 2xl:w-[368px] rounded-2xl bg-white drop-shadow-md"
            >
              <span
                class="flex items-center mx-auto mt-[-1.438rem] w-max rounded-xl py-3 px-5 text-center bg-black-core pointer-events-none"
              >
                <span class="text-white mobile-header-2-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div class="flex flex-wrap mt-7 px-1 xl:px-5 pb-5">
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

                  <span class="sub-h1-regular text-black-87 text-center">{{
                    techs.content
                  }}</span>
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
              class="mt-16 h-auto w-[313px] xl:w-[345px] 2xl:w-[368px] rounded-2xl bg-white drop-shadow-md"
            >
              <span
                class="flex items-center mx-auto mt-[-1.438rem] w-max rounded-xl py-3 px-5 text-center bg-black-core pointer-events-none"
              >
                <span class="text-white mobile-header-2-semibold"
                  >{{ stack.title }}
                </span>
              </span>
              <div class="flex flex-wrap mt-7 px-1 xl:px-5 pb-5">
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

                  <span class="sub-h1-regular text-black-87 text-center">{{
                    techs.content
                  }}</span>
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
import frontend8 from "@/assets/images/services/technology/frontend-8-200w.webp";

import desktopfrontend1 from "@/assets/images/services/technology/desktop-icon/frontend-1-200w.webp";
import desktopfrontend2 from "@/assets/images/services/technology/desktop-icon/frontend-2-200w.webp";
import desktopfrontend3 from "@/assets/images/services/technology/desktop-icon/frontend-3-200w.webp";
import desktopfrontend4 from "@/assets/images/services/technology/desktop-icon/frontend-4-200w.webp";
import desktopfrontend5 from "@/assets/images/services/technology/desktop-icon/frontend-5-200w.webp";
import desktopfrontend6 from "@/assets/images/services/technology/desktop-icon/frontend-6-200w.webp";
import desktopfrontend7 from "@/assets/images/services/technology/desktop-icon/frontend-7-200w.webp";
import desktopfrontend8 from "@/assets/images/services/technology/desktop-icon/frontend-8-200w.webp";

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
            frontend8,
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
            { image: desktopfrontend8, content: "Angular" },
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
