## Implement /deployments API endpoint

1. Adapt client-side caching mechanism to reduce API server calls to `api-server` Kubernetes cluster component.
2. Add `/deployments` API endpoint to FastHTTP server
3. Return JSON array of deployment names from the informer's local cache.

### Informer's local cache

- [cache](https://pkg.go.dev/k8s.io/client-go@v0.33.2/tools/cache#pkg-overview)

    Package cache is a client-side caching mechanism. It is useful for reducing the number of server calls you'd otherwise need to make. Reflector watches a server and updates a Store. Two stores are provided; one that simply caches objects (for example, to allow a scheduler to list currently available nodes), and one that additionally acts as a FIFO queue (for example, to allow a scheduler to process incoming pods).

- [GetStore](https://pkg.go.dev/k8s.io/client-go@v0.33.2/tools/cache#SharedInformer.GetStore) returns the informer's local cache as a Store

    E.g.:
    ```sh
    deploymentInformer.GetStore().List()
    ```

### Usage:

1. To compile controller, run:
```sh
make build
```

2. To start server, run
```sh
./k8s-controller --log-level trace --kubeconfig ~/.kube/config server --namespace <namespace> --labelSelectors labelName=labelValue
```

3. Call `/deployments` API endpoint to list deployment names for specified namespace and labelSelectors in step 2.
```sh
curl http://localhost:<port>/deployments
```