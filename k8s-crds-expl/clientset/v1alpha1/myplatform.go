package v1alpha1

import (
	"github.com/martin-helmich/kubernetes-crd-example/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/scheme"
)

type MyplatformInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.ProjectList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Myplatform, error)
	Create(*v1alpha1.Project) (*v1alpha1.Project, error)
	// Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type MyplatformClient struct {
	restClient dynamic.Interface
	ns         string
}

func (c *MyplatformClient) List(opts metav1.ListOptions) (*v1alpha1.MyplatformList, error) {
	result := v1alpha1.MyplatformList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("myplatforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *MyplatformClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Myplatform, error) {
	result := v1alpha1.Myplatform{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("myplatforms").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *MyplatformClient) Create(myplatform *v1alpha1.Myplatform) (*v1alpha1.Myplatform, error) {
	result := v1alpha1.Myplatform{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("myplatforms").
		Body(myplatform).
		Do().
		Into(&result)

	return &result, err
}

// func (c *projectClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
// 	opts.Watch = true
// 	return c.restClient.
// 		Get().
// 		Namespace(c.ns).
// 		Resource("projects").
// 		VersionedParams(&opts, scheme.ParameterCodec).
// 		Watch()
// }
