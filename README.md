# Authenticating inside the cluster

This example shows you how to configure a client with client-go to authenticate
to the Kubernetes API from an application running inside the Kubernetes cluster.

client-go uses the [Service Account token][sa] mounted inside the Pod at the
`/var/run/secrets/kubernetes.io/serviceaccount` path when the
`rest.InClusterConfig()` is used.

## Running this example

First compile the application for Linux:

    cd flink-kubernetes-api
    make build-and-push

Then package it to a docker image using the provided Dockerfile to run it on
Kubernetes.

If you have RBAC enabled on your cluster, use the following
snippet to create role binding which will grant the default service account view
permissions.

```
kubectl -n flink-operator apply -f rbac.yaml 
```

Then, run the image in a Pod with a single instance Deployment:

    kubectl -n flink-operator apply -f deployment.yaml

    ...

The example now runs on Kubernetes API and successfully queries the number of
pods in the cluster every 10 seconds.

### Clean up

To stop this example and clean up the pod, press <kbd>Ctrl</kbd>+<kbd>C</kbd> on
the `kubectl run` command and then run:

    kubectl -n flink-operator delete deployment flink-kubernetes-dashboard
