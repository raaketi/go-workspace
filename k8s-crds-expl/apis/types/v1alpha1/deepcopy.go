package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Myplatform) DeepCopyInto(out *Myplatform) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = MyplatformSpec{
		AppId:        in.Spec.AppId,
		Language:     in.Spec.Language,
		Os:           in.Spec.Os,
		InstanceSize: in.Spec.InstanceSize,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Myplatform) DeepCopyObject() runtime.Object {
	out := Myplatform{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *MyplatformList) DeepCopyObject() runtime.Object {
	out := MyplatformList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Myplatform, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
