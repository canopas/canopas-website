<template>
  <section class="bg-[#121212] 3xl:outer-container">
    <div class="py-14 lg:py-[7.5rem] container text-white">
      <!-- ---------------------MOBILE UI START------------------ -->
      <div class="flex flex-col gap-8 lg:hidden">
        <h2 class="mobile-header-1 lg:desk-header-1 text-center">
          Why should you partner with Canopas?
        </h2>
        <div class="flex flex-col gap-6" v-for="item in items" :key="item">
          <div class="flex gap-4">
            <img
              :src="item.image"
              class="w-9 h-9 object-contain"
              :alt="item.title"
            />
            <div class="flex flex-col gap-3">
              <span class="sub-h1-semibold lg:mobile-header-2-semibold">{{
                item.title
              }}</span>
              <span
                class="sub-h3-regular text-white-core-80"
                v-html="item.description"
              ></span>
            </div>
          </div>
        </div>
        <nuxt-link
          class="gradient-btn primary-btn mt-0"
          to="/contact"
          @click.native="$mixpanel.track('tap_backend_partner_with_us_cta')"
        >
          <span class="sub-h3-semibold">Let’s Work Together</span></nuxt-link
        >
      </div>
      <!-- ---------------------MOBILE UI END-------------------- -->

      <!-- ---------------------DESKTOP UI START------------------ -->
      <div id="partner-with-us" class="hidden lg:block h-[150vh]">
        <div class="sticky top-0 h-[110vh] xl:top-4 xl:h-[100vh] main">
          <div
            class="absolute -translate-y-1/2 top-1/2 text-white"
            :class="[isAnimated ? '!flex flex-1' : 'left-1/2 -translate-x-1/2']"
          >
            <div class="flex-col my-auto" :class="{ '!flex': isAnimated }">
              <span
                class="desk-header-2 font-normal xl:!font-bold"
                :class="[isAnimated ? 'opacity-100' : 'opacity-[0.12]']"
                >Why should you</span
              >
              <span
                class="block ml-2 transition-all duration-1000"
                :class="[
                  isAnimated
                    ? 'opacity-100 !desk-header-2 font-normal xl:!font-bold !ml-0'
                    : 'text1 opacity-[0.12]',
                ]"
                >partner
              </span>
              <div
                class="flex ml-0 xl:ml-32 gap-6"
                :class="{ '!ml-0 block': isAnimated }"
              >
                <span
                  class="my-auto -rotate-90 desk-header-2 font-normal xl:!font-bold opacity-[0.12]"
                  :class="{
                    rotate: isAnimated,
                  }"
                  >with</span
                ><span
                  class="transition-all duration-1000"
                  :class="[
                    isAnimated
                      ? 'opacity-100 !desk-header-2 font-normal xl:!font-bold'
                      : 'text1 opacity-[0.12]',
                  ]"
                  >Canopas?</span
                >
              </div>
              <nuxt-link
                class="gradient-btn primary-btn ml-0 mt-10 animate-fadeIn"
                :class="[!isAnimated ? 'hidden' : '']"
                to="/contact"
                @click.native="
                  $mixpanel.track('tap_backend_partner_with_us_cta')
                "
              >
                <span class="sub-h3-semibold lg:sub-h1-semibold"
                  >Let’s Work Together</span
                ></nuxt-link
              >
            </div>
            <nuxt-link
              class="gradient-btn primary-btn mt-[5.69rem]"
              :class="{ '!hidden': isAnimated }"
              to="/contact"
              @click.native="$mixpanel.track('tap_backend_partner_with_us_cta')"
            >
              <span class="sub-h3-semibold lg:sub-h1-semibold"
                >Let’s Work Together</span
              ></nuxt-link
            >
          </div>
          <div
            class="flex flex-col absolute w-[50%] left-1/2 -translate-y-1/2 top-1/2 overflow-hidden"
          >
            <div
              v-for="(item, index) in items"
              :key="item"
              class="flex flex-col mb-10 item"
              :class="[
                isVisible
                  ? `item-${index} transform-all duration-1000 opacity-100`
                  : 'translate-x-full',
              ]"
            >
              <div class="flex gap-[1.69rem]">
                <img
                  :src="item.image"
                  class="w-9 h-9 xl:w-14 xl:h-14 object-contain"
                  :alt="item.title"
                />
                <div class="flex flex-col gap-4">
                  <span
                    class="mobile-header-3-semibold xl:mobile-header-2-semibold"
                    >{{ item.title }}</span
                  >
                  <span
                    class="sub-h3-regular xl:mobile-header-3-regular text-white-core-80"
                    v-html="item.description"
                  ></span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- ---------------------DESKTOP UI END-------------------- -->
    </div>
  </section>
</template>
<script setup>
import maintainance from "@/assets/images/backend-development/partnerwithus/maintanance.webp";
import moneyback from "@/assets/images/backend-development/partnerwithus/moneyback.webp";
import user from "@/assets/images/backend-development/partnerwithus/user.webp";
const { $mixpanel } = useNuxtApp();
const isAnimated = ref(false);
const isVisible = ref(false);
const items = [
  {
    title: "Skin In The Game",
    image: maintainance,
    description: `To prove our skin in the game, we offer <span class="text-[#FF9472] sub-h3-semibold xl:mobile-header-3-semibold">6 Months of Free Post Deployment Maintenance</span> for all the work we do on your project. If any bugs arise in the project within 6 months of deployment, we'll do it for free. This approach ensures careful development, as neglect could result in significant expenses.`,
  },
  {
    title: "100% Risk Free Experience / Money Back Guarantee",
    image: moneyback,
    description: `To walk the talk, we offer a <span class="text-[#FF9472] sub-h3-semibold xl:mobile-header-3-semibold">100% Money Back Guarantee for two weeks of work at any time during our engagement</span>. If you work with us, you will get a feedback survey every Monday. If you ain't happy with the work we did in the past week, get your money back.`,
  },
  {
    title: "Achieving Excellence - One step at a time",
    image: user,
    description: `Whether your goal is user growth, user engagement, revenue growth, user satisfaction, improved user experience, unforgettable branding, customer loyalty, or all of them, we know how to achieve it, one step at a time!`,
  },
];

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
function handleScroll() {
  const stickyparent = document.getElementById("partner-with-us");
  const stickyparenthead = stickyparent.getBoundingClientRect();
  if (stickyparenthead.top < 140) {
    isAnimated.value = true;
    isVisible.value = true;
  }
}
</script>
<style scoped lang="postcss">
.item-0 {
  @apply translate-x-0 delay-1000;
}
.item-1 {
  @apply translate-x-0 delay-[2s];
}
.item-2 {
  @apply translate-x-0 delay-[3s];
}
.text1 {
  @apply text-[10.5rem] xl:text-[12.5rem] leading-[14.25rem] xl:leading-[16.25rem] xl:font-bold;
}
.rotate {
  animation: rotate 0.8s linear forwards;
}

@keyframes rotate {
  0% {
    transform: rotate(-90deg);
    position: absolute;
    top: 60%;
    left: 0;
  }
  100% {
    transform: rotate(0deg);
    position: absolute;
    opacity: 1;
    top: 22%;
    left: 51%;
  }
}
</style>
