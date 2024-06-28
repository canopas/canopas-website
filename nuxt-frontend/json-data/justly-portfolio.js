import landingBackground400w from "@/assets/images/portfolio/justly/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/justly/landing/background-800w.webp";
import landingBackground1400w from "@/assets/images/portfolio/justly/landing/background-1400w.webp";
import landingBackground2400w from "@/assets/images/portfolio/justly/landing/background-2400w.webp";

import feature1_400w from "@/assets/images/portfolio/justly/features/feature1-400w.webp";
import feature1_800w from "@/assets/images/portfolio/justly/features/feature1-800w.webp";
import feature1_1200w from "@/assets/images/portfolio/justly/features/feature1-1200w.webp";

import feature2_400w from "@/assets/images/portfolio/justly/features/feature2-400w.webp";
import feature2_800w from "@/assets/images/portfolio/justly/features/feature2-800w.webp";
import feature2_1200w from "@/assets/images/portfolio/justly/features/feature2-1200w.webp";

import feature3_400w from "@/assets/images/portfolio/justly/features/feature3-400w.webp";
import feature3_800w from "@/assets/images/portfolio/justly/features/feature3-800w.webp";
import feature3_1200w from "@/assets/images/portfolio/justly/features/feature3-1200w.webp";

import feature4_400w from "@/assets/images/portfolio/justly/features/feature4-400w.webp";
import feature4_800w from "@/assets/images/portfolio/justly/features/feature4-800w.webp";
import feature4_1200w from "@/assets/images/portfolio/justly/features/feature4-1200w.webp";

import brandingBackground400w from "@/assets/images/portfolio/justly/branding/background-400w.webp";
import brandingBackground800w from "@/assets/images/portfolio/justly/branding/background-800w.webp";
import brandingBackground1400w from "@/assets/images/portfolio/justly/branding/background-1400w.webp";
import brandingBackground2400w from "@/assets/images/portfolio/justly/branding/background-2400w.webp";

import branding1Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_1-400w.webp";
import branding1Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_1-800w.webp";
import branding1Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_1-1200w.webp";

import branding2Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_2-400w.webp";
import branding2Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_2-800w.webp";
import branding2Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_2-1200w.webp";

import branding3Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_3-400w.webp";
import branding3Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_3-800w.webp";
import branding3Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_3-1200w.webp";

import branding4Image400w from "@/assets/images/portfolio/justly/branding/justly_branding_4-400w.webp";
import branding4Image800w from "@/assets/images/portfolio/justly/branding/justly_branding_4-800w.webp";
import branding4Image1200w from "@/assets/images/portfolio/justly/branding/justly_branding_4-1200w.webp";

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
    title: "Justly - Build healthy habits with Justly Habit Tracker",
    description:
      "Developed Android and iOS apps to help people create healthy habits. A habit tracker to improve your life by 1% every day.",
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
        "Justly helps you improve your mental and physical well-being and reach your goals. It offers tools and resources for weight loss, relationship improvement, and career advancement. But it's not just about personal goals - Justly is committed to advancing humanity. Download the app to learn more about our meticulous approach to every aspect of the project, from APP ARCHITECTURE TO DEPLOYMENT.",
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
      alt: "justly-background-image",
      features: {
        gridData1: [
          {
            id: 1,
            aspectRatio: "150%",
            image: [feature1_400w, feature1_800w, feature1_1200w],
            alt: "justly-logo",
          },
          {
            id: 2,
            aspectRatio: "130%",
            image: [feature2_400w, feature2_800w, feature2_1200w],
            alt: "goal-feture",
          },
        ],
        gridData2: [
          {
            id: 3,
            aspectRatio: "115%",
            image: [feature3_400w, feature3_800w, feature3_1200w],
            alt: "rank-feture",
          },
          {
            id: 4,
            aspectRatio: "150%",
            image: [feature4_400w, feature4_800w, feature4_1200w],
            alt: "notes-feature",
          },
        ],
      },
    },
    branding: {
      title: `<span class="border-text">Careful</span> <br/>branding`,
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
            aspectRatio: "173%",
            title:
              "25+ habits from 5 unique domains to help you create the system you want for your identity.",
            image: [
              branding1Image400w,
              branding1Image800w,
              branding1Image1200w,
            ],
            alt: "justly-reminder-screen",
          },
          {
            id: 2,
            aspectRatio: "121%",
            title:
              "Create a community, find a wise mentor, or ask a friend to hold you accountable for the habits you want to cultivate.",
            image: [
              branding2Image400w,
              branding2Image800w,
              branding2Image1200w,
            ],
            alt: "justly-habits-activity-image",
          },
        ],
        gridData2: [
          {
            id: 3,
            aspectRatio: "144%",
            title:
              "Never lose your focus again with a wealth of stats and highlights to keep you on track.",
            image: [
              branding3Image400w,
              branding3Image800w,
              branding3Image1200w,
            ],
            alt: "justly-my-activities-image",
          },
          {
            id: 4,
            aspectRatio: "143%",
            title:
              "Adopt the daily routines of successful people to set yourself apart and achieve your goals.",
            image: [
              branding4Image400w,
              branding4Image800w,
              branding4Image1200w,
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
      class: "mobile-header-2 lg:desk-header-2",
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
