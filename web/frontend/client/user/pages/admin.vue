<template>
  <v-app id="login">
    <page-header></page-header>
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4 lg4>
            <v-card class="pa-3 card-gradient">
              <v-card-text>
                <div class="layout column align-center">
                  <img width="100%" src="/dark-logo.png" alt="trust reserve logo"/>
                </div>
              </v-card-text>
              <v-card-actions class="connect-wallet">
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
import {mapGetters, mapMutations} from 'vuex';
import Metamask from "~/components/Metamask";
import Dialog from "~/components/common/Dialog";
import web3Mixin from "~/mixins/web3Mixin";

export default {
  components: {Dialog, Metamask},
  layout: 'default',
  mixins: [web3Mixin],
  data() {
    return {
      errorMessage: '',
      isErrorDialogActive: false,
    }
  },
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
    async getToken() {
      let token = null
      try {
        const {data} = await this.$axios.get(`/api/token/${this.web3.coinbase}`)
        console.log(data)
        token = data.token
      } catch (e) {
        this.errorMessage = e.message
      }
      return token
    },
  }
}
</script>
<style scoped lang="css">
</style>
