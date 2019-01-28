# Algod REST API.


<a name="overview"></a>
## Overview
API Endpoint for AlgoD Operations.


### Security Tokens

A security token is required by most API calls.

The token is located in the data directory in "algod.token" after execution of any goal command.

The token can also be generated using the following command:

     - Goal node generatetoken -d data

Note that you have to specify the node data directory or have the $ALGORAND_DATA enviornment variable set.
The node must be shutdown to execute this command.

Pass the token in the header of the API call - "X-Algo-API-Token: {token}"

    - Ex. curl http://127.0.0.1:8080/accounts -H "X-Algo-API-Token:{token}"
    - Ex. local data directory curl "http://$(cat algod.net)/status" -H "X-Algo-API-Token: $(cat algod.token)" 


### Version information
*Version* : 0.0.1


### Contact information
*Contact Email* : contact@algorand.com


### URI scheme
*Host* : localhost  
*BasePath* : /  
*Schemes* : HTTP


### Consumes

* `application/json`


### Produces

* `application/json`




<a name="paths"></a>
## Paths

<a name="healthcheck"></a>
### Returns OK if healthy.
```
GET /health
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK.|No Content|
|**401**|Invalid API Token|No Content|
|**404**|Not Found|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="swaggerjson"></a>
### Gets the current swagger spec.
```
GET /swagger.json
```


#### Description
Returns the entire swagger spec in plaintext.


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|The current swagger spec|string|
|**401**|Invalid API Token|No Content|
|**404**|Not Found|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="accountinformation"></a>
### Get account information.
```
GET /v1/account/{address}
```


#### Description
Given a specific account public key, this call returns the accounts status, balance and spendable amounts


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|AccountInformationResponse contains an account information|[Account](#account)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="rewards"></a>
### Get the list of rewards between firstRound and lastRound
```
GET /v1/account/{address}/rewards
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Query**|**firstRound**  <br>*required*|Do not fetch any rewards before this round.|integer (int64)|
|**Query**|**lastRound**  <br>*required*|Do not fetch any rewards after this round.|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|RewardResponse contains a list of rewards for a specific address|[RewardList](#rewardlist)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="transactioninformation"></a>
### Get a specific confirmed transaction.
```
GET /v1/account/{address}/transaction/{txid}
```


#### Description
Given a wallet address and a transaction id, it returns the confirmed transaction information. This call scans up to config.Protocol.MaxTxnLife blocks in the past.


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Path**|**txid**  <br>*required*|A transaction id|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|TransactionResponse contains a transaction information|[Transaction](#transaction)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**404**|Transaction Not Found|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="transactions"></a>
### Get a list of confirmed transactions.
```
GET /v1/account/{address}/transactions
```


#### Description
Returns the list of confirmed transactions between firstRound and lastRound


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Query**|**firstRound**  <br>*required*|Do not fetch any transactions before this round.|integer (int64)|
|**Query**|**lastRound**  <br>*required*|Do not fetch any transactions after this round.|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|TransactionsResponse contains a list of transactions|[TransactionList](#transactionlist)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="getblock"></a>
### Get the block for the given round.
```
GET /v1/block/{round}
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round from which to fetch block information.|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|BlockResponse contains block information|[Block](#block)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="getsupply"></a>
### Get the current supply reported by the ledger.
```
GET /v1/ledger/supply
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|SupplyResponse contains the ledger supply information|[Supply](#supply)|
|**401**|Invalid API Token|No Content|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="getstatus"></a>
### Gets the current node status.
```
GET /v1/status
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|StatusResponse contains the node's status information|[NodeStatus](#nodestatus)|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="waitforblock"></a>
### Gets the node status after waiting for the given round.
```
GET /v1/status/wait-for-block-after/{round}/
```


#### Description
Waits for a block to appear after round {round} and returns the node's status at the time.


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round to wait until returning status|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|StatusResponse contains the node's status information|[NodeStatus](#nodestatus)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="rawtransaction"></a>
### Broadcasts a raw transaction to the network.
```
POST /v1/transactions
```


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Body**|**rawtxn**  <br>*required*|The byte encoded signed transaction to broadcast to network|string (binary)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|TransactionIDResponse contains a transaction information|[transactionID](#transactionid)|
|**400**|Bad Request|string|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Consumes

* `application/x-binary`


#### Produces

* `application/json`


<a name="suggestedfee"></a>
### Get the suggested fee
```
GET /v1/transactions/fee
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|TransactionFeeResponse contains a suggested fee|[TransactionFee](#transactionfee)|
|**401**|Invalid API Token|No Content|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="getpendingtransactions"></a>
### Get a list of unconfirmed transactions currently in the transaction pool.
```
GET /v1/transactions/pending
```


#### Description
Get the list of pending transactions, sorted by priority, in decreasing order, truncated at the end at MAX. If MAX = 0, returns all pending transactions.


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**max**  <br>*optional*|Truncated number of transactions to display. If max=0, returns all pending txns.|integer (int64)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|PendingTransactionsResponse contains a (potentially truncated) list of transactions and<br>the total number of transactions currently in the pool.|[PendingTransactions](#pendingtransactions)|
|**401**|Invalid API Token|No Content|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|


#### Produces

* `application/json`


<a name="getversion"></a>
### GET /versions

#### Description
Retrieves the current version


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|VersionsResponse is the response to `GET /versions`|[Version](#version)|


#### Produces

* `application/json`




<a name="definitions"></a>
## Definitions

<a name="account"></a>
### Account
Account Description


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*required*|Address indicates the account public key|string|
|**amount**  <br>*required*|Amount indicates the total number of Algos in the account|integer (uint64)|
|**status**  <br>*required*|Status indicates the delegation status of the account's Algos<br>Offline - indicates that the associated account is delegated.<br>Online  - indicates that the associated account used as part of the delegation pool.<br>NotParticipating - indicates that the associated account is neither a delegator nor a delegate.|string|


<a name="algos"></a>
### Algos
Algos is our unit of currency.  It is wrapped in a struct to nudge
developers to use an overflow-checking library for any arithmetic.


|Name|Schema|
|---|---|
|**Raw**  <br>*optional*|integer (uint64)|


<a name="block"></a>
### Block
Block contains a block information


|Name|Description|Schema|
|---|---|---|
|**balRoot**  <br>*required*|BalanceRoot is the root of the merkle tree after committing this block|string|
|**currentProtocol**  <br>*required*|CurrentProtocol is a string that represents the current protocol|string|
|**hash**  <br>*required*|Hash is the current block hash|string|
|**nextProtocol**  <br>*required*|NextProtocol is a string that represents the next proposed protocol|string|
|**nextProtocolApprovals**  <br>*required*|NextProtocolApprovals is the number of blocks which approved the protocol upgrade|integer (uint64)|
|**nextProtocolSwitchOn**  <br>*required*|NextProtocolSwitchOn is the round on which the protocol upgrade will take effect|integer (uint64)|
|**nextProtocolVoteBefore**  <br>*required*|NextProtocolVoteBefore is the deadline round for this protocol upgrade (No votes will be consider after this round)|integer (uint64)|
|**period**  <br>*required*|Period is the period on which the block was confirmed|integer (uint64)|
|**previousBlockHash**  <br>*required*|PreviousBlockHash is the previous block hash|string|
|**proposer**  <br>*required*|Proposer is the address of this block proposer|string|
|**round**  <br>*required*|Round is the current round on which this block was appended to the chain|integer (uint64)|
|**seed**  <br>*required*|Seed is the sortition seed|string|
|**timestamp**  <br>*required*|TimeStamp in seconds since epoch|integer (int64)|
|**txnRoot**  <br>*required*|TransactionsRoot authenticates the set of transactions appearing in the block.<br>More specifically, it's the root of a merkle tree whose leaves are the block's Txids, in lexicographic order.<br>For the empty block, it's 0.<br>Note that the TxnRoot does not authenticate the signatures on the transactions, only the transactions themselves.<br>Two blocks with the same transactions but in a different order and with different signatures will have the same TxnRoot.|string|
|**txns**  <br>*optional*||[TransactionList](#transactionlist)|
|**upgradeApprove**  <br>*required*|UpgradeApprove indicates a yes vote for the current proposal|boolean|
|**upgradePropose**  <br>*required*|UpgradePropose indicates a proposed upgrade|string|


<a name="nodestatus"></a>
### NodeStatus
NodeStatus contains the information about a node status


|Name|Description|Schema|
|---|---|---|
|**catchupTime**  <br>*required*|CatchupTime in nanoseconds|integer (int64)|
|**lastConsensusVersion**  <br>*required*|LastVersion indicates the last consensus version supported|string|
|**lastRound**  <br>*required*|LastRound indicates the last round seen|integer (uint64)|
|**nextConsensusVersion**  <br>*required*|NextVersion of consensus protocol to use|string|
|**nextConsensusVersionRound**  <br>*required*|NextVersionRound is the round at which the next consensus version will apply|integer (uint64)|
|**nextConsensusVersionSupported**  <br>*required*|NextVersionSupported indicates whether the next consensus version is supported by this node|boolean|
|**timeSinceLastRound**  <br>*required*|TimeSinceLastRound in nanoseconds|integer (int64)|


<a name="paymenttransactiontype"></a>
### PaymentTransactionType
PaymentTransactionType contains the additional fields for a payment Transaction


|Name|Description|Schema|
|---|---|---|
|**amount**  <br>*required*|Amount is the amount of Algos intended to be transferred|integer (uint64)|
|**to**  <br>*required*|To is the receiver's address|string|


<a name="pendingtransactions"></a>
### PendingTransactions
PendingTransactions represents a potentially truncated list of transactions currently in the
node's transaction pool.


|Name|Description|Schema|
|---|---|---|
|**totalTxns**  <br>*required*|TotalTxns|integer (uint64)|
|**truncatedTxns**  <br>*required*||[TransactionList](#transactionlist)|


<a name="rewardlist"></a>
### RewardList

|Name|Description|Schema|
|---|---|---|
|**awardee**  <br>*required*|Address is the rewardee address|string|
|**rewards**  <br>*required*||[Algos](#algos)|


<a name="supply"></a>
### Supply
Supply represents the current supply of Algos in the system


|Name|Description|Schema|
|---|---|---|
|**onlineMoney**  <br>*required*|OnlineMoney|integer (uint64)|
|**round**  <br>*required*|Round|integer (uint64)|
|**totalMoney**  <br>*required*|TotalMoney|integer (uint64)|


<a name="transaction"></a>
### Transaction
Transaction contains all fields common to all transactions and serves as an envelope to all transactions
type


|Name|Description|Schema|
|---|---|---|
|**fee**  <br>*required*|Fee is the transaction fee|integer (uint64)|
|**first-round**  <br>*required*|FirstRound indicates the first valid round for this transaction|integer (uint64)|
|**from**  <br>*required*|From is the sender's address|string|
|**last-round**  <br>*required*|LastRound indicates the last valid round for this transaction|integer (uint64)|
|**noteb64**  <br>*optional*|Note is a free form data|< integer (uint8) > array|
|**payment**  <br>*optional*||[PaymentTransactionType](#paymenttransactiontype)|
|**round**  <br>*optional*|ConfirmedRound indicates the block number this transaction appeared in|integer (uint64)|
|**tx**  <br>*required*|TxID is the transaction ID|string|
|**type**  <br>*required*||[TxType](#txtype)|


<a name="transactionfee"></a>
### TransactionFee
TransactionFee contains the suggested fee


|Name|Description|Schema|
|---|---|---|
|**fee**  <br>*required*|Fee is transaction fee|integer (uint64)|


<a name="transactionlist"></a>
### TransactionList
TransactionList contains a list of transactions


|Name|Description|Schema|
|---|---|---|
|**transactions**  <br>*required*|TransactionList is a list of rewards|< [Transaction](#transaction) > array|


<a name="txtype"></a>
### TxType
TxType is the type of the transaction written to the ledger

*Type* : string


<a name="version"></a>
### Version
Note that we annotate this as a model so that legacy clients
can directly import a swagger generated Version model.


|Name|Schema|
|---|---|
|**genesis_id**  <br>*required*|string|
|**versions**  <br>*required*|< string > array|


<a name="transactionid"></a>
### transactionID
TransactionID Description


|Name|Description|Schema|
|---|---|---|
|**txId**  <br>*required*|TxId is the string encoding of the transaction hash|string|




<a name="securityscheme"></a>
## Security

<a name="api_key"></a>
### api_key
Generated header parameter. This token can be generated using the Goal command line tool. Example value ='b7e384d0317b8050ce45900a94a1931e28540e1f69b2d242b424659c341b4697'

*Type* : apiKey  
*Name* : X-Algo-API-Token  
*In* : HEADER



