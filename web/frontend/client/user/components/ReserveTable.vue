<template>
  <v-data-table
    :headers="headers"
    :items="reserves"
    :disable-sort="true"
    :hide-default-footer="true"
    class="elevation-1"
  >
    <template v-slot:item.actions="{ item }">
      <v-btn v-if="settlementable(item.status)" x-small @click="settlement(item)" color="primary">Settlement</v-btn>
      <a :href="detailUrl(item)" target="_blank">
        <v-btn x-small :disabled="!item" type="button">Detail</v-btn>
      </a>
      <v-btn v-if="cancelable(item.status)" x-small @click="cancel(item)" color="accent">Cancel</v-btn>
    </template>
  </v-data-table>
</template>

<script>
export default {
  name: "ReserveTable",
  props: {
    reserves: {
      type: Array,
      default: () => ([])
    },
  },
  computed: {
    headers() {
      return [
        {
          text: 'Status',
          align: 'start',
          value: 'status',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Name',
          value: 'merchant.name',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Deposit',
          value: 'merchant.deposit',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Number of Visitors',
          value: 'number',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Cancelable Days (day)',
          value: 'merchant.cancelable_days',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Actions',
          value: 'actions',
          cellClass: 'text-no-wrap',
          sortable: false
        },
      ]
    },
  },
  methods: {
    detailUrl(item) {
      return item ? `/merchant/${item.address}/show` : ''
    },
    cancel(item) {
      this.$emit('cancel', item)
    },
    settlement(item) {
      this.$emit('settlement', item)
    },
    settlementable(status) {
      return status === 'Reserved'
    },
    cancelable(status) {
      return status === 'Reserved'
    }
  }
}
</script>

<style scoped>

</style>
