<template>
  <div id="pageReserve">
    <v-container grid-list-xl fluid>
      RESERVES
      <div v-for="(reservation, i) in reservations">
        {{ reservation.surname }}
        {{ reservation.first_name }}
        {{ reservation.phonenumber }}
        {{ reservation.number }}
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "ReserveIndex",
  async asyncData({ app, error }) {
    let reservations = []
    try {
      const {data} = await app.$axios.get('/api/reserve_list')
      reservations = data.reservations ? data.reservations : []
    }catch (e){
      error({
        statusCode: e.response.status,
        message: e.response.data.message,
      })
    }
    return {
      reservations
    }
  }
}
</script>
