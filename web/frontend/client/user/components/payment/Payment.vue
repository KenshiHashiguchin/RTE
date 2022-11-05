<template>
  <div>
    <template v-if="!displayTokenAmount">
      <b>
        <slot>Pay</slot>
      </b>
      <br>
      <br>
      <v-container
          fluid
          class="pa-0"
      >
        <v-row align="center">
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
      </v-container>
    </template>
    <template v-else>
      <CalcTokenAmount
          :from-name="receiveTokenName"
          :from-decimal="receiveTokenDecimals"
          :from-amount="receiveTokenAmount"
          :from-address="receiveAddress"
          :to-name="paymentTokenName"
          :to-decimal="paymentTokenDecimal"
          :to-address="paymentAddress"
          :to-balance="balanceOf(paymentTokenName)"
          :reservation="reservation"
          :merchant="merchant"
      ></CalcTokenAmount>
    </template>
  </div>
</template>

<script>
import tokenContactAddresses from '@/constants/tokenContactAddresses';
import Web3 from "web3";
import web3Mixin from "@/mixins/web3Mixin";
import CalcTokenAmount from "~~/client/user/components/payment/CalcTokenAmount";

export default {
  name: "Payment",
  components: {CalcTokenAmount},
  mixins: [web3Mixin],
  props: {
    receiveTokenName: {
      type: String,
      required: true,
    },
    receiveTokenAmount: {
      type: Number,
      required: true,
    },
    receiveAddress: {
      type: String,
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
    balances: [],
    payerAddress: null,
    receiveTokenDecimals: 18,
    displayTokenAmount: false,

    // paymentToken
    paymentTokenName: '',
    paymentTokenDecimal: 18,
    paymentAddress: '',
    paymentBalance: 0,
  }),
  computed: {
    erc20Address() {
      let erc20 = Object.keys(tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES).filter((key) => {
        return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES[key].address === this.merchant.received_address
      });

      if (erc20.length > 0) {
        return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES[erc20].name
      }

      return this.merchant.received_address
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
    balanceOf() {
      return (ercName) => {
        const item = this.balances.find(item => {
          return item.name === ercName
        })
        if (!item) {
          return 0
        }
        return Number(item.balance)
      }
    }
  },
  methods: {
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
    calcTokenAmount(tokenName) {
      const item = this.balances.find(item => {
        return item.name === tokenName
      })
      this.paymentTokenDecimal = Number(item.decimals)
      this.paymentTokenName = tokenName
      console.log("calcTokenAmount")
      console.log(item.address)
      this.paymentAddress = item.address
      this.displayTokenAmount = true
    }
  },
  mounted() {
    this.setBalances()
    this.setReceiveTokenDecimals()
  }
}
</script>

<style scoped>

</style>
