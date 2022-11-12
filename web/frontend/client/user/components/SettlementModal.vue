<template>
  <v-dialog v-model="active"
            persistent
            class="text-center"
            max-width="600"
  >
    <v-card class="px-4">
      <template v-if="step === selectTokenStep">
        <v-card-title class="mb-5 text-h4 justify-center">
          Select Token
        </v-card-title>
        <v-container
          fluid
          class="pa-0"
        >
          <v-row v-if="!loading" align="center">
            <v-col
              v-for="token in erc20"
              :key="token.address"
              cols="12"
              sm="12"
            >
              <v-btn block @click="calcTokenAmount(token.name)">
                {{ token.name }}　 {{ balance(token.name) }}
              </v-btn>
            </v-col>
          </v-row>
          <div v-else class="text-center my-4">
            <v-progress-circular
              indeterminate
              color="primary"
            ></v-progress-circular>
          </div>
        </v-container>
      </template>
      <template v-if="step === inputSendAmountStep">
        <v-card-title class="mb-5 text-h4 justify-center">
          Send To {{ toName }}
        </v-card-title>
        <v-card-text class=text-center>
          <div
            class="text-h4 inline-block mb-3"
          >Pay Amount</div>
          <div
            class=" text-h4 inline-block"
          >
            {{ input.current }}
          </div>
          <p>※input your total amount</p>
          <p>Deposit Amount: {{ depositAmount }}</p>
        </v-card-text>
        <template v-if="!loading">
          <v-card-actions class="justify-center my-5">
            <v-btn @click="enter" color="primary">enter</v-btn>
          </v-card-actions>
          <v-card-text class="pa-0">
          <v-container grid-list-xs pa-1>
            <v-layout row wrap pa0>
              <v-flex v-for="button in buttons" :key="button.key" xs4 class="text-center mb-2">
                <v-btn v-if="button.key"
                       :class="['ma-0', button.class]"
                       :color="button.color"
                       @click="inputKey(button.key)"
                       fab
                       large
                >
                  <v-icon v-if="button.icon" dark>{{ button.icon }}</v-icon>
                  <template v-else>{{ button.label }}</template>
                </v-btn>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card-text>
        </template>
        <div v-else class="text-center my-4">
          <v-card-text>Calculating...</v-card-text>
          <v-progress-circular
            indeterminate
            color="primary"
          ></v-progress-circular>
        </div>
      </template>
      <template v-if="step === commitStep">
        <v-card-title class="mb-5 text-h4">
          Confirmation
        </v-card-title>
        <v-card-text class="text-center">
          <div
            class="text-h4 inline-block mb-3"
          >You send {{paymentTokenName}}</div>
          <div
            class="inline-block"
          ><p>{{paymentAmount}}</p></div>
        </v-card-text>
        <template v-if="!loading">
          <v-card-actions class="justify-center my-5">
            <v-btn @click="commit" color="primary">OK</v-btn>
          </v-card-actions>
        </template>
        <div v-else class="text-center my-4">
          <v-card-text>Calculating...</v-card-text>
          <v-progress-circular
            indeterminate
            color="primary"
          ></v-progress-circular>
        </div>
      </template>
      <template v-if="step === completeStep">
        <v-card-title class="mb-5 text-h4 justify-center">
          Result
        </v-card-title>
        <v-card-text>Payment completed!<br/>We look forward to seeing you again!</v-card-text>
      </template>
      <v-card-actions>
        <v-flex class="d-flex justify-space-between">
          <v-btn v-if="step === 2"
                 color="primary"
                 text
                 class=""
                 :disabled="loading"
                 @click="back"
          >
            back
          </v-btn>
          <v-btn
            color="primary"
            text
            class="ml-auto"
            :disabled="loading"
            @click="close"
          >
            close
          </v-btn>
        </v-flex>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import {AlphaRouter} from "@uniswap/smart-order-router";
