<template>
  <section>
    <div class="container my-16 lg:my-60">
      <div
        class="row flex flex-col md:flex-row gap-4 my-0 mx-[2%] md:mx-[6%] md:-ml-8"
      >
        <div class="text-center mobile-header-2 text-black-87 md:hidden mb-8">
          Frequently Asked Questions
        </div>
        <div
          class="basis-[33%] image py-0 px-[20%] md:py-0 md:px-0 md:translate-y-60 lg:translate-y-80 xl:translate-y-40 animate__animated"
          ref="faqImage"
        >
          <aspect-ratio height="100%">
            <img
              :src="faqImage"
              loading="lazy"
              alt="frequently-asked-questions-image"
              class="h-full w-full object-contain"
            />
          </aspect-ratio>
        </div>
        <div class="basis-[67%]">
          <div class="md:pl-[3.75rem]">
            <div class="desk-header-2 text-black-87 hidden md:block">
              Frequently Asked Questions
            </div>
            <transition-group tag="div" :name="'faq-' + faqTransitionName">
              <div class="mt-12 min-h-[28rem] h-auto sm:min-h-[33rem]">
                <div
                  class="mt-6"
                  v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                  :key="faq"
                >
                  <div
                    class="faq-container border border-solid overflow-hidden bg-white-smoke shadow-[0px_1px_6px_1px_rgba(0,0,0,0.25)] rounded-[15px] p-[0.938rem] cursor-pointer sm:p-[1.563rem]"
                    @click="expandListItem(faq.id)"
                  >
                    <div class="flex items-center">
                      <div class="w-[10%] text-pink-300">
                        <Icon
                          class="plus-icon"
                          :name="
                            openList && faq.id == currentIndex
                              ? 'fa6-solid:minus'
                              : 'fa6-solid:plus'
                          "
                        />
                      </div>
                      <div
                        class="sub-h1-semibold lg:mobile-header-2-semibold text-black-87 w-[90%]"
                      >
                        {{ faq.question }}
                      </div>
                    </div>
                    <collapse-transition>
                      <div
                        class="flex items-center"
                        v-if="openList && faq.id == currentIndex"
                        :key="faq.answer"
                      >
                        <div class="w-[10%]"></div>
                        <div
                          class="sub-h3-regular lg:sub-h1-regular text-black-60 w-[90%] mt-4"
                          v-html="faq.answer"
                        ></div>
                      </div>
                    </collapse-transition>
                  </div>
                </div>
              </div>
            </transition-group>
            <collapse-transition>
              <transition-group
                v-if="showAdditionalFAQs"
                tag="div"
                :name="'faq-' + faqTransitionName"
              >
                <div class="min-h-[28rem] h-auto sm:min-h-[33rem]">
                  <div class="mt-6" v-for="faq in faqs.slice(5, 10)" :key="faq">
                    <div
                      class="faq-container border border-solid overflow-hidden bg-white-smoke shadow-[0px_1px_6px_1px_rgba(0,0,0,0.25)] rounded-[15px] p-[0.938rem] cursor-pointer sm:p-[1.563rem]"
                      @click="expandListItem(faq.id)"
                    >
                      <div class="flex items-center">
                        <div class="w-[10%] text-pink-300">
                          <Icon
                            class="plus-icon"
                            :name="
                              openList && faq.id == currentIndex
                                ? 'fa6-solid:minus'
                                : 'fa6-solid:plus'
                            "
                          />
                        </div>
                        <div
                          class="sub-h1-semibold lg:mobile-header-2-semibold text-black-87 w-[90%]"
                        >
                          {{ faq.question }}
                        </div>
                      </div>
                      <collapse-transition>
                        <div
                          class="flex items-center"
                          v-if="openList && faq.id == currentIndex"
                          :key="faq.answer"
                        >
                          <div class="w-[10%]"></div>
                          <div
                            class="sub-h3-regular lg:sub-h1-regular text-black-60 w-[90%] mt-4"
                            v-html="faq.answer"
                          ></div>
                        </div>
                      </collapse-transition>
                    </div>
                  </div>
                </div>
              </transition-group>
            </collapse-transition>
            <div
              class="w-full my-8 lg:hidden mt-8 flex justify-center items-center"
            >
              <div
                class="m-[5px] rounded-[0.6rem] border border-solid border-transparent bg-gradient-to-r from-pink-300 to-orange-300 p-4 text-center shadow-[inset_2px_1000px_1px_white] active:scale-[0.98] border border-black w-[9.688rem] cursor-pointer h-[3.563rem]"
                @click="showAdditionalFAQs = !showAdditionalFAQs"
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
      <div
        class="hidden lg:block container text-right mt-10 xl:pr-16 2xl:pr-20"
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
    </div>
  </section>
</template>

<script type="module">
import AspectRatio from "@/components/utils/AspectRatio.vue";
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import faqImage from "@/assets/images/faq/FAQ.svg";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
      showAdditionalFAQs: false,
      faqImage: faqImage,
      faqs: [
        {
          id: 1,
          question: "What are the top skills you look for in a candidate?",
          answer: `It mostly depends on the job you are applying for. Usually, apart from the technical skills, we look for someone who, </br>
                   - has a deep desire to learn and excel </br>
                   - has a willingness to work hard</br>
                   - is looking for a challenging environment</br>
                   - is not afraid to leave the comfort zone</br>`,
        },
        {
          id: 2,
          question: "Do you allow flexible work schedules?",
          answer: "Yeah, you can come anytime between 7 AM to 9 PM.",
        },
        {
          id: 3,
          question: "Do you have an employee referral program?",
          answer:
            "Yeah, we do have an employee referral program. Reach out to your friends or family working at Canopas so they can get a referral bonus if you’re hired.",
        },
        {
          id: 4,
          question: "How can I know my application status?",
          answer:
            "We do our best to give you a reply to your application as soon as possible.",
        },
        {
          id: 5,
          question: "What should I wear?",
          answer: "Whatever makes you comfortable. We don’t have a dress code.",
        },
        {
          id: 6,
          question: "What does the hiring process look like?",
          answer:
            "You will be greeted by HR when you visit us for an interview. There will be two rounds of interviews. Technical or Non-technical followed by HR interview. ",
        },
        {
          id: 7,
          question: "How much time will the interview take?",
          answer:
            "We won't take too much time from your busy schedule. Your Interview might take an hour or two at max.",
        },
        {
          id: 8,
          question: "How should I apply?",
          answer:
            "Browse through all openings available on this website and apply for positions you’re looking for. If any job doesn’t match with your skills, you can check back later.",
        },
        {
          id: 9,
          question: "How much experience do I need?",
          answer:
            "You will find more details about it in the job descriptions page.",
        },
        {
          id: 10,
          question: "I’m a college graduate. Do you hire freshers?",
          answer:
            "Yes, We are always looking for people who want a challenging career with personal growth. You can apply from the available openings that you find most interesting.",
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
    AspectRatio,
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  inject: ["mixpanel"],
  methods: {
    expandListItem(index) {
      if (this.previousIndex == index && this.openList) {
        this.openList = false;
        this.$mixpanel.track("tap_faq_collapse");
      } else {
        this.openList = true;
        this.$mixpanel.track("tap_faq_expand");
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
      var len = this.faqs.length / 5;
      return (
        (this.current + (diff % Math.ceil(len)) + Math.ceil(len)) %
        Math.ceil(len)
      );
    },
    handleScroll() {
      let data = {
        name: this.$refs.faqImage,
        animation: "animate__zoomIn",
      };
      if (data) {
        this.$emit("add-animation", data);
      }
    },
  },
};
</script>
