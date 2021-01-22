/*
Copyright the Velero contributors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/j4m3s-s/velero/pkg/apis/velero/v1"
	scheme "github.com/j4m3s-s/velero/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupStorageLocationsGetter has a method to return a BackupStorageLocationInterface.
// A group's client should implement this interface.
type BackupStorageLocationsGetter interface {
	BackupStorageLocations(namespace string) BackupStorageLocationInterface
}

// BackupStorageLocationInterface has methods to work with BackupStorageLocation resources.
type BackupStorageLocationInterface interface {
	Create(ctx context.Context, backupStorageLocation *v1.BackupStorageLocation, opts metav1.CreateOptions) (*v1.BackupStorageLocation, error)
	Update(ctx context.Context, backupStorageLocation *v1.BackupStorageLocation, opts metav1.UpdateOptions) (*v1.BackupStorageLocation, error)
	UpdateStatus(ctx context.Context, backupStorageLocation *v1.BackupStorageLocation, opts metav1.UpdateOptions) (*v1.BackupStorageLocation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.BackupStorageLocation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BackupStorageLocationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BackupStorageLocation, err error)
	BackupStorageLocationExpansion
}

// backupStorageLocations implements BackupStorageLocationInterface
type backupStorageLocations struct {
	client rest.Interface
	ns     string
}

// newBackupStorageLocations returns a BackupStorageLocations
func newBackupStorageLocations(c *VeleroV1Client, namespace string) *backupStorageLocations {
	return &backupStorageLocations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupStorageLocation, and returns the corresponding backupStorageLocation object, and an error if there is any.
func (c *backupStorageLocations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.BackupStorageLocation, err error) {
	result = &v1.BackupStorageLocation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupStorageLocations that match those selectors.
func (c *backupStorageLocations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BackupStorageLocationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.BackupStorageLocationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupStorageLocations.
func (c *backupStorageLocations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backupStorageLocation and creates it.  Returns the server's representation of the backupStorageLocation, and an error, if there is any.
func (c *backupStorageLocations) Create(ctx context.Context, backupStorageLocation *v1.BackupStorageLocation, opts metav1.CreateOptions) (result *v1.BackupStorageLocation, err error) {
	result = &v1.BackupStorageLocation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupStorageLocation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backupStorageLocation and updates it. Returns the server's representation of the backupStorageLocation, and an error, if there is any.
func (c *backupStorageLocations) Update(ctx context.Context, backupStorageLocation *v1.BackupStorageLocation, opts metav1.UpdateOptions) (result *v1.BackupStorageLocation, err error) {
	result = &v1.BackupStorageLocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		Name(backupStorageLocation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupStorageLocation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backupStorageLocations) UpdateStatus(ctx context.Context, backupStorageLocation *v1.BackupStorageLocation, opts metav1.UpdateOptions) (result *v1.BackupStorageLocation, err error) {
	result = &v1.BackupStorageLocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		Name(backupStorageLocation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupStorageLocation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backupStorageLocation and deletes it. Returns an error if one occurs.
func (c *backupStorageLocations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupStorageLocations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupstoragelocations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backupStorageLocation.
func (c *backupStorageLocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BackupStorageLocation, err error) {
	result = &v1.BackupStorageLocation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupstoragelocations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
