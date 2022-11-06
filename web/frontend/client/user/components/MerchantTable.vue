<template>
  <v-data-table
    :headers="headers"
    :items="merchants"
    :disable-sort="true"
    :hide-default-footer="true"
    @delete=""
    @delete-confirm=""
    class="elevation-1"
  >
    <template v-slot:item.actions="{ item }">
      <ReserveDialog :merchant="item">Reserve</ReserveDialog>
      <a :href="`/merchant/${item.address}/show`" target="_blank">
        <v-btn x-small type="button">Detail</v-btn>
      </a>
    </template>
  </v-data-table>
</template>

<script>

import ReserveDialog from "~/components/ReserveDialog";
export default {
  name: "MerchantTable",
  components: {ReserveDialog},
  props: {
    merchants: {
      type: Array,
      default: () => ([])
    }
  },
  computed: {
    headers() {
      return [
        {
          text: 'Name',
          align: 'start',
          value: 'name',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Deposit Amount',
          value: 'deposit',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'WalletAddress',
          value: 'address',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'ERC20 ContractAddress',
          value: 'received_address',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'PhoneNumber',
          value: 'tel',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Address',
          value: 'merchant_address',
          class: 'text-no-wrap',
          cellClass: 'text-no-wrap',
        },
        {
          text: 'Cancelable Days (day)',
          value: 'cancelable_days',
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
    reserve(item) {
      this.$emit('reserve', item)
    },
    // goToDetail(item) {
    //   this.$emit('go-to-detail', item)
    // }
  }
}
</script>

<style scoped>

</style>
