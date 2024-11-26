<template>
  <section class="sm:mt-16 md:mb-20 lg:my-60">
    <div class="lg:container">
      <div class="mx-[2%] my-0 flex flex-col gap-4 lg:my-0">
        <div class="text-center">
          <span class="mobile-header-2 lg:desk-header-2 text-black-core/[0.87]"
            >FAQs on iOS app development</span
          >
        </div>

        <div
          class="container -mt-8 flex-[67%] lg:pb-16 lg:w-[65rem] lg:-ml-6 xl:mx-auto"
        >
          <div>
            <transition-group tag="div" :name="'faq-' + faqTransitionName">
              <div class="faq-section mt-12 h-auto">
                <div
                  class="mt-4"
                  v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                  :key="faq"
                >
                  <div
                    class="faq-container cursor-pointer overflow-hidden rounded-[5px] bg-white lg:p-[15px] lg:shadow-md"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="faq-header flex flex-row items-center">
                      <div
                        class="faq-question w-[90%] p-[10px] sub-h1-semibold lg:mobile-header-2-semibold"
                        :class="
                          openList && faq.id == currentIndex
                            ? 'v2-canopas-gradient-text footer-icon'
                            : 'text-black-60 lg:text-black-87'
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
                          class="angle-up text-black-87"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'footer-icon'
                              : 'text-black-87'
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
                          class="faq-answer w-[90%] px-3 sub-h3-regular lg:mobile-header-3-regular"
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
                <div class="faq-section h-auto lg:hidden">
                  <div class="mt-3" v-for="faq in faqs.slice(5, 10)" :key="faq">
                    <div
                      class="faq-container cursor-pointer overflow-hidden rounded-[5px] bg-white lg:p-[15px] lg:shadow-md"
                      @click="expandListItem(faq.id)"
                    >
                      <div class="faq-header flex flex-row items-center">
                        <div
                          class="faq-question w-[90%] p-[10px] sub-h1-semibold lg:mobile-header-2-semibold"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'v2-canopas-gradient-text footer-icon'
                              : 'text-black-60 lg:text-black-87'
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
                            class="angle-up text-black-87"
                            :class="
                              openList && faq.id == currentIndex
                                ? 'footer-icon'
                                : 'text-black-87'
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
                            class="faq-answer w-[90%] px-3 sub-h3-regular lg:mobile-header-3-regular"
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
              class="w-full my-8 lg:hidden mt-8 flex justify-center items-center"
            >
              <div
                class="m-[5px] rounded-[0.6rem] border border-solid border-transparent bg-gradient-to-r from-[#ff9472] via-[#ff909c] to-[#f2709c] p-[1rem] text-center shadow-[inset_2px_1000px_1px_#fff] active:scale-[0.98] border border-black w-[155px] cursor-pointer h-[57px]"
                @click="(showAdditionalFAQs = !showAdditionalFAQs)"
              >
                <span class="sub-h3-semibold v2-canopas-gradient-text"
                  >{{ showAdditionalFAQs ? "View Less" : "View More" }}
                </span>
                <Icon
                  class="ml-2 fab w-4 h-4 footer-icon"
                  :name="
                    showAdditionalFAQs
                      ? 'fa6-solid:angle-up'
                      : 'fa6-solid:angle-down'
                  "
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="hidden lg:block container -mt-8 text-right lg:pt-4 xl:pr-16 2xl:pr-40"
    >
      <button
        v-if="isActivePrev"
        type="button"
        :disabled="!isActivePrev"
        class="h-12 flex ml-auto items-center text-center"
        @click="slide(-1)"
        @click.native="mixpanel.track('tap_ios_app_faq_previous_arrow')"
        aria-label="leftArrow"
      >
        <Icon
          :class="isActivePrev ? 'footer-icon' : ''"
          class="mr-2 h-5 w-5"
          icon="arrow-left"
          id="leftArrow"
        /><span class="v2-canopas-gradient-text sub-h1-semibold">Back</span>
      </button>
      <button
        v-else
        type="button"
        class="h-12 flex ml-auto items-center text-center"
        :disabled="!isActiveNext"
        @click="slide(1)"
        @click.native="mixpanel.track('tap_ios_app_faq_next_arrow')"
        aria-label="rightArrow"
      >
        <span class="v2-canopas-gradient-text sub-h1-semibold"
          >View More Questions</span
        >
        <Icon
          :class="isActiveNext ? 'footer-icon' : ''"
          class="ml-2 h-5 w-5"
          name="fa6-solid:arrow-right"
          id="rightArrow"
        />
      </button>
    </div>
  </section>
</template>

<script type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
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
