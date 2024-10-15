import landingBackground400w from "@/assets/images/portfolio/togness/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/togness/landing/background-800w.webp";
import landingBackground1600w from "@/assets/images/portfolio/togness/landing/background-1600w.webp";
import landingBackground2400w from "@/assets/images/portfolio/togness/landing/background-2400w.webp";

import videoSectionVideo from "@/assets/images/portfolio/togness/video/background_video.mp4";
import videoBackground400w from "@/assets/images/portfolio/togness/video/background-400w.webp";
import videoBackground800w from "@/assets/images/portfolio/togness/video/background-800w.webp";
import videoBackground1600w from "@/assets/images/portfolio/togness/video/background-1600w.webp";
import videoBackground2400w from "@/assets/images/portfolio/togness/video/background-2400w.webp";

import branding1Image400w from "@/assets/images/portfolio/togness/branding/togness_branding_1-400w.webp";
import branding1Image800w from "@/assets/images/portfolio/togness/branding/togness_branding_1-800w.webp";
import branding1Image1200w from "@/assets/images/portfolio/togness/branding/togness_branding_1-1200w.webp";
import branding1Image1600w from "@/assets/images/portfolio/togness/branding/togness_branding_1-1600w.webp";

import branding3Image400w from "@/assets/images/portfolio/togness/branding/togness_branding_2-400w.webp";
import branding3Image800w from "@/assets/images/portfolio/togness/branding/togness_branding_2-800w.webp";
import branding3Image1200w from "@/assets/images/portfolio/togness/branding/togness_branding_2-1200w.webp";
import branding3Image1600w from "@/assets/images/portfolio/togness/branding/togness_branding_2-1600w.webp";

import branding2Image400w from "@/assets/images/portfolio/togness/branding/togness_branding_3-400w.webp";
import branding2Image800w from "@/assets/images/portfolio/togness/branding/togness_branding_3-800w.webp";
import branding2Image1200w from "@/assets/images/portfolio/togness/branding/togness_branding_3-1200w.webp";
import branding2Image1600w from "@/assets/images/portfolio/togness/branding/togness_branding_3-1600w.webp";

import branding4Image400w from "@/assets/images/portfolio/togness/branding/togness_branding_4-400w.webp";
import branding4Image800w from "@/assets/images/portfolio/togness/branding/togness_branding_4-800w.webp";
import branding4Image1200w from "@/assets/images/portfolio/togness/branding/togness_branding_4-1200w.webp";
import branding4Image1600w from "@/assets/images/portfolio/togness/branding/togness_branding_4-1600w.webp";
import draftStoryVideo from "@/assets/images/portfolio/togness/branding/draft_story_video.mp4";

import design400w from "@/assets/images/portfolio/togness/design/background-400w.webp";
import design800w from "@/assets/images/portfolio/togness/design/background-800w.webp";
import design1400w from "@/assets/images/portfolio/togness/design/background-1600w.webp";
import design2400w from "@/assets/images/portfolio/togness/design/background-2400w.webp";

import element1_400w from "@/assets/images/portfolio/togness/element/togness_element_1-400w.webp";
import element1_800w from "@/assets/images/portfolio/togness/element/togness_element_1-800w.webp";
import element1_1200w from "@/assets/images/portfolio/togness/element/togness_element_1-1200w.webp";
import element1_1600w from "@/assets/images/portfolio/togness/element/togness_element_1-2400w.webp";

import element2_400w from "@/assets/images/portfolio/togness/element/togness_element_2-400w.webp";
import element2_800w from "@/assets/images/portfolio/togness/element/togness_element_2-800w.webp";
import element2_1200w from "@/assets/images/portfolio/togness/element/togness_element_2-1200w.webp";
import element2_1600w from "@/assets/images/portfolio/togness/element/togness_element_2-2400w.webp";

import element3_400w from "@/assets/images/portfolio/togness/element/togness_element_3-400w.webp";
import element3_800w from "@/assets/images/portfolio/togness/element/togness_element_3-800w.webp";
import element3_1200w from "@/assets/images/portfolio/togness/element/togness_element_3-1200w.webp";
import element3_1600w from "@/assets/images/portfolio/togness/element/togness_element_3-2400w.webp";
import playStoreImage from "@/assets/images/portfolio/play_store.webp";

