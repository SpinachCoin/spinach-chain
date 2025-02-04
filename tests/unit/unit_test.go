package unit

import (
	"math/big"
	"spinach-chain/src/token"
	"testing"
)

func TestSpinachCoin(t *testing.T) {
	sc := token.NewSpinachCoin()

	// Test initial supply
	expectedSupply := new(big.Int).Mul(big.NewInt(1000000000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	if sc.TotalSupply().Cmp(expectedSupply) != 0 {
		t.Errorf("Initial supply is incorrect, got %s, want %s", sc.TotalSupply().String(), expectedSupply.String())
	}

	// Test minting
	address := "0xTestAddress"
	amount := new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	sc.Mint(address, amount)
	if sc.BalanceOf(address).Cmp(amount) != 0 {
		t.Errorf("Minting failed, got %s, want %s", sc.BalanceOf(address).String(), amount.String())
	}

	// Test transfer
	recipient := "0xRecipientAddress"
	if !sc.Transfer(address, recipient, amount) {
		t.Errorf("Transfer failed")
	}
	if sc.BalanceOf(recipient).Cmp(amount) != 0 {
		t.Errorf("Recipient balance is incorrect, got %s, want %s", sc.BalanceOf(recipient).String(), amount.String())
	}
}