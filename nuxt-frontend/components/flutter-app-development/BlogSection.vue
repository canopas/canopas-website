<template>
  <section class="mt-16 lg:mt-[15rem]">
    <!-- --------------------Mobile UI Start----------------------- -->
    <div class="text-center h-[30rem] sm:!h-[36rem] md:!h-[42rem] lg:hidden">
      <p class="header-2 opacity-[0.87]">Our blogs</p>
      <div class="mt-2">
        <swiper
          :slidesPerView="1.425"
          :spaceBetween="0"
          :centeredSlides="true"
          class="!p-4"
        >
          <swiper-slide v-for="(blog, index) in blogs" :key="index">
            <div
              class="scale-75 text-left transition-all duration-200 ease-out"
              @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
            >
              <img
                :src="blog.image[0]"
                :srcset="`${blog.image[0]} 400w, ${blog.image[1]} 800w`"
                :alt="blog.title"
                class="h-full w-full rounded-xl object-cover"
                loading="lazy"
              />
              <div class="p-4 flex flex-col gap-3">
                <p class="sub-h1-semibold opacity-[0.87]">{{ blog.title }}</p>
                <p class="sub-h4-regular text-black-core/60">
                  {{ blog.readTime }}
                </p>
              </div>
            </div>
          </swiper-slide>
        </swiper>
        <a
          class="white-btn gradient-border-btn primary-btn mt-8"
          :href="canopas"
          target="_blank"
        >
          <span class="sub-h3-semibold v2-canopas-gradient-text">View all</span>
        </a>
      </div>
    </div>
    <!-- --------------------Desktop UI Start----------------------- -->
    <div class="container hidden lg:block lg:min-h-[580px]">
      <p class="mb-14 text-center header-2 text-black-core/[0.87]">Our blogs</p>
      <div class="flex flex-row space-x-5">
        <div
          @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
          v-for="blog in blogs.slice(0, 1)"
          :key="blog"
          class="content group relative flex basis-full"
        >
          <img
            :src="blog.image[2]"
            :srcset="`${blog.image[2]} 800w`"
            :alt="blog.title"
            class="object-fill h-full w-full cursor-pointer rounded-xl"
            loading="lazy"
          />
          <div
            class="bg-[#121212] ease absolute top-1/2 left-1/2 right-1/2 rounded-xl transition-all duration-700 opacity-0 overflow-hidden w-0 h-0 hover:opacity-100 group-hover:h-full group-hover:w-full group-hover:top-0 group-hover:left-0 group-hover:right-0"
          >
            <div
              class="absolute w-full -top-72 left-1/2 text-center -translate-x-1/2 -translate-y-1/2 transition-all duration-700 transform group-hover:top-1/2"
            >
              <span
                v-html="blog.hovertitle"
                @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
              ></span>
            </div>
          </div>
        </div>
        <div class="flex flex-col space-y-5 basis-[49%]">
          <div
            class="cursor-pointer group content relative"
            v-for="blog in blogs.slice(1, 3)"
            :key="blog"
            @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
            :class="blog.className"
          >
            <img
              :src="blog.image[1]"
              :srcset="`${blog.image[1]} 800w`"
              class="object-fill h-full w-full cursor-pointer rounded-xl"
              :alt="blog.title"
              loading="lazy"
            />
            <div
              class="bg-[#121212] ease absolute top-1/2 left-1/2 right-1/2 rounded-xl transition-all duration-700 opacity-0 overflow-hidden w-0 h-0 hover:opacity-100 group-hover:h-full group-hover:w-full group-hover:top-0 group-hover:left-0 group-hover:right-0"
            >
              <div
                class="absolute w-80 -top-72 left-1/2 text-center -translate-x-1/2 -translate-y-1/2 transition-all duration-700 transform group-hover:top-1/2"
              >
                <span
                  v-html="blog.hovertitle"
                  @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
                ></span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="flex space-x-5 mt-5">
        <div
          class="cursor-pointer group content relative mt-4 flex-col"
          v-for="blog in blogs.slice(3, 5)"
          :key="blog"
          @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
        >
          <img
            :src="blog.image[1]"
            :srcset="`${blog.image[1]} 800w`"
            class="object-fit h-full w-full cursor-pointer rounded-xl"
            loading="lazy"
            :alt="blog.title"
          />
          <div
            class="bg-[#121212] ease absolute top-1/2 left-1/2 right-1/2 rounded-xl transition-all duration-700 opacity-0 overflow-hidden w-0 h-0 hover:opacity-100 group-hover:h-full group-hover:w-full group-hover:top-0 group-hover:left-0 group-hover:right-0"
          >
            <div
              class="absolute w-96 -top-72 left-1/2 right-0 text-center -translate-x-1/2 -translate-y-1/2 transition-all duration-700 transform group-hover:top-1/2"
            >
              <span
                v-html="blog.hovertitle"
                @click="openBlog(blog.link, 'tap_flutter_app_blog_section')"
              ></span>
            </div>
          </div>
        </div>
      </div>
      <a
        class="primary-btn gradient-border-btn white-btn mt-[4.5rem] lg:mb-[15rem]"
        :href="canopas"
        target="_blank"
      >
        <span class="sub-h3-semibold">Read More Blogs </span>
        <Icon name="fa6-solid:arrow-right" class="fa ml-2" />
      </a>
    </div>
  </section>
