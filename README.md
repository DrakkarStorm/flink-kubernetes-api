# flink-kubernetes-api

This project is to add action on kubernetes for the interface [flink-kubernetes-dashboard](https://github.com/EnzoDechaene/flink-kubernetes-dashboard)

To have access to manage Flink Operator Object on Kubernetes you will need to apply the rbac.yaml on your kubernetes cluster.

## Getting Started

First, run the development server:

```bash
go build .

./kubernetes-api
```

You can try to acces by requeting http://localhost:8080/deployments

## Use on Kubernetes

If you want to embedded the api and the dashboard you'll need to build and push the docker image to your repository

```bash
make build-and-push
```

The next step is to add the rbac.yaml to have access to your Flink Operator Kubernetes Object.

## Next

Use the Flink Kubernetes Dashboard install to use your api.