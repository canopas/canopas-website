import config from "./config.js";

const robots = {
  UserAgent: "*",
};

if (config.IS_PROD) {
  robots.Allow = "/";
  robots.Sitemap = "https://canopas.com/sitemap.xml";
} else {
  robots.Disallow = "/";
}

export default robots;
