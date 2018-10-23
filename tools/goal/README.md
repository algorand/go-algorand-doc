# GOAL Command-Line Interface (CLI) Tool

GOAL is the CLI for interacting with a running Algorand instance. The binary 'goal' is installed alongside the algod binary and is considered an integral part of the complete installation. The binaries should be used in tandem - you should not try to use a version of goal with a different version of algod.

## GOAL Commands

|  Usage |
|-|
| goal [flags]|
| goal [command] |

Available Commands:

| Command | Usage |
|------|------------------------|
| account | Control and manage Algorand accounts |
| clerk   | Provides the tools to control transactions |
| help    | Help about any command |
| ledger  | Access ledger-related details |
| license | Display license information |
| logging | Control and manage Algorand logging |
| network | Create and manage private, multi-node, locally-hosted networks |
| node    | Init, stop and get the status of the specified algorand node |
| version | The current version of the Algorand daemon (algod) |

| Flags | Usage |
|-|-|
| -d, --datadir string | Data directory for the node |
| -h, --help           | help for goal |

Use "goal [command] --help" for more information about a command.

---

### [`goal account`](https://github.com/algorand/go-algorand-doc/blob/master/tools/goal/account.md)

Create, delete, list and control algorand accounts

| Usage |
|-|
| goal account [flags] |
| goal account [command] |

Available Commands:

| Command | Usage |
|------|------------------------|
| balance      | Retrieve the balance for the specified account |
| import       | Import an Algorand account |
| list         | Show the list of Algorand accounts on this machine |
| new          | Create an new Algorand account |
| onlinestatus | Change online status for the specified account |
| rewards      | Retrieve the recent rewards for the specified account |

| Flags | Usage |
|-|-|
| -f, --default      | Set this account as the default one |
| -h, --help         | help for account |
| -n, --name string  | Rename this account |

---

### [`goal clerk`](https://github.com/algorand/go-algorand-doc/blob/master/tools/goal/clerk.md)

Send and view recent transactions and balances

| Usage |
|-|
| goal clerk [flags] |
| goal clerk [command] |

| Command | Usage |
|------|------------------------|
| recent      | Shows a list of the most recent transaction (defaults to last 10) |
| send        | Send money to an address |

| Flags | Usage |
|-|-|
| -h, --help  | help for clerk |

| Global Flags | Usage |
|-|-|
| -d, --datadir string  | Data directory for the node |

---

### `goal help`

---

### `goal ledger`

Access ledger-related details

| Usage |
|-|
| goal ledger [flags] |
| goal ledger [command] |

| Command | Usage |
|------|------------------------|
| balances    | Show all non-zero account balances |
| rewardstree | Shows the current rewards tree |

| Flags | Usage |
|-|-|
| -h, --help  | help for ledger |

| Global Flags | Usage |
|-|-|
| -d, --datadir string  | Data directory for the node |

---

### `goal license`

Displays license information

| Usage |
|-|
| goal license [flags] |

| Flags | Usage |
|-|-|
| -h, --help  | help for license |

---

### `goal logging`

Enable/disable and configure Algorand remote logging

| Usage |
|-|
| goal logging [flags] |
| goal logging [command] |

| Command | Usage |
|-------------|-|
| disable     | Disable Algorand remote logging |
| enable      | Enable Algorand remote logging |
| send        | Upload logs and data for analysis |

| Flags | Usage |
|-|-|
| -h, --help  | help for logging |

| Global Flags | Usage |
|-|-|
| -d, --datadir string  | Data directory for the node |

---

### [`goal network`](https://github.com/algorand/go-algorand-doc/blob/master/tools/goal/network.md)

Allows creating, starting, stopping, and deleting private, locally-hosted networks

| Usage |
|-|
| goal network [flags] |
| goal network [command] |

| Command | Usage |
|------------|-|
| create     | Create a private named network |
| delete     | Stops and Deletes a deployed private network |
| start      | Start a deployed private network |
| status     | Prints status for all nodes in a deployed private network |
| stop       | Stop a deployed private network |

| Flags | Usage |
|-|-|
| -h, --help            | help for network |
| -r, --rootdir string  | Root directory for the private network directories |

| Global Flags | Usage |
|-|-|
| -d, --datadir string  | Data directory for the node |

---

### [`goal node`](https://github.com/algorand/go-algorand-doc/blob/master/tools/goal/node.md)

Allows to issue, stop and get the status of the specified algorand node

| Usage |
|-|
| goal node [flags] |
| goal node [command] |

| Command | Usage |
|------------|-|
| clone      | Clone the specified node to create another node |
| restart    | stop, and then start, the specified Algorand node |
| start      | Init the specified algorand node |
| status     | Get the current node status |
| stop       | stop the specified Algorand node |

| Flags | Usage |
|-|-|
| -h, --help  | help for node |

| Global Flags | Usage |
|-|-|
| -d, --datadir string  | Data directory for the node |

---

### `goal version`

The current version of the Algorand daemon (algod)

| Usage |
|-|
| goal version [flags] |

| Flags | Usage |
|-|-|
| -h, --help     | help for version |
| -v, --verbose  | Print all version info available |

| Global Flags | Usage |
|-|-|
| -d, --datadir string  | Data directory for the node |
