# Containerized docker algorand private development node

## Description

This folder contains the files necessary to build a container image that can be used to run an algorand private local node.
Once you have built the image, you can run it and simply connect to the REST API to the port 7979 (see scripts below).

## Build

```bash
docker build -t randpool/node:v1.0.0 ./
```

## Run

```bash
docker run -dit --name randpool-node \
    -p 7979:7979 \
    -p 7833:7833 \
    --rm \
    randpool/node:v1.0.0
```

## Test

To test if your container is running correctly, you can run this golang source code, which should give you the node status if successful:

```bash
go run main.go --token=$(docker exec randpool cat /root/randpool/Primary/algod.token) --port=7979
```

```golang
    package main

    import (
        "flag"
        "fmt"

        "github.com/algorand/go-algorand-sdk/client/algod"
    )

    func main() {
        port := flag.Int("port", 7979, "port used for the REST API")
        token := flag.String(
            "token", "-",
            "token used by the node, docker exec randpool cat /root/randpool/Primary/algod.token",
        )
        flag.Parse()

        algodAddress := fmt.Sprintf("http://127.0.0.1:%d", *port)
        algodToken := *token

        // Create an algod client
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            fmt.Printf("failed to make algod client: %s\n", err)
            return
        }
        // Get algod status
        nodeStatus, err := algodClient.Status()
        if err != nil {
            fmt.Printf("error getting algod status: %s\n", err)
            return
        }
        fmt.Printf("algod last round: %d\n", nodeStatus.LastRound)
        fmt.Printf("algod time since last round: %d\n", nodeStatus.TimeSinceLastRound)
        fmt.Printf("algod catchup: %d\n", nodeStatus.CatchupTime)
        fmt.Printf("algod latest version: %s\n", nodeStatus.LastVersion)
    }
```
