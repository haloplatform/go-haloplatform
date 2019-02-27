// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Halo network.
var MainnetBootnodes = []string{
	// Glo bootnode mainnet
	"enode://9cd1270d06552d2053065418aabd39a19c5a27881d49efe2326bc4d6c17b95756cd78ed84b922a1cac51e2ff56f7efdecc7ef839c73478198e2d606f5a6e7442@35.185.178.33:50607",
	"enode://02970aeff18d42d12a27376426b65b9e32c12651bb8fbf962506ee6849509d7fd8a47dda3f2bf773946a21ee1ba91d48fe1bee95a9812c632223b295344bb54a@35.198.4.148:50607",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	// Glo bootnode testnet
	"enode://b959f9d3cd422abc738dd4f015c970f341411b79c8488ef5f68792dcfc722d6c215d5dd5d2581d319e3d41120e30eecf51f347cdadc94dcf3ef03cd4f81413c5@35.197.125.20:50607",
}
