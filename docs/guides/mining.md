# Mining

Just some general tips on mining. You will need a node running so see [running a node](running-a-node.md) or [starting a private network](start-private-network.md) for information to do that.

Currently `Glo` only includes the original `geth` cpu miner for mining. Other miners will be modified by Halo Platform or the community post launch. There are just a few minor adjustments that have to be made to have miners accept non Proof of Work blocks.

## Start Mining

You can simply start mining when attached to the `Glo` console with the following command.

```bash
> miner.start(4)
```

The number passed to the `start` function indicates the number of cpu threads you wish to mine with. More mines faster, but stresses your CPU more. Find a balance.

You can then stop mining by simply running this command in the `Glo` console.

```
> miner.stop()
```

## Getting Rewards

Before you start mining you should set where you are going to receive those rewards. You do that by setting the etherbase where it will go. Do so like this in the `Glo` console.

```bash
> miner.setEtherbase(eth.accounts[1])
```

or
```bash
> miner.setEtherbase("address string")
```

> Note that your etherbase does not need to be an address of a local account, just an existing one.

## Extra Data

There is an option to add extra Data (32 bytes only) to your mined blocks. By convention this is interpreted as a unicode string, so you can set your short vanity tag. 

> Note: 32 bytes roughly equals 32 characters in ASCII


```bash
> miner.setExtra("shadowsdocs")
```

## Hashrate

You can check your hashrate with miner.hashrate, the result is in H/s (Hash operations per second).

```bash
> miner.hashrate
```

would result in say `712000` which would equal `695.3 KH/s` or `0.67 MH/s`

## Checking rewards and using them

After you successfully mined some blocks, you can check the ether balance of your etherbase account. Now assuming your etherbase is a local account:

```bash
> eth.getBalance(eth.coinbase).toNumber();
```

In order to spend your earnings you will need to have this account unlocked.

```bash
> personal.unlockAccount(eth.coinbase)
```

Then you can send a transaction like normal.

```bash
> eth.sendTransaction({to: 'address string', from: eth.coinbase, amount 100})
```

## Success

Mining success depends on the set block difficulty. Block difficulty dynamically adjusts each block in order to regulate the network hashing power to produce a 4 minute blocktime. Your chances of finding a block therefore follows from your hashrate relative to difficulty. 