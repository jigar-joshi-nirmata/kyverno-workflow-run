/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/nirmata/kyverno/pkg/api/kyverno/v1alpha1"
	scheme "github.com/nirmata/kyverno/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterPolicyViolationsGetter has a method to return a ClusterPolicyViolationInterface.
// A group's client should implement this interface.
type ClusterPolicyViolationsGetter interface {
	ClusterPolicyViolations() ClusterPolicyViolationInterface
}

// ClusterPolicyViolationInterface has methods to work with ClusterPolicyViolation resources.
type ClusterPolicyViolationInterface interface {
	Create(*v1alpha1.ClusterPolicyViolation) (*v1alpha1.ClusterPolicyViolation, error)
	Update(*v1alpha1.ClusterPolicyViolation) (*v1alpha1.ClusterPolicyViolation, error)
	UpdateStatus(*v1alpha1.ClusterPolicyViolation) (*v1alpha1.ClusterPolicyViolation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterPolicyViolation, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterPolicyViolationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterPolicyViolation, err error)
	ClusterPolicyViolationExpansion
}

// clusterPolicyViolations implements ClusterPolicyViolationInterface
type clusterPolicyViolations struct {
	client rest.Interface
}

// newClusterPolicyViolations returns a ClusterPolicyViolations
func newClusterPolicyViolations(c *KyvernoV1alpha1Client) *clusterPolicyViolations {
	return &clusterPolicyViolations{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterPolicyViolation, and returns the corresponding clusterPolicyViolation object, and an error if there is any.
func (c *clusterPolicyViolations) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterPolicyViolation, err error) {
	result = &v1alpha1.ClusterPolicyViolation{}
	err = c.client.Get().
		Resource("clusterpolicyviolations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterPolicyViolations that match those selectors.
func (c *clusterPolicyViolations) List(opts v1.ListOptions) (result *v1alpha1.ClusterPolicyViolationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterPolicyViolationList{}
	err = c.client.Get().
		Resource("clusterpolicyviolations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterPolicyViolations.
func (c *clusterPolicyViolations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterpolicyviolations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterPolicyViolation and creates it.  Returns the server's representation of the clusterPolicyViolation, and an error, if there is any.
func (c *clusterPolicyViolations) Create(clusterPolicyViolation *v1alpha1.ClusterPolicyViolation) (result *v1alpha1.ClusterPolicyViolation, err error) {
	result = &v1alpha1.ClusterPolicyViolation{}
	err = c.client.Post().
		Resource("clusterpolicyviolations").
		Body(clusterPolicyViolation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterPolicyViolation and updates it. Returns the server's representation of the clusterPolicyViolation, and an error, if there is any.
func (c *clusterPolicyViolations) Update(clusterPolicyViolation *v1alpha1.ClusterPolicyViolation) (result *v1alpha1.ClusterPolicyViolation, err error) {
	result = &v1alpha1.ClusterPolicyViolation{}
	err = c.client.Put().
		Resource("clusterpolicyviolations").
		Name(clusterPolicyViolation.Name).
		Body(clusterPolicyViolation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterPolicyViolations) UpdateStatus(clusterPolicyViolation *v1alpha1.ClusterPolicyViolation) (result *v1alpha1.ClusterPolicyViolation, err error) {
	result = &v1alpha1.ClusterPolicyViolation{}
	err = c.client.Put().
		Resource("clusterpolicyviolations").
		Name(clusterPolicyViolation.Name).
		SubResource("status").
		Body(clusterPolicyViolation).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterPolicyViolation and deletes it. Returns an error if one occurs.
func (c *clusterPolicyViolations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterpolicyviolations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterPolicyViolations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterpolicyviolations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterPolicyViolation.
func (c *clusterPolicyViolations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterPolicyViolation, err error) {
	result = &v1alpha1.ClusterPolicyViolation{}
	err = c.client.Patch(pt).
		Resource("clusterpolicyviolations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
