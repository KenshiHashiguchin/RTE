<template>
  ERROR
</template>
<script>
export default {
  props: {
    error: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  computed: {
    status() {
      return (code) => {
        return this.error.statusCode === code
      }
    },
  },
  head: {
    title: 'ERROR',
  },
}
</script>
