import landingBackgroundImage from "@/assets/images/portfolio/luxeradio/landing_background.png";
import videoBackgroundImage from "@/assets/images/portfolio/luxeradio/video_background.png";
import videoSectionVideo from "@/assets/images/portfolio/luxeradio/luxeradio_videosection_video.mp4";
import brandingBackgroundImage from "@/assets/images/portfolio/luxeradio/branding/background.png";
import branding1 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_1.png";
import branding2 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_2.png";
import branding3 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_3.png";
import brandVideo from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_video.mp4";
import branding5 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_5.png";
import design1 from "@/assets/images/portfolio/luxeradio/design/luxeradio_design_1.png";
import design2 from "@/assets/images/portfolio/luxeradio/design/luxeradio_design_2.png";
import element1 from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_1.png";
import element2 from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_2.png";
import element3 from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_3.png";
import elementVideo from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_video.mp4";
import footerBackgroundImage from "@/assets/images/portfolio/luxeradio/footer_background.png";

export default {
  name: "luxeradio",
  detail: {
    landing: {
      title: "Luxe Radio",
      subTitle:
        "A radio, music streaming, and podcast platform for Moroccan people.",
      responsiveImages: false,
      backgroundImage: landingBackgroundImage,
      alt: "luxeradio-app-and-desktop-images",
    },
    video: {
      title: "The Problem",
      description:
        "Luxe Radio - A broadcast media company was looking to develop mobile apps for their company. We discussed requirements with the client and understood what they were up to, how they saw the new way of broadcasting their content, and the new concepts they had for their product. From app architecture to deployment, they were looking for our help to bring their idea to life. Though, they had an excellent designer who designed the UI and UX of the product.",
      buttons: [
        {
          name: "Play Store",
          link: "https://play.google.com/store/apps/details?id=app.luxeradio.ma",
        },
        {
          name: "App Store",
          link: "https://apps.apple.com/ma/app/luxe-radio/id1073120504",
        },
      ],
      backgroundImage: videoBackgroundImage,
      video: videoSectionVideo,
    },
    branding: {
      title: "",
      solution: {
        title: "The Solution",
        description:
          "Our team added a feature to stream their content live and on-demand. We included a feature that allows users to create playlists to access the content offline. Users can also interact and post their reactions on live debates. The client wanted to handle advertising personally instead of integrating advertisement APIs. For this, We contributed to building the features that allow them to incorporate ads. <br/>We developed mobile apps for iOS, Android, and Huawei's App Gallery and web apps for Mac and Windows. Plus, the app is also available for CarPlay and Android Auto.",
      },
      responsiveImages: false,
      backgroundImage: brandingBackgroundImage,
      alt: "luxeradio-branding-screen",
      details: {
        firstDetail: [
          {
            id: 1,
            title: "A music player playing the on-demand content.",
            aspectRatio: "134%",
            image: branding1,
            alt: "luxeradio-player-screen",
          },
          {
            id: 2,
            title: "",
            aspectRatio: "135%",
            image: branding3,
            alt: "luxeradio-app-logo",
          },
          {
            id: 3,
            title: "",
            aspectRatio: "130%",
            image: branding5,
            alt: "luxeradio-mood-playlist-image",
          },
        ],

        secondDetail: [
          {
            id: 4,
            title:
              "A playlist users can listen to, follow, and share with others.",
            aspectRatio: "135%",
            image: branding2,
            alt: "luxeradio-cover-playlist-screen",
          },
          {
            id: 5,
            title: "Luxeradio has playlists that match your every mood.",
            aspectRatio: "80%",
            background:
              "linear-gradient(to top right, #afa8be 23%, #f0b4cc 100%)",
            video: brandVideo,
            alt: "luxeradio-genres-playlist-screen",
          },
        ],
      },
    },
    design: [
      {
        title: `What's so <span class="border-text">special</span> about Luxe Radio?`,
        description:
          "We developed Luxe Radio by focusing on the user experience. From app design to development, you will see that we paid attention to every micro detail in the project that makes the user's life easy.",
        responsiveImages: false,
        backgroundImage: design1,
        alt: "luxeradio-features-screen",
      },
      {
        title: "What's the Outcome of the project?",
        description:
          "Canopas has exceeded our expectations. They've understood our ideas and how big our thoughts and our expectations are — they get our vision. I can proudly say that no broadcast media application looks like ours. The Play Store is the hardest app store to get good reviews on, and we've just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas' work.",
        backgroundImage: design2,
        alt: "luxeradio-desktop-app-cover-screen",
      },
    ],
    element: {
      title: `Canopas' applied design is pixel-perfect; they respect the spacing
and everything to the smallest detail.`,
      detail: {
        firstDetail: [
          {
            background:
              "linear-gradient(135.4deg, #101010 -0.02%, #B0DDFF 125.66%);",
            aspectRatio: "100%",
            video: elementVideo,
            title: "App intro: an animation to show essential app features.",
          },
          {
            aspectRatio: "140%",
            image: element1,
            alt: "luxeradio-cover-playlist",
            title: "luxeradio-element-image",
          },
        ],
        secondDetail: [
          {
            aspectRatio: "100%",
            image: element3,
            alt: "luxeradio-element-image",
            title: "",
          },
          {
            aspectRatio: "140%",
            image: element2,
            alt: "luxeradio-element-image",
            title: "",
          },
        ],
      },
    },
    footer: {
      responsiveImages: false,
      backgroundImage: footerBackgroundImage,
      alt: "luxeradio-covers-and-library-screens",
      title: "Smile+",
    },
  },
};
