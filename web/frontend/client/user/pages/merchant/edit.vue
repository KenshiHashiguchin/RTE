<template>
  <div id="pageMerchantEdit">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex lg12>
          <v-card>
            <v-toolbar color="secondary" dark flat dense cad>
              <v-toolbar-title className="subheading">Merchant</v-toolbar-title>
              <v-spacer></v-spacer>
            </v-toolbar>
            <v-divider></v-divider>
            <v-card-text className="">
              <shipment-form></shipment-form>
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VWidget from '@/components/VWidget';
import ShipmentForm from '@/components/widgets/form/ShipmentForm';

export default {
  layout: 'dashboard',
  components: {
    VWidget,
    ShipmentForm,
  },
  data() {
    return {};
  },
  computed: {},
  methods: {}
}
</script>

<style scoped>

</style>
