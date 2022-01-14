<template>
  <div class="consultant-container normal-text">
    <div class="image-container">
      <img
        class="consultant-image abs-fit-parent"
        loading="lazy"
        :src="steps[currentStep].image"
      />
    </div>
    <div class="steps-container">
      <div v-for="(step, i) in steps" :key="step" class="step-item">
        <div class="step-header normal-text text-left" @click="onStepClick(i)">
          <div class="step-number">{{ i + 1 }}</div>
          <div class="step-title">{{ step.title }}</div>
          <font-awesome-icon
            v-if="i != currentStep"
            class="step-plus-icon"
            icon="plus"
          />
        </div>
        <collapse-transition>
          <div
            v-if="i == currentStep"
            :key="step.descriptions"
            class="step-description card-body"
          >
            {{ step.descriptions.join("\n\n") }}
          </div>
        </collapse-transition>
        <div class="step-divider" v-if="i != steps.length"></div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import step1Image from "@/assets/images/home/step1.webp";
import step2Image from "@/assets/images/home/step2.webp";
import step3Image from "@/assets/images/home/step3.webp";
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  components: {
    CollapseTransition,
    FontAwesomeIcon,
  },
  data() {
    return {
      steps: [
        {
          image: step1Image,
          title: "FREE Consultation",
          descriptions: [
            "For Canopas, WHY do you want to develop a product is more IMPORTANT than what do you want to develop?",
            "We are very good at what we do. But, there is also so much we don’t know about you and your business.",
            "The only person who knows more about you and your business is none other than You, Yourself.",
            "Let’s jump on a call or chat for 15 minutes",
          ],
        },
        {
          image: step2Image,
          title: "100% Money back Guarantee",
          descriptions: [
            "A Journey of Thousand Miles Begins with a Single Step - Confucius.",
            "Our first step of the journey consists of TWO WEEKS of Work. If you like the output, the journey will begin.",
            "After a successful collaboration for a couple of weeks, we will lay out a proper plan for your product and start working on it.",
            "If you decide to cancel, you will receive a full refund. We offer a 100% Money-back guarantee. No questions asked.",
          ],
        },
        {
          image: step3Image,
          title: "Journey Begins",
          descriptions: [
            "We will start on a small scale by defining a couple of goals or features based on your requirements. Then, on every next iteration, we design, code, test, and add to your existing product.",
            "The result of this process is that every cycle produces software that is complete in itself and has more features and capabilities than the previous one.",
            "You will get this updated version of your product every week.",
            "This way, we always know if we’re on the right track. We can learn early about the problems without losing time and making costly mistakes.\nOf Course, we provide hassle-free support for your apps after launching your product.",
          ],
        },
      ],
      currentStep: 0,
    };
  },
  mounted() {},
  methods: {
    onStepClick(index) {
      this.currentStep = index;
    },
  },
};
</script>

<style scoped>
.consultant-container {
  margin: 48px 10%;
  display: flex;
  flex-direction: column;
  border-radius: 0.9rem;
  border: 1px solid #e2e2e2;
  overflow: hidden;
}

.image-container {
  position: relative;
  width: 100%;
  padding-top: 66.5%;
}

.consultant-image {
  object-fit: cover;
}

.steps-container {
  display: flex;
  flex-direction: column;
  width: 100%;
  position: relative;
}

.step-item {
  display: flex;
  flex-direction: column;
  width: 100%;
  padding: 0 16px;
  white-space: pre-wrap;
}

.step-divider {
  background: #e2e2e2;
  height: 1px;
  width: 100%;
}

.step-header {
  display: flex;
  flex-direction: row;
  width: 100%;
  align-items: center;
  padding: 20px 16px 20px 16px;
  cursor: pointer;
}

.step-description {
  padding-bottom: 20px;
  line-height: 1.4rem;
}

.step-number {
  background: linear-gradient(180deg, #ff9472 0%, #f2709c 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  font-weight: 700;
  border: 1px solid rgba(0, 0, 0, 0.15);
  border-radius: 0.9rem;
  padding: 5px 0.9rem;
  margin-right: 10px;
}

.step-title {
  font-style: normal;
  font-weight: 700;
  color: #3d3d3d;
}

.step-plus-icon {
  margin-left: auto;
  order: 2;
  color: rgba(61, 61, 61, 0.5);
}

@media (min-width: 992px) {
  .consultant-container {
    flex-direction: row;
  }

  .image-container {
    position: relative;
    width: 100%;
    padding-top: 0;
  }

  .step-title {
    font-size: 1.5rem;
    line-height: 1.8rem;
    font-size: normal;
  }

  .step-description {
    line-height: 1.8rem;
  }
}

@media (min-width: 1200px) {
  .consultant-container {
    margin: 48px 15%;
  }

  .step-plus-icon {
    font-size: 1.5rem;
  }
}
</style>
