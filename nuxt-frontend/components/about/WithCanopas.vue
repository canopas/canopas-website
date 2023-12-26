<template>
  <section class="3xl:outer-container">
    <div
      class="mt-16 lg:mt-[7.5rem] mb-11 lg:mb-24 bg-gradient-to-t from-orange-5 to-white"
    >
      <div class="container pb-14">
        <div
          class="m-auto sm:w-[90%] text-center mobile-header-2 lg:desk-header-2 text-black-87"
        >
          With <span class="primary-color">canopas</span>, you will have
          everything you need under one roof.
        </div>
        <div>
          <div
            ref="service1"
            class="flex flex-wrap m-auto mt-6 lg:mt-20 lg:mx-auto lg:mb-auto lg:even:w-[70%]"
          >
            <div
              v-for="(phase, i) in phases.slice(0, 3)"
              :key="i"
              class="lg:flex-[1_1_30%] lg:flex lg:flex-col py-4 px-0 lg:p-0 lg:text-left last-of-type:lg:p-0 text-center"
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
                class="lg:-mt-5 p-4 sm:px-12 lg:py-4 lg:pl-[1.2rem] xl:pl-[1.9rem] lg:py-4 lg:px-6"
                :class="phase.className"
              >
                <div
                  class="mobile-header-3-semibold lg:desk-header-3 primary-color"
                >
                  {{ phase.title }}
                </div>
                <div
                  class="sub-h3-regular mt-2 lg:mt-3 lg:mobile-header-3-medium text-black-60"
                >
                  {{ phase.description }}
                </div>
              </div>
            </div>
          </div>
          <div
            ref="service2"
            class="flex flex-wrap m-auto lg:mt-20 lg:mx-auto lg:mb-auto lg:even:w-[85%]"
          >
            <div
              v-for="(phase, i) in phases.slice(3, 5)"
              :key="i"
              class="text-center py-4 px-0 lg:flex-[1_1_30%] lg:flex lg:flex-col lg:p-0 lg:text-left"
            >
              <aspect-ratio height="56.25%">
                <img
                  :src="phase.image"
                  class="w-full h-full object-cover"
                  :class="phase.title == 'Marketing' ? 'h-full lg:h-[84%]' : ''"
                  loading="lazy"
                  :alt="phase.alt"
                />
              </aspect-ratio>
              <div
                class="lg:-mt-[3.25rem] p-4 sm:px-12 lg:pl-[1.9rem] 2xl:pl-10 lg:px-6"
                :class="phase.className"
              >
                <div
                  class="mobile-header-3-semibold lg:desk-header-3 primary-color"
                >
                  {{ phase.title }}
                </div>
                <div
                  class="sub-h3-regular mt-2 lg:mt-3 lg:mobile-header-3-medium text-black-60"
                >
                  {{ phase.description }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
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
          className: "lg:ml-[17px] 2xl:ml-6",
        },
        {
          image: development,
          alt: "development",
          title: "Development",
          description:
            "Backend and Frontend. From the beginning, we have performance and maintainability in mind. Always insisting on best practices and keeping up with the latest trends.",
          className: "lg:ml-[35px] xl:ml-[38px] 2xl:ml-[50px]",
        },
        {
          image: marketing,
          alt: "marketing",
          title: "Marketing",
          description:
            "From organic growth to writing irresistible ads for your marketing campaign, we're here to help you. We'll help you build a brand and community of prospects who can give you a boost on your product launch.",
          className: "lg:ml-7 xl:ml-[38px]",
        },
        {
          image: support,
          alt: "support",
          title: "Support",
          description:
            "From bug fixing to feature updates, you can count on us. For the last 7 years, our clients had a 99.87% uptime for their products.",
          className: "-mt-[15px] lg:ml-[25px] xl:ml-[38px]",
        },
      ],
    };
  },
  components: {
    AspectRatio,
  },
};
</script>
