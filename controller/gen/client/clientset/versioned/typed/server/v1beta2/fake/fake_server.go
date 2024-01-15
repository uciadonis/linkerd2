/*
Copyright The Kubernetes Authors.

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

	v1beta2 "github.com/linkerd/linkerd2/controller/gen/apis/server/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServers implements ServerInterface
type FakeServers struct {
	Fake *FakeServerV1beta2
	ns   string
}

var serversResource = v1beta2.SchemeGroupVersion.WithResource("servers")

var serversKind = v1beta2.SchemeGroupVersion.WithKind("Server")

// Get takes name of the server, and returns the corresponding server object, and an error if there is any.
func (c *FakeServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.Server, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serversResource, c.ns, name), &v1beta2.Server{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Server), err
}

// List takes label and field selectors, and returns the list of Servers that match those selectors.
func (c *FakeServers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serversResource, serversKind, c.ns, opts), &v1beta2.ServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ServerList{ListMeta: obj.(*v1beta2.ServerList).ListMeta}
	for _, item := range obj.(*v1beta2.ServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servers.
func (c *FakeServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serversResource, c.ns, opts))

}

// Create takes the representation of a server and creates it.  Returns the server's representation of the server, and an error, if there is any.
func (c *FakeServers) Create(ctx context.Context, server *v1beta2.Server, opts v1.CreateOptions) (result *v1beta2.Server, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serversResource, c.ns, server), &v1beta2.Server{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Server), err
}

// Update takes the representation of a server and updates it. Returns the server's representation of the server, and an error, if there is any.
func (c *FakeServers) Update(ctx context.Context, server *v1beta2.Server, opts v1.UpdateOptions) (result *v1beta2.Server, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serversResource, c.ns, server), &v1beta2.Server{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Server), err
}

// Delete takes name of the server and deletes it. Returns an error if one occurs.
func (c *FakeServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(serversResource, c.ns, name, opts), &v1beta2.Server{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.ServerList{})
	return err
}

// Patch applies the patch and returns the patched server.
func (c *FakeServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.Server, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serversResource, c.ns, name, pt, data, subresources...), &v1beta2.Server{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Server), err
}