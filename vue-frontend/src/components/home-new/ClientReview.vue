<template>
  <div class="container-fluid overlay">
    <img :src="backgroundImage" class="render-background" loading="lazy" />
    <div class="container">
      <div>
        <div class="flex-div">
          <div>
            <div class="v2-header-3-text title">
              Success <span class="border-text">stories</span> of people like
              you.
            </div>
            <div
              v-for="client in clients.slice(0, 2)"
              :key="client"
              class="client-details"
              :style="{ background: client.bgColor }"
            >
              <div
                class="v2-normal-3-text"
                :style="{ color: client.fontColor }"
              >
                {{ client.review }}
              </div>
              <div class="client-image-div mt-4">
                <img :src="client.image" alt="client" loading="lazy" />
                <div class="v2-title-3-text canopas-gradient-text">
                  {{ client.name }}
                </div>
              </div>
            </div>
          </div>
          <div>
            <div
              v-for="client in clients.slice(2, 5)"
              :key="client"
              class="client-details"
              :style="{ background: client.bgColor }"
            >
              <div
                class="v2-normal-3-text"
                :style="{ color: client.fontColor }"
              >
                {{ client.review }}
              </div>
              <div class="client-image-div mt-4">
                <img :src="client.image" alt="client" loading="lazy" />
                <div class="v2-title-3-text canopas-gradient-text">
                  {{ client.name }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import backgroundImage from "@/assets/images/clients/v2-client/quote-mark.svg";
import marcus from "@/assets/images/clients/v2-client/marcus.png";
import jake from "@/assets/images/clients/v2-client/jake.png";
import lisa from "@/assets/images/clients/v2-client/lisa.png";
import maor from "@/assets/images/clients/v2-client/maor.png";
import ramasis from "@/assets/images/clients/v2-client/ramasis.png";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      clients: [
        {
          id: "client-1",
          image: marcus,
          name: "Marcus L.",
          fontColor: "#FFFFFF",
          bgColor: "#3D3D3D",
          review:
            "Canopas has been nothing but wonderful on this project. His\
              communication and ability to advise on the best solutions\
              throughout the project were top-notch. I have very much\
              enjoyed working with Darpan and his team and would surely\
              recommend them to anyone looking for a solid developer.",
        },
        {
          id: "client-2",
          image: jake,
          name: "Jake N.",
          fontColor: "#FFFFFF",
          bgColor: "#3D3D3D",
          review:
            "Canopas team was also incredibly kind and always willing\
              tosolve any problem through research something that we highly\
              appreciate. Many people do not take the time to understand a\
              problem but Canopas did it every time and found a way to fix\
              it. They even thought about the User Experience and the\
              design which not many engineers think about. ",
        },
        {
          id: "client-3",
          image: lisa,
          name: "Lisa W.",
          fontColor: "#3D3D3D",
          bgColor:
            "linear-gradient(270.11deg, rgba(249, 131, 134, 0.1) -39.34%, rgba(249, 131, 134, 0.1) 49.77%);",
          review:
            "There is not enough space to say all the wonderful things I\
              would want to share about Canopas. They are incredibly helpful,\
              stay calm even when we had to deal with tough issues on our\
              app, and always found a way to help us fix whatever was needed\
              or roll out any new features for our app in both the iOS and\
              Android stores. I will absolutely find a way to work with\
              Canopas again!",
        },
        {
          id: "client-4",
          image: maor,
          name: "Maor T.",
          fontColor: "#FFFFFF",
          bgColor: "#3D3D3D",
          review:
            "This is our favorite expert for all mobile and web developing\
              areas. Darpan has a great team who get the job done above our\
              expectations. We would like to hire Canopas for more projects in\
              the future and highly recommend their kind and professional\
              services.",
        },
        {
          id: "client-5",
          image: ramasis,
          name: "Ramsis A.",
          fontColor: "#FFFFFF",
          bgColor: "#3D3D3D",
          review:
            "Canopas has been great to work with. From day 1, they made\
              sure they understood exactly what I wanted and advised and\
              guided me during the process (with patience). This project\
              has ended, but I am sure to work with them again on other\
              projects. I would recommend their services. ",
        },
      ],
    };
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  position: relative;
}

.overlay {
  padding: 7% 0;
}

.render-background {
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  position: absolute;
  object-fit: cover;
  z-index: -5;
}

.border-text {
  -webkit-text-stroke: 1px #3d3d3d;
  color: #fff !important;
}

.flex-div {
  display: flex;
  flex-direction: column;
}

.title {
  padding: 10%;
  text-align: center;
}

.client-details {
  border: transparent;
  padding: 30px;
  margin: 0 1rem 3rem;
  border-radius: 3rem;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.5);
}

.client-image-div {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  flex-direction: column;
}

@include media-breakpoint-up(sm) {
  .client-image-div {
    flex-direction: row;
  }
}

@include media-breakpoint-up(md) {
  .flex-div {
    flex-direction: row wrap;
  }

  .client-details {
    padding: 35px;
    margin: 0 1rem 3rem;
  }
}

@include media-breakpoint-up(lg) {
  .flex-div {
    flex-flow: wrap;

    div {
      flex: 35%;
    }
  }

  .title {
    padding: 35% 20% 25% 10%;
    text-align: start;
  }
}

@include media-breakpoint-up(xl) {
  .title {
    padding: 20% 25% 20% 8%;
  }

  .client-details {
    margin: 0 2rem 3rem;
  }
}
</style>
