<template>
  <div id="pageReserve">
    <v-container grid-list-xl fluid>
      <div v-for="(reserve, i) in reserves">
        {{ reserve.surname }}
        {{ reserve.firstname }}
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "ReserveIndex",
  async asyncData() {
    let reserves = []
    const {data} = await this.$axios.get('/api/reserve_list')
    reserves = data.reserves ? data.reserves : []
    return {
      reserves
    }
  }
}
</script>

<style scoped>

</style>
