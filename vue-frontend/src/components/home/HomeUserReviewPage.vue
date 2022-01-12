<template>
  <div class="user-reviews-div normal-text">
    <div class="header-text canopas-gradient-text text-center">
      Listen what real app users <br />
      are saying...
    </div>
    <div class="horizontal-grid" ref="horizontalGrid">
      <div class="grid-animation">
        <div class="horizontal-grid-1" ref="horizontalGrid1">
          <div
            v-for="userReview in usersReviewsGrid1"
            :key="userReview.id"
            class="users-review"
          >
            {{ userReview.review }}
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div
            v-for="userReview in usersReviewsGrid1"
            :key="userReview.id"
            class="users-review"
          >
            {{ userReview.review }}
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
        </div>
        <div class="horizontal-grid-2" ref="horizontalGrid2">
          <div
            v-for="userReview in usersReviewsGrid2"
            :key="userReview.id"
            class="users-review"
          >
            {{ userReview.review }}
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div
            v-for="userReview in usersReviewsGrid2"
            :key="userReview.id"
            class="users-review"
          >
            {{ userReview.review }}
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      usersReviewsGrid1: [
        {
          id: 1,
          review: "Thank you, developers, for this amazing app experience.",
        },
        {
          id: 2,
          review: "It's smooth, I love, cheers to the team behind it.",
        },
        {
          id: 3,
          review:
            "The best Moroccan app I have ever seen, clean work .. cheers to developers 🤙",
        },
        {
          id: 4,
          review: "Slick and elegant. Nice user interface and great content.",
        },
        {
          id: 5,
          review: "The application is very well done! Keep up guys👏",
        },
        {
          id: 6,
          review: "Impeccable app. Great user experience.",
        },
        {
          id: 7,
          review:
            "A great application, I like the graphics and the way it work 👍",
        },
      ],
      usersReviewsGrid2: [
        {
          id: 1,
          review: "Perfect design, love how classy they made it.",
        },
        {
          id: 2,
          review:
            "The app is well designed, giving a very pleasing and smooth User Experience. Bravo!",
        },
        {
          id: 3,
          review:
            "I would love to say that there may be something wrong with the app, so that I can contribute to its development... It is perfectly designed, minimalistic and easy to use. Cheers.",
        },
        {
          id: 4,
          review: "Keep up the good work you truly deserve more than 5 stars.",
        },
        {
          id: 5,
          review: "I am surprised by how good this app is !! I love it.",
        },
        {
          id: 6,
          review:
            "The app has made incredible changes, I had it since the beginning, always improving !!!",
        },
        {
          id: 7,
          review: "Elegant , splendid well designed app.",
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {},
  methods: {},
};
</script>

<style scoped>
.user-reviews-div {
  padding: 10% 0 2%;
  color: black;
}

.header-text {
  font-weight: 700;
  line-height: 1.813rem;
  font-size: 1.5rem;
}

.checked {
  color: #ffc700;
}

.horizontal-grid {
  margin-top: 32px;
  overflow-x: hidden;
  overflow-y: auto;
  text-align: left;
  scrollbar-width: none;
}

.horizontal-grid::-webkit-scrollbar {
  display: none;
}

.grid-animation {
  margin-bottom: 32px;
  animation: scroll 30s linear infinite;
}

.grid-animation:hover {
  animation-play-state: paused;
}

.horizontal-grid-1,
.horizontal-grid-2 {
  display: flex;
  width: 100%;
  height: fit-content;
}

.horizontal-grid-1 {
  margin-left: 160px;
}

.users-review {
  border: 1px solid #e2e2e2;
  border-radius: 16px;
  padding: 32px;
  margin: 32px 0 0 32px;
  flex: 0 0 320px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  transform: translate3d(0, 0, 0);
}

.users-review:hover {
  cursor: pointer;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.2) !important;
  transform: scale(0.97);
}

@keyframes scroll {
  /*
     We animate from 3.5 review cards to 10.5 review card. We have total 14 cards, so user will 
     see animation as infinite animation.
     The review cards width + margin - multiplied by numbers of gridView cards
     = (320 + 32) * 3.5  to (320 + 32) * 10.5
     */
  0% {
    transform: translate(-1232px);
  }

  100% {
    transform: translate(-3696px);
  }
}

@media (min-width: 481px) {
  .horizontal-grid-1 {
    margin-left: 240px;
  }

  .users-review {
    flex: 0 0 480px;
  }

  @keyframes scroll {
    0% {
      transform: translate(-1792px);
    }

    100% {
      transform: translate(-5376px);
    }
  }
}

@media (min-width: 768px) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }
}

@media (min-width: 992px) {
  .header-text {
    font-size: 3rem;
    line-height: 3.6rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }

  .horizontal-grid-1 {
    margin-left: 264px;
  }

  .users-review {
    flex: 0 0 528px;
  }

  @keyframes scroll {
    0% {
      transform: translate(-1960px);
    }

    100% {
      transform: translate(-5880px);
    }
  }
}
</style>
