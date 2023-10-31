<template>
  <section class="my-8 md:mb-20 md:mt-32">
    <div class="mb-4 md:container">
      <div class="mx-[2%] my-0 flex flex-col gap-4 md:my-0">
        <div class="text-center">
          <span
            class="font-inter-bold text-3xl leading-[2.4375rem] text-black-core/[0.87] md:text-[2.65625rem] lg:text-[3.4375rem] md:leading-[3.453125rem] lg:leading-[4.46875rem]"
            >FAQs on iOS App Development</span
          >
        </div>

        <div
          class="container mt-6 flex-[67%] pb-8 md:pb-16 md:w-[65rem] md:-ml-6 xl:mx-auto"
        >
          <div>
            <transition-group tag="div" :name="'faq-' + faqTransitionName">
              <div class="faq-section h-auto">
                <div
                  class="mt-4"
                  v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                  :key="faq"
                >
                  <div
                    class="faq-container cursor-pointer overflow-hidden rounded-[5px] bg-white md:p-[15px] md:shadow-md"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="faq-header flex flex-row items-center">
                      <div
                        class="faq-question w-[90%] p-[10px] font-inter-semibold tracking-[-0.04rem] text-lg leading-[1.6875rem] md:text-2xl lg:text-[1.75rem] md:leading-8 lg:leading-[2.625rem]"
                        :class="
                          openList && faq.id == currentIndex
                            ? 'v2-canopas-gradient-text footer-icon'
                            : 'text-black-core/[0.60] md:text-black-core/[0.80]'
                        "
                      >
                        {{ faq.question }}
                      </div>
                      <div
                        class="faq-icon w-[10%] text-end"
                        @click.native="
                          mixpanel.track('tap_ios_app_faq_questions')
                        "
                      >
                        <Icon
                          class="xmark text-black-core/[0.87]"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'footer-icon'
                              : 'text-black-core/[0.87]'
                          "
                          :name="
                            openList && faq.id == currentIndex
                              ? 'fa6-solid:angle-up'
                              : 'fa6-solid:angle-down'
                          "
                        />
                      </div>
                    </div>

                    <collapse-transition>
                      <div
                        class="faq-header mt-1.5 flex flex-row animate-fadeIn items-center gradient-border h-auto border-l md:border-l-4 ml-3"
                        v-if="openList && faq.id == currentIndex"
                        :key="faq.answer"
                      >
                        <div
                          class="faq-answer w-[90%] px-3 font-inter-regular text-base leading-normal text-black-core/[0.87] md:text-xl md:leading-[1.875rem]"
                          v-html="faq.answer"
                        ></div>
                      </div>
                    </collapse-transition>
                    <div class="faq-divider" v-if="faq.id != faq.length"></div>
                  </div>
                </div>
              </div>
            </transition-group>
          </div>
        </div>
      </div>
      <div
        class="container py-4 flex justify-center items-center md:justify-end pr-4 text-center md:text-right md:pr-6 md:pt-4 xl:pr-20 2xl:pr-44 md:-mr-8 lg:-mr-10"
      >
        <button
          v-if="isActivePrev"
          type="button"
          :disabled="!isActivePrev"
          class="clients-indicators mx-1 my-0 cursor-pointer sm:mx-2"
          @click="slide(-1)"
          @click.native="$mixpanel.track('tap_ios_app_faq_previous_arrow')"
          aria-label="leftArrow"
        >
          <Icon
            :class="isActivePrev ? 'footer-icon' : ''"
            class="arrow h-3.5 w-3.5 rounded-full border-[#3d3d3d26] bg-white p-2.5 drop-shadow-md md:h-5 md:w-5"
            name="fa6-solid:angle-left"
            id="leftArrow"
          />
        </button>
        <button
          v-else
          type="button"
          :disabled="!isActiveNext"
          class="clients-indicators flex items-center justify-center my-0 cursor-pointer"
          @click="slide(1)"
          @click.native="$mixpanel.track('tap_ios_app_faq_next_arrow')"
          aria-label="rightArrow"
        >
          <div
            :disabled="!isActiveNext"
            class="faq-question v2-canopas-gradient-text w-full font-inter-semibold text-base leading-[1.21rem] md:text-2xl lg:text-3xl md:leading-[1.3234375rem] lg:leading-[1.436875rem]"
          >
            Load More Questions
          </div>
          <Icon
            :class="isActiveNext ? 'footer-icon' : ''"
            class="arrow h-3.5 w-3.5 rounded-full bg-white p-2.5 md:h-5 md:w-5"
            name="fa6-solid:angle-right"
            id="rightArrow"
          />
        </button>
      </div>
    </div>
  </section>
