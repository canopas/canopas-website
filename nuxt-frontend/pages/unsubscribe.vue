<template>
  <div>
    <UnSubscribe
      :company-title="config.COMPANY_TITLE"
      :api-url="config.STRAPI_URL"
    />
  </div>
</template>
<script setup>
import config from "~/config";
</script>
