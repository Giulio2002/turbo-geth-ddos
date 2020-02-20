const privatekey = "87B10F7E1607FA3097666B8265E4B0A2FB77C92BEBBFB1352302AAEFC7B8D581"
const from = "0xbfb0444E5b0e33c8CA38332533D258A6B2Bb3dE8"
const gas = 7900000
const fs = require('fs')


async function deploy(contract, provider) {
    var block = await provider.eth.getBlock("latest");
    console.log("gasLimit: " + block.gasLimit);    
    const options = {
        chainId: 454546356,
        from,
        gasPrice: 20000,
        to: "0x0000000000000000000000000000000000000000",
        data: contract.options.bytecode,
        gas
    };
    const signedTransaction = await provider.eth.accounts.signTransaction(options, privatekey);
    const handle = await provider.eth.sendSignedTransaction(signedTransaction.rawTransaction);
    return new provider.eth.Contract(JSON.parse(abi), handle.contractAddress);
}

const deployContracts = async (provider) => {
    let contracts = []
    files = fs.readdirSync(__dirname + '/../build/');
    //listing all files using forEach
    for (let index = 0; index < files.length; index++) {
        const file = files[index]
        // Do whatever you want to do with the file
        const jsonContract = fs.readFileSync(__dirname + '/../build/' + file)
        const contract = JSON.parse(jsonContract)
        const attackContract = new provider.eth.Contract(contract.abi, options = {data: contract.bytecode})
        contracts.push({attackContract: await deploy(attackContract, provider), name: file})
    }
    return contracts
}

module.exports = {deployContracts}