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
        class="mx-4 flex flex-[70%] flex-col"
      >
        <img
          :src="service.bullet"
          alt="service-bullet"
          loading="lazy"
          class="-mr-5 h-[50px] w-[50px]"
        />
        <a :href="service.url" class="mobile-header-3 v2-canopas-gradient-text">
          {{ service.title }}
        </a>
        <div class="pt-4 sub-h3-regular text-black-60 sm:w-[70%] md:w-[65%]">
          {{ service.content }}
        </div>
      </div>
      <img
        :src="service.image"
        alt="service-background-image"
        loading="lazy"
        class="h-full w-[126px] flex-[30%] sm:flex-[10%] md:flex-[6%]"
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
