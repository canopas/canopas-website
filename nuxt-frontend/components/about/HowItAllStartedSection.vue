<template>
  <section class="bg-deep-charcoal 3xl:outer-container">
    <div
      class="py-[7.5rem] relative container my-0 mx-auto flex gap-[6.5rem] mb-60"
    >
      <div class="sticky w-full top-[25rem] pb-[12.5rem] h-[3vh]">
        <div class="flex flex-col float-left">
          <div class="desk-header-2 text-white">How it all started?</div>
          <nuxt-link
            class="ml-0 gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_landing_cta')"
          >
            <span class="sub-h3-semibold lg:sub-h1-semibold"
              >Talk to Our Expert
            </span>
          </nuxt-link>
        </div>
      </div>
      <div class="flex flex-col gap-14">
        <div
          class="flex gap-10 item opacity-50"
          v-for="(event, index) in events"
          :key="index"
          :class="{ '!opacity-100': index === activeIndex }"
        >
          <div><img :src="event.image" alt="year-image" loading="lazy" /></div>
          <div class="w-4/5 mobile-header-3-medium text-white">
            {{ event.description }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup type="module">
import year1 from "@/assets/images/about/year1.webp";
import year2 from "@/assets/images/about/year2.webp";
import year3 from "@/assets/images/about/year3.webp";
import year4 from "@/assets/images/about/year4.webp";
import year5 from "@/assets/images/about/year5.webp";
import year6 from "@/assets/images/about/year6.webp";

var activeIndex = ref(0);
const events = [
  {
    image: year1,
    description:
      "Once upon a time, in the bustling streets of India, two childhood friends Darpan and Jignesh had a dream of using technology to make a difference in the world. As 90s kids, they watched in awe as technology advanced at lightning speed, and they knew that they had to be a part of this revolution.",
  },
  {
    image: year2,
    description:
      "They started their journey in college, experimenting with every aspect of Android development. They tinkered with different operating systems like CyanogenMod and MIUI, build custom ROMs, and even installed them on their devices just for fun. These experiments taught them the ins and outs of the Android OS and its root structure.",
  },
  {
    image: year3,
    description:
      "Their first product was a lightweight all-in-one file manager, created to address the biggest concern of the time: storage and memory. It was the start of a professional journey that would take them all over the world, working with clients from the USA, Canada, UK, Asia, Africa, and many more.",
  },
  {
    image: year4,
    description:
      "As they gained more experience, they expanded their services to include iOS and backend development. Today, they help businesses generate more revenue, acquire more users, and deliver exceptional user experiences through mobile apps, websites, web apps, desktop apps, UI/UX, and SEO.",
  },
  {
    image: year5,
    description:
      "At Canopas, they see themselves as problem solvers first and foremost. Technology is just the tool they use to make a difference in the world. Whether they're working on a new project or tackling an old problem, you can always count on them to deliver the best solution possible.",
  },
  {
    image: year6,
    description:
      "If you're a business owner or an entrepreneur with a vision to help humanity, we would love to hear from you. With Canopas, you can be sure that your ideas will be transformed into reality and make a real impact in the world.",
  },
];

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
function handleScroll() {
  const child = document.querySelectorAll(".item");
  child.forEach((item, index) => {
    const itemRect = item.getBoundingClientRect();
    const itemTop = itemRect.top;
    if (itemTop < 380) {
      activeIndex.value = index;
    }
  });
}
</script>
