<template>
  <div class="tw-container tw-mb-[10rem]">
    <div
      class="row tw-flex tw-flex-col md:tw-flex-row tw-gap-4 tw-my-0 tw-mx-[2%] md:tw-my-0 md:tw-mx-[6%] md:-tw-ml-[32px]"
    >
      <div
        class="tw-text-center header-text canopas-gradient-text title-mobile tw-block md:tw-hidden tw-mb-12"
      >
        <span
          class="underline-text tw-underline-offset-[0.5rem] sm:tw-underline-offset-[1rem]"
          >Freq</span
        >uently Asked Questions
      </div>
      <div
        class="tw-flex-[33%] image tw-py-0 tw-px-[20%] md:tw-py-0 md:tw-px-0 md:tw-translate-y-60 lg:tw-translate-y-80 xl:tw-translate-y-40 animate__animated"
        ref="faqImage"
      >
        <aspect-ratio height="100%">
          <img
            :src="faqImage"
            loading="lazy"
            alt="frequently-asked-questions-image"
            class="tw-h-full tw-w-full tw-object-contain"
          />
        </aspect-ratio>
      </div>
      <div class="tw-flex-[67%]">
        <div class="faq-content md:tw-pl-[60px]">
          <div
            class="header-text canopas-gradient-text title-desktop tw-hidden md:tw-block"
          >
            <span
              class="underline-text tw-underline-offset-[0.5rem] sm:tw-underline-offset-[1rem]"
              >Freq</span
            >uently Asked Questions
          </div>
          <transition-group tag="div" :name="'faq-' + faqTransitionName">
            <div
              class="tw-mt-12 faq-section tw-min-h-[28rem] tw-h-auto sm:tw-min-h-[33rem]"
            >
              <div
                class="tw-mt-6"
                v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                :key="faq"
              >
                <div
                  class="faq-container tw-border-[1px] tw-border-solid tw-border-[#e2e2e2] tw-overflow-hidden tw-bg-[#3d3d3d08] tw-shadow-[0px_1px_6px_1px_rgba(0,0,0,0.25)] tw-rounded-[15px] tw-p-[15px] tw-cursor-pointer sm:tw-p-[25px] normal-text"
                  @click="expandListItem(faq.id)"
                >
                  <div class="faq-header tw-flex tw-flex-row tw-items-center">
                    <div class="faq-icon tw-w-[10%] tw-text-pink-300">
                      <font-awesome-icon
                        class="plus-icon"
                        :icon="
                          openList && faq.id == currentIndex ? 'minus' : 'plus'
                        "
                      />
                    </div>
                    <div
                      class="faq-question tw-text-black-900 tw-font-bold tw-w-[90%]"
                    >
                      {{ faq.question }}
                    </div>
                  </div>
                  <collapse-transition>
                    <div
                      class="faq-header tw-flex tw-flex-row tw-items-center"
                      v-if="openList && faq.id == currentIndex"
                      :key="faq.answer"
                    >
                      <div class="faq-icon tw-w-[10%]"></div>
                      <div
                        class="normal-2-text faq-answer tw-w-[90%] tw-mt-[1rem]"
                        v-html="faq.answer"
                      ></div>
                    </div>
                  </collapse-transition>
                  <div class="faq-divider" v-if="faq.id != faq.length"></div>
                </div>
              </div>
            </div>
          </transition-group>
          <div class="tw-mt-6 tw-text-center">
            <button
              type="button"
              :disabled="!isActivePrev"
              :class="!isActivePrev ? 'tw-opacity-20' : ''"
              class="clients-indicators tw-bg-none tw-border-none tw-my-0 tw-mx-[8px] tw-cursor-pointer"
              @click="slide(-1)"
              @click.native="mixpanel.track('tap_faq_previous_arrow')"
              aria-label="leftArrow"
            >
              <font-awesome-icon
                class="arrow tw-border-[1px] tw-border-solid tw-border-[#3d3d3d26] tw-rounded-[15px] tw-h-[25px] tw-w-[25px] tw-p-[10px] tw-text-[#fff] tw-bg-black-900"
                icon="arrow-left"
                id="leftArrow"
              />
            </button>
            <button
              type="button"
              :disabled="!isActiveNext"
              :class="!isActiveNext ? 'tw-opacity-20' : ''"
              class="clients-indicators tw-bg-none tw-border-none tw-my-0 tw-mx-[8px] tw-cursor-pointer"
              @click="slide(1)"
              @click.native="mixpanel.track('tap_faq_next_arrow')"
              aria-label="rightArrow"
            >
              <font-awesome-icon
                class="arrow tw-border-[1px] tw-border-solid tw-border-[#3d3d3d26] tw-rounded-[15px] tw-h-[25px] tw-w-[25px] tw-p-[10px] tw-text-[#fff] tw-bg-black-900"
                icon="arrow-right"
                id="rightArrow"
              />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import AspectRatio from "@/components/utils/AspectRatio.vue";
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import faqImage from "@/assets/images/faq/FAQ.svg";

const TOTAL_FAQ_IN_SLIDE = 5;

export default {
  data() {
    return {
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
            "You will be greeted by HR when you visit us for an interview. There will be two rounds of interviews. Technical or Non-technical followed by an HR interview. ",
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
    FontAwesomeIcon,
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
        this.mixpanel.track("tap_faq_collapse");
      } else {
        this.openList = true;
        this.mixpanel.track("tap_faq_expand");
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
