FROM debian:buster-slim as downloader

LABEL maintainer="vincent_serpoul@manulife.com"

RUN mkdir inst
WORKDIR /inst

RUN apt-get update &&\
    apt-get install -y ca-certificates

COPY ./install_master_linux-amd64.tar.gz /inst/
RUN tar -xf /inst/install_master_linux-amd64.tar.gz -C /inst

RUN /inst/update.sh -i -c stable -p /app -d /var/nodedata -n

COPY ./genesis.json /app/randpool.genesis.json

RUN mkdir -p /var/kmddata && \
    mkdir -p /var/nodedata

# Create randpool network
RUN /app/goal network create -r /app/nodes/randpool -n private -t /app/randpool.genesis.json

# desactivate telemetry
RUN /app/diagcfg telemetry disable

FROM debian:buster-slim

COPY --from=downloader /app /app

# Copy the right config for the node
COPY ./node_config.json /app/nodes/randpool/Primary/config.json

# Copy the right config for kmd
COPY ./kmd_config.json /app/nodes/randpool/Primary/kmd-v0.5/kmd_config.json

RUN groupadd -r algo && useradd -r -g algo algo
RUN chown -R algo:algo /app && \
    chown -R algo:algo /app/nodes/randpool/Primary

USER algo

# import wallet1 so it's avail
RUN /app/goal account importrootkey -u -d /app/nodes/randpool/Primary

VOLUME ["/app/nodes/randpool/Primary"]

COPY ./start.sh /app/start.sh

EXPOSE 7979 7833


ENTRYPOINT ["/app/start.sh"]

# docker build -t randpool/node:v1.0.0 infra/node