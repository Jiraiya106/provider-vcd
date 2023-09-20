//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataEntryObservation) DeepCopyInto(out *MetadataEntryObservation) {
	*out = *in
	if in.IsSystem != nil {
		in, out := &in.IsSystem, &out.IsSystem
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserAccess != nil {
		in, out := &in.UserAccess, &out.UserAccess
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataEntryObservation.
func (in *MetadataEntryObservation) DeepCopy() *MetadataEntryObservation {
	if in == nil {
		return nil
	}
	out := new(MetadataEntryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataEntryParameters) DeepCopyInto(out *MetadataEntryParameters) {
	*out = *in
	if in.IsSystem != nil {
		in, out := &in.IsSystem, &out.IsSystem
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserAccess != nil {
		in, out := &in.UserAccess, &out.UserAccess
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataEntryParameters.
func (in *MetadataEntryParameters) DeepCopy() *MetadataEntryParameters {
	if in == nil {
		return nil
	}
	out := new(MetadataEntryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutedV2) DeepCopyInto(out *RoutedV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutedV2.
func (in *RoutedV2) DeepCopy() *RoutedV2 {
	if in == nil {
		return nil
	}
	out := new(RoutedV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoutedV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutedV2List) DeepCopyInto(out *RoutedV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoutedV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutedV2List.
func (in *RoutedV2List) DeepCopy() *RoutedV2List {
	if in == nil {
		return nil
	}
	out := new(RoutedV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoutedV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutedV2Observation) DeepCopyInto(out *RoutedV2Observation) {
	*out = *in
	if in.DNSSuffix != nil {
		in, out := &in.DNSSuffix, &out.DNSSuffix
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dns1 != nil {
		in, out := &in.Dns1, &out.Dns1
		*out = new(string)
		**out = **in
	}
	if in.Dns2 != nil {
		in, out := &in.Dns2, &out.Dns2
		*out = new(string)
		**out = **in
	}
	if in.EdgeGatewayID != nil {
		in, out := &in.EdgeGatewayID, &out.EdgeGatewayID
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InterfaceType != nil {
		in, out := &in.InterfaceType, &out.InterfaceType
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MetadataEntry != nil {
		in, out := &in.MetadataEntry, &out.MetadataEntry
		*out = make([]MetadataEntryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Org != nil {
		in, out := &in.Org, &out.Org
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.StaticIPPool != nil {
		in, out := &in.StaticIPPool, &out.StaticIPPool
		*out = make([]StaticIPPoolObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vdc != nil {
		in, out := &in.Vdc, &out.Vdc
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutedV2Observation.
func (in *RoutedV2Observation) DeepCopy() *RoutedV2Observation {
	if in == nil {
		return nil
	}
	out := new(RoutedV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutedV2Parameters) DeepCopyInto(out *RoutedV2Parameters) {
	*out = *in
	if in.DNSSuffix != nil {
		in, out := &in.DNSSuffix, &out.DNSSuffix
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dns1 != nil {
		in, out := &in.Dns1, &out.Dns1
		*out = new(string)
		**out = **in
	}
	if in.Dns2 != nil {
		in, out := &in.Dns2, &out.Dns2
		*out = new(string)
		**out = **in
	}
	if in.EdgeGatewayID != nil {
		in, out := &in.EdgeGatewayID, &out.EdgeGatewayID
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.InterfaceType != nil {
		in, out := &in.InterfaceType, &out.InterfaceType
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MetadataEntry != nil {
		in, out := &in.MetadataEntry, &out.MetadataEntry
		*out = make([]MetadataEntryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Org != nil {
		in, out := &in.Org, &out.Org
		*out = new(string)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.StaticIPPool != nil {
		in, out := &in.StaticIPPool, &out.StaticIPPool
		*out = make([]StaticIPPoolParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vdc != nil {
		in, out := &in.Vdc, &out.Vdc
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutedV2Parameters.
func (in *RoutedV2Parameters) DeepCopy() *RoutedV2Parameters {
	if in == nil {
		return nil
	}
	out := new(RoutedV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutedV2Spec) DeepCopyInto(out *RoutedV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutedV2Spec.
func (in *RoutedV2Spec) DeepCopy() *RoutedV2Spec {
	if in == nil {
		return nil
	}
	out := new(RoutedV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutedV2Status) DeepCopyInto(out *RoutedV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutedV2Status.
func (in *RoutedV2Status) DeepCopy() *RoutedV2Status {
	if in == nil {
		return nil
	}
	out := new(RoutedV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticIPPoolObservation) DeepCopyInto(out *StaticIPPoolObservation) {
	*out = *in
	if in.EndAddress != nil {
		in, out := &in.EndAddress, &out.EndAddress
		*out = new(string)
		**out = **in
	}
	if in.StartAddress != nil {
		in, out := &in.StartAddress, &out.StartAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticIPPoolObservation.
func (in *StaticIPPoolObservation) DeepCopy() *StaticIPPoolObservation {
	if in == nil {
		return nil
	}
	out := new(StaticIPPoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticIPPoolParameters) DeepCopyInto(out *StaticIPPoolParameters) {
	*out = *in
	if in.EndAddress != nil {
		in, out := &in.EndAddress, &out.EndAddress
		*out = new(string)
		**out = **in
	}
	if in.StartAddress != nil {
		in, out := &in.StartAddress, &out.StartAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticIPPoolParameters.
func (in *StaticIPPoolParameters) DeepCopy() *StaticIPPoolParameters {
	if in == nil {
		return nil
	}
	out := new(StaticIPPoolParameters)
	in.DeepCopyInto(out)
	return out
}
