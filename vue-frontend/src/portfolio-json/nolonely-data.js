import landingBackground400w from "@/assets/images/portfolio/nolonely/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/nolonely/landing/background-800w.webp";
import landingBackground1400w from "@/assets/images/portfolio/nolonely/landing/background-1400w.webp";
import landingBackground2400w from "@/assets/images/portfolio/nolonely/landing/background-2400w.webp";

import videoBackgroundImage from "@/assets/images/portfolio/nolonely/swiper/mobile.webp";
import cooking from "@/assets/images/portfolio/nolonely/swiper/cooking.svg";
import meditation from "@/assets/images/portfolio/nolonely/swiper/meditation.svg";
import painting from "@/assets/images/portfolio/nolonely/swiper/painting.svg";
import pet from "@/assets/images/portfolio/nolonely/swiper/pet.svg";
import singing from "@/assets/images/portfolio/nolonely/swiper/singing.svg";

import brandingBackground400w from "@/assets/images/portfolio/nolonely/branding/background-400w.webp";
import brandingBackground800w from "@/assets/images/portfolio/nolonely/branding/background-800w.webp";
import brandingBackground1400w from "@/assets/images/portfolio/nolonely/branding/background-1400w.webp";
import brandingBackground2400w from "@/assets/images/portfolio/nolonely/branding/background-2400w.webp";

import branding1Image400w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_1-400w.webp";
import branding1Image800w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_1-800w.webp";
import branding1Image1200w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_1-1200w.webp";
import branding1Image1600w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_1-1600w.webp";

import branding2Image400w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_2-400w.webp";
import branding2Image800w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_2-800w.webp";
import branding2Image1200w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_2-1200w.webp";
import branding2Image1600w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_2-1600w.webp";

import branding4Image400w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_4-400w.webp";
import branding4Image800w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_4-800w.webp";
import branding4Image1200w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_4-1200w.webp";
import branding4Image1600w from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_4-1600w.webp";
import brandingAnimation from "@/assets/lottie/brandingAnimation.json";

import design400w from "@/assets/images/portfolio/nolonely/design/design-400w.webp";
import design800w from "@/assets/images/portfolio/nolonely/design/design-800w.webp";
import design1400w from "@/assets/images/portfolio/nolonely/design/design-1400w.webp";
import design2400w from "@/assets/images/portfolio/nolonely/design/design-2400w.webp";
import cyclingAnimation from "@/assets/lottie/cyclingAnimation.json";

import footerBackground400w from "@/assets/images/portfolio/nolonely/footer/background-400w.webp";
import footerBackground800w from "@/assets/images/portfolio/nolonely/footer/background-800w.webp";
import footerBackground1400w from "@/assets/images/portfolio/nolonely/footer/background-1400w.webp";
import footerBackground2400w from "@/assets/images/portfolio/nolonely/footer/background-2400w.webp";

import config from "@/config.js";

