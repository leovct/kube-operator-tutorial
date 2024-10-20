# 🛠️ Build a Kubernetes Operator in 10 minutes

> **👋 The source code has been updated in May 2024 to use the latest version of kubebuilder ([v3.15.0](https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v3.15.0)). Expect the code to be kept up to date with the latest kubebuilder releases!**

## Table of Contents

- [Introduction](#introduction)
- [Architecture Diagram](#architecture-diagram)
- [Differences between versions](#differences-between-versions)
- [Contributing](#contributing)

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
$ diff --exclude=bin --exclude=README.md -r operator-v1 operator-v2
diff --color --exclude=bin --exclude=README.md -r operator-v1/api/v1/foo_types.go operator-v2/api/v1/foo_types.go
33a34,36
>
> 	// Foo's favorite colour
> 	Colour string `json:"colour,omitempty"`
diff --color --exclude=bin --exclude=README.md -r operator-v1/config/crd/bases/tutorial.my.domain_foos.yaml operator-v2/config/crd/bases/tutorial.my.domain_foos.yaml
50a51,53
>               colour:
>                 description: Foo's favorite colour
>                 type: string
Only in operator-v2/internal: color
diff --color --exclude=bin --exclude=README.md -r operator-v1/internal/controller/foo_controller.go operator-v2/internal/controller/foo_controller.go
31a32
> 	"my.domain/tutorial/internal/color"
76a78
> 	foo.Status.Colour = color.ConvertStrToColor(foo.Name + foo.Namespace)
```

### `v2` <> `v2-with-tests`

```diff
$ diff --exclude=bin --exclude=README.md -r operator-v2 operator-v2-with-tests
Only in operator-v2-with-tests/internal/color: color_test.go
diff --color --exclude=bin --exclude=README.md -r operator-v2/internal/controller/foo_controller_test.go operator-v2-with-tests/internal/controller/foo_controller_test.go
24,26d23
< 	"k8s.io/apimachinery/pkg/api/errors"
< 	"k8s.io/apimachinery/pkg/types"
< 	"sigs.k8s.io/controller-runtime/pkg/reconcile"
27a25
> 	corev1 "k8s.io/api/core/v1"
29c27
<
---
> 	"k8s.io/apimachinery/pkg/types"
33,35c31
< var _ = Describe("Foo Controller", func() {
< 	Context("When reconciling a resource", func() {
< 		const resourceName = "test-resource"
---
> var _ = Describe("Foo controller", func() {
37c33,35
< 		ctx := context.Background()
---
> 	const (
> 		foo1Name   = "foo-1"
> 		foo1Friend = "jack"
39,43c37,38
< 		typeNamespacedName := types.NamespacedName{
< 			Name:      resourceName,
< 			Namespace: "default", // TODO(user):Modify as needed
< 		}
< 		foo := &tutorialv1.Foo{}
---
> 		foo2Name   = "foo-2"
> 		foo2Friend = "joe"
45,52c40,88
< 		BeforeEach(func() {
< 			By("creating the custom resource for the Kind Foo")
< 			err := k8sClient.Get(ctx, typeNamespacedName, foo)
< 			if err != nil && errors.IsNotFound(err) {
< 				resource := &tutorialv1.Foo{
< 					ObjectMeta: metav1.ObjectMeta{
< 						Name:      resourceName,
< 						Namespace: "default",
---
> 		namespace = "default"
> 	)
>
> 	Context("When setting up the test environment", func() {
> 		It("Should create Foo custom resources", func() {
> 			By("Creating a first Foo custom resource")
> 			ctx := context.Background()
> 			foo1 := tutorialv1.Foo{
> 				ObjectMeta: metav1.ObjectMeta{
> 					Name:      foo1Name,
> 					Namespace: namespace,
> 				},
> 				Spec: tutorialv1.FooSpec{
> 					Name: foo1Friend,
> 				},
> 			}
> 			Expect(k8sClient.Create(ctx, &foo1)).Should(Succeed())
>
> 			By("Creating another Foo custom resource")
> 			foo2 := tutorialv1.Foo{
> 				ObjectMeta: metav1.ObjectMeta{
> 					Name:      foo2Name,
> 					Namespace: namespace,
> 				},
> 				Spec: tutorialv1.FooSpec{
> 					Name: foo2Friend,
> 				},
> 			}
> 			Expect(k8sClient.Create(ctx, &foo2)).Should(Succeed())
> 		})
> 	})
>
> 	Context("When creating a pod with the same name as one of the Foo custom resources' friends", func() {
> 		It("Should update the status of the first Foo custom resource", func() {
> 			By("Creating the pod")
> 			ctx := context.Background()
> 			pod := corev1.Pod{
> 				ObjectMeta: metav1.ObjectMeta{
> 					Name:      foo1Friend,
> 					Namespace: namespace,
> 				},
> 				Spec: corev1.PodSpec{
> 					Containers: []corev1.Container{
> 						{
> 							Name:    "ubuntu",
> 							Image:   "ubuntu:latest",
> 							Command: []string{"sleep"},
> 							Args:    []string{"infinity"},
> 						},
54c90,102
< 					// TODO(user): Specify other spec details if needed.
---
> 				},
> 			}
> 			Expect(k8sClient.Create(ctx, &pod)).Should(Succeed())
>
> 			By("Updating the status of the first Foo custom resource")
> 			var foo1 tutorialv1.Foo
> 			foo1Request := types.NamespacedName{
> 				Name:      foo1Name,
> 				Namespace: namespace,
> 			}
> 			Eventually(func() bool {
> 				if err := k8sClient.Get(ctx, foo1Request, &foo1); err != nil {
> 					return false
56c104,111
< 				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
---
> 				return foo1.Status.Happy
> 			}).Should(BeTrue())
>
> 			By("Not updating the status of the other Foo custom resource")
> 			var foo2 tutorialv1.Foo
> 			foo2Request := types.NamespacedName{
> 				Name:      foo2Name,
> 				Namespace: namespace,
57a113,118
> 			Consistently(func() bool {
> 				if err := k8sClient.Get(ctx, foo2Request, &foo2); err != nil {
> 					return false
> 				}
> 				return foo2.Status.Happy
> 			}).Should(BeFalse())
58a120
> 	})
60,64c122,131
< 		AfterEach(func() {
< 			// TODO(user): Cleanup logic after each test, like removing the resource instance.
< 			resource := &tutorialv1.Foo{}
< 			err := k8sClient.Get(ctx, typeNamespacedName, resource)
< 			Expect(err).NotTo(HaveOccurred())
---
> 	Context("When updating the name of a Foo custom resource's friend", func() {
> 		It("Should update the status of the Foo custom resource", func() {
> 			By("Getting the second Foo custom resource")
> 			ctx := context.Background()
> 			var foo2 tutorialv1.Foo
> 			foo2Request := types.NamespacedName{
> 				Name:      foo2Name,
> 				Namespace: namespace,
> 			}
> 			Expect(k8sClient.Get(ctx, foo2Request, &foo2)).To(Succeed())
66,67c133,156
< 			By("Cleanup the specific resource instance Foo")
< 			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
---
> 			By("Updating the name of a Foo custom resource's friend")
> 			foo2.Spec.Name = foo1Friend
> 			Expect(k8sClient.Update(ctx, &foo2)).To(Succeed())
>
> 			By("Updating the status of the other Foo custom resource")
> 			Eventually(func() bool {
> 				if err := k8sClient.Get(ctx, foo2Request, &foo2); err != nil {
> 					return false
> 				}
> 				return foo2.Status.Happy
> 			}).Should(BeTrue())
>
> 			By("Not updating the status of the first Foo custom resource")
> 			var foo1 tutorialv1.Foo
> 			foo1Request := types.NamespacedName{
> 				Name:      foo1Name,
> 				Namespace: namespace,
> 			}
> 			Consistently(func() bool {
> 				if err := k8sClient.Get(ctx, foo1Request, &foo1); err != nil {
> 					return false
> 				}
> 				return foo1.Status.Happy
> 			}).Should(BeTrue())
69,73c158,168
< 		It("should successfully reconcile the resource", func() {
< 			By("Reconciling the created resource")
< 			controllerReconciler := &FooReconciler{
< 				Client: k8sClient,
< 				Scheme: k8sClient.Scheme(),
---
> 	})
>
> 	Context("When deleting a pod with the same name as one of the Foo custom resourcess' friends", func() {
> 		It("Should update the status of the first Foo custom resource", func() {
> 			By("Deleting the pod")
> 			ctx := context.Background()
> 			pod := corev1.Pod{
> 				ObjectMeta: metav1.ObjectMeta{
> 					Name:      foo1Friend,
> 					Namespace: namespace,
> 				},
74a170
> 			Expect(k8sClient.Delete(ctx, &pod)).Should(Succeed())
76,81c172,196
< 			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
< 				NamespacedName: typeNamespacedName,
< 			})
< 			Expect(err).NotTo(HaveOccurred())
< 			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
< 			// Example: If you expect a certain status condition after reconciliation, verify it here.
---
> 			By("Updating the status of the first Foo custom resource")
> 			var foo1 tutorialv1.Foo
> 			foo1Request := types.NamespacedName{
> 				Name:      foo1Name,
> 				Namespace: namespace,
> 			}
> 			Eventually(func() bool {
> 				if err := k8sClient.Get(ctx, foo1Request, &foo1); err != nil {
> 					return false
> 				}
> 				return foo1.Status.Happy
> 			}).Should(BeFalse())
>
> 			By("Updating the status of the other Foo custom resource")
> 			var foo2 tutorialv1.Foo
> 			foo2Request := types.NamespacedName{
> 				Name:      foo2Name,
> 				Namespace: namespace,
> 			}
> 			Consistently(func() bool {
> 				if err := k8sClient.Get(ctx, foo2Request, &foo2); err != nil {
> 					return false
> 				}
> 				return foo2.Status.Happy
> 			}).Should(BeFalse())
diff --color --exclude=bin --exclude=README.md -r operator-v2/internal/controller/suite_test.go operator-v2-with-tests/internal/controller/suite_test.go
19a20
> 	"context"
24a26,27
> 	ctrl "sigs.k8s.io/controller-runtime"
>
44a48,49
> var ctx context.Context
> var cancel context.CancelFunc
53a59
> 	ctx, cancel = context.WithCancel(context.TODO())
83a90,106
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
86a110
> 	cancel()
```

## Contributing

Contributions are welcome! Feel free to open issues or reach out if you want more details! :)

### Bump kubebuilder version

Simple steps to follow to upgrade the tutorial to the latest `kubebuilder` version.

Note: this is an example with `operator-v1`. Repeat the same steps for all the other versions of the operator...

```bash
# 1) Scaffold the projects.
./scripts/bump.sh operator-v1
./scripts/bump.sh operator-v2
./scripts/bump.sh operator-v2-with-tests

# 2) Test that the new version works (for each folder: operator-v1, operator-v2 and operator-v2-with-tests).
# Note: for this step, you will need a running Kubernetes cluster.
make test

kind create cluster
kubectl cluster-info --context kind-kind
kubectl get nodes

make install
kubectl get crds
make run

kubectl apply -k config/samples
# Check the logs of the controller, it should detect the creation events.
# Also check the status of the CRDs, it should be empty at this point.
kubectl describe foos

kubectl apply -f config/samples/pod.yaml
# Again, check the logs of the controller, it should throw some logs.
# The foo-1 CRD should now have an happy status.
kubectl describe foos

# Update the pod name from `jack` to `joe`.
sed -i '' "s/jack/joe/" config/samples/pod.yaml
kubectl apply -f config/samples/pod.yaml
# Both CRDs should now have an happy status.
kubectl describe foos
kubectl delete pod jack --force
# Only the foo-2 CRD should have an empty status.
kubectl describe foos

# Once you're done, clean up the environment.
kind delete cluster --name kind

# 3) Compare the diffs between the new and the old projects.
# Also make sure to compare diffs between projects and keep the `README` updated!

# 4) Release a new tag!

# 5) Update the website articles and Medium articles too!
# - https://leovct.github.io/
# - https://medium.com/@leovct/list/kubernetes-operators-101-dcfcc4cb52f6
```
