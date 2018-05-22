# :cake: Feather Client (aka Featherlite) Integration

The featherlite client is built to work with a running `Glo` light node. Either Featherlite can start this node for you, or you can have the node already started before Featherlite Starts.

If you are wanting to start a `Glo` node outside of featherlite for local development on say a :lock: private network, then you will need to pass the additional commands to your `Glo` node. 

```bash
--rpc --rpccorsdomain="*"
```

Featherlite will detect your `Glo` client running and use it for blockchain access. 

----
> Side note for pre-launch development. The client detection isn't working fully so you will need to change `HAPI` to look at the `localhost:8545` port.

---

## Getting Feather Client

The full feather client repo can be found here: https://github.com/haloplatform/feather-client

The `Develop` branch will hold the latest greatest version of feather. 

Once you look at that repo follow the instructions there to get it started. If splashUI isn't completed when you run it (it boots straight into onboarding) then the client is pre-launch development. See side note above.
