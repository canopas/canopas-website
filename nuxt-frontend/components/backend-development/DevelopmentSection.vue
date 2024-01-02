<template>
  <section class="pt-16 lg:pt-[7.5rem] 3xl:outer-container lg:bg-[#F8F8F8]">
    <div
      class="container text-center flex flex-col gap-4 lg:gap-6 mb-6 lg:mb-8"
    >
      <h2 class="header-2 text-black-core/[0.87]">
        Backend development services
      </h2>
      <p class="sub-h1-regular text-black-core/[0.6]">
        Our backend development services ensure that your software runs
        flawlessly, scales effortlessly, and remains bulletproof. We craft
        unique solutions that cater to your distinct business challenges and
        aspirations.
      </p>
    </div>
    <div id="stickyParent" class="sticky-parent h-[700vh]">
      <div
        class="bg-[#F8F8F8] sticky top-0 h-[100vh] main overflow-hidden snap-mandatory snap-y"
      >
        <div
          v-for="(item, index) in items"
          :key="index"
          class="section container flex h-screen snap-start pt-10"
        >
          <div class="flex flex-col pb-10">
            <div class="flex-1 relative">
              <div class="absolute">
                <div class="circle">
                  <svg
                    width="574"
                    height="575"
                    viewBox="0 0 574 575"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      v-for="(path, index) in paths"
                      :key="index"
                      :d="path.d"
                      :stroke="path.strokeColor"
                      :strokeWidth="
                        index === activeIndex ? '35px' : path.strokeWidth
                      "
                      :fill="path.fill"
                      :opacity="path.opacity"
                    />
                  </svg>
                </div>

                <div
                  class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-[50%] bg-white w-[18.204rem] h-[18.17563rem]"
                >
                  <div
                    class="absolute top-[45px] left-[3rem] rounded-[50%] bg-[#F8F8F8] w-[12.204rem] h-[12.17563rem]"
                  >
                    <div
                      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-[50%] bg-white w-[3rem] h-[3rem]"
                    >
                      <div
                        class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
                      >
                        <img
                          :src="item.image"
                          class="w-10 lg:w-16"
                          :alt="item.title"
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div
            class="flex-1 flex flex-col mt-[24rem]"
            :class="{ 'animate-fadeIn': isAnimated }"
          >
            <h3 class="header-3" :class="item.className">
              {{ item.title }}
            </h3>
            <span
              v-html="item.description"
              class="sub-h3-regular lg:sub-h3-medium flex flex-col gap-4 text-black-core/[0.6]"
            ></span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup>
import { Autoplay, Zoom } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { ref, computed } from "vue";
import custom from "@/assets/images/backend-development/development/customarchitecture.webp";
import development from "@/assets/images/backend-development/development/Development.webp";
import database from "@/assets/images/backend-development/development/DatabaseDesign.webp";
import performance from "@/assets/images/backend-development/development/PerformanceOptimization.webp";
import migration from "@/assets/images/backend-development/development/Migration.webp";
import security from "@/assets/images/backend-development/development/Security.webp";

