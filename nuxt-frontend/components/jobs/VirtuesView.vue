<template>
  <div
    class="from-[#ffad9326] to-[#f2709c26] 3xl:outer-container bg-gradient-247 px-0 overflow-hidden text-center rounded-tr-none rounded-bl-none rounded-tl-[100px] rounded-br-[100px] sm:rounded-tl-[150px] sm:rounded-br-[150px] md:rounded-tl-[200px] md:rounded-br-[200px] lg:rounded-tl-[300px] lg:rounded-br-[300px]"
  >
    <div
      class="pt-10 lg:pt-14 mobile-header-3-semibold lg:desk-header-3 text-black-87"
    >
      Canopas is built on the 10 virtues
    </div>
    <div class="text-center container mt-8 mb-14 lg:mb-28 sm:p-0">
      <div class="flex flex-wrap m-auto lg:w-[85%]">
        <div
          class="relative basis-full lg:basis-[33%]"
          v-for="virtue in virtues.slice(0, 3)"
          :key="virtue.id"
          @touchstart.passive="(activeIndex = virtue.id)"
          @hover.native="$mixpanel.track('hover_virtue')"
        >
          <div
            class="flex flex-col items-center my-5 mx-auto py-0 px-10 sm:my-2 sm:px-0 lg:my-10"
          >
            <div
              class="rounded-full bg-white relative h-0 w-1/2 pb-[50%] mb-5 sm:w-[30%] sm:pb-[30%] lg:w-1/2 lg:pb-[50%]"
            >
              <img
                class="absolute w-[95%] h-[95%] left-0"
                :src="virtue.image"
                :alt="virtue.title"
              />
            </div>
            <div
              class="w-3/5 sub-h1-semibold secondary-color lg:mobile-header-3-medium lg:text-black-87"
            >
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="hover:opacity-100 absolute -translate-x-2/4 transition duration-[0.4s] ease-in-out h-auto w-[290px] left-2/4 top-0 bottom-0 opacity-0 bg-white p-2 m-auto rounded-xl flex flex-col justify-center sm:w-3/5 lg:w-full 2xl:p-3"
            :class="activeIndex == virtue.id ? 'opacity-100' : ''"
          >
            <div class="w-full secondary-color mobile-header-3-medium">
              {{ virtue.title }}
            </div>
            <div
              class="sub-h3-regular lg:sub-h1-regular text-black-60 mt-4 mt-2.5 mb-3 w-full lg:mb-0 2xl:mb-2"
            >
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-wrap m-auto lg:w-3/5">
        <div
          class="relative basis-full lg:basis-1/2"
          v-for="virtue in virtues.slice(3, 5)"
          :key="virtue.id"
          @touchstart.passive="(activeIndex = virtue.id)"
          @hover.native="$mixpanel.track('hover_virtue')"
        >
          <div
            class="flex flex-col items-center my-5 mx-auto py-0 px-10 sm:my-2 sm:px-0 lg:my-10"
          >
            <div
              class="rounded-full bg-white relative h-0 w-1/2 pb-[50%] mb-5 sm:w-[30%] sm:pb-[30%] lg:w-1/2 lg:pb-[50%]"
            >
              <img
                class="absolute w-[95%] h-[95%] left-0"
                :src="virtue.image"
                :alt="virtue.title"
                loading="lazy"
              />
            </div>
            <div
              class="w-3/5 sub-h1-semibold secondary-color lg:mobile-header-3-medium lg:text-black-87"
            >
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="hover:opacity-100 absolute -translate-x-2/4 transition duration-[0.4s] ease-in-out h-auto w-[290px] left-2/4 top-0 bottom-0 opacity-0 bg-white p-2 m-auto rounded-xl flex flex-col justify-center sm:w-3/5 lg:w-full 2xl:p-3"
            :class="activeIndex == virtue.id ? 'opacity-100' : ''"
          >
            <div class="w-full secondary-color mobile-header-3-medium">
              {{ virtue.title }}
            </div>
            <div
              class="sub-h3-regular lg:sub-h1-regular text-black-60 mt-4 mb-3 w-full lg:mb-0 2xl:mb-2"
            >
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-wrap m-auto lg:w-[85%]">
        <div
          class="relative basis-full lg:basis-[33%]"
          v-for="virtue in virtues.slice(5, 8)"
          :key="virtue.id"
          @touchstart.passive="(activeIndex = virtue.id)"
          @hover.native="$mixpanel.track('hover_virtue')"
        >
          <div
            class="flex flex-col items-center my-5 mx-auto py-0 px-10 sm:my-2 sm:px-0 lg:my-10"
          >
            <div
              class="rounded-full bg-white relative h-0 w-1/2 pb-[50%] mb-5 sm:w-[30%] sm:pb-[30%] lg:w-1/2 lg:pb-[50%]"
            >
              <img
                class="absolute w-[95%] h-[95%] left-0"
                :src="virtue.image"
                :alt="virtue.title"
                loading="lazy"
              />
            </div>
            <div
              class="w-3/5 sub-h1-semibold secondary-color lg:mobile-header-3-medium lg:text-black-87"
            >
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="hover:opacity-100 absolute -translate-x-2/4 transition duration-[0.4s] ease-in-out h-auto w-[290px] left-2/4 top-0 bottom-0 opacity-0 bg-white p-2 m-auto rounded-xl flex flex-col justify-center sm:w-3/5 lg:w-full 2xl:p-3"
            :class="activeIndex == virtue.id ? 'opacity-100' : ''"
          >
            <div class="w-full secondary-color mobile-header-3-medium">
              {{ virtue.title }}
            </div>
            <div
              class="sub-h3-regular lg:sub-h1-regular text-black-60 mt-4 mt-2.5 mb-3 w-full lg:mb-0 2xl:mb-2"
            >
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-wrap m-auto lg:w-3/5">
        <div
          class="relative basis-full lg:basis-1/2"
          v-for="virtue in virtues.slice(8, 10)"
          :key="virtue.id"
          @touchstart.passive="(activeIndex = virtue.id)"
          @hover.native="$mixpanel.track('hover_virtue')"
        >
          <div
            class="flex flex-col items-center my-5 mx-auto py-0 px-10 sm:my-2 sm:px-0 lg:my-10"
          >
            <div
              class="rounded-full bg-white relative h-0 w-1/2 pb-[50%] mb-5 sm:w-[30%] sm:pb-[30%] lg:w-1/2 lg:pb-[50%]"
            >
              <img
                class="absolute w-[95%] h-[95%] left-0"
                :src="virtue.image"
                :alt="virtue.title"
                loading="lazy"
              />
            </div>
            <div
              class="w-3/5 sub-h1-semibold secondary-color lg:mobile-header-3-medium lg:text-black-87"
            >
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="hover:opacity-100 absolute -translate-x-2/4 transition duration-[0.4s] ease-in-out h-auto w-[290px] left-2/4 top-0 bottom-0 opacity-0 bg-white p-2 m-auto rounded-xl flex flex-col justify-center sm:w-3/5 lg:w-full 2xl:p-3"
            :class="activeIndex == virtue.id ? 'opacity-100' : ''"
          >
            <div class="w-full secondary-color mobile-header-3-medium">
              {{ virtue.title }}
            </div>
            <div
              class="sub-h3-regular lg:sub-h1-regular text-black-60 mt-4 mt-2.5 mb-3 w-full md:mb-0 2xl:mb-2"
            >
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
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
