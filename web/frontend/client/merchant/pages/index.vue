<template>
  <div id="pageDashboard">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>
        <v-flex lg4 sm12 xs12>
        <vue-qrcode :value="qrCodeUrl" :options="{ width: 200 }"></vue-qrcode>
        <!--TODO download-->
        </v-flex>
        <!-- social/weather card start -->
        <v-flex lg4 sm12 xs12>
          <profile-card>
          </profile-card>
        </v-flex>
        <!--&lt;!&ndash; statistic section &ndash;&gt;-->
        <!--<v-flex lg4 sm12 xs12>-->
        <!--  <linear-statistic-->
        <!--    title="Sales"-->
        <!--    sub-title="Sales increase"-->
        <!--    icon="trending_up"-->
        <!--    color="success"-->
        <!--    :value="15"-->
        <!--  >-->
        <!--  </linear-statistic>-->
        <!--  <linear-statistic class="my-4"-->
        <!--                    title="Orders"-->
        <!--                    sub-title="Increase"-->
        <!--                    icon="trending_up"-->
        <!--                    color="pink"-->
        <!--                    :value="30"-->
        <!--  >-->
        <!--  </linear-statistic>-->
        <!--  <linear-statistic class="my-4"-->
        <!--                    title="Revenue"-->
        <!--                    sub-title="Revenue increase"-->
        <!--                    icon="trending_up"-->
        <!--                    color="primary"-->
        <!--                    :value="50"-->
        <!--  >-->
        <!--  </linear-statistic>-->
        <!--  <linear-statistic class="mt-4"-->
        <!--                    title="Cost"-->
        <!--                    sub-title="Cost reduce"-->
        <!--                    icon="trending_down"-->
        <!--                    color="orange"-->
        <!--                    :value="25"-->
        <!--  >-->
        <!--  </linear-statistic>-->
        <!--</v-flex>-->
        <!--&lt;!&ndash; Circle statistic &ndash;&gt;-->
        <!--<v-flex lg4 sm12 xs12 v-for="(item,index) in trending" :key="'c-trending'+index">-->
        <!--  <circle-statistic-->
        <!--    :title="item.subheading"-->
        <!--    :sub-title="item.headline"-->
        <!--    :caption="item.caption"-->
        <!--    :icon="item.icon.label"-->
        <!--    :color="item.linear.color"-->
        <!--    :value="item.linear.value"-->
        <!--  >-->
        <!--  </circle-statistic>-->
        <!--</v-flex>-->
        <!--&lt;!&ndash; acitivity/chat widget &ndash;&gt;-->
        <!--&lt;!&ndash;<v-flex lg6 sm12 xs12>&ndash;&gt;-->
        <!--&lt;!&ndash;  <chat-window height="308px"></chat-window>&ndash;&gt;-->
        <!--&lt;!&ndash;</v-flex>&ndash;&gt;-->
        <!--<v-flex lg6 sm12 xs12>-->
        <!--  <v-widget title="Activities" contentBg="white">-->
        <!--    <div slot="widget-content">-->
        <!--      <ol class="timeline timeline-activity timeline-point-sm timeline-content-right">-->
        <!--        <li class="timeline-block" v-for="(item, index) in activity" :key="index">-->
        <!--          &lt;!&ndash;<div class="timeline-point">&ndash;&gt;-->
        <!--          &lt;!&ndash;  <v-circle dot large :color="item.color"></v-circle>&ndash;&gt;-->
        <!--          &lt;!&ndash;</div>&ndash;&gt;-->
        <!--          <div class="timeline-content">-->
        <!--            <time datetime="2018" class="subheading">{{item.timeString}}</time>-->
        <!--            <div class="py-2 text&#45;&#45;secondary" v-html="item.text"></div>-->
        <!--          </div>-->
        <!--        </li>-->
        <!--      </ol>-->
        <!--    </div>-->
        <!--  </v-widget>-->
        <!--</v-flex>-->
        <!--<v-flex lg7 sm12 xs12>-->
        <!--  <plain-table></plain-table>-->
        <!--</v-flex>-->
        <!--<v-flex lg5 sm12 xs12>-->
        <!--  <plain-table-order></plain-table-order>-->
        <!--</v-flex>-->
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import VueQrcode from '@chenfengyuan/vue-qrcode';
import API from '@/api';
// import EChart from '@/components/chart/echart';
import MiniStatistic from '@/components/widgets/statistic/MiniStatistic';
import PostListCard from '@/components/widgets/card/PostListCard';
import ProfileCard from '@/components/widgets/card/ProfileCard';
import PostSingleCard from '@/components/widgets/card/PostSingleCard';
import WeatherCard from '@/components/widgets/card/WeatherCard';
import PlainTable from '@/components/widgets/list/PlainTable';
import PlainTableOrder from '@/components/widgets/list/PlainTableOrder';
import VWidget from '@/components/VWidget';
import Material from 'vuetify/es5/util/colors';
// import VCircle from '@/components/circle/VCircle';
// import BoxChart from '@/components/widgets/chart/BoxChart';
// import ChatWindow from '@/components/chat/ChatWindow';
import CircleStatistic from '@/components/widgets/statistic/CircleStatistic';
import LinearStatistic from '@/components/widgets/statistic/LinearStatistic';

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
  data: () => ({
    color: Material,
    selectedTab: 'tab-1',
    linearTrending: [
      {
        subheading: 'Sales',
        headline: '2,55',
        caption: 'increase',
        percent: 15,
        icon: {
          label: 'trending_up',
          color: 'success'
        },
        linear: {
          value: 15,
          color: 'success'
        }
      },
      {
        subheading: 'Revenue',
        headline: '6,553',
        caption: 'increase',
        percent: 10,
        icon: {
          label: 'trending_down',
          color: 'error'
        },
        linear: {
          value: 15,
          color: 'error'
        }
      },
      {
        subheading: 'Orders',
        headline: '5,00',
        caption: 'increase',
        percent: 50,
        icon: {
          label: 'arrow_upward',
          color: 'info'
        },
        linear: {
          value: 50,
          color: 'info'
        }
      }
    ],
    trending: [
      {
        subheading: 'Email',
        headline: '15+',
        caption: 'email opens',
        percent: 15,
        icon: {
          label: 'email',
          color: 'info'
        },
        linear: {
          value: 15,
          color: 'info'
        }
      },
      {
        subheading: 'Tasks',
        headline: '90%',
        caption: 'tasks completed.',
        percent: 90,
        icon: {
          label: 'list',
          color: 'primary'
        },
        linear: {
          value: 90,
          color: 'success'
        }
      },
      {
        subheading: 'Issues',
        headline: '100%',
        caption: 'issues fixed.',
        percent: 100,
        icon: {
          label: 'bug_report',
          color: 'primary'
        },
        linear: {
          value: 100,
          color: 'error'
        }
      },
    ]
  }),
  computed: {
    activity () {
      return API.getActivity();
    },
    posts () {
      return [];
    },
    siteTrafficData () {
      return [];
    },
    locationData () {
      return [];
    },
    qrCodeUrl() {
      return 'https://google.com'
    }
  },
};
</script>
