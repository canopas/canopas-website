<template>
  <div class="tw-p-0 tw-m-0 tw-relative">
    <img
      class="tw-sticky tw-top-0 tw-w-full tw-h-screen tw--z-[5] tw-object-cover"
      :src="developerBackgroundImage"
      loading="lazy"
      alt="developer-background-image"
    />
    <div
      class="tw-sticky tw-top-0 tw-w-full tw-h-screen tw--z-[5] tw-object-cover tw--mt-[100vh] tw-bg-black-900 tw-opacity-80"
    ></div>

    <div
      class="tw--mt-[100vh] tw-mx-auto tw-mb-0 tw-flex tw-flex-col tw-px-12 md:tw-pl-20 md:tw-pr-16 lg:tw-px-36 xl:tw-px-44 tw-container"
    >
      <div class="tw-pb-[10%] tw-pt-[10%]" id="item1">
        <div
          class="header-text canopas-gradient-text tw-my-8 tw-mx-0"
          data-scroll-class="item1"
        >
          What's more painful than hiring a developer or team who misunderstands
          every little thing about your project?
        </div>
        <div class="normal-text resources-text" data-scroll-class="item2">
          <div class="tw-text-white">
            <b> THREE REASONS WHY </b> You can’t afford to outsource your
            project to someone who struggles to understand you and your
            business.
          </div>
          <ul class="tw-text-white tw-pl-0 tw-mt-5 tw-my-8 tw-mx-0">
            <li class="tw-list-none tw-py-2.5 tw-px-0 tw-flex tw-opacity-75">
              <span
                class="tw-content-none tw-bg-gradient-to-t tw-from-orange-300 tw-to-pink-300 tw-inline-block tw-leading-none tw-mt-[3px] tw-mr-6 tw-mb-0 tw-ml-0; tw-h-auto tw-w-[6px] tw-rounded-none"
              ></span>
              <div class="tw-w-full tw-p-2.5">
                You might end up wasting your precious resources like Time,
                Money, and Energy.
              </div>
            </li>
            <li class="tw-list-none tw-py-2.5 tw-px-0 tw-flex tw-opacity-75">
              <span
                class="tw-content-none tw-bg-gradient-to-t tw-from-orange-300 tw-to-pink-300 tw-inline-block tw-leading-none tw-mt-[3px] tw-mr-6 tw-mb-0 tw-ml-0; tw-h-auto tw-w-[6px] tw-rounded-none"
              ></span>
              <div class="tw-w-full tw-p-2.5">
                Your final product might have issues like bad quality, poor
                performance, awful design, and sometimes you won’t receive
                anything.
              </div>
            </li>
            <li class="tw-list-none tw-py-2.5 tw-px-0 tw-flex tw-opacity-75">
              <span
                class="tw-content-none tw-bg-gradient-to-t tw-from-orange-300 tw-to-pink-300 tw-inline-block tw-leading-none tw-mt-[3px] tw-mr-6 tw-mb-0 tw-ml-0; tw-h-auto tw-w-[6px] tw-rounded-none"
              ></span>
              <div class="tw-w-full tw-p-2.5">
                Most IMPORTANTLY, unsuccessful business relationships might
                affect your mood and behaviour. You might experience stress,
                anxiety, negativity, insecurity, uncertainty, etc. It might
                influence your behaviour, and you may end up doubting everyone
                around you including YOURSELF.
              </div>
            </li>
          </ul>
        </div>
      </div>
      <div class="tw-pb-[10%] tw-my-8 tw-mx-0" id="item2">
        <div class="header-text canopas-gradient-text tw-my-8 tw-mx-0">
          You need a solution that will reward you with money, peace, and
          satisfaction, not headache and nightmares.
        </div>
        <div class="normal-text resources-text">
          <div class="tw-text-white">
            You probably need a team that helps you to develop a digital product
            that is…
          </div>

          <ul class="tw-p-2.5 tw-pl-0 tw-text-white mt-20 tw-my-8 tw-mx-0">
            <li class="tw-list-none tw-py-2.5 tw-px-0 tw-flex tw-opacity-75">
              <span
                class="tw-content-none tw-bg-gradient-to-t tw-from-orange-300 tw-to-pink-300 tw-inline-block tw-leading-none tw--mt-[8px] tw-mr-6 tw--mb-[9px] tw-ml-0; tw-h-auto tw-w-[6px] tw-rounded-none"
              ></span>
              <div class="tw-w-full">
                Simple<br />
                Easy to use<br />
                Beautifully Crafted<br />
                And Users love it.<br />
              </div>
            </li>
          </ul>
          <div class="tw-text-white">
            But, I guess that’s not enough
            <br />
            <br />
            You want somebody who asks questions, suggests ways to improve your
            apps and tries to understand YOU, your vision(or helps you achieve
            your goals or dreams), and the value of your app or product.
            <br />
            <br />
            You’re looking for a team who knows that your time is limited and
            knows how to use it wisely to build a high-quality product.
            <br />
            <br />
            Canopas understands these issues and knows how to handle them. So,
            you don't have to.
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import developerBackgroundImage from "@/assets/images/theme/developer.webp";

export default {
  data() {
    return {
      developerBackgroundImage: developerBackgroundImage,
    };
  },
};
</script>
