<template>
  <section
    class="tw-relative tw-mt-[30px] xl:tw-mt-[100px] 2xl:tw-mt-[130px] tw-mb-[150px]"
  >
    <h1
      class="tw-mt-8 tw-mb-8 v2-canopas-gradient-text tw-text-center tw-text-[2.5rem] md:tw-text-[3.5rem] lg:tw-text-[5rem] tw-leading-[3rem] md:tw-leading-[4rem] lg:tw-leading-[5.5rem] tw-font-inter-bold"
    >
      WHY US?
    </h1>
    <p
      class="tw-container tw-font-inter-medium tw-text-black-core/[0.87] tw-text-center tw-text-[1.563rem] md:tw-text-[2.188rem] lg:tw-text-[3rem] tw-leading-[2.188rem] md:tw-leading-[2.813rem] lg:tw-leading-[4.063rem]"
    >
      Here are some of the things you will get by working with us that make us
      special according to our past customers.
    </p>
    <div class="tw-absolute tw-w-full tw-h-full lg:tw-hidden">
      <img
        :src="benefitbgImages[0]"
        :srcset="`${benefitbgImages[0]} 400w`"
        loading="lazy"
        class="tw-object-cover tw-w-full tw-h-[80%]"
        alt="background-image"
      />
    </div>
    <div
      class="lg:tw-hidden tw-container tw-flex tw-flex-row tw-mt-20"
      v-for="(benefit, index) in benefits"
      :key="index"
    >
      <div
        class="tw-mt-[-50px] tw-bg-clip-text tw-text-transparent tw-from-[#FF835B]/[0.2] tw-to-[#F2709C]/[0.2] tw-bg-gradient-[270.11deg] tw-font-inter-bold tw-text-[6.25rem] tw-leading-[8.125rem]"
      >
        {{ benefit.id }}
      </div>
      <div class="tw-ml-[-20px]">
        <div
          class="tw-font-inter-semibold tw-text-[1.5rem] md:tw-text-[2rem] tw-leading-[2.25rem] md:tw-leading-[2.813rem] tw-text-black-core/[0.87]"
        >
          {{ benefit.title }}
        </div>
        <div
          class="tw-mt-2 tw-pr-2 tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[1rem] md:tw-text-[1.125rem] tw-leading-[1.5rem] md:tw-leading-[1.625rem]"
        >
          {{ benefit.description }}
        </div>
      </div>
    </div>
    <div class="lg:tw-mt-[8rem]">
      <div
        class="xll:tw-container tw-absolute 2xl:tw-inset-0 tw-hidden lg:tw-block"
      >
        <img
          :src="width > 2440 ? benefitbg2Images[2] : benefitbg2Images[0]"
          :srcset="
            width > 2440
              ? `${benefitbg2Images[2]} 1400w`
              : `${benefitbg2Images[0]} 400w, ${benefitbg2Images[1]} 800w, ${benefitbg2Images[3]} 1800w, ${benefitbg2Images[4]} 2400w`
          "
          loading="lazy"
          class="tw-w-full tw-h-full xll:tw-object-contain"
          alt="background-image"
        />
      </div>
      <div
        class="tw-container tw-hidden lg:tw-flex tw-flex-col tw-mt-20"
        v-for="(benefit, index) in benefits"
        :key="index"
        :class="[
          index == 0 || index == 2 || index == 4 || index == 5
            ? 'tw-items-start'
            : 'tw-items-end',
          benefit.mainClass,
        ]"
      >
        <div
          class="tw-flex tw-flex-row tw-gap-4 tw-w-[75%] tw-relative"
          :class="[
            index == 2 || index == 4 ? 'tw-container' : '',
            index == 4 ? 'tw-max-w-[67%]' : '',
          ]"
        >
          <div
            class="tw-font-inter-semibold tw-text-[3.75rem] tw-leading-[5.625rem] tw-text-black-core/[0.87] tw-mt-[-10px]"
          >
            {{ benefit.number }}
          </div>

          <div class="">
            <div
              class="tw-font-inter-semibold tw-text-[2.5rem] tw-leading-[3.75rem] tw-text-black-core/[0.87]"
            >
              {{ benefit.title }}
            </div>
            <div
              class="xl:tw-w-[80%] tw-mt-2 tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[1.25rem] tw-leading-[1.875rem]"
              :class="index == 4 ? 'tw-w-[100%] xl:tw-w-[84%]' : ''"
            >
              {{ benefit.description }}
            </div>
          </div>
          <div class="tw-absolute gradient-border" :class="benefit.className">
            <div
              class="tw-absolute gradient-border"
              :class="benefit.secondBorder"
            ></div>
            <div
              class="tw-absolute gradient-border"
              :class="benefit.thirdBorder"
            ></div>
            <div
              class="tw-absolute tw-border-[#FF835B]"
              :class="benefit.arrowClass"
            ></div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import benefitbg400w from "@/assets/images/client-thankyou/benefit-bg-400w.webp";
