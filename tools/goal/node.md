# goal node

The `goal node` collection of commands are provided to support the creation and management of Algorand node instances, where each instance corresponds to a unique data directory.

| Command | Usage |
|------------|-|
| clone      | Clone the specified node to create another node |
| restart    | stop, and then start, the specified Algorand node |
| start      | Init the specified algorand node |
| status     | Get the current node status |
| stop       | stop the specified Algorand node |

## To clone an existing node instance (without wallets)
> `goal node clone -d <path_to_data_dir> -t <path_to_cloned_dir>`

Optionally you can control whether the clone includes the current ledger, or if it starts with an uninitialized one.  The default is to clone the ledger as well.  Specify `-n` or `--noledger` to start with an uninitialized ledger.

## To clone an existing instance without the ledger
> `goal node clone -n -d <path_to_data_dir> -t <path_to_cloned_dir>`

## To start the node instance
> `goal node start -d <path_to_data_dir>`

## To stop the node instance
> `goal node stop -d <path_to_data_dir>`

## To restart the node instance (quicker version of stop/start)
> `goal node restart -d <path_to_data_dir>`

## To check the status of the node instance
> `goal node status -d <path_to_data_dir>`
