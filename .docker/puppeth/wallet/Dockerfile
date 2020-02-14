FROM silesiacoin/infra-client:latest as client

# Inject libraries into
FROM puppeth/wallet:latest

COPY --from=client /usr/local/bin/ssc /usr/local/bin/ssc
RUN rm -rf /usr/local/bin/geth; ln -s /usr/local/bin/ssc /usr/local/bin/geth

EXPOSE 80 8545