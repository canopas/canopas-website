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
      title: "Luxeradio",
      subTitle: "Lorem Ipsum is simply dummy text of the printing.",
      backgroundImage: landingBackgroundImage,
    },
    video: {
      title: "The Challenge",
      description:
        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the i ndustry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
      buttons: [
        { name: "Launch Project", link: "" },
        { name: "App Store", link: "" },
      ],
      backgroundImage: videoBackgroundImage,
      video: videoSectionVideo,
    },
    branding: {
      title: `<span class="border-text">Careful</span> <br/>branding`,

      backgroundImage: brandingBackgroundImage,
      details: {
        firstDetail: [
          {
            id: 1,
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
            aspectRatio: "90%",
            image: branding1,
          },
          {
            id: 2,
            title: "",
            aspectRatio: "90%",
            image: branding3,
          },
          {
            id: 3,
            title: "",
            aspectRatio: "90%",
            image: branding5,
          },
        ],

        secondDetail: [
          {
            id: 4,
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
            aspectRatio: "135%",
            image: branding2,
          },
          {
            id: 5,
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard.",
            aspectRatio: "80%",
            background:
              "linear-gradient(to top right, #afa8be 23%, #f0b4cc 100%)",
            video: brandVideo,
          },
        ],
      },
    },
    design: [
      {
        title: `The <span class="border-text">design</span> will not leave you cold`,
        description:
          "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
        backgroundImage: design1,
      },
      {
        title: "Extra attention to detail development",
        description:
          "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
        backgroundImage: design2,
      },
    ],
    element: {
      title: `Every <span class="border-text">element</span><br/> made with love`,
      detail: {
        firstDetail: [
          {
            background:
              "linear-gradient(135.4deg, #101010 -0.02%, #B0DDFF 125.66%);",
            aspectRatio: "100%",
            video: elementVideo,
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
          },
          {
            aspectRatio: "140%",
            image: element1,
            title: "",
          },
        ],
        secondDetail: [
          {
            aspectRatio: "100%",
            image: element3,
            title: "",
          },
          {
            aspectRatio: "140%",
            image: element2,
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
          },
        ],
      },
    },

    footer: { backgroundImage: footerBackgroundImage, title: "Smile+" },
  },
};
