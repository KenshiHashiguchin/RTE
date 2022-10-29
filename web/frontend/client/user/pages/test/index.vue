<template>
  <div>
    <div>
      API TEST
    </div>
    <div>
      <v-btn @click="getToken">getToken</v-btn>
      <v-btn @click="signature">signature</v-btn>
      <v-btn @click="merchants">merchants</v-btn>
      <v-btn @click="merchant">merchant</v-btn>
      <v-btn @click="reserve">reserve</v-btn>
      <v-btn @click="reserves">reserves</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: "TestIndex",
  methods: {
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

<style scoped>

</style>
