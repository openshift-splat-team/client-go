// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/openshift/api/network/v1alpha1"
	networkv1alpha1 "github.com/openshift/client-go/network/applyconfigurations/network/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSNameResolvers implements DNSNameResolverInterface
type FakeDNSNameResolvers struct {
	Fake *FakeNetworkV1alpha1
	ns   string
}

var dnsnameresolversResource = v1alpha1.SchemeGroupVersion.WithResource("dnsnameresolvers")

var dnsnameresolversKind = v1alpha1.SchemeGroupVersion.WithKind("DNSNameResolver")

// Get takes name of the dNSNameResolver, and returns the corresponding dNSNameResolver object, and an error if there is any.
func (c *FakeDNSNameResolvers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DNSNameResolver, err error) {
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(dnsnameresolversResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}

// List takes label and field selectors, and returns the list of DNSNameResolvers that match those selectors.
func (c *FakeDNSNameResolvers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DNSNameResolverList, err error) {
	emptyResult := &v1alpha1.DNSNameResolverList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(dnsnameresolversResource, dnsnameresolversKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DNSNameResolverList{ListMeta: obj.(*v1alpha1.DNSNameResolverList).ListMeta}
	for _, item := range obj.(*v1alpha1.DNSNameResolverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSNameResolvers.
func (c *FakeDNSNameResolvers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(dnsnameresolversResource, c.ns, opts))

}

// Create takes the representation of a dNSNameResolver and creates it.  Returns the server's representation of the dNSNameResolver, and an error, if there is any.
func (c *FakeDNSNameResolvers) Create(ctx context.Context, dNSNameResolver *v1alpha1.DNSNameResolver, opts v1.CreateOptions) (result *v1alpha1.DNSNameResolver, err error) {
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(dnsnameresolversResource, c.ns, dNSNameResolver, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}

// Update takes the representation of a dNSNameResolver and updates it. Returns the server's representation of the dNSNameResolver, and an error, if there is any.
func (c *FakeDNSNameResolvers) Update(ctx context.Context, dNSNameResolver *v1alpha1.DNSNameResolver, opts v1.UpdateOptions) (result *v1alpha1.DNSNameResolver, err error) {
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(dnsnameresolversResource, c.ns, dNSNameResolver, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSNameResolvers) UpdateStatus(ctx context.Context, dNSNameResolver *v1alpha1.DNSNameResolver, opts v1.UpdateOptions) (result *v1alpha1.DNSNameResolver, err error) {
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(dnsnameresolversResource, "status", c.ns, dNSNameResolver, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}

// Delete takes name of the dNSNameResolver and deletes it. Returns an error if one occurs.
func (c *FakeDNSNameResolvers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dnsnameresolversResource, c.ns, name, opts), &v1alpha1.DNSNameResolver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSNameResolvers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(dnsnameresolversResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DNSNameResolverList{})
	return err
}

// Patch applies the patch and returns the patched dNSNameResolver.
func (c *FakeDNSNameResolvers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DNSNameResolver, err error) {
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(dnsnameresolversResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied dNSNameResolver.
func (c *FakeDNSNameResolvers) Apply(ctx context.Context, dNSNameResolver *networkv1alpha1.DNSNameResolverApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DNSNameResolver, err error) {
	if dNSNameResolver == nil {
		return nil, fmt.Errorf("dNSNameResolver provided to Apply must not be nil")
	}
	data, err := json.Marshal(dNSNameResolver)
	if err != nil {
		return nil, err
	}
	name := dNSNameResolver.Name
	if name == nil {
		return nil, fmt.Errorf("dNSNameResolver.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(dnsnameresolversResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDNSNameResolvers) ApplyStatus(ctx context.Context, dNSNameResolver *networkv1alpha1.DNSNameResolverApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DNSNameResolver, err error) {
	if dNSNameResolver == nil {
		return nil, fmt.Errorf("dNSNameResolver provided to Apply must not be nil")
	}
	data, err := json.Marshal(dNSNameResolver)
	if err != nil {
		return nil, err
	}
	name := dNSNameResolver.Name
	if name == nil {
		return nil, fmt.Errorf("dNSNameResolver.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.DNSNameResolver{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(dnsnameresolversResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DNSNameResolver), err
}
