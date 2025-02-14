/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/galasa-dev/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	scheme "github.com/galasa-dev/galasa-kubernetes-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GalasaEngineControllerComponentsGetter has a method to return a GalasaEngineControllerComponentInterface.
// A group's client should implement this interface.
type GalasaEngineControllerComponentsGetter interface {
	GalasaEngineControllerComponents(namespace string) GalasaEngineControllerComponentInterface
}

// GalasaEngineControllerComponentInterface has methods to work with GalasaEngineControllerComponent resources.
type GalasaEngineControllerComponentInterface interface {
	Create(ctx context.Context, galasaEngineControllerComponent *v2alpha1.GalasaEngineControllerComponent, opts v1.CreateOptions) (*v2alpha1.GalasaEngineControllerComponent, error)
	Update(ctx context.Context, galasaEngineControllerComponent *v2alpha1.GalasaEngineControllerComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaEngineControllerComponent, error)
	UpdateStatus(ctx context.Context, galasaEngineControllerComponent *v2alpha1.GalasaEngineControllerComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaEngineControllerComponent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaEngineControllerComponent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaEngineControllerComponentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaEngineControllerComponent, err error)
	GalasaEngineControllerComponentExpansion
}

// galasaEngineControllerComponents implements GalasaEngineControllerComponentInterface
type galasaEngineControllerComponents struct {
	client rest.Interface
	ns     string
}

// newGalasaEngineControllerComponents returns a GalasaEngineControllerComponents
func newGalasaEngineControllerComponents(c *GalasaV2alpha1Client, namespace string) *galasaEngineControllerComponents {
	return &galasaEngineControllerComponents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the galasaEngineControllerComponent, and returns the corresponding galasaEngineControllerComponent object, and an error if there is any.
func (c *galasaEngineControllerComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaEngineControllerComponent, err error) {
	result = &v2alpha1.GalasaEngineControllerComponent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GalasaEngineControllerComponents that match those selectors.
func (c *galasaEngineControllerComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaEngineControllerComponentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.GalasaEngineControllerComponentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested galasaEngineControllerComponents.
func (c *galasaEngineControllerComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a galasaEngineControllerComponent and creates it.  Returns the server's representation of the galasaEngineControllerComponent, and an error, if there is any.
func (c *galasaEngineControllerComponents) Create(ctx context.Context, galasaEngineControllerComponent *v2alpha1.GalasaEngineControllerComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaEngineControllerComponent, err error) {
	result = &v2alpha1.GalasaEngineControllerComponent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaEngineControllerComponent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a galasaEngineControllerComponent and updates it. Returns the server's representation of the galasaEngineControllerComponent, and an error, if there is any.
func (c *galasaEngineControllerComponents) Update(ctx context.Context, galasaEngineControllerComponent *v2alpha1.GalasaEngineControllerComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaEngineControllerComponent, err error) {
	result = &v2alpha1.GalasaEngineControllerComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		Name(galasaEngineControllerComponent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaEngineControllerComponent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *galasaEngineControllerComponents) UpdateStatus(ctx context.Context, galasaEngineControllerComponent *v2alpha1.GalasaEngineControllerComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaEngineControllerComponent, err error) {
	result = &v2alpha1.GalasaEngineControllerComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		Name(galasaEngineControllerComponent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaEngineControllerComponent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the galasaEngineControllerComponent and deletes it. Returns an error if one occurs.
func (c *galasaEngineControllerComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *galasaEngineControllerComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched galasaEngineControllerComponent.
func (c *galasaEngineControllerComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaEngineControllerComponent, err error) {
	result = &v2alpha1.GalasaEngineControllerComponent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("galasaenginecontrollercomponents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
