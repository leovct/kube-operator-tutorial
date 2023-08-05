# 🛠️ Build a Kubernetes Operator in 10 minutes

## Table of Contents

- [Introduction](#introduction)
- [Architecture Diagram](#architecture-diagram)
- [Differences between versions](#differences-between-versions)

## Introduction

This repository serves as a valuable reference for all tutorial followers aspiring to master the art of building Kubernetes Operators. Here, you will find the complete source code and resources used in the [articles](https://medium.com/@leovct/list/kubernetes-operators-101-dcfcc4cb52f6).

Below, you'll find the mapping of each tutorial article to its corresponding code directory.

Directory | Purpose | Article
------ | ------- | -------
[`operator-v1`](operator-v1/README.md) | First version of the Kubernetes operator | [Build a Kubernetes Operator in 10 minutes](https://medium.com/better-programming/build-a-kubernetes-operator-in-10-minutes-11eec1492d30)
[`operator-v2`](operator-v2/README.md) | Second version of the Kubernetes operator with color status | [How to Write Tests for your Kubernetes Operator](https://betterprogramming.pub/write-tests-for-your-kubernetes-operator-d3d6a9530840)
[`operator-v2-with-tests`](operator-v2-with-tests/README.md) | Second version of the Kubernetes operator with unit and integration tests | [How to Write Tests for your Kubernetes Operator](https://betterprogramming.pub/write-tests-for-your-kubernetes-operator-d3d6a9530840)

Happy coding and learning! 🚀

## Architecture Diagram

Here's the architecture diagram of the `Foo` operator that you'll design following the articles.

Note that it's a very simple operator which has no real use, except to demonstrate the capabilities of an operator.

<p><img src="doc/overview.png" alt="operator-overview" width="700px"/></p>

## Differences between versions

Below are examples of `diff` outputs between different versions of the operator.

### `v1` <> `v2`

```diff
$ diff -r operator-v1 operator-v2
diff --color -r operator-v1/README.md operator-v2/README.md
1c1
< # operator-v1
---
> # operator-v2
diff --color -r operator-v1/api/v1/foo_types.go operator-v2/api/v1/foo_types.go
32a33,35
> 
>       // Foo's favorite colour
>       Colour string `json:"colour,omitempty"`
Only in operator-v2: color
diff --color -r operator-v1/config/crd/bases/tutorial.my.domain_foos.yaml operator-v2/config/crd/bases/tutorial.my.domain_foos.yaml
46a47,49
>               colour:
>                 description: Foo's favorite colour
>                 type: string
diff --color -r operator-v1/controllers/foo_controller.go operator-v2/controllers/foo_controller.go
32a33
>       color "my.domain/tutorial/color"
77a79
>       foo.Status.Colour = color.ConvertStrToColor(foo.Name + foo.Namespace)
```

### `v2` <> `v2-with-tests`

```diff
$ diff -r operator-v2 operator-v2-with-tests
diff --color -r operator-v2/README.md operator-v2-with-tests/README.md
1c1
< # operator-v2
---
> # operator-v2-with-tests
Only in operator-v2-with-tests/color: color_test.go
Only in operator-v2-with-tests/controllers: foo_controller_test.go
diff --color -r operator-v2/controllers/suite_test.go operator-v2-with-tests/controllers/suite_test.go
19a20
>       "context"
22a24,25
>       ctrl "sigs.k8s.io/controller-runtime"
> 
40,42c43,49
< var cfg *rest.Config
< var k8sClient client.Client
< var testEnv *envtest.Environment
---
> var (
>       cfg       *rest.Config
>       k8sClient client.Client
>       testEnv   *envtest.Environment
>       ctx       context.Context
>       cancel    context.CancelFunc
> )
53a61
>       ctx, cancel = context.WithCancel(context.TODO())
75a84,100
>       // Register and start the Foo controller
>       k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
>               Scheme: scheme.Scheme,
>       })
>       Expect(err).ToNot(HaveOccurred())
> 
>       err = (&FooReconciler{
>               Client: k8sManager.GetClient(),
>               Scheme: k8sManager.GetScheme(),
>       }).SetupWithManager(k8sManager)
>       Expect(err).ToNot(HaveOccurred())
> 
>       go func() {
>               defer GinkgoRecover()
>               err = k8sManager.Start(ctx)
>               Expect(err).ToNot(HaveOccurred(), "failed to run manager")
>       }()
78a104
>       cancel()
```
