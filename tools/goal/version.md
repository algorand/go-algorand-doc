# goal version

`goal version` is used to determine node version information.

| Command | Usage |
|------|------------------|
| version    | print version information and exit |

## To print version information
> `goal version -d <path_to_datadir>`

## To print extra version information
> `goal version -v -d <path_to_datadir>`

In addition to printing the version, this will print the GenesisID, the update channel, and the commit branch and hash.