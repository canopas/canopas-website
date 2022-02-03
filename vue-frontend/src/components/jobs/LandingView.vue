<template>
  <div class="text-center landing-container">
    <img :src="backgroundImage" class="render-background" />
    <div class="overlay flex-container" style="margin-bottom: 0px">
      <div class="description">
        <div class="para1-text">
          Are you looking for a job where you can leverage your knowledge and
          experience to build digital products that will be used by millions of
          customers.
        </div>
        <div class="para2-text">
          You've come to the right place. We're looking for an AVENGER who can
          join our existing team of AVENGERS.
        </div>
      <div style="padding: 2%;">
         <button class="button" style="vertical-align: middle">
          <span>SEE ALL POSITIONS</span>
          </button>
      </div>
      </div>
    </div>
   
  </div>
</template>


<style lang="scss" scoped>
.landing-container {
  position: relative;
}

.para1-text {
  position: relative;
  font-weight: bold;
  line-height: 1.5em;
  text-align: center;
  color: #3d3d3d;
  font-family: Inter, sans-serif;
}
.para2-text {
  position: relative;
  padding-bottom: 10%;
  margin-top:6%;
  font-family: Inter, sans-serif;
  font-size: 75%;
  line-height: 1.5em;
  text-align: center;
  color: rgba(61, 61, 61, 0.8);
}
.flex-container {
  display: flex;
  flex-direction: column;
}

.button {
  display: inline-block;
  border-radius: 7px;
  background: linear-gradient(90deg, #f2709c 0%, #ff9472 100%);
  font-family: Product Sans;
  border: none;
  margin-top: -8%;
  text-align: center;
  font-weight: bold;
  font-size: 75%;
  padding:1%;
  width: 32%;
  color: white;
}
.overlay {
  background: rgba(255, 255, 255, 0.95);
  padding: 7% 0;
}

.description {
  margin: auto;
  width: 72%;
}

@include media-breakpoint-up(sm) {
  .description {
    width: 49%;
  }
}

@include media-breakpoint-up(md) {
  .description{
    width: 25%;

  .para2-text {
   padding-bottom: 8%;

    }
  }
}



</style>