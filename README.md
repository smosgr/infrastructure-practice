
# Platform engineering Test

We have a team with a new app that they are trying to deploy into our kubernetes cluster. They are having
a few issues getting it deployed and wondered if you could help. They also need some help adding a health check
and serving some files from a ConfigMap.

## Environment

You will need a running Kubernetes 'cluster' in order to complete this test. We recommend the follow:

   - [minikube](https://github.com/kubernetes/minikube): a cross-platform tool for running a local
     Kubernetes cluster;

You should be running **Kubernetes 1.19 or later**. If you have previously used minikube you may need to
delete and recreate your local kubernetes cluster. You can check the version of Kube that's running on
Minikube with `kubectl version` and, if appropriate, delete the local cluster with `minikube delete`.

You will need to enable the ingress addon:

```
$ minikube addons enable ingress
```

In order to build and run the HTTP server and container you will need:

   - [Docker](https://docs.docker.com/install/)
   - [Golang](https://golang.org/doc/install)
   - `make`

You can compile the binary with the following command:

```
$ make
```

To build a container so it is available to minikube use:

```
$ make docker
```

You can apply the configuration to run our greet service in minikube with:

```
$ kubectl apply -f kubernetes.yaml
```

## Steps

   1. Find and fix any issues with the server and configuration.
   2. Write a new `/healthz/` health check endpoint and setup a readiness probe.
   3. Serve the included files in `assets/` from a ConfigMap through a new handler at `/assets` through the HTTP server.
