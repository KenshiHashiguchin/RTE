<template>
  <div>
    <div>
      API TEST
    </div>
    <div>
      <v-btn @click="getToken">getToken</v-btn>
      <v-btn @click="signature">signature</v-btn>
      <v-btn @click="register">register</v-btn>
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
     * GET /api/token/:address
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
     * 認証OKであればuser_sessionというキーでcookieにセット(Set-Cookie)する"
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
     * 店舗登録
     * POST	/api/register
     *
     * received_address(string)：受け取りアドレス
     * name(string) ：店舗名
     * tel(string)：電話番号
     * merchant_address(string)：住所
     * deposit(int)：必要デポジット金額
     * cancelable_time(int)：キャンセル可能期間
     * @returns {Promise<void>}
     */
    async register() {
      try {
        const {data} = await this.$axios.post('/api/register', {
            params: {
              received_address: 'hogehoge',
              name: 'fugafuga',
              tel: '08064263884',
              merchant_address: 'hoge',
              deposit: 3000,
              cancelable_time: '2022-12-01 18:00:00',
            }
          }
        )
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
    }
  },
}
</script>

<style scoped>

</style>
