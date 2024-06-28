<template>
  <div class="mt-14 lg:hidden w-full flex justify-center">
    <hr class="w-[95%] sm:w-[75%]" />
  </div>
  <section
    v-if="activeBlog"
    class="mt-16 mb-24 lg:my-60 text-black-87 h-full mx-5 sm:mx-0"
  >
    <div class="container flex flex-col items-center m-auto w-full">
      <div class="container">
        <p class="text-white-smoke background-text text-center">Blogs</p>
        <h2
          class="-mt-5 md:mt-[-50px] mx-auto lg:w-[71%] xl:w-[65%] mobile-header-2 lg:desk-header-2 text-center"
        >
          Do we have a team with the right skills?
        </h2>
      </div>
      <div class="mt-4 lg:w-[80%]">
        <div
          class="w-full text-center text-black-60 sub-h1-regular lg:mobile-header-2-regular"
        >
          Our
          <span
            class="v2-canopas-gradient-text sub-h1-semibold lg:mobile-header-2-semibold"
            >5-star Clutch ratings</span
          >
          and love from the developer community speak for us. Our blogs attract
          <span
            class="v2-canopas-gradient-text sub-h1-semibold lg:mobile-header-2-semibold"
            >50K+ monthly visitors</span
          >, and the number keeps growing.
        </div>
      </div>
    </div>

    <div class="xll:container relative">
      <div class="flex gap-8 container">
        <div class="md:w-[64%]">
          <div
            v-for="(blog, index) in blogs"
            :key="blog"
            @mouseenter="
              activeIndex = index;
              activeBlog = blog;
              animate = true;
            "
            @touchstart.passive="activeIndex = index"
            @click.native="$mixpanel.track('tap_blog_post')"
            class="w-full"
          >
            <div class="flex flex-col">
              <div
                class="flex gap-8 md:gap-4 lg:gap-8 justify-end mt-[18px] md:mt-4 lg:mt-8"
              >
                <div
                  class="w-[3%] flex flex-col items-center text-black-87 lg:text-black-60"
                  :class="index == 0 ? 'mt-4 lg:mt-8' : ''"
                >
                  <span
                    class="mobile-header-2 lg:desk-header-2"
                    @mouseleave="animate = false"
                  >
                    {{ blog.published_on[1] }}
                  </span>
                  <span
                    class="text-center sub-h4-regular lg:sub-h4-medium"
                    @mouseleave="animate = false"
                  >
                    {{ blog.published_on[0] }}
                  </span>
                </div>

                <div
                  class="flex flex-col gap-5 w-full"
                  @touchstart.passive="
                    activeBlog = blog;
                    animate = true;
                  "
                  @touchend="animate = false"
                >
                  <hr
                    v-if="index == 0"
                    class="h-[0.063rem] mb-4 lg:mb-8 bg-[#C6C6C6]"
                  />
                  <div>
                    <div @mouseleave="animate = false">
                      <nuxt-link
                        :to="'/' + blog.slug"
                        @click.native="$mixpanel.track('tap_blog_post')"
                      >
                        <span
                          :class="
                            activeIndex == index
                              ? 'bg-gradient-underline-out box-decoration-clone bg-no-repeat pb-[5px] transition-all duration-500 bg-[length:100%] hover:bg-gradient-underline-out'
                              : 'bg-[length:0%]'
                          "
                          class="sub-h1-semibold lg:mobile-header-2-semibold xl:desk-header-3 text-black-87 cursor-pointer"
                        >
                          {{ blog.title }}</span
                        >
                      </nuxt-link>
                    </div>
                    <div
                      class="flex justify-between mt-[0.9375rem] md:mt-4 lg:mt-8 text-black-60"
                    >
                      <span class="w-auto sub-h4-regular lg:sub-h1-semibold">
                        {{ blog.author.name }}
                      </span>
                      <nuxt-link
                        :to="'/' + blog.slug"
                        @click.native="$mixpanel.track('tap_blog_post')"
                      >
                        <span
                          class="lg:w-[7.938rem] cursor-pointer sub-h3-semibold lg:mobile-header-3-semibold lg:v2-canopas-gradient-text"
                        >
                          Read more
                          <Icon
                            class="arrow fa w-4 h-4 text-pink-300 !hidden lg:!inline-block"
                            name="fa6-solid:arrow-right-long"
                            id="leftArrow"
                          />
                        </span>
                      </nuxt-link>
                    </div>
                  </div>

                  <hr class="h-[0.063rem] bg-[#C6C6C6]" />
                  <nuxt-link
                    v-if="activeIndex == index"
                    :to="'/' + blog.slug"
                    @click.native="$mixpanel.track('tap_blog_post')"
                    class="md:hidden"
                  >
                    <div class="animate-fadeInRight">
                      <img
                        :src="blog.image_url"
                        class="h-full w-full rounded-lg"
                        loading="lazy"
                        :alt="blog.title"
                      />
                      <hr v-if="index != blogs.length - 1" class="mt-5" />
                    </div>
                  </nuxt-link>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="hidden md:flex w-[36%] items-center justify-center">
          <nuxt-link
            :to="'/' + activeBlog.slug"
            @click.native="$mixpanel.track('tap_blog_post')"
          >
            <img
              v-if="activeBlog !== null"
              :src="activeBlog.image_url"
              class="w-full cursor-pointer rounded-lg"
              :class="animate ? 'animate-fadeInRight' : ''"
              loading="lazy"
              :alt="activeBlog.title"
            />
          </nuxt-link>
        </div>
      </div>
      <img
        class="absolute top-0 -right-5 sm:right-0 h-full z-[-1]"
        :src="bg"
        :srcset="`${bg} 400w, ${bg400} 800w, ${bg2400} 1200w`"
        alt="blog-background-image"
        loading="lazy"
      />
    </div>
  </section>
</template>

<script setup>
import config from "@/config.js";
import { useBlogListStore } from "@/stores/resources";
import bg from "@/assets/images/blog/bg/bg.webp";
import bg400 from "@/assets/images/blog/bg/bg400.svg";
import bg2400 from "@/assets/images/blog/bg/bg2400.svg";

const { $mixpanel } = useNuxtApp();

const blogs = ref([]);
const width = ref(680);
const activeBlog = ref(null);
const activeIndex = ref(0);
const animate = ref(false);

const store = useBlogListStore();
const resources = computed(() => store.items);
const status = computed(() => store.status);

await useAsyncData("blogs", () =>
  store.loadResources(config.SHOW_DRAFT_POSTS, false, 0, 3),
);

if (status.value === config.SUCCESS) {
  blogs.value = resources.value;
  activeBlog.value = blogs.value[0];
}

onMounted(() => {
  for (const element of blogs.value) {
    element.published_on = element.published_on.replace(",", "").split(" ", 2);
  }
  width.value = window.innerWidth;
});
</script>
