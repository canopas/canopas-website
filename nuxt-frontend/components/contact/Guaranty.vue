<template>
  <div class="container py-16 lg:py-28 3xl:pb-[23rem]">
    <div
      class="rounded-xl border bg-white-smoke shadow-[0px_2px_35px_rgba(0,0,0,0.12)] px-4 py-8 lg:py-14"
    >
      <div class="flex items-center lg:items-end gap-1">
        <img
          :src="badge"
          class="w-14 lg:w-[6.75rem] h-[3.188rem] lg:h-[6.188rem] ml-[-0.313rem] lg:ml-0 lg:mt-[-2.188rem]"
          alt="badge-image"
        />
        <p
          class="sub-h3-semibold lg:mobile-header-2-semibold w-full lg:text-center lg:w-[82%] lg:mb-[-0.5rem] lg:ml-6"
        >
          In two weeks — at any time, if you feel you won't get the results you
          want, you don't have to pay a penny.
        </p>
      </div>
      <div
        class="flex flex-col sub-h3-medium lg:mobile-header-3-medium text-black-60 lg:px-14 mt-6"
      >
        <div v-for="benefit in benefits" class="mt-4" :key="benefit">
          <span v-html="benefit.title"></span>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import badge from "@/assets/images/guaranty-badge.webp";
export default {
  data() {
    return {
      badge,
      benefits: [
        {
          title: ` We don't want you to decide yes or no today... We want you to make a
          fully informed decision, that is all.
          <span class="sub-h3-semibold lg:mobile-header-3-semibold text-black-87"
            >The only way you can make a fully informed decision is on the
            inside, not the outside.</span
          >
          So you get on the ride with us and see if everything we say on this
          website is true and valuable to you.`,
        },
        {
          title: `Then, if it is, that's when you decide to keep it. If not, no hard
          feelings. By working with us for two weeks, you will be able to make
          an informed decision that this isn't for you. But you can't make that
          decision right now for the same reason you don't buy a car without a
          test drive.`,
        },
        {
          title: `And remember,
          <span class="sub-h3-semibold lg:mobile-header-3-semibold text-black-87"
            >whether it's 79 minutes or 79 hours from now, if you ain't happy,
            we ain't happy.</span
          >
          For any reason whatsoever, if you want your money back you can get it
          because we only want to keep your money if you're happy.`,
        },
        {
          title: `All you have to do is
          <span class="sub-h3-semibold lg:mobile-header-3-semibold text-black-87"
            >just tell us "This isn't for me" and you got it, No Questions
            Asked.</span
          >
          And our response times to any support request average 59min over a
          daytime period.`,
        },
        {
          title: `You can only make such a guarantee when you're confident that what you
          have is the real deal and we are damn confident that when you sign up
          with us you're getting exactly what you need to
          <span class="sub-h3-semibold lg:mobile-header-3-semibold text-black-87"
            >GROW YOUR BUSINESS.</span
          >`,
        },
      ],
    };
  },
};
</script>
