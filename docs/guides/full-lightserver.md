# Running a full node as a server for light nodes

Light nodes use the 'les' subprotocol for communication and need special endpoints as a servers. A full node with --lightserv option can be such a server for light nodes.


## :one: - Clone `Glo` Repo

```bash
git clone https://github.com/haloplatform/glo.git

cd glo

git checkout testnet_qa
```

## :two: - Install

To install the `glo` client you need to copy it to your operating systems command path.

**Linux** :penguin:
```bash
sudo cp pre_built/Ubuntu/geth /usr/local/bin/
```

**Windows** 

Can't easily install in windows, instead just simply `cd` to the correct folder and call the `.exe` directly.

```bash
cd pre_built\Windows\
```

**MacOS**
```bash
sudo cp pre_built/Mac/geth /usr/local/bin/
```

> Note for MacOS users: If your terminal tells you that `/usr/local/bin` doesn't exist run `sudo mkdir /usr/local/bin`

## :three: - Run Node

```
cd local_test_net
```

**Linux/MacOS**

_init_

```
rm -rf qdata
mkdir -p qdata/logs
mkdir -p qdata/dd1/{keystore,geth}
cp keys/key6 qdata/dd1/keystore
cp keys/key7 qdata/dd1/keystore
geth --datadir qdata/dd1 init genesis.json
```

_run_

```
nohup geth --lightserv 50 --datadir qdata/dd1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --rpcport 9565 --port 50608 --ipcpath qdata/dd1/geth.ipc --unlock "0,1" --password passwords.txt 2>>qdata/logs/1.log &
```

_attach_

```
geth attach qdata/dd1/geth.ipc
```

**Windows**

_init_

```
RMDIR /S /Q qdata
MKDIR qdata\logs
mkdir qdata\dd1\keystore
mkdir qdata\dd1\geth
copy keys\key6 qdata\dd1\keystore
copy keys\key7 qdata\dd1\keystore
geth --datadir qdata\dd1 init genesis.json
```

_run_

```
start cmd /K geth --lightserv 50 --datadir qdata\dd1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --rpcport 9565 --port 50608 --ipcpath "node1" --unlock "0,1" --password passwords.txt
```

_attach_

```
geth attach \\.\pipe\node1
```