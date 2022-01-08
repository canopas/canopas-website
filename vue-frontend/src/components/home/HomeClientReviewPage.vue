<template>
  <div class="success-stories-bg text-center">
    <div>
      <div class="header-text canopas-gradient-text">
        Success stories of <br />
        people like You!
      </div>
      <div class="horizontal-slider normal-text text-center">
        <div
          id="clientsIndicators"
          class="carousel slide"
          ref="carousel"
          data-bs-ride="carousel"
        >
          <ol class="clients__list text-center">
            <li
              data-bs-target="#clientsIndicators"
              data-bs-slide-to="0"
              class="clients__list__item gradient-circular-border"
            >
              <img
                src="@/assets/images/clients/lisa.jpg"
                alt="client"
                draggable="false"
              />
            </li>
            <li
              data-bs-target="#clientsIndicators"
              data-bs-slide-to="1"
              class="clients__list__item gradient-circular-border"
            >
              <img
                src="@/assets/images/clients/marcus.jpg"
                alt="client"
                draggable="false"
              />
            </li>
            <li
              data-bs-target="#clientsIndicators"
              data-bs-slide-to="2"
              class="clients__list__item gradient-circular-border"
            >
              <img
                src="@/assets/images/clients/jake.jpg"
                alt="client"
                draggable="false"
              />
            </li>
            <li
              data-bs-target="#clientsIndicators"
              data-bs-slide-to="3"
              class="clients__list__item gradient-circular-border"
            >
              <img
                src="@/assets/images/clients/maor.jpg"
                alt="client"
                draggable="false"
              />
            </li>
            <li
              data-bs-target="#clientsIndicators"
              data-bs-slide-to="4"
              class="clients__list__item gradient-circular-border"
            >
              <img
                src="@/assets/images/clients/ramasis.jpg"
                alt="client"
                draggable="false"
              />
            </li>
            <li
              data-bs-target="#clientsIndicators"
              data-bs-slide-to="5"
              class="clients__list__item gradient-circular-border"
            >
              <img
                src="@/assets/images/clients/jake.jpg"
                alt="client"
                draggable="false"
              />
            </li>
          </ol>
          <div class="carousel-inner">
            <div class="carousel-item">
              <div class="client">
                <div class="normal-text canopas-gradient-text mt-4">
                  Lisa W.
                </div>
                <div class="normal-text text-left mt-4">
                  There is not enough space to say all the wonderful things I
                  would want to share about Canopas. They are incredibly
                  helpful, stay calm even when we had to deal with tough issues
                  on our app, and always found a way to help us fix whatever was
                  needed or roll out any new features for our app in both the
                  iOS and Android stores. I will absolutely find a way to work
                  with Canopas again!
                </div>
              </div>
            </div>
            <div class="carousel-item">
              <div class="client">
                <div class="normal-text canopas-gradient-text mt-4">
                  Marcus L.
                </div>
                <div class="normal-text text-left mt-4">
                  Canopas has been nothing but wonderful on this project. His
                  communication and ability to advise on the best solutions
                  throughout the project were top-notch. I have very much
                  enjoyed working with Darpan and his team and would surely
                  recommend them to anyone looking for a solid developer.
                </div>
              </div>
            </div>
            <div class="carousel-item active">
              <div class="client">
                <div class="normal-text canopas-gradient-text mt-4">
                  Jake N.
                </div>
                <div class="normal-text text-left mt-4">
                  Canopas team was also incredibly kind and always willing to
                  solve any problem through research something that we highly
                  appreciate. Many people do not take the time to understand a
                  problem but Canopas did it every time and found a way to fix
                  it. They even thought about the User Experience and the design
                  which not many engineers think about. We are so lucky to have
                  worked with them and look forward to working with them again
                  in the future!
                </div>
              </div>
            </div>
            <div class="carousel-item">
              <div class="client">
                <div class="normal-text canopas-gradient-text mt-4">
                  Maor T.
                </div>
                <div class="normal-text text-left mt-4">
                  This is our favorite expert for all mobile and web developing
                  areas. Darpan has a great team who get the job done above our
                  expectations. We would like to hire Canopas for more projects
                  in the future and highly recommend their kind and professional
                  services.
                </div>
              </div>
            </div>
            <div class="carousel-item">
              <div class="client">
                <div class="normal-text canopas-gradient-text mt-4">
                  Ramsis A.
                </div>
                <div class="normal-text text-left mt-4">
                  Canopas has been great to work with. From day 1, they made
                  sure they understood exactly what I wanted and advised and
                  guided me during the process (with patience). This project has
                  ended, but I am sure to work with them again on other
                  projects. I would recommend their services.
                </div>
              </div>
            </div>
            <div class="carousel-item">
              <div class="client">
                <div class="normal-text canopas-gradient-text mt-4">
                  Jake N.
                </div>
                <div class="normal-text text-left mt-4">
                  Canopas team was unbelievable. They did everything and above.
                  We had initially hired them just for engineering the app, but
                  they were able to jump between app development, web
                  development, different integrations, and more. It was really
                  impressive, and we were so lucky to work with them. They
                  thought about everything, even things we did not bring up.
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-4 client-arrow">
        <button type="button" class="clients-indicators" @click="leftIndicator">
          <font-awesome-icon class="arrow" icon="arrow-left" id="leftArrow" />
        </button>
        <button
          type="button"
          class="clients-indicators"
          @click="rightIndicator"
        >
          <font-awesome-icon class="arrow" icon="arrow-right" id="rightArrow" />
        </button>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import * as Bootstrap from "bootstrap";

