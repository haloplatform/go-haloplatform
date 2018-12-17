package params

import (
	"github.com/ethereum/go-ethereum/common"
)

type AccountInfo struct {
	Address    common.Address /// Account address
	AddBalance string         /// Convert to big.Int when AddBalance()
}

// Account list to inflate for Mainnet
func MainnetCoinSplitAccount() []AccountInfo {
	return []AccountInfo{
		{common.HexToAddress("0x67b7301878bc38559e118e9958d1bb8b4eb6afb2"), "199000000000000000000"}, /// 199 Halo
		{common.HexToAddress("0x928e5528aaa3c19640836c802c2743a0346acc56"), "299000000000000000000"}, /// 299 Halo
	}
}

// Account list to inflate for testnet
func TestnetCoinSplitAccount() []AccountInfo {
	return []AccountInfo{
		{common.HexToAddress("0x67b7301878bc38559e118e9958d1bb8b4eb6afb2"), "199000000000000000000"}, /// 199 Halo
		{common.HexToAddress("0x928e5528aaa3c19640836c802c2743a0346acc56"), "299000000000000000000"}, /// 299 Halo
	}
}

// Account list to inflate for local development
func LocalCoinSplitAccount() []AccountInfo {
	return []AccountInfo{
		{common.HexToAddress("0x67b7301878bc38559e118e9958d1bb8b4eb6afb2"), "499000000000000000000"}, /// 199 Halo
		{common.HexToAddress("0x928e5528aaa3c19640836c802c2743a0346acc56"), "599000000000000000000"}, /// 299 Halo
	}
}
