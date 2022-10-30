<template>
  <div>
    <div v-for="(merchant,i) in merchants" class="">
      {{ merchant.name }}
      {{ merchant.phonenumber }}
      {{ merchant.phonenumber }}
      <NuxtLink :to=`/merchant/${merchant.id}`>go to detail</NuxtLink>
    </div>
  </div>
</template>

<script>
export default {
  name: "ReserveCreate",
  async asyncData(app) {
    let merchants = []
    try {
      const {data} = await app.$axios.get('/api/merchants')
      merchants = data.merchants ? data.merchants : []
    } catch (e) {
      console.log(e)
    }
    return {
      merchants
    }
  }
}
</script>

<style scoped>

</style>
