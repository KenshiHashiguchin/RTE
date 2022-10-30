<template>
  <div>
    <div>
      <v-text-field
        label="SurName"
        hide-details="auto"
        v-model="surname"
      ></v-text-field>
      <v-text-field
        label="FirstName"
        hide-details="auto"
        v-model="firstname"
      ></v-text-field>
      <v-text-field
        label="PhoneNumber"
        hide-details="auto"
        v-model="phonenumber"
      ></v-text-field>
      <v-text-field
        label="Number"
        hide-details="auto"
        v-model="number"
      ></v-text-field>
      <v-text-field
        label="Time"
        hide-details="auto"
        v-model="time"
      ></v-text-field>
      <v-btn @click="reserve">reserve</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: "ReserveCreate",
  async asyncData({app, route}) {
    let merchant = null
    console.log(route.params)
    try {
      const {data} = await app.$axios.get(`/api/merchant/${route.params.address}`)
      merchant = data.merchant
    } catch (e) {
      console.log(e)
    }
    return {
      merchant
    }
  },
  data() {
    return {
      surname: '',
      firstname: '',
      phonenumber: '',
      number: '',
      time: '',
    }
  },
  methods: {
    async reserve() {
      try {
        const {data} = await this.$axios.post('/api/reserve', {
          merchant_address: this.merchant.address,
          surname: 'hogehoge',
          firstname: 'fugafuga',
          phonenumber: '08064263884',
          number: 2,
          time: '2022-12-01 18:00:00',
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
