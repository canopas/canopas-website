<template>
  <div class="container relative m-0 mt-16 max-w-full p-0 text-center">
    <div>
      <p class="text-center background-text text-[#F8F8F8]">Services</p>
      <p
        class="-mt-5 pb-8 text-center mobile-header-2 lg:desk-header-2 text-black-87"
      >
        How we can help you
      </p>
    </div>
    <div></div>
    <div
      v-for="(service, index) in services"
      :key="index"
      class="relative flex px-0 py-5 even:flex-row-reverse"
    >
      <div
        :class="
          index % 2 == 0 ? 'items-end text-right' : 'items-start text-left'
        "
        class="mx-4 flex flex-[70%] flex-col"
      >
        <img
          :src="service.bullet"
          alt="service-bullet"
          loading="lazy"
          class="-mr-5 h-[50px] w-[50px]"
          height="50"
          width="50"
        />
        <nuxt-link
          :href="service.url"
          class="mobile-header-3 v2-canopas-gradient-text"
        >
          {{ service.title }}
        </nuxt-link>
        <div
          class="pt-4 sub-h3-regular text-black-60 sm:w-[70%] md:w-[65%]"
          v-html="service.content"
        ></div>
        <div
          class="absolute bottom-10 h-3 w-3 bg-[#F98386] rounded-full"
          :class="index % 2 == 0 ? 'left-5' : 'right-5'"
        ></div>
        <div
          class="absolute top-10 h-1.5 w-1.5 bg-[#F37299] rounded-full"
          :class="index % 2 == 0 ? 'left-40' : 'right-40'"
        ></div>
      </div>
      <img
        :src="service.image"
        alt="service-background-image"
        loading="lazy"
        class="h-full w-[126px] flex-[30%] sm:flex-[10%] md:flex-[6%]"
        height="180"
        width="130"
      />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      services: [
        {
          title: "Mobile Development",
          content:
            "Whether your goal is to acquire new users, improve your retention rate, or increase in-app sales, our team will guide you based on our decade-long experience. <span class='sub-h3-semibold'>From native to hybrid, we've got you covered.</span>",
          image: "/images/why/why-canopas-mobile-1.webp",
          bullet: "/images/why/why-canopas-dot-1.svg",
          url: "/mobile-app-development",
        },
        {
          title: "Web Development",
          content:
            "Whether you need help with websites, web apps, an admin panel, an analytical dashboard, or desktop apps, our engineers will develop <span class='sub-h3-semibold'>scalable backend and front-end solutions.</span>",
          image: "/images/why/why-canopas-mobile-2.webp",
          bullet: "/images/why/why-canopas-dot-2.svg",
        },
        {
          title: "Custom Software Development",
          content:
            "<span class='sub-h3-semibold'>Unlock your business growth with tailored solutions</span> for unique business needs, enabling scalability, seamless integration, competitive edge, long-term cost efficiency, and enhanced security.",
          image: "/images/why/why-canopas-mobile-3.webp",
          bullet: "/images/why/why-canopas-dot-3.svg",
        },
        {
          title: "UI/UX Design",
          content:
            "We design user-centered UI/UX followed by user testing. <span class='sub-h3-semibold'>We make design decisions based on only one thing: Data.</span> Because it ensures our clients get value from our work.",
          image: "/images/why/why-canopas-mobile-4.webp",
          bullet: "/images/why/why-canopas-dot-1.svg",
        },
        {
          title: "Bug Fixing and Maintenance",
          content:
            "Looking to fix bugs in your app? We've helped multiple clients fix bugs and later added new features to match their vision. Plus, <span class='sub-h3-semibold'>We offer 6 months of FREE maintenance on our work.</span>",
          image: "/images/why/why-canopas-mobile-5.webp",
          bullet: "/images/why/why-canopas-dot-1.svg",
        },
      ],
    };
  },
};
</script>
