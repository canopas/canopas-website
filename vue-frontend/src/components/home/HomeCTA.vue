<template>
  <div class="cta-container text-center">
    <div class="header-text canopas-gradient-text mt-2">
      How do you know we’re a good fit for you or not?
    </div>
    <div class="normal-text mt-3">
      Let’s have a short chat or call to find out.
    </div>
    <div class="communication">
      <div class="item communication-item1">
        I love talking with people
        <button class="gradient-btn" @click="openContact">
          <font-awesome-icon class="fa" icon="phone-alt" aria-hidden="true" />
          <span>Call Now</span>
        </button>
      </div>
      <div class="item-divider"></div>
      <div class="item communication-item2">
        I am an Introvert. I prefer to chat or writing a mail.
        <div class="chat-btn-container">
          <button class="gradient-btn" @click="openContact">
            <font-awesome-icon
              class="fa"
              icon="comment-alt"
              aria-hidden="true"
            />
            <span>Chat Now</span>
          </button>
          <button class="gradient-btn" @click="openContact">
            <font-awesome-icon
              class="fa"
              icon="paper-plane"
              aria-hidden="true"
            />
            <span>Drop a mail</span>
          </button>
        </div>
      </div>
    </div>
    <div class="normal-text text-left reason-text">
      Four reasons why do we do this call?
    </div>
    <ul class="normal-text chat-list text-left">
      <li>
        <span class="bullet"></span>
        <div>It gives an opportunity to learn more about each other.</div>
      </li>
      <li>
        <span class="bullet"></span>
        <div>
          We want to give you a chance to put your idea into action. You might
          have faced this situation before where you wish you had someone to
          help you.
        </div>
      </li>
      <li>
        <span class="bullet"></span>
        <div>
          Sometimes, you want to check the feasibility of your idea. We will
          give you honest feedback based on our industry experience.
        </div>
      </li>
      <li>
        <span class="bullet"></span>
        <div>
          As you read above, we offer 100% Money-back guarantee. so, there’s no
          risk to you.
        </div>
      </li>
    </ul>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {};
  },
  components: {
    FontAwesomeIcon,
  },
  methods: {
    openContact() {
      this.$router.push("/contact");
    },
  },
};
</script>

<style lang="scss" scoped>
.cta-container {
  display: flex;
  flex-direction: column;
  margin: 10% auto;
  width: 80%;
  padding: 0 5%;
}

.communication {
  display: flex;
  flex-direction: column;
  margin: 1.25rem 0;
  border: 2px solid #e2e2e2;
  border-radius: 0.9rem;
}

.communication > .item {
  padding: 1rem;
  margin: auto;
  text-align: -webkit-center;
}

.item-divider {
  width: 100%;
  height: 2px;
  background: #e2e2e2;
}

.chat-btn-container {
  display: flex;
  flex-direction: column;
  margin: auto;
}

.gradient-btn {
  border-radius: 0.6rem;
  padding: 1rem 1.25rem;
  margin: 1.25rem auto 0 auto;
  display: flex;
  text-align: center;
  color: white;
  background: linear-gradient(91.53deg, #f2709c 3.91%, #ff9472 100%);
  border: 1px solid transparent;
}

.gradient-btn > span {
  margin: 0 6px;
  font-weight: 700;
}

.communication .fa {
  font-size: 1.5rem;
}

.header-text {
  margin: auto;
  width: 90%;
}

.reason-text {
  color: black;
  margin-top: 1.25rem;
}

.text-left {
  text-align: left;
}

.bullet {
  content: "";
  background: linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  width: 1rem;
  height: 1rem;
  border-radius: 30%;
  display: inline-block;
  line-height: 1;
  margin: 0.3rem 1.5rem 0 0;
}

.chat-list {
  list-style-type: none;
  margin: 0;
  padding: 0;
}

.chat-list > li {
  padding: 0.6rem 0;
  display: flex;
}

.chat-list > li > div {
  width: 90%;
}

@include media-breakpoint-up(md) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }

  .communication {
    flex-direction: row;
  }

  .communication > .item {
    padding: 2.5rem 1rem;
  }

  .item-divider {
    width: 2px;
    height: auto;
  }

  .chat-btn-container {
    flex-direction: row;
    justify-content: center;
  }

  .communication-item1 {
    flex: 35%;
  }

  .communication-item2 {
    flex: 65%;
  }

  .gradient-btn {
    padding: 1rem 1rem;
    margin: 1.25rem 0.5rem 0 0.5rem;
  }
}

@include media-breakpoint-up(lg) {
  .header-text {
    font-size: 3rem;
    line-height: 3.6rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }

  .cta-container {
    width: 800px;
    margin: 7% auto;
  }
}
</style>
