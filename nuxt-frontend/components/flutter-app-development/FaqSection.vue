<template>
  <div class="bg-[#121212] text-white 3xl:outer-container">
    <div class="container pt-14 pb-20 lg:pt-[7.5rem] lg:pb-10">
      <p class="text-center header-2 mb-6 lg:mb-[4.5rem]">
        FAQs on Flutter app development
      </p>
      <transition-group tag="div" :name="'faq-' + faqTransitionName">
        <div class="faq-section">
          <div
            class="mt-4 lg:mt-6"
            v-for="faq in faqs.slice(sliceFrom, sliceTo)"
            :key="faq"
          >
            <div
              class="faq-container cursor-pointer overflow-hidden rounded-xl bg-[#F2709C]/[0.08] py-6 px-4 lg:py-8 lg:px-10"
              @click="expandListItem(faq.id)"
            >
              <div class="faq-header flex flex-row gap-4 lg:gap-6">
                <div
                  class="my-auto w-[6%] text-end"
                  @click.native="
                    $mixpanel.track('tap_flutter_app_faq_questions')
                  "
                >
                  <Icon
                    class="minus"
                    :name="
                      openList && faq.id == currentIndex
                        ? 'fa6-solid:minus'
                        : 'fa6-solid:plus'
                    "
                  />
                </div>
                <div class="sub-h1-semibold w-[90%]">
                  {{ faq.question }}
                </div>
              </div>
              <collapse-transition>
                <div
                  class="faq-header flex flex-row gap-4 lg:gap-6"
                  v-if="openList && faq.id == currentIndex"
                  :key="faq.answer"
                >
                  <div class="w-[6%]"></div>
                  <div
                    class="faq-answer mt-4 w-[90%] animate-fadeIn sub-h3-regular text-white items-end"
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
          <div class="faq-section">
            <div
              class="mt-4 lg:mt-6"
              v-for="faq in faqs.slice(5, 10)"
              :key="faq"
            >
              <div
                class="faq-container cursor-pointer overflow-hidden rounded-xl bg-[#F2709C]/[0.08] py-6 px-4 lg:py-8 lg:px-10"
                @click="expandListItem(faq.id)"
              >
                <div class="faq-header flex flex-row gap-4 lg:gap-6">
                  <div
                    class="my-auto w-[6%] text-end"
                    @click.native="
                      $mixpanel.track('tap_flutter_app_faq_questions')
                    "
                  >
                    <Icon
                      class="minus"
                      :name="
                        openList && faq.id == currentIndex
                          ? 'fa6-solid:minus'
                          : 'fa6-solid:plus'
                      "
                    />
                  </div>
                  <div class="sub-h1-semibold w-[90%]">
                    {{ faq.question }}
                  </div>
                </div>
                <collapse-transition>
                  <div
                    class="faq-header flex flex-row items-center gap-4 lg:gap-6"
                    v-if="openList && faq.id == currentIndex"
                    :key="faq.answer"
                  >
                    <div class="w-[6%]"></div>
                    <div
                      class="faq-answer mt-4 w-[90%] animate-fadeIn sub-h3-regular text-white"
                      v-html="faq.answer"
                    ></div>
                  </div>
                </collapse-transition>
                <div class="faq-divider" v-if="faq.id != faq.length"></div>
              </div>
            </div>
          </div>
        </transition-group>
      </collapse-transition>
      <div
        class="border primary-btn mt-8 lg:hidden"
        @click="showAdditionalFAQs = !showAdditionalFAQs"
      >
        <span class="sub-h3-semibold"
          >{{ showAdditionalFAQs ? "View Less" : "View More" }}
        </span>
        <Icon
          class="ml-2"
          :name="
            showAdditionalFAQs ? 'fa6-solid:angle-up' : 'fa6-solid:angle-down'
          "
        />
      </div>
    </div>
    <div class="hidden lg:block container text-right pb-[8.75rem]">
      <button
        v-if="isActivePrev"
        type="button"
        :disabled="!isActivePrev"
        class="h-12 flex ml-auto items-center text-center"
        @click="slide(-1)"
        @click.native="$mixpanel.track('tap_flutter_app_faq_previous_arrow')"
      >
        <Icon
          class="mr-2 h-5 w-5"
          name="fa6-solid:arrow-left"
          id="leftArrow"
        /><span class="sub-h3-semibold xl:sub-h4-semibold">Back</span>
      </button>
      <button
        v-else
        type="button"
        class="h-12 flex ml-auto items-center text-center"
        :disabled="!isActiveNext"
        @click="slide(1)"
        @click.native="$mixpanel.track('tap_flutter_app_faq_next_arrow')"
      >
        <span class="sub-h3-semibold xl:sub-h4-semibold"
          >View More Questions</span
        >
        <Icon
          class="ml-2 h-5 w-5"
          name="fa6-solid:arrow-right"
          id="rightArrow"
        />
      </button>
    </div>
  </div>
