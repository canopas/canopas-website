<template>
  <div class="container overlay">
    <div class="v2-header-3-text text-center title-width">
      We say our products are
      <span class="v2-canopas-gradient-text">amazing</span>
      but
      <span class="v2-canopas-gradient-text">user reviews </span>confirm it.
    </div>

    <div v-if="backgroundImage !== null">
      <LottieAnimation :jsonData="backgroundImage" class="background-image" />
    </div>

    <div class="pt-4">
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
import { defineAsyncComponent } from "vue";

export default {
  data() {
    return {
      backgroundImage: null,
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
    LottieAnimation: defineAsyncComponent(() =>
      import("@/components/LottieAnimation.vue")
    ),
  },
  mounted() {
    if (window.innerWidth >= 768 && window.innerWidth <= 992) {
      this.slpitReview = true;
    }
    if (window.innerWidth < 768) {
      this.usersReviews = this.usersReviews.slice(0, 6);
    } else {
      import("@/assets/lottie/globe.json").then((data) => {
        this.backgroundImage = data.default;
      });
    }
  },
};
</script>

<style lang="scss" scoped>
.container {
  position: relative;
}

.overlay {
  padding-top: 50px;
  padding-bottom: 80px;
}

.background-image {
  display: none;
}

.v2-header-3-text {
  margin: auto;
}

.title-width {
  width: 80%;
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
  .title-width {
    width: 100%;
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

  .overlay {
    padding-top: 100px;
    padding-bottom: 150px;
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

  .title-width {
    width: 80%;
  }

  .overlay {
    padding-top: 150px;
    padding-bottom: 250px;
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
@supports (-webkit-touch-callout: none) {
  .v2-header-3-text {
    letter-spacing: -2px;
  }
}
</style>
