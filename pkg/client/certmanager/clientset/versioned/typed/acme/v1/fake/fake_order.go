/*
Copyright 2022 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOrders implements OrderInterface
type FakeOrders struct {
	Fake *FakeAcmeV1
	ns   string
}

var ordersResource = v1.SchemeGroupVersion.WithResource("orders")

var ordersKind = v1.SchemeGroupVersion.WithKind("Order")

// Get takes name of the order, and returns the corresponding order object, and an error if there is any.
func (c *FakeOrders) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Order, err error) {
	emptyResult := &v1.Order{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(ordersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Order), err
}

// List takes label and field selectors, and returns the list of Orders that match those selectors.
func (c *FakeOrders) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OrderList, err error) {
	emptyResult := &v1.OrderList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(ordersResource, ordersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OrderList{ListMeta: obj.(*v1.OrderList).ListMeta}
	for _, item := range obj.(*v1.OrderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested orders.
func (c *FakeOrders) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(ordersResource, c.ns, opts))

}

// Create takes the representation of a order and creates it.  Returns the server's representation of the order, and an error, if there is any.
func (c *FakeOrders) Create(ctx context.Context, order *v1.Order, opts metav1.CreateOptions) (result *v1.Order, err error) {
	emptyResult := &v1.Order{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(ordersResource, c.ns, order, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Order), err
}

// Update takes the representation of a order and updates it. Returns the server's representation of the order, and an error, if there is any.
func (c *FakeOrders) Update(ctx context.Context, order *v1.Order, opts metav1.UpdateOptions) (result *v1.Order, err error) {
	emptyResult := &v1.Order{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(ordersResource, c.ns, order, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Order), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrders) UpdateStatus(ctx context.Context, order *v1.Order, opts metav1.UpdateOptions) (result *v1.Order, err error) {
	emptyResult := &v1.Order{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(ordersResource, "status", c.ns, order, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Order), err
}

// Delete takes name of the order and deletes it. Returns an error if one occurs.
func (c *FakeOrders) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ordersResource, c.ns, name, opts), &v1.Order{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrders) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(ordersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OrderList{})
	return err
}

// Patch applies the patch and returns the patched order.
func (c *FakeOrders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Order, err error) {
	emptyResult := &v1.Order{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(ordersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Order), err
}
