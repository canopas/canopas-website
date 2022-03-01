<template>
  <div class="container">
    <div class="contact-title">Hey smarty,</div>
    <div class="normal-text contact-description">
      <p>
        Thanks for reaching out to us. We’re happy to help you achieve your
        goals and dreams.
      </p>
      <p>
        As you’re landed here, I guess you’re looking for help with your idea or
        existing product and you’re wondering how our team can help you.
      </p>
      <p>
        We will guide you through the process so you don’t have to worry about
        it. Before that, you gotta fill up the form so we can know more about
        you before we approach you.
      </p>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
};
</script>

<style lang="scss" scoped>
.container {
  margin: 48px auto;
  padding: 0 5%;
}

.contact-title {
  margin-bottom: 32px;
  font-size: 2.25rem;
  font-weight: bold;
  letter-spacing: 1px;
  background: linear-gradient(89.88deg, #f2709c -1.25%, #ff9472 28.15%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.contact-description {
  line-height: 1.5rem;
}

.contact-description > p {
  margin-bottom: 24px;
}
</style>
