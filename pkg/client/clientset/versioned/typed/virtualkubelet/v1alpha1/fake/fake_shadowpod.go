// Copyright 2019-2021 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/liqotech/liqo/apis/virtualkubelet/v1alpha1"
)

// FakeShadowPods implements ShadowPodInterface
type FakeShadowPods struct {
	Fake *FakeVirtualkubeletV1alpha1
	ns   string
}

var shadowpodsResource = schema.GroupVersionResource{Group: "virtualkubelet", Version: "v1alpha1", Resource: "shadowpods"}

var shadowpodsKind = schema.GroupVersionKind{Group: "virtualkubelet", Version: "v1alpha1", Kind: "ShadowPod"}

// Get takes name of the shadowPod, and returns the corresponding shadowPod object, and an error if there is any.
func (c *FakeShadowPods) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShadowPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shadowpodsResource, c.ns, name), &v1alpha1.ShadowPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowPod), err
}

// List takes label and field selectors, and returns the list of ShadowPods that match those selectors.
func (c *FakeShadowPods) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShadowPodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shadowpodsResource, shadowpodsKind, c.ns, opts), &v1alpha1.ShadowPodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShadowPodList{ListMeta: obj.(*v1alpha1.ShadowPodList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShadowPodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shadowPods.
func (c *FakeShadowPods) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shadowpodsResource, c.ns, opts))

}

// Create takes the representation of a shadowPod and creates it.  Returns the server's representation of the shadowPod, and an error, if there is any.
func (c *FakeShadowPods) Create(ctx context.Context, shadowPod *v1alpha1.ShadowPod, opts v1.CreateOptions) (result *v1alpha1.ShadowPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shadowpodsResource, c.ns, shadowPod), &v1alpha1.ShadowPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowPod), err
}

// Update takes the representation of a shadowPod and updates it. Returns the server's representation of the shadowPod, and an error, if there is any.
func (c *FakeShadowPods) Update(ctx context.Context, shadowPod *v1alpha1.ShadowPod, opts v1.UpdateOptions) (result *v1alpha1.ShadowPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shadowpodsResource, c.ns, shadowPod), &v1alpha1.ShadowPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowPod), err
}

// Delete takes name of the shadowPod and deletes it. Returns an error if one occurs.
func (c *FakeShadowPods) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shadowpodsResource, c.ns, name), &v1alpha1.ShadowPod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShadowPods) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shadowpodsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShadowPodList{})
	return err
}

// Patch applies the patch and returns the patched shadowPod.
func (c *FakeShadowPods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShadowPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shadowpodsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShadowPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowPod), err
}
