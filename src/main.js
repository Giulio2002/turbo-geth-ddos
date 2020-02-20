const Web3 = require('web3');
const attack =  require('./attack.js')

const main = async () => {
    var web3 = new Web3('http://127.0.0.1:8545');
    const datas = []
    console.log(`Current Block Number ${await web3.eth.getBlockNumber()}`)
    const contracts = await attack.deployContracts(web3)
    console.log(contracts)
}

main()