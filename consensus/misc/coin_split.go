package misc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
)

// ApplyCoinSplitHardFork modifies the state database according to the CoinSplit hard-fork
// rules, inflate all balances of a set of accounts
func ApplyMainnetCoinSplitHardFork(statedb *state.StateDB) {
	// Inflate by adding balance for all accounts
	for _, acc := range params.MainnetCoinSplitAccount() {
		tmpBig := big.NewInt(0)
		tmpBig.SetString(acc.AddBalance, 10)
		statedb.AddBalance(acc.Address, tmpBig) /// StateProcessor use the reference, need to new BigInt object
	}
}

func ApplyTestnetCoinSplitHardFork(statedb *state.StateDB) {
	// Inflate by adding balance for all accounts
	for _, acc := range params.TestnetCoinSplitAccount() {
		tmpBig := big.NewInt(0)
		tmpBig.SetString(acc.AddBalance, 10)
		statedb.AddBalance(acc.Address, tmpBig) /// StateProcessor use the reference, need to new BigInt object
	}
}

// For local development
func ApplyLocalCoinSplitHardFork(statedb *state.StateDB) {
	for _, acc := range params.LocalCoinSplitAccount() {
		tmpBig := big.NewInt(0)
		tmpBig.SetString(acc.AddBalance, 10)
		statedb.AddBalance(acc.Address, tmpBig) /// StateProcessor use the reference, need to new BigInt object
	}
}
