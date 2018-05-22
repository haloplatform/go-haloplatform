# Private Send

`web3.eth.sendTransaction(object)` was modified to support private transactions.

You can send a transaction like normal to the network using 

```javascript
web3.eth.sendTransaction(transactionObject [, callback])
```

## `transactionObject`

- `from`:`string` - The address for the sending account. Uses the [web3.eth.defaultAccount](web-eth.md#DefaultAccount) property, if not specified.
- `to`:`string` - (optional) The destination address of the message, left undefined for a contract-creation transaction.
- `value`:`number|string|BigNumber` - (optional) The value transferred for the transaction in Halo Coin, also the endowment if it's a contract-creation transaction.
- `gas`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The amount of gas to use for the transaction (unused gas is refunded).
- `gasPrice`:`Number|String|BigNumber` - (optional, default: To-Be-Determined) The price of gas for this transaction in Halo Coin, defaults to the mean network gas price.
- `data`:`String` - (optional) Either a byte string containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code.
- `nonce`:`Number` - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
- `privateFrom`:`String` - (optional) When sending a private transaction, the sending party's base64-encoded public key to use. If not present and passing privateFor, use the default key as configured in the TransactionManager.
- `privateFor`:`Array<String>|List<String>` - (optional) When sending a private transaction, an array of the recipients' base64-encoded public keys.

## `callback`

- `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. It is recommended to always call asynchronous.

## Returns

- `txHash`:`string` - The 32 Bytes transaction hash as HEX string. 

> If the transaction was a contract creation use web3.eth.getTransactionReceipt() to get the contract address, after the transaction was mined.

## Example

Send Contract

```javascript
// compiled solidity source code using https://chriseth.github.io/cpp-ethereum/
var code = "603d80600c6000396000f3007c01000000000000000000000000000000000000000000000000000000006000350463c6888fa18114602d57005b6007600435028060005260206000f3";

web3.eth.sendTransaction({    
    data: code,
    privateFor: ['ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc=']
  },
  (err, address) => {
    if (!err) {
      console.log(address); // '0x7f9fade1c0d57a7af66ab4ead7c2eb7b11a91385'
    }
  }
});
```

Send Transaction

```javascript
// address' used are not valid network address. They are used for examples only.

web3.eth.sendTransaction({  
    from: '0x01',  
    amount: '0.01',    
    gas: '0.00001',
    privateFrom: 'ROAZ=',
    privateFor: ['ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc=']
  },
  function(err, address) {
    if (!err) {
      console.log(address); // '0x7f9fade1c0d57a7af66ab4ead7c2eb7b11a91385'
    }
  }
});
