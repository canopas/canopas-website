import { useElementVisibility } from "@vueuse/core";

const events = {
  budget: "view_ad_budget_section",
  portfolio: "view_home_portfolio_section",
  parallax: "view_parallax_section",
  clientreview: "view_client_review_section",
  userreview: "view_user_review_section",
  blog: "view_blog_post_section",
  contribution: "view_open_source_contribution",
  CTA: "view_home_page_end_cta",
  footer: "view_home_page_footer",
  service1: "view_service_section_first_row",
  service2: "view_service_section_second_row",
  portfolios: {
    0: "view_home_luxe_radio_portfolio",
    1: "view_home_togness_portfolio",
    2: "view_home_justly_portfolio",
    3: "view_home_smile_portfolio",
  },
};

export function analyticsEvent(refs) {
  let event = "";
  Object.keys(refs).forEach((key) => {
    if (key == "portfolios") {
      refs[key].forEach((ref, index) => {
        if (useElementVisibility(refs[key][index]).value) {
          event = events[key][index];
          return;
        }
      });
    } else {
      if (useElementVisibility(refs[key]).value) {
        event = events[key];
        return;
      }
    }
  });
  return event;
}
