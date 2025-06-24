## GitHub workflow with build stage

#### CI workflow builds artifact (docker image) and publishes to GitHub container registry

[ci pipeline](https://github.com/AdamDubnytskyy/k8s-controller/actions/workflows/ci.yml)

[GitHub container registry](https://github.com/AdamDubnytskyy/k8s-controller/pkgs/container/app)


## Makefile

#### Commands:
- To compile kubernetes controller app binary, run:
  ```sh
  make build
  ```
- To run tests:
  ```sh
  make test
  ```
- To run kubernetes controller application:
  ```sh
  make run
  ```
- To build docker image artifact, run:
  ```sh
  make docker-build
  ```
- To cleanup kubernetes controller app binary, run:
  ```sh
  make clean
  ```