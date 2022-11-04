/* global ethers */
/* eslint prefer-const: "off" */

async function deployTres () {
  const Tres = await ethers.getContractFactory('TRES')
  const tres = await Tres.deploy(
    'TrustReserve Coin',
    'TRES',
    '0x86935f11c86623dec8a25696e1c19a8659cbf95d'
  )
  await tres.deployed()
  console.log('TRES deployed:', tres.address)
  return tres.address
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
if (require.main === module) {
  deployTres()
    .then(() => process.exit(0))
    .catch(error => {
      console.error(error)
      process.exit(1)
    })
}

exports.deployTres = deployTres
