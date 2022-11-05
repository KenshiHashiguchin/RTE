export default {
  data() {
    return {
      errorMessage: '',
      isErrorDialogActive: false,
    }
  },
  methods: {
    async login() {
      try {
        await this.connect()
        const token = await this.getToken()
        await this.signature(token)
        await this.$router.push('/dashboard')
      } catch (e) {
        console.log(e)
        this.isErrorDialogActive = true
        this.errorMessage = e.message
      }
    },
    async getToken() {
      let token = null
      try {
        const {data} = await this.$axios.get(`/api/token/${this.web3.coinbase}`)
        console.log(data)
        token = data.token
      } catch (e) {
        this.errorMessage = e.message
      }
      return token
    },
  },
}
