#!/bin/bash
set -u
set -e

GLOBAL_ARGS="--rpc --rpcaddr 0.0.0.0 --rpcport 9565 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --ws --wsport 9566 --wsaddr 127.0.0.1 --wsorigins 127.0.0.1 --emitcheckpoints --wsapi eth,web3 --rpcvhosts rpctest.haloplatform.tech

echo "[*] Starting full node"
geth --syncmode full --networkid 4256338638 --datadir /go-halo/qdata $GLOBAL_ARGS --port 50607 --ipcpath /go-halo/qdata/geth.ipc