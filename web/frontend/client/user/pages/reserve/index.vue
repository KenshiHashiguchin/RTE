<template>
  <div id="pageReserveIndex">
    <page-header></page-header>
    <v-main>
      <v-container fluid>
        <breadcrumbs/>
        <h1>Reservations</h1>
        <v-layout>
          <v-flex text-xs-center>
            <v-row dense>
              <v-col v-for="(reservation, i) in reservations"
                     :cols="12" :key="`reservation-list-item-${i}`">
                <v-list one-line>
                  <v-list-item>
                    <v-list-item-content>
                      <v-list-item-title>
                        <v-chip
                          class="ma-2"
                          label
                          color="cyan"
                          text-color="white"
                        >
                          Status
                        </v-chip>
                        {{ status }}
                      </v-list-item-title>
                      <v-divider></v-divider>
                      <!-- todo merchant name-->
                      <v-list-item-title>
                        <v-chip
                          class="ma-2"
                          label
                          text-color="white"
                        >
                          Merchant Id
                        </v-chip>
                        {{ reservation.merchant_id }}
                      </v-list-item-title>
                      <v-list-item-title>
                        <v-chip
                          class="ma-2"
                          label
                          text-color="white"
                        >
                          Number of Visitors
                        </v-chip>
                        {{ reservation.number }}
                      </v-list-item-title>
                      <!--{{ reservation.surname }}-->
                      <!--{{ reservation.first_name }}-->
                      <!--{{ reservation.phonenumber }}-->
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-col>
            </v-row>
          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </div>
</template>

<script>
import Breadcrumbs from "~/components/common/Breadcrumbs";

export default {
  name: "ReserveIndex",
  components: {Breadcrumbs},
  async asyncData({app, error}) {
    let reservations = []
    try {
      const {data} = await app.$axios.get('/api/reserve_list')
      reservations = data.reservations ? data.reservations : []
    } catch (e) {
      error({
        statusCode: e.response.status,
        message: e.response.data.message,
      })
    }
    return {
      reservations
    }
  },
  computed: {
    status() {
      const statuses = ['None',
      'Reserved',
      'Canceled',
      'Settled']
      return statuses[Math.floor(Math.random() * statuses.length)]
    }
  }
}
</script>
