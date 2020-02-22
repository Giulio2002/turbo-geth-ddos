package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Giulio2002/turbo-geth-ddos/contracts"
	"github.com/ledgerwatch/turbo-geth/accounts/abi/bind"
	"github.com/ledgerwatch/turbo-geth/accounts/abi/bind/backends"
	"github.com/ledgerwatch/turbo-geth/consensus/ethash"
	"github.com/ledgerwatch/turbo-geth/core"
	"github.com/ledgerwatch/turbo-geth/core/types"
	"github.com/ledgerwatch/turbo-geth/core/vm"
	"github.com/ledgerwatch/turbo-geth/crypto"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/params"
)

func main() {
	// Configure and generate a sample block chain
	var (
		db, _   = ethdb.NewMemDatabase2()
		key, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		address = crypto.PubkeyToAddress(key.PublicKey)
		gspec   = &core.Genesis{
			Config: &params.ChainConfig{
				ChainID:             big.NewInt(1),
				HomesteadBlock:      new(big.Int),
				EIP150Block:         big.NewInt(1),
				EIP155Block:         big.NewInt(1),
				EIP158Block:         big.NewInt(1),
				ByzantiumBlock:      big.NewInt(1),
				ConstantinopleBlock: big.NewInt(1),
			},
			Alloc: core.GenesisAlloc{
				address: {Balance: big.NewInt(9000000000000000000)},
			},
		}
	)
	genesis := gspec.MustCommit(db)
	genesisDb := db.MemCopy()

	engine := ethash.NewFaker()
	blockchain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	var suicide *contracts.Suicide
	var extcode *contracts.Extcodesize
	blockchain.EnableReceipts(true)
	contractBackend := backends.NewSimulatedBackendWithConfig(gspec.Alloc, gspec.Config, gspec.GasLimit)

	transactOpts := bind.NewKeyedTransactor(key)
	transactOpts.GasPrice = big.NewInt(1)
	transactOpts.GasLimit = 790000 // 7.9 Million of gas

	blocks, _ := core.GenerateChain(context.Background(), gspec.Config, genesis, engine, genesisDb, 2, func(i int, block *core.BlockGen) {
		var (
			tx  *types.Transaction
			txs []*types.Transaction
		)

		switch i {
		case 0:
			_, tx, suicide, _ = contracts.DeploySuicide(transactOpts, contractBackend)
		case 1:
			_, tx, extcode, _ = contracts.DeployExtcodesize(transactOpts, contractBackend)
		}
		if txs == nil && tx != nil {
			txs = append(txs, tx)
		}
		for _, tx := range txs {
			block.AddTx(tx)
		}
		contractBackend.Commit()
	})
	fmt.Println("Gas for Selfdestruct")
	blockchain.InsertChain(context.Background(), types.Blocks{blocks[0]})
	_, _ = suicide.Start(transactOpts)
	receipts := blockchain.GetReceiptsByHash(blockchain.GetBlockByNumber(1).Hash())
	fmt.Println(receipts[0].GasUsed - 21000)
	fmt.Println("Gas for Extcodesize")
	blockchain.InsertChain(context.Background(), types.Blocks{blocks[1]})
	gas, _ := extcode.Start(&bind.CallOpts{})
	fmt.Println(gas.Uint64())
}
