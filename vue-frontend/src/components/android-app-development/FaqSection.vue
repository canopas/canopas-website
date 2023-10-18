<template>
  <section class="sm:tw-mt-16 md:tw-mt-40 xl:tw-mt-60">
    <div
      class="tw-mb-4 tw-bg-pink-gradient-background md:tw-container md:tw-bg-none"
    >
      <div
        class="tw-mx-[2%] tw-my-0 tw-flex tw-flex-col tw-gap-4 md:tw-my-0 md:-tw-ml-8"
      >
        <div class="tw-text-center">
          <span class="header-2 tw-text-black-core/[0.87]"
            >FAQs on android app development</span
          >
        </div>
        <div class="tw-mt-6 tw-mx-auto">
          <img
            :src="faqImage"
            loading="lazy"
            alt="frequently-asked-questions-image"
            class="tw-h-full tw-w-[20.25rem] tw-object-cover md:-tw-mb-10 md:tw-w-[25.75rem] md:tw-w-[412px]"
          />
        </div>
        <div
          class="tw-container tw--mt-8 tw-flex-[67%] md:tw-pb-16 md:tw-w-[65rem]"
        >
          <div class="md:tw-pl-[30px]">
            <transition-group tag="div" :name="'faq-' + faqTransitionName">
              <div class="faq-section tw-mt-12 tw-h-auto">
                <div
                  class="tw-mt-3 md:tw-mt-6"
                  v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                  :key="faq"
                >
                  <div
                    class="faq-container tw-cursor-pointer tw-overflow-hidden tw-rounded-[5px] tw-bg-white tw-p-[15px] tw-shadow-md sm:tw-p-[25px]"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="faq-header tw-flex tw-flex-row tw-items-center">
                      <div
                        class="faq-question tw-w-[90%] sub-h1-semibold"
                        :class="
                          openList && faq.id == currentIndex
                            ? 'v2-canopas-gradient-text footer-icon'
                            : 'tw-text-black-core/[0.87]'
                        "
                      >
                        {{ faq.question }}
                      </div>
                      <div
                        class="faq-icon tw-w-[10%] tw-text-end"
                        @click.native="
                          mixpanel.track('tap_android_app_faq_questions')
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
                              ? 'xmark'
                              : 'plus'
                          "
                        />
                      </div>
                    </div>
                    <collapse-transition>
                      <div
                        class="faq-header tw-flex tw-flex-row tw-items-center"
                        v-if="openList && faq.id == currentIndex"
                        :key="faq.answer"
                      >
                        <div
                          class="faq-answer tw-mt-4 tw-w-[90%] tw-animate-fadeIn sub-h3-regular tw-text-black-core/[0.87]"
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
                <div class="faq-section tw-h-auto">
                  <div
                    class="tw-mt-3"
                    v-for="faq in faqs.slice(5, 10)"
                    :key="faq"
                  >
                    <div
                      class="faq-container tw-cursor-pointer tw-overflow-hidden tw-rounded-[5px] tw-bg-white tw-p-[15px] tw-shadow-md sm:tw-p-[25px]"
                      @click="expandListItem(faq.id)"
                    >
                      <div
                        class="faq-header tw-flex tw-flex-row tw-items-center"
                      >
                        <div
                          class="faq-question tw-w-[90%] sub-h1-semibold"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'v2-canopas-gradient-text footer-icon'
                              : 'tw-text-black-core/[0.87]'
                          "
                        >
                          {{ faq.question }}
                        </div>
                        <div
                          class="faq-icon tw-w-[10%] tw-text-end"
                          @click.native="
                            mixpanel.track('tap_android_app_faq_questions')
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
                                ? 'xmark'
                                : 'plus'
                            "
                          />
                        </div>
                      </div>
                      <collapse-transition>
                        <div
                          class="faq-header tw-flex tw-flex-row tw-items-center"
                          v-if="openList && faq.id == currentIndex"
                          :key="faq.answer"
                        >
                          <div
                            class="faq-answer tw-mt-4 tw-w-[90%] tw-animate-fadeIn sub-h3-regular tw-text-black-core/[0.87]"
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
            <div class="tw-mx-auto tw-my-8 md:tw-hidden">
              <div
                class="tw-border tw-border-white primary-btn"
                @click="showAdditionalFAQs = !showAdditionalFAQs"
              >
                <span class="tw-text-white sub-h3-semibold"
                  >{{ showAdditionalFAQs ? "View Less" : "View More" }}
                </span>
                <font-awesome-icon
                  class="tw-ml-2 tw-mt-1 fab tw-w-4 tw-h-4 tw-text-white"
                  :icon="showAdditionalFAQs ? faAngleUp : faAngleDown"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="tw-hidden md:tw-block tw-container tw--mt-8 tw-pr-4 tw-text-right md:tw-pr-6 md:tw-pt-4 xl:tw-pr-20 2xl:tw-pr-44"
    >
      <button
        v-if="isActivePrev"
        type="button"
        :disabled="!isActivePrev"
        class="tw-h-12 tw-flex tw-ml-auto tw-items-center tw-text-center"
        @click="slide(-1)"
        @click.native="mixpanel.track('tap_android_app_faq_previous_arrow')"
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
        @click.native="mixpanel.track('tap_android_app_faq_next_arrow')"
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
import faqImage from "@/assets/images/andriod-app-development/faq.webp";
import { faAngleUp, faAngleDown } from "@fortawesome/free-solid-svg-icons";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
      faAngleUp,
      faAngleDown,
      showAdditionalFAQs: false,
      faqImage: faqImage,
      faqs: [
        {
          id: 1,
          question:
            " What is the typical timeline for developing an Android app?",
          answer:
            " The timeline for app development varies depending on the complexity of the app, its features, and the specific requirements of the project. On average, a simple app can take around 2-3 months, while a \
more complex app can take 6-9 months or more.",
        },
        {
          id: 2,
          question: "How much does it cost to develop an Android app?",
          answer:
            " The cost of app development can vary greatly depending on the features, complexity, and resources required. We would be happy to provide a quote after discussing your specific app needs",
        },
        {
          id: 3,
          question: "Can you help if I only have an app idea?",
          answer:
            "Absolutely! We provide consultation services where we can help shape your idea, define clear goals, identify potential challenges and provide solutions, and offer a roadmap for development.",
        },
        {
          id: 4,
          question: "Do you provide app design (UI/UX) services? ",
          answer:
            "Yes,we have a dedicated team of UI/UX designers who will work closely with you to create an intuitive and visually appealing app that enhances user experience.",
        },
        {
          id: 5,
          question: "How do you ensure the quality of the app?",
          answer:
            " We follow a rigorous testing process which includes unit testing, integration testing, and functional testing to ensure that the app is bug-free and offers a smooth user experience.\
           We also perform usability testing to ensure the app is user-friendly.",
        },
        {
          id: 6,
          question: "What technologies do you use for Android app development?",
          answer:
            "We use a variety of technologies depending on the requirements of the project. This includes Java, Kotlin, Firebase, Android Studio, SQLite, Retrofit, and more.",
        },
        {
          id: 7,
          question: "Will my app be compatible with all versions of Android?",
          answer:
            " We aim to make apps compatible with the most recent and widely used versions of Android to reach the majority of Android users. However, if you have specific requirements, we can discuss those during the \
consultation phase.",
        },
        {
          id: 8,
          question: "Do you offer post-development support and maintenance? ",
          answer:
            " Yes,we provide maintenance and support services after the app is launched. This includes fixing any bugs,updating the app for new Android versions, and adding new features as required.",
        },
        {
          id: 9,
          question:
            "Can you integrate third-party services or APIs into my app? ",
          answer:
            " Yes, we have extensive experience in integrating third-party services and APIs into Android apps based on project requirements.",
        },
        {
          id: 10,
          question: "How will my app be tested? ",
          answer:
            "We perform multiple rounds of testing including unit testing, integration testing, functional testing, and usability testing. We also conduct \
performance testing to ensure your app runs smoothly under varying loads.",
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
        this.mixpanel.track("tap_android_app_faq_collapse");
      } else {
        this.openList = true;
        this.mixpanel.track("tap_android_app_faq_expand");
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
