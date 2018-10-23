# goal clerk

The `goal clerk` collection of commands are provided to support the mangement of transaction information.

For all intents and purposes, account and wallet can be used interchangeably when discussing Algorand accounts.

| Command | Usage |
|------|------------------------|
| recent      | Shows a list of the most recent transaction (defaults to last 10) |
| send        | Send money to an address |

## To see recent transactions
> `goal clerk recent -t <num> -d <path_to_data_dir>`

Returns the `-t` most recent transactions that the node knows about. If no `-t` is provided, defaults to `10`.

## To send tokens
> `goal clerk send -f <fromAddr> -t <toAddr> -a <amount> --fee <fee> --firstvalid <roundNum> --lastvalid <roundNum> -d <path_to_data_dir>`

Creates a transaction sending `amount` tokens from `fromAddr` to `toAddr`. If the optional `--fee` is not provided, the transaction will use the recommended amount. If the optional `--firstvalid` and `--lastvalid` are provided, the transaction will only be valid from round `firstValid` to round `lastValid`.
If broadcast of the transaction is successful, the transaction ID will be returned. 