export default {
  name: "nolonely",
  seoData: {
    title: "Canopas Portfolio - Nolonely - Build Healthy Habits.",
    description:
      "Nolonely is a startup to help people improve their daily life by creating healthy habits. Our team developed mobile apps for android, iOS, and the admin panel for management.",
    type: "Website",
    url: config.BASE_URL + "/portfolio/nolonely",
  },
  detail: {
    landing: {
      title: "Nolonely",
      subTitle: "Overcome loneliness and depression by building healthy habits",
      responsiveImages: true,
      backgroundImage: [
        landingBackground400w,
        landingBackground800w,
        landingBackground1400w,
        landingBackground2400w,
      ],
      alt: "nolonely-background-image",
    },
    video: {
      title: "The Problem",
      description:
        "Nolonely - A startup to help people improve their daily life by creating new habits. Habits are the core of any change people want to make in their lives. Inspired by the atomic habits book, the startup wanted to create a product that helps people focus on progress instead of goals and talents. The project is in the development mode and the habit module is live on both app stores.",
      buttons: [
        {
          name: "Play Store",
          link: "https://play.google.com/store/apps/details?id=com.canopas.nolonely",
        },
        {
          name: "App Store",
          link: "https://apps.apple.com/us/app/nolonely/id1570951174",
        },
      ],
      backgroundImage: "",
      videoBackgroundImage: videoBackgroundImage,
      alt: "nolonely-background-image",
      slider: [
        {
          id: 1,
          image: cooking,
          backgroundColor: "#FD6429",
          alt: "nolonely-cooking-activity",
        },
        {
          id: 2,
          image: meditation,
          backgroundColor: "#5C3C8C",
          alt: "nolonely-meditation-activity",
        },
        {
          id: 3,
          image: painting,
          backgroundColor: "#96BC9F",
          alt: "nolonely-painting-activity",
        },
        {
          id: 4,
          image: pet,
          backgroundColor: "#F9626B",
          alt: "nolonely-spend-time-with-pet-activity",
        },
        {
          id: 5,
          image: singing,
          backgroundColor: "#A1AEC3",
          alt: "nolonely-singing-activity",
        },
      ],
    },
    branding: {
      title: `<span class="border-text">Nolonely</span> <br/>branding`,
      responsiveImages: true,
      backgroundImage: [
        brandingBackground400w,
        brandingBackground800w,
        brandingBackground1400w,
        brandingBackground2400w,
      ],
      alt: "nolonely-branding-screen",
      details: {
        firstDetail: [
          {
            id: 1,
            aspectRatio: "134%",
            title: "",
            image: [
              branding1Image400w,
              branding1Image800w,
              branding1Image1200w,
              branding1Image1600w,
            ],
            alt: "nolonely-app-logo-image",
          },
          {
            id: 2,
            aspectRatio: "90%",
            title:
              "25+ habits from 5 unique domains to help you create the system you want for your identity.",
            image: [
              branding2Image400w,
              branding2Image800w,
              branding2Image1200w,
              branding2Image1600w,
            ],
            alt: "nolonely-habits-activity-image",
          },
        ],
        secondDetail: [
          {
            id: 3,
            background:
              "linear-gradient(140deg,hsl(240deg 53% 81%) 0%,hsl(270deg 43% 82%) 4%,hsl(305deg 34% 83%) 11%,hsl(330deg 45% 88%) 22%,hsl(347deg 45% 92%) 39%,hsl(0deg 20% 95%) 57%,hsl(357deg 76% 92%) 71%,hsl(354deg 92% 87%) 82%,hsl(350deg 96% 83%) 90%,hsl(346deg 95% 78%) 96%,hsl(342deg 92% 72%) 100%)",
            aspectRatio: "94%",
            title: "",
            animation: brandingAnimation,
            alt: "nolonely-branding-image",
          },
          {
            id: 4,
            aspectRatio: "140%",
            title: "Learn how the activity helps you improve yourself.",
            image: [
              branding4Image400w,
              branding4Image800w,
              branding4Image1200w,
              branding4Image1600w,
            ],
            alt: "nolonely-maditation-activity-screen",
          },
        ],
      },
    },
    design: [
      {
        title: `<span class="border-text">WHY</span> habits are so important?`,
        description:
          "Because it's the portion of your life you can influence to achieve your goals or desired outcome. It helps you build a system that increases the likelihood of your success or achieving goals.",
        responsiveImages: true,
        backgroundImage: [design400w, design800w, design1400w, design2400w],
        alt: "nolonely-activities-screens",
        animation: cyclingAnimation,
      },
    ],
    element: {
      class: "v2-header-2-text tw-font-bold lg:tw-w-4/5",
      title: `Much more than just a <span class="border-text">habit tracker...</span>`,
      detail: "",
    },

    footer: {
      backgroundImage: [
        footerBackground400w,
        footerBackground800w,
        footerBackground1400w,
        footerBackground2400w,
      ],
      alt: "nolonely-habit-tracker-screens",
      responsiveImages: true,
      title: config.IS_SHOW_PORTFOLIO_PAGES ? "Smile+" : "Luxe Radio",
      url: config.IS_SHOW_PORTFOLIO_PAGES ? "smileplus" : "luxeradio",
    },
  },
};
