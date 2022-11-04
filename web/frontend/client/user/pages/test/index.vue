<template>
  <v-main>
    <v-container fluid>
      <h1>
        API&CONTRACT TEST
      </h1>
      <div class="mx-auto">
        <h2>CONTRACT</h2>
        <v-text-field label="paymentId(string)" v-model="paymentId"/>
        <div v-if="reservation">
          <p>additionalAmount : {{ reservation.additionalAmount }}</p>
          <p>cancelableTime : {{ reservation.cancelableTime }}</p>
          <p>depositAmount : {{ reservation.depositAmount }}</p>
          <p>merchant : {{ reservation.merchant }}</p>
          <p>status : {{ reservation.status }}</p>
          <p>subscriber : {{ reservation.subscriber }}</p>
          <p>token : {{ reservation.token }}</p>
          <p>withdrawableTime : {{ reservation.withdrawableTime }}</p>
        </div>
        <v-btn @click="getReservation">getReservation</v-btn>
        <v-divider class="my-2"></v-divider>
        <h2>API</h2>
        <v-btn @click="getToken">getToken</v-btn>
        <v-btn @click="signature">signature</v-btn>
        <v-btn @click="merchants">merchants</v-btn>
        <v-btn @click="merchant">merchant</v-btn>
        <v-btn @click="reserve">reserve</v-btn>
        <v-btn @click="reserves">reserves</v-btn>
      </div>
    </v-container>
  </v-main>
</template>

<script>
import Web3 from "web3";
import web3Mixin from "@/mixins/web3Mixin";

export default {
  name: "TestIndex",
  mixins: [web3Mixin],
  data() {
    return {
      paymentId: 'test',
      reservation: null,
    }
  },
  methods: {
    async getReservation() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const contract = await this.getContract(instance)
        const res = await contract.methods.getReservation(this.paymentId).call()
        console.log(res)
        this.reservation = res
      } catch (error) {
        console.log(error)
      }
    },
    /**
     * 署名用トークンを取得
     * GET  /api/token/:address
     *
     */
    async getToken() {
      try {
        const address = 'hoge'
        const {data} = await this.$axios.get(`/api/token/${address}`)
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
    /**
     * 認証 取得したtokenを署名
     * POST  /api/auth
     *
     * 署名するメッセージ：「Signature for login authentication(token:{{署名用トークン}})」
     *
     * 認証OKであればuser_sessionというキーでcookieにセット(Set-Cookie)する
     * payload : address(string), signature(string)
     */
    async signature() {
      try {
        const {data} = await this.$axios.post('/api/auth', {
            params: {}
          }
        )
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
    /**
     * 店舗一覧取得
     * GET  /api/merchants
     *
     */
    async merchants() {
      try {
        const {data} = await this.$axios.get('/api/merchants')
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
    /**
     * 店舗取得
     * GET  /api/merchant/:address
     */
    async merchant(address) {
      try {
        const {data} = await this.$axios.get(`/api/merchant/${address}`)
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
    /**
     * 予約登録
     * POST  /api/reserve
     *
     * params :{
     *    merchant_id(uint): 予約先事業者ID
     *    surname(string): 予約者姓
     *    firstname(string): 予約者名
     *    phonenumber(string): 予約者電話番号
     *    number(int): 予約人数
     *    time(timestamp):予約日時"
     * }
     */
    async reserve() {
      try {
        const {data} = await this.$axios.post('/api/reserve', {
          params: {
            merchant_id: 1,
            surname: 'hogehoge',
            firstname: 'fugafuga',
            phonenumber: '08064263884',
            number: 2,
            time: '2022-12-01 18:00:00',
          }
        })
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
    /**
     * 予約一覧
     * GET  /api/reserve_list
     */
    async reserves() {
      try {
        const {data} = await this.$axios.get('/api/reserve_list')
        console.log(data)
      } catch (e) {
        // todo
        this.error = e
      }
    },
  }
}
</script>
