<template>
  <div id="pageIndex">
    <page-header></page-header>
    <v-main>
      <v-container fluid>
        <v-card class="mx-auto" max-width="400">
          <v-list-item two-line>
            <v-list-item-content>
              <v-list-item-title class="text-h5">
                Your Score
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-card-text>
          </v-card-text>
          <v-list class="transparent">
            <v-list-item>
              <v-list-item-title>Your Score</v-list-item-title>
              <v-list-item-subtitle class="text-right" v-if="!scoreLoading">
                {{ score }}
              </v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>Your Token Amount</v-list-item-title>
              <v-list-item-subtitle class="text-right">
                {{ amount }}
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
import web3Mixin from "@/mixins/web3Mixin";
import Web3 from "web3";

export default {
  // layouts: 'dashboard',
  components: {
    PageHeader,
    VueQrcode,
  },
  mixins: [web3Mixin],
  data: () => ({
    scoreLoading: true,
    execPoint: 0,
    failurePoint: 0,
    amount: 0,
  }),
  computed: {
    ...mapGetters('web3', ['getInstance']),
    web3() {
      return this.getInstance
    },
    score() {
      /**
       * n = 0かつ m > 100：Exceptional
       n = 0かつ m > 50：Excellent
       n = 0かつ m > 20：Great
       n = 0かつ m >= 0：Good
       n > 1かつ m - n >= 0： Sometimes Bothering
       n > 1かつ m - n < 0： Bothering
       */
      if (this.failurePoint === 0) {
        if (this.execPoint > 100) {
          return "Exceptional"
        }
        if (this.execPoint > 50) {
          return "Excellent"
        }
        if (this.execPoint > 20) {
          return "Great"
        }
        return "Good"
      } else {
        if (this.execPoint - this.failurePoint >= 0) {
          return "Sometimes Bothering"
        }
        return "Bothering"
      }
    }
  },
  methods: {
    async setScore() {
      const instance = this.createWeb3Instance(Web3.givenProvider)
      const address = await instance.eth.getCoinbase();
      console.log("address")
      console.log(address)
      const contract = await this.getContract(instance)
      this.executionPoint = await contract.methods.executionPoint(address).call()
      this.failurePoint = await contract.methods.failurePoint(address).call()
      console.log(this.executionPoint)
      console.log(this.failurePoint)
      this.scoreLoading = false
    },
    async setAmount () {
      const instance = this.createWeb3Instance(Web3.givenProvider)
      const address = await instance.eth.getCoinbase();
      console.log("address")
      console.log(address)
      const contract = await this.getUtilityTokenContract(instance)
      const balance = await contract.methods.balanceOf(address).call()
      const decimals = await contract.methods.decimals().call()
      this.amount = balance / (10 ** decimals)
      console.log(this.amount)
    }
  },
  mounted() {
    this.setScore()
    this.setAmount()
  }
};
</script>
