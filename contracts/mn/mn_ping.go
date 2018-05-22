// Copyright 2016 The go-ethereum Authors
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

package mn

//go:generate abigen --sol contract/MNReward.sol --pkg contract --out contract/mnreward.go

import (
	"context"
	"crypto/ecdsa"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	mnsc "github.com/ethereum/go-ethereum/contracts/mn/contract"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type Config struct {
	InstAddr     common.Address // MNInstance address
	RewardAddr   common.Address // MNReward address
	PingInterVal int
	PingKey      *ecdsa.PrivateKey
}

type MNPing struct {
	config    *Config
	txOpt     *bind.TransactOpts
	ethClient *ethclient.Client
	quitSync  chan struct{}
	lastTx    *types.Transaction
	pingIP    *net.IP
}

func NewMNPing(rpc *rpc.Client, conf *Config, exIP *net.IP) (self *MNPing, err error) {

	// rpcClient, err := stack.Attach()
	// if err != nil {
	// 	log.Crit("Failed to attach to self: %v", err)
	// }

	mn_ping := &MNPing{
		config:    conf,
		txOpt:     bind.NewKeyedTransactor(conf.PingKey),
		ethClient: ethclient.NewClient(rpc),
		quitSync:  make(chan struct{}),
		pingIP:    exIP,
	}
	/// There is a problem in gas estimation (EstimateGas()) -> fixed value will work
	mn_ping.txOpt.GasLimit = 1000000

	return mn_ping, nil
}

func (self *MNPing) Start() {
	log.Info("Start masternode PING loop", "inst", self.config.InstAddr, "reward", self.config.RewardAddr, "interval(m)", self.config.PingInterVal, "ip", self.pingIP.String())
	self.pingLoop()
}

func (self *MNPing) Stop() {
	close(self.quitSync)
	log.Info("Masternode PING stopped")
}

func (self *MNPing) doPing() {
	log.Debug("Do Ping", "inst_addr", self.config.InstAddr)

	/// Check the result of previous PING
	if self.lastTx != nil {
		receipt, err := self.ethClient.TransactionReceipt(context.Background(), self.lastTx.Hash())
		if err != nil {
			log.Warn("Get PING receipt error", "error", err)
		}
		/// We can't use receipt.Status because it's always failed
		if receipt.CumulativeGasUsed == self.txOpt.GasLimit {
			log.Warn("The previous PING is failed", "tx", self.lastTx.Hash().Hex())
		} else {
			log.Info("The previous PING is successful", "tx", self.lastTx.Hash().Hex())
		}
	}

	/// Check
	mn_inst, err := mnsc.NewMNInstance(self.config.InstAddr, self.ethClient)
	if err != nil {
		log.Crit("Failed to retrieve SC", "err", err)
	}

	inst_session := &mnsc.MNInstanceSession{
		Contract:     mn_inst,
		TransactOpts: *self.txOpt,
	}

	/// Try to get to check whether SC is existing or not
	node_info, err := inst_session.GetInfo()
	if err != nil {
		/// Need to wait for synchronization
		log.Warn("Failed to get MNInstance info", "err", err)
		return
	}

	log.Info("MNInstance info", "info", node_info)

	mn_reward, err := mnsc.NewMNReward(self.config.RewardAddr, self.ethClient)
	if err != nil {
		log.Crit("Failed to retrieve SC", "err", err)
	}

	reward_session := &mnsc.MNRewardSession{
		Contract:     mn_reward,
		TransactOpts: *self.txOpt,
	}

	/// UT: lastTx, err := reward_session.DoPing(self.config.InstAddr, "192.168.0.1")
	lastTx, err := reward_session.DoPing(self.config.InstAddr, self.pingIP.String())
	if err != nil {
		/// Need to wait for synchronization
		log.Warn("Failed to do ping", "err", err)
		return
	}

	log.Info("Ping txn", "hash", lastTx.Hash().Hex())
	self.lastTx = lastTx
}

func (self *MNPing) pingLoop() {
	pingTicker := time.NewTicker(time.Duration(self.config.PingInterVal) * time.Minute)
	defer pingTicker.Stop()

	// PING when starting
	self.doPing()

	for {
		select {
		case <-pingTicker.C:
			self.doPing()
		// quit
		case <-self.quitSync:
			log.Debug("Masternode PING loop stopped")
			return
		}
	}
}
