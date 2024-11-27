<template>
  <section class="sm:mt-16 md:mt-40 xl:mt-60">
    <div class="mb-4 bg-pink-gradient-background lg:container lg:bg-none">
      <div class="mx-[2%] my-0 flex flex-col gap-4 lg:my-0 lg:-ml-8">
        <div class="text-center">
          <span class="mobile-header-2 lg:desk-header-2 text-black-87"
            >FAQs on android app development</span
          >
        </div>
        <div class="mt-6 mx-auto">
          <img
            :src="faqImage"
            loading="lazy"
            alt="frequently-asked-questions-image"
            class="h-full w-[20.25rem] object-cover lg:-mb-10 lg:w-[25.75rem] lg:w-[412px]"
          />
        </div>
        <div class="container -mt-8 flex-[67%] lg:pb-16 lg:w-[65rem]">
          <div class="lg:pl-[30px]">
            <transition-group tag="div" :name="'faq-' + faqTransitionName">
              <div class="faq-section mt-12 h-auto">
                <div
                  class="mt-3 lg:mt-6"
                  v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                  :key="faq"
                >
                  <div
                    class="faq-container cursor-pointer overflow-hidden rounded-[5px] bg-white p-[15px] shadow-md sm:p-[25px]"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="faq-header flex flex-row items-center">
                      <div
                        class="faq-question w-[90%] sub-h1-semibold lg:mobile-header-2-semibold"
                        :class="
                          openList && faq.id == currentIndex
                            ? 'v2-canopas-gradient-text footer-icon'
                            : 'text-black-87'
                        "
                      >
                        {{ faq.question }}
                      </div>
                      <div
                        class="faq-icon w-[10%] text-end"
                        @click.native="
                          $mixpanel.track('tap_android_app_faq_questions')
                        "
                      >
                        <Icon
                          class="xmark text-black-87"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'footer-icon'
                              : 'text-black-87'
                          "
                          :name="
                            openList && faq.id == currentIndex
                              ? 'fa6-solid:xmark'
                              : 'fa6-solid:plus'
                          "
                        />
                      </div>
                    </div>
                    <collapse-transition>
                      <div
                        class="faq-header flex flex-row items-center"
                        v-if="openList && faq.id == currentIndex"
                        :key="faq.answer"
                      >
                        <div
                          class="faq-answer mt-4 w-[90%] animate-fadeIn sub-h3-regular lg:mobile-header-3-regular text-black-87"
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
                <div class="faq-section h-auto">
                  <div class="mt-3" v-for="faq in faqs.slice(5, 10)" :key="faq">
                    <div
                      class="faq-container cursor-pointer overflow-hidden rounded-[5px] bg-white p-[15px] shadow-md sm:p-[25px]"
                      @click="expandListItem(faq.id)"
                    >
                      <div class="faq-header flex flex-row items-center">
                        <div
                          class="faq-question w-[90%] sub-h1-semibold lg:mobile-header-2-semibold"
                          :class="
                            openList && faq.id == currentIndex
                              ? 'v2-canopas-gradient-text footer-icon'
                              : 'text-black-87'
                          "
                        >
                          {{ faq.question }}
                        </div>
                        <div
                          class="faq-icon w-[10%] text-end"
                          @click.native="
                            $mixpanel.track('tap_android_app_faq_questions')
                          "
                        >
                          <Icon
                            class="xmark text-black-87"
                            :class="
                              openList && faq.id == currentIndex
                                ? 'footer-icon'
                                : 'text-black-87'
                            "
                            :name="
                              openList && faq.id == currentIndex
                                ? 'fa6-solid:xmark'
                                : 'fa6-solid:plus'
                            "
                          />
                        </div>
                      </div>
                      <collapse-transition>
                        <div
                          class="faq-header flex flex-row items-center"
                          v-if="openList && faq.id == currentIndex"
                          :key="faq.answer"
                        >
                          <div
                            class="faq-answer mt-4 w-[90%] animate-fadeIn sub-h3-regular lg:mobile-header-3-regular text-black-87"
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
            <div class="mx-auto my-8 lg:hidden">
              <div
                class="border border-white primary-btn"
                @click="showAdditionalFAQs = !showAdditionalFAQs"
              >
                <span class="text-white sub-h3-semibold"
                  >{{ showAdditionalFAQs ? "View Less" : "View More" }}
                </span>
                <Icon
                  class="ml-2 mt-1 fab w-4 h-4 text-white"
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
      class="hidden lg:block container -mt-8 pr-4 text-right lg:pr-6 lg:pt-4 xl:pr-20 2xl:pr-44"
    >
      <button
        v-if="isActivePrev"
        type="button"
        :disabled="!isActivePrev"
        class="h-12 flex ml-auto items-center text-center"
        @click="slide(-1)"
        @click.native="$mixpanel.track('tap_android_app_faq_previous_arrow')"
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
        @click.native="$mixpanel.track('tap_android_app_faq_next_arrow')"
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
import faqImage from "@/assets/images/andriod-app-development/faq.webp";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
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
        this.$mixpanel.track("tap_android_app_faq_collapse");
      } else {
        this.openList = true;
        this.$mixpanel.track("tap_android_app_faq_expand");
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
      let len = this.faqs.length / TOTAL_FAQ_IN_SLIDE;
      return (
        (this.current + (diff % Math.ceil(len)) + Math.ceil(len)) %
        Math.ceil(len)
      );
    },
  },
};
</script>
