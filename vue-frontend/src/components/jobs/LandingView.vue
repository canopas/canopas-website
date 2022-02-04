<template>
  <div class="text-center landing-container">
      <div class="overlay flex-container">
      <div class="description">
        <div class="para1-text normal-text ">
          Are you looking for a job where you can leverage your knowledge and
          experience to build digital products that will be used by millions of
          customers.
        </div>
        <div class="para2-text normal-text">
          You've come to the right place. We're looking for an AVENGER who can
          join our existing team of AVENGERS.
        </div>
      <div>
         <button class="button normal-text">
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
.normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }
 .para1-text {
     position: relative;
     font-weight: bold;
     text-align: center;
     color: #3d3d3d;
}
 .para2-text {
     position: relative;
     padding-bottom: 10%;
     margin-top:6%;
     font-size: 1rem;
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
     border: none;
     text-align: center;
     font-weight: bold;
     font-size: 85%;
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
         width: 36%;
    }
     .para1-text {
         padding: 0%;
         width: 93%;
         padding-left: 10%;
    }
     .para2-text {
         padding-bottom: 8%;
         
    }
     .button{
         font-size: 95%;
         padding: 1.5% 0%;
         width: 24%;
    }
}
 @include media-breakpoint-between(md, xl) {
     .description{
         width: 46%;
    }
     .para1-text {
         width: 100%;
    }
     .para2-text{
         width: 100%;
         margin-left: 6%;
    }
     .button{
         width: 42%;
    }
     @include media-breakpoint-between(xl, xxl) {
         .description{
             width: 60%;
             margin-left: 28%;
        }
         .para1-text {
             padding-left: 10%;
        }
         .para2-text{
             width: 75%;
        }
         .button{
             width: 22%;
             margin-left: -30%;
        }
    }
}


</style>