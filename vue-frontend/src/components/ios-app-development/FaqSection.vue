<template>
  <section class="tw-my-8 md:tw-mb-20 md:tw-mt-32">
    <div class="tw-mb-4 md:tw-container">
      <div class="tw-mx-[2%] tw-my-0 tw-flex tw-flex-col tw-gap-4 md:tw-my-0">
        <div class="tw-text-center">
          <span
            class="tw-font-inter-bold tw-text-3xl tw-leading-[2.4375rem] tw-text-black-core/[0.87] md:tw-text-[2.65625rem] lg:tw-text-[3.4375rem] md:tw-leading-[3.453125rem] lg:tw-leading-[4.46875rem]"
            >FAQs on iOS App Development</span
          >
        </div>

        <div
          class="tw-container tw-mt-6 tw-flex-[67%] tw-pb-8 md:tw-pb-16 md:tw-w-[65rem] md:tw--ml-6 xl:tw-mx-auto"
        >
          <div>
            <transition-group tag="div" :name="'faq-' + faqTransitionName">
              <div class="faq-section tw-h-auto">
                <div
                  class="tw-mt-4"
                  v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                  :key="faq"
                >
                  <div
                    class="faq-container tw-cursor-pointer tw-overflow-hidden tw-rounded-[5px] tw-bg-white md:tw-p-[15px] md:tw-shadow-md"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="faq-header tw-flex tw-flex-row tw-items-center">
                      <div
                        class="faq-question tw-w-[90%] tw-p-[10px] tw-font-inter-semibold tw-tracking-[-0.04rem] tw-text-lg tw-leading-[1.6875rem] md:tw-text-2xl lg:tw-text-[1.75rem] md:tw-leading-8 lg:tw-leading-[2.625rem]"
                        :class="
                          openList && faq.id == currentIndex
                            ? 'v2-canopas-gradient-text footer-icon'
                            : 'tw-text-black-core/[0.60] md:tw-text-black-core/[0.80]'
                        "
                      >
                        {{ faq.question }}
                      </div>
                      <div
                        class="faq-icon tw-w-[10%] tw-text-end"
                        @click.native="
                          mixpanel.track('tap_ios_app_faq_questions')
                        "
                      >
                        <font-awesome-icon
                          class="xmark tw-text-black-core/[0.87]"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'footer-icon'
                              : 'tw-text-black-core/[0.87]'
                          "
                          :icon="
                            openList && faq.id == currentIndex
                              ? 'angle-up'
                              : 'angle-down'
                          "
                        />
                      </div>
                    </div>

                    <collapse-transition>
                      <div
                        class="faq-header tw-mt-1.5 tw-flex tw-flex-row tw-animate-fadeIn tw-items-center gradient-border tw-h-auto tw-border-l md:tw-border-l-4 tw-ml-3"
                        v-if="openList && faq.id == currentIndex"
                        :key="faq.answer"
                      >
                        <div
                          class="faq-answer tw-w-[90%] tw-px-3 tw-font-inter-regular tw-text-base tw-leading-normal tw-text-black-core/[0.87] md:tw-text-xl md:tw-leading-[1.875rem]"
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
        class="tw-container tw-py-4 tw-flex tw-justify-center tw-items-center md:tw-justify-end tw-pr-4 tw-text-center md:tw-text-right md:tw-pr-6 md:tw-pt-4 xl:tw-pr-20 2xl:tw-pr-44 md:tw--mr-8 lg:tw--mr-10"
      >
        <button
          v-if="isActivePrev"
          type="button"
          :disabled="!isActivePrev"
          class="clients-indicators tw-mx-1 tw-my-0 tw-cursor-pointer sm:tw-mx-2"
          @click="slide(-1)"
          @click.native="mixpanel.track('tap_ios_app_faq_previous_arrow')"
          aria-label="leftArrow"
        >
          <font-awesome-icon
            :class="isActivePrev ? 'footer-icon' : ''"
            class="arrow tw-h-3.5 tw-w-3.5 tw-rounded-full tw-border-[#3d3d3d26] tw-bg-white tw-p-2.5 tw-drop-shadow-md md:tw-h-5 md:tw-w-5"
            icon="fa-angle-left"
            id="leftArrow"
          />
        </button>
        <button
          v-else
          type="button"
          :disabled="!isActiveNext"
          class="clients-indicators tw-flex tw-items-center tw-justify-center tw-my-0 tw-cursor-pointer"
          @click="slide(1)"
          @click.native="mixpanel.track('tap_ios_app_faq_next_arrow')"
          aria-label="rightArrow"
        >
          <div
            :disabled="!isActiveNext"
            class="faq-question v2-canopas-gradient-text tw-w-full tw-font-inter-semibold tw-text-base tw-leading-[1.21rem] md:tw-text-2xl lg:tw-text-3xl md:tw-leading-[1.3234375rem] lg:tw-leading-[1.436875rem]"
          >
            Load More Questions
          </div>
          <font-awesome-icon
            :class="isActiveNext ? 'footer-icon' : ''"
            class="arrow tw-h-3.5 tw-w-3.5 tw-rounded-full tw-bg-white tw-p-2.5 md:tw-h-5 md:tw-w-5"
            icon="fa-angle-right"
            id="rightArrow"
          />
        </button>
      </div>
    </div>
  </section>
</template>

<script type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

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
    FontAwesomeIcon,
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
        this.mixpanel.track("tap_ios_app_faq_collapse");
      } else {
        this.openList = true;
        this.mixpanel.track("tap_ios_app_faq_expand");
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
