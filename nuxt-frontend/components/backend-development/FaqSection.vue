<template>
  <section class="my-16 lg:my-60 3xl:outer-container">
    <div class="container">
      <h2 class="mobile-header-2 lg:desk-header-2 text-black-87 text-center">
        FAQs on Backend development
      </h2>
      <transition-group
        tag="div"
        class="mt-6 mx-2 lg:mt-[4.5rem]"
        :name="'faq-' + faqTransitionName"
      >
        <div class="faq-section">
          <div
            class="mb-4"
            v-for="faq in faqs.slice(sliceFrom, sliceTo)"
            :key="faq"
          >
            <div
              class="faq-container cursor-pointer overflow-hidden rounded-xl py-2 lg:py-8 lg:px-10 lg:shadow-md"
              @click="expandListItem(faq.id)"
            >
              <div class="faq-header flex items-center">
                <div
                  class="faq-question w-[90%] sub-h1-semibold lg:mobile-header-2-semibold"
                  :class="
                    openList && faq.id == currentIndex
                      ? 'v2-canopas-gradient-text footer-icon'
                      : 'text-black-60'
                  "
                >
                  {{ faq.question }}
                </div>
                <div
                  class="faq-icon w-[10%] text-end"
                  @click.native="$mixpanel.track('tap_backend_faq_questions')"
                >
                  <Icon
                    class="angle-up text-black-60"
                    :class="
                      openList && faq.id == currentIndex
                        ? 'footer-icon'
                        : 'text-black-60'
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
                  class="faq-header mt-4 flex animate-fadeIn items-center gradient-border h-auto border-l md:border-l-4 ml-3 lg:ml-1"
                  v-if="openList && faq.id == currentIndex"
                  :key="faq.answer"
                >
                  <div
                    class="faq-answer w-[90%] px-3 sub-h3-regular lg:mobile-header-3-regular text-black-60"
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
          class="mt-4 mx-2"
          :name="'faq-' + faqTransitionName"
          v-if="showAdditionalFAQs"
        >
          <div class="faq-section lg:hidden">
            <div class="mb-4" v-for="faq in faqs.slice(5, 10)" :key="faq">
              <div
                class="faq-container cursor-pointer overflow-hidden rounded-xl bg-white py-2 lg:py-8 lg:px-10 lg:shadow-md"
                @click="expandListItem(faq.id)"
              >
                <div class="faq-header flex items-center">
                  <div
                    class="faq-question w-[90%] sub-h1-semibold lg:mobile-header-2-semibold"
                    :class="
                      openList && faq.id == currentIndex
                        ? 'v2-canopas-gradient-text footer-icon'
                        : 'text-black-60'
                    "
                  >
                    {{ faq.question }}
                  </div>
                  <div
                    class="faq-icon w-[10%] text-end"
                    @click.native="$mixpanel.track('tap_backend_faq_questions')"
                  >
                    <Icon
                      class="angle-up text-black-60"
                      :class="
                        openList && faq.id == currentIndex
                          ? 'footer-icon'
                          : 'text-black-60'
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
                    class="faq-header mt-4 flex animate-fadeIn items-center gradient-border h-auto border-l md:border-l-4 ml-3"
                    v-if="openList && faq.id == currentIndex"
                    :key="faq.answer"
                  >
                    <div
                      class="faq-answer w-[90%] px-3 sub-h3-regular lg:mobile-header-3-regular text-black-60"
                      v-html="faq.answer"
                    ></div>
                  </div>
                </collapse-transition>
                <div class="faq-divider" v-if="faq.id != faq.length"></div>
              </div>
            </div>
          </div> </transition-group
      ></collapse-transition>
      <div
        class="primary-btn gradient-border-btn white-btn mt-8 lg:hidden mx-auto"
        @click="(showAdditionalFAQs = !showAdditionalFAQs)"
      >
        <span class="sub-h3-semibold"
          >{{ showAdditionalFAQs ? "View Less" : "View More" }}
        </span>
        <Icon
          :name="
            showAdditionalFAQs ? 'fa6-solid:angle-up' : 'fa6-solid:angle-down'
          "
          class="fa ml-2 w-4 h-4"
        />
      </div>
    </div>
    <div class="hidden lg:block container text-right">
      <button
        v-if="isActivePrev"
        type="button"
        :disabled="!isActivePrev"
        class="h-12 flex ml-auto items-center text-center"
        @click="slide(-1)"
        @click.native="$mixpanel.track('tap_backend_faq_previous_arrow')"
      >
        <Icon
          :class="isActivePrev ? 'footer-icon' : ''"
          class="mr-2 h-5 w-5"
          name="fa6-solid:arrow-left"
          id="leftArrow"
        /><span class="v2-canopas-gradient-text sub-h1-semibold">Back</span>
      </button>
      <button
        v-else
        type="button"
        class="h-12 flex ml-auto items-center text-center"
        :disabled="!isActiveNext"
        @click="slide(1)"
        @click.native="$mixpanel.track('tap_backend_faq_next_arrow')"
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
<script setup type="module">
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
const { $mixpanel } = useNuxtApp();
const TOTAL_FAQ_IN_SLIDE = 5;
let showAdditionalFAQs = ref(false);
let faqs = [
  {
    id: 1,
    question: "What is backend development and why is it important?",
    answer:
      "Backend development refers to server-side development, where developers focus on databases, scripting, and the architecture of a website or application. It's crucial as it powers the client side (frontend) and ensures data, application logic, and server operations run smoothly.",
  },
  {
    id: 2,
    question:
      "Which technologies and languages do you specialize in for backend development?",
    answer:
      "We specialize in a range of backend technologies such as GoLang, Node.js,PHP, Ruby on Rails, Python, etc.",
  },
  {
    id: 3,
    question: "How do you handle post-deployment maintenance and updates?",
    answer:
      "We offer 3 months of FREE post-deployment support to ensure your backend continues to operate efficiently. This includes routine checks, updates, and immediate assistance should any issues arise.",
  },
  {
    id: 4,
    question: "How do you ensure the scalability of the backend?",
    answer:
      "Our backend solutions are designed with scalability in mind. We adopt scalable architecture patterns, use cloud solutions when appropriate, and ensure the database is optimized for growth.",
  },
  {
    id: 5,
    question: "How long does a typical backend development project take?",
    answer:
      "The timeline for backend development varies based on the project's complexity, features, and specific requirements. After a detailed assessment, we can provide a more accurate timeframe tailored to your project.",
  },
  {
    id: 6,
    question: "How do you ensure the security of my application's backend?",
    answer:
      "Our team implements industry-leading security practices, including encryption, secure coding practices, and regular security audits. We also ensure compliance with industry-specific regulations.",
  },
  {
    id: 7,
    question: "What's the difference between frontend and backend development?",
    answer:
      "Frontend (or client-side) development focuses on what users see and interact with in a browser or application. Backend (or server-side) development manages the application, database, and system infrastructure, ensuring everything runs seamlessly behind the scenes.",
  },
  {
    id: 8,
    question: "How do you handle data migrations?",
    answer:
      "We have a systematic approach to data migration that includes backup, data integrity checks, and rigorous testing post-migration to ensure no data is lost or compromised.",
  },
  {
    id: 9,
    question:
      "Can you integrate third-party services/APIs into my existing system?",
    answer:
      "Absolutely! We have extensive experience in integrating a wide range of third-party services and APIs, ensuring that they work seamlessly with your existing system.",
  },
  {
    id: 10,
    question: "What is your experience with Agile/Scrum methodologies?",
    answer:
      "Our team is well-versed in Agile/Scrum methodologies, ensuring timely delivery, adaptability, and effective communication throughout the project.",
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
  var len = faqs.length / TOTAL_FAQ_IN_SLIDE;
  return (
    (current.value + (diff % Math.ceil(len)) + Math.ceil(len)) % Math.ceil(len)
  );
};
</script>
<style scoped lang="postcss">
.gradient-border {
  border-image: linear-gradient(180deg, #f2709c 50%, #ff835b 50%) 1;
}
.gradient-border-btn > .fa {
  @apply text-[#f2709c];
}
.gradient-border-btn:hover > .fa {
  @apply text-white;
}
</style>
