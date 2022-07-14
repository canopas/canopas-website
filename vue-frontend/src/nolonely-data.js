import landingBackground400w from "@/assets/images/portfolio/nolonely/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/nolonely/landing/background-800w.webp";
import landingBackground1200w from "@/assets/images/portfolio/nolonely/landing/background-1200w.webp";
import landingBackground1600w from "@/assets/images/portfolio/nolonely/landing/background-1600w.webp";
import videoBackgroundImage from "@/assets/images/portfolio/nolonely/swiper/mobile.webp";
import cooking from "@/assets/images/portfolio/nolonely/swiper/cooking.svg";
import meditation from "@/assets/images/portfolio/nolonely/swiper/meditation.svg";
import painting from "@/assets/images/portfolio/nolonely/swiper/painting.svg";
import pet from "@/assets/images/portfolio/nolonely/swiper/pet.svg";
import singing from "@/assets/images/portfolio/nolonely/swiper/singing.svg";
import brandingBackground400w from "@/assets/images/portfolio/nolonely/branding/background-400w.webp";
import brandingBackground800w from "@/assets/images/portfolio/nolonely/branding/background-800w.webp";
import brandingBackground1200w from "@/assets/images/portfolio/nolonely/branding/background-1200w.webp";
import brandingBackground1600w from "@/assets/images/portfolio/nolonely/branding/background-1600w.webp";
import branding1 from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_1.webp";
import branding2 from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_2.webp";
import branding3 from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_3.webp";
import branding4 from "@/assets/images/portfolio/nolonely/branding/nolonely_branding_4.webp";
import design400w from "@/assets/images/portfolio/nolonely/design/design-400w.webp";
import design800w from "@/assets/images/portfolio/nolonely/design/design-800w.webp";
import design1200w from "@/assets/images/portfolio/nolonely/design/design-1200w.webp";
import design1600w from "@/assets/images/portfolio/nolonely/design/design-1600w.webp";
import cyclingAnimation from "@/assets/images/portfolio/nolonely/design/cycling.gif";
import footerBackground400w from "@/assets/images/portfolio/nolonely/footer/background-400w.webp";
import footerBackground800w from "@/assets/images/portfolio/nolonely/footer/background-800w.webp";
import footerBackground1200w from "@/assets/images/portfolio/nolonely/footer/background-1200w.webp";
import footerBackground1600w from "@/assets/images/portfolio/nolonely/footer/background-1600w.webp";

export default {
  name: "noLonely",
  detail: {
    landing: {
      title: "NoLonely",
      subTitle: "Overcome loneliness and depression by building healthy habits",
      responsiveImages: true,
      backgroundImage: [
        landingBackground400w,
        landingBackground800w,
        landingBackground1200w,
        landingBackground1600w,
      ],
      alt: "nolonely-background-image",
    },
    video: {
      title: "The Problem",
      description:
        "NoLonely - A startup to help people improve their daily life by creating new habits. Habits are the core of any change people want to make in their lives. Inspired by the atomic habits book, the startup wanted to create a product that helps people focus on progress instead of goals and talents. The project is in the development mode and the habit module is live on both app stores.",
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
        brandingBackground1200w,
        brandingBackground1600w,
      ],
      alt: "nolonely-branding-screen",
      details: {
        firstDetail: [
          {
            id: 1,
            aspectRatio: "134%",
            title: "",
            image: branding1,
            alt: "nolonely-app-logo-image",
          },
          {
            id: 2,
            aspectRatio: "90%",
            title:
              "25+ habits from 5 unique domains to help you create the system you want for your identity.",
            image: branding2,
            alt: "nolonely-habits-activity-image",
          },
        ],
        secondDetail: [
          {
            id: 3,
            background:
              "linear-gradient(152.55deg, #F9779F -5.7%, rgba(249, 119, 159, 0) 26.47%)",
            aspectRatio: "94%",
            title: "",
            image: branding3,
            alt: "nolonely-branding-image",
          },
          {
            id: 4,
            aspectRatio: "134%",
            title: "Learn how the activity helps you improve yourself.",
            image: branding4,
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
        backgroundImage: [design400w, design800w, design1200w, design1600w],
        alt: "nolonely-activities-screens",
        gif: cyclingAnimation,
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
        footerBackground1200w,
        footerBackground1600w,
      ],
      alt: "nolonely-habit-tracker-screens",
      responsiveImages: true,
      title: "Togness",
    },
  },
};
