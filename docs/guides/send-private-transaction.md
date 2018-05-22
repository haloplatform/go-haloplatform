# :lock: Sending Private Transacitons :lock:

With Halo Platform you can easily send private transactions because of our :star:**constellation**:star: of masternodes.

In the geth console you can do this via:

```bash
> eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[1], value: 100, privateFrom: btoa(eth.accounts[0]), privateFor: [btoa(eth.accounts[1]])});
```

In Web3 it is:

```javascript
web3.eth.sendTransaction({
    from: 'address string',
    to: 'address string',
    value: '100.01',
    gas: '0.0001',
    privateFrom: 'base64 encoded address string'
    privateFor: ['base64 encoded address string array']
})
```

_New transactionObject parameters:_

`privateFor` - The `Array<string>` of addresses to send the transaction to.
`privateFrom` - The address the transaction is sent from.

Only the node that is supposed to receive the transaction will be able to see the contents of it. It will be garbled junk for others.


> For More Information please refer to [the private send api](../api/private-send.md)
