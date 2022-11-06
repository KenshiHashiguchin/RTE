import erc20Abi from "~/constants/erc20_abi.json";

const TOKEN_CONTRACT_ADDRESSES = [
  {
    name: 'DAI',
    address: '0x001B3B4d0F3714Ca98ba10F6042DaEbF0B1B7b6F',
    decimals: 6,
  },
  {
    name: 'WETH',
    address: '0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa',
    decimals: 18,
  },
  {
    name: 'USDT',
    address: '0x3813e82e6f7098b9583FC0F33a962D02018B6803',
    decimals: 18,
  },
  {
    name: 'WMATIC',
    address: '0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889',
    decimals: 18,
  },
]

function getAddress(name){
  const tokenAddress = TOKEN_CONTRACT_ADDRESSES.find(v => v.name === name)
  return tokenAddress ? tokenAddress.address : ''
}
function getAbi(name){
  let abi = ''
  switch (name) {
    case 'DAI':
    case 'WETH':
    case 'USDT':
    case 'WMATIC':
      abi = erc20Abi
      break
    default:
      break;
  }
  return abi
}

export default {
  TOKEN_CONTRACT_ADDRESSES,
  getAddress,
  getAbi,
}
