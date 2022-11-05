<template>
  <v-btn block color="primary" :loading="loading" @click="$emit('connect')">
    Enter
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
