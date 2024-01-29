<template>
  <section class="bg-deep-charcoal mb-16">
    <div class="container pb-8">
      <h1 class="pt-14 text-center mobile-header-2 text-white">
        How it all started?
      </h1>
      <div
        class="container flex flex-col items-center w-full md:w-[82%] h-[24.25rem] mt-8 overflow-y-scroll hidden-scrollbar before:-mt-5 before:blur-lg before:absolute before:left-0 before:h-[6%] before:w-full before:bg-black-gradient"
      >
        <div v-for="(event, index) in events" :key="event">
          <div :class="[event.className]" class="flex">
            <div class="flex flex-col">
              <h1 class="mobile-header-2 text-primary-1">
                {{ event.year }}
              </h1>
              <div
                :class="index == events.length - 1 ? 'hidden' : ''"
                class="ml-6 h-60 sm:h-40 md:h-48 border-dotted border-l-2 border-black-87"
              ></div>
            </div>
            <div class="ml-6 sub-h3-medium text-white">
              {{ event.description }}
            </div>
          </div>
        </div>
      </div>
      <div
        class="-mt-6 w-full h-12 after:absolute after:left-0 after:h-[5%] after:w-full after:blur-lg after:bg-black-gradient-bottom"
      ></div>
    </div>
  </section>
</template>

<script type="module">
export default {
  data() {
    return {
      events: [
        {
          year: "2012",
          description:
            "Once upon a time, in the bustling streets of India, two childhood friends Darpan and Jignesh had a dream of using technology to make a difference in the world. As 90s kids, they watched in awe as technology advanced at lightning speed, and they knew that they had to be a part of this revolution.",
          className: "mt-9",
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
          className: "mb-9",
        },
      ],
    };
  },
  components: {},
};
</script>
