<template>
  <v-breadcrumbs :items="items"></v-breadcrumbs>
</template>

<script>
import breadcrumbs from '~/assets/js/breadcrumbs'

export default {
  name: "Breadcrumbs",
  computed: {
    items () {
      return breadcrumbs.get(this.$route.name)
    }
  }
}
</script>

<style scoped>

</style>