</template>
<script setup type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
const TOTAL_FAQ_IN_SLIDE = 5;
const { $mixpanel } = useNuxtApp();
let showAdditionalFAQs = ref(false);
let faqs = [
  {
    id: 1,
    question: "What is Flutter?",
    answer:
      "Flutter is an open-source UI software development kit created by Google. It allows developers to build natively compiled applications for mobile, web, and desktop from a single codebase.",
  },
  {
    id: 2,
    question: "What language does Flutter use?",
    answer:
      "Flutter uses the Dart programming language, also developed by Google. Dart syntax is easy to understand for JavaScript or Java developers.",
  },
  {
    id: 3,
    question: "What are the advantages of using Flutter?",
    answer:
      "Flutter offers several benefits, including faster code writing and testing, beautiful UI, high performance, and cost-effectiveness due to its single codebase for both iOS and Android platforms.",
  },
  {
    id: 4,
    question: "Is Flutter suitable for any kind of application?",
    answer:
      "Yes, Flutter is versatile and can be usedfor different types of apps, ranging from simple projects to complex applications.",
  },
  {
    id: 5,
    question: "Is Flutter only for mobile applications?",
    answer:
      "While Flutter is primarily known for developing mobile applications, it can also be used to build web and desktop applications.",
  },
  {
    id: 6,
    question: "How long does it take to develop a Flutter app?",
    answer:
      "The development time can vary basedon the complexity of the app, but Flutter alows for faster development due to its hot-reload feature and single codebase.",
  },
  {
    id: 7,
    question: "Is Flutter good for MVP development?",
    answer:
      "Yes, because of its fast development process,Flutter is an excelent choice for building Minimum Viable Products (MVPs).",
  },
  {
    id: 8,
    question: "How does Flutter compare to React Native?",
    answer:
      "Both are popular choices for cross-platform development. Flutter offers native performance, rich UI, and comes with a larger set of widgets out of the box, while React Native has a larger community and more third-party plugins.",
  },
  {
    id: 9,
    question: "Can you integrate APIs in a Flutter application?",
    answer:
      "Yes, Flutter alows you to easily connect your application with different APIs.",
  },
  {
    id: 10,
    question: "Can I use Firebase with Flutter?",
    answer:
      "Yes, Firebase provides a suite of cloud-based services that are commonly used in Flutter apps, including databases, analytics, crash reporting, and authentication.",
  },
];
let previousIndex = ref(null);
let currentIndex = ref(null);
let openList = ref(true);
let current = ref(null);
let faqTransitionName = ref(null);
let sliceFrom = ref(null);
let sliceTo = TOTAL_FAQ_IN_SLIDE;
var isActivePrev = ref(false);
var isActiveNext = ref(true);

const expandListItem = (index) => {
  if (previousIndex.value === index && openList.value) {
    openList.value = false;
    $mixpanel.track("tap_flutter_app_faq_collapse");
  } else {
    openList.value = true;
    $mixpanel.track("tap_flutter_app_faq_expand");
  }
  currentIndex.value = index;
  previousIndex.value = currentIndex.value;
};
let slide = (dir) => {
  faqTransitionName.value = dir === 1 ? "next" : "prev";
  sliceFrom.value =
    dir === 1
      ? sliceFrom.value + TOTAL_FAQ_IN_SLIDE
      : sliceFrom.value - TOTAL_FAQ_IN_SLIDE;
  sliceTo =
    dir === 1 ? sliceTo + TOTAL_FAQ_IN_SLIDE : sliceTo - TOTAL_FAQ_IN_SLIDE;
  isActivePrev.value = sliceFrom.value >= TOTAL_FAQ_IN_SLIDE;
  isActiveNext.value = sliceTo < faqs.length;
  current.value = getRoundedIndex(dir);
};
const getRoundedIndex = (diff) => {
  let len = faqs.length / TOTAL_FAQ_IN_SLIDE;
  return (
    (current.value + (diff % Math.ceil(len)) + Math.ceil(len)) % Math.ceil(len)
  );
};
</script>
