# Setting up a TestNet Node or Relay

1. TestNet nodes and the update process
2. Installing on a Mac
3. Installing on Ubuntu (we require 18.04)
4. Installing in Docker
5. Installing in a VM (see installing on Ubuntu)
6. Configuring your node as a Relay
7. Configuring auto-update
8. Running algod as a service
9. Notes on installing on other Linux Distros

## 1. TestNet nodes and the update process
The current release of our testnet package is designed around simple installation and a simplified auto-update mechanism.  A node installation consists of 2 folders: the Binaries (bin) and the Data (data) folders.  The bin folder can be anywhere you choose.  A recommended location is `~/node`.  We currently assume the folder is dedicated to algorand binaries, as we archive the folder before each update (we do not currently delete anything, but will overwrite our own binaries and add new ones).  The data folder is assumed to be `~/.algorand-testnet`, but can and should be overridden.  We will create it if it doesn’t exist.  We recommend using something under your node folder, e.g. `~/node/data`.
The bin folder holds our executables (`goal`, `algod`, `carpenter`, `doberman`, and `updater`), our update script (`update.sh`), and our optional configuration files (`config.json` and `phonebook.json`) -- we only distribute a sample `config.json.example` now.
The data folder holds `genesis.json`, used for creating the initial ledger (`ledger.sqlite`) in a network-specific sub-folder, generated log files (`node.log`, and a `latest.log` symlinked to the latest log file).  You can optionally create a `phonebook.json` and/or `config.json` file and add configuration options that are merged into the defaults read from the bin folder.

#### Installing a new node is generally a 3-step process:
1. Prepare your environment by installing dependencies
2. Retrieve the installation package appropriate for your platform
3. Optionally, verify the hash of `the tar.gz` file
4. Run `update.sh -i -c stable -p ~/node -d ~/node/data -n`

When the installer runs, it will pull down the latest update package from S3 for your platform and install it.
The `'-n'` option above tells the installed to not auto-start the node.  If installation succeeds you'll be instructed to start the node manually.
Before starting for the first time, you need to enable telemetry and configure a Host name for your machine (to help identify sources of telemetry).

    cd ~/node
    ./goal logging enable -n MeaningfulHostName

Please run `./goal logging` and send the output to Algorand so we can correlate you with your telemetry.

Then you can start the node!

    ./goal node start -d data

To check for, and install, the latest update, you can simply run `./update.sh -d ~/node/data` at any time.  It will query S3 for available builds and see if any are newer than your installed version (to force an update, you can delete algod and run `./update.sh -i -c stable -d ~/node/data` again).  If there is a newer version, it’s downloaded and unpacked before we shut down your node, back up your files, and install the update.  If any part of the process fails, we attempt to restore your previous version (bin and data) and restart the node.  If it succeeds, we’ll start the new version of the node (even if it wasn’t running when you initiated the update).

## 2. Installing on a Mac
To install on a Mac (OSX v10.13.4 / High Sierra):
* Open a terminal window
* Install the libsodium package
  * Follow these instructions if you do not have homebrew installed: http://macappstore.org/libsodium/
  * Otherwise `brew install libsodium`
* Install the libgmp package
  * `brew install gmp`
* Create a temporary folder to hold the install package and files
  * `mkdir ~/inst`
  * `cd ~/inst`
  * copy the installer (install_master_darwin-amd64.tar.gz) from github: downloads/installers/darwin_amd64
    * \(optional\) Verify SHA for the file
  * Unzip the package (use the appropriate filename)
    * `tar -xf install_master_darwin-amd64.tar.gz`
  * `./update.sh -i -c stable -p ~/node -d ~/node/data -n`

When the installer runs, it will pull down the latest update package from S3 for your platform and install it.
The `'-n'` option above tells the installed to not auto-start the node.  If installation succeeds you'll be instructed to start the node manually.
Before starting for the first time, you need to enable telemetry and configure a Host name for your machine (to help identify sources of telemetry).

    cd ~/node
    ./goal logging enable -n MeaningfulHostName

Please run `./goal logging` and send the output to Algorand so we can correlate you with your telemetry.

Then you can start the node!

    ./goal node start -d data

You should now have a running algorand node!  You can verify the daemon is running:

    pgrep algod

