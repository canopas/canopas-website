<template>
  <section class="container my-16 lg:my-60">
    <div class="text-center mobile-header-2 lg:desk-header-2 text-black-87">
      We are here to contribute
    </div>
    <div
      class="mt-4 mb-8 lg:mb-[4.5rem] text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
    >
      and that starts with this site itself,
      <nuxt-link
        class="sub-h3-semibold lg:mobile-header-2-semibold v2-canopas-gradient-text border-b border-[#ff835b]"
        :to="websiteOpenSourceUrl"
        target="_blank"
        @click.native="$mixpanel.track('tap_canopas_website_github')"
        >canopas is open source!</nuxt-link
      >
    </div>
    <div
      class="flex flex-col border mt-2 lg:mt-0"
      v-for="(contribution, index) in contributions"
      :key="index"
    >
      <nuxt-link
        :to="contribution.link"
        target="_blank"
        @click.native="$mixpanel.track('tap_home_github_contribution_section')"
        class="flex-col p-4 lg:p-6"
      >
        <div class="mobile-header-3-semibold lg:desk-header-3 text-black-87">
          {{ contribution.title }}
        </div>
        <div
          class="mt-2 sub-h3-regular lg:mobile-header-2-regular text-black-60"
        >
          {{ contribution.description }}
        </div>
        <div
          v-if="contribution.tags.length > 0"
          class="my-6 overflow-hidden overflow-x-scroll scrollbar-hidden"
        >
          <div class="w-full min-w-[1000px] lg:min-w-full flex-wrap flex gap-2">
            <div
              class="mt-1 inline-flex"
              v-for="(tag, index) in contribution.tags"
              :key="index"
            >
              <p
                class="bg-white-smoke rounded-[1.875rem] py-1 px-4 lg:sub-h1-regular text-black-4"
              >
                {{ tag }}
              </p>
            </div>
          </div>
        </div>
        <div class="flex gap-4 mt-4 lg:mt-6">
          <div class="flex items-center">
            <Icon
              class="fab footer-icon h-[1.625rem] w-[1.625rem] pr-[0.313rem]"
              name="fa6-solid:star"
            />
            <span
              class="v2-canopas-gradient-text mt-0.5 sub-h3-semibold lg:mobile-header-3-semibold"
              >{{ contribution.stars }}</span
            >
          </div>
          <div
            v-if="contribution.forks > 0"
            class="flex items-center gap-0.5 text-black-87"
          >
            <span> <Icon class="fa h-5 w-5" name="fa6-solid:code-fork" /></span
            ><span class="sub-h3-semibold lg:mobile-header-3-semibold">{{
              contribution.forks
            }}</span>
          </div>
          <div class="flex items-center gap-2">
            <span
              class="h-4 w-4 rounded-full"
              :style="{ backgroundColor: contribution.color }"
            ></span
            ><span
              class="sub-h3-semibold lg:mobile-header-3-semibold text-black-87"
              >{{ contribution.language }}</span
            >
          </div>
        </div>
      </nuxt-link>
    </div>
  </section>
</template>
<script type="module">
import contributionsJson from "@/json-data/contribution";

export default {
  data() {
    return {
      width: 0,
      websiteOpenSourceUrl: "https://github.com/canopas/canopas-website",
      contributions: contributionsJson,
    };
  },
  mounted() {
    this.width = window.innerWidth;
  },
  inject: ["mixpanel"],
};
</script>

<style scoped>
.scrollbar-hidden {
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}

/* Hide scrollbar for Chrome, Safari and Opera */
.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
</style>
