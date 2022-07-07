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
      subTitle: "Lorem Ipsum is simply dummy text of the printing.",
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
        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the i ndustry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
      buttons: [
        { name: "Launch Project", link: "" },
        { name: "App Store", link: "" },
      ],
      videoBackgroundImage: videoBackgroundImage,
      slider: [
        { id: 1, image: cooking, backgroundColor: "#FD6429" },
        { id: 2, image: meditation, backgroundColor: "#5C3C8C" },
        { id: 3, image: painting, backgroundColor: "#7A8494" },
        { id: 4, image: pet, backgroundColor: "#F9626B" },
        { id: 5, image: singing, backgroundColor: "#D7EFFF" },
      ],
    },
    branding: {
      title: `<span class="border-text">Careful</span> <br/>branding`,
      responsiveImages: true,
      backgroundImage: [
        brandingBackground400w,
        brandingBackground800w,
        brandingBackground1200w,
        brandingBackground1600w,
      ],
      details: {
        firstDetail: [
          {
            id: 1,
            aspectRatio: "132%",
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
            image: branding1,
          },
          {
            id: 2,
            aspectRatio: "90%",
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard.",
            image: branding2,
          },
        ],
        secondDetail: [
          {
            id: 3,
            background:
              "linear-gradient(152.55deg, #F9779F -5.7%, rgba(249, 119, 159, 0) 26.47%)",
            aspectRatio: "94%",
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard.",
            image: branding3,
          },
          {
            id: 4,
            aspectRatio: "132%",
            title:
              "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
            image: branding4,
          },
        ],
      },
    },
    design: [
      {
        title: `The <span class="border-text">design</span> will not leave you cold`,
        description:
          "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard. Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
        responsiveImages: true,
        backgroundImage: [design400w, design800w, design1200w, design1600w],
        gif: cyclingAnimation,
      },
    ],
    element: {
      title: `Every <span class="border-text">element</span><br/> made with love`,
      detail: "",
    },

    footer: {
      backgroundImage: [
        footerBackground400w,
        footerBackground800w,
        footerBackground1200w,
        footerBackground1600w,
      ],
      responsiveImages: true,
      title: "Togness",
    },
  },
};
