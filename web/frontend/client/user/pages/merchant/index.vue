<template>
  <v-container fluid>
    <v-layout>
      <v-flex text-xs-center>
        <v-row dense>
          <v-col
            v-for="(merchant,i) in merchants"
            :key="`merchant-card-${i}`"
            :cols="4"
          >
            <v-card
              max-width="375"
              class="mx-auto"
              @click="goToDetail(merchant)">
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
            </v-card>
          </v-col>
        </v-row>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: "MerchantIndex",
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
  },
  methods: {
    goToDetail(merchant) {
      console.log('goToDetail')
      console.log(merchant.address)
      this.$router.push(`/merchant/${merchant.address}/show`)
    }
  }
}
</script>

<style scoped>

</style>