If it outputs a number (process id) then algod is running.
You can watch the agreement activity by running the carpenter utility (it’s a long-running program, so ctrl+c to stop watching):

    ./carpenter -file data/latest.log

## 3. Installing on Ubuntu
Nodes have been verified on Ubuntu 18.04.  Other Debian-based distros should work as well (use apt-get install rather than apt install)

* Open a terminal window
* Install the libsodium-dev package
  * `sudo apt update`
  * `sudo apt install libsodium-dev`
* Create a temporary folder to hold the install package and files
  * `mkdir ~/inst`
  * `cd ~/inst`
  * copy the installer (install_master_linux-amd64.tar.gz) from github: downloads/installers/linux_amd64
    * \(optional\) Verify SHA for the file
  * Unzip the package (use the appropriate filename)
    * `tar -xf install_master_linux-amd64.tar.gz`
  * `./update.sh -i -c stable -p ~/node -d ~/node/data -n`
    * If you get an error trying to download the update:
      * `sudo apt install ca-certificates`
      * Try `./update.sh -i -c stable -p ~/node -d ~/node/data` again

When the installer runs, it will pull down the latest update package from S3 for your platform and install it.
The `'-n'` option above tells the installed to not auto-start the node.  If installation succeeds you'll be instructed to start the node manually.
Before starting for the first time, you need to enable telemetry and configure a Host name for your machine (to help identify sources of telemetry).

    cd ~/node
    ./goal logging enable -n MeaningfulHostName

Please run `./goal logging` and send the output to Algorand so we can correlate you with your telemetry.

Then you can start the node!

    ./goal node start -d data

You should now have a running algorand node!  You can verify the daemon is running:

    pgrep algod

If it outputs a number (process id) then algod is running.
You can watch the agreement activity by running the carpenter utility (it’s a long-running program, so ctrl+c to stop watching):

    ./carpenter -file data/latest.log

## 4. Installing in Docker
A Docker image was created with Ubuntu 18.04 and the required libsodium-dev and wget packages pre-installed.  We also install a certificates package as the certs required for S3 are outdated.
* Open a terminal window
* Create a Docker container
  * `docker run -it algorand/node`
* This should put you at the bash prompt inside the Docker container (the remaining steps are identical to Ubuntu above, starting after the ‘Install the libsodium-dev package’ step)

Note we do not currently configure the container for any persistent storage, so when you stop the container your node is wiped - there is no persistence!

## 5. Installing in a VM
See installing on Ubuntu

## 6. Configuring your node as a Relay
A benefit of our decentralized network implementation is that a relay is effectively the same as any other node.  The distinction currently is made by configuring a node to actively listen for connections from other nodes, and having itself advertised using SRV records available through DNS. As a precaution is not recommended that relay nodes interact with accounts or participate in consensus. The primary reason for this recommendation is because the relay address is publically known and more likely to be subject to malicious attacks. If you want to run a relay and a normal node, just install the normal node on a different machine without a publicly advertised ip address.

Please contact Algorand if you are interested in running a Relay that will be able to support the network.  This requires a substantially more powerful machine than a typical Node.

## 7. Configuring auto-update
To ensure your node is running the latest software, we recommend using CRON to check for updates on a schedule.   Since checking for updates is fairly lightweight and does not shut down your node unless an update is available, there is no harm in checking frequently.
We recommend the following:

    crontab -e

Add a line that looks like this (run update.sh every hour of every day), where ‘user’ is the name of the account you’re using to install / run the node:

`30 * * * * /home/user/node/update.sh -d /home/user/node/data >/home/user/node/update.log 2>&1`

## 8. Running algod as a service
To ensure your node has maximum availability on the network, we have added the ability to run it as a service.

We are currently editing the configuration files for this and will update shortly.

## 9. Notes on installing on other Linux Distros
### AWS Linux (Red Hat 4.8.5-11)
 * sudo yum install https://rpmfind.net/linux/remi/enterprise/6/remi/x86_64/libsodium23-1.0.16-1.el6.remi.x86_64.rpm

### CentOS
 * sudo yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
 * sudo yum install libsodium

### Ensure your machine’s time is correct.  Ensure you have an NTP service running.
 
On CentOS:
 * `sudo yum install ntp`
 * `sudo chkconfig ntpd on`
 * `sudo ntpdate pool.ntp.org`
 * `sudo service ntpd start`
