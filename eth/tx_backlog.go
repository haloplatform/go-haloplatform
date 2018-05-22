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

package eth

import (
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type TxBacklog struct {
	mu            sync.Mutex
	txRemoteNum   int                     /// The current number of remote txns
	txRemoteQueue chan *types.Transaction /// The buffered channel like as
	maxRemoteBuff int
}

func NewTxBackLog(maxBuff int) *TxBacklog {
	// Create the protocol manager with the base fields
	txBacklog := &TxBacklog{
		maxRemoteBuff: maxBuff,
		txRemoteQueue: make(chan *types.Transaction, maxBuff),
	}
	log.Info("Create transaction backlog", "remote", txBacklog.maxRemoteBuff)
	return txBacklog
}

func (tb *TxBacklog) QueueRemoteTxs(txs []*types.Transaction) {
	tb.mu.Lock()

	if tb.txRemoteNum == tb.maxRemoteBuff {
		/// queue is full, drop them
		tb.mu.Unlock()
		return
	}

	addNum := len(txs)
	if tb.txRemoteNum+addNum > tb.maxRemoteBuff {
		/// Fill to full queue
		addNum = tb.maxRemoteBuff - tb.txRemoteNum
	}
	tb.txRemoteNum += addNum
	tb.mu.Unlock()

	/// Write into buffered channel
	for i := 0; i < addNum; i++ {
		tb.txRemoteQueue <- txs[i]
	}
}

func (tb *TxBacklog) DeQueueRemoteTxs(num int) []*types.Transaction {
	tb.mu.Lock()

	/// Empty
	if tb.txRemoteNum == 0 {
		tb.mu.Unlock()
		return nil
	}

	var getRemote int = 0
	if num > tb.txRemoteNum {
		getRemote = tb.txRemoteNum
	} else {
		getRemote = num
	}
	tb.txRemoteNum -= getRemote
	tb.mu.Unlock()

	result := make([]*types.Transaction, getRemote)
	for i := 0; i < getRemote; i++ {
		result[i] = <-tb.txRemoteQueue
	}
	return result
}

func (tb *TxBacklog) GetStatus() int {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	return tb.txRemoteNum
}
