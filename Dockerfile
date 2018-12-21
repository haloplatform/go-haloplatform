FROM golang:1.9
ADD . /go-ethereum
WORKDIR /go-ethereum
RUN make geth
RUN mv /go-ethereum/build/bin/geth /usr/local/bin/
RUN chmod +x /go-ethereum/scripts/startup.sh \
    /go-ethereum/scripts/full-node-testnet-start.sh \
    /go-ethereum/scripts/full-node-mainnet-start.sh  \
    /go-ethereum/scripts/full-node-devnet-start.sh
EXPOSE 9565 9566 50607
ENTRYPOINT ["/go-ethereum/scripts/startup.sh"]