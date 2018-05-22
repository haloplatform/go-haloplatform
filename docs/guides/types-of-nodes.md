# Types of Nodes

There are several types of nodes that can be ran within the Halo Platform network. 

## Light Node

This is the way that most users should run their node. To run your node as a light node you should add the following command line argument to the full command.

```
--syncmode "light"
```

This will signal the `Glo` node to only sync the last 512 blocks and connect to the Masternode Network to operate and serve operations against the network.

## Full Node

This is the default mode of the `Glo` node. Simply following the instructions in [running a node](running-a-node.md)

## Full Node - Light Server

A full node can operate as a server for light clients. To initiate this simply add the following command line argument to the full command. :warning: Cannot be used in conjunction with light sync :warning:

```
--lightserv 50
```

This will open up the `les` protocol for light nodes to connect to.

## Masternode - Raft Server

Masternodes operate as both a Full node and a Light Server in addition to using a Proof of Stake raft fault tolerance consensus between Masternodes. The `--raft` command line argument starts it in raft node but we need a few more command line arguments to get it fully started. :warning: Cannot be used in conjunction with light sync :warning:

The extra command line arguments are:

```
--raft --lightserv 50 --raftport 50401
```

> :bulb: Masternodes and Full Nodes work together to maintain network stability.

## Mining Node

While it is recommended for learning purposes to use a full node and the `Glo` console to manage mining you can start a full node as a dedicated mining node with the following command line arguments.

```
--mine --minerthreads=1 --etherbase 'address string'
```

`minerthreads` denotes the number of CPU threads you want to use with your miner. 

`etherbase` is the address you wish to deposite your mining rewards into.

For full docs on mining please refer to [the mining guide](mining.md)