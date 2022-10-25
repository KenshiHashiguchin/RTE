<template>
  <v-btn block color="primary" :loading="loading" @click="$emit('connect')">
    Connect to Metamask
  </v-btn>
</template>

<script>
export default {
  name: 'Metamask',
  data() {
    return {
      loading: false,
    }
  },
  computed: {},
  methods: {}
}
</script>
