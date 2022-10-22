package v1alpha1

import meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Myplatform struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               MyplatformSpec   `json:"spec"`
	Status             MyplatformStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyplatformList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []Myplatform `json:"items"`
}
type MyplatformSpec struct {
	AppId        string `json:"appId"`
	Language     string `json:"language"`
	Os           string `json:"os"`
	InstanceSize string `json:"instanceSize"`
}
type MyplatformStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}
