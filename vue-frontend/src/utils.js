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
  landingview: "view_jobs_landing_section",
  virtue: "view_virtues_section",
  perks: "view_perks_benefits",
  whycanopas: "view_why_canopas",
  joblist: "view_jobs_list",
  readmore: "view_job_read_more",
  applyjob: "view_job_apply",
  jobfailed: "job_apply_failed",
  jobsuccess: "job_apply_success",
  portfoliolist: {
    0: "view_portfolio_page_luxe_radio",
    1: "view_portfolio_page_togness",
    2: "view_portfolio_page_justly",
    3: "view_portfolio_page_smile",
  },
  portfoliocta: "view_portfolio_page_end_cta",

  luxelanding: "view_luxe_radio_landing",
  luxebanner: "view_luxe_first_banner_image",
  luxevideo: "view_luxe_radio_video",
  luxesolution: "view_luxe_radio_solution",
  luxeparallax1: "view_luxe__first_parallax_section",
  luxeui1: "view_luxe_ui_section1",
  luxefeedback1: "view_luxe_feedback1",
  luxeparallax2: "view_luxe_last_parallax_section",
  luxepagecta: "view_luxe_page_end_cta",

  tognesslanding: "view_togness_landing",
  tognessbanner: "view_togness_first_banner_image",
  tognessvideo: "view_togness_problem",
  tognesssolution: "view_togness_solution",
  tognessparallax1: "view_togness_parallax1",
  tognessui1: "view_togness_ui_section1",
  tognessfeedback1: "view_togness_feedback1",
  tognessparallax2: "view_togness_last_parallax",
  tognesspagecta: "view_togness_page_end_cta",

  justlylanding: "view_justly_landing",
  justlybanner: "view_justly_first_banner_image",
  justlyvideo: "view_justly_horizontal_slider",
  justlyparallax1: "view_justly_parallax",
  justlyui1: "view_justly_ui_section1",
  justlyfeedback1: "view_justly_feedback1",
  justlyparallax2: "view_justly_last_parallax",
  justlypagecta: "view_justly_page_end_cta",
};

export function analyticsEvent(refs) {
  let event = "";
  Object.keys(refs).forEach((key) => {
    if (key == "portfolios" || key == "portfoliolist") {
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
