<template>
  <div>
    <br>
    <v-container
        fluid
        class="pa-0"
    >
      <v-row>
        <v-col cols="12">
          <div v-if="loading" class="text-center">
            <v-progress-circular
                indeterminate
                color="primary"
            ></v-progress-circular>
          </div>
          <div v-else class="text-center">
            <h2 class="text-center">
              Deposit amount for {{ toName }}: {{ paymentAmount }}
            </h2>
            <br>
            <p v-if="isBalanceErr" style="color: red">
              Balance is not enough
            </p>
            <v-btn class="text-center" @click="back">
              back
            </v-btn>
            <v-btn v-if="!isApprove && !isBalanceErr" color="primary" class="text-center" @click="reserve"
                   :loading="reserveLoading"
                   :disabled="reserveLoading"
            >
              Pay token for the deposit
            </v-btn>
            <v-btn v-else-if="!isBalanceErr" color="teal lighten-4" class="text-center" @click="approve"
                   :loading="approveLoading"
                   :disabled="approveLoading"
            >
              approve token for the deposit
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
    <v-dialog
        v-model="completeDialog"
        persistent
        light
        max-width="500px"
    >
      <v-card>
        <v-card-title class="text-center">
          <v-spacer></v-spacer>
        </v-card-title>
        <v-card-text class="text-center">
          <h2 class="text-center" style="display: block">Reservation made!</h2>
          <v-btn color="primary" class="text-center" style="margin-top: 30px"
                 @click="goReserve"
          >
            reservation list
          </v-btn>
        </v-card-text>
      </v-card>
    </v-dialog>
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
    approveLoading: false,
    reserveLoading: false,
    paymentAmount: 0,
    isApprove: false,
    allowance: 0,
    completeDialog: false,
    paymentId: null,
    isBalanceErr: false,
  }),
  computed: {
    unixtime() {
      const datetime = new Date(this.reservation.date + ' ' + this.reservation.time)
      return datetime.getTime() / 1000
    },
  },
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
      if (this.toBalance < this.paymentAmount * (10 ** this.toDecimal)) {
        this.isBalanceErr = true
        this.loading = false
        return
      }
      console.log("this.allowance")
      console.log(this.allowance)
      if (this.paymentAmount * (10 ** this.toDecimal) > this.allowance) {
        this.isApprove = true
      }
      this.loading = false
    },
    async approve() {
      try {
        this.approveLoading = true
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance, this.toName)
        const contractAddress = trustReserveContractAddress.address

        const approve = await contract.methods.approve(contractAddress, this.paymentAmount * (10 ** this.toDecimal))
            .send({from: address})
        console.log(approve)
        this.isApprove = false
        this.approveLoading = true
      } catch (error) {
        console.log(error)
        this.approveLoading = false
      }
    },
    async reserve() {
      try {
        this.reserveLoading = true
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance)

        this.paymentId = Math.random().toString(32).substring(2)
        console.log(this.paymentId)
        const reserve = await contract.methods.reserve(
            this.paymentId,
            this.merchant.address,
            this.toAddress,
            this.paymentAmount * (10 ** this.toDecimal),
            this.unixtime - this.merchant.cancelable_time,
            this.unixtime + 100000
        )
            .send({from: address})
        console.log(this.unixtime)
        console.log(this.unixtime - this.merchant.cancelable_time)
        console.log(this.unixtime + 100000)
        console.log(reserve)
        await this.saveReserve()
        this.reserveLoading = false
        this.completeDialog = true

      } catch (error) {
        console.log(error)
        this.reserveLoading = false
      }
    },
    async saveReserve() {
      try {
        const datetime = new Date(this.reservation.date + ' ' + this.reservation.time)
        const {data} = await this.$axios.post('/api/reserve', {
          payment_id: this.paymentId,
          merchant_address: this.merchant.address,
          surname: this.reservation.surname,
          firstname: this.reservation.firstname,
          phonenumber: this.reservation.phonenumber,
          number: this.reservation.number,
          time: this.unixtime,
        })
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
    goReserve() {
      this.completeDialog = false
      this.$router.push('/reserve')
    },
    back() {
      this.$emit("close")
    }
  },
  mounted() {
    this.setAllowance()
    this.calcRequiredAmount()
  }
}
</script>

<style scoped>

</style>
