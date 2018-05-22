# Start Private Network

This will start a preconfigured private network with several debugging flags. It is intended for local development against the blockchain such as web3 testing and contract testing.

## :one: - Clone Repo

First step as always:

```bash
git clone https://github.com/haloplatform/glo.git

cd glo

git checkout testnet_qa
```

## :two: - Install or Go To Binary Directory

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

## :three: - Start network

**Linux**

```bash
geth --dev --ipcpath ~/.ethereum/geth.ipc console
```

**MacOS**

```bash
geth --dev --ipcpath ~/Library/Ethereum/geth.ipc console
```

**Windows**

```bash
geth.exe --dev --ipcpath geth.ipc console
```

## :four: - Start Miner

In the console run: `miner.start()`

You can use `miner.stop()` to stop the miner.

The mining will cause blocks to be generated for you to actually perform transactions.

## Commands

Review the [quick commands](quick-commands.md) for instructions to get started in geth.

---
This should get you started with a private network for deploying and testing smart contracts and interacting with the chain via web3.

To turn on web3 interaction remember to set your rpccorsdomain flags and other items.

