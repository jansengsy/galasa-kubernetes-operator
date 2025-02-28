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

// FakeGalasaToolboxComponents implements GalasaToolboxComponentInterface
type FakeGalasaToolboxComponents struct {
	Fake *FakeGalasaV2alpha1
	ns   string
}

var galasatoolboxcomponentsResource = schema.GroupVersionResource{Group: "galasa.dev", Version: "v2alpha1", Resource: "galasatoolboxcomponents"}

var galasatoolboxcomponentsKind = schema.GroupVersionKind{Group: "galasa.dev", Version: "v2alpha1", Kind: "GalasaToolboxComponent"}

// Get takes name of the galasaToolboxComponent, and returns the corresponding galasaToolboxComponent object, and an error if there is any.
func (c *FakeGalasaToolboxComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(galasatoolboxcomponentsResource, c.ns, name), &v2alpha1.GalasaToolboxComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaToolboxComponent), err
}

// List takes label and field selectors, and returns the list of GalasaToolboxComponents that match those selectors.
func (c *FakeGalasaToolboxComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaToolboxComponentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(galasatoolboxcomponentsResource, galasatoolboxcomponentsKind, c.ns, opts), &v2alpha1.GalasaToolboxComponentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.GalasaToolboxComponentList{ListMeta: obj.(*v2alpha1.GalasaToolboxComponentList).ListMeta}
	for _, item := range obj.(*v2alpha1.GalasaToolboxComponentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested galasaToolboxComponents.
func (c *FakeGalasaToolboxComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(galasatoolboxcomponentsResource, c.ns, opts))

}

// Create takes the representation of a galasaToolboxComponent and creates it.  Returns the server's representation of the galasaToolboxComponent, and an error, if there is any.
func (c *FakeGalasaToolboxComponents) Create(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(galasatoolboxcomponentsResource, c.ns, galasaToolboxComponent), &v2alpha1.GalasaToolboxComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaToolboxComponent), err
}

// Update takes the representation of a galasaToolboxComponent and updates it. Returns the server's representation of the galasaToolboxComponent, and an error, if there is any.
func (c *FakeGalasaToolboxComponents) Update(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(galasatoolboxcomponentsResource, c.ns, galasaToolboxComponent), &v2alpha1.GalasaToolboxComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaToolboxComponent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGalasaToolboxComponents) UpdateStatus(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaToolboxComponent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(galasatoolboxcomponentsResource, "status", c.ns, galasaToolboxComponent), &v2alpha1.GalasaToolboxComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaToolboxComponent), err
}

// Delete takes name of the galasaToolboxComponent and deletes it. Returns an error if one occurs.
func (c *FakeGalasaToolboxComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(galasatoolboxcomponentsResource, c.ns, name), &v2alpha1.GalasaToolboxComponent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGalasaToolboxComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(galasatoolboxcomponentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.GalasaToolboxComponentList{})
	return err
}

// Patch applies the patch and returns the patched galasaToolboxComponent.
func (c *FakeGalasaToolboxComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaToolboxComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(galasatoolboxcomponentsResource, c.ns, name, pt, data, subresources...), &v2alpha1.GalasaToolboxComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaToolboxComponent), err
}
