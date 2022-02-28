<template>
  <div class="container">
    <div class="row">
      <div
        class="text-center header-text canopas-gradient-text title-mobile mb-5"
      >
        <span class="underline-text">Freq</span>uently Asked Questions
      </div>
      <div class="col-md-4 image">
        <img
          :src="faqImage"
          loading="lazy"
          alt="frequently-asked-questions-image"
        />
      </div>
      <div class="col-md-8">
        <div class="faq-content">
          <div class="header-text canopas-gradient-text title-desktop">
            <span class="underline-text">Freq</span>uently Asked Questions
          </div>
          <transition-group tag="div" :name="'faq-' + faqTransitionName">
            <div class="mt-5 faq-section">
              <div
                class="mt-4"
                v-for="faq in faqs.slice(sliceFrom, sliceTo)"
                :key="faq"
              >
                <div
                  class="faq-container jobs-normal-text"
                  @click="expandListItem(faq.id)"
                >
                  <div class="faq-header">
                    <div class="faq-icon gradient-icon">
                      <font-awesome-icon
                        class="plus-icon"
                        :icon="
                          openList && faq.id == currentIndex ? 'minus' : 'plus'
                        "
                      />
                    </div>
                    <div class="faq-question">{{ faq.question }}</div>
                  </div>
                  <collapse-transition>
                    <div
                      class="faq-header"
                      v-if="openList && faq.id == currentIndex"
                      :key="faq.answer"
                    >
                      <div class="faq-icon"></div>
                      <div class="normal-2-text faq-answer mt-3">
                        {{ faq.answer }}
                      </div>
                    </div>
                  </collapse-transition>
                  <div class="faq-divider" v-if="faq.id != faq.length"></div>
                </div>
              </div>
            </div>
          </transition-group>
          <div class="mt-4 text-center">
            <button
              type="button"
              :disabled="!isActivePrev"
              class="clients-indicators"
              @click="slide(-1)"
            >
              <font-awesome-icon
                class="arrow"
                icon="arrow-left"
                id="leftArrow"
              />
            </button>
            <button
              type="button"
              :disabled="!isActiveNext"
              class="clients-indicators"
              @click="slide(1)"
            >
              <font-awesome-icon
                class="arrow"
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
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import faqImage from "@/assets/images/faq/FAQ.svg";

const TOTAL_FAQ_IN_SLIDE = 6;

export default {
  data() {
    return {
      faqImage: faqImage,
      faqs: [
        {
          id: 1,
          question: "What are the top skills you look for in a candidate?",
          answer:
            "It mostly depends on the job you’re applying for. Usually, we look for a deep desire to learn, willingness to work hard, always looking for a challenge, not afraid to leave your comfort zone, and a desire to excel.",
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
            "You will receive status about your application within 24 hours.",
        },
        {
          id: 5,
          question: "How can I prepare?",
          answer:
            "Write answers for the engineering, HR, and marketing team separately.",
        },
        {
          id: 6,
          question: "What should I wear?",
          answer: "Whatever makes you comfortable. We don’t have a dress code.",
        },
        {
          id: 7,
          question: "What does the hiring process look like?",
          answer:
            "You will be greeted by HR when you visit us for an interview. There will be two rounds of interviews. Technical or Non-technical followed by an HR interview. ",
        },
        {
          id: 8,
          question: "How much time will the interview take?",
          answer: "Max 1 hour.",
        },
        {
          id: 9,
          question: "How should I apply?",
          answer:
            "Browse through all openings available on this website and apply for positions you’re looking for. If any job doesn’t match with your skills, you can check back later.",
        },
        {
          id: 10,
          question: "How much experience do I need?",
          answer:
            "You will find more details about it in the job descriptions page.",
        },
        {
          id: 11,
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
  },
  methods: {
    expandListItem(index) {
      if (this.previousIndex == index && this.openList) {
        this.openList = false;
      } else {
        this.openList = true;
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
      this.isActiveNext = this.sliceTo <= this.faqs.length;
      this.current = this.getRoundedIndex(dir);
    },
    getRoundedIndex(diff) {
      var len = this.faqs.length / 5;
      return (
        (this.current + (diff % Math.ceil(len)) + Math.ceil(len)) %
        Math.ceil(len)
      );
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin-bottom: 10rem;
}
.row {
  margin: 0 2%;

  * {
    padding: 0;
  }
}

.image {
  padding: 0 20%;
}

.underline-text {
  text-underline-offset: 0.5rem;
}

.faq-container {
  border: 1px solid #e2e2e2;
  overflow: hidden;
  background: rgba(61, 61, 61, 0.03);
  box-shadow: 0px 1px 6px 1px rgba(0, 0, 0, 0.25);
  border-radius: 15px;
  padding: 15px;
}

.faq-header {
  display: flex;
  flex-direction: row;
  cursor: pointer;
}

.faq-icon {
  width: 10%;
}

.gradient-icon {
  color: #f2709c;
}

.faq-question {
  color: #3d3d3d;
  font-weight: 700;
  width: 90%;
}

.faq-answer {
  width: 90%;
}

.clients-indicators {
  background: none;
  border: none;
  margin: 0 8px;
  cursor: pointer;
}

button:disabled,
button[disabled] {
  opacity: 0.2;
}

.faq-section {
  min-height: 33rem;
  height: auto;
}

.arrow {
  border: 1px solid rgba(61, 61, 61, 0.15);
  border-radius: 15px;
  height: 45px;
  width: 45px;
  padding: 10px;
  color: #fff;
  background: #3d3d3d;
}

.title-mobile {
  display: block;
}
.title-desktop {
  display: none;
}

/* NEXT QUESTION */
.faq-next-enter-active,
.faq-next-leave-active {
  transition: transform 0.5s ease-in-out;
}
.faq-next-enter-active {
  transform: translate(150%);
}
.faq-next-enter-to {
  transform: translate(0%);
}
.faq-next-leave-to {
  transform: translate(-150%);
}

/* PREVIOUS QUESTION */
.faq-prev-enter-active,
.faq-prev-leave-active {
  transition: transform 0.5s ease-in-out;
}
.faq-prev-enter-active {
  transform: translate(-150%);
}
.faq-prev-enter-to {
  transform: translate(0%);
}
.faq-prev-leave-to {
  transform: translate(150%);
}

@include media-breakpoint-up(sm) {
  .faq-container {
    padding: 25px;
  }

  .faq-section {
    min-height: 40rem;
  }

  .underline-text {
    text-underline-offset: 1rem;
  }
}

@include media-breakpoint-up(md) {
  .row {
    margin: 0 6%;
    margin-left: -32px;
  }

  .faq-content {
    padding-left: 60px;
  }

  .image {
    transform: translateY(15rem);
    padding: 0;
  }

  .title-mobile {
    display: none;
  }
  .title-desktop {
    display: block;
  }
}

@include media-breakpoint-up(lg) {
  .image {
    transform: translateY(20rem);
  }
}

@include media-breakpoint-up(xl) {
  .image {
    transform: translateY(10rem);
  }
}
</style>
