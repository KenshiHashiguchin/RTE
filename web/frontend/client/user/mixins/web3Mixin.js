import Web3 from "web3";
import trustReserveAbi from "~/constants/trust_reserve_abi.json";
import daiAbi from "~/constants/dai_abi.json";
import wethAbi from "~/constants/weth_abi.json";
import usdtAbi from "~/constants/usdt_abi.json";
import wmaticAbi from "~/constants/wmatic_abi.json";
import contractAddresses from "~/constants/contractAddresses.json";
import {mapGetters, mapMutations} from "vuex";

export default {
  computed: {
    ...mapGetters('web3', ['getInstance']),
    web3() {
      return this.getInstance
    }
  },
  methods: {
    ...mapMutations('web3', ['registerWeb3Instance']),
    async connect() {
      if (typeof window !== 'undefined' && typeof window.ethereum !== 'undefined') {
        try {
          // Ask to connect
          await window.ethereum.request({method: 'eth_requestAccounts'})
          const instance = this.createWeb3Instance(window.ethereum)
          const networkId = await instance.eth.net.getId();
          const coinbase = await instance.eth.getCoinbase();
          const balance = await instance.eth.getBalance(coinbase);
          this.registerWeb3Instance({
            networkId,
            coinbase,
            balance
          })
        } catch (error) {
          throw new Error('Please connect to your Ethereum address on Metamask before connecting to this website')
        }
      }
      // No web3 provider
      else {
        throw new Error("No web3 provider detected.\nDid you install the Metamask extension on this browser?")
      }
    },
    async getContract(web3Instance, contract = 'trustReserve') {
      console.log(contractAddresses)
      let abi = ''
      switch (contract) {
        case 'DAI':
          abi = daiAbi
          break
        case 'WETH':
          abi = wethAbi
          break
        case 'USDT':
          abi = usdtAbi
          break
        case 'WMATIC':
          abi = wmaticAbi
          break
        default:
          abi = trustReserveAbi
          break
      }
      return new web3Instance.eth.Contract(
        abi,
        contractAddresses.address
      )
    },
    createWeb3Instance(provider) {
      if (typeof window !== 'undefined' && typeof window.ethereum !== 'undefined') {
        return new Web3(provider)
      } else {
        throw new Error("No web3 provider detected.\nDid you install the Metamask extension on this browser?")
      }
    },
    async signature(token) {
      try {
        const message = `Signature for login authentication(token:${token})`
        // sign
        const instance = this.createWeb3Instance(window.ethereum)
        let signature = await instance.eth.personal.sign(message, this.web3.coinbase);

        // req
        await this.$axios.post('/api/auth', {
          address: this.web3.coinbase,
          signature: signature,
        })
      } catch (e) {
        console.log(e)
        throw new Error("failed signature")
      }
    },
    disconnect() {
      //  remove cookie
      this.$cookies.remove('user_token')
    }
  }
}
