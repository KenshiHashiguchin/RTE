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

export default {
  computed: {
    ...mapState('web3', ['isListening'])
  },
  methods: {
    connect() {
      // todo
      // call common func
    },
    disconnect() {
      // todo
      //  remove cookie
    }
  }
};
</script>
