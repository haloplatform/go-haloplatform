#!/bin/bash
set -u
set -e

GLOBAL_ARGS="--rpc --rpcaddr 0.0.0.0 --rpcport 9565 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --ws --wsport 9555 --wsaddr 0.0.0.0 --wsorigins" * --wsapi eth,web3"

echo "[*] Starting full node"
geth --syncmode full --networkid 42568378638 --datadir /root/.go-haloplatform/qdata $GLOBAL_ARGS --port 50607 --ipcpath /root/.go-haloplatform/qdata/geth.ipc