<template>
  <section>
    <div class="tw-container tw-mt-10 tw-mb-20">
      <h1
        class="tw-pt-5 sm:tw-pt-10 tw-text-center tw-text-[1.875rem] md:tw-text-[2.813rem] tw-leading-[2.25rem] md:tw-leadig-[3.875rem] tw-font-inter-bold tw-text-black-core/[0.87]"
      >
        How it all started?
      </h1>
      <div
        class="tw-container tw-flex tw-flex-col tw-items-center tw-w-full md:tw-w-[82%] tw-h-[24.25rem] tw-mt-[5rem] tw-overflow-y-scroll hidden-scrollbar before:tw--mt-5 before:tw-blur-lg before:tw-absolute before:tw-left-0 before:tw-h-[6%] before:tw-w-full before:tw-bg-white-gradient"
      >
        <div
          v-for="(event, index) in events"
          :key="event"
          class="tw-flex tw-flex-row tw-items-center"
        >
          <div :class="[event.className]" class="tw-flex tw-flex-row">
            <div class="tw-flex tw-flex-col">
              <h1
                class="tw-text-[1.5rem] tw-leading-[2.25rem] tw-font-inter-semibold v2-canopas-gradient-text"
              >
                {{ event.year }}
              </h1>
              <div
                :class="index == events.length - 1 ? 'tw-hidden' : ''"
                class="tw-ml-[1.5rem] tw-h-60 sm:tw-h-40 md:tw-h-48 tw-border-dotted tw-border-l-2 tw-border-[#000000]/[0.87]"
              ></div>
            </div>
            <div
              class="tw-ml-[1.5rem] tw-text-[1rem] md:tw-text-[1.125rem] tw-leading-[1.5rem] md:tw-leading-[1.625rem] tw-font-inter-medium tw-text-black-core/[0.87]"
            >
              {{ event.description }}
            </div>
          </div>
        </div>
      </div>
      <div
        class="tw-mt-[-1.5rem] tw-w-full tw-h-12 after:tw-absolute after:tw-left-0 after:tw-h-[5%] after:tw-w-full after:tw-blur-lg after:tw-bg-white-gradient-bottom"
      ></div>
    </div>
  </section>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      events: [
        {
          year: "2012",
          description:
            "Once upon a time, in the bustling streets of India, two childhood friends Darpan and Jignesh had a dream of using technology to make a difference in the world. As 90s kids, they watched in awe as technology advanced at lightning speed, and they knew that they had to be a part of this revolution.",
          className: "tw-mt-9",
        },
        {
          year: "2013",
          description:
            "They started their journey in college, experimenting with every aspect of Android development. They tinkered with different operating systems like CyanogenMod and MIUI, build custom ROMs, and even installed them on their devices just for fun. These experiments taught them the ins and outs of the Android OS and its root structure.",
        },
        {
          year: "2014",
          description:
            "Their first product was a lightweight all-in-one file manager, created to address the biggest concern of the time: storage and memory. It was the start of a professional journey that would take them all over the world, working with clients from the USA, Canada, UK, Asia, Africa, and many more.",
        },
        {
          year: "2016",
          description:
            "As they gained more experience, they expanded their services to include iOS and backend development. Today, they help businesses generate more revenue, acquire more users, and deliver exceptional user experiences through mobile apps, websites, web apps, desktop apps, UI/UX, and SEO.",
        },
        {
          year: "2020",
          description:
            "At Canopas, they see themselves as problem solvers first and foremost. Technology is just the tool they use to make a difference in the world. Whether they're working on a new project or tackling an old problem, you can always count on them to deliver the best solution possible.",
        },
        {
          year: "2023",
          description:
            "If you're a business owner or an entrepreneur with a vision to help humanity, we would love to hear from you. With Canopas, you can be sure that your ideas will be transformed into reality and make a real impact in the world.",
          className: "tw-mb-9",
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>