const modules = [Autoplay, Zoom];
const paths = [
  {
    opacity: "0.2",
    d: "M181.107 467.768C149.995 449.49 124.135 423.471 106.039 392.237C87.9434 361.003 78.2264 325.616 77.8322 289.514L89.9435 289.381C90.3149 323.393 99.4693 356.731 116.517 386.157C133.566 415.582 157.928 440.095 187.238 457.315L181.107 467.768Z",
    strokeColor: "#F2709C",
    fill: "#F2709C",
    strokeWidth: 1,
  },
  {
    opacity: "0.2",
    d: "M392.32 468.112C361.049 486.347 325.586 496.15 289.396 496.564C253.207 496.978 217.529 487.989 185.849 470.475L191.706 459.864C221.551 476.365 255.163 484.834 289.258 484.444C323.352 484.053 356.762 474.817 386.222 457.638L392.32 468.112Z",
    strokeColor: "#FF9472",
    fill: "#FF9472",
    strokeWidth: 1,
  },
  {
    opacity: "0.2",
    d: "M496.18 289.585C495.785 324.716 486.561 359.183 469.359 389.81C452.156 420.437 427.529 446.238 397.745 464.837L391.333 454.554C419.392 437.032 442.594 412.724 458.8 383.87C475.007 355.017 483.697 322.546 484.069 289.448L496.18 289.585Z",
    strokeColor: "#F2709C",
    fill: "#F2709C",
    strokeWidth: 1,
  },
  {
    d: "M403.995 90.2581C438.824 108.798 467.704 136.816 487.304 171.078C506.904 205.341 516.423 244.451 514.764 283.896L480.508 282.453C481.895 249.473 473.936 216.773 457.549 188.126C441.162 159.479 417.015 136.053 387.894 120.552L403.995 90.2581Z",
    strokeColor: "#FF9472",
    fill: "#FF9472",
    strokeWidth: 1,
  },
  {
    opacity: "0.2",
    d: "M77.8467 283.759C78.4498 247.324 88.5451 211.678 107.134 180.345C125.724 149.012 152.163 123.078 183.838 105.106L189.812 115.651C159.971 132.582 135.062 157.015 117.549 186.533C100.036 216.052 90.5253 249.634 89.9571 283.96L77.8467 283.759Z",
    strokeColor: "#FF9472",
    fill: "#FF9472",
    strokeWidth: 1,
  },
  {
    opacity: "0.2",
    d: "M188.317 102.635C219.385 85.9993 254.149 77.4874 289.382 77.8887C324.615 78.2901 359.176 87.5917 389.858 104.93L383.903 115.485C354.997 99.1505 322.437 90.3875 289.244 90.0094C256.051 89.6313 223.3 97.6504 194.03 113.322L188.317 102.635Z",
    strokeColor: "#F2709C",
    fill: "#F2709C",
    strokeWidth: 1,
  },
  // Add other paths data here
];

const items = [
  {
    title: "Custom Backend Architecture",
    image: custom,
    className: "text-[#FF9472]",
    description: `<span>In today's dynamic digital landscape, a one-size-fits-all approach often falls short. Recognizing the unique challenges and objectives of each business,we offer tailored backend solutions built from the ground up.</span>
    <span>Our expertise ensures that your application functions with precision, aligning seamlessly with your front-end while meeting specific business logic and operational needs.</span>
    <span>At Canopas, we craft unique solutions that cater to your distinct business challenges and aspirations.</span>`,
  },
  {
    title: "Seamless API Development & Integration",
    image: development,
    className: "text-[#F2709C]",
    description: `<span>As businesses increasingly rely on a myriad of software applications, the need for efficient communication between these platforms becomes paramount.</span>
    <span>Our team crafts seamless and efficient APIs, ensuring that your various applications connect and relay information without hitches. This focus on integration fosters smooth data flow, ultimately enhancing operational efficiency.</span>
    <span>From monolithic to serverless microservices, we've got you covered with skills that help you scale your business. We put a great emphasis on API documentation for future integrations and updates.</span>`,
  },
  {
    title: "Database Design & Management",
    image: database,
    className: "text-[#FF9472]",
    description: `<span>In the digital era, data is gold. Our approach to database design prioritizes structure, optimization, and scalability.</span>
    <span>We ensure that you have a robust system in place to securely store, manage, and retrieve vast amounts of data. By choosing us, you're not only ensuring immediate data efficiency but also preparing your business for future growth.</span>
    <span>For database selection, we analyze project requirements first and then make a decision based on the feature requirements. Sometimes, we use relational and non-relational databases in a single project..</span>`,
  },
  {
    title: "Performance Optimization",
    image: performance,
    className: "text-[#F2709C]",
    description: `<span>A slow server response can be the chink in your digital armor, deterring potential users and impacting user experience.</span>
    <span>We go beyond mere backend development by enhancing server responses and actively working on reducing load times. This ensures that your users enjoy an uninterrupted and swift experience, increasing satisfaction and engagement.</span>
    <span>We use techniques like data sharding, caching, and database optimization. It helps you bring more repeat customers and happy customers who advocate your brand to others.</span>`,
  },
  {
    title: "Migration & Upgrades",
    image: migration,
    className: "text-[#FF9472]",
    description: `<span>Legacy systems, while familiar, often hinder growth and efficiency. Our team specializes in transitioning your old systems to contemporary, more efficient platforms.</span>
    <span>Through meticulous planning and execution, we ensure no data loss ordowntime, making the migration process seamless and painless.</span>
    <span>We also utilize CI/CD platforms and automate deployment pipelines to push changes efficiently. We always make sure that new updates don't break the experience for users on older versions or systems.</span>`,
  },
  {
    title: "Security & Compliance",
    image: security,
    className: "text-[#F2709C]",
    description: `<span>In an age where cyber threats are prevalent, the security of your backend is non-negotiable.</span>
    <span>We embed robust security measures from the ground up, ensuring databreaches remain a concern of the past.</span>
    <span>Additionally, with evolving industry standards, we ensure that your backend is not only secure but also compliant, giving you peace of mind and earning the trust of your users.</span>
    <span>We promote regular training sessions and workshops on the latest developments in the industry and encourage the team to stay updated with all the security and regulatory compliances.</span>`,
  },
];

