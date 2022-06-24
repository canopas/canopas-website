import landingBackgroundImage from "@/assets/images/portfolio/luxeradio/landing_background.png";
import videoBackgroundImage from "@/assets/images/portfolio/luxeradio/video_background.png";
import brandingBackgroundImage from "@/assets/images/portfolio/luxeradio/branding/background.png";
import branding1 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_1.png";
import branding2 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_2.png";
import branding3 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_3.png";
import branding4 from "@/assets/images/portfolio/luxeradio/branding/luxeradio_branding_4.png";
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
      video: "",
    },
    branding: {
      title: "Careful branding",
      backgroundImage: brandingBackgroundImage,
      details: [
        {
          id: 1,
          title:
            "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
          image: branding1,
        },
        {
          id: 2,
          title:
            "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
          image: branding2,
        },
        {
          id: 3,
          title: "",
          image: branding3,
        },
        {
          id: 4,
          title: "",
          image: branding4,
        },
        {
          id: 5,
          title:
            "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard.",
          image: branding5,
        },
      ],
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

    footer: { backgroundImage: footerBackgroundImage },
  },
};