import footerBackground400w from "@/assets/images/portfolio/togness/footer/background-400w.webp";
import footerBackground800w from "@/assets/images/portfolio/togness/footer/background-800w.webp";
import footerBackground1600w from "@/assets/images/portfolio/togness/footer/background-1600w.webp";
import footerBackground2400w from "@/assets/images/portfolio/togness/footer/background-2400w.webp";

import config from "@/config.js";

export default {
  name: "togness",
  seoData: {
    title:
      "Togness - One app for all your life's occasions | Photo and Video Editing App | Canopas",
    description:
      "The only photo and video editor for your memorable occasions. Canopas developed native Android and iOS apps with a backend and admin panel.",
    type: "Website",
    url: config.BASE_URL + "/portfolio/togness",
  },
  detail: {
    landing: {
      landingref1: "tognesslanding",
      title: "Togness",
      subTitle:
        "The best photo editor and slideshow maker app for your life's occasions.",
      backgroundImage: [
        landingBackground400w,
        landingBackground800w,
        landingBackground1600w,
        landingBackground2400w,
      ],
      landingref2: "tognessbanner",
      alt: "togness-screen-images",
    },
    video: {
      ref: "tognessvideo",
      title: "The Problem",
      description:
        "Togness - A client was on a mission to salvage their idea. Initially, they hired a team and all went well, at least for a while! When faced with challenges, the team that said they didn't had the development skills to complete the app as initially agreed upon. That's when we came into the picture. We started with client to fix a few issues to demonstrate our capability. Later, once the client was satisfied, full project was handed over to us. The client wanted to develop apps for both Android and iOS platforms.",
      buttons: "",
      backgroundImage: "",
      animation: videoSectionVideo,
      videoBackground: [
        videoBackground400w,
        videoBackground800w,
        videoBackground1600w,
        videoBackground2400w,
      ],
      alt: "togness-update-quote-and-select-quote-video-image",
      solution: {
        ref: "tognesssolution",
        title: "The Solution",
        description:
          "We started with the complex issue that the previous team was unable to achieve. Once the android development was well underway, we started iOS. As a team, we also helped the client with REST API, admin panel, and DevOps services. Some of the main features of the app are: use a photo gallery to add media files, merge images and videos, add quotes, add music, merge audio of video and music, apply filters, apply transitions effect, export in different aspect ratios, customize photo slide durations, etc. ",
        buttons: [
          {
            name: "Play Store",
            link: "https://play.google.com/store/apps/details?id=com.togness.story",
            image: playStoreImage,
            event: "tap_togness_play_store",
          },
          // {
          //   name: "App Store",
          //   link: "https://apps.apple.com/au/app/togness-slide-show-maker/id1528833797",
          //   image: appStoreImage,
          //   event: "tap_togness_app_store",
          // },
        ],
      },
    },
    branding: {
      title: ``,
      backgroundImage: "",
      details: {
        ref: "tognessui1",
        gridData1: [
          {
            background: "",
            aspectRatio: "145%",
            image: [
              element1_400w,
              element1_800w,
              element1_1200w,
              element1_1600w,
            ],
            alt: "togness-music-selection-screen",
            title: "Music tracks in four different lengths to match the story.",
          },
          {
            aspectRatio: "145%",
            image: [
              element2_400w,
              element2_800w,
              element2_1200w,
              element2_1600w,
            ],
            alt: "togness-apply-transition-slides-screen",
            title: "",
          },
        ],
        gridData2: [
          {
            aspectRatio: "122%",
            image: [
              element3_400w,
              element3_800w,
              element3_1200w,
              element3_1600w,
            ],
            alt: "togness-all-screens-images",
            title: "Overview of the app that shows multiple features.",
          },
          {
            aspectRatio: "164%",
            alt: "togness-apply-filters-slide-video-screen",
            title: "togness",
            video: draftStoryVideo,
          },
        ],
      },
    },
    design: [
      {
        title: `Togness : One app for all your life's occasions`,
        backgroundImage: [design400w, design800w, design1400w, design2400w],
        designref1: "tognessparallax1",
        alt: "togness-branding-screens",
        subTitle: "What's the outcome of the project?",
        subDescription: `We have only just launched, so unable to provide metrics but it works and works really well. Any issues are rectified quickly.<br/><br/>
        What I found most impressive or unique about Canopas is that there was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they understood the first time. This proved they knew exactly what they were doing. Very reassuring!`,
        designref2: "tognessfeedback1",
      },
    ],
    element: {
      title: ``,
      detail: {
        gridData1: [
          {
            id: 1,
            title: "",
            aspectRatio: "164%",
            image: [
              branding1Image400w,
              branding1Image800w,
              branding1Image1200w,
              branding1Image1600w,
            ],
            alt: "togness-video-edit-screen",
          },
          {
            id: 2,
            title: "",
            aspectRatio: "118%",
            image: [
              branding2Image400w,
              branding2Image800w,
              branding2Image1200w,
              branding2Image1600w,
            ],
            alt: "togness-select-moods-screen",
          },
        ],

        gridData2: [
          {
            id: 3,
            title: "",
            aspectRatio: "118%",
            image: [
              branding3Image400w,
              branding3Image800w,
              branding3Image1200w,
              branding3Image1600w,
            ],
            alt: "togness-update-to-premium-popup-screen",
          },
          {
            id: 4,
            title: "",
            aspectRatio: "164%",
            image: [
              branding4Image400w,
              branding4Image800w,
              branding4Image1200w,
              branding4Image1600w,
            ],
            alt: "togness-selection-story-screen",
          },
        ],
      },
    },
    footer: {
      ref: "tognessparallax2",
      backgroundImage: [
        footerBackground400w,
        footerBackground800w,
        footerBackground1600w,
        footerBackground2400w,
      ],
      alt: "togness-dark-mode-screens",
      title: "Justly",
      url: "justly",
      event: "tap_next_project_justly",
      review: {
        ref: "tognessreview",
        usersReviews: [
          {
            id: 1,
            review:
              "Makes slideshows easy. Really helped me create a slideshow to remember my grandfather so much easier. Recommend it for anyone wanting to share memories with their family.",
          },
          {
            id: 2,
            review:
              "It's a great app. Some useful feature and free premium is a plus.Highly recommended",
          },
          {
            id: 3,
            review:
              "Good easy to use apps, And can make slideshow look very nice",
          },
          {
            id: 4,
            review:
              "I use this app a lot. I've already made some slideshows of my pets. So cute!",
          },
          {
            id: 5,
            review:
              "It had amazing quotes, And was super easy to use! I would say it's one of a kind!",
          },
          {
            id: 6,
            review:
              "This is a really neat app. I've always been terrible with producing slideshows but this app is very easy to use and the results are beautiful.",
          },
          {
            id: 7,
            review: "Great UI. Easy and fun to use.",
          },
          {
            id: 8,
            review:
              "Easy to use and fun to use. Nowwww I would say that this app is probably one of the best slideshow makers I know! I highly recommend it to any of you guys that need a app for making slide shows.",
          },
          {
            id: 9,
            review:
              "Very user friendly! This App is straightforward and simple to use. The step by step process with it’s categories, music, photo & quote options makes it easy for those who are not overly- tech savvy.",
          },
          {
            id: 10,
            review:
              "I really enjoy putting together our special moments using Togness. It’s simple and easy to use.",
          },
          {
            id: 11,
            review:
              "I’ve always struggled to make something special because it is always difficult. I found this very easy to produce just what I wanted.",
          },
          {
            id: 12,
            review:
              "Great and user-friendly. Love it! Super easy to make slide shows and photo montages, complete with quotes and music.",
          },
          {
            id: 13,
            review:
              "I can use this for different occasions and it is quick to put something together that looks quite professional.",
          },
          {
            id: 14,
            review:
              "Unique and easy. Very easy and intuitive putting together photos, videos with music and quotes.",
          },
        ],
      },
    },
  },
};
