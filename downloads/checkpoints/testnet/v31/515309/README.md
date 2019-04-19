<h1>Bootstrapping your node from a checkpoint.</h1>

Currently we only support bootstrapping non-archival nodes.

To bootstrap, download the tar.gz file and extract is (eg `tar -xf NonArchival-Round515309.tar.gz`).
Verify the SHA256 values for the two files.
Stop your node if it is running.
Move the *.sqlite files into the ledger directory (eg `~/node/data/testnet-v31.0`)

Start your node.