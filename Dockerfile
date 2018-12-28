FROM golang:1.9

ENV user halo
RUN mkdir /go-ethereum
RUN useradd -m -d /go-halo ${user} 

ADD . /go-ethereum
WORKDIR /go-ethereum

RUN chown -R ${user} /go-ethereum

USER ${user}
RUN make geth
USER root
RUN mv /go-ethereum/build/bin/geth /usr/local/bin/
RUN chmod +x /go-ethereum/scripts/startup.sh \
    /go-ethereum/scripts/full-node-testnet-start.sh \
    /go-ethereum/scripts/full-node-mainnet-start.sh  \
    /go-ethereum/scripts/full-node-devnet-start.sh
EXPOSE 9565 9566 50607
RUN chown -R ${user} /go-halo
ENTRYPOINT ["/go-ethereum/scripts/startup.sh"]
