/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/galasa-dev/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGalasaApiComponents implements GalasaApiComponentInterface
type FakeGalasaApiComponents struct {
	Fake *FakeGalasaV2alpha1
	ns   string
}

var galasaapicomponentsResource = schema.GroupVersionResource{Group: "galasa.dev", Version: "v2alpha1", Resource: "galasaapicomponents"}

var galasaapicomponentsKind = schema.GroupVersionKind{Group: "galasa.dev", Version: "v2alpha1", Kind: "GalasaApiComponent"}

// Get takes name of the galasaApiComponent, and returns the corresponding galasaApiComponent object, and an error if there is any.
func (c *FakeGalasaApiComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaApiComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(galasaapicomponentsResource, c.ns, name), &v2alpha1.GalasaApiComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaApiComponent), err
}

// List takes label and field selectors, and returns the list of GalasaApiComponents that match those selectors.
func (c *FakeGalasaApiComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaApiComponentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(galasaapicomponentsResource, galasaapicomponentsKind, c.ns, opts), &v2alpha1.GalasaApiComponentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.GalasaApiComponentList{ListMeta: obj.(*v2alpha1.GalasaApiComponentList).ListMeta}
	for _, item := range obj.(*v2alpha1.GalasaApiComponentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested galasaApiComponents.
func (c *FakeGalasaApiComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(galasaapicomponentsResource, c.ns, opts))

}

// Create takes the representation of a galasaApiComponent and creates it.  Returns the server's representation of the galasaApiComponent, and an error, if there is any.
func (c *FakeGalasaApiComponents) Create(ctx context.Context, galasaApiComponent *v2alpha1.GalasaApiComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaApiComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(galasaapicomponentsResource, c.ns, galasaApiComponent), &v2alpha1.GalasaApiComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaApiComponent), err
}

// Update takes the representation of a galasaApiComponent and updates it. Returns the server's representation of the galasaApiComponent, and an error, if there is any.
func (c *FakeGalasaApiComponents) Update(ctx context.Context, galasaApiComponent *v2alpha1.GalasaApiComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaApiComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(galasaapicomponentsResource, c.ns, galasaApiComponent), &v2alpha1.GalasaApiComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaApiComponent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGalasaApiComponents) UpdateStatus(ctx context.Context, galasaApiComponent *v2alpha1.GalasaApiComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaApiComponent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(galasaapicomponentsResource, "status", c.ns, galasaApiComponent), &v2alpha1.GalasaApiComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaApiComponent), err
}

// Delete takes name of the galasaApiComponent and deletes it. Returns an error if one occurs.
func (c *FakeGalasaApiComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(galasaapicomponentsResource, c.ns, name), &v2alpha1.GalasaApiComponent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGalasaApiComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(galasaapicomponentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.GalasaApiComponentList{})
	return err
}

// Patch applies the patch and returns the patched galasaApiComponent.
func (c *FakeGalasaApiComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaApiComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(galasaapicomponentsResource, c.ns, name, pt, data, subresources...), &v2alpha1.GalasaApiComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaApiComponent), err
}