import benefitbg800w from "@/assets/images/client-thankyou/benefit-bg-800w.webp";
import benefit2bg400w from "@/assets/images/client-thankyou/benefit-bg-2-400w.webp";
import benefit2bg800w from "@/assets/images/client-thankyou/benefit-bg-2-800w.webp";
import benefit2bg1400w from "@/assets/images/client-thankyou/benefit-bg-2-1400w.webp";
import benefit2bg1800w from "@/assets/images/client-thankyou/benefit-bg-2-1800w.webp";
import benefit2bg2400w from "@/assets/images/client-thankyou/benefit-bg-2-2400w.webp";
export default {
  data() {
    return {
      width: 0,
      benefitbgImages: [benefitbg400w, benefitbg800w],
      benefitbg2Images: [
        benefit2bg400w,
        benefit2bg800w,
        benefit2bg1400w,
        benefit2bg1800w,
        benefit2bg2400w,
      ],
      benefits: [
        {
          id: 1,
          number: "01.",
          title: `Personalized attention`,
          description: `You will notice this in our intro call where we will learn more about your business, idea, goals, vision, and customers. Our team of experts will take the time to understand your specific needs and work with you to create 
a customized plan of action.`,
          className:
            "tw-top-[4.5rem] tw-left-10  tw-w-[25%] xl:tw-w-[26%] 2xl:tw-w-[27%]  tw-h-[234px] xl:tw-h-[263px] 2xl:tw-h-[235px] tw-border-l-[3px] tw-border-b-[3px] ",
          arrowClass:
            "tw-top-[13.9rem] xl:tw-top-[15.7rem] 2xl:tw-top-[14rem] tw-left-[10.5rem] xl:tw-left-[13rem]  2xl:tw-left-[16rem] tw-w-0 tw-h-0 tw-border-t-[10px] tw-border-t-transparent tw-border-l-[20px] tw-border-b-[10px] tw-border-b-transparent",
        },
        {
          id: 2,
          number: "02.",
          title: `Quick response time`,
          description: `We pride ourselves on our ability to respond quickly to our clients' inquiries, concerns, 
or complaints. Our team is dedicated to providing exceptional customer service, 
and we always aim to resolve any issues promptly.`,
          className:
            "tw-top-[4.5rem] tw-left-[2rem]  tw-w-[20%] tw-h-[131px] xl:tw-h-[130px] 2xl:tw-h-[99px]   tw-border-l-[3px]",
          secondBorder:
            "tw-top-[4.6rem] xl:tw-top-[3.8rem] 2xl:tw-top-[1.3rem] tw-left-[-3.7rem] 2xl:tw-left-[-5rem] xl:tw-left-[-4.5rem]  tw-h-[112px] xl:tw-h-[135px] 2xl:tw-h-[155px] tw-border-l-[3px]   tw--rotate-90",
          thirdBorder:
            "tw-top-[8.1rem] xl:tw-top-[7.9rem] 2xl:tw-top-[6.05rem] tw-left-[-7.1rem] xl:tw-left-[-8.7rem] 2xl:tw-left-[-9.9rem] tw-h-[66px] xl:tw-h-[73px] 2xl:tw-h-[99px]  tw-border-l-[3px] ",
          arrowClass:
            "tw-top-[11.5rem] xl:tw-top-[11.7rem] 2xl:tw-top-[12rem] tw-left-[-7.6rem] xl:tw-left-[-9.2rem] 2xl:tw-left-[-10.4rem] tw-w-0 tw-h-0  tw-border-t-[15px] tw-border-l-transparent tw-border-l-[10px] tw-border-r-[10px] tw-border-r-transparent",
        },
        {
          id: 3,
          number: "03.",
          title: `Going above and beyond`,
          description: `At our core, we are dedicated to surpassing our clients' wildest expectations. We take pride in always pushing the envelope and finding new, innovative ways to provide an unforgettable experience. 
Our team is known to surprise and delight clients with personalized suggestions and recommendations that will leave your customers overjoyed.`,
          mainClass: "tw-pl-[0%]",
          className:
            "tw-top-[4.5rem] tw-left-[3rem]  tw-w-[8%] xl:tw-w-[10%] 2xl:tw-w-[11%] tw-h-[294px] 2xl:tw-h-[266px] tw-border-l-[3px] tw-border-b-[3px] ",
          arrowClass:
            "tw-top-[17.7rem] xl:tw-top-[17.6rem] 2xl:tw-top-[15.9rem] tw-left-[3.2rem] xl:tw-left-[4.2rem] 2xl:tw-left-[6.2rem] tw-w-0 tw-h-0  tw-border-t-[10px] tw-border-t-transparent tw-border-l-[20px] tw-border-b-[10px] tw-border-b-transparent",
        },
        {
          id: 4,
          number: "04.",
          title: `Seamless omnichannel experiences`,
          description: `We understand that customers interact with businesses through multiple channels. 
That's why we offer seamless omnichannel experiences, so you can interact with us through your preferred channel, whether 
it's email, phone, or chat. 
For the sake of simplicity, we prefer using one channel once the contract is established for ease of access and project management.`,
          className:
            "tw-top-[4.5rem] tw-left-[2.4rem] xl:tw-left-[2.5rem] tw-w-[20%]  tw-h-[131px] 2xl:tw-h-[134px]  tw-border-l-[3px]",
          secondBorder:
            "tw-top-[5.9rem] xl:tw-top-[5.4rem] 2xl:tw-top-[5.1rem]  tw-left-[-2.5rem]  xl:tw-left-[-2.9rem] 2xl:tw-left-[-3.4rem] tw-h-[71px] xl:tw-h-[88px] 2xl:tw-h-[105px] tw-border-l-[3px]  tw--rotate-90",
          thirdBorder:
            "tw-top-[8rem] xl:tw-top-[8.1rem] 2xl:tw-top-[8.3rem] tw-left-[-4.7rem] xl:tw-left-[-5.6rem] 2xl:tw-left-[-6.6rem] tw-h-[220px] xl:tw-h-[157px] 2xl:tw-h-[125px] tw-border-l-[3px] ",
          arrowClass:
            "tw-top-[21.5rem] xl:tw-top-[17.5rem] 2xl:tw-top-[15.5rem] tw-left-[-5.2rem] xl:tw-left-[-6.1rem] 2xl:tw-left-[-7.2rem] tw-w-0 tw-h-0  tw-border-t-[15px] tw-border-l-transparent tw-border-l-[10px] tw-border-r-[10px] tw-border-r-transparent",
        },
        {
          id: 5,
          number: "05.",
          title: `Transparency and honesty`,
          description: `We believe in building trust with our clients. That's why we are transparent and honest about our products or services. 
You can expect clear and upfront pricing, honest product descriptions, and easy-to understand terms and conditions.`,
          className:
            "tw-top-[4.5rem] tw-left-[3rem] tw-w-[20%] tw-h-[131px]   tw-border-l-[3px]",
          secondBorder:
            "tw-top-[2.95rem] xl:tw-top-[1.9rem] 2xl:tw-top-[0.9rem]  tw-left-[-5.4rem] xl:tw-left-[-6.4rem] 2xl:tw-left-[-7.335rem] tw-h-[166px] xl:tw-h-[199px] 2xl:tw-h-[231px] tw-border-l-[3px]  tw--rotate-90",
          thirdBorder:
            "tw-top-[8rem] xl:tw-top-[8.1rem] 2xl:tw-top-[8.1rem] tw-left-[-10.5rem] xl:tw-left-[-12.5rem] 2xl:tw-left-[-14.5rem] tw-h-[162px] xl:tw-h-[94px] 2xl:tw-h-[67px] tw-border-l-[3px]",
          arrowClass:
            " tw-top-[17.5rem] xl:tw-top-[13.5rem] 2xl:tw-top-[11.5rem] tw-left-[-11rem] xl:tw-left-[-13rem] 2xl:tw-left-[-15rem] tw-w-0 tw-h-0 tw-border-t-[15px] tw-border-l-transparent tw-border-l-[10px] tw-border-r-[10px] tw-border-r-transparent",
        },
        {
          id: 6,
          number: "06.",
          title: `Flexibility and adaptability`,
          description: `We understand that businesses have unique needs and requirements. That's why we offer flexible and adaptable solutions to meet your specific needs. We are always willing to work with our clients to find the best solution for their business.`,
          className:
            "tw-top-[4.5rem] tw-left-10 tw-w-[25%] 2xl:tw-w-[27%] tw-h-[262px] 2xl:tw-h-[237px] tw-border-l-[3px] tw-border-b-[3px] ",
          arrowClass:
            "tw-top-[15.7rem] 2xl:tw-top-[14.1rem] tw-left-[10.5rem] xl:tw-left-[12.5rem] 2xl:tw-left-[15.5rem] tw-w-0 tw-h-0  tw-border-t-[10px] tw-border-t-transparent tw-border-l-[20px] tw-border-b-[10px] tw-border-b-transparent",
        },
        {
          id: 7,
          number: "07.",
          title: `Giving back to the community`,
          description: `We believe in giving back to the community in whatever way we can. As Tony Robbins says "Giving back to society will give you a sense of fulfillment you just can’t get anywhere else." We took a chance and tried it once, and now it's been an integral part of our operations with no regrets. For the tech community, we do open-source contributions to help people learn new things based on our experiences. For others, we do the best we can.`,
        },
      ],
    };
  },
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>

<style scoped>
.gradient-border {
  border-image: linear-gradient(to right, #f2709c, #ff835b) 1;
}
</style>
