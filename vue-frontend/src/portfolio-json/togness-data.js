import landingBackground400w from "@/assets/images/portfolio/togness/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/togness/landing/background-800w.webp";
import landingBackground1600w from "@/assets/images/portfolio/togness/landing/background-1600w.webp";
import landingBackground2400w from "@/assets/images/portfolio/togness/landing/background-2400w.webp";

import videoSectionVideo from "@/assets/images/portfolio/togness/video/background_video.gif";
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

import element4_400w from "@/assets/images/portfolio/togness/element/togness_element_4-400w.webp";
import element4_800w from "@/assets/images/portfolio/togness/element/togness_element_4-800w.webp";
import element4_1200w from "@/assets/images/portfolio/togness/element/togness_element_4-1200w.webp";
import element4_1600w from "@/assets/images/portfolio/togness/element/togness_element_4-2400w.webp";
import draftStroyVideo from "@/assets/images/portfolio/togness/element/draft_stroy_video.gif";

import footerBackground400w from "@/assets/images/portfolio/togness/footer/background-400w.webp";
import footerBackground800w from "@/assets/images/portfolio/togness/footer/background-800w.webp";
import footerBackground1600w from "@/assets/images/portfolio/togness/footer/background-1600w.webp";
import footerBackground2400w from "@/assets/images/portfolio/togness/footer/background-2400w.webp";

import config from "@/config.js";

export default {
  name: "togness",
  seoData: {
    title:
      "Canopas Portfolio - Togness - One app for all your life's occasions.",
    description:
      "Togness is a photo editing and slideshow maker app for all your life's occasions. We developed android and iOS apps with an admin panel for content management and analytics.",
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
            event: "tap_togness_play_store",
          },
          {
            name: "App Store",
            link: "https://apps.apple.com/au/app/togness-slide-show-maker/id1528833797",
            event: "tap_togness_app_store",
          },
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
            image: [
              element4_400w,
              element4_800w,
              element4_1200w,
              element4_1600w,
            ],
            alt: "togness-apply-filters-slide-video-screen",
            title: "",
            gif: draftStroyVideo,
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
    },
  },
};
