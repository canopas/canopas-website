<template>
  <div class="success-stories-bg text-center">
    <div class="container">
      <div class="header-text canopas-gradient-text">
        Success stories of <br />
        people like You!
      </div>
      <div class="horizontal-slider normal-text text-center">
        <div class="clients-list-wrapper">
          <transition-group
            :name="'client-image-' + reviewTransitionName"
            tag="ol"
            class="clients-list text-center"
          >
            <div
              v-for="(client, i) in currentClients"
              :key="client"
              class="client-image-item"
              :class="
                'client-image-item-' + Math.abs(i - (clientImageCount - 1) / 2)
              "
            >
              <img
                :src="client.image"
                alt="client"
                draggable="false"
                loading="lazy"
              />
            </div>
          </transition-group>
        </div>
        <div class="client-review-slider">
          <transition-group tag="div" :name="'review-' + reviewTransitionName">
            <div class="client-review" :key="clients[current].id">
              <div class="normal-text canopas-gradient-text">
                {{ clients[current].name }}
              </div>
              <div class="normal-text text-left mt-4">
                {{ clients[current].review }}
              </div>
            </div>
          </transition-group>
        </div>
      </div>
      <div class="mt-4">
        <button type="button" class="clients-indicators" @click="slide(-1)">
          <font-awesome-icon class="arrow" icon="arrow-left" id="leftArrow" />
        </button>
        <button type="button" class="clients-indicators" @click="slide(1)">
          <font-awesome-icon class="arrow" icon="arrow-right" id="rightArrow" />
        </button>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      clients: [
        {
          id: "client-1",
          image: require("@/assets/images/clients/lisa.webp"),
          name: "Lisa W.",
          review:
            "There is not enough space to say all the wonderful things I\
                  would want to share about Canopas. They are incredibly\
                  helpful, stay calm even when we had to deal with tough issues\
                  on our app, and always found a way to help us fix whatever was\
                  needed or roll out any new features for our app in both the\
                  iOS and Android stores. I will absolutely find a way to work\
                  with Canopas again!",
        },
        {
          id: "client-2",
          image: require("@/assets/images/clients/marcus.webp"),
          name: "Marcus L.",
          review:
            "Canopas has been nothing but wonderful on this project. His\
                  communication and ability to advise on the best solutions\
                  throughout the project were top-notch. I have very much\
                  enjoyed working with Darpan and his team and would surely\
                  recommend them to anyone looking for a solid developer.",
        },
        {
          id: "client-3",
          image: require("@/assets/images/clients/jake.webp"),
          name: "Jake N.",
          review:
            "Canopas team was also incredibly kind and always willing to\
                  solve any problem through research something that we highly\
                  appreciate. Many people do not take the time to understand a\
                  problem but Canopas did it every time and found a way to fix\
                  it. They even thought about the User Experience and the design\
                  which not many engineers think about. We are so lucky to have\
                  worked with them and look forward to working with them again\
                  in the future!",
        },
        {
          id: "client-4",
          image: require("@/assets/images/clients/maor.webp"),
          name: "Maor T.",
          review:
            "This is our favorite expert for all mobile and web developing\
                  areas. Darpan has a great team who get the job done above our\
                  expectations. We would like to hire Canopas for more projects\
                  in the future and highly recommend their kind and professional\
                  services.",
        },
        {
          id: "client-5",
          image: require("@/assets/images/clients/ramasis.webp"),
          name: "Ramsis A.",
          review:
            "Canopas has been great to work with. From day 1, they made\
                  sure they understood exactly what I wanted and advised and\
                  guided me during the process (with patience). This project has\
                  ended, but I am sure to work with them again on other\
                  projects. I would recommend their services.",
        },
        {
          id: "client-6",
          image: require("@/assets/images/clients/jake.webp"),
          name: "Jake N.",
          review:
            "Canopas team was unbelievable. They did everything and above.\
                  We had initially hired them just for engineering the app, but\
                  they were able to jump between app development, web\
                  development, different integrations, and more. It was really\
                  impressive, and we were so lucky to work with them. They\
                  thought about everything, even things we did not bring up.",
        },
      ],
      current: 0,
      currentClients: [],
      reviewTransitionName: "",
      clientImageCount: 0,
    };
  },
  components: {
    FontAwesomeIcon,
  },
  methods: {
    slide(dir) {
      dir === 1
        ? (this.reviewTransitionName = "next")
        : (this.reviewTransitionName = "prev");
      this.current = this.getRoundedIndex(dir);
      this.refreshCurrentClients();
    },
    getRoundedIndex(diff) {
      var len = this.clients.length;
      return (this.current + (diff % len) + len) % len;
    },
    refreshCurrentClients() {
      if (this.clientImageCount == 3) {
        this.currentClients.splice(
          0,
          3,
          this.clients[this.getRoundedIndex(-1)],
          this.clients[this.getRoundedIndex(0)],
          this.clients[this.getRoundedIndex(1)]
        );
      } else {
        this.currentClients.splice(
          0,
          5,
          this.clients[this.getRoundedIndex(-2)],
          this.clients[this.getRoundedIndex(-1)],
          this.clients[this.getRoundedIndex(0)],
          this.clients[this.getRoundedIndex(1)],
          this.clients[this.getRoundedIndex(2)]
        );
      }
    },
  },
  mounted() {
    if (window.innerWidth >= 992) {
      this.clientImageCount = 5;
      this.currentClients.splice(0, 0, "", "", "", "", "");
    } else {
      this.clientImageCount = 3;
      this.currentClients.splice(0, 0, "", "", "");
    }
    this.refreshCurrentClients();
  },
};
</script>

