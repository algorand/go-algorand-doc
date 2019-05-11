# Spinning up a docker dev node

## Build

```bash
docker build -t randpool/devnode:v1.0.0 ./infra/devnode
```

## Run

```bash
docker run -dit --name randpool \
    -p 7979:7979 \
    --rm \
    randpool/devnode:v1.0.0
```
