# goal logging

`goal logging` is used to control the optional logging and telemetry.

| Command | Usage |
|------|------------------------|
| disable    | Disable Algorand remote logging |
| enable     | Enable Algorand remote logging |
| send       | Upload logs and data for analysis |

## To disable logging
> `goal logging disable`

## To enable logging
> `goal logging enable -n nodename`

This will turn on remote logging. The "friendly name" for the node, used by logging, will be determined by `-n nodename`.

## To upload logs to Algorand
> `goal logging send -d <path_to_datadir>`

This will zip and timestamp the node logs, and upload them to Algorand.