<style lang="scss" scoped>
.success-stories-bg {
  background: rgba(255, 148, 114, 0.05);
  padding: 6% 0;
  backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
  overflow: hidden;
}

.horizontal-slider {
  margin: 48px auto 0;
  width: 85%;
}

.clients-list-wrapper {
  position: relative;
  width: 100%;
  padding-top: 24%;
}

.clients-list {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  padding-left: 0 !important;
}

.client-image-item {
  border-radius: 50%;
  position: relative;
  cursor: pointer;
  margin: 0 16px;
  background: linear-gradient(180deg, #ff9472 0%, #f2709c 100%);
}

.client-image-item > img {
  border-radius: 50%;
  margin: 0 auto;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border: 5px solid transparent;
}

.client-image-item-0 {
  opacity: 1;
  width: 24%;
  padding-top: 24%;
}

.client-image-item-1 {
  opacity: 0.4;
  width: 21%;
  padding-top: 21%;
}

.client-image-item-2 {
  opacity: 0.15;
  width: 18%;
  padding-top: 18%;
}

.client-image-next-enter-active {
  transform: translate(120%);
}
.client-image-next-enter-to {
  transform: translate(0%);
}
.client-image-next-leave-active {
  display: none;
}

.client-image-prev-enter-active {
  transform: translate(-120%);
}
.client-image-prev-enter-to {
  transform: translate(0%);
}
.client-image-prev-leave-active {
  display: none;
}

.clients-indicators {
  background: none;
  border: none;
}

.arrow {
  border: 1px solid rgba(61, 61, 61, 0.15);
  border-radius: 15px;
  height: 45px;
  width: 45px;
  padding: 10px;
  color: #f2709c;
}

.client-image-item {
  transition: all 0.6s ease-in-out;
}

.client-review-slider {
  margin: 0 auto;
  position: relative;
  height: 430px;
  width: 100%;
}

.client-review {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  right: 0;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* GO TO NEXT SLIDE */
.review-next-enter-active,
.review-next-leave-active {
  transition: transform 0.5s ease-in-out;
}
.review-next-enter-active {
  transform: translate(150%);
}
.review-next-enter-to {
  transform: translate(0%);
}
.review-next-leave-to {
  transform: translate(-150%);
}

/* GO TO PREVIOUS SLIDE */
.review-prev-enter-active,
.review-prev-leave-active {
  transition: transform 0.5s ease-in-out;
}
.review-prev-enter-active {
  transform: translate(-150%);
}
.review-prev-enter-to {
  transform: translate(0%);
}
.review-prev-leave-to {
  transform: translate(150%);
}

@include media-breakpoint-up(sm) {
  .client-review-slider {
    height: 350px;
  }
}

@include media-breakpoint-up(md) {
  .client-review {
    width: 70%;
  }

  .clients-list-wrapper {
    padding-top: 15%;
  }

  .client-image-item-0 {
    opacity: 1;
    width: 15%;
    padding-top: 15%;
  }

  .client-image-item-1 {
    opacity: 0.4;
    width: 12%;
    padding-top: 12%;
  }

  .client-image-item-2 {
    opacity: 0.15;
    width: 9%;
    padding-top: 9%;
  }
}

@include media-breakpoint-up(lg) {
  .horizontal-slider {
    width: 60%;
  }

  .client-review {
    width: 100%;
  }

  .client-review-slider {
    width: 640px;
    overflow: hidden;
  }
}

@include media-breakpoint-up(xl) {
  .success-stories-bg {
    border-radius: 400px 0;
  }
}
</style>
