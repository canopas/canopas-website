<template>
  <div
    class="tw-relative tw-w-full tw-mx-auto tw-py-0 tw-px-[0.75rem] lg:tw-py-[50px]"
  >
    <img
      :src="backgroundImage"
      class="tw-w-full tw-h-full tw-top-0 tw-left-0 tw-absolute tw-object-cover -tw-z-[5]"
      loading="lazy"
      alt="double-quote-mark"
    />
    <div class="tw-container">
      <div>
        <div class="tw-flex tw-flex-col lg:tw-flex-row">
          <div class="lg:tw-flex-[35%]">
            <div
              class="v2-header-3-text tw-text-center tw-p-[10%] lg:tw-text-start lg:tw-pt-[35%] lg:tw-pr-[20%] lg:tw-pb-[25%] lg:tw-pl-[10%] xl:tw-py-[20%] xl:tw-pr-[25%] xl:tw-pl-[8%]"
            >
              Success
              <span class="tw-text-white border-text">stories</span>
              of people like you.
            </div>
            <div
              v-for="client in clients.slice(0, 2)"
              :key="client"
              class="tw-border-transparent tw-p-[30px] tw-mt-0 tw-mx-4 tw-mb-12 tw-rounded-[3rem] tw-shadow-[0_4px_4px_rgba(0,0,0,0.5)] md:tw-p-[35px] md:tw-mt-0 md:tw-mx-4 md:tw-mb-12 xl:tw-mt-0 xl:tw-mx-[2rem] xl:tw-mb-12"
              :class="[client.bgColor]"
            >
              <div class="v2-normal-3-text" :class="[client.fontColor]">
                {{ client.review }}
              </div>
              <div
                class="tw-flex tw-flex-wrap tw-items-center tw-flex-col sm:tw-flex-row tw-mt-6 v2-title-3-text"
              >
                <img :src="client.image" alt="client" loading="lazy" />
                <div class="canopas-gradient-text">
                  {{ client.name }}
                </div>
              </div>
            </div>
          </div>
          <div class="lg:tw-flex-[35%]">
            <div
              v-for="client in clients.slice(2, 5)"
              :key="client"
              class="tw-border-transparent tw-p-[30px] tw-mt-0 tw-mx-4 tw-mb-12 tw-rounded-[3rem] tw-shadow-[0_4px_4px_rgba(0,0,0,0.5)] md:tw-p-[35px] md:tw-mt-0 md:tw-mx-4 md:tw-mb-12 xl:tw-mt-0 xl:tw-mx-[2rem] xl:tw-mb-12"
              :class="[client.bgColor]"
            >
              <div class="v2-normal-3-text" :class="[client.fontColor]">
                {{ client.review }}
              </div>
              <div
                class="tw-flex tw-flex-wrap tw-items-center tw-flex-col sm:tw-flex-row tw-mt-6 v2-title-3-text"
              >
                <img :src="client.image" alt="client" loading="lazy" />

                <div class="canopas-gradient-text">
                  {{ client.name }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import backgroundImage from "@/assets/images/clients/v2-client/quote-mark.svg";
import marcus from "@/assets/images/clients/v2-client/marcus.png";
import jake from "@/assets/images/clients/v2-client/jake.png";
import lisa from "@/assets/images/clients/v2-client/lisa.png";
import maor from "@/assets/images/clients/v2-client/maor.png";
import ramasis from "@/assets/images/clients/v2-client/ramasis.png";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      clients: [
        {
          id: "client-1",
          image: marcus,
          name: "Marcus L.",
          fontColor: "tw-text-white",
          bgColor: "tw-bg-black-900",
          review:
            "Canopas has been nothing but wonderful on this project. His\
              communication and ability to advise on the best solutions\
              throughout the project were top-notch. I have very much\
              enjoyed working with Darpan and his team and would surely\
              recommend them to anyone looking for a solid developer.",
        },
        {
          id: "client-2",
          image: jake,
          name: "Jake N.",
          fontColor: "tw-text-white",
          bgColor: "tw-bg-black-900",
          review:
            "Canopas team was also incredibly kind and always willing\
              tosolve any problem through research something that we highly\
              appreciate. Many people do not take the time to understand a\
              problem but Canopas did it every time and found a way to fix\
              it. They even thought about the User Experience and the\
              design which not many engineers think about. ",
        },
        {
          id: "client-3",
          image: lisa,
          name: "Lisa W.",
          fontColor: "tw-text-black-900",
          bgColor:
            "tw-from-[#f98386]/10 tw-to-[#f98386]/10 tw-bg-gradient-247 ",
          review:
            "There is not enough space to say all the wonderful things I\
              would want to share about Canopas. They are incredibly helpful,\
              stay calm even when we had to deal with tough issues on our\
              app, and always found a way to help us fix whatever was needed\
              or roll out any new features for our app in both the iOS and\
              Android stores. I will absolutely find a way to work with\
              Canopas again!",
        },
        {
          id: "client-4",
          image: maor,
          name: "Maor T.",
          fontColor: "tw-text-white",
          bgColor: "tw-bg-black-900",
          review:
            "This is our favorite expert for all mobile and web developing\
              areas. Darpan has a great team who get the job done above our\
              expectations. We would like to hire Canopas for more projects in\
              the future and highly recommend their kind and professional\
              services.",
        },
        {
          id: "client-5",
          image: ramasis,
          name: "Ramsis A.",
          fontColor: "tw-text-white",
          bgColor: "tw-bg-black-900",
          review:
            "Canopas has been great to work with. From day 1, they made\
              sure they understood exactly what I wanted and advised and\
              guided me during the process (with patience). This project\
              has ended, but I am sure to work with them again on other\
              projects. I would recommend their services. ",
        },
      ],
    };
  },
};
</script>
