<template>
  <section class="relative mt-[30px] xl:mt-[100px] 2xl:mt-[130px] mb-[150px]">
    <h1
      class="mt-8 mb-8 v2-canopas-gradient-text text-center text-[2.5rem] md:text-[3.5rem] lg:text-[5rem] leading-[3rem] md:leading-[4rem] lg:leading-[5.5rem] font-inter-bold"
    >
      WHY US?
    </h1>
    <p
      class="container font-inter-medium text-black-core/[0.87] text-center text-[1.563rem] md:text-[2.188rem] lg:text-[3rem] leading-[2.188rem] md:leading-[2.813rem] lg:leading-[4.063rem]"
    >
      Here are some of the things you will get by working with us that make us
      special according to our past customers.
    </p>
    <div class="absolute w-full h-full lg:hidden">
      <img
        :src="benefitbgImages[0]"
        :srcset="`${benefitbgImages[0]} 400w`"
        loading="lazy"
        class="object-cover w-full h-[80%]"
        alt="background-image"
      />
    </div>
    <div
      class="lg:hidden container flex flex-row mt-20"
      v-for="(benefit, index) in benefits"
      :key="index"
    >
      <div
        class="mt-[-50px] bg-clip-text text-transparent from-[#FF835B]/[0.2] to-[#F2709C]/[0.2] bg-gradient-[270.11deg] font-inter-bold text-[6.25rem] leading-[8.125rem]"
      >
        {{ benefit.id }}
      </div>
      <div class="-ml-5">
        <div
          class="font-inter-semibold text-[1.5rem] md:text-[2rem] leading-9 md:leading-[2.813rem] text-black-core/[0.87]"
        >
          {{ benefit.title }}
        </div>
        <div
          class="mt-2 pr-2 font-inter-regular text-black-core/[0.87] text-[1rem] md:text-[1.125rem] leading-6 md:leading-[1.625rem]"
        >
          {{ benefit.description }}
        </div>
      </div>
    </div>
    <div class="lg:mt-32">
      <div class="xll:container absolute 2xl:inset-0 hidden lg:block">
        <img
          :src="width > 2440 ? benefitbg2Images[2] : benefitbg2Images[0]"
          :srcset="
            width > 2440
              ? `${benefitbg2Images[2]} 1400w`
              : `${benefitbg2Images[0]} 400w, ${benefitbg2Images[1]} 800w, ${benefitbg2Images[3]} 1800w, ${benefitbg2Images[4]} 2400w`
          "
          loading="lazy"
          class="w-full h-full xll:object-contain"
          alt="background-image"
        />
      </div>
      <div
        class="container hidden lg:flex flex-col mt-20"
        v-for="(benefit, index) in benefits"
        :key="index"
        :class="[
          index == 0 || index == 2 || index == 4 || index == 5
            ? 'items-start'
            : 'items-end',
          benefit.mainClass,
        ]"
      >
        <div
          class="flex flex-row gap-4 w-[75%] relative"
          :class="[
            index == 2 || index == 4 ? 'container' : '',
            index == 4 ? 'max-w-[67%]' : '',
          ]"
        >
          <div
            class="font-inter-semibold text-[3.75rem] leading-[5.625rem] text-black-core/[0.87] -mt-2.5"
          >
            {{ benefit.number }}
          </div>

          <div class="">
            <div
              class="font-inter-semibold text-[2.5rem] leading-[3.75rem] text-black-core/[0.87]"
            >
              {{ benefit.title }}
            </div>
            <div
              class="xl:w-[80%] mt-2 font-inter-regular text-black-core/[0.87] text-[1.25rem] leading-[1.875rem]"
              :class="index == 4 ? 'w-[100%] xl:w-[84%]' : ''"
            >
              {{ benefit.description }}
            </div>
          </div>
          <div class="absolute gradient-border" :class="benefit.className">
            <div
              class="absolute gradient-border"
              :class="benefit.secondBorder"
            ></div>
            <div
              class="absolute gradient-border"
              :class="benefit.thirdBorder"
            ></div>
            <div
              class="absolute border-[#FF835B]"
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
            "top-[4.5rem] left-10  w-[25%] xl:w-[26%] 2xl:w-[27%]  h-[234px] xl:h-[263px] 2xl:h-[235px] border-l-[3px] border-b-[3px] ",
          arrowClass:
            "top-[13.9rem] xl:top-[15.7rem] 2xl:top-56 left-[10.5rem] xl:left-52  2xl:left-64 w-0 h-0 border-t-[10px] border-t-transparent border-l-[20px] border-b-[10px] border-b-transparent",
        },
        {
          id: 2,
          number: "02.",
          title: `Quick response time`,
          description: `We pride ourselves on our ability to respond quickly to our clients' inquiries, concerns, 
or complaints. Our team is dedicated to providing exceptional customer service, 
and we always aim to resolve any issues promptly.`,
          className:
            "top-[4.5rem] left-8  w-[20%] h-[131px] xl:h-[130px] 2xl:h-[99px]   border-l-[3px]",
          secondBorder:
            "top-[4.6rem] xl:top-[3.8rem] 2xl:top-[1.3rem] left-[-3.7rem] 2xl:-left-20 xl:left-[-4.5rem]  h-[112px] xl:h-[135px] 2xl:h-[155px] border-l-[3px]   -rotate-90",
          thirdBorder:
            "top-[8.1rem] xl:top-[7.9rem] 2xl:top-[6.05rem] left-[-7.1rem] xl:left-[-8.7rem] 2xl:left-[-9.9rem] h-[66px] xl:h-[73px] 2xl:h-[99px]  border-l-[3px] ",
          arrowClass:
            "top-[11.5rem] xl:top-[11.7rem] 2xl:top-48 left-[-7.6rem] xl:left-[-9.2rem] 2xl:left-[-10.4rem] w-0 h-0  border-t-[15px] border-l-transparent border-l-[10px] border-r-[10px] border-r-transparent",
        },
        {
          id: 3,
          number: "03.",
          title: `Going above and beyond`,
          description: `At our core, we are dedicated to surpassing our clients' wildest expectations. We take pride in always pushing the envelope and finding new, innovative ways to provide an unforgettable experience. 
Our team is known to surprise and delight clients with personalized suggestions and recommendations that will leave your customers overjoyed.`,
          mainClass: "pl-[0%]",
          className:
            "top-[4.5rem] left-[3rem]  w-[8%] xl:w-[10%] 2xl:w-[11%] h-[294px] 2xl:h-[266px] border-l-[3px] border-b-[3px] ",
          arrowClass:
            "top-[17.7rem] xl:top-[17.6rem] 2xl:top-[15.9rem] left-[3.2rem] xl:left-[4.2rem] 2xl:left-[6.2rem] w-0 h-0  border-t-[10px] border-t-transparent border-l-[20px] border-b-[10px] border-b-transparent",
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
            "top-[4.5rem] left-[2.4rem] xl:left-10 w-[20%]  h-[131px] 2xl:h-[134px]  border-l-[3px]",
          secondBorder:
            "top-[5.9rem] xl:top-[5.4rem] 2xl:top-[5.1rem]  -left-10  xl:left-[-2.9rem] 2xl:left-[-3.4rem] h-[71px] xl:h-[88px] 2xl:h-[105px] border-l-[3px]  -rotate-90",
          thirdBorder:
            "top-32 xl:top-[8.1rem] 2xl:top-[8.3rem] left-[-4.7rem] xl:left-[-5.6rem] 2xl:left-[-6.6rem] h-[220px] xl:h-[157px] 2xl:h-[125px] border-l-[3px] ",
          arrowClass:
            "top-[21.5rem] xl:top-[17.5rem] 2xl:top-[15.5rem] left-[-5.2rem] xl:left-[-6.1rem] 2xl:left-[-7.2rem] w-0 h-0  border-t-[15px] border-l-transparent border-l-[10px] border-r-[10px] border-r-transparent",
        },
        {
          id: 5,
          number: "05.",
          title: `Transparency and honesty`,
          description: `We believe in building trust with our clients. That's why we are transparent and honest about our products or services. 
You can expect clear and upfront pricing, honest product descriptions, and easy-to understand terms and conditions.`,
          className: "top-[4.5rem] left-12 w-[20%] h-[131px]   border-l-[3px]",
          secondBorder:
            "top-[2.95rem] xl:top-[1.9rem] 2xl:top-[0.9rem]  left-[-5.4rem] xl:left-[-6.4rem] 2xl:left-[-7.335rem] h-[166px] xl:h-[199px] 2xl:h-[231px] border-l-[3px]  -rotate-90",
          thirdBorder:
            "top-32 xl:top-[8.1rem] 2xl:top-[8.1rem] left-[-10.5rem] xl:left-[-12.5rem] 2xl:left-[-14.5rem] h-[162px] xl:h-[94px] 2xl:h-[67px] border-l-[3px]",
          arrowClass:
            " top-[17.5rem] xl:top-[13.5rem] 2xl:top-[11.5rem] -left-44 xl:-left-52 2xl:left-60 w-0 h-0 border-t-[15px] border-l-transparent border-l-[10px] border-r-[10px] border-r-transparent",
        },
        {
          id: 6,
          number: "06.",
          title: `Flexibility and adaptability`,
          description: `We understand that businesses have unique needs and requirements. That's why we offer flexible and adaptable solutions to meet your specific needs. We are always willing to work with our clients to find the best solution for their business.`,
          className:
            "top-[4.5rem] left-10 w-[25%] 2xl:w-[27%] h-[262px] 2xl:h-[237px] border-l-[3px] border-b-[3px] ",
          arrowClass:
            "top-[15.7rem] 2xl:top-[14.1rem] left-[10.5rem] xl:left-[12.5rem] 2xl:left-[15.5rem] w-0 h-0  border-t-[10px] border-t-transparent border-l-[20px] border-b-[10px] border-b-transparent",
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
