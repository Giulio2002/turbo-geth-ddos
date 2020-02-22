
const gas = 6000000
const fs = require('fs')
const Attack = require('../Attack.json')
let from;

const runAttacks = async (provider) => {
    from = (await provider.eth.getAccounts())[0]
    files = fs.readdirSync(__dirname + '/../build/');
    //listing all files using forEach
    for (let index = 0; index < files.length; index++) {
        const file = files[index]
        if (file.indexOf('Attack') === -1) {
            continue
        }
        console.log("Now doing " + file.replace('.json', ''))
        // Do whatever you want to do with the file
        const jsonContract = fs.readFileSync(__dirname + '/../build/' + file)
        const contract = JSON.parse(jsonContract)
        const attackContract = new provider.eth.Contract(contract.abi, options = {data: '0x' + contract.bytecode})

        const attack = await attackContract.deploy({
            arguments: []
        })
        .send({
            from,
            gas,
            gasPrice: '30000000000000'
        })
        try { 
        await attack.methods.loop().call({
            from,
            gas,
            gasPrice: '30000000000000'
        }) } catch(e) {
            console.error(e)
        }
        console.log("End of " + file.replace('.json', ''))
    }

}

module.exports = {runAttacks}