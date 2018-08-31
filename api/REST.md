# REST API (Remote)

| Call | Parameter | Returns | Type | Description |
|------|-----------|---------|------|-------------|
|/account/{pk string}||number|GET|Returns the balance of \{ pk string \}|/account/{pk string}/passphrase|passphrase string	GET|Returns the passphrase of \{ pk string\}|
|/account/{pk string}/transaction/{txid}||transaction|GET|Returns the specific transaction information|
|/account/{pk string}/transactions/{limit}|(optional) start, end string|list of transaction (format TBD)|GET|Return a list of public key's n recent transactions within the given range. The number of results is limited by the limit parameter.|
/accounts||pk string|POST|Creates a new account and returns the corresponding public key|
|/accounts||list of string|GET|Returns the list of accounts registered on the node|
|/accounts|passphrase string|pk string|PUT|Given a passphrase returns the user account|
|/block/{round}||block data|GET|Given a round number, returns the block information|
|/passphrase||string|GET|Returns a cryptographically secure passphrase.|
|/transactions|from, to string, (optional) fee, amount number|txid string|POST|Posts transaction to the network|
|/transactions/fee||number|GET|Returns the current suggested fee in the network|
|/transactions/{limit}||list of transactions|GET|Returns the list of confirmed transactions from all accounts on the node.|
|/version||string|GET|Returns Algorand's node version|
