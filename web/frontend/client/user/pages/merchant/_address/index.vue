<template>
  <v-container fluid>
    <v-layout>
      <v-flex text-xs-center>
        <v-row dense>
          <v-col :cols="6">
            <v-card
              max-width="375"
              class="mx-auto">
              <v-img
                src="https://cdn.vuetifyjs.com/images/lists/ali.png"
                height="150px"
                dark
              />
              <v-list>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.name : {{ merchant.name }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.address : {{ merchant.address }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.received_address : {{ merchant.received_address }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.merchant_address : {{ merchant.merchant_address }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.tel : {{ merchant.tel }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.deposit : {{ merchant.deposit }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>merchant.cancelable_time : {{ merchant.cancelable_time }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
              <v-btn @click="goToReserve">Reserve</v-btn>
              <v-btn @click="cancel">cancel</v-btn>
            </v-card>
          </v-col>
        </v-row>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "MerchantAddress",
  async asyncData(app, query) {
    let merchant = null
    console.log(query.address)
    try {
      const {data} = await this.$axios.get(`/api/merchant/${query.address}`)
      merchant = data.merchant
    } catch (e) {
      console.log(e)
    }
    return {
      merchant
    }
  },
  methods: {
    goToReserve(merchant) {
      this.$router.push(`/merchant/${merchant.address}/reserve`)
    },
    cancel() {
      // todo open modal
    }
  }
}
</script>

<style scoped>

</style>
