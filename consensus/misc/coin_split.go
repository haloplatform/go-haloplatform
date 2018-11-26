package misc

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

// ApplyCoinSplitHardFork modifies the state database according to the CoinSplit hard-fork
// rules, inflate all balances of a set of accounts
func ApplyCoinSplitHardFork(statedb *state.StateDB) {
	/// TODO: May change for new structure after Shannon start his task
	// Inflate by adding balance for all accounts
	for _, acc := range params.CoinSplitAccount() {
		tmpBig := big.NewInt(0)
		tmpBig.SetString(acc.AddBalance, 10)
		statedb.AddBalance(acc.Address, tmpBig) /// StateProcessor use the reference, need to new BigInt object
	}
}
