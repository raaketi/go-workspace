package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "contoso.com"
const GroupVersion = "v1alpha1"

var SchemeGroupVersion = schema.GroupVersion{Group: "contoso.com", Version: "v1alpha1"}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Myplatform{},
		&MyplatformList{},
	)
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// func Newclient(cfg *rest.Config) (*rest.RESTClient, error) {
// 	scheme := runtime.NewScheme()
// 	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
// 	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
// 		return nil, err
// 	}
// 	config := *cfg
// 	config.GroupVersion = &SchemeGroupVersion
// 	config.APIPath = "/apis"
// 	config.ContentType = runtime.ContentTypeJSON
// 	// config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}
// 	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme)
// 	config.UserAgent = rest.DefaultKubernetesUserAgent()
// 	client, err := rest.RESTClientFor(&config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }
