# Kubernetes controller

![k8s-controller release (latest SemVer)](https://img.shields.io/github/v/tag/AdamDubnytskyy/k8s-controller?sort=semver)
[![Go Reference](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/AdamDubnytskyy/k8s-controller)
[![Test Status](https://github.com/AdamDubnytskyy/k8s-controller/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/AdamDubnytskyy/k8s-controller/actions/workflows/tests.yml)

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
- [up & running control-plane](docs/control-plane/README.md)

## CI
To setup GitHub Actions workflow follow steps described below:
1. Generate PAT (personal access token). Select listed scopes: write:packages, repo. Copy secret.
2. Go to `Settings` tab of your repository, choose `Security` section and select `Actions`. Create a `New repository secret` with name `K8S_CONTROLLER_TOKEN`, paste secret copied generated on step 1.

---

## Steps
1. [Cobra CLI](docs/cobra-cli/README.md) ✅
2. [Zerolog logging](docs/zerolog-logging/README.md) ✅
3. [Pflag, implementing POSIX/GNU-style --flags](docs/pflag/README.md) ✅
4. [Fast HTTP](docs/fast-http-server/README.md) ✅
5. [Makefile, CI](docs/ci/README.md) ✅
6. [kubernetes/client-go](docs/go-client/README.md) ✅
7. [Informer](docs/informer/README.md) ✅
8. [API handler](docs/api-handler/README.md) ✅
9. [controller-runtime](docs/controller-runtime/README.md) ✅
10. [Leader election](docs/leader-election/README.md) ✅
11. [FrontendPage CRD and Advanced Controller Implementation]()