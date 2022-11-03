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
              <merchant-form></merchant-form>
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VWidget from '@/components/VWidget';
import MerchantForm from '@/components/widgets/form/MerchantForm';

export default {
  layout: 'dashboard',
  components: {
    VWidget,
    MerchantForm,
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
