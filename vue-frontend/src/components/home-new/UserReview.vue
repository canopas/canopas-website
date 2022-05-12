<template>
  <div class="container overlay">
    <div class="v2-header-3-text text-center">
      We say our products are
      <span class="v2-canopas-gradient-text">amazing</span>
      but
      <span class="v2-canopas-gradient-text"> user reviews </span>confirm it.
    </div>

    <LottieAnimation :jsonData="backgroundImage" class="background-image" />
    <div>
      <div class="flex-div">
        <div
          v-for="userReview in usersReviews.slice(0, 3)"
          :key="userReview.id"
          class="users-review v2-normal-3-text"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-3">
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
          </div>
        </div>
      </div>
      <div v-if="slpitReview" class="flex-div">
        <div
          v-for="userReview in usersReviews.slice(3, 5)"
          :key="userReview.id"
          class="users-review v2-normal-3-text"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-3">
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
          </div>
        </div>
      </div>
      <div v-if="slpitReview" class="flex-div">
        <div
          v-for="userReview in usersReviews.slice(5, 7)"
          :key="userReview.id"
          class="users-review v2-normal-3-text"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-3">
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
          </div>
        </div>
      </div>
      <div v-else class="flex-div">
        <div
          v-for="userReview in usersReviews.slice(3, 7)"
          :key="userReview.id"
          class="users-review v2-normal-3-text"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-3">
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
          </div>
        </div>
      </div>
      <div class="flex-div">
        <div
          v-for="userReview in usersReviews.slice(7, 10)"
          :key="userReview.id"
          class="users-review v2-normal-3-text"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-3">
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
            <font-awesome-icon icon="star" class="checked" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import backgroundImage from "@/assets/lottie/globe.json";
import LottieAnimation from "@/components/utils/LottieAnimation.vue";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      usersReviews: [
        {
          id: 1,
          review: "It's smooth, I love, cheers to the team behind it.",
        },
        {
          id: 2,
          review:
            "The best Moroccan app I have ever seen, clean work.. cheers to developers 🤙",
        },
        {
          id: 3,
          review: "Elegant , splendid well designed app.",
        },
        {
          id: 4,
          review: "I am surprised by how good this app is !! I love it. ",
        },
        {
          id: 5,
          review:
            "The app is well designed, giving a very pleasing and smooth User Experience. Bravo!",
        },
        {
          id: 6,
          review: "Perfect design, love how classy they made it",
        },
        {
          id: 7,
          review: "Slick and elegant.Nice user interface and great content.",
        },
        {
          id: 8,
          review: "Thank you, developers, for this amazing app experience.",
        },
        {
          id: 9,
          review: "Impeccable app. Great user experience. ",
        },
        {
          id: 10,
          review:
            "The app has made incredible changes, I had it since the beginning, always improving !!! ",
        },
      ],
      slpitReview: false,
    };
  },
  components: {
    FontAwesomeIcon,
    LottieAnimation,
  },
  mounted() {
    if (window.innerWidth >= 768 && window.innerWidth <= 992) {
      this.slpitReview = true;
    }
  },
};
</script>

<style lang="scss" scoped>
.container {
  position: relative;
}

.overlay {
  padding: 150px 0 250px;
}

.background-image {
  display: none;
}

.v2-header-3-text {
  margin: auto;
}

.flex-div {
  display: flex;
  flex-flow: wrap;
  margin: auto;

  div {
    flex: 100%;
  }
}

.users-review {
  border: 1.5px solid #fa7b78 !important;
  border-radius: 16px;
  padding: 20px;
  margin: 15px 15px 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  transform: translate3d(0, 0, 0);
}

.users-review:hover {
  cursor: pointer;
  box-shadow: 0px 1px 6px 2px #f98188;
}

.checked {
  color: #ffc700;
}

@include media-breakpoint-up(sm) {
  .v2-header-3-text {
    width: 80%;
  }
}

@include media-breakpoint-up(md) {
  .background-image {
    display: block;
    height: 70%;
    width: 100%;
    position: absolute;
    object-fit: cover;
    z-index: -5;
    margin-left: auto;
    margin-right: auto;
    opacity: 40%;
  }

  .flex-div {
    align-items: center;

    div {
      flex: 20%;
    }
  }
  .users-review {
    padding: 20px 15px;
    margin: 10px;
    text-align: center;
  }

  .flex-div:nth-of-type(1) {
    width: 90%;

    .users-review:nth-of-type(1) {
      margin-top: 15%;
    }

    .users-review:nth-of-type(3) {
      margin-top: 12%;
    }
  }

  .flex-div:nth-of-type(2) {
    width: 100%;

    .users-review:nth-of-type(1) {
      margin-top: -8%;
    }

    .users-review:nth-of-type(2) {
      margin-top: 6%;
    }
  }

  .flex-div:nth-of-type(3) {
    width: 80%;

    .users-review:nth-of-type(1) {
      margin-top: -17%;
    }
  }

  .flex-div:nth-of-type(4) {
    width: 95%;

    .users-review:nth-of-type(1) {
      margin-top: -17%;
    }

    .users-review:nth-of-type(2) {
      margin-top: -5%;
    }
  }
}

@include media-breakpoint-up(lg) {
  .flex-div:nth-of-type(1) {
    width: 80%;

    .users-review:nth-of-type(1) {
      margin-top: 15%;
    }

    .users-review:nth-of-type(3) {
      margin-top: 12%;
    }
  }

  .flex-div:nth-of-type(2) {
    width: 100%;

    .users-review:nth-of-type(1) {
      margin-top: -4%;
    }

    .users-review:nth-of-type(2) {
      margin-top: 0;
    }

    .users-review:nth-of-type(3) {
      margin-top: -8%;
    }

    .users-review:nth-of-type(4) {
      margin-top: 6%;
    }
  }

  .flex-div:nth-of-type(3) {
    width: 75%;
    margin-left: 20%;

    .users-review:nth-of-type(1) {
      margin-top: -5%;
    }

    .users-review:nth-of-type(2) {
      margin-top: -5%;
    }
  }
}

@include media-breakpoint-up(xl) {
  .flex-div:nth-of-type(1) {
    width: 70%;
  }

  .flex-div:nth-of-type(2) {
    width: 90%;
  }

  .flex-div:nth-of-type(3) {
    width: 70%;
  }
}
</style>
