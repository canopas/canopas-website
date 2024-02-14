<template>
  <section v-if="activeBlog" class="my-16 lg:my-60 text-black-87 h-full">
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
          Well, at least the community says Hell Yeah. Our blogs hosted on
          medium have
          <span
            class="v2-canopas-gradient-text sub-h1-semibold lg:mobile-header-2-semibold"
            >100k+</span
          >
          minutes monthly reading time and it’s only rising.
        </div>
      </div>
    </div>
    <div
      class="relative mt-8 h-full"
      :class="width > 767 && width < 992 ? 'container' : ''"
    >
      <div class="w-full h-full mr-0">
        <img
          :src="bgMobile"
          class="md:hidden sm:ml-[153px] w-full sm:w-[80%] h-full object-cover"
          alt="blog-background-image"
        />
      </div>
      <div
        class="absolute top-0 left-0 right-0 flex flex-col md:flex-row md:mt-8 lg:mt-16 z-[1] lg:min-h-[580px] md:pb-[15%] lg:py-0"
        :class="width < 768 || width > 2000 ? 'container' : ''"
      >
        <div
          class="relative flex flex-col items-center md:items-start md:justify-center md:flex-[60%] lg:ml-[7%] xl:ml-[10%] 2xl:ml-[15%] xll:ml-[1%]"
          :class="width > 2000 ? 'ml-0' : ''"
        >
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
            class="md:w-1/2 2xl:w-[55%] xll:w-[65%] 3xl:w-[55%] w-full"
          >
            <div class="flex flex-col 2xl:pr-16">
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
                  class="flex flex-col w-full"
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
                      {{ blog.authorName }}
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
                          class="!hidden lg:!inline-block arrow fa w-4 h-4 text-pink-300"
                          name="fa6-solid:arrow-right-long"
                          id="leftArrow"
                        />
                      </span>
                    </nuxt-link>
                  </div>
                  <hr class="h-[0.063rem] mt-4 lg:mt-8 bg-[#C6C6C6]" />
                  <nuxt-link
                    v-if="activeIndex == index"
                    :to="'/' + blog.slug"
                    @click.native="$mixpanel.track('tap_blog_post')"
                  >
                    <div
                      class="block md:!hidden mt-4 w-full animate-fadeInRight"
                    >
                      <img
                        :src="blog.image_url"
                        class="object-cover"
                        loading="lazy"
                        :alt="blog.title"
                      />
                      <hr
                        v-if="index != blogs.length - 1"
                        class="md:hidden mt-4 h-[0.063rem] bg-[#C6C6C6]"
                      />
                    </div>
                  </nuxt-link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="hidden md:flex lg:justify-end 3xl:justify-center flex-[40%]">
        <div>
          <img
            v-if="(width > 767 && width < 991) || width > 3839"
            :src="bgNew"
            loading="lazy"
            class="3xl:w-[55%] 3xl:mr-0 relative ml-[300px] 3xl:ml-[1156px] mt-[30px] 3xl:-mt-[24px]"
            alt="blog-background-image"
          />
          <img
            v-else
            :src="bg"
            loading="lazy"
            class="relative mt-14 lg:mt-[4.5rem]"
            alt="blog-background-image"
          />
          <nuxt-link
            :to="'/' + activeBlog.slug"
            @click.native="$mixpanel.track('tap_blog_post')"
          >
            <img
              v-if="activeBlog !== null"
              :src="activeBlog.image_url"
              class="absolute inline-block inset-0 m-auto w-[90%] md:top-8 lg:top-[5.5rem] xl:top-[7.5rem] 3xl:-top-[3.5rem] md:w-[47%] lg:w-[37%] xl:w-[35%] 2xl:w-[30%] xll:w-[26%] 3xl:w-[17%] md:-right-[387px] lg:-right-[570px] xl:-right-[680px] 2xl:-right-[945px] xll:-right-[1600px] 3xl:-right-[1246px] cursor-pointer"
              :class="animate ? 'animate-fadeInRight' : ''"
              loading="lazy"
              :alt="activeBlog.title"
            />
          </nuxt-link>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import bgMobile from "@/assets/images/blog/bg/bg400.svg";
import bg from "@/assets/images/blog/bg/bg2400.svg";
import bgNew from "@/assets/images/blog/bg/bg.webp";
import config from "@/config.js";
import { useBlogListStore } from "@/stores/resources";

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
