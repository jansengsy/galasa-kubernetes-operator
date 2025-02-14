/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/galasa-dev/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GalasaToolboxComponentLister helps list GalasaToolboxComponents.
// All objects returned here must be treated as read-only.
type GalasaToolboxComponentLister interface {
	// List lists all GalasaToolboxComponents in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.GalasaToolboxComponent, err error)
	// GalasaToolboxComponents returns an object that can list and get GalasaToolboxComponents.
	GalasaToolboxComponents(namespace string) GalasaToolboxComponentNamespaceLister
	GalasaToolboxComponentListerExpansion
}

// galasaToolboxComponentLister implements the GalasaToolboxComponentLister interface.
type galasaToolboxComponentLister struct {
	indexer cache.Indexer
}

// NewGalasaToolboxComponentLister returns a new GalasaToolboxComponentLister.
func NewGalasaToolboxComponentLister(indexer cache.Indexer) GalasaToolboxComponentLister {
	return &galasaToolboxComponentLister{indexer: indexer}
}

// List lists all GalasaToolboxComponents in the indexer.
func (s *galasaToolboxComponentLister) List(selector labels.Selector) (ret []*v2alpha1.GalasaToolboxComponent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.GalasaToolboxComponent))
	})
	return ret, err
}

// GalasaToolboxComponents returns an object that can list and get GalasaToolboxComponents.
func (s *galasaToolboxComponentLister) GalasaToolboxComponents(namespace string) GalasaToolboxComponentNamespaceLister {
	return galasaToolboxComponentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GalasaToolboxComponentNamespaceLister helps list and get GalasaToolboxComponents.
// All objects returned here must be treated as read-only.
type GalasaToolboxComponentNamespaceLister interface {
	// List lists all GalasaToolboxComponents in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.GalasaToolboxComponent, err error)
	// Get retrieves the GalasaToolboxComponent from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.GalasaToolboxComponent, error)
	GalasaToolboxComponentNamespaceListerExpansion
}

// galasaToolboxComponentNamespaceLister implements the GalasaToolboxComponentNamespaceLister
// interface.
type galasaToolboxComponentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GalasaToolboxComponents in the indexer for a given namespace.
func (s galasaToolboxComponentNamespaceLister) List(selector labels.Selector) (ret []*v2alpha1.GalasaToolboxComponent, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.GalasaToolboxComponent))
	})
	return ret, err
}

// Get retrieves the GalasaToolboxComponent from the indexer for a given namespace and name.
func (s galasaToolboxComponentNamespaceLister) Get(name string) (*v2alpha1.GalasaToolboxComponent, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("galasatoolboxcomponent"), name)
	}
	return obj.(*v2alpha1.GalasaToolboxComponent), nil
}
