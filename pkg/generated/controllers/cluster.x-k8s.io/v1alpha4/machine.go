/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha4

import (
	"context"
	"sync"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

type MachineHandler func(string, *v1alpha4.Machine) (*v1alpha4.Machine, error)

type MachineController interface {
	generic.ControllerMeta
	MachineClient

	OnChange(ctx context.Context, name string, sync MachineHandler)
	OnRemove(ctx context.Context, name string, sync MachineHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() MachineCache
}

type MachineClient interface {
	Create(*v1alpha4.Machine) (*v1alpha4.Machine, error)
	Update(*v1alpha4.Machine) (*v1alpha4.Machine, error)
	UpdateStatus(*v1alpha4.Machine) (*v1alpha4.Machine, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha4.Machine, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha4.MachineList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha4.Machine, err error)
}

type MachineCache interface {
	Get(namespace, name string) (*v1alpha4.Machine, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha4.Machine, error)

	AddIndexer(indexName string, indexer MachineIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha4.Machine, error)
}

type MachineIndexer func(obj *v1alpha4.Machine) ([]string, error)

type machineController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewMachineController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) MachineController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &machineController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromMachineHandlerToHandler(sync MachineHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha4.Machine
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha4.Machine))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *machineController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha4.Machine))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateMachineDeepCopyOnChange(client MachineClient, obj *v1alpha4.Machine, handler func(obj *v1alpha4.Machine) (*v1alpha4.Machine, error)) (*v1alpha4.Machine, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *machineController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *machineController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *machineController) OnChange(ctx context.Context, name string, sync MachineHandler) {
	c.AddGenericHandler(ctx, name, FromMachineHandlerToHandler(sync))
}

func (c *machineController) OnRemove(ctx context.Context, name string, sync MachineHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromMachineHandlerToHandler(sync)))
}

func (c *machineController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *machineController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *machineController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *machineController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *machineController) Cache() MachineCache {
	return &machineCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *machineController) Create(obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	result := &v1alpha4.Machine{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *machineController) Update(obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	result := &v1alpha4.Machine{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *machineController) UpdateStatus(obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	result := &v1alpha4.Machine{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *machineController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *machineController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha4.Machine, error) {
	result := &v1alpha4.Machine{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *machineController) List(namespace string, opts metav1.ListOptions) (*v1alpha4.MachineList, error) {
	result := &v1alpha4.MachineList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *machineController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *machineController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha4.Machine, error) {
	result := &v1alpha4.Machine{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type machineCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *machineCache) Get(namespace, name string) (*v1alpha4.Machine, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha4.Machine), nil
}

func (c *machineCache) List(namespace string, selector labels.Selector) (ret []*v1alpha4.Machine, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha4.Machine))
	})

	return ret, err
}

func (c *machineCache) AddIndexer(indexName string, indexer MachineIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha4.Machine))
		},
	}))
}

func (c *machineCache) GetByIndex(indexName, key string) (result []*v1alpha4.Machine, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha4.Machine, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha4.Machine))
	}
	return result, nil
}

// MachineStatusHandler is executed for every added or modified Machine. Should return the new status to be updated
type MachineStatusHandler func(obj *v1alpha4.Machine, status v1alpha4.MachineStatus) (v1alpha4.MachineStatus, error)

// MachineGeneratingHandler is the top-level handler that is executed for every Machine event. It extends MachineStatusHandler by a returning a slice of child objects to be passed to apply.Apply
type MachineGeneratingHandler func(obj *v1alpha4.Machine, status v1alpha4.MachineStatus) ([]runtime.Object, v1alpha4.MachineStatus, error)

// RegisterMachineStatusHandler configures a MachineController to execute a MachineStatusHandler for every events observed.
// If a non-empty condition is provided, it will be updated in the status conditions for every handler execution
func RegisterMachineStatusHandler(ctx context.Context, controller MachineController, condition condition.Cond, name string, handler MachineStatusHandler) {
	statusHandler := &machineStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromMachineHandlerToHandler(statusHandler.sync))
}

// RegisterMachineGeneratingHandler configures a MachineController to execute a MachineGeneratingHandler for every events observed, passing the returned objects to the provided apply.Apply.
// If a non-empty condition is provided, it will be updated in the status conditions for every handler execution
func RegisterMachineGeneratingHandler(ctx context.Context, controller MachineController, apply apply.Apply,
	condition condition.Cond, name string, handler MachineGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &machineGeneratingHandler{
		MachineGeneratingHandler: handler,
		apply:                    apply,
		name:                     name,
		gvk:                      controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterMachineStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type machineStatusHandler struct {
	client    MachineClient
	condition condition.Cond
	handler   MachineStatusHandler
}

// sync is executed on every resource addition or modification. Executes the configured handlers and sends the updated status to the Kubernetes API
func (a *machineStatusHandler) sync(key string, obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type machineGeneratingHandler struct {
	MachineGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
	seen  sync.Map
}

// Remove handles the observed deletion of a resource, cascade deleting every associated resource previously applied
func (a *machineGeneratingHandler) Remove(key string, obj *v1alpha4.Machine) (*v1alpha4.Machine, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1alpha4.Machine{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	if a.opts.UniqueApplyForResourceVersion {
		a.seen.Delete(key)
	}

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

// Handle executes the configured MachineGeneratingHandler and pass the resulting objects to apply.Apply, finally returning the new status of the resource
func (a *machineGeneratingHandler) Handle(obj *v1alpha4.Machine, status v1alpha4.MachineStatus) (v1alpha4.MachineStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.MachineGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}
	if !a.isNewResourceVersion(obj) {
		return newStatus, nil
	}

	err = generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
	if err != nil {
		return newStatus, err
	}
	a.storeResourceVersion(obj)
	return newStatus, nil
}

// isNewResourceVersion detects if a specific resource version was already successfully processed.
// Only used if UniqueApplyForResourceVersion is set in generic.GeneratingHandlerOptions
func (a *machineGeneratingHandler) isNewResourceVersion(obj *v1alpha4.Machine) bool {
	if !a.opts.UniqueApplyForResourceVersion {
		return true
	}

	// Apply once per resource version
	key := obj.Namespace + "/" + obj.Name
	previous, ok := a.seen.Load(key)
	return !ok || previous != obj.ResourceVersion
}

// storeResourceVersion keeps track of the latest resource version of an object for which Apply was executed
// Only used if UniqueApplyForResourceVersion is set in generic.GeneratingHandlerOptions
func (a *machineGeneratingHandler) storeResourceVersion(obj *v1alpha4.Machine) {
	if !a.opts.UniqueApplyForResourceVersion {
		return
	}

	key := obj.Namespace + "/" + obj.Name
	a.seen.Store(key, obj.ResourceVersion)
}
