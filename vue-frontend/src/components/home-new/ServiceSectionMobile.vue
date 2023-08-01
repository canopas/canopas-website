<template>
  <div
    class="tw-container tw-relative tw-m-0 tw-mt-[40px] tw-max-w-full tw-p-0 tw-text-center"
  >
    <div>
      <p
        class="tw-text-center tw-text-[2.5rem] tw-leading-[1.875rem] tw-text-black-core/[.03]"
      >
        Services
      </p>
      <p
        class="tw-mt-[-20px] tw-pb-8 tw-text-center tw-font-roboto-bold tw-text-[1.875rem] tw-leading-[2.198rem]"
      >
        How We Can Help You
      </p>
    </div>
    <img
      :src="ServiceBackgroundImage"
      alt="service-background-image"
      loading="lazy"
      class="tw-absolute tw-left-0 tw-top-0 tw-h-full tw-w-full tw-object-contain sm:tw-object-cover"
    />
    <div
      v-for="(service, index) in services"
      :key="index"
      class="tw-relative tw-flex tw-px-0 tw-py-[20px] even:tw-flex-row-reverse"
    >
      <div
        :class="
          index % 2 == 0
            ? 'tw-items-end tw-text-right'
            : 'tw-items-start tw-text-left'
        "
        class="tw-mx-[20px] tw-flex tw-flex-[70%] tw-flex-col tw-text-black-900"
      >
        <img
          :src="service.bullet"
          alt="service-bullet"
          loading="lazy"
          class="-tw-mr-[20px] tw-h-[50px] tw-w-[50px]"
        />
        <a
          :href="service.url"
          class="tw-py-[15px] tw-font-inter-bold tw-text-[1.5rem] tw-leading-[1.816rem]"
        >
          {{ service.title }}
        </a>
        <div
          class="tw-py-[15px] tw-font-inter-regular tw-text-[1rem] tw-leading-[1.25rem] sm:tw-w-[70%]"
        >
          {{ service.content }}
        </div>
      </div>
      <img
        :src="service.image"
        alt="service-background-image"
        loading="lazy"
        class="tw-h-full tw-w-[126px] tw-flex-[30%]"
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
import config from "@/config.js";

export default {
  data() {
    return {
      ServiceBackgroundImage,
      services: [
        {
          title: "Mobile App Development",
          content:
            "Whether your goal is to acquire new users, improve your retention rate, or increase in-app sales, our team will guide you based on our decade-long experience. From native to hybrid, we've got you covered.",
          image: circle1,
          bullet: bullet1,
          url: "/mobile-app-development",
        },
        {
          title: "Web Development",
          content:
            "Whether you need help with websites, web apps, an admin panel, an analytical dashboard, or desktop apps, our engineers will develop scalable backend and front-end solutions.",
          image: circle2,
          bullet: bullet2,
        },
        {
          title: "Custom Software Development",
          content:
            "Unlock your business growth with tailored solutions for unique business needs, enabling scalability, seamless integration, competitive edge, long-term cost efficiency, and enhanced security.",
          image: circle3,
          bullet: bullet3,
        },
        {
          title: "UI/UX Design",
          content:
            "We design user-centered UI/UX followed by user testing. We make design decisions based on only one thing: Data. Because it ensures our clients get value from our work.",
          image: circle4,
          bullet: bullet1,
        },
        {
          title: "DevOps and Maintenance",
          content:
            "Our team values prevention over cure, writes unit tests from an early stage of the project for future maintainability, extendibility, and quick development, and uses auto-deployment to save time.",
          image: circle5,
          bullet: bullet1,
        },
      ],
    };
  },
};
</script>
