<template>
  <div>
    <div class="flex items-center lg:pt-0">
      <h1 class="mobile-header-2 lg:desk-header-2 primary-color pr-2">
        Hey Smarty
      </h1>
      <img
        :src="HandWing"
        alt="HandWing"
        class="w-[32px] md:w-[49.38px] h-[32px] md:h-[50.29px]"
      />
    </div>

    <div class="my-6 sub-h1-regular lg:mobile-header-2-regular text-black-87">
      Thanks for reaching out to us. This form requires
      <span class="sub-h1-semibold lg:mobile-header-2-semibold primary-color"
        >less than 60 seconds</span
      >
      to fill up.
    </div>
    <div class="sub-h3-semibold lg:sub-h2-medium text-black-60">
      Here's what you will get by filling out this form.
    </div>
    <div>
      <ul class="list-none m-auto p-0 text-left">
        <li
          v-for="reason in reasons"
          :key="reason"
          class="flex items-start px-0 pt-4"
        >
          <span
            class="mt-2 md:mt-3 mr-[15px] w-2 md:w-[8.81px] h-2 md:h-[9px] rounded-full bg-pink-300"
          ></span>
          <div
            class="w-11/12 lg:w-[90%] sub-h3-regular lg:mobile-header-3-regular text-black-87"
          >
            <span v-html="reason.title"></span>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import HandWing from "@/assets/images/theme/HandWing.png";
export default {
  data() {
    return {
      HandWing: HandWing,
      reasons: [
        {
          title: `Free consultation and QA session to know more about each other.`,
        },
        {
          title: `  The call is <span class='sub-h3-semibold lg:mobile-header-3-semibold'>only 15 minutes</span> and we promise not to give you a hard\
              sell. We will talk more ONLY if you find the conversation\
              interesting.`,
        },
        {
          title: ` To prove our <span class='sub-h3-semibold lg:mobile-header-3-semibold'>Skin In the Game,</span> we offer 6 Months of Free Post Deployment Maintenance for all the work we do on your project.`,
        },
      ],
    };
  },
};
</script>
