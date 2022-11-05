<template>
  <v-app id="pageIndex">
    <page-header></page-header>
    <v-main>
      <v-container fill-height fluid>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4 lg4>
            <v-card class="pa-3 card-gradient">
              <v-card-text>
                <div class="layout column align-center">
                  <img width="100%" src="/dark-logo.png" alt="trust reserve logo"/>
                </div>
                <v-row align="center">
                  <v-col>Create public goods that keep power equitable.</v-col>
                </v-row>
              </v-card-text>
              <v-card-actions class="connect-wallet">
                <metamask @connect="login"/>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
      <Dialog :active="isErrorDialogActive"
              :message="errorMessage"
              @close="isErrorDialogActive = false"></Dialog>
    </v-main>
  </v-app>
</template>

<script>
import Metamask from "~/components/Metamask";
import Dialog from "~/components/common/Dialog";
import web3Mixin from "~/mixins/web3Mixin";
import loginMixin from "~/mixins/LoginMixin";

export default {
  components: {Dialog, Metamask},
  layout: 'default',
  mixins: [web3Mixin, loginMixin],
}
</script>