const activeIndex = ref(0);
const isAnimated = ref(false);
const handleScroll = () => {
  const stickyParent = document.getElementById("stickyParent");
  const stickyTop = stickyParent.offsetTop;
  var position = stickyParent.getBoundingClientRect();

  const sticky = document.querySelector(".main");
  const items = document.querySelectorAll(".section");
  const scrollPosition = window.scrollY || window.pageYOffset;

  const windowHeight = window.innerHeight;
  const paths = document.querySelectorAll("path");

  if (scrollPosition >= stickyTop) {
    const scrollAmount = scrollPosition - stickyTop;
    sticky.scrollTop = scrollAmount;
  }
  items.forEach((item) => {
    const react = item.getBoundingClientRect();
    const itemOffset = react.top + scrollPosition;
    // You might need to adjust the condition based on your layout and when you want the animation to trigger
    if (scrollPosition >= itemOffset - windowHeight / 2) {
      // Find the corresponding circle element for the current item
      const circleElement = item.querySelector(".circle");

      // Set the animation flag to true for other elements
      isAnimated.value = true;
      circleElement.style.transform = `rotate(${window.scrollY * 0.1}deg)`;
      paths.forEach((path, index) => {
        activeIndex.value = index;
      });
    } else {
      isAnimated.value = false;
    }
  });
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});
onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
</script>
<style scoped>
.circle {
  overflow: hidden;
}
li {
  overflow: hidden;
  position: absolute;
  top: 0;
  right: 0;
  width: 50%;
  height: 50%;
  transform-origin: 0% 100%;
  margin-bottom: 10px; /* Add margin bottom here */
}
.text {
  position: absolute;
  left: -100%;
  width: 200%;
  height: 200%;
  text-align: center;
  -webkit-transform: skewY(30deg) rotate(15deg);
  -ms-transform: skewY(30deg) rotate(15deg);
  transform: skewY(30deg) rotate(15deg);
  padding-top: 20px;
}

li:first-child {
  -webkit-transform: rotate(0deg) skewY(-31deg);
  -ms-transform: rotate(0deg) skewY(-31deg);
  transform: rotate(0deg) skewY(-31deg);
}
li:nth-child(2) {
  -webkit-transform: rotate(60deg) skewY(-31deg);
  -ms-transform: rotate(60deg) skewY(-31deg);
  transform: rotate(60deg) skewY(-31deg);
}
li:nth-child(3) {
  -webkit-transform: rotate(120deg) skewY(-31deg);
  -ms-transform: rotate(120deg) skewY(-31deg);
  transform: rotate(120deg) skewY(-31deg);
}
li:nth-child(4) {
  -webkit-transform: rotate(180deg) skewY(-31deg);
  -ms-transform: rotate(180deg) skewY(-31deg);
  transform: rotate(180deg) skewY(-31deg);
}
li:nth-child(5) {
  -webkit-transform: rotate(240deg) skewY(-31deg);
  -ms-transform: rotate(240deg) skewY(-31deg);
  transform: rotate(240deg) skewY(-31deg);
}
li:nth-child(6) {
  -webkit-transform: rotate(300deg) skewY(-31deg);
  -ms-transform: rotate(300deg) skewY(-31deg);
  transform: rotate(300deg) skewY(-31deg);
}

li:first-child .text {
  background: #f2709c;
}
li:nth-child(2) .text {
  background: #ff9472;
}
li:nth-child(3) .text {
  background: #f2709c;
}
li:nth-child(4) .text {
  background: #ff9472;
}
li:nth-child(5) .text {
  background: #f2709c;
}
li:nth-child(6) .text {
  background: #ff9472;
}
</style>
