//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutRestartTarget) DeepCopyInto(out *RolloutRestartTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutRestartTarget.
func (in *RolloutRestartTarget) DeepCopy() *RolloutRestartTarget {
	if in == nil {
		return nil
	}
	out := new(RolloutRestartTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageEncryption) DeepCopyInto(out *StorageEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageEncryption.
func (in *StorageEncryption) DeepCopy() *StorageEncryption {
	if in == nil {
		return nil
	}
	out := new(StorageEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Token) DeepCopyInto(out *Token) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(v1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Token.
func (in *Token) DeepCopy() *Token {
	if in == nil {
		return nil
	}
	out := new(Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuth) DeepCopyInto(out *VaultAuth) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuth.
func (in *VaultAuth) DeepCopy() *VaultAuth {
	if in == nil {
		return nil
	}
	out := new(VaultAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthConfigAppRole) DeepCopyInto(out *VaultAuthConfigAppRole) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthConfigAppRole.
func (in *VaultAuthConfigAppRole) DeepCopy() *VaultAuthConfigAppRole {
	if in == nil {
		return nil
	}
	out := new(VaultAuthConfigAppRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthConfigJwt) DeepCopyInto(out *VaultAuthConfigJwt) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(Token)
		(*in).DeepCopyInto(*out)
	}
	if in.TokenAudiences != nil {
		in, out := &in.TokenAudiences, &out.TokenAudiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthConfigJwt.
func (in *VaultAuthConfigJwt) DeepCopy() *VaultAuthConfigJwt {
	if in == nil {
		return nil
	}
	out := new(VaultAuthConfigJwt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthConfigKubernetes) DeepCopyInto(out *VaultAuthConfigKubernetes) {
	*out = *in
	if in.TokenAudiences != nil {
		in, out := &in.TokenAudiences, &out.TokenAudiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthConfigKubernetes.
func (in *VaultAuthConfigKubernetes) DeepCopy() *VaultAuthConfigKubernetes {
	if in == nil {
		return nil
	}
	out := new(VaultAuthConfigKubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthList) DeepCopyInto(out *VaultAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthList.
func (in *VaultAuthList) DeepCopy() *VaultAuthList {
	if in == nil {
		return nil
	}
	out := new(VaultAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthSpec) DeepCopyInto(out *VaultAuthSpec) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(VaultAuthConfigKubernetes)
		(*in).DeepCopyInto(*out)
	}
	if in.Jwt != nil {
		in, out := &in.Jwt, &out.Jwt
		*out = new(VaultAuthConfigJwt)
		(*in).DeepCopyInto(*out)
	}
	if in.AppRole != nil {
		in, out := &in.AppRole, &out.AppRole
		*out = new(VaultAuthConfigAppRole)
		**out = **in
	}
	if in.StorageEncryption != nil {
		in, out := &in.StorageEncryption, &out.StorageEncryption
		*out = new(StorageEncryption)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthSpec.
func (in *VaultAuthSpec) DeepCopy() *VaultAuthSpec {
	if in == nil {
		return nil
	}
	out := new(VaultAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthStatus) DeepCopyInto(out *VaultAuthStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthStatus.
func (in *VaultAuthStatus) DeepCopy() *VaultAuthStatus {
	if in == nil {
		return nil
	}
	out := new(VaultAuthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConnection) DeepCopyInto(out *VaultConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConnection.
func (in *VaultConnection) DeepCopy() *VaultConnection {
	if in == nil {
		return nil
	}
	out := new(VaultConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConnectionList) DeepCopyInto(out *VaultConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConnectionList.
func (in *VaultConnectionList) DeepCopy() *VaultConnectionList {
	if in == nil {
		return nil
	}
	out := new(VaultConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConnectionSpec) DeepCopyInto(out *VaultConnectionSpec) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConnectionSpec.
func (in *VaultConnectionSpec) DeepCopy() *VaultConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(VaultConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConnectionStatus) DeepCopyInto(out *VaultConnectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConnectionStatus.
func (in *VaultConnectionStatus) DeepCopy() *VaultConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(VaultConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultDynamicSecret) DeepCopyInto(out *VaultDynamicSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultDynamicSecret.
func (in *VaultDynamicSecret) DeepCopy() *VaultDynamicSecret {
	if in == nil {
		return nil
	}
	out := new(VaultDynamicSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultDynamicSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultDynamicSecretList) DeepCopyInto(out *VaultDynamicSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultDynamicSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultDynamicSecretList.
func (in *VaultDynamicSecretList) DeepCopy() *VaultDynamicSecretList {
	if in == nil {
		return nil
	}
	out := new(VaultDynamicSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultDynamicSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultDynamicSecretSpec) DeepCopyInto(out *VaultDynamicSecretSpec) {
	*out = *in
	if in.RolloutRestartTargets != nil {
		in, out := &in.RolloutRestartTargets, &out.RolloutRestartTargets
		*out = make([]RolloutRestartTarget, len(*in))
		copy(*out, *in)
	}
	in.Destination.DeepCopyInto(&out.Destination)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultDynamicSecretSpec.
func (in *VaultDynamicSecretSpec) DeepCopy() *VaultDynamicSecretSpec {
	if in == nil {
		return nil
	}
	out := new(VaultDynamicSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultDynamicSecretStatus) DeepCopyInto(out *VaultDynamicSecretStatus) {
	*out = *in
	out.SecretLease = in.SecretLease
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultDynamicSecretStatus.
func (in *VaultDynamicSecretStatus) DeepCopy() *VaultDynamicSecretStatus {
	if in == nil {
		return nil
	}
	out := new(VaultDynamicSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPKISecret) DeepCopyInto(out *VaultPKISecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPKISecret.
func (in *VaultPKISecret) DeepCopy() *VaultPKISecret {
	if in == nil {
		return nil
	}
	out := new(VaultPKISecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPKISecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPKISecretList) DeepCopyInto(out *VaultPKISecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultPKISecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPKISecretList.
func (in *VaultPKISecretList) DeepCopy() *VaultPKISecretList {
	if in == nil {
		return nil
	}
	out := new(VaultPKISecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPKISecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPKISecretSpec) DeepCopyInto(out *VaultPKISecretSpec) {
	*out = *in
	if in.RolloutRestartTargets != nil {
		in, out := &in.RolloutRestartTargets, &out.RolloutRestartTargets
		*out = make([]RolloutRestartTarget, len(*in))
		copy(*out, *in)
	}
	in.Destination.DeepCopyInto(&out.Destination)
	if in.AltNames != nil {
		in, out := &in.AltNames, &out.AltNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPSans != nil {
		in, out := &in.IPSans, &out.IPSans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.URISans != nil {
		in, out := &in.URISans, &out.URISans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPKISecretSpec.
func (in *VaultPKISecretSpec) DeepCopy() *VaultPKISecretSpec {
	if in == nil {
		return nil
	}
	out := new(VaultPKISecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPKISecretStatus) DeepCopyInto(out *VaultPKISecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPKISecretStatus.
func (in *VaultPKISecretStatus) DeepCopy() *VaultPKISecretStatus {
	if in == nil {
		return nil
	}
	out := new(VaultPKISecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSecretLease) DeepCopyInto(out *VaultSecretLease) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSecretLease.
func (in *VaultSecretLease) DeepCopy() *VaultSecretLease {
	if in == nil {
		return nil
	}
	out := new(VaultSecretLease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStaticSecret) DeepCopyInto(out *VaultStaticSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStaticSecret.
func (in *VaultStaticSecret) DeepCopy() *VaultStaticSecret {
	if in == nil {
		return nil
	}
	out := new(VaultStaticSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultStaticSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStaticSecretList) DeepCopyInto(out *VaultStaticSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultStaticSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStaticSecretList.
func (in *VaultStaticSecretList) DeepCopy() *VaultStaticSecretList {
	if in == nil {
		return nil
	}
	out := new(VaultStaticSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultStaticSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStaticSecretSpec) DeepCopyInto(out *VaultStaticSecretSpec) {
	*out = *in
	if in.RolloutRestartTargets != nil {
		in, out := &in.RolloutRestartTargets, &out.RolloutRestartTargets
		*out = make([]RolloutRestartTarget, len(*in))
		copy(*out, *in)
	}
	in.Destination.DeepCopyInto(&out.Destination)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStaticSecretSpec.
func (in *VaultStaticSecretSpec) DeepCopy() *VaultStaticSecretSpec {
	if in == nil {
		return nil
	}
	out := new(VaultStaticSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultStaticSecretStatus) DeepCopyInto(out *VaultStaticSecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultStaticSecretStatus.
func (in *VaultStaticSecretStatus) DeepCopy() *VaultStaticSecretStatus {
	if in == nil {
		return nil
	}
	out := new(VaultStaticSecretStatus)
	in.DeepCopyInto(out)
	return out
}
