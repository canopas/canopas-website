<template>
  <section class="relative mt-16 lg:mt-60">
    <p
      class="container text-center mobile-header-2 lg:desk-header-2 text-black-87 2xl:max-w-[67%] xll:max-w-[45%] 3xl:max-w-[28%] 2xl:mx-auto"
    >
      Why Canopas?
    </p>
    <p class="container text-center mt-2 lg:mt-6 text-black-4 lg:text-2xl">
      Discover what makes us unique through the experiences of our satisfied
      customers.
    </p>
    <div class="absolute w-full h-full lg:hidden">
      <img
        :src="benefitbgImages[0]"
        :srcset="`${benefitbgImages[0]} 400w`"
        loading="lazy"
        class="object-cover w-full h-4/5"
        alt="background-image"
      />
    </div>
    <div
      class="lg:hidden container flex mt-20"
      v-for="(benefit, index) in benefits"
      :key="index"
    >
      <div
        class="-mt-10 bg-clip-text text-transparent from-orange-2 to-pink-2 bg-gradient-[270.11deg] desk-header-1"
      >
        {{ benefit.id }}
      </div>
      <div>
        <div class="mobile-header-3-semibold text-black-87">
          {{ benefit.title }}
        </div>
        <div class="mt-2 pr-2 sub-h3-regular text-black-60">
          {{ benefit.description }}
        </div>
      </div>
    </div>
    <div class="lg:mt-16">
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
          class="flex gap-4 w-[75%] relative"
          :class="[
            index == 2 || index == 4 ? 'container' : '',
            index == 4 ? 'max-w-[67%]' : '',
          ]"
        >
          <div class="desk-header-2">
            {{ benefit.number }}
          </div>

          <div>
            <div class="desk-header-3 text-black-87">
              {{ benefit.title }}
            </div>
            <div
              class="xl:w-[80%] mt-2 mobile-header-3-regular text-black-60"
              :class="index == 4 ? 'w-full xl:w-[84%]' : ''"
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
              class="absolute border-orange-300"
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
            "top-[4.5rem] left-6  w-[27%] xl:w-[28%] 2xl:w-[28%]  h-[13.688rem] xl:h-[13.813rem] 2xl:h-[13.75rem] border-l-[3px] border-b-[3px] ",
          arrowClass:
            "top-[12.9rem] xl:top-[13.1rem] 2xl:top-52 left-44 xl:left-56  2xl:left-[16.5rem] w-0 h-0 border-t-[10px] border-t-transparent border-l-[20px] border-b-[10px] border-b-transparent",
        },
        {
          id: 2,
          number: "02.",
          title: `Quick response time`,
          description: `We pride ourselves on our ability to respond quickly to our clients' inquiries, concerns, 
or complaints. Our team is dedicated to providing exceptional customer service, 
and we always aim to resolve any issues promptly.`,
          className:
            "top-16 left-8 xl:left-[1.4rem] w-[20%] h-[8.188rem] xl:h-[8.063rem] 2xl:h-[6.188rem]   border-l-[3px]",
          secondBorder:
            "top-[4.5rem] xl:top-[3.8rem] 2xl:top-[1.3rem] left-[-3.9rem] 2xl:-left-20 xl:left-[-4.5rem]  h-[7.5rem] xl:h-[8.438rem] 2xl:h-[9.688rem] border-l-[3px] -rotate-90",
          thirdBorder:
            "top-[8.2rem] xl:top-[7.9rem] 2xl:top-[6.05rem] left-[-7.6rem] xl:left-[-8.7rem] 2xl:left-[-9.9rem] h-14 xl:h-[4.063rem] 2xl:h-[3.688rem]  border-l-[3px] ",
          arrowClass:
            "top-[11.5rem] 2xl:top-[9.5rem] left-[-8.1rem] xl:left-[-9.2rem] 2xl:left-[-10.4rem] w-0 h-0  border-t-[15px] border-l-transparent border-l-[10px] border-r-[10px] border-r-transparent",
        },
        {
          id: 3,
          number: "03.",
          title: `Going above and beyond`,
          description: `At our core, we are dedicated to surpassing our clients' wildest expectations. We take pride in always pushing the envelope and finding new, innovative ways to provide an unforgettable experience. 
Our team is known to surprise and delight clients with personalized suggestions and recommendations that will leave your customers overjoyed.`,
          mainClass: "pl-0",
          className:
            "top-14 left-10  w-[8%] xl:w-[11%] 2xl:w-[10%] h-[18.375rem] 2xl:h-[16.625rem] border-l-[3px] border-b-[3px] ",
          arrowClass:
            "top-[17.7rem] 2xl:top-[15.9rem] left-[3.2rem] xl:left-20 2xl:left-[5.9rem] w-0 h-0  border-t-[10px] border-t-transparent border-l-[20px] border-b-[10px] border-b-transparent",
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
            "top-14 left-[1.6rem] xl:left-10 w-[20%] xl:left-6  h-[8.188rem] 2xl:h-[8.375rem]  border-l-[3px]",
          secondBorder:
            "top-[5.9rem] xl:top-[5.4rem] 2xl:top-[5.1rem]  -left-10  xl:left-[-2.9rem] 2xl:left-[-3.4rem] h-[4.438rem] xl:h-[5.5rem] 2xl:h-[6.563rem] border-l-[3px]  -rotate-90",
          thirdBorder:
            "top-32 xl:top-[8.1rem] 2xl:top-[8.3rem] left-[-4.7rem] xl:left-[-5.6rem] 2xl:left-[-6.6rem] h-[7.938rem] xl:h-[9.813rem] 2xl:h-[5.563rem] border-l-[3px] ",
          arrowClass:
            "top-[15.5rem] xl:top-[17.5rem] 2xl:top-[13.5rem] left-[-5.2rem] xl:left-[-6.1rem] 2xl:left-[-7.2rem] w-0 h-0  border-t-[15px] border-l-transparent border-l-[10px] border-r-[10px] border-r-transparent",
        },
        {
          id: 5,
          number: "05.",
          title: `Transparency and honesty`,
          description: `We believe in building trust with our clients. That's why we are transparent and honest about our products or services. 
You can expect clear and upfront pricing, honest product descriptions, and easy-to understand terms and conditions.`,
          className:
            "top-14 left-[2.2rem] w-[20%] h-[8.188rem]   border-l-[3px]",
          secondBorder:
            "top-[2.95rem] xl:top-[1.9rem] 2xl:top-[0.9rem]  left-[-5.4rem] xl:left-[-6.4rem] 2xl:left-[-7.335rem] h-[10.375rem] xl:h-[12.438rem] 2xl:h-[14.438rem] border-l-[3px]  -rotate-90",
          thirdBorder:
            "top-32 xl:top-[8.1rem] 2xl:top-[8.1rem] left-[-10.5rem] xl:left-[-12.5rem] 2xl:left-[-14.5rem] h-[6.438rem] xl:h-[5.875rem] 2xl:h-[4.188rem] border-l-[3px]",
          arrowClass:
            " top-56 xl:top-[13.5rem] 2xl:top-[11.5rem] -left-44 xl:-left-52 2xl:-left-60 w-0 h-0 border-t-[15px] border-l-transparent border-l-[10px] border-r-[10px] border-r-transparent",
        },
        {
          id: 6,
          number: "06.",
          title: `Flexibility and adaptability`,
          description: `We understand that businesses have unique needs and requirements. That's why we offer flexible and adaptable solutions to meet your specific needs. We are always willing to work with our clients to find the best solution for their business.`,
          className:
            "top-14 left-[1.6rem] w-[28%] 2xl:w-[29%] h-[14.813rem] border-l-[3px] border-b-[3px] ",
          arrowClass:
            "top-[14.1rem] 2xl:top-[14.1rem] left-44 xl:left-[13.5rem] 2xl:left-[16.5rem] w-0 h-0  border-t-[10px] border-t-transparent border-l-[20px] border-b-[10px] border-b-transparent",
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
