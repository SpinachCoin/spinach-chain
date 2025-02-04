package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Transaction struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Amount *big.Int `json:"amount"`
}

type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Coinbase     Transaction   `json:"coinbase"` // Block reward transaction
	PrevHash     string        `json:"prevHash"`
	Hash         string        `json:"hash"`
}

func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp + b.PrevHash
	for _, tx := range b.Transactions {
		txData, _ := json.Marshal(tx)
		record += string(txData)
	}
	coinbaseData, _ := json.Marshal(b.Coinbase)
	record += string(coinbaseData)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}