export default {
  data() {
    return {
      totalLength: null,
      direction: "left",
      index: 0,
      clicked: false,
      clientList: null,
      clientListItem: null,
      carouselElem: null,
      carousel: null,
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted: function () {
    this.totalLength =
      document.querySelector("#clientsIndicators .carousel-item").length - 1;
    this.clientList = document.querySelector(".clients__list");
    this.clientListItem = document.querySelector(".clients__list__item");

    this.carouselElem = this.$refs.carousel;
    this.carousel = new Bootstrap.Carousel(this.carouselElem, {
      interval: false,
      touch: false,
    });

    this.carouselSlider();
  },
  methods: {
    carouselSlider() {
      this.carouselElem.addEventListener("data-bs-slide-to", (e) => {
        console.log("call slied to function");
        this.slideTo(e, this.direction);
      });

      document.querySelector(".clients__list__item").onclick = () => {
        this.myFunction();
      };
    },
    myFunction() {
      console.log(this);
      this.clicked = true;
      this.index = this.index();
      this.direction = this.index > this.totalLength / 2 ? "left" : "right";
      this.carousel.to(parseInt(this.attr("data-bs-slide-to")));
    },
    leftIndicator() {
      this.clicked = false;
      this.direction = "left";
      this.carousel.prev();
    },
    rightIndicator() {
      this.clicked = false;
      this.direction = "right";
      this.carousel.next();
    },
    slideTo(event, direction) {
      var indicators = document.querySelector(".clients__list");
      var items = document.querySelector(".carousel-item");
      var active = document.querySelector(".carousel-item.active");
      var item = event.relatedTarget;
      var itemIndex = item.index();
      console.log(itemIndex);

      if (typeof direction === "undefined") direction = "left";

      if (event.direction !== direction) {
        event.preventDefault();

        if (direction === "right") {
          item.insertBefore(active);
          this.carousel.next();

          setTimeout(function () {
            if (itemIndex === items.length - 1) {
              item.insertAfter(items.eq(itemIndex - 1));
            } else {
              item.insertBefore(items.eq(itemIndex + 1));
            }
          }, 600);
        } else {
          item.insertAfter(active);
          this.carousel.prev();

          setTimeout(function () {
            if (itemIndex === 0) {
              item.insertBefore(items.eq(1));
            } else {
              item.insertAfter(items.eq(itemIndex - 1));
            }
          }, 600);
        }

        indicators.find(".active").classList.remove("active");
        indicators.children().eq(itemIndex).classList.add("active");
      }
    },
  },
};
</script>

<style scoped>
.success-stories-bg {
  background: rgba(255, 148, 114, 0.05);
  padding: 6% 0;

  backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
}

.horizontal-slider {
  margin: 3rem auto 0;
  width: 60%;
}

.clients__list {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 10rem;
  margin: 0 auto;
  padding-left: 0 !important;
}

.clients__list > li:last-child {
  display: none;
}

.clients__list__item {
  overflow: hidden;
  border-radius: 50%;
  position: relative;
  -webkit-mask: none;
  cursor: pointer;
  margin: 0 1rem;
}

.clients__list__item > img {
  border-radius: 50%;
  margin: 0 auto;
  width: 100%;
  height: 100%;
  border: 5px solid transparent;
}

.clients__list__item:first-child {
  opacity: 0.15;
  width: 9%;
}

.clients__list__item:nth-child(2) {
  opacity: 0.4;
  width: 12%;
}

.clients__list__item:nth-child(3) {
  opacity: 1;
  width: 15%;
}

.clients__list__item:nth-child(3) > img {
  pointer-events: none;
}

.clients__list__item:nth-child(4) {
  opacity: 0.4;
  width: 12%;
}

.clients__list__item:nth-child(5) {
  opacity: 0.15;
  width: 9%;
}

.clients__list__item:last-child {
  opacity: 0.15;
  width: 9%;
}

.gradient-circular-border {
  content: "";
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 2px;
  background: linear-gradient(180deg, #ff9472 0%, #f2709c 100%);
  -webkit-mask-composite: destination-out;
  mask-composite: exclude !important;
}

.carousel-inner {
  height: 19rem;
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

.pushleft {
  animation: pushleft 0.6s;
}

.pushright {
  animation: pushright 0.6s;
}

@-webkit-keyframes pushleft {
  0% {
    -webkit-transform: translateX(10%);
  }
  100% {
    -webkit-transform: translateX(0);
  }
}

@-webkit-keyframes pushright {
  0% {
    -webkit-transform: translateX(-10%);
  }
  100% {
    -webkit-transform: translateX(0);
  }
}

@keyframes expand {
  from {
    transform: scale(0.8);
    opacity: 0.15;
  }
}

@keyframes expand-in {
  from {
    transform: scale(1.1);
    opacity: 0.4;
  }
}

/* responsive */

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

  .carousel-inner {
    height: 15rem;
  }

  .client {
    width: 40rem;
    margin: 2rem auto;
  }

  .clients__list__item:first-child {
    opacity: 0.15;
    width: 9%;
  }

  .clients__list__item:last-child {
    opacity: 0.15;
    width: 9%;
  }
}

@media (min-width: 1025px) {
  .success-stories-bg {
    border-radius: 25rem 0;
  }

  .carousel-inner {
    height: 20rem;
  }
}
</style>