</template>

<script type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
      faqs: [
        {
          id: 1,
          question:
            "What programming languages do you use for iOS app development?",
          answer:
            "Our primary languages for iOS app development are Swift and Objective-C.",
        },
        {
          id: 2,
          question: "How long will it take to develop my iOS app?",
          answer:
            "The development time for an iOS app can vary greatly based on the complexity of the app, the number of features, and your specific requirements. We can provide a more accurate estimate after discussing your project.",
        },
        {
          id: 3,
          question: "Can you build an app that works on both iOS and Android?",
          answer:
            "Yes, we also provide cross-platform app development services that can create apps for both iOS and Android.",
        },
        {
          id: 4,
          question: "Do you follow Apple's design guidelines?",
          answer:
            "Yes,we adhere strictly to Apple's Human Interface Guidelines to ensure the best user experience and facilitate approval on the App Store.",
        },
        {
          id: 5,
          question:
            "How do I track the progress of my app development project?",
          answer:
            "We will provide regular updates throughout the project, and we're always available to discuss any questions or concerns you might have.",
        },
        {
          id: 6,
          question: "What happens if I find a bug after the app is launched?",
          answer:
            "We offer post-launch support, during which we will fix any bugs free of charge. After this period, we offer maintenance and support services.",
        },
        {
          id: 7,
          question: "Can you help launch my app on the App Store?",
          answer:
            "Yes, we can guide you through the whole process of launching your app on the App Store, including setting up all necessary configurations and ensuring your app complies with all of Apple's guidelines.",
        },
        {
          id: 8,
          question: "Can you integrate my iOS app with other systems or APIs?",
          answer:
            "Absolutely, we have extensive experience in integrating iOS apps with various systems, databases, and APIs.",
        },
        {
          id: 9,
          question: "Will my app be optimized for all iOS devices?",
          answer:
            "Yes, we will  ensure your app works well on all supported iPhone and iPad models.",
        },
        {
          id: 10,
          question: "Can my app be updated or modified after it's launched?",
          answer:
            "Yes, we offer app update and modification services to keep your app updated with the latest iOS features or any changes you require.",
        },
      ],
      currentIndex: 0,
      previousIndex: 0,
      openList: true,
      current: 0,
      faqTransitionName: null,
      sliceFrom: 0,
      sliceTo: TOTAL_FAQ_IN_SLIDE,
      isActivePrev: false,
      isActiveNext: true,
    };
  },
  components: {
    CollapseTransition,
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll, { passive: false });
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll, { passive: false });
  },
  inject: ["mixpanel"],
  methods: {
    expandListItem(index) {
      if (this.previousIndex == index && this.openList) {
        this.openList = false;
        this.$mixpanel.track("tap_ios_app_faq_collapse");
      } else {
        this.openList = true;
        this.$mixpanel.track("tap_ios_app_faq_expand");
      }

      this.currentIndex = index;
      this.previousIndex = this.currentIndex;
    },
    slide(dir) {
      this.faqTransitionName = dir === 1 ? "next" : "prev";
      this.sliceFrom =
        dir === 1
          ? this.sliceFrom + TOTAL_FAQ_IN_SLIDE
          : this.sliceFrom - TOTAL_FAQ_IN_SLIDE;
      this.sliceTo =
        dir === 1
          ? this.sliceTo + TOTAL_FAQ_IN_SLIDE
          : this.sliceTo - TOTAL_FAQ_IN_SLIDE;

      this.isActivePrev = this.sliceFrom >= TOTAL_FAQ_IN_SLIDE;
      this.isActiveNext = this.sliceTo < this.faqs.length;
      this.current = this.getRoundedIndex(dir);
    },
    getRoundedIndex(diff) {
      var len = this.faqs.length / TOTAL_FAQ_IN_SLIDE;
      return (
        (this.current + (diff % Math.ceil(len)) + Math.ceil(len)) %
        Math.ceil(len)
      );
    },
  },
};
</script>
<style scoped>
.gradient-border {
  border-image: linear-gradient(180deg, #f2709c 50%, #ff835b 50%) 1;
}
</style>
