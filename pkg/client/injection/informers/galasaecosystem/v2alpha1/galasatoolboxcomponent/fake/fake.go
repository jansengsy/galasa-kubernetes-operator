/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/galasa-dev/galasa-kubernetes-operator/pkg/client/injection/informers/factory/fake"
	galasatoolboxcomponent "github.com/galasa-dev/galasa-kubernetes-operator/pkg/client/injection/informers/galasaecosystem/v2alpha1/galasatoolboxcomponent"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = galasatoolboxcomponent.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Galasa().V2alpha1().GalasaToolboxComponents()
	return context.WithValue(ctx, galasatoolboxcomponent.Key{}, inf), inf.Informer()
}
