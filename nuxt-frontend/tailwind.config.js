const plugin = require("tailwindcss/plugin");

module.exports = {
  content: [
    "./index.html",
    "./pages/**/*.{vue,js,ts,jsx,tsx}",
    "./components/**/*.{vue,js,ts,jsx,tsx}",
    "./error.vue",
  ],
  theme: {
    extend: {
      display: ["group-hover"],
      screens: {
        xs: "420px",
        sm: "576px",
        md: "768px",
        lg: "992px",
        xl: "1200px",
        xl2: "1330px",
        "2xl": "1400px",
        "2xll": "1800px",
        xll: "2440px",
        "3xl": "3840px",
        "4xl": "5000px",
        "hover-hover": {
          raw: "(hover: hover) and (pointer: fine)",
        },
      },
      flex: {
        1: "1 0 0%",
        2: "2 0 0%",
      },
      colors: {
        primary: {
          1: "#FF9472",
        },
        black: {
          core: "#000000",
          900: "#3d3d3d",
          87: "#000000DE",
          60: "#00000099",
          80: "#000000cc",
          8: "#00000014",
          4: "#00000066",
        },
        white: "#ffffff",
        "white-core": {
          1: "#FFFFFF1A",
          12: "#FFFFFF1F",
          80: "#FFFFFFCC",
          60: "#FFFFFF99",
          90: "#FFFFFF14",
          5: "#FFFFFF0d",
          8: "#ffffff14",
        },
        pink: {
          0: "#f2709c00",
          1: "#f2709c1a",
          300: "#f2709c",
          40: "#f2709c0a",
          80: "#f2709c14",
          90: "#f2709ccc",
          16: "#f2709c29",
          5: "#f2709c0d",
          2: "#f2709c33",
        },
        orange: {
          300: "#ff835b",
          1: "#ff835b1a",
          5: "#ff835b0d",
          80: "#ff835b14",
          2: "#ff835b33",
        },
        slate: {
          200: "#e2e2e2",
          400: "#8d8a8a",
        },
        "deep-charcoal": "#121212",
        "white-smoke": "#F8F8F8",
        "black-suede": "#43434366",
        "dodger-blue": "#26AEF9",
        "cool-purple": "#2671FA",
        "white-smoke-1": "#F2F2F2",
        "lavender-blush": "#FFEDF0",
        yellow: "#FFC700",
        "orange-peel": "#f985840d",
        "light-crimson": "#f476950d",
        "royal-purple": "#A97BFF",
        flamingo: "#F05138",
        "soft-yellow": "#F1E05A",
        "sea-green": "#41B883",
        "eerie-black": "#1E1E1D",
      },
      borderRadius: {
        "x-lg": "20px",
      },
      borderWidth: {
        1: "1px",
      },
      borderColor: {
        gray: "#00000033",
      },
      boxShadow: {
        "3xl": "0px 4px 45px #0000001a",
        "4xl": "2px 1000px 1px #3d3d3d inset",
      },
      lineHeight: {
        11: "2.75rem",
        17: "4.25rem",
        23: "5.75rem",
      },
      fontFamily: {
        "product-sans": ["Product Sans"],
        "inter-light": ["Inter-Light"],
        "inter-bold": ["Inter-Bold"],
        "inter-medium": ["Inter-Medium"],
        "inter-semibold": ["Inter-SemiBold"],
        "inter-regular": ["Inter-Regular"],
        "opensans-bold": ["Opensans-Bold"],
        "opensans-semibold": ["Opensans-SemiBold"],
        "source-codepro": ["Source CodePro"],
        "poppins-regular": ["Poppins Regular"],
        "poppins-medium": ["Poppins Medium"],
        "inter-extralight": ["Inter-ExtraLight"],
        "comme-light": ["Comme-Light"],
        "comme-regular": ["Comme-Regular"],
        "comme-medium": ["Comme-Medium"],
        "comme-semibold": ["Comme-SemiBold"],
      },
      letterSpacing: {
        "extra-wider": "0.0625em",
      },
      keyframes: {
        zoomIn: {
          "0%": { transform: "scale(1, 1);" },
          "50%": { transform: "scale(1.5, 1.5);" },
          "100%": { transform: "scale(1, 1);" },
        },
        zoomOut: {
          "0%": { transform: "scale(1.5, 1.5);" },
          "50%": { transform: "scale(1, 1);" },
          "100%": { transform: "scale(1.5, 1.5);" },
        },
        typingErase: {
          "0%": { width: "0;" },
          "50%,60%": { width: "100%;" },
          "95%,100%": { width: "0;" },
        },
        menuSticky: {
          "0%": { top: "-120px", opacity: "0" },
          "100%": { top: " 0", opacity: "1" },
        },
        fadeInLeft: {
          "0%": { transform: "translateX(100%)", opacity: "0" },
          "50%": { transform: "translateX(50%)", opacity: "0" },
          "100%": { transform: "translateX(0)", opacity: "1" },
        },
        fadeInRight: {
          "0%": { transform: "translateX(200px);" },
          "100%": { transform: "translateX(0px);" },
        },
        fadeIn: {
          "0%": { opacity: 0 },
          "100%": { opacity: 1 },
        },
        fadeOut: {
          "0%": { opacity: 1 },
          "100%": { opacity: 0 },
        },
        fadeUp: {
          "0%": {
            transform: "translateX(0) translateY(0)",
            opacity: 1,
          },
          "50%": {
            transform: "translateY(-300%) translateX(-50%)",
            opacity: 0,
          },
          "100%": {
            transform: "translateY(-600%) translateX(-50%)",
            opacity: 0,
          },
        },
        moveIn: {
          "0%": {
            transform: "translateY(100%);",
            opacity: 0,
          },
          "100%": {
            transform: "translateY(0);",
            opacity: 1,
          },
        },
        moveUp: {
          "0%": {
            transform: "translateY(3rem)",
          },
          "100%": {
            transform: "translateY(0)",
          },
        },
        underlineOut: {
          "0%": { transform: "scaleX(0)" },
          "100%": { transform: "scaleX(1)" },
        },
        slideDown: {
          "0%": {
            transform: "translateX(100%) translateY(-300%)",
            opacity: 1,
          },
          "50%": {
            transform: "translateX(0) translateY(0)",
            opacity: 1,
          },
        },
        zoomEffect: {
          "0%": {
            transform: " scale(1) translateX(0)",
          },
          "100%": {
            transform: " scale(1.2) translateX(15%)",
          },
        },
        mulitpleUnderlineOut: {
          "0%": {
            backgroundSize: "0%",
          },
          "100%": {
            backgroundSize: "100%",
          },
        },
      },
      animation: {
        zoomIn: "zoomIn 4s ease-in infinite",
        zoomOut: "zoomOut 4s ease-in infinite",
        typingErase: "typingErase 4s steps(40, end) infinite",
        menuSticky: "0.6s ease-in-out",
        fadeInRight: "fadeInRight 150ms ease-in",
        underlineOut: "underlineOut 1s",
        mulitpleUnderlineOut: "mulitpleUnderlineOut 1s",
        gridAnimation: "scroll 40s linear infinite",
        gridAnimationReverse: "scroll-reverse 20s linear infinite",
        gridAnimationReverse2: "scroll-reverse 40s linear infinite",
        fadeIn: "fadeIn 0.6s ease-in",
        slideDown: "slideDown 7s infinite",
        fadeUp: "fadeUp 4s ease-in",
        fadeSlide: "fadeIn 4s ease-in-out",
        fadeSlideOut: "fadeOut 2.5s ease-in-out",
        zoomEffect: "zoomEffect 0.5s ease forwards",
        moveUp: "moveUp 1s ease-in-out",
        fadeInLeft: "fadeInLeft 1s ease-in-out",
        moveIn: "moveIn 0.5s ease-in-out",
      },
      zIndex: {
        "-1": "-1",
      },
      height: {
        "3/3": "79%",
        "5/5": "97%",
      },
      inset: {
        "1/6": "15%",
      },
      bgGradientDeg: {
        247: "247.22deg",
        270: "270.11deg",
        180: "180deg",
      },
      backgroundImage: {
        "gradient-underline":
          "linear-gradient(transparent, transparent),linear-gradient(transparent, transparent),linear-gradient(to right, #f2709c, #ff835b)",
        "single-color-underline":
          "linear-gradient(to bottom, transparent 54%, #FFECEC 50%)",
        "white-gradient":
          "linear-gradient(to top, rgba(255, 255, 255, 0.6) 5%, #FFFFFF 51.63% )",
        "white-gradient-bottom":
          "linear-gradient(to bottom, rgba(255, 255, 255, 0.6) 5%, #FFFFFF 51.63%)",
        "black-gradient":
          "linear-gradient(to top, rgba(0, 0, 0, 1) 5%, #000000 51.63%)",
        "black-gradient-bottom":
          "linear-gradient(to bottom, rgba(0, 0, 0, 1) 5%, #000000 51.63%)",
        "multi-gradient-background":
          "linear-gradient(180deg, #121212 11.98%, #A65B5F 36.02%, #F98188 55.21%, #9B565A 73.96%, #121212 91.67%)",
        "pink-gradient-background":
          "linear-gradient(183deg, #FFF 36.43%, rgba(249, 133, 132, 0.50) 58.99%, #F47695 79.68%);",
        "gradient-background":
          "linear-gradient(to bottom, #FFFFFF , #FDE0E2 23%, #FDE0E2 88%, #FFFFFF 100%);",
        "gradient-L": "linear-gradient(178deg, #FFFFFF  68%, #f77783 40%);",
        "gradient-W": "linear-gradient(257deg, #FFFFFF  65%, #f77783 68%);",
        "gradient-T": "linear-gradient(540deg, #FFFFFF  33%, #f77783 10%);",
        "multiple-underline-out":
          "linear-gradient(transparent calc(100% - 5px), white 1px)",
        "gradient-underline-out":
          "linear-gradient(transparent calc(100% - 5px), #f77783 1px)",
        "black-underline-out":
          "linear-gradient(transparent calc(100% - 4px), #000000DE 1px)",
        "light-pink-gradient-background":
          "linear-gradient(180deg, #FFFEFE 3.13%, #FFF6F0 47.54%, #FFF 95.59%)",
        "togness-blue-gradient-background":
          "linear-gradient(180deg, #FFFFFF 2.84%, #E7F2FF 40.99%, #E7F2FF 48.58%, #FFFFFF 97.07%)",
        "justly-pink-gradient-background":
          "linear-gradient(180deg, #FFFFFF 2.84%, #FFE1E4 35.92%, #FFE1E4 48.58%, #FFFFFF 97.07%)",
        "luxeradio-yellow-gradient-background":
          "linear-gradient(180deg, #FFFFFF 2.84%, #FEF1DE 40.99%, #FEF1DE 52.9%, #FFFFFF 97.07%)",
        "justly-radical-gradient":
          "radial-gradient(84% 76.48% at 50% 50%, #E4949C 0%, rgba(255, 191, 171, 0) 100%)",
        "luxeradio-radical-gradient":
          "radial-gradient(84% 76.48% at 50% 50%, #FEF1DE 0%, rgba(255, 191, 171, 0) 100%)",
        "togness-radical-gradient":
          "radial-gradient(84% 76.48% at 50% 50%, #E7F2FF 0%, rgba(255, 191, 171, 0) 100%)",
        "ios-landing-background":
          "linear-gradient(184deg, #FFF 18.82%, rgba(249, 133, 132, 0.10) 52.46%, rgba(244, 118, 149, 0.10)67.02%)",
      },
    },
  },
  corePlugins: {
    container: false,
  },
  plugins: [
    plugin(function ({ matchUtilities, theme }) {
      matchUtilities(
        {
          "bg-gradient": (angle) => ({
            "background-image": `linear-gradient(${angle}, var(--tw-gradient-stops))`,
          }),
        },
        {
          values: Object.assign(theme("bgGradientDeg", {}), {
            10: "10deg",
            15: "15deg",
            20: "20deg",
            25: "25deg",
            30: "30deg",
            45: "45deg",
            60: "60deg",
            90: "90deg",
            120: "120deg",
            135: "135deg",
          }),
        },
      );
    }),
    function ({ addComponents }) {
      addComponents({
        ".container": {
          margin: "auto",
          padding: "0 0.75rem",
          maxWidth: "100%",
          "@screen sm": {
            maxWidth: "540px",
          },
          "@screen md": {
            maxWidth: "720px",
          },
          "@screen lg": {
            maxWidth: "960px",
          },
          "@screen xl": {
            maxWidth: "1140px",
          },
          "@screen 2xl": {
            maxWidth: "1320px",
          },
        },
        ".outer-container": {
          margin: "auto",
          padding: "0 0.75rem",
          maxWidth: "100%",
          "@screen 3xl": {
            maxWidth: "2550px",
          },
        },
      });
    },
  ],
};
