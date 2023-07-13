<template>
  <section class="tw-py-[50px] md:tw-pb-[100px] xl:tw-pt-[100px]">
    <!-- MOBILE UI START  -->
    <div class="container tw-flex tw-flex-col tw-items-center md:tw-hidden">
      <p
        class="tw-w-[60%] tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.269rem] tw-text-black-core/[0.85] sm:tw-w-full"
      >
        DEVELOPMENT PROCESS
      </p>
      <div
        class="cards tw-relative tw-h-full tw-w-full tw-pt-[550px]"
        ref="cards"
      >
        <div
          class="tw-absolute tw-bottom-0 tw-left-0 tw-right-0 tw-top-0 tw-m-auto tw-h-[22.80088rem] tw-w-[22.625rem] tw-rounded-[1.25rem] tw-border tw-border-[#f77881] tw-bg-[#3d3d3d] tw-pt-[20px] tw-transition-all tw-duration-500 sm:tw-w-[24.625rem]"
          :style="{ zIndex: process.zIndex }"
          v-for="(process, index) in processes"
          :key="index"
          @touchstart="onTouchStart"
          @touchmove="onTouchMove"
          @touchend="onTouchEnd"
        >
          <div class="tw-flex tw-flex-col tw-px-4">
            <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
              <div>
                <img
                  :src="process.image"
                  alt="process"
                  class="tw-h-12 tw-w-12 sm:tw-h-14 sm:tw-w-14"
                />
              </div>
              <p
                class="tw-font-inter-semibold tw-text-[1.125rem] tw-leading-[1.688rem] tw-text-[#F77881]"
              >
                {{ process.title }}
              </p>
            </div>
            <p
              class="tw-mt-[1.25rem] tw-text-center tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-text-white"
            >
              {{ process.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
    <!-- MOBILE UI END  -->

    <!-- Desktop UI START -->
    <div class="tw-container tw-hidden md:tw-block">
      <p
        class="tw-text-center tw-font-inter-bold tw-text-[3.125rem] tw-leading-[3.625rem] tw-text-black-core/[0.85] xl:tw-text-[4.125rem] xl:tw-leading-[6.188rem]"
      >
        Development Process
      </p>
      <div
        class="tw-flex tw-flex-row-reverse tw-items-center 2xl:tw-gap-[7rem]"
      >
        <transition-group>
          <div class="tw-w-[60%] xl:tw-w-[50%]">
            <div v-for="(process, index) in processes" :key="process">
              <div class="tw-mt-12">
                <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
                  <div
                    :class="
                      activeIndex == index
                        ? 'gradient-border-text'
                        : 'border-text '
                    "
                    class="tw-font-futura-bold tw-text-[2.875rem] tw-leading-[3.375rem] xl:tw-text-[4.125rem] xl:tw-leading-[6.188rem]"
                  >
                    {{ process.id }}
                  </div>
                  <div
                    @mouseover="
                      (activeProcess = process), (activeIndex = index)
                    "
                    @touchstart.passive="
                      (activeIndex = index), (activeProcess = process)
                    "
                    class="tw-relative tw-w-max tw-cursor-pointer tw-rounded-r-[15px] tw-bg-[#FFF1ED] tw-py-6 tw-pl-4 tw-pr-8 tw-font-inter-medium tw-text-[1.25rem] tw-leading-[1.625rem] tw-text-black-core xl:tw-py-9 xl:tw-pr-12 xl:tw-text-[1.5rem] xl:tw-leading-[2.25rem]"
                    :class="process.className"
                  >
                    <span
                      class="tw-absolute tw-right-0 tw-top-0 tw-h-3 tw-w-3 tw-rounded-full tw-bg-[#3D3D3D] xl:tw-h-5 xl:tw-w-5"
                    ></span>
                    {{ process.title }}
                  </div>
                </div>
                <collapse-transition>
                  <div
                    v-if="activeIndex == index"
                    :key="process.description"
                    class="tw-ml-12 tw-mt-[1rem] tw-font-inter-regular tw-text-[1.25rem] tw-leading-[1.625rem] tw-text-black-core/[0.70] xl:tw-text-[1.5rem] xl:tw-leading-[2.25rem]"
                    v-html="process.description"
                  ></div>
                </collapse-transition>
              </div>
            </div></div
        ></transition-group>
        <div class="tw-w-[40%] xl:tw-w-[50%]">
          <img
            v-if="activeProcess !== null"
            :src="activeProcess.image"
            alt="process"
            class="tw-h-full tw-w-full tw-object-contain"
          />
        </div>
      </div>
    </div>
    <!-- DESKTOP UI END -->
  </section>
</template>

<script type="module">
import process1 from "@/assets/images/mobile-app-development/development-process/process-1.gif";
import process2 from "@/assets/images/mobile-app-development/development-process/process-2.gif";
import process3 from "@/assets/images/mobile-app-development/development-process/process-3.gif";
import process4 from "@/assets/images/mobile-app-development/development-process/process-4.gif";
import process5 from "@/assets/images/mobile-app-development/development-process/process-5.gif";
import CollapseTransition from "@ivanv/vue-collapse-transition/src/CollapseTransition.vue";

export default {
  data() {
    return {
      touchStartX: 0,
      touchEndX: 0,
      minSwipeDistance: 50,
      activeProcess: null,
      activeCard: 0,
      totalProcesses: 5,
      activeIndex: 0,
      processes: [
        {
          id: 1,
          title: "Idea Incubation",
          description:
            "Our journey begins with your vision. We don't just listen; we delve deep into your idea, understanding the 'why' behind it. What's the core problem your app aims to solve? Who's your target audience? These insights help us shape a mobile app solution that truly resonates with your goals.",
          image: process1,
        },
        {
          id: 2,
          title: "UI/UX Design",
          description:
            "Once we've nailed down the essence of your idea, our design team steps in. Crafting an immersive user interface and a seamless user experience isn't just about aesthetics - it's about designing an app your users will love to use. We strike the perfect balance between functionality and visual appeal, making your app irresistible to users.",
          image: process2,
        },
        {
          id: 3,
          title: "Crafting Intelligent Solutions",
          description:
            "Now, it's time to breathe life into the design. Our team of seasoned developers focuses on creating an app that's fast, thoughtful, and smart. We meticulously code to ensure that your app runs smoothly, provides quick response times, and offers intelligent solutions to your users.",
          image: process3,
        },
        {
          id: 4,
          title: "Taking Flight: App Deployment",
          description:
            "Once we're confident of the app's performance and reliability, it's time to introduce it to the world. We take care of the entire deployment process, ensuring your app finds its rightful place on the App Store or Google Play. But our journey doesn't end there.",
          image: process4,
        },
        {
          id: 5,
          title: "Enduring Excellence: Maintenance and Evolution",
          className: "lg:!tw-w-[70%] xl:!tw-w-[85%]",
          description:
            "The lifecycle of an app doesn't end after launching the product. As your business grows, so should your app. We provide ongoing maintenance and continuous feature enhancement to ensure your app stays relevant, scalable, and in tune with your evolving business needs.",
          image: process5,
        },
      ],
    };
  },
  components: {
    CollapseTransition,
  },
  mounted() {
    this.generateCards();
    this.activeProcess = this.processes[0];
  },

  methods: {
    onTouchStart(event) {
      this.touchStartX = event.touches[0].clientX;
    },
    onTouchMove(event) {
      this.touchEndX = event.touches[0].clientX;
    },
    onTouchEnd() {
      const distance = this.touchEndX - this.touchStartX;
      if (distance > this.minSwipeDistance) {
        this.prevCard();
      } else if (distance < -this.minSwipeDistance) {
        this.nextCard();
      }
    },
    generateCards() {
      for (var i = 0; i < this.totalProcesses; i++) {
        this.processes[i].zIndex = i * -2; // Set z-index for the card
      }
      this.arrangeCards();
    },
    arrangeCards() {
      var order = 0;
      for (var i = this.activeCard; i < this.totalProcesses; i++) {
        var cardElement = this.$refs.cards.children[i];
        if (cardElement) {
          cardElement.style.transform = `translate3d(0px, 0px, ${
            order * -50
          }px) rotateX(0deg)`;
          order++;
        }
      }
    },

    nextCard() {
      if (this.activeCard < this.totalProcesses - 1) {
        const activeCard = this.$refs.cards.children[this.activeCard];
        if (activeCard) activeCard.classList.add("go-away");
        this.activeCard++;
        this.arrangeCards();
      }
    },
    prevCard() {
      if (this.activeCard > 0) {
        this.activeCard--;
        const activeCard = this.$refs.cards.children[this.activeCard];
        if (activeCard) activeCard.classList.remove("go-away");
        this.arrangeCards();
      }
    },
  },
};
</script>
<style scoped>
.gradient-border-text {
  -webkit-text-stroke: 0px;
  color: #f77881 !important;
}

.cards {
  perspective: 1455px;
  perspective-origin: 50% -50%;
}

.cards div.go-away {
  opacity: 0;
  left: -100px;
}
</style>
