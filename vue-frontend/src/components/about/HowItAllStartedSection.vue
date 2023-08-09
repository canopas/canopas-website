<template>
  <div class="!tw-mt-20 xl:!tw-mt-[14.063rem] xl:tw-mb-20 2xl:tw-mb-48">
    <h1
      class="tw-text-center tw-text-[1.875rem] md:tw-text-[2.813rem] tw-leading-[2.25rem] md:tw-leading-[3.875rem] tw-font-inter-bold tw-text-black-core/[0.87]"
    >
      How it all started?
    </h1>
    <div
      :class="
        width > 1399 && width < 1536 ? '2xl:!tw-w-full' : '2xl:!tw-w-[1535px]'
      "
      class="tw-flex tw-flex-row tw-items-center tw-mx-auto tw-w-full xll:tw-w-[1530px] 3xl:tw-w-[1595px] xl:tw-pt-20"
    >
      <div
        class="tw-flex tw-flex-col tw-w-[60%] xll:tw-w-[70%] tw-overflow-hidden tw-py-40"
      >
        <div class="tw-relative">
          <ul
            ref="years"
            :style="{
              left: `${left}px`,
            }"
            class="tw-absolute tw-flex tw-flex-nowrap tw-justify-between tw-w-full tw-ml-[16%] 3xl:tw-ml-[14%] tw-w-[84%] 2xl:tw-w-[75%] xll:tw-w-[72%] tw-transition-all tw-ease-in-out tw-duration-1000"
          >
            <li
              v-for="(story, index) in stories"
              :key="story"
              :ref="'year-' + index"
              :class="
                index == activeIndex
                  ? 'v2-canopas-gradient-text'
                  : index < activeIndex
                  ? 'tw-text-black-core/[0.2]'
                  : 'tw-text-black-core/[0.6]'
              "
              class="tw-text-[1.5rem] tw-leading-[2.5rem] xl:tw-text-[1.875rem] xl:tw-leading-[2.812rem] tw-font-inter-medium hover:tw-bg-clip-text hover:tw-text-transparent hover:tw-from-[#FF835B] hover:tw-to-[#F2709C] hover:tw-bg-gradient-[270.11deg]"
              @click="slide(index, story)"
            >
              <span>{{ story.year }}</span>
            </li>
          </ul>
        </div>
        <div
          v-if="activeStory !== null"
          class="tw-ml-[16%] 3xl:tw-ml-[14%] tw-mt-[4.5rem] tw-w-[84%] 2xl:tw-w-[76%] xll:tw-w-[67%] tw-h-[200px]"
          :key="activeStory.year"
          :class="animate ? 'tw-animate-fadeIn ' : ''"
        >
          <p
            class="tw-text-[1rem] tw-leading-[1.5rem] xl:tw-text-[1.187rem] xl:tw-leading-[1.813rem] tw-font-inter-medium"
          >
            {{ activeStory.story }}
          </p>
        </div>
      </div>
      <div
        class="tw-relative tw-flex tw-flex-col tw-items-center tw-justify-center tw-mt-[4.25rem] 2xl:tw-mt-[14.25rem] xll:tw-mt-[21.25rem] 3xl:tw-mt-[26.25rem] tw-w-[40%] xll:tw-w-[30%] tw-pr-[5%] xll:tw-pr-0"
      >
        <img
          :src="image1"
          alt="workspace-image-1"
          class="tw-absolute tw-bottom-[8.5rem] xl:tw-bottom-[9.5rem] 2xl:tw-bottom-[12.5rem] xll:tw-bottom-[12.5rem] tw-left-[1.5rem] xll:tw-left-[-2.5rem] tw-w-[60%] xll:tw-w-[80%] tw-z-[3]"
        />
        <img
          :src="image2"
          alt="workspace-image-2"
          loading="lazy"
          class="tw-relative tw-self-end tw-w-[70%] xll:tw-w-[80%]"
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
      this.mixpanel.track(story.event);
    },
  },
  inject: ["mixpanel"],
  mounted() {
    this.activeStory = this.stories[0];
    this.width = window.innerWidth;
  },
};
</script>
