<template>
  <div>
    <UnSubscribe
      :companyTitle="config.COMPANY_TITLE"
      :apiUrl="config.STRAPI_URL"
    />
  </div>
</template>
<script setup>
import config from "~/config";
</script>
