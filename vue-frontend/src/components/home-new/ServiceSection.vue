<template>
  <div class="md:tw-pt-[100px] tw-font-inter-regular">
    <div
      class="tw-pb-16 md:tw-pb-24 tw-font-roboto-bold tw-text-black-core/[.87] tw-leading-[4.4375rem] tw-text-[3.438rem] tw-text-center"
    >
      How We Can Help You
    </div>

    <div
      v-for="(service, index) in services"
      :key="service"
      @mouseover="activeIndex = index"
      @mouseleave="activeIndex = null"
      @touchstart.passive="activeIndex = index"
      :class="
        activeIndex === index ? 'tw-ease-in-out hover:tw-text-transparent ' : ''
      "
      class="tw-relative tw-flex tw-items-center tw-h-36 tw-pl-16 xl:tw-pl-48 xll:tw-pl-[40rem]"
    >
      <div
        v-if="activeIndex === index || (activeIndex == null && index == 0)"
        class="tw-absolute tw-left-0 tw-top-0 tw-flex tw-justify-center tw-items-center tw-ml-0 tw-w-full tw-h-full tw-z-[2] tw-pr-8 xl:tw-pr-12 tw-shadow-[0px_2px_50px_4px_rgba(61,61,61,0.1)]"
      >
        <div
          class="tw-flex tw-items-center tw-basis-[44%] lg:tw-basis-[49%] tw-h-36 tw-bg-gradient-to-r tw-from-[#F2709C] tw-to-[#FF835B] tw-pl-16 xl:tw-pl-48 xll:tw-pl-[40rem] tw-font-inter-semibold tw-text-white tw-text-2xl xl:tw-text-[1.875rem] xl:tw-leading-[2.4375rem]"
        >
          0{{ index + 1 }}
          <span class="tw-pl-[20px] xl:tw-pl-[62px]">
            {{ service.title }}
          </span>
        </div>
        <div
          class="tw-basis-[1%] tw-ml-[-1px] tw-border-transparent tw-border-r-[70px] tw-border-t-[144px] tw-border-t-[#ff835b]"
        ></div>
        <div
          class="tw-flex tw-flex-col tw-basis-[55%] lg:tw-basis-[50%] tw-justify-center tw-items-start tw-px-4 tw-font-inter-regular tw-text-base xl:tw-text-[1.25rem] xl:tw-leading-[1.775rem] tw-text-black-core/[.8]"
        >
          <p>
            {{ service.content }}
          </p>
        </div>
      </div>
      <div
        v-else
        class="tw-ml-0 tw-text-xl xl:tw-text-[1.5625rem] xl:tw-leading-[2rem]"
      >
        0{{ index + 1 }}
        <span class="tw-pl-[25px] xl:tw-pl-[70px]">
          {{ service.title }}
        </span>
      </div>
    </div>
  </div>
</template>

<script type="module">
export default {
  data() {
    return {
      activeIndex: null,
      services: [
        {
          title: "Mobile app Development",
          content:
            "Whether your goal is to acquire new users, improve your retention rate, or increase in-app sales, our team will guide you based on our decade-long experience. From native to hybrid, we've got you covered.",
        },
        {
          title: "Web Development",
          content:
            "Whether you need help with websites, web apps, an admin panel, an analytical dashboard, or desktop apps, our engineers will develop scalable backend and front-end solutions.",
        },
        {
          title: "UI/UX Design ",
          content:
            "We design user-centered UI/UX followed by user testing. We make design decisions based on only one thing: Data. Because it ensures our clients get value from our work.",
        },
        {
          title: "SEO",
          content:
            "To scale your business exponentially you have to get organic reach for your businesses and we know how to do it. In just one year, we grew from 16 daily impressions to 24369 daily impressions. 1500X impressions growth under one year",
        },
        {
          title: "Software Testing",
          content:
            "Prevention is better than cure. Our development team lives by this motto. We write unit tests from the very early stages of development to ensure future maintainability, extendibility, and rapid development.",
        },
      ],
    };
  },
};
</script>
