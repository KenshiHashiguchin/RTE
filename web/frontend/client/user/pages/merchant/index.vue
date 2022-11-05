<template>
  <div id="pageMerchantIndex">
    <page-header></page-header>
    <v-main>
      <v-container fluid>
        <breadcrumbs/>
        <h1>Merchant</h1>
        <v-layout>
          <v-flex text-xs-center>
            <MerchantTable
              :merchants="merchants"
              @reserve="reserve"
            >
            </MerchantTable>
          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </div>
</template>

<script>
import Breadcrumbs from "~/components/common/Breadcrumbs";
export default {
  name: "MerchantIndex",
  components: {Breadcrumbs},
  async asyncData({app}) {
    let merchants = []
    try {
      const {data} = await app.$axios.get('/api/merchants')
      merchants = data.merchants ? data.merchants.map(v => {
        v.cancelable_days = `before ${Math.floor(v.cancelable_time / 60 / 60)} days`
        return v
      }) : []
    } catch (e) {
      console.log(e)
    }
    return {
      merchants
    }
  },
  methods: {
    reserve(merchant) {
      console.log(merchant)
    },
  }
}
</script>

<style scoped>

</style>
