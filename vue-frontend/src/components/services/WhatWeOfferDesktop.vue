<template>
  <section>
    <div
      class="tw-py-12 xl:tw-py-24 tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.25rem] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem]"
    >
      <span>What We Offer</span>
    </div>
    <div class="tw-container"></div>
    <div class="tw-container tw-text-center tw-text-black-core-900/[0.87]">
      <div class="tw-flex tw-justify-center tw-mb-16">
        <div class="tw-flex tw-flex-col tw-justify-end tw-w-[50%] tw-text-left">
          <div
            v-for="(service, index) in services"
            :key="service"
            class="tw-relative"
            :style="{
              height: `${height}px`,
            }"
            :ref="'image-' + index"
          >
            <div class="tw-absolute tw-top-[50%] tw-translate-y-[-50%]">
              <div class="tw-py-5">
                <span
                  class="tw-font-inter-semibold tw-text-[1.5rem] tw-leading-[2.25rem] md:tw-text-[2.25rem] md:tw-leading-[3.43rem] xl:tw-text-[3.125rem] xl:tw-leading-[4.6875rem]"
                  >{{ service.title }}</span
                >
              </div>
              <div
                class="tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.094rem] xl:tw-text-[1.188rem] tw-leading-[1.5rem] md:tw-leading-[1.625rem] xl:tw-leading-[1.781rem]"
              >
                <p>{{ service.description }}</p>
              </div>
            </div>
          </div>
        </div>
        <div
          class="tw-sticky tw-flex tw-items-center tw-top-0 -tw-translate-[50%] tw-w-[50%]"
          :style="{
            height: `${height}px`,
          }"
        >
          <div
            v-for="service in services"
            :key="service"
            class="tw-absolute tw-w-[100%] tw-overflow-hidden tw-pb-[100%]"
            :style="{
              clipPath: `polygon(0% ${service.path[0]}%, 100% ${service.path[1]}%, 100% 100%, 0% 100%)`,
            }"
          >
            <img
              :src="service.image"
              alt="Service-images"
              class="tw-absolute tw-top-0 tw-right-0 tw-w-[80%] tw-h-full tw-z-[1] tw-object-cover tw-ml-auto"
              loading="lazy"
            />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script type="module">
import mobileApp800w from "@/assets/images/service/mobileapp-800w.webp";
import seo800w from "@/assets/images/service/seo-800w.webp";
import softwaretesting800w from "@/assets/images/service/softwaretesting-800w.webp";
import uiuxdesign800w from "@/assets/images/service/uiuxdesign-800w.webp";
import webapp800w from "@/assets/images/service/webapp-800w.webp";
import { elementInViewPort } from "@/utils";

export default {
  data() {
    return {
      services: [
        {
          path: [0, 0],
          image: mobileApp800w,
          title: "Mobile App Development",
          description:
            "When you've spent over a decade building apps you learn a lot about the ins and outs of the mobile world. It's for that reason that at Canopas we've crafted specialized app solutions for our clients that help ensure \
             their mobile apps are next level when they launch, and stay that way as they grow. From native to hybrid, we've got you covered.",
        },
        {
          path: [100, 100],
          image: webapp800w,
          title: "Web App Development",
          description:
            "Web apps are complex, so you'll be wanting some proven experience in your corner. We've been specialized app developers for over a decade; you're in safe hands. Our team of web engineers will save your website from the evil forces of slow load times and outdated design.\
             So, if you're in need of a web superhero, call us today, and let the magic begin!",
        },
        {
          path: [100, 100],
          image: uiuxdesign800w,
          title: "UI/UX Design",
          description:
            "We help our clients see themselves through the eyes of the people whom they are trying to reach. Then we design something amazing.\
            Product design is about how it works, but we make sure it look nice too. Then, we'll design a mobile or web experience that feels great.",
        },
        {
          path: [100, 100],
          image: seo800w,
          title: "SEO / Digital Marketer",
          description:
            "To scale your business exponentially you have to get organic reach for your business and we know how to do it. In just one year, \
            we grew from 16 daily impressions to 24369 daily impressions. 1500X impressions growth under one year. We know how it works and what it takes.",
        },
        {
          path: [100, 100],
          image: softwaretesting800w,
          title: "Software Testing",
          description:
            "From unit testing to Test Driven Development, Manual Testing to Automation Testing, our team of QA specialists will work with engineers to make sure you get products built to last, \
            with the highest standards, so that you can rely on your solution for years to come.",
        },
      ],
      lastScrollY: 0,
      prevIndex: 0,
      height: 0,
    };
  },
  methods: {
    handleScroll() {
      const scrollY = window.scrollY;
      const scrollUp = this.lastScrollY > scrollY && window.pageYOffset > 100;

      const elementIdx = elementInViewPort(this.$refs);
      const index = parseInt(
        elementIdx ? elementIdx.charAt(elementIdx.length - 1) : 0
      );

      if (this.prevIndex != index) {
        // when image will change, prevImage should be at path 0 or 100
        this.services[this.prevIndex].path = [
          scrollUp ? 100 : 0,
          scrollUp ? 100 : 0,
        ];
        this.prevIndex = index;
      }

      // random offset used when scrolling
      const rect = this.$refs[elementIdx][0].getBoundingClientRect();
      const scrollOffset =
        ((window.innerHeight - rect.top) * 100) / window.innerHeight;

      // if firstImage then path should be 0
      if (index == 0) {
        this.services[index].path = [0, 0];
      } else if (
        !scrollUp &&
        index == this.services.length - 1 &&
        rect.top < window.innerHeight / 2
      ) {
        this.services[index].path = [scrollUp ? 100 : 0, scrollUp ? 100 : 0];
      } else {
        this.services[index].path = [100 - scrollOffset, 100 - scrollOffset];
      }

      this.lastScrollY = scrollY;
    },
  },
  mounted() {
    this.height = window.innerHeight;
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
};
</script>
