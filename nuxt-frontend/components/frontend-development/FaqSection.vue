<template>
  <div class="bg-deep-charcoal text-white">
    <div class="container pt-14 pb-20 lg:pt-[7.5rem] lg:pb-14">
      <p
        class="text-center mobile-header-2 lg:desk-header-2 mb-6 lg:mb-[4.5rem]"
      >
        FAQs on Frontend development
      </p>
      <transition-group tag="div" :name="'faq-' + faqTransitionName">
        <div class="faq-section">
          <div
            class="mt-4 lg:mt-6"
            v-for="faq in faqs.slice(sliceFrom, sliceTo)"
            :key="faq"
          >
            <div
              class="faq-container cursor-pointer overflow-hidden rounded-xl bg-pink-80 py-6 px-4 lg:py-8 lg:px-10"
              @click="expandListItem(faq.id)"
            >
              <div class="faq-header flex gap-4 lg:gap-6">
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
                <div
                  class="sub-h1-semibold lg:mobile-header-2-semibold w-[90%]"
                >
                  {{ faq.question }}
                </div>
              </div>
              <collapse-transition>
                <div
                  class="faq-header flex gap-4 lg:gap-6"
                  v-if="openList && faq.id == currentIndex"
                  :key="faq.answer"
                >
                  <div class="w-[6%]"></div>
                  <div
                    class="faq-answer mt-4 w-[90%] animate-fadeIn sub-h3-regular lg:mobile-header-3-regular text-white items-end"
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
                class="faq-container cursor-pointer overflow-hidden rounded-xl bg-pink-80 py-6 px-4 lg:py-8 lg:px-10"
                @click="expandListItem(faq.id)"
              >
                <div class="faq-header flex gap-4 lg:gap-6">
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
                  <div
                    class="sub-h1-semibold lg:mobile-header-2-semibold w-[90%]"
                  >
                    {{ faq.question }}
                  </div>
                </div>
                <collapse-transition>
                  <div
                    class="faq-header flex items-center gap-4 lg:gap-6"
                    v-if="openList && faq.id == currentIndex"
                    :key="faq.answer"
                  >
                    <div class="w-[6%]"></div>
                    <div
                      class="faq-answer mt-4 w-[90%] animate-fadeIn sub-h3-regular lg:mobile-header-3-regular text-white"
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
    <div
      class="hidden lg:block container text-right pb-[3.75rem] xll:pb-3 3xl:pb-9"
    >
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
        /><span class="sub-h1-semibold">Back</span>
      </button>
      <button
        v-else
        type="button"
        class="h-12 flex ml-auto items-center text-center"
        :disabled="!isActiveNext"
        @click="slide(1)"
        @click.native="$mixpanel.track('tap_flutter_app_faq_next_arrow')"
      >
        <span class="sub-h1-semibold">View More Questions</span>
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
    question: "What is frontend development?",
    answer:
      "Frontend development focuses on building the visual and interactive aspects of a website or web application. It's everything the user sees and interacts with directly, including design, layout, buttons, and forms.",
  },
  {
    id: 2,
    question: "How does frontend differ from backend development?",
    answer:
      "While frontend development deals with the user interface and user experience, backend development handles the server, database, andapplication logic, ensuring data processing and storag",
  },
  {
    id: 3,
    question: "What technologies do you use for frontend development?",
    answer:
      "We utilize a range of cutting-edge technologies including HTML, CSS, JavaScript, and frameworks like React, Angular, Vue.js, and more, depending on the project's requirements.",
  },
  {
    id: 4,
    question: "Why is responsive design important?",
    answer:
      "Responsive design ensures your website or application looks and functions optimally across all devices, from desktops to smartphones, enhancing user experience and reach.",
  },
  {
    id: 5,
    question: "How can performance optimization benefit my website?",
    answer:
      "Performance optimization ensures faster loading times, better user retention, improved SEO rankings, and a boost in overall user engagement and conversion rates.",
  },
  {
    id: 6,
    question: "What is involved in a frontend audit?",
    answer:
      "A frontend audit evaluates the quality, performance, and efficiency of your website or application's frontend code. It identifies areas for improvement, potential optimizations, and ensures best practices are followed.",
  },
  {
    id: 7,
    question: "How do you ensure the UI/UX design aligns with our brand?",
    answer:
      "We collaborate closely with clients to understand their brand ethos, objectives, and audience. This insight, combined with our design expertise, ensures a UI/UX that resonates with your brand's voice and vision.",
  },
  {
    id: 8,
    question: "Can you help migrate my website to a modern frontend framework?",
    answer:
      "Absollutely! We can assist in migrating your existing website to a modern frontend framework, enhancing performance, scalability, and maintainability.",
  },
  {
    id: 9,
    question:
      "Do you provide ongoing support and maintenance for frontend projects?",
    answer:
      "Yes, we offer post-development support and maintenance packages to ensure your site remains updated, optimized, and free from issues as technologies evolve.",
  },
  {
    id: 10,
    question: "How do frontend optimizations impact SEO?",
    answer:
      "Optimizing the frontend can lead to faster page load times, improved user experience, and reduced bounce rates—all of which are key factors search engines consider when ranking websites.",
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
