<template>
  <div>
    <div>
      <v-btn @clink="reserve">reserve</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: "ReserveCreate",
  methods: {
    async reserve() {
      try {
        const {data} = await this.$axios.post('/api/reserve', {
          params: {
            merchant_id: 1,
            surname: 'hogehoge',
            firstname: 'fugafuga',
            phonenumber: '08064263884',
            number: 2,
            time: '2022-12-01 18:00:00',
          }
        })
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
  }
}
</script>
