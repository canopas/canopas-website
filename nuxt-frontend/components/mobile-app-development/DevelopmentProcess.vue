<template>
  <section class="lg:bg-light-pink-gradient-background mt-16 lg:mt-60">
    <!-- MOBILE UI START  -->
    <div class="container flex flex-col items-center lg:hidden">
      <p class="text-center mobile-header-2 text-black-87">
        Development process
      </p>
      <div class="cards relative h-full w-full pt-[34.375rem]" ref="cards">
        <div
          class="absolute bottom-0 left-0 right-0 top-0 m-auto h-[22.80088rem] w-[22.625rem] rounded-[1.25rem] border border-pink-300 bg-deep-charcoal pt-5 transition-all duration-500 sm:w-[24.625rem]"
          :style="{ zIndex: process.zIndex }"
          v-for="(process, index) in processes"
          :key="index"
          @touchstart="onTouchStart"
          @touchmove="onTouchMove"
          @touchend="onTouchEnd"
        >
          <div class="flex flex-col px-4">
            <div class="flex items-center gap-2">
              <div>
                <img
                  :src="process.image"
                  alt="process"
                  class="h-12 w-12 sm:h-14 sm:w-14"
                />
              </div>
              <p class="sub-h1-semibold text-white">
                {{ process.title }}
              </p>
            </div>
            <p class="mt-5 text-center sub-h3-regular text-white-core-80">
              {{ process.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
    <!-- MOBILE UI END  -->

    <!-- Desktop UI START -->
    <div
      class="container hidden h-[61.563rem] overflow-hidden lg:block lg:h-[70.813rem] xl:h-[77.813rem]"
    >
      <p class="text-center desk-header-2">Development Process</p>
      <div class="flex flex-row-reverse items-center 2xl:gap-28">
        <transition-group>
          <div class="w-3/5 xl:w-1/2">
            <div
              v-for="(process, index) in processes"
              :key="process"
              class="flex flex-col"
            >
              <div class="mb-4 mt-12 flex items-center gap-2">
                <div
                  :class="
                    activeIndex == index ? 'secondary-color' : 'border-text '
                  "
                  class="desk-header-1"
                >
                  {{ process.id }}
                </div>
                <div
                  @mouseover="(activeProcess = process), (activeIndex = index)"
                  @touchstart.passive="
                    (activeIndex = index), (activeProcess = process)
                  "
                  class="relative w-max cursor-pointer rounded-r-[15px] bg-pink-80 sub-h2-medium text-black-87 py-6 pl-4 pr-8 xl:py-9 xl:pr-12"
                  :class="process.className"
                >
                  <span
                    class="absolute right-0 top-0 h-3 w-3 rounded-full bg-black-60 xl:h-5 xl:w-5"
                  ></span>
                  {{ process.title }}
                </div>
              </div>
              <collapse-transition>
                <div
                  v-if="activeIndex == index"
                  :key="process.description"
                  class="ml-[2.4rem] xl:ml-[3.2rem] mobile-header-2-regular text-black-60"
                  v-html="process.description"
                ></div>
              </collapse-transition>
            </div></div
        ></transition-group>
        <div class="w-2/5 xl:w-1/2">
          <img
            v-if="activeProcess !== null"
            :src="activeProcess.desktopImage"
            alt="process"
            class="h-full w-full object-contain"
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
import desktop_process1 from "@/assets/images/mobile-app-development/development-process/desktop/process-1.gif";
import desktop_process2 from "@/assets/images/mobile-app-development/development-process/desktop/process-2.gif";
import desktop_process3 from "@/assets/images/mobile-app-development/development-process/desktop/process-3.gif";
import desktop_process4 from "@/assets/images/mobile-app-development/development-process/desktop/process-4.gif";
import desktop_process5 from "@/assets/images/mobile-app-development/development-process/desktop/process-5.gif";
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
          desktopImage: desktop_process1,
        },
        {
          id: 2,
          title: "UI/UX Design",
          description:
            "Once we've nailed down the essence of your idea, our design team steps in. Crafting an immersive user interface and a seamless user experience isn't just about aesthetics - it's about designing an app your users will love to use. We strike the perfect balance between functionality and visual appeal, making your app irresistible to users.",
          image: process2,
          desktopImage: desktop_process2,
        },
        {
          id: 3,
          title: "Crafting Intelligent Solutions",
          description:
            "Now, it's time to breathe life into the design. Our team of seasoned developers focuses on creating an app that's fast, thoughtful, and smart. We meticulously code to ensure that your app runs smoothly, provides quick response times, and offers intelligent solutions to your users.",
          image: process3,
          desktopImage: desktop_process3,
        },
        {
          id: 4,
          title: "Taking Flight: App Deployment",
          description:
            "Once we're confident of the app's performance and reliability, it's time to introduce it to the world. We take care of the entire deployment process, ensuring your app finds its rightful place on the App Store or Google Play. But our journey doesn't end there.",
          image: process4,
          desktopImage: desktop_process4,
        },
        {
          id: 5,
          title: "Enduring Excellence: Maintenance and Evolution",
          className: "lg:!w-[70%] xl:!w-[85%]",
          description:
            "The lifecycle of an app doesn't end after launching the product. As your business grows, so should your app. We provide ongoing maintenance and continuous feature enhancement to ensure your app stays relevant, scalable, and in tune with your evolving business needs.",
          image: process5,
          desktopImage: desktop_process5,
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
        if (activeCard) {
          activeCard.style.transform =
            "translate(-300px, -300px) rotate(-30deg)";
          setTimeout(() => {
            activeCard.classList.add("go-away");
            this.activeCard++;
            this.arrangeCards();
          }, 300);
        }
      }
    },

    prevCard() {
      if (this.activeCard > 0) {
        const activeCard = this.$refs.cards.children[this.activeCard - 1];
        if (activeCard) {
          activeCard.classList.remove("go-away");
          this.activeCard--;
          this.arrangeCards();
        }
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
  transition: transform 0.5s;
}
</style>
