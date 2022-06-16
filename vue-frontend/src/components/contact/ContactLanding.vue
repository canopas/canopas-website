<template>
  <div class="background">
    <div class="container tw-py-32">
      <div class="v2-header-3-text v2-canopas-gradient-text">Hey Smarty,</div>
      <div class="v2-title-2-text tw-pt-20">
        Thanks for reaching out to us. We're happy to help you achieve your
        goals and dreams.
      </div>
      <div class="v2-title-2-text tw-pt-12">
        This form requires
        <span class="canopas-gradient-text">less then 60 seconds</span>
        to fill up. Here's what you will get by filling out this form.
      </div>
      <div>
        <ul class="chat-list text-left mt-4">
          <li>
            <font-awesome-icon
              class="tw-h-6 tw-w-4 tw-text-pink-300 tw-mr-6 tw-mt-2 lg:tw-mt-3"
              icon="chevron-right"
              aria-hidden="true"
            />
            <div class="tw-w-11/12 lg:tw-w-10/12 v2-normal-text">
              Free consultation for 30 Minutes.
            </div>
          </li>
          <li>
            <font-awesome-icon
              class="tw-h-6 tw-w-4 tw-text-pink-300 tw-mr-6 tw-mt-2 lg:tw-mt-3"
              icon="chevron-right"
              aria-hidden="true"
            />
            <div class="tw-w-11/12 lg:tw-w-10/12 v2-normal-text">
              100% MONEY BACK GUARANTEE for the first TWO WEEKS of work if
              you're not satisfied. No questions asked.
            </div>
          </li>
          <li>
            <font-awesome-icon
              class="tw-h-6 tw-w-4 tw-text-pink-300 tw-mr-6 tw-mt-2 lg:tw-mt-3"
              icon="chevron-right"
              aria-hidden="true"
            />
            <div class="tw-w-11/12 lg:tw-w-10/12 v2-normal-text">
              An opportunity to grow your business exponentially.
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {};
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style lang="postcss" scoped>
.background {
  background: linear-gradient(
    270.11deg,
    rgba(249, 131, 134, 0.05) -24.42%,
    rgba(249, 131, 134, 0.05) 101.76%
  );
  box-shadow: 0px 4px 2px 1px rgba(0, 0, 0, 5%);
  z-index: -1;
}

.chat-list {
  list-style-type: none;
  margin: auto;
  padding: 0;
}

.chat-list > li {
  padding: 10px 0;
  display: flex;
  align-items: flex-start;
}

.v2-title-2-text {
  @apply tw-leading-9 lg:tw-leading-11;
}
</style>
