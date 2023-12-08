# gmx-go-sdk

Golang SDK for the GMX protocol

## Examples

Check out the [examples](./examples) folder for some examples on how to use the SDK.

## Generating the SDK

1. abigen

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

2. forge

```bash
curl -L https://foundry.paradigm.xyz | bash
foundryup
```

3. yarn

```bash
yarn install
```

4. Run the generator

```bash
make build && make generate
```
