<template>
  <div class="user-reviews-div normal-text">
    <div class="header-text canopas-gradient-text text-center">
      Listen what real app users <br />
      are saying...
    </div>
    <div class="horizontal-grid">
      <div>
        <div class="horizontal-grid-1 animation" ref="horizontalGrid1">
          <div class="first">
            Thank you, developers, for this amazing app experience.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            It's smooth, I love, cheers to the team behind it.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            The best Moroccan app I have ever seen, clean work .. cheers to
            developers 🤙
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            Slick and elegant. Nice user interface and great content.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            The application is very well done! Keep up guys👏
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            Impeccable app. Great user experience.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            A great application, I like the graphics and the way it work 👍

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
      <div>
        <div class="horizontal-grid-2 animation">
          <div>
            Perfect design, love how classy they made it.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            The app is well designed, giving a very pleasing and smooth User
            Experience. Bravo!
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            I would love to say that there may be something wrong with the app,
            so that I can contribute to its development... It is perfectly
            designed, minimalistic and easy to use. Cheers.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            Keep up the good work you truly deserve more than 5 stars.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            I am surprised by how good this app is !! I love it.
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            The app has made incredible changes, I had it since the beginning,
            always improving !!!
            <div>
              <font-awesome-icon icon="star" class="checked mt-3" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
              <font-awesome-icon icon="star" class="checked" />
            </div>
          </div>
          <div>
            Elegant , splendid well designed app.
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
      margin: 0,
      animationWidth: 0,
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {},
};
</script>

<style scoped>
.user-reviews-div {
  padding: 10% 0 2%;
  color: black;
}

.checked {
  color: #ffc700;
}

.horizontal-grid {
  margin-top: 2rem;
  overflow-x: hidden;
  overflow-y: hidden;
  text-align: left;
  scrollbar-width: none;
  /* white-space: nowrap; */
  animation: marquee 10s linear infinite;
}

.horizontal-grid:hover {
  animation-play-state: paused;
}

.first {
  margin-left: 12% !important;
}

@keyframes marquee {
  0% {
    transform: translate3d(0, 0, 0);
  }
  100% {
    transform: translate3d(-40%, 0, 0);
  }
}

.horizontal-grid::-webkit-scrollbar {
  display: none;
}

.horizontal-grid-1,
.horizontal-grid-2 {
  display: flex;
  width: 100%;
  height: fit-content;
  backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  -ms-backface-visibility: hidden;
}

.horizontal-grid-2 {
  margin-bottom: 2rem;
}

.horizontal-grid-1 > div,
.horizontal-grid-2 > div {
  border: 1px solid #e2e2e2;
  border-radius: 1rem;
  padding: 2rem;
  margin: 2rem 0 0 2rem;
  flex: 0 0 20rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
  transition: transform 250ms linear;
}

.horizontal-grid-1 > div:hover,
.horizontal-grid-2 > div:hover {
  cursor: pointer;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.2) !important;
  transform: scale(0.97);
}

@media (min-width: 481px) {
  .horizontal-grid-1 > div,
  .horizontal-grid-2 > div {
    flex: 0 0 30rem;
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

  .horizontal-grid-1 > div,
  .horizontal-grid-2 > div {
    flex: 0 0 33rem !important;
  }
}
</style>
