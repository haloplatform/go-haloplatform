# Running A Light Node connected to Testnet

This will get you up and going with running a light node locally and connecting it to a network of masternodes running the testnet.

---

## Using Prebuilt `Glo`

### :one: - Clone `Glo` Github Repo

```bash
git clone https://github.com/haloplatform/glo.git

cd glo

git checkout testnet_qa
```

### :two: - Install the `Glo` client code

To install the `glo` client you need to copy it to your operating systems command path.

**Linux** :penguin:
```bash
sudo cp pre_built/Ubuntu/geth /usr/local/bin/
```

**Windows** 
```bash
copy pre_built\Windows\geth.exe local_test_net\light_node
```

**MacOS**
```bash
sudo cp pre_built/Mac/geth /usr/local/bin/
```

> Note for MacOS users: If your terminal tells you that `/usr/local/bin` doesn't exist run `sudo mkdir /usr/local/bin`

After copying the `Glo` executable you need to:

```bash
cd local_test_net/light_node
```

### :three: - Init your genesis block - _Do Only Once_

**Linux/MacOS**
```bash
rm -rf node_data/geth

geth --datadir node_data init genesis.json
```

**Windows**
```bash
RMDIR /S /Q node_data\geth

geth --datadir node_data init genesis.json
```

### :four: - Start the light node

**Linux/MacOS**

Run in Terminal: 
```bash
geth --datadir node_data --syncmode "light" --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --rpcport 9565 --port 50608 --unlock "0,1" --password passwords.txt --verbosity 5 2>>node.log
```

Open another Terminal and run:
```bash
geth attach node_data/geth.ipc
```

**Windows**

Run in Command Line: 
```bash
geth --datadir node_data --syncmode "light" --ipcpath "light" --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints --rpcport 9565 --port 50608 --unlock "0,1" --password passwords.txt --verbosity 5 >> node.log
```

Open another Command Line and run:

```bash
geth attach "\\.\pipe\light"
```

---

## Using Source

Using `Glo` from source is basically the same. Instead of running the prebuilt binaries you will need to follow the following guide to compile your source code.

[Compile Glo from source](compile-glo-from-source.md)

After you do that move the resulting binary to the same locations as talked about above and follow steps 2-4 to get connected.