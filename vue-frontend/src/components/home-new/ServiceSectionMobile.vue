<template>
  <div
    class="tw-container tw-relative tw-m-0 tw-mt-[66px] tw-max-w-full tw-p-0 tw-text-center"
  >
    <h1
      class="tw-pb-8 tw-font-roboto-bold tw-text-center tw-leading-[2.198rem] tw-text-[1.875rem]"
    >
      How We Can Help You
    </h1>
    <img
      :src="ServiceBackgroundImage"
      alt="service-background-image"
      loading="lazy"
      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-full tw-object-contain sm:tw-object-cover"
    />
    <div
      v-for="(service, index) in services"
      :key="index"
      class="tw-relative tw-flex even:tw-flex-row-reverse tw-py-[20px] tw-px-0"
    >
      <div
        :class="
          index % 2 == 0
            ? 'tw-items-end tw-text-right'
            : 'tw-items-start tw-text-left'
        "
        class="tw-flex tw-flex-col tw-flex-[70%] tw-mx-[20px] tw-text-black-900"
      >
        <img
          :src="service.bullet"
          alt="service-bullet"
          loading="lazy"
          class="tw-w-[50px] tw-h-[50px] -tw-mr-[20px]"
        />
        <div
          class="tw-py-[15px] tw-font-inter-bold tw-text-[1.5rem] tw-leading-[1.816rem]"
        >
          {{ service.title }}
        </div>
        <div
          class="sm:tw-w-[70%] tw-py-[15px] tw-font-inter-regular tw-text-[1rem] tw-leading-[1.25rem]"
        >
          {{ service.content }}
        </div>
      </div>
      <img
        :src="service.image"
        alt="service-background-image"
        loading="lazy"
        class="tw-flex-[30%] tw-w-[126px] tw-h-full"
      />
    </div>
  </div>
</template>

<script>
import ServiceBackgroundImage from "@/assets/images/why/why-canopas-mobile-bg.svg";

import bullet1 from "@/assets/images/why/why-canopas-dot-1.svg";
import bullet2 from "@/assets/images/why/why-canopas-dot-2.svg";
import bullet3 from "@/assets/images/why/why-canopas-dot-3.svg";

import circle1 from "@/assets/images/why/why-canopas-mobile-1.webp";
import circle2 from "@/assets/images/why/why-canopas-mobile-2.webp";
import circle3 from "@/assets/images/why/why-canopas-mobile-3.webp";
import circle4 from "@/assets/images/why/why-canopas-mobile-4.webp";
import circle5 from "@/assets/images/why/why-canopas-mobile-5.webp";

export default {
  data() {
    return {
      ServiceBackgroundImage,
      services: [
        {
          title: "Mobile Development",
          content:
            "Whether your goal is to acquire new users, improve your retention rate, or increase in-app sales, our team will guide you based on our decade-long experience. From native to hybrid, we've got you covered.",
          image: circle1,
          bullet: bullet1,
        },
        {
          title: "Web Development",
          content:
            "Whether you need help with websites, web apps, an admin panel, an analytical dashboard, or desktop apps, our engineers will develop scalable backend and front-end solutions.",
          image: circle2,
          bullet: bullet2,
        },
        {
          title: "UI/UX Design",
          content:
            "We design user-centered UI/UX followed by user testing. We make design decisions based on only one thing: Data. Because it ensures our clients get value from our work.",
          image: circle3,
          bullet: bullet3,
        },
        {
          title: "SEO",
          content:
            "To scale your business exponentially you have to get organic reach for your businesses and we know how to do it. In just one year, we grew from 16 daily impressions to 24369 daily impressions. 1500X impressions growth under one year.",
          image: circle4,
          bullet: bullet1,
        },
        {
          title: "Software Testing",
          content:
            "Prevention is better than cure. Our development team lives by this motto. We write unit tests from the very early stages of development to ensure future maintainability, extendibility, and rapid development.",
          image: circle5,
          bullet: bullet1,
        },
      ],
    };
  },
};
</script>
