export const state = () => ({
  isListening: false,
  web3: {
    networkId: null,
    coinbase: null,
    balance: null,
  },
})

export const mutations = {
  init(state) {
    state.isListening = false
    state.web3.networkId = null
    state.web3.coinbase = null
    state.web3.balance = null
  },
  registerWeb3Instance (state, payload) {
    console.log('registerWeb3instance Mutation being executed', payload)
    let result = payload
    let web3Copy = state.web3
    console.log(result)
    web3Copy.coinbase = result.coinbase
    web3Copy.networkId = result.networkId
    web3Copy.balance = parseInt(result.balance, 10)
    state.web3 = web3Copy
  },
  updateIsListening(state, payload){
    state.isListening = payload
  }
}

export const actions = {}


export const getters = {
  getInstance: (state) => {
    return state.web3
  },
}
