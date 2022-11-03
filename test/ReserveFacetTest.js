/* global describe it before ethers */

const {deployDiamond} = require('../scripts/deploy.js')

const {expect} = require('chai')
const {waffle, ethers} = require('hardhat');
const {deployMockContract, provider} = waffle;

describe('ReserveFacetTest', async function () {
  let diamondAddress
  let diamondCutFacet
  let diamondLoupeFacet
  let ownershipFacet
  let reserveFacet
  let owner
  let swapAddress
  let subscriberAddress
  let merchantAddress
  let addrs
  let paymentId = 'asNfY93lMNh'
  let token
  let depositAmount = 300
  let cancelableTime = 20
  let withdrawableTime = 40
  let mockedERC20Contract
  let status = {None: 0, Reserved: 1, Canceled: 2, Settled: 3};
  let now = parseInt(Date.now() / 1000)
  let mockedRouterContract

  before(async function () {
    [owner, swapAddress, subscriberAddress, merchantAddress, token, ...addrs] = await ethers.getSigners()
    diamondAddress = await deployDiamond()
    diamondCutFacet = await ethers.getContractAt('DiamondCutFacet', diamondAddress)
    diamondLoupeFacet = await ethers.getContractAt('DiamondLoupeFacet', diamondAddress)
    ownershipFacet = await ethers.getContractAt('OwnershipFacet', diamondAddress)
    reserveFacet = await ethers.getContractAt('ReserveFacet', diamondAddress)

    // interface
    const [deployerOfERC20Contract, deployerOfRouterContract] = provider.getWallets();
    // deploy the contract to Mock
    const ERC20Contract = require('../artifacts/@openzeppelin/contracts/token/ERC20/IERC20.sol/IERC20.json');
    mockedERC20Contract = await deployMockContract(deployerOfERC20Contract, ERC20Contract.abi);
    await mockedERC20Contract.mock.transfer.returns(true);
    await mockedERC20Contract.mock.transferFrom.returns(true);
    await mockedERC20Contract.mock.approve.returns(true);

    const IUniswapV2Router02 = require('../artifacts/contracts/interfaces/IUniswapV2Router02.sol/IUniswapV2Router02.json');
    mockedRouterContract = await deployMockContract(deployerOfRouterContract, IUniswapV2Router02.abi);
    await mockedRouterContract.mock.swapExactTokensForTokens.returns([]);
    await mockedRouterContract.mock.swapTokensForExactTokens.returns([]);
  })

  it('owner can change swapSubmitAddress', async () => {
    console.log(await reserveFacet.swapSubmitAddress());
    expect(await reserveFacet.swapSubmitAddress()).not.equal(mockedRouterContract.address)
    await reserveFacet.transferSwapSubmitAddress(mockedRouterContract.address)
    expect(await reserveFacet.swapSubmitAddress()).equal(mockedRouterContract.address)
  })

  it('reserve', async () => {
    await reserveFacet.connect(subscriberAddress).reserve(paymentId, merchantAddress.address, mockedERC20Contract.address, depositAmount, now + cancelableTime, now + withdrawableTime)

    // check
    // Already exist paymentId
    await expect(reserveFacet.reserve(paymentId, merchantAddress.address, mockedERC20Contract.address, depositAmount, now + cancelableTime, now + withdrawableTime)).revertedWith("Already exist paymentId")

    let reservation = await reserveFacet.getReservation(paymentId)
    expect(reservation.merchant).equal(merchantAddress.address)
    expect(reservation.subscriber).equal(subscriberAddress.address)
    expect(reservation.token).equal(mockedERC20Contract.address)
    expect(reservation.depositAmount).equal(depositAmount)
    expect(reservation.additionalAmount).equal(0)
    expect(reservation.cancelableTime).equal(now + cancelableTime)
    expect(reservation.withdrawableTime).equal(now + withdrawableTime)
    expect(reservation.status).equal(status.Reserved)
  })

  it('cancel', async () => {
    await expect(reserveFacet.cancel(paymentId)).revertedWith("Must be subscriber owner")
    await ethers.provider.send("evm_increaseTime", [cancelableTime]);
    await ethers.provider.send("evm_mine", []);
    await expect(reserveFacet.connect(subscriberAddress).cancel(paymentId)).revertedWith("Cancellation period has passed")
    await ethers.provider.send("evm_increaseTime", [-cancelableTime]);
    await ethers.provider.send("evm_mine", []);
    await reserveFacet.connect(subscriberAddress).cancel(paymentId)
    let reservation = await reserveFacet.getReservation(paymentId)
    expect(reservation.status).equal(status.Canceled)
  })

  it('withdrawDeposit', async () => {
    // create reservation
    let withdrawPaymentId = "test_withdrawDeposit"
    await reserveFacet.connect(subscriberAddress).reserve(withdrawPaymentId, merchantAddress.address, mockedERC20Contract.address, depositAmount, now + cancelableTime, now + withdrawableTime)
    let reservation = await reserveFacet.getReservation(withdrawPaymentId)
    expect(reservation.status).equal(status.Reserved)

    // withdrawDeposit
    await expect(reserveFacet.connect(subscriberAddress).withdrawDeposit(10, 10000, withdrawPaymentId, [])).revertedWith("Must be merchant owner")
    await expect(reserveFacet.connect(merchantAddress).withdrawDeposit(10, 10000, withdrawPaymentId, [])).revertedWith("Still can't pull it out")
    await ethers.provider.send("evm_increaseTime", [withdrawableTime]);
    await ethers.provider.send("evm_mine", []);
    await reserveFacet.connect(merchantAddress).withdrawDeposit(10, 10000, withdrawPaymentId, [])
    reservation = await reserveFacet.getReservation(withdrawPaymentId)
    expect(reservation.status).equal(status.Canceled)
    await ethers.provider.send("evm_increaseTime", [-withdrawableTime]);
    await ethers.provider.send("evm_mine", []);
  })

  it('settleReservation', async () => {
    // already canceled
    await expect(reserveFacet.connect(subscriberAddress).settleReservation(5, 5, 10000, paymentId, [])).revertedWith("This transaction has been closed")

    // create reservation
    let settlePaymentId = "test_settlePaymentId"
    await reserveFacet.connect(subscriberAddress).reserve(settlePaymentId, merchantAddress.address, mockedERC20Contract.address, depositAmount, now + cancelableTime, now + withdrawableTime)
    let reservation = await reserveFacet.getReservation(settlePaymentId)
    expect(reservation.status).equal(status.Reserved)

    let additionalAmount = 500
    await reserveFacet.connect(subscriberAddress).settleReservation(additionalAmount+depositAmount, 5, 10000, settlePaymentId, [])
    reservation = await reserveFacet.getReservation(settlePaymentId)
    expect(reservation.additionalAmount).equal(additionalAmount)
    expect(reservation.status).equal(status.Settled)
  })
})