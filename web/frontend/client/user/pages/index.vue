<template>
  <div id="pageIndex">
    <page-header></page-header>
    <v-main>
      <v-container fill-height class="py-5">
        <v-card width="400px" class="py-5 px-2 mx-auto mt-5 card-gradient">
          <v-card-text>
            <v-row align="center">
              <v-col
                class="text-h2"
                cols="8">Trust<br/>Reserve</v-col>
            </v-row>
            <v-row align="center">
              <v-col>Create public goods that keep power equitable.</v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-btn block dark rounded color="primary" @click="enter">Enter</v-btn>
          </v-card-actions>
        </v-card>
      </v-container>
    </v-main>
  </div>
</template>

<script>
import PageHeader from '@/components/PageHeader'

export default {
  // layouts: 'dashboard',
  components: {
    PageHeader,
  },
  methods: {
    enter() {
      this.$router.push('/admin')
    }
  }
};
</script>
