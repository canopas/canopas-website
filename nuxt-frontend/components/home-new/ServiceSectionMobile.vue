<template>
  <div class="container relative m-0 mt-10 max-w-full p-0 text-center">
    <div>
      <p
        class="text-center text-[2.5rem] leading-[1.875rem] text-black-core/[.03]"
      >
        Services
      </p>
      <p
        class="-mt-5 pb-8 text-center font-roboto-bold text-[1.875rem] leading-[2.198rem]"
      >
        How We Can Help You
      </p>
    </div>
    <img
      :src="ServiceBackgroundImage"
      alt="service-background-image"
      loading="lazy"
      class="absolute left-0 top-0 h-full w-full object-contain sm:object-cover"
    />
    <div
      v-for="(service, index) in services"
      :key="index"
      class="relative flex px-0 py-5 even:flex-row-reverse"
    >
      <div
        :class="
          index % 2 == 0 ? 'items-end text-right' : 'items-start text-left'
        "
        class="mx-5 flex flex-[70%] flex-col text-black-900"
      >
        <img
          :src="service.bullet"
          alt="service-bullet"
          loading="lazy"
          class="-mr-5 h-[50px] w-[50px]"
        />
        <a
          :href="service.url"
          class="py-[15px] font-inter-bold text-[1.5rem] leading-[1.816rem]"
        >
          {{ service.title }}
        </a>
        <div
          class="py-[15px] font-inter-regular text-[1rem] leading-5 sm:w-[70%]"
        >
          {{ service.content }}
        </div>
      </div>
      <img
        :src="service.image"
        alt="service-background-image"
        loading="lazy"
        class="h-full w-[126px] flex-[30%]"
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
