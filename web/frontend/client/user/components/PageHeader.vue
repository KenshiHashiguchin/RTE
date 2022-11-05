<template>
  <v-app-bar app>
    <v-spacer></v-spacer>
    <client-only>
      <v-btn v-if="!isListening" dark color="primary" @click="connect">connect</v-btn>
      <v-btn v-else dark color="primary" @click="disconnect">disconnect</v-btn>
    </client-only>
  </v-app-bar>
</template>

<script>
import {mapState} from "vuex";
import web3Mixin from "~/mixins/web3Mixin";

export default {
  computed: {
    ...mapState('web3', ['isListening'])
  },
  mixins: [web3Mixin],
};
</script>
