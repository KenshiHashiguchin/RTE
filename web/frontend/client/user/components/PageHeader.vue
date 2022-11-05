<template>
  <v-app-bar app>
    <img src="/dark-logo.png" height="36" alt="Trust Reserve">
    <v-spacer></v-spacer>
    <client-only>
      <v-btn v-if="!isListening" dark color="primary" @click="login">connect</v-btn>
      <v-btn v-else dark color="primary" @click="disconnect">disconnect</v-btn>
    </client-only>
    <Dialog :active="isErrorDialogActive"
            :message="errorMessage"
            @close="isErrorDialogActive = false"></Dialog>
  </v-app-bar>
</template>

<script>
import {mapState} from "vuex";
import web3Mixin from "~/mixins/web3Mixin";
import loginMixin from "~~/client/user/mixins/LoginMixin";

export default {
  computed: {
    ...mapState('web3', ['isListening'])
  },
  mixins: [web3Mixin, loginMixin],
};
</script>
