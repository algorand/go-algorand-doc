# REST API (Remote)

## API Token update

The testnet build released 11/21/18 added security to most API calls.

The token is located in the data directory in "algod.token" after execution of any goal command.

Pass the token in the header of the API call - "X-Algo-API-Token={token}"

    - Ex. curl http://127.0.0.1:8080/accounts -H "X-Algo-API-Token:{token}"

## API Reference

| Call | Parameter | Returns | Type | Description |
|------|-----------|---------|------|-------------|
|/account/{pk string}|pubkey string|number|GET|Returns the balance of \{ pk string \}|/account/{pk string}/passphrase|passphrase string	GET|Returns the passphrase of \{ pk string\}|
|/account/{pk string}|pubkey string; pubkey string, transactionFee uint64, validRounds uint64, byte status|txid string|PATCH| Creates and broadcasts either a "go online" or "go offline" transaction, and returns the txid of that transaction
|/account/{pk string}/partkey|pubkey string; feeLimit uint64, keyFirstValid uint64, keyLastValid uint64|-|POST| Attempts to create and store a participation key with the passed fee limit. Key will be valid from round keyFirstValid to round keyLastValid. Key will be used for going online/participating in consensus.
|/account/{pk string}/transaction/{txid}|pubkey string, txid string|transaction|GET|Returns the specific transaction information|
|/account/{pk string}/transactions/{limit}|pubkey string, (optional) int limit|list of transaction (format TBD)|GET|Return a list of public key's n recent transactions within the given range. The number of results is limited by the limit parameter.|
|/accounts|-|pk string|POST|Creates a new account and returns the corresponding public key|
|/accounts|-|list of string|GET|Returns the list of accounts registered on the node|
|/accounts|passphrase string|pk string|PUT|Given a passphrase returns the user account|
|/block/{round}|round int|block data|GET|Given a round number, returns the block information|
|/blocks/{limit}|(optional) int limit|list of block data|GET|gives the block information about the {limit} most recent blocks
|/ledger/balances|-|list of account statuses: string address, uint64 balance, byte onlineStatus|GET|Returns the list of all accounts with balances, and whether they are Online
|/ledger/rewardsclaims|-|list of rewards statuses: uint64 round, string address, int16 numRewards|GET|Returns the list of all outstanding, valid, Rewards claims
|/ledger/rewardstree|-|list of account statuses: string address, uint64 balance, byte onlineStatus|GET|Returns the list of all accounts in the current rewards tree
|/ledger/supply|-|uint64 round, uint64 totalMoney, uint64 onlineMoney|GET|returns the current token supply as tracked by the ledger
|/net/peers|-|list of string peers|GET|get the node's peer list
|/net/peer|ip:port string|bool success|POST|attempt to add the passed peer to the node's peer list
|/passphrase|-|string|POST|Generates, and then returns, a cryptographically secure passphrase.|
|/rewards|-|list of (address string, uint64 round, uint16 numRewards)|returns the list of accounts receiving rewards, how much they were rewarded, and on what round
|/status|-|JSON containing string and int64|GET|Returns the current status of the node: the last round seen, the last consensus version supported, the next consensus version to use, the round at which the next consensus version will apply, whether the next consesnsus version is supported by this node, the time since last round, and the catchup time.
|/status/wait-for-block-after/{round}|round int|same as /status|GET|waits for the passed round, and then returns the status
|/transactions|from, to string, (optional) fee, amount number|txid string|POST|Posts transaction to the network|
|/transactions/fee|-|number|GET|Returns the current suggested fee in the network|
|/transactions/{limit}|(optional) int limit|list of transactions|GET|Returns the list of confirmed transactions from all accounts on the node.|
|/transactions/pool/{limit}|(optional) int limit|list of transactions|GET|Returns the list of pending transactions (from any account) in the transaction pool
|/transactions/pool/stats|-|int, int, int|GET|Returns the current transaction pool stats: the number of transactions confirmed, outstanding, and expired.
|/version|-|string|GET|Returns Algorand's node version|