import {CurrencyAmount, Percent, Token, TradeType, Pair} from "@uniswap/sdk-core";
import {ethers} from "ethers";
import {get} from 'lodash'
import tokenContactAddresses from "~/constants/tokenContactAddresses";
import Web3 from "web3";
import web3Mixin from "~/mixins/web3Mixin";
import trustReserveContractAddress from "~/constants/trustReserveContractAddress.json";

export default {
  name: "SettlementModal",
  mixins: [web3Mixin],
  data() {
    return {
      input: {
        current: "0",
        operator: "",
        prev: "",
      },
      paymentAmount: 0,
      receiveTokenDecimals: 18,
      paymentTokenName: '',
      paymentTokenDecimal: 18,
      paymentAddress: '',
      paymentBalance: 0,
      balances: [],

      allowance: 0,
      loading: false,
      pathAddresses: [],
    }
  },
  props: {
    step: {
      type: Number,
      default: 1,
    },
    active: {
      type: Boolean,
      default: false,
    },
    reservation: {
      type: Object,
      default: () => ({})
    }
  },
  computed: {
    receiveAddress() {
      return get(this.reservation, 'merchant.received_address', '')
    },
    depositAmount() {
      return get(this.reservation, 'depositAmount', '0')
    },
    erc20Address() {
      let erc20 = Object.keys(tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES).filter((key) => {
        return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES[key].address === this.receiveAddress
      });

      if (erc20.length > 0) {
        return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES[erc20].name
      }

      return this.receiveAddress
    },
    erc20() {
      return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES
    },
    balance() {
      return (ercName) => {
        const item = this.balances.find(item => {
          return item.name === ercName
        })
        if (!item) {
          return "--"
        }
        return Math.round(item.balance / (10 ** item.decimals) * 10000) / 10000
      }
    },
    selectTokenStep() {
      return 1
    },
    inputSendAmountStep() {
      return 2
    },
    commitStep() {
      return 3
    },
    completeStep() {
      return 4
    },
    receiveTokenAmount() {
      const amount = this.input.current
      return amount > 0 ? amount : 0
    },
    toName() {
      return get(this.reservation, 'merchant.name', '')
    },
    buttons() {
      return [
        {key: "1", label: "1", color: "", class: "headline"},
        {key: "2", label: "2", color: "", class: "headline"},
        {key: "3", label: "3", color: "", class: "headline"},
        {key: "4", label: "4", color: "", class: "headline"},
        {key: "5", label: "5", color: "", class: "headline"},
        {key: "6", label: "6", color: "", class: "headline"},
        {key: "7", label: "7", color: "", class: "headline"},
        {key: "8", label: "8", color: "", class: "headline"},
        {key: "9", label: "9", color: "", class: "headline"},
        {key: "0", label: "0", color: "", class: "headline"},
        {key: "ac", label: "C", color: "grey lighten-2", class: ""},
        {
          key: "back",
          label: "BACK",
          icon: "backspace",
          color: "grey lighten-2",
          class: "",
        },
      ];
    }
  },
  methods: {
    close() {
      this.$emit('close')
    },
    inputKey(key) {
      const isNumber = /^\d$/.test(key);
      if (isNumber) {
        if (this.input.current.length >= 12) {
          return;
        }
        if (this.input.current !== "0") {
          this.input.current += key;
        } else {
          this.input.current = key;
        }
        return
      }
      switch (key) {
        case "back":
          this.input.current =
            this.input.current.substring(0, this.input.current.length - 1) || "0";
          return;
        case "c":
          this.input.current = "0";
          this.input.operator = "";
          this.input.prev = "";
          return;
        case "ce":
          this.input.current = "0";
          return;
      }
    },
    next(nextStep) {
      this.$emit('next', nextStep)
    },
    back() {
      this.$emit('back')
    },
    async calcTokenAmount(tokenName) {
      const item = this.balances.find(item => {
        return item.name === tokenName
      })
      this.paymentTokenDecimal = Number(item.decimals)
      this.paymentTokenName = tokenName
      this.paymentAddress = item.address
      this.next('inputSendAmountStep')
    },
    async setBalances() {
      this.erc20.forEach(token => {
        this.setBalanceOf(token)
      })
    },
    async setBalanceOf(token) {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance, token.name)
        const balance = await contract.methods.balanceOf(address).call()
        const decimals = await contract.methods.decimals().call()
        console.log(balance)
        console.log(decimals)
        this.balances.push({name: token.name, balance: balance, decimals: decimals, address: contract._address})
      } catch (error) {
        console.log(error)
      }
    },
    async setReceiveTokenDecimals() {
      const instance = this.createWeb3Instance(Web3.givenProvider)
      const contract = await this.getContract(instance, this.receiveTokenName)
      this.receiveTokenDecimals = Number(await contract.methods.decimals().call())
    },
    async setAllowance() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance, this.paymentTokenName)
        const contractAddress = trustReserveContractAddress.address

        this.allowance = await contract.methods.allowance(address, contractAddress)
          .call()
        console.log(this.allowance)
      } catch (error) {
        console.log(error)
      }
    },
    async calcRequiredAmount() {
      try {
        await window.ethereum.request({method: 'eth_requestAccounts'})
        const instance = this.createWeb3Instance(window.ethereum)
        const networkId = await instance.eth.net.getId();
        const from = new Token(networkId, this.receiveAddress, this.receiveTokenDecimals, this.receiveTokenName, this.receiveTokenName)
        const to = new Token(networkId, this.paymentAddress, this.paymentTokenDecimal, this.paymentTokenName, this.paymentTokenName)

        const requireAmount = CurrencyAmount.fromRawAmount(from, this.receiveTokenAmount)
        const web3Provider = new ethers.providers.JsonRpcProvider(process.env.infuraUrl)
        const router = new AlphaRouter({chainId: networkId, provider: web3Provider})
        const route = await router.route(
          requireAmount,
          to,
          TradeType.EXACT_OUTPUT,
          {
            recipient: this.receiveAddress, // FIXME
            slippageTolerance: new Percent(25, 100),
            deadline: Math.floor(Date.now() / 1000 + 1800)
          }
        )
        this.paymentAmount = route.quote.toFixed(this.paymentTokenDecimal)
        if (this.toBalance < this.paymentAmount * (10 ** this.toDecimal)) {
          console.log('balance err')
        }
        // swapするpath(settleReservationの_path引数)
        console.log('route.trade.routes[0].path')
        console.log(route.trade.routes[0].path)
        this.pathAddresses = route.trade.routes[0].path.map(v => {
          return v.address
        })
        if (this.paymentAmount * (10 ** this.paymentTokenDecimal) > this.allowance) {
          this.isApprove = true
        }
      }catch(e){
        console.log(e)
      }
    },
    async enter() {
      this.loading = true
      try {
        await this.setAllowance()
        await this.calcRequiredAmount()

        const instance = this.createWeb3Instance(Web3.givenProvider)
        const address = await instance.eth.getCoinbase();
        const contract = await this.getContract(instance, this.paymentTokenName)
        const contractAddress = trustReserveContractAddress.address

        const approve = await contract.methods.approve(contractAddress, this.paymentAmount * (10 ** this.paymentTokenDecimal))
          .send({from: address})
        console.log(approve)
      } catch (error) {
        console.log(error)
      }
      this.loading = false
      this.next('commit')
    },
    async commit () {
      this.loading = true
      try {
        await this.settlement()
      }catch(e){
        console.log(e)
      }
      this.loading = false
      this.next('complete')
    },
    async settlement() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const accounts = await instance.eth.getAccounts()
        const account = accounts[0]
        const contract = await this.getContract(instance)
        const datetime = new Date()
        const amountIn = this.paymentAmount * (10 ** this.paymentTokenDecimal)
        const res = await contract.methods.settleReservation(
          amountIn,
          this.input.current,
          datetime.getTime(),
          this.reservation.payment_id,
          this.pathAddresses,
        ).send({from: account})
        console.log(res)
        this.reservation = res
      } catch (error) {
        console.log(error)
      }
    },
  },
  mounted() {
    this.setBalances()
    this.setReceiveTokenDecimals()
  }
}
</script>

<style scoped>

</style>
