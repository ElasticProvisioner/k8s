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

package fake

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1alpha1 "k8s.io/api/certificates/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	certificatesv1alpha1 "k8s.io/client-go/applyconfigurations/certificates/v1alpha1"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTrustBundles implements ClusterTrustBundleInterface
type FakeClusterTrustBundles struct {
	Fake *FakeCertificatesV1alpha1
}

var clustertrustbundlesResource = v1alpha1.SchemeGroupVersion.WithResource("clustertrustbundles")

var clustertrustbundlesKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterTrustBundle")

// Get takes name of the clusterTrustBundle, and returns the corresponding clusterTrustBundle object, and an error if there is any.
func (c *FakeClusterTrustBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterTrustBundle, err error) {
	emptyResult := &v1alpha1.ClusterTrustBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(clustertrustbundlesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterTrustBundle), err
}

// List takes label and field selectors, and returns the list of ClusterTrustBundles that match those selectors.
func (c *FakeClusterTrustBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterTrustBundleList, err error) {
	emptyResult := &v1alpha1.ClusterTrustBundleList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clustertrustbundlesResource, clustertrustbundlesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterTrustBundleList{ListMeta: obj.(*v1alpha1.ClusterTrustBundleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterTrustBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTrustBundles.
func (c *FakeClusterTrustBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clustertrustbundlesResource, opts))
}

// Create takes the representation of a clusterTrustBundle and creates it.  Returns the server's representation of the clusterTrustBundle, and an error, if there is any.
func (c *FakeClusterTrustBundles) Create(ctx context.Context, clusterTrustBundle *v1alpha1.ClusterTrustBundle, opts v1.CreateOptions) (result *v1alpha1.ClusterTrustBundle, err error) {
	emptyResult := &v1alpha1.ClusterTrustBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clustertrustbundlesResource, clusterTrustBundle, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterTrustBundle), err
}

// Update takes the representation of a clusterTrustBundle and updates it. Returns the server's representation of the clusterTrustBundle, and an error, if there is any.
func (c *FakeClusterTrustBundles) Update(ctx context.Context, clusterTrustBundle *v1alpha1.ClusterTrustBundle, opts v1.UpdateOptions) (result *v1alpha1.ClusterTrustBundle, err error) {
	emptyResult := &v1alpha1.ClusterTrustBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clustertrustbundlesResource, clusterTrustBundle, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterTrustBundle), err
}

// Delete takes name of the clusterTrustBundle and deletes it. Returns an error if one occurs.
func (c *FakeClusterTrustBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustertrustbundlesResource, name, opts), &v1alpha1.ClusterTrustBundle{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTrustBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clustertrustbundlesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterTrustBundleList{})
	return err
}

// Patch applies the patch and returns the patched clusterTrustBundle.
func (c *FakeClusterTrustBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTrustBundle, err error) {
	emptyResult := &v1alpha1.ClusterTrustBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clustertrustbundlesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterTrustBundle), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterTrustBundle.
func (c *FakeClusterTrustBundles) Apply(ctx context.Context, clusterTrustBundle *certificatesv1alpha1.ClusterTrustBundleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterTrustBundle, err error) {
	if clusterTrustBundle == nil {
		return nil, fmt.Errorf("clusterTrustBundle provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterTrustBundle)
	if err != nil {
		return nil, err
	}
	name := clusterTrustBundle.Name
	if name == nil {
		return nil, fmt.Errorf("clusterTrustBundle.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.ClusterTrustBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clustertrustbundlesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterTrustBundle), err
}
