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
      },
    },
    plugins: [],
  },
};
