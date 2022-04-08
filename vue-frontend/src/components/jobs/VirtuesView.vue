<template>
  <div class="background">
    <div class="normal-text title">Canopas is built on the 10 virtues</div>
    <div class="text-center container">
      <div class="flex-div">
        <div
          v-for="virtue in virtues.slice(0, 3)"
          :key="virtue.id"
          @touchstart="activeIndex = virtue.id"
        >
          <div class="details">
            <div class="image">
              <img :src="virtue.image" :alt="virtue.title" loading="lazy" />
            </div>
            <div class="virtue-text red-text normal-text">
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="description"
            :class="activeIndex == virtue.id ? 'active' : ''"
          >
            <div class="red-text normal-text">
              {{ virtue.title }}
            </div>
            <div class="normal-2-text">
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex-div">
        <div
          v-for="virtue in virtues.slice(3, 5)"
          :key="virtue.id"
          @touchstart="activeIndex = virtue.id"
        >
          <div class="details">
            <div class="image">
              <img :src="virtue.image" :alt="virtue.title" loading="lazy" />
            </div>
            <div class="virtue-text red-text normal-text">
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="description"
            :class="activeIndex == virtue.id ? 'active' : ''"
          >
            <div class="red-text normal-text">
              {{ virtue.title }}
            </div>
            <div class="normal-2-text">
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex-div">
        <div
          v-for="virtue in virtues.slice(5, 8)"
          :key="virtue.id"
          @touchstart="activeIndex = virtue.id"
        >
          <div class="details">
            <div class="image">
              <img :src="virtue.image" :alt="virtue.title" loading="lazy" />
            </div>
            <div class="virtue-text red-text normal-text">
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="description"
            :class="activeIndex == virtue.id ? 'active' : ''"
          >
            <div class="red-text normal-text">
              {{ virtue.title }}
            </div>
            <div class="normal-2-text">
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex-div">
        <div
          v-for="virtue in virtues.slice(8, 10)"
          :key="virtue.id"
          @touchstart="activeIndex = virtue.id"
        >
          <div class="details">
            <div class="image">
              <img :src="virtue.image" :alt="virtue.title" loading="lazy" />
            </div>
            <div class="virtue-text red-text normal-text">
              {{ virtue.title }}
            </div>
          </div>
          <div
            class="description"
            :class="activeIndex == virtue.id ? 'active' : ''"
          >
            <div class="red-text normal-text">
              {{ virtue.title }}
            </div>
            <div class="normal-2-text">
              {{ virtue.description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import virtue1 from "@/assets/images/virtues/jobs_canopas_virtue_1.svg";
import virtue2 from "@/assets/images/virtues/jobs_canopas_virtue_2.svg";
import virtue3 from "@/assets/images/virtues/jobs_canopas_virtue_3.svg";
import virtue4 from "@/assets/images/virtues/jobs_canopas_virtue_4.svg";
import virtue5 from "@/assets/images/virtues/jobs_canopas_virtue_5.svg";
import virtue6 from "@/assets/images/virtues/jobs_canopas_virtue_6.svg";
import virtue7 from "@/assets/images/virtues/jobs_canopas_virtue_7.svg";
import virtue8 from "@/assets/images/virtues/jobs_canopas_virtue_8.svg";
import virtue9 from "@/assets/images/virtues/jobs_canopas_virtue_9.svg";
import virtue10 from "@/assets/images/virtues/jobs_canopas_virtue_10.svg";
export default {
  data() {
    return {
      activeIndex: 0,
      virtues: [
        {
          id: 1,
          title: "We make big bold bets",
          image: virtue1,
          description:
            "Sometimes we fail, but failure makes us smarter. We get back up, we make the next bet, and we go!",
        },
        {
          id: 2,
          title: "We do what matters the most",
          image: virtue2,
          description:
            "Almost always, what we do matters so much compared to how we do it. There's no point doing bad things in a very efficient way.",
        },
        {
          id: 3,
          title: "We tell the truth even if it hurts",
          image: virtue3,
          description:
            "Being honest to others, even if it would hurt them, just shows them how much they can trust you. It’s good be honest, but don’t be rude.",
        },
        {
          id: 4,
          title: "We value time",
          image: virtue4,
          description:
            "'The one that values his time can value the time of others as well'- well said. Always respect others' time even more than yours.",
        },
        {
          id: 5,
          title: "We move fast and break things",
          image: virtue5,
          description:
            "Making mistakes is a natural consequence of innovation in a highly competitive and complex environment..",
        },
        {
          id: 6,
          title: "We celebrate differences",
          image: virtue6,
          description:
            "We stand apart from the average. We ensure people of diverse backgrounds feel welcome. We encourage different opinions and approaches to be heard, and then we come together and build.",
        },
        {
          id: 7,
          title: "We are owners",
          image: virtue7,
          description:
            "We seek out problems and we solve them. We help each other and those who matter to us. We have a bias for action and accountability.",
        },
        {
          id: 8,
          title: "We don't have titles",
          image: virtue8,
          description: "We don't need titles to be true leaders.",
        },
        {
          id: 9,
          title: "We value ideas over hierarchy",
          image: virtue9,
          description:
            "We believe that the best ideas can come from anywhere, both inside and outside our company. Our job is to seek out those ideas, to shape and improve them through candid debate, and to take them from concept to action.",
        },
        {
          id: 10,
          title: "We are playing the infinite game",
          image: virtue10,
          description:
            "An infinite mindset is the recognition that there is no practical end to our work. The goal is not to win or be the best in the infinite game; it is to strive to be better, to experience constant improvement.",
        },
      ],
    };
  },
};
</script>

<style scoped lang="scss">
.background {
  background: linear-gradient(
    247.22deg,
    rgba(255, 173, 147, 0.15) 11.18%,
    rgba(242, 112, 156, 0.15) 80.3%
  );
  padding: 100px 0;
  overflow: hidden;
  text-align: center;
  border-radius: 100px 0;
}

.container {
  margin-top: 80px;
}

.red-text {
  color: #e8505b;
}

.title {
  color: #3d3d3d;
  font-weight: 700;
}

.flex-div {
  display: flex;
  flex-flow: wrap;
  margin: auto;
}

.flex-div > div {
  position: relative;
  flex: 100%;
}

.details {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 5% auto;
  padding: 0 10%;

  div {
    margin-bottom: 5%;
  }

  .virtue-text {
    width: 60%;
  }
}

.description {
  position: absolute;
  height: auto;
  width: 290px;
  left: 50%;
  top: 0;
  bottom: 0;
  opacity: 0;
  background: white;
  padding: 2%;
  margin: auto;
  border-radius: 15px;
  display: flex;
  flex-direction: column;
  justify-content: center;

  transform: translateX(-50%);
  -webkit-transition: all 0.4s ease-in-out 0s;
  -moz-transition: all 0.4s ease-in-out 0s;
  transition: all 0.4s ease-in-out 0s;

  div {
    margin-bottom: 3%;
    width: 100%;
  }
}

.active {
  opacity: 1;
}

@media (hover: hover) and (pointer: fine) {
  .flex-div > div:hover .description {
    opacity: 1;
  }
}

.image {
  border-radius: 50%;
  background: white;
  position: relative;
  height: 0;
  width: 50%;
  padding-bottom: 50%;

  img {
    position: absolute;
    width: 95%;
    height: 95%;
    left: 0;
  }
}

@include media-breakpoint-up(sm) {
  .container {
    padding: 0%;
  }

  .background {
    border-radius: 150px 0;
  }

  .description {
    width: 50%;
  }

  .details {
    margin: 2% 0;
    padding: 0;
  }

  .image {
    width: 30%;
    padding-bottom: 30%;
  }
}

@include media-breakpoint-up(md) {
  .background {
    border-radius: 200px 0;
  }

  .flex-div:nth-of-type(even) {
    width: 70%;
  }

  .flex-div > div {
    flex: 33%;
  }

  .description {
    width: 90%;

    div {
      margin-bottom: 0;
    }
  }

  .details {
    margin: 10% 0;
  }

  .image {
    width: 50%;
    padding-bottom: 50%;
  }

  .normal-2-text {
    margin-top: 10px;
  }
}

@include media-breakpoint-up(lg) {
  .background {
    border-radius: 300px 0;
  }

  .flex-div {
    width: 85%;
  }

  .flex-div:nth-of-type(even) {
    width: 60%;
  }

  .red-text {
    font-size: 1.3125rem;
    line-height: 1.5918rem;
  }
}

@include media-breakpoint-up(xxl) {
  .description {
    padding: 3%;

    div {
      margin-bottom: 2%;
    }
  }
}
</style>
