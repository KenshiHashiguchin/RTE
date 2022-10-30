export const state = () => ({
  web3: {
    networkId: null,
    coinbase: null,
    balance: null,
  },
})

export const mutations = {
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
}

export const actions = {}


export const getters = {
  getInstance: (state) => {
    return state.web3
  },
}
