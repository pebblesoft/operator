/*
Copyright 2020 The Tekton Authors

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

package tektonconfig

import (
	"context"

	"github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	openshiftpipelinesascodeinformer "github.com/tektoncd/operator/pkg/client/injection/informers/operator/v1alpha1/openshiftpipelinesascode"
	tektonAddoninformer "github.com/tektoncd/operator/pkg/client/injection/informers/operator/v1alpha1/tektonaddon"
	"github.com/tektoncd/operator/pkg/reconciler/shared/tektonconfig"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
)

// NewController initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
func NewController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	ctrl := tektonconfig.NewExtensibleController(OpenShiftExtension)(ctx, cmw)
	tektonAddoninformer.Get(ctx).Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterController(&v1alpha1.TektonConfig{}),
		Handler:    controller.HandleAll(ctrl.EnqueueControllerOf),
	})
	openshiftpipelinesascodeinformer.Get(ctx).Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterController(&v1alpha1.TektonConfig{}),
		Handler:    controller.HandleAll(ctrl.EnqueueControllerOf),
	})
	return ctrl
}
