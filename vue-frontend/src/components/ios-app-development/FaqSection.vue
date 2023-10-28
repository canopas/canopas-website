<template>
  <section class="tw-my-8 sm:tw-mt-16 md:tw-mb-20 lg:tw-my-60">
    <div class="tw-mb-4 lg:tw-container">
      <div class="tw-mx-[2%] tw-my-0 tw-flex tw-flex-col tw-gap-4 lg:tw-my-0">
        <div class="tw-text-center">
          <span class="header-2 tw-text-black-core/[0.87]"
            >FAQs on iOS app development</span
          >
        </div>

        <div
          class="tw-container tw-mt-6 tw-flex-[67%] tw-pb-8 lg:tw-pb-16 lg:tw-w-[65rem] lg:tw--ml-6 xl:tw-mx-auto"
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
                    class="faq-container tw-cursor-pointer tw-overflow-hidden tw-rounded-[5px] tw-bg-white lg:tw-p-[15px] lg:tw-shadow-md"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="faq-header tw-flex tw-flex-row tw-items-center">
                      <div
                        class="faq-question tw-w-[90%] tw-p-[10px] sub-h1-semibold"
                        :class="
                          openList && faq.id == currentIndex
                            ? 'v2-canopas-gradient-text footer-icon'
                            : 'tw-text-black-core/[0.60] lg:tw-text-black-core/[0.80]'
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
                          class="faq-answer tw-w-[90%] tw-px-3 sub-h3-regular"
                          v-html="faq.answer"
                        ></div>
                      </div>
                    </collapse-transition>
                    <div class="faq-divider" v-if="faq.id != faq.length"></div>
                  </div>
                </div>
              </div>
            </transition-group>
            <collapse-transition>
              <transition-group
                tag="div"
                :name="'faq-' + faqTransitionName"
                v-if="showAdditionalFAQs"
              >
                <div class="faq-section tw-h-auto lg:tw-hidden">
                  <div
                    class="tw-mt-3"
                    v-for="faq in faqs.slice(5, 10)"
                    :key="faq"
                  >
                    <div
                      class="faq-container tw-cursor-pointer tw-overflow-hidden tw-rounded-[5px] tw-bg-white lg:tw-p-[15px] lg:tw-shadow-md"
                      @click="expandListItem(faq.id)"
                    >
                      <div
                        class="faq-header tw-flex tw-flex-row tw-items-center"
                      >
                        <div
                          class="faq-question tw-w-[90%] tw-p-[10px] sub-h1-semibold"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'v2-canopas-gradient-text footer-icon'
                              : 'tw-text-black-core/[0.60] lg:tw-text-black-core/[0.80]'
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
                            class="faq-answer tw-w-[90%] tw-px-3 sub-h3-regular"
                            v-html="faq.answer"
                          ></div>
                        </div>
                      </collapse-transition>
                      <div
                        class="faq-divider"
                        v-if="faq.id != faq.length"
                      ></div>
                    </div>
                  </div>
                </div>
              </transition-group>
            </collapse-transition>
            <div
              class="tw-w-full tw-my-8 lg:tw-hidden tw-mt-8 tw-flex tw-justify-center tw-items-center"
            >
              <div
                class="tw-m-[5px] tw-rounded-[0.6rem] tw-border tw-border-solid tw-border-transparent tw-bg-gradient-to-r tw-from-[#ff9472] tw-via-[#ff909c] tw-to-[#f2709c] tw-p-[1rem] tw-text-center tw-shadow-[inset_2px_1000px_1px_#fff] active:tw-scale-[0.98] tw-border tw-border-black tw-w-[155px] tw-cursor-pointer tw-h-[57px]"
                @click="showAdditionalFAQs = !showAdditionalFAQs"
              >
                <span class="sub-h3-semibold v2-canopas-gradient-text"
                  >{{ showAdditionalFAQs ? "View Less" : "View More" }}
                </span>
                <font-awesome-icon
                  class="fa tw-ml-2 tw-mt-1 footer-icon"
                  :icon="showAdditionalFAQs ? faAngleUp : faAngleDown"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="tw-hidden lg:tw-block tw-container tw--mt-8 tw-text-right lg:tw-pt-4 xl:tw-pr-16 2xl:tw-pr-40"
    >
      <button
        v-if="isActivePrev"
        type="button"
        :disabled="!isActivePrev"
        class="tw-h-12 tw-flex tw-ml-auto tw-items-center tw-text-center"
        @click="slide(-1)"
        @click.native="mixpanel.track('tap_ios_app_faq_previous_arrow')"
        aria-label="leftArrow"
      >
        <font-awesome-icon
          :class="isActivePrev ? 'footer-icon' : ''"
          class="tw-mr-2 tw-h-5 tw-w-5"
          icon="arrow-left"
          id="leftArrow"
        /><span class="v2-canopas-gradient-text sub-h4-semibold">Back</span>
      </button>
      <button
        v-else
        type="button"
        class="tw-h-12 tw-flex tw-ml-auto tw-items-center tw-text-center"
        :disabled="!isActiveNext"
        @click="slide(1)"
        @click.native="mixpanel.track('tap_ios_app_faq_next_arrow')"
        aria-label="rightArrow"
      >
        <span class="v2-canopas-gradient-text sub-h4-semibold"
          >View More Questions</span
        >
        <font-awesome-icon
          :class="isActiveNext ? 'footer-icon' : ''"
          class="tw-ml-2 tw-h-5 tw-w-5"
          icon="arrow-right"
          id="rightArrow"
        />
      </button>
    </div>
  </section>
</template>

<script type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faAngleUp, faAngleDown } from "@fortawesome/free-solid-svg-icons";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
      faAngleUp,
      faAngleDown,
      showAdditionalFAQs: false,
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
<style lang="postcss">
.gradient-border {
  border-image: linear-gradient(180deg, #f2709c 50%, #ff835b 50%) 1;
}
</style>
