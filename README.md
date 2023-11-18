# Grab
Kubernetes extension to grab schema

## Getting Started

1. Start a local cluster

```bash
kind create cluster
```

1. Apply CRDs

```bash
kubectl apply -f ./example/all-source-controller.crds.yaml
```

2. Run main

```bash
go run main.go
go run main.go --resource buckets.source.toolkit.fluxcd.io -r gitrepositories.source.toolkit.fluxcd.io
```

3. Clean up

```bash
kind delete cluster
```

## Running
```bash
# assumes you have a working KUBECONFIG
GO111MODULE="on" go build -o kubectl-grabschema
GO111MODULE="on" go build -o grabschema

# place the built binary somewhere in your PATH
sudo cp ./kubectl-grabschema /usr/local/bin
sudo cp ./kubectl_complete-grabschema /usr/local/bin

# you can now begin using this plugin as a regular kubectl command:
# Generate an example CR from existing CRDs
kubectl grabschema crd-name
```

### Cleanup
```bash
sudo rm -rf /usr/local/bin/kubectl-grabschema
sudo rm -rf /usr/local/bin/kubectl_complete-grabschema 
```
