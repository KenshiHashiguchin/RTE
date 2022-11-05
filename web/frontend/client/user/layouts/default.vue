<template>
  <div id="appRoot">
    <v-app>
      <nuxt />
    </v-app>
  </div>
</template>

<script>
import Web3 from "web3";
import {mapMutations} from "vuex";

export default {
  name: 'defaultLayout',
  methods: {
    ...mapMutations('web3', ['registerWeb3Instance']),
  },
  async mounted() {
    if (typeof window !== 'undefined' && typeof window.ethereum !== 'undefined') {
      const instance = new Web3(window.ethereum)
      instance.eth.net.isListening()
        .then(async (accounts) => {
          console.log(accounts)
          this.$store.commit('web3/updateIsListening', true)
          await window.ethereum.request({method: 'eth_requestAccounts'})
          const networkId = await instance.eth.net.getId();
          const coinbase = await instance.eth.getCoinbase();
          const balance = await instance.eth.getBalance(coinbase);
          this.registerWeb3Instance({
            networkId,
            coinbase,
            balance
          })
        })
        .catch(e => {
          console.log('Wow. Something went wrong: ' + e)
          this.$store.commit('web3/updateIsListening', false)
        })
    }
    console.log('isLoggedIn end')
  },
}
</script>
