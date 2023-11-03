# 🛠️ Build a Kubernetes Operator in 10 minutes

> **👋 The source code has been updated in early November 2023 to use the latest version of kubebuilder ([v3.13.0](https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v3.13.0)). Expect the code to be kept up to date with the latest kubebuilder releases!**

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
diff --color -r operator-v1/api/v1/foo_types.go operator-v2/api/v1/foo_types.go
33a34,36
>
> 	// Foo's favorite colour
> 	Colour string `json:"colour,omitempty"`
Only in operator-v2/bin: k8s
Only in operator-v1/bin: kustomize
Binary files operator-v1/bin/manager and operator-v2/bin/manager differ
Only in operator-v2/bin: setup-envtest
diff --color -r operator-v1/config/crd/bases/tutorial.my.domain_foos.yaml operator-v2/config/crd/bases/tutorial.my.domain_foos.yaml
45a46,48
>               colour:
>                 description: Foo's favorite colour
>                 type: string
Only in operator-v2: cover.out
Only in operator-v2/internal: color
diff --color -r operator-v1/internal/controller/foo_controller.go operator-v2/internal/controller/foo_controller.go
31a32
> 	"my.domain/tutorial/internal/color"
76a78
> 	foo.Status.Colour = color.ConvertStrToColor(foo.Name + foo.Namespace)
diff --color -r operator-v1/internal/controller/suite_test.go operator-v2/internal/controller/suite_test.go
66c66
< 			fmt.Sprintf("1.28.3-%s-%s", runtime.GOOS, runtime.GOARCH)),
---
> 			fmt.Sprintf("1.28.0-%s-%s", runtime.GOOS, runtime.GOARCH)),
```

### `v2` <> `v2-with-tests`

```diff
$ diff -r operator-v2 operator-v2-with-tests
Binary files operator-v2/bin/manager and operator-v2-with-tests/bin/manager differ
Only in operator-v2-with-tests/internal/color: color_test.go
Only in operator-v2-with-tests/internal/controller: foo_controller_test.go
diff --color -r operator-v2/internal/controller/suite_test.go operator-v2-with-tests/internal/controller/suite_test.go
19a20
> 	"context"
24a26,27
> 	ctrl "sigs.k8s.io/controller-runtime"
>
42,44c45,51
< var cfg *rest.Config
< var k8sClient client.Client
< var testEnv *envtest.Environment
---
> var (
> 	cfg       *rest.Config
> 	k8sClient client.Client
> 	testEnv   *envtest.Environment
> 	ctx       context.Context
> 	cancel    context.CancelFunc
> )
53a61
> 	ctx, cancel = context.WithCancel(context.TODO())
66c74
< 			fmt.Sprintf("1.28.0-%s-%s", runtime.GOOS, runtime.GOARCH)),
---
> 			fmt.Sprintf("1.28.3-%s-%s", runtime.GOOS, runtime.GOARCH)),
83a92,108
> 	// Register and start the Foo controller
> 	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
> 		Scheme: scheme.Scheme,
> 	})
> 	Expect(err).ToNot(HaveOccurred())
>
> 	err = (&FooReconciler{
> 		Client: k8sManager.GetClient(),
> 		Scheme: k8sManager.GetScheme(),
> 	}).SetupWithManager(k8sManager)
> 	Expect(err).ToNot(HaveOccurred())
>
> 	go func() {
> 		defer GinkgoRecover()
> 		err = k8sManager.Start(ctx)
> 		Expect(err).ToNot(HaveOccurred(), "failed to run manager")
> 	}()
86a112
> 	cancel()
```
