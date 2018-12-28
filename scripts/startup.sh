#!/bin/bash
chown -R halo /go-halo/qdata
while test $# -gt 0; do
    case "$1" in
        --testnet)
            flags=${*/--testnet/""}
            echo "[Testnet] Starting full node"
            runuser -l halo -c "/go-ethereum/scripts/full-node-testnet-start.sh ${flags}"
            break
            ;;
        --devnet)
            flags=${*/--devnet/""}
            echo "[Devnet] Starting full node"
            runuser -l halo -c "/go-ethereum/scripts/full-node-devnet-start.sh ${flags}"
            break
            ;;
        --mainnet)
            flags=${*/--mainnet/""}
            echo "[Mainnet] Starting full node"
            runuser -l halo -c "/go-ethereum/scripts/full-node-mainnet-start.sh ${flags}"
            break
            ;;
        *)
            flags=${*/--mainnet/""}
            runuser -l halo -c "/go-ethereum/scripts/full-node-mainnet-start.sh ${flags}"
            break
            ;;
    esac
done
