package main

import (
	"fmt"
	"math/big"
	"spinach-chain/src/blockchain"
	"spinach-chain/src/token"
)

func main() {
	fmt.Println("Starting Spinach Chain...")

	// Initialize SpinachCoin
	spinachCoin := token.NewSpinachCoin()

	// Initialize blockchain
	bc := blockchain.NewBlockchain(spinachCoin)

	// Mine a block with a miner address
	minerAddress := "0xMinerAddress"
	bc.AddBlock([]blockchain.Transaction{}, minerAddress)

	// Mine another block
	bc.AddBlock([]blockchain.Transaction{
		{From: "0xAlice", To: "0xBob", Amount: big.NewInt(10 * 1e18)},
	}, minerAddress)

	// Print blockchain
	bc.Print()

	// Check miner's balance
	fmt.Printf("Miner Balance: %s SPIN\n", spinachCoin.BalanceOf(minerAddress).String())
}