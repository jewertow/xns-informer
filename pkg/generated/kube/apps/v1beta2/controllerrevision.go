/*
Copyright Red Hat, Inc.

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

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta2

import (
	"context"
	time "time"

	internalinterfaces "github.com/maistra/xns-informer/pkg/generated/kube/internalinterfaces"
	informers "github.com/maistra/xns-informer/pkg/informers"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta2 "k8s.io/client-go/listers/apps/v1beta2"
	cache "k8s.io/client-go/tools/cache"
)

// ControllerRevisionInformer provides access to a shared informer and lister for
// ControllerRevisions.
type ControllerRevisionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.ControllerRevisionLister
}

type controllerRevisionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespaces       informers.NamespaceSet
}

// NewControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewControllerRevisionInformer(client kubernetes.Interface, namespaces informers.NamespaceSet, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, namespaces, resyncPeriod, indexers, nil)
}

// NewFilteredControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredControllerRevisionInformer(client kubernetes.Interface, namespaces informers.NamespaceSet, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	newInformer := func(namespace string) cache.SharedIndexInformer {
		return cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.AppsV1beta2().ControllerRevisions(namespace).List(context.TODO(), options)
				},
				WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.AppsV1beta2().ControllerRevisions(namespace).Watch(context.TODO(), options)
				},
			},
			&appsv1beta2.ControllerRevision{},
			resyncPeriod,
			indexers,
		)
	}

	return informers.NewMultiNamespaceInformer(namespaces, resyncPeriod, newInformer)
}

func (f *controllerRevisionInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, f.namespaces, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1beta2.ControllerRevision{}, f.defaultInformer)
}

func (f *controllerRevisionInformer) Lister() v1beta2.ControllerRevisionLister {
	return v1beta2.NewControllerRevisionLister(f.Informer().GetIndexer())
}
