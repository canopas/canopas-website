import landingBackground400w from "@/assets/images/portfolio/justly/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/justly/landing/background-800w.webp";
import landingBackground1400w from "@/assets/images/portfolio/justly/landing/background-1400w.webp";
import landingBackground2400w from "@/assets/images/portfolio/justly/landing/background-2400w.webp";

import videoBackgroundImage from "@/assets/images/portfolio/justly/swiper/mobile.webp";
import cooking from "@/assets/images/portfolio/justly/swiper/cooking.svg";
import meditation from "@/assets/images/portfolio/justly/swiper/meditation.svg";
import painting from "@/assets/images/portfolio/justly/swiper/painting.svg";
import pet from "@/assets/images/portfolio/justly/swiper/pet.svg";
import singing from "@/assets/images/portfolio/justly/swiper/singing.svg";

import brandingBackground400w from "@/assets/images/portfolio/justly/branding/background-400w.webp";
import brandingBackground800w from "@/assets/images/portfolio/justly/branding/background-800w.webp";
import brandingBackground1400w from "@/assets/images/portfolio/justly/branding/background-1400w.webp";
import brandingBackground2400w from "@/assets/images/portfolio/justly/branding/background-2400w.webp";

import branding1Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_1-400w.webp";
import branding1Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_1-800w.webp";
import branding1Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_1-1200w.webp";
import branding1Image1600w from "@/assets/images/portfolio/justly/branding/justly_branding_1-1600w.webp";

import branding2Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_2-400w.webp";
import branding2Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_2-800w.webp";
import branding2Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_2-1200w.webp";
import branding2Image1600w from "@/assets/images/portfolio/justly/branding/justly_branding_2-1600w.webp";

import branding3Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_3-400w.webp";
import branding3Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_3-800w.webp";
import branding3Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_3-1200w.webp";
import branding3Image1600w from "@/assets/images/portfolio/justly/branding/justly_branding_3-1600w.webp";

import branding4Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_4-400w.webp";
import branding4Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_4-800w.webp";
import branding4Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_4-1200w.webp";
import branding4Image1600w from "@/assets/images/portfolio/justly/branding/justly_branding_4-1600w.webp";

import playStoreImage from "@/assets/images/portfolio/play_store.webp";
import appStoreImage from "@/assets/images/portfolio/app_store.webp";

import design400w from "@/assets/images/portfolio/justly/design/design-400w.webp";
import design800w from "@/assets/images/portfolio/justly/design/design-800w.webp";
import design1400w from "@/assets/images/portfolio/justly/design/design-1400w.webp";
import design2400w from "@/assets/images/portfolio/justly/design/design-2400w.webp";
import cyclingAnimation from "@/assets/lottie/cyclingAnimation.json";

import footerBackground400w from "@/assets/images/portfolio/justly/footer/background-400w.webp";
import footerBackground800w from "@/assets/images/portfolio/justly/footer/background-800w.webp";
import footerBackground1400w from "@/assets/images/portfolio/justly/footer/background-1400w.webp";
import footerBackground2400w from "@/assets/images/portfolio/justly/footer/background-2400w.webp";

import config from "@/config.js";

