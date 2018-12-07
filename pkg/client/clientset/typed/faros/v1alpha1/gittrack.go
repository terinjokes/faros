/*
Copyright 2018 Pusher Ltd.

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

package v1alpha1

import (
	v1alpha1 "github.com/pusher/faros/pkg/apis/faros/v1alpha1"
	scheme "github.com/pusher/faros/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GitTracksGetter has a method to return a GitTrackInterface.
// A group's client should implement this interface.
type GitTracksGetter interface {
	GitTracks(namespace string) GitTrackInterface
}

// GitTrackInterface has methods to work with GitTrack resources.
type GitTrackInterface interface {
	Create(*v1alpha1.GitTrack) (*v1alpha1.GitTrack, error)
	Update(*v1alpha1.GitTrack) (*v1alpha1.GitTrack, error)
	UpdateStatus(*v1alpha1.GitTrack) (*v1alpha1.GitTrack, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GitTrack, error)
	List(opts v1.ListOptions) (*v1alpha1.GitTrackList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitTrack, err error)
	GitTrackExpansion
}

// gitTracks implements GitTrackInterface
type gitTracks struct {
	client rest.Interface
	ns     string
}

// newGitTracks returns a GitTracks
func newGitTracks(c *FarosV1alpha1Client, namespace string) *gitTracks {
	return &gitTracks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitTrack, and returns the corresponding gitTrack object, and an error if there is any.
func (c *gitTracks) Get(name string, options v1.GetOptions) (result *v1alpha1.GitTrack, err error) {
	result = &v1alpha1.GitTrack{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gittracks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitTracks that match those selectors.
func (c *gitTracks) List(opts v1.ListOptions) (result *v1alpha1.GitTrackList, err error) {
	result = &v1alpha1.GitTrackList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gittracks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitTracks.
func (c *gitTracks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gittracks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gitTrack and creates it.  Returns the server's representation of the gitTrack, and an error, if there is any.
func (c *gitTracks) Create(gitTrack *v1alpha1.GitTrack) (result *v1alpha1.GitTrack, err error) {
	result = &v1alpha1.GitTrack{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gittracks").
		Body(gitTrack).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gitTrack and updates it. Returns the server's representation of the gitTrack, and an error, if there is any.
func (c *gitTracks) Update(gitTrack *v1alpha1.GitTrack) (result *v1alpha1.GitTrack, err error) {
	result = &v1alpha1.GitTrack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gittracks").
		Name(gitTrack.Name).
		Body(gitTrack).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gitTracks) UpdateStatus(gitTrack *v1alpha1.GitTrack) (result *v1alpha1.GitTrack, err error) {
	result = &v1alpha1.GitTrack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gittracks").
		Name(gitTrack.Name).
		SubResource("status").
		Body(gitTrack).
		Do().
		Into(result)
	return
}

// Delete takes name of the gitTrack and deletes it. Returns an error if one occurs.
func (c *gitTracks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gittracks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitTracks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gittracks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gitTrack.
func (c *gitTracks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitTrack, err error) {
	result = &v1alpha1.GitTrack{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gittracks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}