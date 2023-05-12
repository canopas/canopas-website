<template>
  <div class="sm:tw-container tw-text-center tw-text-black-core-900/[0.87]">
    <div
      class="tw-py-12 sm:tw-py-16 tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.25rem]"
    >
      <span>What We Offer</span>
    </div>

    <div class="tw-flex tw-flex-col">
      <div
        v-for="service in services"
        :key="service"
        class="md:tw-w-[70%] tw-m-auto"
      >
        <div>
          <img
            :src="service.images[0]"
            :srcset="`${service.images[0]} 400w, ${service.images[1]} 800w`"
            alt="Service-images"
            class="tw-w-full tw-h-full tw-object-cover"
            loading="lazy"
          />
        </div>

        <div :class="service.class" class="tw-px-2.5 sm:tw-px-20">
          <div class="tw-py-5 sm:tw-py-6">
            <span
              class="tw-text-center tw-font-inter-semibold tw-text-[1.5rem] tw-leading-[2.25rem]"
              >{{ service.title }}</span
            >
          </div>
          <div
            class="tw-mb-10 sm:tw-mb-14 tw-text-center tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem]"
          >
            <p>{{ service.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import mobilApp400w from "@/assets/images/service/mobileapp-400w.webp";
import mobilApp800w from "@/assets/images/service/mobileapp-800w.webp";

import seo400w from "@/assets/images/service/seo-400w.webp";
import seo800w from "@/assets/images/service/seo-800w.webp";

import softwaretesting400w from "@/assets/images/service/softwaretesting-400w.webp";
import softwaretesting800w from "@/assets/images/service/softwaretesting-800w.webp";

import uiuxdesign400w from "@/assets/images/service/uiuxdesign-400w.webp";
import uiuxdesign800w from "@/assets/images/service/uiuxdesign-800w.webp";

import webapp400w from "@/assets/images/service/webapp-400w.webp";
import webapp800w from "@/assets/images/service/webapp-800w.webp";

export default {
  data() {
    return {
      services: [
        {
          images: [mobilApp400w, mobilApp800w],
          title: "Mobile App Development",
          description:
            "When you've spent over a decade building apps you learn a lot about the ins and outs of the mobile world. It's for that reason that at Canopas we've crafted specialized app solutions for our clients that help ensure \
             their mobile apps are next level when they launch, and stay that way as they grow. From native to hybrid, we've got you covered.",
          class:
            "tw-bg-gradient-to-b tw-from-[#FADFB7]  via-[#FADFB7]  tw-to-white tw-bg-gradient-[180deg] ",
        },
        {
          images: [webapp400w, webapp800w],
          title: "Web App Development",
          description:
            "Web apps are complex, so you'll be wanting some proven experience in your corner. We've been specialized app developers for over a decade, you're in safe hands. Our team of web engineers will save your website from the evil forces of slow load times and outdated design.\
             So, if you're in need of a web superhero, call us today, and let the magic begin!",
          class:
            "tw-bg-gradient-to-b tw-from-[#BBDEFF]  via-[#BBDEFF]  tw-to-white tw-bg-gradient-[180deg] ",
        },
        {
          images: [uiuxdesign400w, uiuxdesign800w],
          title: "UI/UX Design",
          description:
            "We help our clients see themselves through the eyes of the people whom they are trying to reach. Then we design something amazing.\
            Product design is about how it works, but we make sure it look nice too. Then, we'll design a mobile or web experience that feels great.",
          class:
            "tw-bg-gradient-to-b tw-from-[#E3D9EE]  via-[#E3D9EE]  tw-to-white tw-bg-gradient-[180deg] ",
        },
        {
          images: [seo400w, seo800w],
          title: "SEO / Digital Marketer",
          description:
            "To scale your business exponentially you have to get organic reach for your business and we know how to do it. In just one year, \
            we grew from 16 daily impressions to 24369 daily impressions. 1500X impressions growth under one year. We know how it works and what it takes.",
          class:
            "tw-bg-gradient-to-b tw-from-[#EBEDF3]  via-[#EBEDF3]  tw-to-white tw-bg-gradient-[180deg] ",
        },
        {
          images: [softwaretesting400w, softwaretesting800w],
          title: "Software Testing",
          description:
            "From unit testing to Test Driven Development, Manual Testing to Automation Testing, our team of QA specialists will work with engineers to make sure you get products built to last, \
            with the highest standards, so that you can rely on your solution for years to come.",
          class:
            "tw-bg-gradient-to-b tw-from-[#E8FAF0] via-[#E8FAF0] tw-to-white tw-bg-gradient-[180deg] ",
        },
      ],
    };
  },
};
</script>
