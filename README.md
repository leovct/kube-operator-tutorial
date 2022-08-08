# Kubernetes Operator Tutorial
If you want to know more about Kubernetes operators, I invite you to browse [this list of articles](https://medium.com/@leovct/list/kubernetes-operators-101-dcfcc4cb52f6) I have written on the subject.

## Some context on the different branches of the project
Branch | Purpose | Article
------ | ------- | -------
operator-v1 | First version of the Kubernetes operator | [Build a Kubernetes Operator in 10 minutes](https://medium.com/better-programming/build-a-kubernetes-operator-in-10-minutes-11eec1492d30)
operator-v2 | Second version of the Kubernetes operator with color status | [How to Write Tests for your Kubernetes Operator](https://betterprogramming.pub/write-tests-for-your-kubernetes-operator-d3d6a9530840)
tests | Second version of the Kubernetes operator with unit and integration tests | [How to Write Tests for your Kubernetes Operator](https://betterprogramming.pub/write-tests-for-your-kubernetes-operator-d3d6a9530840)

## Overview of the operator's functioning
<p><img src="doc/overview.png" alt="operator-overview" width="700px"/></p>

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/tutorial:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/tutorial:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

