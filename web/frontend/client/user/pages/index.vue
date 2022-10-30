<template>
  <div id="pageDashboard">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex lg4 sm12 xs12>
          <vue-qrcode :value="qrCodeUrl"></vue-qrcode>
          <div>
            <p>Network: {{ web3.networkId }}</p>
            <p>Account: {{ web3.coinbase }}</p>
            <p>Balance: {{ web3.balance }}</p>
          </div>
        </v-flex>
        <v-flex lg4 sm12 xs12>
          <NuxtLink to="/merchant">merchants</NuxtLink>
          <NuxtLink to="/reserve">reserves</NuxtLink>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VueQrcode from '@chenfengyuan/vue-qrcode';
import API from '@/api';
import MiniStatistic from '@/components/widgets/statistic/MiniStatistic';
import PostListCard from '@/components/widgets/card/PostListCard';
import ProfileCard from '@/components/widgets/card/ProfileCard';
import PostSingleCard from '@/components/widgets/card/PostSingleCard';
import WeatherCard from '@/components/widgets/card/WeatherCard';
import PlainTable from '@/components/widgets/list/PlainTable';
import PlainTableOrder from '@/components/widgets/list/PlainTableOrder';
import VWidget from '@/components/VWidget';
import CircleStatistic from '@/components/widgets/statistic/CircleStatistic';
import LinearStatistic from '@/components/widgets/statistic/LinearStatistic';
import {mapGetters} from "vuex";

export default {
  // layouts: 'dashboard',
  components: {
    VueQrcode,
    VWidget,
    MiniStatistic,
    // ChatWindow,
    // VCircle,
    WeatherCard,
    PostSingleCard,
    PostListCard,
    ProfileCard,
    // EChart,
    // BoxChart,
    CircleStatistic,
    LinearStatistic,
    PlainTable,
    PlainTableOrder
  },
  data: () => ({
  }),
  computed: {
    ...mapGetters('web3', ['getInstance']),
    web3() {
      return this.getInstance
    },
    qrCodeUrl() {
      return 'https://google.com'
    }
  },
};
</script>
