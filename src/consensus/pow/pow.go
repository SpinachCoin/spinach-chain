package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

const targetBits = 24
const maxNonce = 1<<31 - 1

type Block struct {
	// Define the structure of Block here
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// Implement the method to prepare data for hashing
	return []byte{}
}

func (pow *ProofOfWork) Run() (int, string) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hex.EncodeToString(hash[:])
}