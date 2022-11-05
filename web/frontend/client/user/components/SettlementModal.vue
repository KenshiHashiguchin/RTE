<template>
  <v-dialog v-model="active"
            @click:outside="close"
            @keydown.esc="close"
            width="350"
            class="text-center"
  >
    <v-card>
      <!-- todo merchant name-->
      <v-card-title class="mb-5 text-h4 justify-center">
        Send To {{toName}}
      </v-card-title>

      <v-card-text class=text-center>
        <div
          class=" text-h2 inline-block"
        >
          {{ input.current }}
        </div>
      </v-card-text>
      <v-card-actions class="justify-center my-5">
        <v-btn @click="next" color="primary">next</v-btn>
      </v-card-actions>
      <v-card-text class="pa-0">
        <v-container grid-list-xs pa-1>
          <v-layout row wrap pa0>
            <v-flex v-for="button in buttons" :key="button.key" xs4 class="text-center mb-2">
              <v-btn v-if="button.key"
                     :class="['ma-0', button.class]"
                     :color="button.color"
                     @click="inputKey(button.key)"
                     fab
                     large
              >
                <v-icon v-if="button.icon" dark>{{ button.icon }}</v-icon>
                <template v-else>{{ button.label }}</template>
              </v-btn>
            </v-flex>
          </v-layout>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          color="primary"
          text
          @click="close"
        >
          close
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import {get} from 'lodash'

export default {
  name: "SettlementModal",
  data() {
    return {
      input: {
        current: "0",
        operator: "",
        prev: "",
      },
    }
  },
  props: {
    active: {
      type: Boolean,
      default: false,
    },
    reservation: {
      type: Object,
      default: () => ({})
    }
  },
  computed: {
    toName() {
      return get(this.reservation, 'merchant.name', '')
    },
    buttons() {
      return [
        {key: "1", label: "1", color: "", class: "headline"},
        {key: "2", label: "2", color: "", class: "headline"},
        {key: "3", label: "3", color: "", class: "headline"},
        {key: "4", label: "4", color: "", class: "headline"},
        {key: "5", label: "5", color: "", class: "headline"},
        {key: "6", label: "6", color: "", class: "headline"},
        {key: "7", label: "7", color: "", class: "headline"},
        {key: "8", label: "8", color: "", class: "headline"},
        {key: "9", label: "9", color: "", class: "headline"},
        {key: "0", label: "0", color: "", class: "headline"},
        {key: "ac", label: "C", color: "grey lighten-2", class: ""},
        {
          key: "back",
          label: "BACK",
          icon: "backspace",
          color: "grey lighten-2",
          class: "",
        },
      ];
    }
  },
  methods: {
    close() {
      this.$emit('close')
    },
    inputKey(key) {
      const isNumber = /^\d$/.test(key);
      if (isNumber) {
        if (this.input.current.length >= 12) {
          return;
        }
        if (this.input.current !== "0") {
          this.input.current += key;
        } else {
          this.input.current = key;
        }
        return
      }
      switch (key) {
        case "back":
          this.input.current =
            this.input.current.substring(0, this.input.current.length - 1) || "0";
          return;
        case "c":
          this.input.current = "0";
          this.input.operator = "";
          this.input.prev = "";
          return;
        case "ce":
          this.input.current = "0";
          return;
      }
    },
    next() {

    }
  },
}
</script>

<style scoped>

</style>
