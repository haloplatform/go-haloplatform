# :computer: Quick Commands

Just a few commands to get you quickly started in `Glo`

## Attach to `Glo` Console

Run this command in your terminal.

```shell
geth.attach
```

---

## Firewall

First make sure the following ports are open on your firewall

- `50608` - Used for peer to peer connections
- `50609` - Masternode light server peer to peer connections
- `9565` - Used for RPC
- `50401` - Used for raft consensus

----

## Commands

### List Accounts

```bash
> personal
```

or

```bash
> eth.accounts
```

### Create Account

```bash
> personal.newAccount("passphrase")
```

> :warning: :warning: Remember to write down and keep your passphrase in a secure location. :warning: :warning:

Console should respond with your address like so: 

```bash
Address: {168bc315a2ee09042d83d7c5811b533620531f67}
```

### Unlock Wallet

```bash
> personal.unlockAccount("address string")
```

### Number of Synced Blocks

```bash
> eth.blockNumber
```

### Number of connected peers

```bash
> net.peerCount
```

### Check Account Balance

```bash
> eth.getBalance(eth.accounts[0])
```

or

```bash
> eth.getBalance("address string")
```

### Send Transaction

```bash
> eth.sendTransaction({from: "address string", to: "address string", value: "100.01" });
```

or

```bash
> eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[1], value: 100.01});
```

### Exit Console

```bash
> Exit
```