<template>
  <div class="!mt-20 xl:!mt-[14.063rem] xl:mb-20 2xl:mb-48">
    <h1
      class="text-center text-[1.875rem] md:text-[2.813rem] leading-[2.25rem] md:leading-[3.875rem] font-inter-bold text-black-core/[0.87]"
    >
      How it all started?
    </h1>
    <div
      :class="width > 1399 && width < 1536 ? '2xl:!w-full' : '2xl:!w-[1535px]'"
      class="flex flex-row items-center mx-auto w-full xll:w-[1530px] 3xl:w-[1595px] xl:pt-20"
    >
      <div class="flex flex-col w-[60%] xll:w-[70%] overflow-hidden py-40">
        <div class="relative">
          <ul
            ref="years"
            :style="{
              left: `${left}px`,
            }"
            class="absolute flex flex-nowrap justify-between w-full ml-[16%] 3xl:ml-[14%] w-[84%] 2xl:w-[75%] xll:w-[72%] transition-all ease-in-out duration-1000"
          >
            <li
              v-for="(story, index) in stories"
              :key="story"
              :ref="'year-' + index"
              :class="
                index == activeIndex
                  ? 'v2-canopas-gradient-text'
                  : index < activeIndex
                    ? 'text-black-core/[0.2]'
                    : 'text-black-core/[0.6]'
              "
              class="text-[1.5rem] leading-[2.5rem] xl:text-[1.875rem] xl:leading-[2.812rem] font-inter-medium hover:bg-clip-text hover:text-transparent hover:from-[#FF835B] hover:to-[#F2709C] hover:bg-gradient-[270.11deg]"
              @click="slide(index, story)"
            >
              <span>{{ story.year }}</span>
            </li>
          </ul>
        </div>
        <div
          v-if="activeStory !== null"
          class="ml-[16%] 3xl:ml-[14%] mt-[4.5rem] w-[84%] 2xl:w-[76%] xll:w-[67%] h-[200px]"
          :key="activeStory.year"
          :class="animate ? 'animate-fadeIn ' : ''"
        >
          <p
            class="text-[1rem] leading-[1.5rem] xl:text-[1.187rem] xl:leading-[1.813rem] font-inter-medium"
          >
            {{ activeStory.story }}
          </p>
        </div>
      </div>
      <div
        class="relative flex flex-col items-center justify-center mt-[4.25rem] 2xl:mt-[14.25rem] xll:mt-[21.25rem] 3xl:mt-[26.25rem] w-[40%] xll:w-[30%] pr-[5%] xll:pr-0"
      >
        <img
          :src="image1"
          alt="workspace-image-1"
          loading="lazy"
          class="absolute bottom-[8.5rem] xl:bottom-[9.5rem] 2xl:bottom-[12.5rem] xll:bottom-[12.5rem] left-[1.5rem] xll:left-[-2.5rem] w-[60%] xll:w-[80%] z-[3]"
        />
        <img
          :src="image2"
          alt="workspace-image-2"
          loading="lazy"
          class="relative self-end w-[70%] xll:w-[80%]"
        />
      </div>
    </div>
  </div>
</template>

<script type="module">
import image1 from "@/assets/images/about/workspace-image-1.webp";
import image2 from "@/assets/images/about/workspace-image-2.webp";

export default {
  data() {
    return {
      image1,
      image2,
      activeStory: null,
      activeIndex: 0,
      left: 0,
      animate: false,
      width: 1400,
      stories: [
        {
          year: "2012",
          story:
            "Once upon a time, in the bustling streets of India, two childhood friends Darpan and Jignesh had a dream of using technology to make a difference in the world. As 90s kids, they watched in awe as technology advanced at lightning speed, and they knew that they had to be a part of this revolution.",
          event: "tap_about_timeline_2012",
        },
        {
          year: "2013",
          story:
            "They started their journey in college, experimenting with every aspect of Android development. They tinkered with different operating systems like CyanogenMod and MIUI, build custom ROMs, and even installed them on their devices just for fun. These experiments taught them the ins and outs of the Android OS and its root structure.",
          event: "tap_about_timeline_2013",
        },
        {
          year: "2014",
          story:
            "Their first product was a lightweight all-in-one file manager, created to address the biggest concern of the time: storage and memory. It was the start of a professional journey that would take them all over the world, working with clients from the USA, Canada, UK, Asia, Africa, and many more.",
          event: "tap_about_timeline_2014",
        },
        {
          year: "2016",
          story:
            "As they gained more experience, they expanded their services to include iOS and backend development. Today, they help businesses generate more revenue, acquire more users, and deliver exceptional user experiences through mobile apps, websites, web apps, desktop apps, UI/UX, and SEO.",
          event: "tap_about_timeline_2016",
        },
        {
          year: "2020",
          story:
            "At Canopas, they see themselves as problem solvers first and foremost. Technology is just the tool they use to make a difference in the world. Whether they're working on a new project or tackling an old problem, you can always count on them to deliver the best solution possible.",
          event: "tap_about_timeline_2020",
        },
        {
          year: "2023",
          story:
            "If you're a business owner or an entrepreneur with a vision to help humanity, we would love to hear from you. With Canopas, you can be sure that your ideas will be transformed into reality and make a real impact in the world.",
          event: "tap_about_timeline_2023",
        },
      ],
    };
  },
  methods: {
    slide(index, story) {
      this.animate = true;
      // ul width
      let totalWidth = Math.ceil(this.$refs.years.clientWidth);

      // li width
      let listWidth = Math.ceil(this.$refs["year-" + index][0].clientWidth);

      // space between list elements
      let space = Math.ceil(
        (totalWidth - listWidth * this.stories.length) /
          (this.stories.length - 1),
      );

      // move left with total of li width and space
      // multiplied with element index to move accurately
      this.left -= (listWidth + space) * (index - this.activeIndex);

      this.activeIndex = index;
      this.activeStory = story;
      this.$mixpanel.track(story.event);
    },
  },
  inject: ["mixpanel"],
  mounted() {
    this.activeStory = this.stories[0];
    this.width = window.innerWidth;
  },
};
</script>
