<template>
  <div id="appRoot">
    <v-app>
      <nuxt />
    </v-app>
  </div>
</template>

<script>
import Web3 from "web3";

export default {
  name: 'defaultLayout',
  async mounted() {
    if (typeof window.ethereum !== 'undefined') {
      let currentAccount = null
      await window.ethereum.request({method: 'eth_accounts'})
      const instance = new Web3(window.ethereum)
      instance.eth.net.isListening()
        .then((accounts) => {
          // if(accounts[0] !== currentAccount) {
          //   currentAccount = accounts[0];
          console.log(accounts)
          this.$store.commit('web3/updateIsListening', true)
          // }
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
