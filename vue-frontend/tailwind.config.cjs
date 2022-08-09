module.exports = {
  prefix: "tw-",
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      container: {
        center: true,
        padding: {
          DEFAULT: "1rem",
          sm: "2rem",
          lg: "4rem",
          xl: "5rem",
          "2xl": "6rem",
        },
      },
      screens: {
        sm: "576px",
        md: "764px",
        lg: "992px",
        xl: "1200px",
        "2xl": "1400px",
      },
      flex: {
        1: "1 0 0%",
        2: "2 0 0%",
      },
      colors: {
        black: {
          900: "#3d3d3d",
        },
        white: "#ffffff ",
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
      borderWidth: {
        1: "1px",
      },
      boxShadow: {
        "3xl": "0px 4px 45px rgba(0, 0, 0, 0.1)",
      },
      lineHeight: {
        11: "2.75rem",
        17: "4.25rem",
        23: "5.75rem",
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
      },
      animation: {
        zoomIn: "zoomIn 4s ease-in infinite",
        zoomOut: "zoomOut 4s ease-in infinite",
        typingErase: "typingErase 4s steps(40, end) infinite",
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
    },
    plugins: [],
  },
};
