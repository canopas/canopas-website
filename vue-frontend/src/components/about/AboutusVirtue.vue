<template>
  <section>
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center tw-gap-[1.5rem] md:tw-gap-[2.5rem] tw-py-11 md:tw-py-20"
    >
      <div
        class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.25rem] md:tw-leading-[5.15625rem] tw-tracking-[-4%]"
      >
        Our Why?
      </div>
      <div
        class="tw-w-[95%] lg:tw-w-[62%] tw-text-center tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.875rem] tw-leading-[1.5rem] md:tw-leading-[2.8125rem] tw-tracking-[-4%]"
      >
        To put humanity into the better position compared to what it was before
        our presence in the world.
      </div>
    </div>
    <div
      class="tw-from-[#ffad9326] tw-to-[#f2709c26] tw-bg-gradient-247 tw-px-0 tw-py-[100px] tw-overflow-hidden tw-text-center tw-rounded-tr-none tw-rounded-bl-none tw-rounded-tl-[100px] tw-rounded-br-[100px] sm:tw-rounded-tl-[150px] sm:tw-rounded-br-[150px] md:tw-rounded-tl-[200px] md:tw-rounded-br-[200px] lg:tw-rounded-tl-[300px] lg:tw-rounded-br-[300px]"
    >
      <div
        class="tw-text-center tw-text-black-900 tw-font-inter-semibold tw-text-[1.125rem] md:tw-text-[2.5rem] tw-leading-[1.36125rem] md:tw-leading-[3.025625rem]"
      >
        Canopas is built on the 10 virtues
      </div>
      <div class="tw-text-center tw-container tw-mt-20 sm:tw-p-0">
        <div class="flex-div tw-flex tw-flex-wrap tw-m-auto lg:tw-w-[85%]">
          <div
            class="tw-relative tw-flex-[100%] md:tw-flex-[33%]"
            v-for="virtue in virtues.slice(0, 3)"
            :key="virtue.id"
            @touchstart.passive="activeIndex = virtue.id"
            @hover.native="mixpanel.track('hover_virtue')"
          >
            <div
              class="tw-flex tw-flex-col tw-items-center tw-my-[5%] tw-mx-auto tw-py-0 tw-px-[10%] sm:tw-my-[2%] sm:tw-px-0 md:tw-my-[10%]"
            >
              <div
                class="image tw-rounded-[50%] tw-bg-white tw-relative tw-h-0 tw-w-1/2 tw-pb-[50%] tw-mb-[5%] sm:tw-w-[30%] sm:tw-pb-[30%] md:tw-w-1/2 md:tw-pb-[50%]"
              >
                <img
                  class="tw-absolute tw-w-[95%] tw-h-[95%] tw-left-0"
                  :src="virtue.image"
                  :alt="virtue.title"
                  loading="lazy"
                />
              </div>
              <div
                class="tw-w-3/5 tw-mb-[5%] tw-text-[#e8505b] tw-font-inter-regular tw-text-[1.3125rem] tw-leading-[1.588125rem]"
              >
                {{ virtue.title }}
              </div>
            </div>
            <div
              class="hover:tw-opacity-100 tw-absolute tw-translate-x-[-50%] tw-transition tw-duration-[0.4s] tw-ease-in-out tw-h-auto tw-w-[290px] tw-left-[50%] tw-top-0 tw-bottom-0 tw-opacity-0 tw-bg-white tw-p-[2%] tw-m-auto tw-rounded-[15px] tw-flex tw-flex-col tw-justify-center sm:tw-w-1/2 md:tw-w-[90%] 2xl:tw-p-[3%]"
              :class="activeIndex == virtue.id ? 'tw-opacity-100' : ''"
            >
              <div
                class="tw-text-[#e8505b] tw-mb-[3%] tw-w-full tw-font-inter-regular md:tw-mb-0 tw-text-[1.3125rem] tw-leading-[1.588125rem] 2xl:tw-mb-[2%]"
              >
                {{ virtue.title }}
              </div>
              <div
                class="normal-2-text tw-mt-2.5 tw-mb-[3%] tw-w-full md:tw-mb-0 2xl:tw-mb-[2%]"
              >
                {{ virtue.description }}
              </div>
            </div>
          </div>
        </div>
        <div
          class="flex-div tw-flex tw-flex-wrap tw-m-auto md:tw-w-[70%] lg:tw-w-[60%]"
        >
          <div
            class="tw-relative tw-flex-[100%] md:tw-flex-[33%]"
            v-for="virtue in virtues.slice(3, 5)"
            :key="virtue.id"
            @touchstart.passive="activeIndex = virtue.id"
            @hover.native="mixpanel.track('hover_virtue')"
          >
            <div
              class="tw-flex tw-flex-col tw-items-center tw-my-[5%] tw-mx-auto tw-py-0 tw-px-[10%] sm:tw-my-[2%] sm:tw-px-0 md:tw-my-[10%]"
            >
              <div
                class="image tw-rounded-[50%] tw-bg-white tw-relative tw-h-0 tw-w-1/2 tw-pb-[50%] tw-mb-[5%] sm:tw-w-[30%] sm:tw-pb-[30%] md:tw-w-1/2 md:tw-pb-[50%]"
              >
                <img
                  class="tw-absolute tw-w-[95%] tw-h-[95%] tw-left-0"
                  :src="virtue.image"
                  :alt="virtue.title"
                  loading="lazy"
                />
              </div>
              <div
                class="tw-w-3/5 tw-mb-[5%] tw-text-[#e8505b] tw-font-inter-regular tw-text-[1.3125rem] tw-leading-[1.588125rem]"
              >
                {{ virtue.title }}
              </div>
            </div>
            <div
              class="hover:tw-opacity-100 tw-absolute tw-translate-x-[-50%] tw-transition tw-duration-[0.4s] tw-ease-in-out tw-h-auto tw-w-[290px] tw-left-[50%] tw-top-0 tw-bottom-0 tw-opacity-0 tw-bg-white tw-p-[2%] tw-m-auto tw-rounded-[15px] tw-flex tw-flex-col tw-justify-center sm:tw-w-1/2 md:tw-w-[90%] 2xl:tw-p-[3%]"
              :class="activeIndex == virtue.id ? 'tw-opacity-100' : ''"
            >
              <div
                class="tw-text-[#e8505b] tw-mb-[3%] tw-w-full md:tw-mb-0 tw-font-inter-regular tw-text-[1.3125rem] tw-leading-[1.588125rem] 2xl:tw-mb-[2%]"
              >
                {{ virtue.title }}
              </div>
              <div
                class="normal-2-text tw-mt-2.5 tw-mb-[3%] tw-w-full md:tw-mb-0 2xl:tw-mb-[2%]"
              >
                {{ virtue.description }}
              </div>
            </div>
          </div>
        </div>
        <div class="flex-div tw-flex tw-flex-wrap tw-m-auto lg:tw-w-[85%]">
          <div
            class="tw-relative tw-flex-[100%] md:tw-flex-[33%]"
            v-for="virtue in virtues.slice(5, 8)"
            :key="virtue.id"
            @touchstart.passive="activeIndex = virtue.id"
            @hover.native="mixpanel.track('hover_virtue')"
          >
            <div
              class="tw-flex tw-flex-col tw-items-center tw-my-[5%] tw-mx-auto tw-py-0 tw-px-[10%] sm:tw-my-[2%] sm:tw-px-0 md:tw-my-[10%]"
            >
              <div
                class="image tw-rounded-[50%] tw-bg-white tw-relative tw-h-0 tw-w-1/2 tw-pb-[50%] tw-mb-[5%] sm:tw-w-[30%] sm:tw-pb-[30%] md:tw-w-1/2 md:tw-pb-[50%]"
              >
                <img
                  class="tw-absolute tw-w-[95%] tw-h-[95%] tw-left-0"
                  :src="virtue.image"
                  :alt="virtue.title"
                  loading="lazy"
                />
              </div>
              <div
                class="tw-w-3/5 tw-mb-[5%] tw-text-[#e8505b] tw-font-inter-regular tw-text-[1.3125rem] tw-leading-[1.588125rem]"
              >
                {{ virtue.title }}
              </div>
            </div>
            <div
              class="hover:tw-opacity-100 tw-absolute tw-translate-x-[-50%] tw-transition tw-duration-[0.4s] tw-ease-in-out tw-h-auto tw-w-[290px] tw-left-[50%] tw-top-0 tw-bottom-0 tw-opacity-0 tw-bg-white tw-p-[2%] tw-m-auto tw-rounded-[15px] tw-flex tw-flex-col tw-justify-center sm:tw-w-1/2 md:tw-w-[90%] 2xl:tw-p-[3%]"
              :class="activeIndex == virtue.id ? 'tw-opacity-100' : ''"
            >
              <div
                class="tw-text-[#e8505b] tw-mb-[3%] tw-w-full tw-font-inter-regular md:tw-mb-0 tw-text-[1.3125rem] tw-leading-[1.588125rem] 2xl:tw-mb-[2%]"
              >
                {{ virtue.title }}
              </div>
              <div
                class="normal-2-text tw-mt-2.5 tw-mb-[3%] tw-w-full md:tw-mb-0 2xl:tw-mb-[2%]"
              >
                {{ virtue.description }}
              </div>
            </div>
          </div>
        </div>
        <div
          class="flex-div tw-flex tw-flex-wrap tw-m-auto md:tw-w-[70%] lg:tw-w-[60%]"
        >
          <div
            class="tw-relative tw-flex-[100%] md:tw-flex-[33%]"
            v-for="virtue in virtues.slice(8, 10)"
            :key="virtue.id"
            @touchstart.passive="activeIndex = virtue.id"
            @hover.native="mixpanel.track('hover_virtue')"
          >
            <div
              class="tw-flex tw-flex-col tw-items-center tw-my-[5%] tw-mx-auto tw-py-0 tw-px-[10%] sm:tw-my-[2%] sm:tw-px-0 md:tw-my-[10%]"
            >
              <div
                class="image tw-rounded-[50%] tw-bg-white tw-relative tw-h-0 tw-w-1/2 tw-pb-[50%] tw-mb-[5%] sm:tw-w-[30%] sm:tw-pb-[30%] md:tw-w-1/2 md:tw-pb-[50%]"
              >
                <img
                  class="tw-absolute tw-w-[95%] tw-h-[95%] tw-left-0"
                  :src="virtue.image"
                  :alt="virtue.title"
                  loading="lazy"
                />
              </div>
              <div
                class="tw-w-3/5 tw-mb-[5%] tw-text-[#e8505b] tw-font-inter-regular tw-text-[1.3125rem] tw-leading-[1.588125rem]"
              >
                {{ virtue.title }}
              </div>
            </div>
            <div
              class="hover:tw-opacity-100 tw-absolute tw-translate-x-[-50%] tw-transition tw-duration-[0.4s] tw-ease-in-out tw-h-auto tw-w-[290px] tw-left-[50%] tw-top-0 tw-bottom-0 tw-opacity-0 tw-bg-white tw-p-[2%] tw-m-auto tw-rounded-[15px] tw-flex tw-flex-col tw-justify-center sm:tw-w-1/2 md:tw-w-[90%] 2xl:tw-p-[3%]"
              :class="activeIndex == virtue.id ? 'tw-opacity-100' : ''"
            >
              <div
                class="tw-text-[#e8505b] tw-mb-[3%] tw-w-full md:tw-mb-0 tw-font-inter-regular tw-text-[1.3125rem] tw-leading-[1.588125rem] 2xl:tw-mb-[2%]"
              >
                {{ virtue.title }}
              </div>
              <div
                class="normal-2-text tw-mt-2.5 tw-mb-[3%] tw-w-full md:tw-mb-0 2xl:tw-mb-[2%]"
              >
                {{ virtue.description }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script type="module">
import virtue1 from "@/assets/images/virtues/jobs_canopas_virtue_1.svg";
import virtue2 from "@/assets/images/virtues/jobs_canopas_virtue_2.svg";
import virtue3 from "@/assets/images/virtues/jobs_canopas_virtue_3.svg";
import virtue4 from "@/assets/images/virtues/jobs_canopas_virtue_4.svg";
import virtue5 from "@/assets/images/virtues/jobs_canopas_virtue_5.svg";
import virtue6 from "@/assets/images/virtues/jobs_canopas_virtue_6.svg";
import virtue7 from "@/assets/images/virtues/jobs_canopas_virtue_7.svg";
import virtue8 from "@/assets/images/virtues/jobs_canopas_virtue_8.svg";
import virtue9 from "@/assets/images/virtues/jobs_canopas_virtue_9.svg";
import virtue10 from "@/assets/images/virtues/jobs_canopas_virtue_10.svg";

export default {
  data() {
    return {
      activeIndex: 0,
      virtues: [
        {
          id: 1,
          title: "We make big bold bets",
          image: virtue1,
          description:
            "Sometimes we fail, but failure makes us smarter. We get back up, we make the next bet, and we go!",
        },
        {
          id: 2,
          title: "We do what matters the most",
          image: virtue2,
          description:
            "Almost always, what we do matters so much compared to how we do it. There's no point doing bad things in a very efficient way.",
        },
        {
          id: 3,
          title: "We tell the truth even if it hurts",
          image: virtue3,
          description:
            "Being honest to others, even if it would hurt them, just shows them how much they can trust you. It’s good be honest, but don’t be rude.",
        },
        {
          id: 4,
          title: "We value time",
          image: virtue4,
          description:
            "'The one that values his time can value the time of others as well'- well said. Always respect others' time even more than yours.",
        },
        {
          id: 5,
          title: "We move fast and break things",
          image: virtue5,
          description:
            "Making mistakes is a natural consequence of innovation in a highly competitive and complex environment..",
        },
        {
          id: 6,
          title: "We celebrate differences",
          image: virtue6,
          description:
            "We stand apart from the average. We ensure people of diverse backgrounds feel welcome. We encourage different opinions and approaches to be heard, and then we come together and build.",
        },
        {
          id: 7,
          title: "We are owners",
          image: virtue7,
          description:
            "We seek out problems and we solve them. We help each other and those who matter to us. We have a bias for action and accountability.",
        },
        {
          id: 8,
          title: "We don't have titles",
          image: virtue8,
          description: "We don't need titles to be true leaders.",
        },
        {
          id: 9,
          title: "We value ideas over hierarchy",
          image: virtue9,
          description:
            "We believe that the best ideas can come from anywhere, both inside and outside our company. Our job is to seek out those ideas, to shape and improve them through candid debate, and to take them from concept to action.",
        },
        {
          id: 10,
          title: "We are playing the infinite game",
          image: virtue10,
          description:
            "An infinite mindset is the recognition that there is no practical end to our work. The goal is not to win or be the best in the infinite game; it is to strive to be better, to experience constant improvement.",
        },
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
