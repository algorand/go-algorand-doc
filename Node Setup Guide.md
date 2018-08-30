Setting up a Testnet Node or Relay

Testnet nodes and the update process
Installing on a Mac
Installing on Ubuntu (we require 18.04)
Installing in Docker
Installing in a VM (see installing on Ubuntu)
Configuring your node as a Relay
Testnet nodes and the update process
The current release of our testnet package is designed around simple installation and a simplified auto-update mechanism.  A node installation consists of 2 folders: the Binaries (bin) and the Data (data) folders.  The bin folder can be anywhere you choose.  A recommended location is ~/node.  We currently assume the folder is dedicated to algorand binaries, as we archive the folder before each update (we do not currently delete anything, but will overwrite our own binaries and add new ones).  The data folder is assumed to be ~/.algorand-testnet, but can be overridden.  We will create it if it doesn’t exist.
The bin folder holds our executables (goal, algod, carpenter, doberman, and updater), our update script (update.sh), and our optional configuration files (config.json and phonebook.json) -- we only distribute a sample config.json.example now.
The data folder holds the ledger (algorand.ledger), a hash file for the genesis ledger, generated log files (node-timestamp.log, and a latest.log symlinked to the latest log file).  You can optionally create a phonebook.json and/or config.json file and add configuration options that are merged into the defaults read from the bin folder.  We also generate an algorand.ledger.sha file containing the hash of the current genesis ledger (we generate the hash when a new one is installed and use it to know when the genesis ledger has been regenerated).

Installing a new node is generally a 3-step process:
Prepare your environment by installing dependencies
Retrieve the installation package appropriate for your platform
Optionally, verify the hash of the tar.gz file
Run update.sh -i -c stable -p ~/node

When the installer runs, it will pull down the latest update package from S3 for your platform and install it.  When it has finished, it will launch the node and you should be up and running.
Note: After the installation script completes, your terminal may appear hung but it’s just waiting for you to hit a key.

To check for, and install, the latest update, you can simply run update.sh (without arguments) at any time.  It will query S3 for available builds and see if any are newer than your installed version (to force an update, you can delete algod and run update -i -c stable again).  If there is a newer version, it’s downloaded and unpacked before we shut down your node, back up your files, and install the update.  If any part of the process fails, we attempt to restore your previous version (bin and data) and restart the node.  If it succeeds, we’ll start the new version of the node (even if it wasn’t running when you initiated the update).
Installing on a Mac
To install on a Mac (OSX v10.13.4 / High Sierra):
Open a terminal window
Install the libsodium package
Follow these instructions if you do not have homebrew installed: http://macappstore.org/libsodium/
Otherwise brew install libsodium
Install the libgmp package
brew install gmp
Create a temporary folder to hold the install package and files
mkdir ~/inst
cd ~/inst
Use wget to grab the installation package
wget TBD
Optionally
wget TBD
Verify SHA for the file
Unzip the package (use the appropriate filename)
tar -xf install_darwin-amd64.tar.gz
Create a bin folder for your node binaries and move the installation files there
mkdir ~/node
mv update* ~/node
cd ~/node
Run the installation script from your bin folder
./update.sh -i -c master

As an alternative to the 4 steps above, you can try this: 
./update.sh -i -c master -p ~/node

You should now have a running algorand node!  You can verify the daemon is running:
pgrep algod

If it outputs a number (process id) then algod is running.
You can watch the agreement activity by running the carpenter utility (it’s a long-running program, so ctrl+c to stop watching):
./carpenter -file ~/.algorand-testnet/latest.log
Installing on Ubuntu
Nodes have been verified on Ubuntu 18.04.  Other Debian-based distros should work as well (use apt-get install rather than apt install)
Open a terminal window
Install the libsodium-dev package
sudo apt update
sudo apt install libsodium-dev
Create a temporary folder to hold the install package and files
mkdir ~/inst
cd ~/inst
NOTE: 
Use wget to grab the installation package
wget https://mail.shoots.us/_algo/install_linux-amd64.tar.gz
Optionally
wget https://mail.shoots.us/_algo/sha.txt
Verify SHA for the file
If wget isn’t installed:
sudo apt install wget and try again
Unzip the package
tar -xf install_linux-amd64.tar.gz
Create a bin folder for your node binaries and move the installation files there
mkdir ~/node
mv update* ~/node
cd ~/node
Run the installation script from your bin folder
./update.sh -i -c master
If you get an error trying to download the update:
sudo apt install ca-certificates
Try ./update.sh -i -c master again

You should now have a running algorand node!  You can verify the daemon is running:
pgrep algod

If it outputs a number (process id) then algod is running.
You can watch the agreement activity by running the carpenter utility (it’s a long-running program, so ctrl+c to stop watching):
./carpenter -file ~/.algorand-testnet/latest.log
Installing in Docker
A Docker image was created with Ubuntu 18.04 and the required libsodium-dev and wget packages pre-installed.  We also install a certificates package as the certs required for S3 are outdated.
Open a terminal window
Create a Docker container
docker run -it algorand/node
This should put you at the bash prompt inside the Docker container (the remaining steps are identical to Ubuntu above, starting after the ‘Install the libsodium-dev package’ step)

Note we do not currently configure the container for any persistent storage, so when you stop the container your node is wiped - there is no persistence!
Installing in a VM
See installing on Ubuntu
Configuring your node as a Relay
A benefit of our decentralized network implementation is that a relay is effectively the same as any other node.  The distinction currently is made by configuring a node to actively listen for connections from other nodes, and having itself advertised using SRV records available through DNS.

So, to configure your node as a relay:
Create a config.json to your ~/.algorand-testnet folder with the following json (by default, our relays listen on port 4160, but there is nothing requiring that):

{
    “NetAddress”: “:4160”
}

Configure your firewall / port-forwarding so the endpoint is accessible
Update our SRV records for the phonebook (when nodes start up they will query and see the updated list including your FQDN or IP address).  The phonebook is stored under <network>.algorand.network, which is managed by Google Domains, where <network> can currently be either devnet or testnet.  David, Nickolai, Naveed, Rotem, and perhaps others have access to algorand.network under Google Domains.
Start / restart your node.  NOTE: See the ! IMPORTANT ! note above before starting your node.
Configuring auto-update
To ensure your node is running the latest software, we recommend using CRON to check for updates on a schedule.   Since checking for updates is fairly lightweight and does not shut down your node unless an update is available, there is no harm in checking frequently.
We recommend the following:
> crontab -e

Add a line that looks like this (run update.sh every hour of every day), where ‘user’ is the name of the account you’re using to install / run the node:
0 * * * * /home/user/node/update.sh -d /home/user/.algorand-testnet >/home/user/node/update.log 2>&1
Notes on installing on other platforms
CentOS:
sudo yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum install libsodium

Ensure your machine’s time is correct.  Ensure you have an NTP service running.

On CentOS:
sudo yum install ntp
sudo chkconfig ntpd on
sudo ntpdate pool.ntp.org
sudo service ntpd start