export default {
  name: "justly",
  seoData: {
    title: "Canopas Portfolio - Justly - Build Healthy Habits.",
    description:
      "Justly is a startup to help people improve their daily life by creating healthy habits. Our team developed mobile apps for android, iOS, and the admin panel for management.",
    type: "Website",
    url: config.BASE_URL + "/portfolio/justly",
  },
  detail: {
    landing: {
      landingref1: "justlylanding",
      title: "Justly",
      subTitle: "Overcome loneliness and depression by building healthy habits",
      responsiveImages: true,
      backgroundImage: [
        landingBackground400w,
        landingBackground800w,
        landingBackground1400w,
        landingBackground2400w,
      ],
      landingref2: "justlybanner",
      alt: "justly-background-image",
    },
    video: {
      ref: "justlyvideo",
      title: "About",
      description:
        "Justly - A startup to help people improve their daily life by creating new habits. Habits are the core of any change people want to make in their lives. Inspired by the atomic habits book, the startup wanted to create a product that helps people focus on progress instead of goals and talents. The project is in the development mode and the habit module is live on both app stores.",
      buttons: [
        {
          name: "Play Store",
          link: "https://play.google.com/store/apps/details?id=com.canopas.nolonely",
          image: playStoreImage,
          event: "tap_justly_play_store",
        },
        {
          name: "App Store",
          link: "https://apps.apple.com/us/app/nolonely/id1570951174",
          image: appStoreImage,
          event: "tap_justly_app_store",
        },
      ],
      backgroundImage: "",
      videoBackgroundImage: videoBackgroundImage,
      alt: "justly-background-image",
      slider: [
        {
          id: 1,
          image: cooking,
          backgroundColor: "#FD6429",
          alt: "justly-cooking-activity",
        },
        {
          id: 2,
          image: meditation,
          backgroundColor: "#5C3C8C",
          alt: "justly-meditation-activity",
        },
        {
          id: 3,
          image: painting,
          backgroundColor: "#96BC9F",
          alt: "justly-painting-activity",
        },
        {
          id: 4,
          image: pet,
          backgroundColor: "#F9626B",
          alt: "justly-spend-time-with-pet-activity",
        },
        {
          id: 5,
          image: singing,
          backgroundColor: "#A1AEC3",
          alt: "justly-singing-activity",
        },
      ],
    },
    branding: {
      title: `<span class="border-text">Justly</span> <br/>branding`,
      responsiveImages: true,
      backgroundImage: [
        brandingBackground400w,
        brandingBackground800w,
        brandingBackground1400w,
        brandingBackground2400w,
      ],
      brandingref: "justlyparallax1",
      alt: "justly-branding-screen",
      details: {
        ref: "justlyui1",
        gridData1: [
          {
            id: 1,
            aspectRatio: "135%",
            title: "",
            image: [
              branding1Image400w,
              branding1Image800w,
              branding1Image1200w,
              branding1Image1600w,
            ],
            alt: "justly-reminder-screen",
          },
          {
            id: 2,
            aspectRatio: "94%",
            title:
              "25+ habits from 5 unique domains to help you create the system you want for your identity.",
            image: [
              branding2Image400w,
              branding2Image800w,
              branding2Image1200w,
              branding2Image1600w,
            ],
            alt: "justly-habits-activity-image",
          },
        ],
        gridData2: [
          {
            id: 3,
            aspectRatio: "94%",
            title: "",
            image: [
              branding3Image400w,
              branding3Image800w,
              branding3Image1200w,
              branding3Image1600w,
            ],
            alt: "justly-my-activities-image",
          },
          {
            id: 4,
            aspectRatio: "135%",
            title: "Learn how the activity helps you improve yourself.",
            image: [
              branding4Image400w,
              branding4Image800w,
              branding4Image1200w,
              branding4Image1600w,
            ],
            alt: "justly-maditation-activity-screen",
          },
        ],
      },
    },
    design: [
      {
        designref: "justlyfeedback1",
        title: `<span class="border-text">WHY</span> habits are so important?`,
        description:
          "Because it's the portion of your life you can influence to achieve your goals or desired outcome. It helps you build a system that increases the likelihood of your success or achieving goals.",
        responsiveImages: true,
        backgroundImage: [design400w, design800w, design1400w, design2400w],
        alt: "justly-activities-screens",
        animation: cyclingAnimation,
      },
    ],
    element: {
      class: "v2-header-2-text tw-font-bold lg:tw-w-4/5",
      title: `Much more than just a <span class="border-text">habit tracker...</span>`,
      detail: "",
    },
    footer: {
      ref: "justlyparallax2",
      backgroundImage: [
        footerBackground400w,
        footerBackground800w,
        footerBackground1400w,
        footerBackground2400w,
      ],
      alt: "justly-habit-tracker-screens",
      responsiveImages: true,
      title: "Luxe Radio",
      url: "luxeradio",
      event: "tap_next_project_luxe",
      review: {
        ref: "justlyreview",
        usersReviews: [
          {
            id: 1,
            review:
              "Kudos to the designers who designed the UI and the developers who implemented it.",
          },
          {
            id: 2,
            review:
              "Best app for making or restoring habits. Truly life changing.",
          },
          {
            id: 3,
            review:
              "Honestly, 5 stars is an underrating for this app! It's too good. The UI is so much attractive and easy to use. Keep up the Good Work🎯️",
          },
          {
            id: 4,
            review:
              "Simple and very easy to use! Very simple and easy to use application.",
          },
          {
            id: 5,
            review:
              "Great app to create new habit in your life! It’s a very good app to create a new habit.",
          },
          {
            id: 6,
            review:
              "Simply Brilliant! Everything about the app is so appealing and this is probably one of the best habit trackers out there.",
          },
          {
            id: 7,
            review:
              "It is a really amazing app that helped me be more productive and it has so many useful features.",
          },
          {
            id: 8,
            review:
              "Really amazing app. I have started doing small things regularly and it's been an awesome experience. Highly recommended!",
          },
          {
            id: 9,
            review:
              "I love this app. It really make it easy to start a new habit",
          },
          {
            id: 10,
            review:
              "If you want to track your healthy habits and want to gain regularity gradually, this app is just for you",
          },
          {
            id: 11,
            review: "Excellent!!!",
          },
          {
            id: 12,
            review:
              "I just came across it accidentally and installed it, but then after I just felt wow! It's been a very smooth experience. User's perspective from each and every degree is completely at the heart of this app. Definitely gonna use it forever!",
          },
        ],
      },
    },
  },
};
