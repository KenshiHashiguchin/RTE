<template>
  <div id="pageHistory">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex lg5 sm12 xs12>
          <plain-table-order></plain-table-order>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import PlainTableOrder from '@/components/widgets/list/PlainTableOrder';

export default {
  layout: 'dashboard',
  components: {
    PlainTableOrder,
  },
}
</script>

<style scoped>

</style>
