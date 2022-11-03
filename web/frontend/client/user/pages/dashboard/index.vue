<template>
  <div id="pageIndex">
    <page-header></page-header>
    <v-main>
      <v-container fluid>
        <v-card class="mx-auto" max-width="400">
          <v-list-item two-line>
            <v-list-item-content>
              <v-list-item-title class="text-h5">
                Your Account
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-card-text>
          <div class="text-center">
            <client-only>
              <vue-qrcode :value="qrCodeUrl"></vue-qrcode>
            </client-only>
          </div>
          </v-card-text>
          <v-list class="transparent">
            <v-list-item>
              <v-list-item-title>Network</v-list-item-title>
              <v-list-item-subtitle class="text-right">
                {{ web3.networkId }}
              </v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>Address</v-list-item-title>
              <v-list-item-subtitle class="text-right">
                {{ web3.coinbase }}
              </v-list-item-subtitle>
            </v-list-item>
          </v-list>

        </v-card>
        <v-layout wrap class="mt-5">
          <v-flex md4>
            <div class="text-center">
              <NuxtLink class="d-block" to="/merchant">
                <v-icon>store</v-icon>
                <div class="mt-2">
                  <span class="d-inline-block">merchants</span>
                </div>
              </NuxtLink>
            </div>
          </v-flex>
          <v-flex md4>
            <div class="text-center">
              <NuxtLink class="d-block" to="/reserve">
                <v-icon>mdi-calendar-check-outline</v-icon>
                <div class="mt-2">
                  <span class="d-inline-block">reservations</span>
                </div>
              </NuxtLink>
            </div>
          </v-flex>
          <v-flex md4>
            <div class="text-center">
              <NuxtLink class="d-block" to="/setting">
                <v-icon>mdi-cog-outline</v-icon>
                <div class="mt-2">
                  <span class="d-inline-block">settings</span>
                </div>
              </NuxtLink>
            </div>
          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </div>
</template>

<script>
import PageHeader from '@/components/PageHeader'
import VueQrcode from '@chenfengyuan/vue-qrcode';
import {mapGetters} from "vuex";

export default {
  // layouts: 'dashboard',
  components: {
    PageHeader,
    VueQrcode,
  },
  data: () => ({}),
  computed: {
    ...mapGetters('web3', ['getInstance']),
    web3() {
      return this.getInstance
    },
    qrCodeUrl() {
      return 'https://google.com'
    }
  },
};
</script>
