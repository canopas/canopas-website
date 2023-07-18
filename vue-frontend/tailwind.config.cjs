const plugin = require("tailwindcss/plugin");

module.exports = {
  prefix: "tw-",
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
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
        "hover-hover": { raw: "(hover: hover) and (pointer: fine)" },
      },
      flex: {
        1: "1 0 0%",
        2: "2 0 0%",
      },
      colors: {
        black: {
          core: "#000000",
          900: "#3d3d3d",
        },
        white: "#ffffff",
        pink: {
          300: "#f2709c",
        },
        orange: {
          300: "#ff835b",
        },
        slate: {
          200: "rgb(226, 226, 226)",
          400: "rgb(141, 138, 138)",
        },
      },
      borderRadius: {
        "x-lg": "20px",
      },
      borderWidth: {
        1: "1px",
      },
      borderColor: {
        gray: "rgba(0, 0, 0, 0.2)",
      },
      boxShadow: {
        "3xl": "0px 4px 45px rgba(0, 0, 0, 0.1)",
        "4xl": "2px 1000px 1px #3d3d3d inset",
      },
      lineHeight: {
        11: "2.75rem",
        17: "4.25rem",
        23: "5.75rem",
      },
      fontFamily: {
        "product-sans": ["Product Sans"],
        "futura-bold": ["FuturaLT-Bold"],
        "roboto-bold": ["Roboto-Bold"],
        "roboto-medium": ["Roboto-Medium"],
        "inter-light": ["Inter-Light"],
        "inter-bold": ["Inter-Bold"],
        "inter-medium": ["Inter-Medium"],
        "inter-semibold": ["Inter-SemiBold"],
        "inter-regular": ["Inter-Regular"],
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
      },
      animation: {
        zoomIn: "zoomIn 4s ease-in infinite",
        zoomOut: "zoomOut 4s ease-in infinite",
        typingErase: "typingErase 4s steps(40, end) infinite",
        menuSticky: "0.6s ease-in-out",
        fadeInRight: "fadeInRight 150ms ease-in",
        underlineOut: "underlineOut 1s",
        gridAnimation: "scroll 40s linear infinite",
        gridAnimationReverse: "scroll-reverse 20s linear infinite",
        gridAnimationReverse2: "scroll-reverse 40s linear infinite",
        fadeIn: "fadeIn 0.6s ease-in",
        slideDown: "slideDown 7s infinite",
        fadeUp: "fadeUp 4s ease-in",
        fadeSlide: "fadeIn 4s ease-in-out",
        fadeSlideOut: "fadeOut 2.5s ease-in-out",
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
        "multi-gradient-background":
          "linear-gradient(180deg, #1A1A1A 11.98%, #A65B5F 36.02%, #F98188 55.21%, #9B565A 73.96%, #1A1A1A 91.67%)",
        "pink-gradient-background":
          "linear-gradient(to top , #FFDCE8 -36%, #F2709C 19.02%, #FF9472 38.21%, #FFFFFF 65.96%, #FFFFFF 91.67%)",
        "gradient-background":
          "linear-gradient(to bottom, #FFFFFF , #FDE0E2 23%, #FDE0E2 88%, #FFFFFF 100%);",
        "gradient-L": "linear-gradient(178deg, #FFFFFF  68%, #f77783 40%);",
        "gradient-W": "linear-gradient(257deg, #FFFFFF  65%, #f77783 68%);",
        "gradient-T": "linear-gradient(540deg, #FFFFFF  30%, #f77783 10%);",
        "multiple-underline-out":
          "linear-gradient(transparent calc(100% - 5px), white 1px)",
        "light-pink-gradient-background":
          "linear-gradient(180deg, #FFFEFE 3.13%, #FFF6F0 47.54%, #FFF 95.59%)",
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
        }
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
      });
    },
  ],
};
