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
