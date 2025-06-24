# List Kubernetes Deployments with client-go

- Added a new `list` command using [k8s.io/client-go](https://github.com/kubernetes/client-go).
- Lists deployments in the specified namespace
- Supports:
   - `--kubeconfig` flag to specify the kubeconfig file for authentication
   - `--namespace` flag to specify Kubernetes namespace
- Uses zerolog for error logging

**Usage:**

```sh
go run main.go --log-level <log-level> --kubeconfig ~/.kube/config --namespace <namespace> list
```
E.g.:
```sh
go run main.go --log-level debug --kubeconfig ~/.kube/config --namespace default list
```


**What it does:**
- Connects to the Kubernetes cluster using the provided kubeconfig file.
- Lists all deployments in the respective namespace and prints their names.