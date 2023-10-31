<template>
  <div
    class="mt-[7.813rem] mb-11 md:mb-24 bg-gradient-to-t from-orange-300/[0.15] to-white"
  >
    <div class="container pb-4 pt-20 lg:pt-32">
      <div
        class="m-auto sm:w-[90%] text-center font-inter-bold text-[1.875rem] lg:text-[3.4375rem] leading-[2.5rem] lg:leading-[5.875rem] tracking-[-4%] text-black-900"
      >
        With <span class="canopas-gradient-text">canopas</span>, you will have
        everything you need under one roof.
      </div>
      <div>
        <div
          ref="service1"
          class="flex flex-wrap m-auto md:mt-20 md:mx-auto md:mb-auto lg:even:w-[70%]"
        >
          <div
            v-for="(phase, i) in phases.slice(0, 3)"
            :key="i"
            class="md:flex-[1_1_30%] md:flex md:flex-col py-4 px-0 md:p-0 md:text-left last-of-type:md:py-0 last-of-type:md:px-40 last-of-type:lg:p-0 text-center"
          >
            <aspect-ratio height="56.25%">
              <img
                :src="phase.image"
                class="w-full h-full object-cover"
                loading="lazy"
                :alt="phase.alt"
              />
            </aspect-ratio>
            <div
              class="md:-mt-5 p-4 sm:py-4 sm:px-12 md:py-4 md:pr-12 md:pl-6 lg:py-4 lg:px-6"
              :class="phase.className"
            >
              <div
                class="bg-gradient-to-r from-pink-300 to-orange-300 p-2 bg-clip-text text-transparent font-inter-medium text-[1.5rem] md:text-[2.1875rem] leading-[1.95rem] md:leading-[4.28rem] tracking-[1%]"
              >
                {{ phase.title }}
              </div>
              <div
                class="md:p-2 font-inter-regular text-[1rem] md:text-[1.375rem] md:leading-8 text-black-900"
              >
                {{ phase.description }}
              </div>
            </div>
          </div>
        </div>
        <div
          ref="service2"
          class="flex flex-wrap m-auto md:mt-20 md:mx-auto md:mb-auto lg:even:w-[70%]"
        >
          <div
            v-for="(phase, i) in phases.slice(3, 5)"
            :key="i"
            class="text-center py-4 px-0 md:flex-[1_1_30%] md:flex md:flex-col md:p-0 md:text-left"
          >
            <aspect-ratio height="56.25%">
              <img
                :src="phase.image"
                class="w-full h-full object-cover"
                :class="phase.title == 'Marketing' ? 'h-full md:h-[84%]' : ''"
                loading="lazy"
                :alt="phase.alt"
              />
            </aspect-ratio>
            <div
              class="md:-mt-10 p-4 sm:py-4 sm:px-12 md:py-4 md:pr-12 md:pl-6 lg:py-4 lg:px-6"
              :class="phase.className"
            >
              <div
                class="bg-gradient-to-r from-pink-300 to-orange-300 p-2 bg-clip-text text-transparent font-inter-medium text-[1.5rem] md:text-[2.1875rem] leading-[1.95rem] md:leading-[4.28rem] tracking-[1%]"
              >
                {{ phase.title }}
              </div>
              <div
                class="md:p-2 font-inter-regular text-[1rem] md:text-[1.375rem] md:leading-8 text-black-900"
              >
                {{ phase.description }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import AspectRatio from "@/components/utils/AspectRatio.vue";
import planning from "@/assets/images/phases/canopas_phases_planning.gif";
import designing from "@/assets/images/phases/canopas_phases_designing.gif";
import development from "@/assets/images/phases/canopas_phases_development.gif";
import marketing from "@/assets/images/phases/canopas_phases_marketing.gif";
import support from "@/assets/images/phases/canopas_phases_support.gif";

export default {
  data() {
    return {
      phases: [
        {
          image: planning,
          alt: "planning",
          title: "Planning",
          description:
            "We research to understand the vision. The audience. The product. The goals. We collaborate, we re-examine, we ask and conclude.",
          className: "",
        },
        {
          image: designing,
          alt: "designing",
          title: "Design",
          description:
            "Working passionately with you, our designers create aesthetically pleasing  and user friendly digital experiences.",
          className: "md:ml-[17px] 2xl:ml-[25px]",
        },
        {
          image: development,
          alt: "development",
          title: "Development",
          description:
            "Backend and Frontend. From the beginning, we have performance and maintainability in mind. Always insisting on best practices and keeping up with the latest trends.",
          className: "md:ml-[35px] 2xl:ml-[45px]",
        },
        {
          image: marketing,
          alt: "marketing",
          title: "Marketing",
          description:
            "From organic growth to writing irresistible ads for your marketing campaign, we're here to help you. We'll help you build a brand and community of prospects who can give you a boost on your product launch.",
          className: "md:ml-[28px] 2xl:ml-[38px]",
        },
        {
          image: support,
          alt: "support",
          title: "Support",
          description:
            "From bug fixing to feature updates, you can count on us. For the last 7 years, our clients had a 99.87% uptime for their products.",
          className: "-mt-[15px] md:ml-[25px] 2xl:ml-[35px]",
        },
      ],
    };
  },
  components: {
    AspectRatio,
  },
};
</script>
