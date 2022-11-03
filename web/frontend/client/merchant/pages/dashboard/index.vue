<template>
  <div id="pageDashboard">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex lg4 sm12 xs12>
          <client-only>
            <vue-qrcode :value="qrCodeUrl"></vue-qrcode>
          </client-only>
        </v-flex>
        <!-- social/weather card start -->
        <v-flex lg4 sm12 xs12>
          <profile-card>
          </profile-card>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VueQrcode from '@chenfengyuan/vue-qrcode';
import API from '~~/client/merchant/api';
import MiniStatistic from '~~/client/merchant/components/widgets/statistic/MiniStatistic';
import PostListCard from '~~/client/merchant/components/widgets/card/PostListCard';
import ProfileCard from '~~/client/merchant/components/widgets/card/ProfileCard';
import PostSingleCard from '~~/client/merchant/components/widgets/card/PostSingleCard';
import WeatherCard from '~~/client/merchant/components/widgets/card/WeatherCard';
import PlainTable from '~~/client/merchant/components/widgets/list/PlainTable';
import PlainTableOrder from '~~/client/merchant/components/widgets/list/PlainTableOrder';
import VWidget from '~~/client/merchant/components/VWidget';
import CircleStatistic from '~~/client/merchant/components/widgets/statistic/CircleStatistic';
import LinearStatistic from '~~/client/merchant/components/widgets/statistic/LinearStatistic';

export default {
  layout: 'dashboard',
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
  computed: {
    qrCodeUrl() {
      return 'https://google.com'
    }
  },
};
</script>
