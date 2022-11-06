<template>
  <div>
    <br>
    <v-container
        fluid
        class="pa-0"
    >
      <v-row>
        <v-col cols="12">
          <h2 class="text-center">
            {{ toName }}
          </h2>
        </v-col>
        <v-col cols="12">
          <div v-if="loading" class="text-center">
            <v-progress-circular
                indeterminate
                color="primary"
            ></v-progress-circular>
          </div>
          <div v-else class="text-center">
            <h3 class="text-center">
              {{ paymentAmount }}
            </h3>
            <br>
            <v-btn v-if="!isApprove" color="primary" class="text-center" @click="reserve">
              Pay token for the deposit
            </v-btn>
            <v-btn v-else color="teal lighten-4" class="text-center" @click="approve">
              approve token for the deposit
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import web3Mixin from "@/mixins/web3Mixin";
import {AlphaRouter} from "@uniswap/smart-order-router";
import {CurrencyAmount, Percent, Token, TradeType} from "@uniswap/sdk-core";
import {ethers} from "ethers";
import Web3 from "web3";
import trustReserveContractAddress from "~~/client/user/constants/trustReserveContractAddress.json";

export default {
  name: "CalcTokenAmount",
  mixins: [web3Mixin],
  props: {
    fromName: {
      type: String,
      required: true,
    },
    fromDecimal: {
      type: Number,
      required: true,
    },
    fromAmount: {
      type: Number,
      required: true,
    },
    fromAddress: {
      type: String,
      required: true,
    },
    toName: {
      type: String,
      required: true,
    },
    toDecimal: {
      type: Number,
      required: true,
    },
    toAddress: {
      type: String,
      required: true,
    },
    toBalance: {
      type: Number,
      required: true,
    },
    reservation: {
      type: Object,
      required: true,
    },
    merchant: {
      type: Object,
      required: true,
    }
  },
  data: () => ({
    loading: true,
    paymentAmount: 0,
    isApprove: false,
    allowance: 0,
  }),
  computed: {},
  methods: {
    async setAllowance() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance, this.toName)
        const contractAddress = trustReserveContractAddress.address
        this.allowance = await contract.methods.allowance(address, contractAddress).call()
        console.log(this.allowance)
      } catch (error) {
        console.log(error)
      }
    },
    async calcRequiredAmount() {
      await window.ethereum.request({method: 'eth_requestAccounts'})
      const instance = this.createWeb3Instance(window.ethereum)
      const networkId = await instance.eth.net.getId();
      const from = new Token(networkId, this.fromAddress, this.fromDecimal, this.fromName, this.fromName)
      const to = new Token(networkId, this.toAddress, this.toDecimal, this.toName, this.toName)

      const requireAmount = CurrencyAmount.fromRawAmount(from, this.fromAmount)
      const web3Provider = new ethers.providers.JsonRpcProvider(process.env.infuraUrl)
      const router = new AlphaRouter({chainId: networkId, provider: web3Provider})
      const route = await router.route(
          requireAmount,
          to,
          TradeType.EXACT_OUTPUT,
          {
            recipient: this.fromAddress, // FIXME
            slippageTolerance: new Percent(25, 100),
            deadline: Math.floor(Date.now() / 1000 + 1800)
          }
      )
      this.paymentAmount = route.quote.toFixed(this.toDecimal)
      // swapするpath(settleReservationの_path引数)
      console.log(route.trade.routes[0].path)

      console.log("this.allowance")
      console.log(this.allowance)
      if (this.paymentAmount * (10 ** this.toDecimal) > this.allowance) {
        this.isApprove = true
      }
      this.loading = false
    },
    async approve() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance, this.toName)
        const contractAddress = trustReserveContractAddress.address

        const approve = await contract.methods.approve(contractAddress, this.paymentAmount * (10 ** this.toDecimal))
            .send({from: address})
        console.log(approve)

      } catch (error) {
        console.log(error)
      }
    },
    async reserve() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance)

        const paymentId = Math.random().toString(32).substring(2)
        console.log('---paymentId----')
        console.log(paymentId)

        // calc timestamp
        const datetime = new Date(this.reservation.date +' '+this.reservation.time)
        console.log('---datetime.getTime()---')
        console.log(datetime.getTime())

        const reserve = await contract.methods.reserve(
            paymentId,
            this.merchant.address,
            this.toAddress,
            this.paymentAmount * (10 ** this.toDecimal),
            datetime.getTime(),
            // datetime.getTime() - this.merchant.cancelable_time,
            datetime.getTime() + 100000
        )
            .send({from: address})
        console.log('---reserve----')
        console.log(reserve)

      } catch (error) {
        console.log(error)
      }
    },
  },
  mounted() {
    this.setAllowance()
    this.calcRequiredAmount()
  }
}
</script>

<style scoped>

</style>
