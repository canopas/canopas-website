import landingBackground400w from "@/assets/images/portfolio/nolonely/landing/background-400w.webp";
import landingBackground800w from "@/assets/images/portfolio/nolonely/landing/background-800w.webp";
import landingBackground1200w from "@/assets/images/portfolio/nolonely/landing/background-1200w.webp";
import landingBackground1600w from "@/assets/images/portfolio/nolonely/landing/background-1600w.webp";
import videoSectionVideo from "@/assets/images/portfolio/luxeradio/luxeradio_videosection_video.mp4";
import brandingBackgroundImage from "@/assets/images/portfolio/togness/branding/background.png";
import design400w from "@/assets/images/portfolio/nolonely/design/design-400w.webp";
import design800w from "@/assets/images/portfolio/nolonely/design/design-800w.webp";
import design1200w from "@/assets/images/portfolio/nolonely/design/design-1200w.webp";
import design1600w from "@/assets/images/portfolio/nolonely/design/design-1600w.webp";
import element1 from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_1.png";
import element2 from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_2.png";
import element3 from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_3.png";
import elementVideo from "@/assets/images/portfolio/luxeradio/element/luxeradio_element_video.mp4";
import footerBackground400w from "@/assets/images/portfolio/togness/footer/background-400w.png";
import footerBackground800w from "@/assets/images/portfolio/togness/footer/background-800w.png";
import footerBackground1200w from "@/assets/images/portfolio/togness/footer/background-1200w.png";
import footerBackground1600w from "@/assets/images/portfolio/togness/footer/background-1600w.png";

export default {
  name: "togness",
  detail: {
    landing: {
      title: "Togness",
      subTitle:
        "Slideshow maker, video editing, photos, video collage with music and quotes!",
      responsiveImages: true,
      backgroundImage: [
        landingBackground400w,
        landingBackground800w,
        landingBackground1200w,
        landingBackground1600w,
      ],
    },
    video: {
      title: "The Challenge",
      description:
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sit morbi turpis lectus. Mi adipiscing sit orci, risus, fringilla ac. In consectetur mollis amet purus nisl mauris convallis. Dolor integer purus etiam turpis placerat facilisis tincidunt sed egestas. Sit volutpat lectus duis faucibus lobortis. In lectus suspendisse ultrices sagittis rhoncus arcu pellentesque ultricies vestibulum. Quis amet vel elit erat.",
      buttons: "",
      backgroundImage: "",
      video: videoSectionVideo,
      solution: {
        title: "The Solutions",
        description:
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sit morbi turpis lectus. Mi adipiscing sit orci, risus, fringilla ac. In consectetur mollis amet purus nisl mauris convallis. Dolor integer purus etiam turpis placerat facilisis tincidunt sed egestas. Sit volutpat lectus duis faucibus lobortis. In lectus suspendisse ultrices sagittis rhoncus arcu pellentesque ultricies vestibulum. Quis amet vel elit erat.",
      },
    },
    branding: {
      title: `<span class="border-text">B</span>randing`,
      backgroundImage: brandingBackgroundImage,
      solution: {
        title: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
        description:
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Justo netus accumsan volutpat interdum ac.Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      },
      details: {
        firstDetail: [],

        secondDetail: [],
      },
    },
    design: [
      {
        title: `<span class="border-text">M</span>ain Features`,
        description: "",
        backgroundImage: [design400w, design800w, design1200w, design1600w],
        subTitle: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
        description:
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Justo netus accumsan volutpat interdum ac.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Justo netus accumsan volutpat interdum ac.Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
        technology: {
          title: "The technology we use to support Togness",
          details: [
            {
              title: "SWIFT",
            },
            {
              title: "C++",
            },
            {
              title: "JAVA",
            },
            {
              title: "HTML5",
            },
          ],
        },
      },
    ],
    element: {
      title: "Main Features",
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
      subTitle: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      description:
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Justo netus accumsan volutpat interdum ac.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Justo netus accumsan volutpat interdum ac.Lorem ipsum dolor sit amet, consectetur adipiscing elit. ",
    },
    footer: {
      backgroundImage: [
        footerBackground400w,
        footerBackground800w,
        footerBackground1200w,
        footerBackground1600w,
      ],
      responsiveImages: true,
      title: "Next portfolio",
    },
  },
};
