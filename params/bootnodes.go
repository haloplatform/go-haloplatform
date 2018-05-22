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
	"enode://a4370e6b836af96ecabbf0e3f5f272b9f282f016f3c2fc5812f67ba2f05d14918fea2a1f6b18ff7b4b23f8fe9e20e90eee8e7315e6d6200721d9ccce32364794@18.207.35.156:50607",
	"enode://50acafdafe5807cd9198eb40f550c91fa596d19f787210fd34a93fe8a64c019f557a8d8ff93ef46b1c90eec8e7419ad4f5af266197c3f40506d0f925984c5cc5@54.183.50.199:50607",
	"enode://7af6741161924313ad3030721c671d502f1bdc9e8021b8299e75b80a99072b2fd8a06e63218e9b2eddc517fe3b1c8d0b0249771477b06abdb1dab922f493fde4@13.236.73.230:50607",
	"enode://4a7a2e319b0cabe6225f0d785348aa8d7587b5ae2599c9bb4b275607f3c3900b32c053928cb0a4c279d8cb963d6e03c8634ddddedcdce2af84a2f9ab44954606@13.228.189.214:50607",
	"enode://00dbec11de2d5dda6d6ef9e12ed4d1a9fda2a24ab19588993a4fc33bdeaa4b727c86d9fa71e18917ae13faa17691da0cb867204bb5d0ba4b20f3a4e7223500e8@52.67.72.214:50607",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	// Glo bootnode testnet
	"enode://95d87fedf849fb04594fe3783cde4aa3421ba66442f4cfd358e85ba899292c0a04edf8e5361553486d89155c029adf0426dc9609e4fe677e7fa0a7d629fec350@54.191.87.237:50607",
}
