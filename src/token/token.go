package token

import (
	"math/big"
)

const (
	InitialSupply = 1000000000 // 1 billion SPIN
	Decimals      = 18         // 18 decimal places
)

type SpinachCoin struct {
	TotalSupply *big.Int
	Balances    map[string]*big.Int
}

func NewSpinachCoin() *SpinachCoin {
	totalSupply := big.NewInt(InitialSupply)
	totalSupply.Mul(totalSupply, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(Decimals), nil))

	return &SpinachCoin{
		TotalSupply: totalSupply,
		Balances:    make(map[string]*big.Int),
	}
}

func (sc *SpinachCoin) Mint(to string, amount *big.Int) {
	if sc.Balances[to] == nil {
		sc.Balances[to] = big.NewInt(0)
	}
	sc.Balances[to].Add(sc.Balances[to], amount)
}

func (sc *SpinachCoin) Transfer(from, to string, amount *big.Int) bool {
	if sc.Balances[from] == nil || sc.Balances[from].Cmp(amount) < 0 {
		return false // Insufficient balance
	}

	if sc.Balances[to] == nil {
		sc.Balances[to] = big.NewInt(0)
	}

	sc.Balances[from].Sub(sc.Balances[from], amount)
	sc.Balances[to].Add(sc.Balances[to], amount)
	return true
}

func (sc *SpinachCoin) BalanceOf(address string) *big.Int {
	if sc.Balances[address] == nil {
		return big.NewInt(0)
	}
	return sc.Balances[address]
}