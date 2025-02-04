package blockchain

import (
	"math/big"
	"spinach-chain/src/token"
	"time"
)

const (
	InitialBlockReward = 50 * 1e18 // 50 SPIN (with 18 decimals)
	HalvingInterval    = 210000    // Halve reward every 210,000 blocks
)

type Blockchain struct {
	Blocks       []*Block
	SpinachCoin  *token.SpinachCoin
	BlockReward  *big.Int
}

func NewBlockchain(spinachCoin *token.SpinachCoin) *Blockchain {
	return &Blockchain{
		Blocks:      []*Block{GenesisBlock()},
		SpinachCoin: spinachCoin,
		BlockReward: big.NewInt(InitialBlockReward),
	}
}

func GenesisBlock() *Block {
	return &Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		Coinbase:     Transaction{To: "0xGenesisAddress", Amount: big.NewInt(0)},
		PrevHash:     "",
		Hash:         "",
	}
}

func (bc *Blockchain) AddBlock(transactions []Transaction, minerAddress string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	// Create coinbase transaction (block reward)
	coinbaseTx := Transaction{
		To:     minerAddress,
		Amount: new(big.Int).Set(bc.BlockReward),
	}

	// Mint new SPIN for the block reward
	bc.SpinachCoin.Mint(minerAddress, bc.BlockReward)

	// Create new block
	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		Coinbase:     coinbaseTx,
		PrevHash:     prevBlock.Hash,
		Hash:         "",
	}
	newBlock.Hash = newBlock.CalculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)

	// Halve the block reward if necessary
	if newBlock.Index%HalvingInterval == 0 {
		bc.BlockReward.Div(bc.BlockReward, big.NewInt(2))
	}
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Transactions: %+v\n", block.Transactions)
		fmt.Printf("Coinbase: %+v\n", block.Coinbase)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}