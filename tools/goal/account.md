# goal account

The `goal account` collection of commands are provided to support the creation and management of accounts / wallets tied to a specific Algorand node instance.

For all intents and purposes, account and wallet can be used interchangeably when discussing Algorand accounts.

| Command | Usage |
|------------|-|
| balance      | Retrieve the balance for the specified account |
| import       | Import an Algorand account |
| list         | Show the list of Algorand accounts on this machine |
| new          | Create an new Algorand account |
| onlinestatus | Change online status for the specified account |
| rewards      | Retrieve the recent rewards for the specified account |

## To create a new account
> `goal account new -n <name>`

Creates a new account with its corresponding *.rootkey wallet file in the current ledger directory (e.g. data/testnet-v10).  The name specified here is store in a local configuration file and is only used by `goal` when working against that specific node instance.

## To report the balances for the specific account
> `goal account balance -d <path_to_data_dir> -a <account_address>`

## To report details for all accounts tied to the specific node instance
> `goal account list -d <path_to_data_dir>`

Also indicates if the account is `[offline]` or `[online]`, and if the account is the Default account for `goal`.

## To report the recent participation rewards for the specific account for the specific range of blocks
> `goal account rewards -d <path_to_data_dir> -a <account_address> -s <start_block> -e <end_block>`

Note: this reports the participation rewards for the account between start and end, inclusive.  If the account was offline during some or all of this time, it will not accrue participation rewards for those offline rounds.  It may accrue delegation / offline rewards, but they will not be reported by this command.

## To change the online status of an account
*** NOT IMPLEMENTED ***

> `goal account onlinestatus -d <path_to_data_dir> -a <account_address> -o <set online>`

Set online should be 1 to set online, 0 to set offline.

## To import the wallet from another location into the node instance
*** NOT IMPLEMENTED ***

> `goal account import -d <path_to_data_dir> -p <path_to_wallet_files>`
