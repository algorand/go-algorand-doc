# Relay Provisioning
## Basic Process based on real world example

* Create instance, open ports, configure static IP (or static DNS name)
* Configure Environment (install libsodium/dependencies)
* Install ALGORAND
* Enable telemetry, metrics
* Update config.json
* Start algod
* Install as service: systemd-setup.sh
* Automate Updates: crontab -e

## Create Instance
Create instance (e.g. at least EC2 i3.large); we use Ubuntu 18.04 LTS  
Open ports: SSH (for your own access), 4161, 9100  
Start Instance  

## Configure Environment
Provision storage if necessary so algod can be installed on large, fast storage device  
Install dependencies

    sudo apt update
    sudo apt install -y libsodium-dev ca-certificates wget --no-install-recommends

## Install Algorand
Grab the installation bootstrapper [from here](https://github.com/algorand/go-algorand-doc/tree/master/downloads/installers/linux_amd64)  
Copy it to the new machine.

    # From local machine
    scp /path/to/installer/install_master_linux-amd64.tar.gz ubuntu@<instance-address>:/tmp

Connect to instance again and install algod.  Note we follow a naming convention of ~/algorand/<network>/<nodetype> for our installations; you can follow whatever convention works for you, but make sure to do the appropriate substitutions.  In this case, <network> is testnet and <nodetype> is relay.

    cd ~/
    tar -xf /tmp/install_master_linux-amd64.tar.gz
    ./update.sh -i -c stable -p ~/algorand/testnet -d ~/algorand/testnet/data/relay -n
    cp ~/algorand/testnet/data/relay/config.json.example ~/algorand/testnet/data/relay/config.json
    # Edit config.json and update NetAddress to ":4161"

## Setup DNS and Telemetry
> `diagcfg telemetry` is similar to `goal logging enable`

    cd ~/algorand/testnet
    ./diagcfg telemetry name -n <DNS-Name-For-This-Machine>
    # e.g. ./diagcfg telemetry name -n <r1.algorand.network>

    # <DNSNAME> should be the full name like r1.network.algorand
    ./diagcfg metric enable -e <DNSNAME> -d data/relay

## Register algod with Systemd, start service
    # Edit ~/algorand/testnet/algorand@.service.template, if necessary
    # Specify the correct user account in these commands – ubuntu is default on ubuntu instances (ec2user on AWS linux)
    sudo ./systemd-setup.sh ubuntu ubuntu
    sudo systemctl enable algorand@$(systemd-escape /home/ubuntu/algorand/testnet/data/relay)
    sudo systemctl start algorand@$(systemd-escape /home/ubuntu/algorand/testnet/data/relay)

## Setup cronjob
    crontab -e
    # Check at every 30 minute mark:
    30 * * * * /home/ubuntu/algorand/testnet/update.sh -d /home/ubuntu/algorand/testnet/data/relay >/home/ubuntu/algorand/testnet/update.log 2>&1

## Send Algorand the DNS Name and Port for your relay (e.g. r1.algorand.network:4161)
