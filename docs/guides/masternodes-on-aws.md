# Starting a Masternode on AWS

Using three servers to run MN network, one for bootnode and two for MN's.

## :one: - Clone `Glo` Github Repo

```bash
git clone https://github.com/haloplatform/glo.git

cd glo

git checkout testnet_qa
```

## :two: - Install executables

**Masternode**

```
sudo cp build/bin/geth /usr/local/bin/
```

**Bootnode**

```
sudo cp build/bin/bootnode /usr/local/bin/
```

## :three: - Execute commands on appropriate servers according to their roles

**Bootnode**

```
nohup bootnode -addr [::]:50608 -nodekey btnode --verbosity 5 &
```

**Masternode 1**

_init_
```
rm -rf qdata
mkdir -p qdata/logs
mkdir -p qdata/dd1/{keystore,geth}
cp permissioned-nodes.json qdata/dd1/static-nodes.json
cp keys/key6 qdata/dd1/keystore
cp keys/key7 qdata/dd1/keystore
cp raft/nodekey1 qdata/dd1/geth/nodekey
geth --datadir qdata/dd1 init genesis.json
```

_run_
```
nohup geth --lightserv 50 --raft --datadir qdata/dd1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --raftport 50401 --rpcport 9565 --port 50608 --ipcpath qdata/dd1/geth.ipc --unlock 0 --password passwords.txt 2>>qdata/logs/1.log &

```

**Masternode 2**

_init_
```
rm -rf qdata
mkdir -p qdata/logs
mkdir -p qdata/dd2/{keystore,geth}
cp permissioned-nodes.json qdata/dd2/static-nodes.json
cp keys/key6 qdata/dd2/keystore
cp keys/key7 qdata/dd2/keystore
cp raft/nodekey2 qdata/dd2/geth/nodekey
geth --datadir qdata/dd2 init genesis.json
```

_run_
```
nohup geth --lightserv 50 --raft --datadir qdata/dd2 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --raftport 50401 --rpcport 9565 --port 50608 --ipcpath qdata/dd2/geth.ipc --unlock 0 --password passwords.txt 2>>qdata/logs/2.log &
```

## :four: - Perform first transaction on any MN server

```bash
geth attach qdata/dd{1,2}/geth.ipc
> eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[1], value: 100});
```

_make sure balance has changed_

```bash
> eth.getBalance(eth.accounts[0])
```

_Get list of connected peers_

```bash
> admin.peers
```

