# Creating a Private Network

Use the `goal network` commands to create a private network.

The `goal network` collection of commands are provided to support the creation and management of 'private networks'.  These are fully-formed Algorand networks with private, custom Genesis ledgers running the current build of Algorand software.  Rather than creating a node instance based on the released genesis.json, these networks have their own and need to be manually connected.

The basic idea is that we create one or more data directories and wallets to form this network, specify which node owns which wallets, and can start/stop the network as a unit.  Each node is just like any other node running on TestNet or DevNet.

| Command | Usage |
|------------|-|
| create     | Create a private named network |
| delete     | Stops and Deletes a deployed private network |
| start      | Start a deployed private network |
| status     | Prints status for all nodes in a deployed private network |
| stop       | Stop a deployed private network |

## To create a new named network
> `goal network create -r ~/net1 -n private -t <path_to_template.json>`

This creates a collection of folders under ~/net1 that make up the entire private network named 'private' (simplifying cleanup).

## To start the private network
> `goal network start -r ~/net1`

## To stop the private network
> `goal network stop -r ~/net1`

## To check the status of all of the nodes running on the private network
> `goal network status -r ~/net1`

## And to delete the private network (stopping first if necessary)
> `goal network delete -r ~/net1`
#### NOTE: This does not prompt first - so be careful before you do this!

Once you have a private network running, you can create more accounts, create transactions, and otherwise treat it like any other set of running nodes.  It's entirely private, however.  If you want to open it up and allow others to join it, you can modify the config.json file and phonebook.json files to explicitly define listening ports and peer addresses.

## Sample Network Template JSON

To create a network, you first create a template file that defines the wallets and nodes comprising the network.  The wallet stake is specified in percent, and the percents should total 100%. Fractional percentages like 0.01 are allowed.
Online = 0 for wallets that will be marked as offline - they have no participation keys generated and cannot participate in consensus.  They will be eligible for offline rewards / incentives.  Online wallets are created with corresponding participation keys, good for the first 10 million blocks (for now).
IsRelay indicates the node is intended to be a relay - there must be at least one relay included in any network. Non-relay nodes will connect to the nodes marked as relays.
ParticipationOnly indicates the wallet only has access to participation keys - not rootkeys. The default behavior is for ParticipationOnly to be false.

```json
{
    "Genesis": {
        "NetworkName": "",
        "Wallets": [
            {
                "Name": "Wallet1",
                "Stake": 50,
                "Online": 1
            },
            {
                "Name": "Wallet2",
                "Stake": 40,
                "Online": 1
            },
            {
                "Name": "Wallet3",
                "Stake": 10,
                "Online": 0
            }
        ],
        "Nodes": [
            {
                "Name": "Primary",
                "IsRelay": true,
                "Wallets": [
                    { "Name": "Wallet1",
                      "ParticipationOnly": false }
                ]
            },
            {
                "Name": "Node",
                "Wallets": [
                    { "Name": "Wallet2",
                      "ParticipationOnly": false },
                    { "Name": "Wallet3",
                      "ParticipationOnly": false }
                ]
            }
        ]
    }
}
```

## Cloning a node in the network

Once a network is created you can add nodes by cloning existing ones. To do this, open a terminal and change the directory to networks directory. For the above examples this would be ~/net1. To clone the node enter
> 'goal node clone -d Primary -n -t NewNode'

Any of the nodes in the directory can be cloned. The above example clones the Primary node. This will create the configuration files in a NewNode directory for the new node. If the network is running you can start the node using goal node start -d path_to_newnode or stop the entire network and restart it. 

## Stopping/Starting a node in the network
In addition to starting or stopping the network (which is the preferred method), you can individually start or stop a node by using goal and passing the proper data directory for the specific node. When starting the node you must provide the peers list of listening relays in the network. This should be a semi-colon separated list. You can get these by looking at the contents of the algod-listen.net file of nodes in the network. These files will be in the data directories of each relay in the Private network. In the template above this will only be the Primary node (based on isRelay property) but in the third example below we show starting a node with more than one relay. Note that if you only have one relay and stop it the network will not continue until you start that node again. In this case you would not use the -p command line option as this is the relay node and would only require the start command with the data directory.  Note also that if you restart a relay, it will likely get assigned a different listening port, so the network will not recover unless you restart the other nodes with the new peer address. In this case, use `goal network stop` and `goal network start` to restart everything so they can connect.

> 'goal node stop -d ~/net1/Node'

> 'goal node start -d ~/net1/Node' -p "127.0.0.1:52530"

> 'goal node start -d ~/net1/Node' -p "$(cat ~/net1/Primary/algod-listen.net)"

> 'goal node start -d ~/net1/ClonedNode' -p "$(cat ~/net1/Primary/algod-listen.net);$(cat ~/net1/Node/algod-listen.net)"

