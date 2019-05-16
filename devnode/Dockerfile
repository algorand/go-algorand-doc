FROM ubuntu:19.04

LABEL maintainer="vincent_serpoul@manulife.com"

RUN mkdir inst
WORKDIR /inst

COPY ./install_master_linux-amd64.tar.gz /inst/

RUN tar -xf install_master_linux-amd64.tar.gz

RUN apt-get update &&\
    apt-get install -y ca-certificates

RUN mkdir -p /root/node

RUN ./update.sh -i -c stable -p ~/node -d ~/node/data -n

######

COPY ./genesis.json /root/randpool.genesis.json

# Create randpool network
RUN /root/node/goal network create -r /root/randpool -n private -t /root/randpool.genesis.json

# desactivate telemetry
RUN /root/node/diagcfg telemetry disable

# Copy the right config
# COPY ./nodeconfig.json /root/randpool/Node/config.json
COPY ./nodeconfig.json /root/randpool/Primary/config.json

VOLUME ["root/randpool"]

COPY ./start.sh /root/start.sh

EXPOSE 7979

WORKDIR /root

ENTRYPOINT ["/root/start.sh"]

# docker build -t randpool/devnode .