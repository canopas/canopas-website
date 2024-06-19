<template>
  <div class="mt-60">
    <p class="text-center background-text text-[#F8F8F8]">Services</p>
    <p
      class="-mt-10 pb-16 text-center mobile-header-2 lg:desk-header-2 text-black-87 md:pb-24 xl:mt-[-50px]"
    >
      How we can help you
    </p>

    <nuxt-link
      v-for="(service, index) in services"
      :key="service"
      @mouseover="activeIndex = index"
      @touchstart.passive="activeIndex = index"
      :href="service.url"
      :class="[
        activeIndex === index
          ? 'transform transition-all duration-500  ease-in-out hover:text-transparent '
          : '',
      ]"
      class="cursor-pointer relative flex h-36 items-center pl-16 2xl:pl-48 xll:pl-[39rem] 3xl:outer-container"
    >
      <div
        v-if="activeIndex === index || (activeIndex == null && index == 0)"
        class="absolute left-0 top-0 z-[2] ml-0 flex h-full w-full items-center justify-center pr-8 shadow-[0px_2px_50px_4px_rgba(61,61,61,0.1)] xl:pr-12"
      >
        <div
          class="flex sub-h1-regular lg:mobile-header-2-regular xl:desk-header-3 h-36 basis-[44%] items-center bg-gradient-to-r from-[#F2709C] to-[#FF835B] pl-16 text-white lg:basis-[49%] 2xl:pl-48 xll:pl-[39rem]"
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
          class="flex basis-[55%] flex-col items-start justify-center px-4 xll:pr-[29rem] sub-h3-regular xl:mobile-header-3-regular text-black-87 lg:basis-[50%]"
        >
          <p v-html="service.content"></p>
        </div>
      </div>
      <div
        v-else
        class="ml-0 sub-h1-regular lg:mobile-header-2-regular text-black-87"
      >
        0{{ index + 1 }}
        <span class="pl-[25px] xl:pl-[70px]">
          {{ service.title }}
        </span>
      </div>
    </nuxt-link>
  </div>
</template>

<script type="module">
export default {
  data() {
    return {
      activeIndex: null,
      services: [
        {
          title: "Mobile Development",
          content:
            "Whether your goal is to acquire new users, improve your retention rate, or increase in-app sales, our team will guide you based on our decade-long experience. <span class='font-bold'>From native to hybrid, we've got you covered.</span>",
          url: "/mobile-app-development",
        },
        {
          title: "Web Development",
          content:
            "Whether you need help with websites, web apps, an admin panel, an analytical dashboard, or desktop apps, our engineers will develop <span class='font-bold'>scalable backend and front-end solutions.</span>",
        },
        {
          title: "Custom Software Development",
          content:
            "<span class='font-bold'>Unlock your business growth with tailored solutions</span> for unique business needs, enabling scalability, seamless integration, competitive edge, long-term cost efficiency, and enhanced security.",
        },
        {
          title: "UI/UX Design",
          content:
            "We design user-centered UI/UX followed by user testing. <span class='font-bold'>We make design decisions based on only one thing: Data.</span> Because it ensures our clients get value from our work.",
        },
        {
          title: "Bug Fixing and Maintenance",
          content:
            "Looking to fix bugs in your app? We've helped multiple clients fix bugs and later added new features to match their vision. Plus, <span class='font-bold'>We offer 6 months of FREE maintenance on our work.</span>",
        },
      ],
    };
  },
};
</script>
