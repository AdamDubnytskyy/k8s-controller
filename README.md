# kubernetes controller

<div align="center">
  <p><em>Build your own Kubernetes controller from scratch</em></p>
</div>

# Overview
This project helps you understand Kubernetes controller concept by providing step-by-step instructions on how to build a Kubernetes controller from scratch. 

It's perfect for developers who want to:
- Learn and understand how to build a Kubernetes conroller

#### Inspired by [Golang Kubernetes Controller Tutorial](https://github.com/den-vasyliev/k8s-controller-tutorial-ref)
Authors: [Den Vasyliev](https://github.com/den-vasyliev) and [Alex Sudom](https://github.com/Alex0M)

## Features
- A starter template for building Kubernetes controllers or CLI tools in Go using [cobra-cli](docs/cobra-cli/README.md).
- [Zerolog logging](docs/zerolog-logging/README.md)
- POSIX/GNU-style --flags with [Pflag](docs/pflag/README.md)

## Prerequisites
- Basic understanding of Kubernetes concepts

---

## Steps
- [x] [Cobra CLI](docs/cobra-cli/README.md)
- [x] [Zerolog logging](docs/zerolog-logging/README.md)
- [x] [Pflag, implementing POSIX/GNU-style --flags](docs/pflag/README.md)
- [x] [Fast HTTP](docs/fast-http-server/README.md)
- [x] [Makefile, CI](docs/ci/README.md)
- [x] [kubernetes/client-go](docs/go-client/README.md)
- [ ] [Informer]()
- [ ] [API handler]()
- [ ] [sigs.k8s.io/controller-runtime]()
- [ ] [Leader election]()