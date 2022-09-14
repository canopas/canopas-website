const plugin = require("tailwindcss/plugin");
const defaultTheme = require("tailwindcss/defaultTheme");
module.exports = {
  prefix: "tw-",
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      screens: {
        xs: "440px",
        sm: "576px",
        md: "768px",
        lg: "992px",
        xl: "1200px",
        "2xl": "1400px",
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
      },
      letterSpacing: {
        "extra-wider": "0.0625em",
      },
      keyframes: {
        zoomIn: {
          "0%": { transform: "scale(0, 0);" },
          "50%": { transform: "scale(1, 1);" },
          "100%": { transform: "scale(0, 0);" },
        },
        zoomOut: {
          "0%": { transform: "scale(1, 1);" },
          "50%": { transform: "scale(0, 0);" },
          "100%": { transform: "scale(1, 1);" },
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
      },
      animation: {
        zoomIn: "zoomIn 4s ease-in infinite",
        zoomOut: "zoomOut 4s ease-in infinite",
        typingErase: "typingErase 4s steps(40, end) infinite",
        menuSticky: "0.6s ease-in-out",
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
          "@screen xs": {
            maxWidth: "440px",
          },
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
