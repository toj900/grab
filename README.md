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

1. Run main

    ```bash
    go run main.go
    ```

    output

    ```bash
    Create example CustomResources from CRDs

    Usage:
    grabschema [crd-name] [flags]

    Examples:

    Create example CustomResources from CRDs
    $ grabschema


    Flags:
    -h, --help   help for grabschema
    ```

1. Grab example CRs for the `buckets.source.toolkit.fluxcd.io` CRD

    ```bash
    go run main.go buckets.source.toolkit.fluxcd.io
    ```

    output

    ```bash
    metadata:
    name: name
    namespace: namespace
    spec:
    # region: "string" # The bucket region.
    # suspend: true # This flag tells the controller to suspend the reconciliation of this source.
    # timeout: "string" # The timeout for download operations, defaults to 60s.
    # accessFrom:  # AccessFrom defines an Access Control List for allowing cross-namespace references to this object.
        # namespaceSelectors:  # NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.
        # - matchLabels:  # MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
    endpoint: "string" # The bucket endpoint address.
    # ignore: "string" # Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.
    # insecure: true # Insecure allows connecting to a non-TLS S3 HTTP endpoint.
    interval: "string" # The interval at which to check for bucket updates.
    bucketName: "string" # The bucket name.
    # provider: "string" # The S3 compatible storage provider name, default ('generic').
    # secretRef:  # The name of the secret containing authentication credentials for the Bucket.
        # name: "string" # Name of the referent.
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
