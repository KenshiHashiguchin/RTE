<template>
  <v-app id="login">
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4 lg4>
            <v-card class="elevation-1 pa-3">
              <v-card-text>
                <div class="layout column align-center">
                  <img
                    src="https://upload.wikimedia.org/wikipedia/commons/3/36/MetaMask_Fox.svg" alt="Metamask"
                    width="120" height="120">
                  <h1 class="flex my-4 primary--text">RTE Admin</h1>
                </div>
              </v-card-text>
              <v-card-actions>
                <metamask @connect="initWeb3"/>
                <p class="italic text-red-600">{{ errorMessage }}</p>
              </v-card-actions>
              <div>
                <p>//TODO move to another page</p>
                <p>Network: {{ web3.networkId }}</p>
                <p>Account: {{ web3.coinbase }}</p>
                <p>Balance: {{ web3.balance }}</p>
              </div>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Web3 from 'web3'
import { mapGetters, mapMutations } from 'vuex';
import Metamask from "~~/client/merchant/components/Metamask";

export default {
  components: { Metamask },
  layout: 'default',
  data: () => ({
    errorMessage: '',
  }),
  computed: {
    ...mapGetters('web3', ['getInstance']),
    web3() {
      return this.getInstance
    }
  },
  methods: {
    ...mapMutations('web3', ['registerWeb3Instance']),
    async initWeb3() {
      // Check for web3 provider
      if (typeof window.ethereum !== 'undefined') {
        try {
          // Ask to connect
          await window.ethereum.send('eth_requestAccounts');
          const instance = new Web3(window.ethereum)
          // Get necessary info on your node
          const networkId = await instance.eth.net.getId();
          const coinbase = await instance.eth.getCoinbase();
          const balance = await instance.eth.getBalance(coinbase);
          // Save it to store
          this.registerWeb3Instance({
            networkId,
            coinbase,
            balance
          });
          this.errorMessage = '';
        } catch (error) {
          // User denied account access
          console.error('User denied web3 access', error);
          this.errorMessage = 'Please connect to your Ethereum address on Metamask before connecting to this website';
        }
      }
      // No web3 provider
      else {
        console.error('No web3 provider detected');
        this.errorMessage = "No web3 provider detected. Did you install the Metamask extension on this browser?";
      }
    }
  }
}
</script>
<style scoped lang="css">
#login {
  background-color: #00BCD4;
  height: 50%;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
  content: "";
  z-index: 0;
}
</style>
