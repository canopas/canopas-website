<template>
  <div class="tw-font-inter-regular md:tw-mt-[7.813rem]">
    <p
      class="tw-text-center tw-text-[3.75rem] tw-leading-[4.0625rem] tw-text-black-core/[.03] xl:tw-text-[5rem] xl:tw-leading-[5rem]"
    >
      Services
    </p>
    <p
      class="-tw-mt-10 tw-pb-16 tw-text-center tw-font-roboto-bold tw-text-[3.438rem] tw-leading-[4.4375rem] tw-text-black-core/[.87] md:tw-pb-24 xl:tw-mt-[-50px]"
    >
      How We Can Help You
    </p>

    <a
      v-for="(service, index) in services"
      :key="service"
      @mouseover="activeIndex = index"
      @touchstart.passive="activeIndex = index"
      :href="service.url"
      :class="[
        activeIndex === index
          ? 'tw-transform tw-transition-all tw-duration-500  tw-ease-in-out hover:tw-text-transparent'
          : '',
      ]"
      class="tw-cursor-pointer tw-relative tw-flex tw-h-36 tw-items-center tw-pl-16 xl:tw-pl-48 xll:tw-pl-[39rem] 3xl:tw-pl-[84rem]"
    >
      <div
        v-if="activeIndex === index || (activeIndex == null && index == 0)"
        class="tw-absolute tw-left-0 tw-top-0 tw-z-[2] tw-ml-0 tw-flex tw-h-full tw-w-full tw-items-center tw-justify-center tw-pr-8 tw-shadow-[0px_2px_50px_4px_rgba(61,61,61,0.1)] xl:tw-pr-12"
      >
        <div
          class="tw-flex tw-h-36 tw-basis-[44%] tw-items-center tw-bg-gradient-to-r tw-from-[#F2709C] tw-to-[#FF835B] tw-pl-16 tw-font-inter-semibold tw-text-2xl tw-text-white lg:tw-basis-[49%] xl:tw-pl-48 xl:tw-text-[1.875rem] xl:tw-leading-[2.4375rem] xll:tw-pl-[39rem] 3xl:tw-pl-[84rem]"
        >
          0{{ index + 1 }}
          <span class="tw-pl-5 xl:tw-pl-[62px]">
            {{ service.title }}
          </span>
        </div>
        <div
          class="-tw-ml-px tw-basis-[1%] tw-border-r-[70px] tw-border-t-[144px] tw-border-transparent tw-border-t-[#ff835b]"
        ></div>
        <div
          class="tw-flex tw-basis-[55%] tw-flex-col tw-items-start tw-justify-center tw-px-4 xll:tw-pr-[29rem] 3xl:tw-pr-[78rem] tw-font-inter-regular tw-text-base tw-text-black-core/[.8] lg:tw-basis-[50%] xl:tw-text-[1.25rem] xl:tw-leading-[1.775rem]"
        >
          <p>
            {{ service.content }}
          </p>
        </div>
      </div>
      <div
        v-else
        class="tw-ml-0 tw-text-xl xl:tw-text-[1.5625rem] xl:tw-leading-8"
      >
        0{{ index + 1 }}
        <span class="tw-pl-[25px] xl:tw-pl-[70px]">
          {{ service.title }}
        </span>
      </div>
    </a>
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
          url: "/mobile-app-development",
        },
        {
          title: "Web Development",
          content:
            "Whether you need help with websites, web apps, an admin panel, an analytical dashboard, or desktop apps, our engineers will develop scalable backend and front-end solutions.",
        },
        {
          title: "Custom Software Development",
          content:
            "Unlock your business growth with tailored solutions for unique business needs, enabling scalability, seamless integration, competitive edge, long-term cost efficiency, and enhanced security.",
        },
        {
          title: "UI/UX Design",
          content:
            "We design user-centered UI/UX followed by user testing. We make design decisions based on only one thing: Data. Because it ensures our clients get value from our work.",
        },
        {
          title: "DevOps and Maintenance",
          content:
            "Our team values prevention over cure, writes unit tests from an early stage of the project for future maintainability, extendibility, and quick development, and uses auto-deployment to save time.",
        },
      ],
    };
  },
};
</script>
