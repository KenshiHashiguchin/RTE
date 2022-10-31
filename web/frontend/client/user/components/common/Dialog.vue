<template>
  <div class="text-center">
    <v-dialog v-model="active"
      width="500"
    >
      <v-card>
        <v-card-title class="text-h5">
          Error
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
            @click="close"
          >
            close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  name: "Dialog",
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
    close() {
      this.$emit('close')
    }
  }
}
</script>

<style scoped>

</style>