</template>
<script>
import { Swiper, SwiperSlide } from "swiper/vue";
import { openBlog } from "@/utils.js";
import inherit_400w from "@/assets/images/flutter-app-development/blog/1-400w.webp";
import inherit_800w from "@/assets/images/flutter-app-development/blog/1-800w.webp";
import desk800w from "@/assets/images/flutter-app-development/blog/desk800w.webp";
import custom_400w from "@/assets/images/flutter-app-development/blog/2-400w.webp";
import custom_800w from "@/assets/images/flutter-app-development/blog/2-800w.webp";
import maintain_400w from "@/assets/images/flutter-app-development/blog/3-800w.webp";
import maintain_800w from "@/assets/images/flutter-app-development/blog/3-800w.webp";
import animate_400w from "@/assets/images/flutter-app-development/blog/4-800w.webp";
import animate_800w from "@/assets/images/flutter-app-development/blog/4-800w.webp";
import master_400w from "@/assets/images/flutter-app-development/blog/5-400w.webp";
import master_800w from "@/assets/images/flutter-app-development/blog/5-800w.webp";

export default {
  data() {
    return {
      openBlog,
      blogs: [
        {
          id: 1,
          image: [inherit_400w, inherit_800w, desk800w],
          title: "How to Use InheritedWidget in Flutter",
          readTime: "7 min",
          hovertitle: `<span class="sub-h1-semibold text-center text-white">How to Use InheritedWidget in Flutter</span>`,
          link: "https://blog.canopas.com/how-to-use-inheritedwidget-in-flutter-5fd64ce52b66",
        },
        {
          id: 2,
          image: [animate_400w, animate_800w],
          title:
            "Flutter Fun Animation A smiley that reacts to your every move",
          readTime: "8 min",
          hovertitle: `<span class="sub-h1-semibold text-center text-white">Flutter Fun Animation A smiley that reacts to your every move</span>`,
          link: "https://blog.canopas.com/flutter-animations-made-fun-a-smiley-that-reacts-to-your-every-move-9ebfb0a0657f",
        },
        {
          id: 3,
          image: [maintain_400w, maintain_800w],
          title: "Maintain the state of BottomNavigationBar across tabs",
          readTime: "9 min",
          hovertitle: `<span class="sub-h1-semibold text-center text-white">Maintain the state of BottomNavigationBar across tabs</span>`,
          link: "https://blog.canopas.com/how-to-maintain-the-state-of-bottomnavigationbar-across-tabs-528b3699f78b",
        },
        {
          id: 4,
          image: [custom_400w, custom_800w],
          title: "Customizing Serialization and Deserialization in Flutter",
          readTime: "10 min",
          hovertitle: `<span class="sub-h1-semibold text-center text-white">Customizing Serialization and Deserialization in Flutter</span>`,
          link: "https://blog.canopas.com/flutter-customizing-serialization-and-deserialization-with-json-serializable-db17e3aa90b0",
        },
        {
          id: 5,
          image: [master_400w, master_800w],
          title: "Mastering Flutter Inspector DevTools for Effective Debugging",
          readTime: "11 min",
          hovertitle: `<span class="sub-h1-semibold text-center text-white">Mastering Flutter Inspector DevTools for Effective Debugging</span>`,
          link: "https://blog.canopas.com/optimizing-flutter-apps-mastering-flutter-inspector-devtools-for-effective-debugging-493c03f68835",
        },
      ],
      canopas: "https://blog.canopas.com",
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
};
</script>
<style lang="postcss" scoped>
.swiper-slide-active div {
  @apply scale-[1] rounded-xl shadow-lg;
}
.gradient-border-btn > .fa {
  @apply text-[#f2709c];
}
.gradient-border-btn:hover > .fa {
  @apply text-white;
}
</style>
