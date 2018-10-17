# go-algorand-doc

## Getting started with TestNet

We are building up this repo as we roll out TestNet and on-board external users.

To get started, refer to the [Node Setup Guide](https://github.com/algorand/go-algorand-doc/blob/master/Node%20Setup%20Guide.md).

We are using many different communication channels for discussing TestNet.  Our official channel for TestNet support is https://community.algorand.com/.  We will also be monitoring this Github repo for issues and discussions.

Please contact us at testnet@algorand.com if you're not sure who to contact directly.

If you want to suggest an alternative communications channel, feel free to do so and we'll consider it.

## Node hardware requirements (subject to change)
At this time, we're expecting participants to run standalone Nodes and not Relays, so the hardware requirements are fairly minimal.  You need 4-8GB RAM, 100GB HDD/SSD, and 10Mbit broadband.  The more cores in your CPU the better, but generally 4 cores are more than enough for a single node.  There are diminishing returns after that.  There is no specific GPU-optimized code, so your graphics card should have no impact.

## Once you have a running node
Ensure you have enabled telemetry and send us your Node name and GUID so we can correlate telemetry properly.

It's important that you are configured to update regularly or you risk being disconnected from the network and unable to connect until after you update. Not to mention falling behind in features and bug fixes.  We recommend setting up a CRON job as outlined in the [Node Setup Guide](https://github.com/algorand/go-algorand-doc/blob/master/Node%20Setup%20Guide.md).  If you want to manually check for an update, use `./update.sh -d ~/node/data` as discussed in the Setup Guide.

## Using GOAL
Run `goal --help` to get help.

Most commands need you to specify the location of the data folder.  You can set ALGORAND_DATA in the environment if you want to skip that step when using `goal`.

e.g.

    ./goal node status -d ~/node/data
    ./goal account new -d ~/node/data

Once you create an account, you can use our [Dispenser](https://bank.testnet.algorand.network) to transfer tokens to your account.

The `./goal clerk` command is used to generate your own transactions.

We currently have a [dashboard](http://r1.algorand.network:5001) running for TestNet, which displays the view of the blockchain from one of our Relays.

We sometimes have a script running that's generating random transactions between some test accounts, generating ~4 TPS on the network.

## Writing your own client
Refer to our [REST API](https://github.com/algorand/go-algorand-doc/blob/master/api/REST.md) documentation as a starting point for writing your own clients.

Let us know where you would like more documentation and we'll look at prioritizing that.

Thanks, and **Welcome to Algorand's TestNet!**
