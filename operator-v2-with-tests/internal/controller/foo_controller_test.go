/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	tutorialv1 "my.domain/tutorial/api/v1"
)

var _ = Describe("Foo controller", func() {

	const (
		foo1Name   = "foo-1"
		foo1Friend = "jack"

		foo2Name   = "foo-2"
		foo2Friend = "joe"

		namespace = "default"
	)

	Context("When setting up the test environment", func() {
		It("Should create Foo custom resources", func() {
			By("Creating a first Foo custom resource")
			ctx := context.Background()
			foo1 := tutorialv1.Foo{
				ObjectMeta: metav1.ObjectMeta{
					Name:      foo1Name,
					Namespace: namespace,
				},
				Spec: tutorialv1.FooSpec{
					Name: foo1Friend,
				},
			}
			Expect(k8sClient.Create(ctx, &foo1)).Should(Succeed())

			By("Creating another Foo custom resource")
			foo2 := tutorialv1.Foo{
				ObjectMeta: metav1.ObjectMeta{
					Name:      foo2Name,
					Namespace: namespace,
				},
				Spec: tutorialv1.FooSpec{
					Name: foo2Friend,
				},
			}
			Expect(k8sClient.Create(ctx, &foo2)).Should(Succeed())
		})
	})

	Context("When creating a pod with the same name as one of the Foo custom resources' friends", func() {
		It("Should update the status of the first Foo custom resource", func() {
			By("Creating the pod")
			ctx := context.Background()
			pod := corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      foo1Friend,
					Namespace: namespace,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "ubuntu",
							Image:   "ubuntu:latest",
							Command: []string{"sleep"},
							Args:    []string{"infinity"},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, &pod)).Should(Succeed())

			By("Updating the status of the first Foo custom resource")
			var foo1 tutorialv1.Foo
			foo1Request := types.NamespacedName{
				Name:      foo1Name,
				Namespace: namespace,
			}
			Eventually(func() bool {
				if err := k8sClient.Get(ctx, foo1Request, &foo1); err != nil {
					return false
				}
				return foo1.Status.Happy
			}).Should(BeTrue())

			By("Not updating the status of the other Foo custom resource")
			var foo2 tutorialv1.Foo
			foo2Request := types.NamespacedName{
				Name:      foo2Name,
				Namespace: namespace,
			}
			Consistently(func() bool {
				if err := k8sClient.Get(ctx, foo2Request, &foo2); err != nil {
					return false
				}
				return foo2.Status.Happy
			}).Should(BeFalse())
		})
	})

	Context("When updating the name of a Foo custom resource's friend", func() {
		It("Should update the status of the Foo custom resource", func() {
			By("Getting the second Foo custom resource")
			ctx := context.Background()
			var foo2 tutorialv1.Foo
			foo2Request := types.NamespacedName{
				Name:      foo2Name,
				Namespace: namespace,
			}
			Expect(k8sClient.Get(ctx, foo2Request, &foo2)).To(Succeed())

			By("Updating the name of a Foo custom resource's friend")
			foo2.Spec.Name = foo1Friend
			Expect(k8sClient.Update(ctx, &foo2)).To(Succeed())

			By("Updating the status of the other Foo custom resource")
			Eventually(func() bool {
				if err := k8sClient.Get(ctx, foo2Request, &foo2); err != nil {
					return false
				}
				return foo2.Status.Happy
			}).Should(BeTrue())

			By("Not updating the status of the first Foo custom resource")
			var foo1 tutorialv1.Foo
			foo1Request := types.NamespacedName{
				Name:      foo1Name,
				Namespace: namespace,
			}
			Consistently(func() bool {
				if err := k8sClient.Get(ctx, foo1Request, &foo1); err != nil {
					return false
				}
				return foo1.Status.Happy
			}).Should(BeTrue())
		})
	})

	Context("When deleting a pod with the same name as one of the Foo custom resourcess' friends", func() {
		It("Should update the status of the first Foo custom resource", func() {
			By("Deleting the pod")
			ctx := context.Background()
			pod := corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      foo1Friend,
					Namespace: namespace,
				},
			}
			Expect(k8sClient.Delete(ctx, &pod)).Should(Succeed())

			By("Updating the status of the first Foo custom resource")
			var foo1 tutorialv1.Foo
			foo1Request := types.NamespacedName{
				Name:      foo1Name,
				Namespace: namespace,
			}
			Eventually(func() bool {
				if err := k8sClient.Get(ctx, foo1Request, &foo1); err != nil {
					return false
				}
				return foo1.Status.Happy
			}).Should(BeFalse())

			By("Updating the status of the other Foo custom resource")
			var foo2 tutorialv1.Foo
			foo2Request := types.NamespacedName{
				Name:      foo2Name,
				Namespace: namespace,
			}
			Consistently(func() bool {
				if err := k8sClient.Get(ctx, foo2Request, &foo2); err != nil {
					return false
				}
				return foo2.Status.Happy
			}).Should(BeFalse())
		})
	})
})
