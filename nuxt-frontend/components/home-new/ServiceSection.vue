<template>
  <div class="font-inter-regular md:mt-[7.813rem]">
    <p
      class="text-center text-[3.75rem] leading-[4.0625rem] text-black-core/[.03] xl:text-[5rem] xl:leading-[5rem]"
    >
      Services
    </p>
    <p
      class="-mt-10 pb-16 text-center font-roboto-bold text-[3.438rem] leading-[4.4375rem] text-black-core/[.87] md:pb-24 xl:mt-[-50px]"
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
          ? 'transform transition-all duration-500  ease-in-out hover:text-transparent'
          : '',
      ]"
      class="cursor-pointer relative flex h-36 items-center pl-16 xl:pl-48 xll:pl-[39rem] 3xl:pl-[84rem]"
    >
      <div
        v-if="activeIndex === index || (activeIndex == null && index == 0)"
        class="absolute left-0 top-0 z-[2] ml-0 flex h-full w-full items-center justify-center pr-8 shadow-[0px_2px_50px_4px_rgba(61,61,61,0.1)] xl:pr-12"
      >
        <div
          class="flex h-36 basis-[44%] items-center bg-gradient-to-r from-[#F2709C] to-[#FF835B] pl-16 font-inter-semibold text-2xl text-white lg:basis-[49%] xl:pl-48 xl:text-[1.875rem] xl:leading-[2.4375rem] xll:pl-[39rem] 3xl:pl-[84rem]"
        >
          0{{ index + 1 }}
          <span class="pl-5 xl:pl-[62px]">
            {{ service.title }}
          </span>
        </div>
        <div
          class="-ml-px basis-[1%] border-r-[70px] border-t-[144px] border-transparent border-t-[#ff835b]"
        ></div>
        <div
          class="flex basis-[55%] flex-col items-start justify-center px-4 xll:pr-[29rem] 3xl:pr-[78rem] font-inter-regular text-base text-black-core/[.8] lg:basis-[50%] xl:text-[1.25rem] xl:leading-[1.775rem]"
        >
          <p>
            {{ service.content }}
          </p>
        </div>
      </div>
      <div v-else class="ml-0 text-xl xl:text-[1.5625rem] xl:leading-8">
        0{{ index + 1 }}
        <span class="pl-[25px] xl:pl-[70px]">
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
