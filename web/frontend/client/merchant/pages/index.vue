<template>
  <v-app id="login">
    <page-header></page-header>
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4 lg4>
            <v-card class="elevation-1 pa-3">
              <v-card-text>
                <div class="layout column align-center">
                  <img width="100%" src="/dark-logo.png" alt="trust reserve logo"/>
                </div>
              </v-card-text>
              <v-card-actions>
                <metamask @connect="login"/>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
        <Dialog :active="isErrorDialogActive"
                :message="errorMessage"
                @close="isErrorDialogActive = false"></Dialog>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Web3 from 'web3'
import { mapGetters, mapMutations } from 'vuex';
import Metamask from "~/components/Metamask";
import Dialog from "~/components/common/Dialog";

export default {
  components: {Dialog, Metamask},
  layout: 'default',
  data: () => ({
    errorMessage: '',
    isErrorDialogActive: false,
  }),
  computed: {
    ...mapGetters('web3', ['getInstance']),
    web3() {
      return this.getInstance
    }
  },
  methods: {
    ...mapMutations('web3', ['registerWeb3Instance']),
    async login() {
      try {
        await this.connect()
        const token = await this.getToken()
        await this.signature(token)
        await this.$router.push('/dashboard')
      } catch (e) {
        console.log(e)
        this.isErrorDialogActive = true
        this.errorMessage = e.message
      }
    },
    async connect() {
      // Check for web3 provider
      if (typeof window.ethereum !== 'undefined') {
        try {
          // Ask to connect
          await window.ethereum.request({method: 'eth_requestAccounts'})
          const instance = new Web3(window.ethereum)
          const networkId = await instance.eth.net.getId();
          const coinbase = await instance.eth.getCoinbase();
          const balance = await instance.eth.getBalance(coinbase);
          this.registerWeb3Instance({
            networkId,
            coinbase,
            balance
          });

          this.errorMessage = '';
        } catch (error) {
          throw new Error('Please connect to your Ethereum address on Metamask before connecting to this website')
        }
      }
      // No web3 provider
      else {
        throw new Error("No web3 provider detected.\nDid you install the Metamask extension on this browser?")
      }
    },
    async getToken() {
      console.log('--getToken Start--')
      let token = null
      try {
        const {data} = await this.$axios.get(`/api/token/${this.web3.coinbase}`)
        console.log(data)
        token = data.token
      } catch (e) {
        // todo
        this.error = e
      }
      console.log('--getToken End--')
      return token
    },
    async signature(token) {
      console.log('--signature Start--')
      const message = `Signature for login authentication(token:${token})`
      try {
        // sign
        const instance = new Web3(window.ethereum)
        let signature = await instance.eth.personal.sign(message, this.web3.coinbase);

        // req
        await this.$axios.post('/api/auth', {
          address: this.web3.coinbase,
          signature: signature,
        })
      } catch (e) {
        this.errorMessage = ''
      }
      console.log('--signature End--')
    }
  }
}
</script>
