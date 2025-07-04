## Leader Election and Metrics for Controller Manager

- Added leader election support using a Lease resource (enabled by default, can be disabled with a flag).
- Added a flag to set the metrics port for the controller manager.
- Both features are configurable via CLI flags.

#### New flags:

- `--enable-leader-election` (default: true) — Enable/disable leader election for the controller manager.
- `--metrics-port` (default: 8081) — Port for controller manager metrics endpoint.

#### What it does:

Ensures only one instance of the controller manager is active at a time (HA support).
Exposes controller metrics on the specified port.


#### Usage:

1. To compile controller, run:
```sh
make build
```

2. To start server, run
```sh
./k8s-controller --log-level trace --kubeconfig ~/.kube/config server --enable-leader-election=false --metrics-port=9090
```