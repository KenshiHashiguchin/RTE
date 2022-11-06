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
        <div v-if="step === cancelingStep" class="text-center my-4">
          <v-progress-circular
            indeterminate
            color="primary"
          ></v-progress-circular>
        </div>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          v-if="step === firstStep"
          color="primary"
          text
          :disabled="step === cancelingStep"
          @click="next"
        >
          yes
        </v-btn>
        <v-btn
          color="primary"
          text
          :disabled="step === cancelingStep"
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
    step: {
      type: Number,
      default: 1,
    },
    active: {
      type: Boolean,
      default: false,
    },
    message: {
      type: String,
      default: ''
    }
  },
  computed: {
    firstStep() {
      return 1
    },
    cancelingStep() {
      return 2
    },
    completeStep() {
      return 3
    },
    failedStep() {
      return 9
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
