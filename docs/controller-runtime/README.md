## controller-runtime Deployment Controller
- Integrate [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime) into project.

- Add deployment controller that runs business logic on each reconcile event for Deployments in specified namespace.

- Controller is started alongside the FastHTTP server.

#### What it does:

- Uses controller-runtime's manager to run a controller for Deployments.
- Adds OTEL environment variables to the deployment and creates OTEL collector ConfigMap object. Deployment **MUST** have label set: `app=nodejs`

#### Usage

1. To compile controller, run:
```sh
make build
```

2. To start server, run
```sh
./k8s-controller --log-level trace --kubeconfig ~/.kube/config server --namespace <namespace> --labelSelectors labelName=labelValue
```

3. Create a new deployment or modify existing one by setting `app=nodejs` label set.

    - E.g.: create a new deployment.
    ```sh
    kubectl create deploy nodejs-app --image=<image> --replicas=1 --labels=app=nodejs -n <namespace>
    ```

    - E.g.: Add required label set `app=nodejs` to existing deployment.
    ```sh
    kubectl label deployment test app=nodejs -n <namespace>
    ```

4. Verify specified namespace, and observe newly added environment variables and annotations for OTEL automatic instrumentation to the deployment and new ConfigMap object done by controller's business logic.

```sh
kubectl describe deploy <name> -n <namespace>
kubectl get cm -n <namespace>
```