# Relay Provisioning
## Basic Process based on real world example

* [Create Instance, open ports, configure static IP (or static DNS name)](#create-instance)
* [Configure Environment (install libsodium/dependencies)](#configure)
* [Install Algorand, update config.json](#install)
* [Enable Telemetry and Metrics](#telemetry)
* [Intall and run as a service](#service)
* [Configure for Automate Updates: crontab -e](#updates)
* [Run algod](#run)
* [Send DNS Name and Port to Algorand](#send)

## <a id="create-instance">Create Instance</a>
Create instance (e.g. at least EC2 i3.large); we use Ubuntu 18.04 LTS  
Open ports: SSH (for your own access), 4161, 9100  
Start Instance  

## <a id="configure">Configure Environment</a>
Provision storage if necessary so algod can be installed on large, fast storage device  
Install dependencies

    sudo apt update
    sudo apt install -y libsodium-dev ca-certificates wget --no-install-recommends

## <a id="install">Install Algorand</a>
Grab the installation bootstrapper [from here](https://github.com/algorand/go-algorand-doc/tree/master/downloads/installers/linux_amd64)  
Copy it to the new machine.

    # From local machine
    scp /path/to/installer/install_master_linux-amd64.tar.gz ubuntu@<instance-address>:/tmp

Connect to instance again and install algod.  Note we follow a naming convention of ~/algorand/\<network\>/\<nodetype\> for our installations; you can follow whatever convention works for you, but make sure to do the appropriate substitutions.  In this case, \<network\> is testnet and \<nodetype\> is relay.

    cd ~/
    tar -xf /tmp/install_master_linux-amd64.tar.gz
    ./update.sh -i -c stable -p ~/algorand/testnet -d ~/algorand/testnet/data/relay -n
    cp ~/algorand/testnet/data/relay/config.json.example ~/algorand/testnet/data/relay/config.json
    # Edit config.json and update NetAddress to ":4161"

## <a id="telemetry">Enable Telemetry and Metrics</a>
> 'diagcfg telemetry' is similar to 'goal logging enable'

    cd ~/algorand/testnet
    ./diagcfg telemetry name -n <DNS-Name-For-This-Machine>
    # e.g. ./diagcfg telemetry name -n r1.algorand.network

    # <DNSNAME> should be the full name like r1.network.algorand
    ./diagcfg metric enable -e <DNSNAME> -d data/relay

## <a id="service">Intall and run as a service</a>

    # Edit ~/algorand/testnet/algorand@.service.template, if necessary
    # Specify the correct user account in these commands – ubuntu is default on ubuntu instances (ec2user on AWS linux)
    sudo ./systemd-setup.sh ubuntu ubuntu
    sudo systemctl enable algorand@$(systemd-escape /home/ubuntu/algorand/testnet/data/relay)
    sudo systemctl start algorand@$(systemd-escape /home/ubuntu/algorand/testnet/data/relay)

## <a id="updates">Configure for Automate Updates: crontab -e</a>
    crontab -e
    # Check at every 30 minute mark:
    30 * * * * /home/ubuntu/algorand/testnet/update.sh -d /home/ubuntu/algorand/testnet/data/relay >/home/ubuntu/algorand/testnet/update.log 2>&1

## <a id="run">Run algod</a>
    goal node start -d ~/algorand/testnet/data/relay

## <a id="send">Send DNS Name and Port to Algorand</a>
Email the DNS Name and Port for your relay (e.g. r1.algorand.network:4161) to testnet-team@algorand.com  
We will evaluate your request and add your relay to the official list at our discretion.
