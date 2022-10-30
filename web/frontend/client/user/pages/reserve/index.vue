<template>
  <div id="pageReserve">
    <v-container grid-list-xl fluid>
      RESERVES
      <div v-for="(reservation, i) in reservations">
        {{ reservation.surname }}
        {{ reservation.firstname }}
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "ReserveIndex",
  async asyncData({ app }) {
    let reservations = []
    try {
      const {data} = await app.$axios.get('/api/reserve_list')
      reservations = data.reservations ? data.reservations : []
    }catch (e){
      console.log(e)
    }
    return {
      reservations
    }
  }
}
</script>

<style scoped>

</style>
