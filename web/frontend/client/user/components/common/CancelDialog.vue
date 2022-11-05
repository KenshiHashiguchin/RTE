<template>
  <v-dialog v-model="active"
            @click:outside="close"
            @keydown.esc="close"
            width="500"
            class="text-center"
  >
    <v-card>
      <v-card-title class="text-h5">
        Confirmation
      </v-card-title>
      <v-card-text class="py-2">
        <p style="white-space:pre-wrap;" v-text="message"></p>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          color="primary"
          text
          @click="next"
        >
          yes
        </v-btn>
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
export default {
  name: "CancelDialog",
  props: {
    active: {
      type: Boolean,
      default: false,
    },
    message: {
      type: String,
      default: ''
    }
  },
  methods: {
    next() {
      this.$emit('next')
    },
    close() {
      this.$emit('close')
    }
  }
}
</script>

<style scoped>

</style>
