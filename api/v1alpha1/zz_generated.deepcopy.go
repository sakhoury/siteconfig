//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	metal3_iov1alpha1 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
	"github.com/openshift/assisted-service/api/v1beta1"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BmcCredentialsName) DeepCopyInto(out *BmcCredentialsName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BmcCredentialsName.
func (in *BmcCredentialsName) DeepCopy() *BmcCredentialsName {
	if in == nil {
		return nil
	}
	out := new(BmcCredentialsName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstance) DeepCopyInto(out *ClusterInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstance.
func (in *ClusterInstance) DeepCopy() *ClusterInstance {
	if in == nil {
		return nil
	}
	out := new(ClusterInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceList) DeepCopyInto(out *ClusterInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceList.
func (in *ClusterInstanceList) DeepCopy() *ClusterInstanceList {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceSpec) DeepCopyInto(out *ClusterInstanceSpec) {
	*out = *in
	out.PullSecretRef = in.PullSecretRef
	if in.ApiVIPs != nil {
		in, out := &in.ApiVIPs, &out.ApiVIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IngressVIPs != nil {
		in, out := &in.IngressVIPs, &out.IngressVIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalNTPSources != nil {
		in, out := &in.AdditionalNTPSources, &out.AdditionalNTPSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MachineNetwork != nil {
		in, out := &in.MachineNetwork, &out.MachineNetwork
		*out = make([]MachineNetworkEntry, len(*in))
		copy(*out, *in)
	}
	if in.ClusterNetwork != nil {
		in, out := &in.ClusterNetwork, &out.ClusterNetwork
		*out = make([]ClusterNetworkEntry, len(*in))
		copy(*out, *in)
	}
	if in.ServiceNetwork != nil {
		in, out := &in.ServiceNetwork, &out.ServiceNetwork
		*out = make([]ServiceNetworkEntry, len(*in))
		copy(*out, *in)
	}
	if in.ExtraAnnotations != nil {
		in, out := &in.ExtraAnnotations, &out.ExtraAnnotations
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.DiskEncryption != nil {
		in, out := &in.DiskEncryption, &out.DiskEncryption
		*out = new(DiskEncryption)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(v1beta1.Proxy)
		**out = **in
	}
	if in.ExtraManifestsRefs != nil {
		in, out := &in.ExtraManifestsRefs, &out.ExtraManifestsRefs
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.SuppressedManifests != nil {
		in, out := &in.SuppressedManifests, &out.SuppressedManifests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PruneManifests != nil {
		in, out := &in.PruneManifests, &out.PruneManifests
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.TemplateRefs != nil {
		in, out := &in.TemplateRefs, &out.TemplateRefs
		*out = make([]TemplateRef, len(*in))
		copy(*out, *in)
	}
	if in.CaBundleRef != nil {
		in, out := &in.CaBundleRef, &out.CaBundleRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Reinstall != nil {
		in, out := &in.Reinstall, &out.Reinstall
		*out = new(ReinstallSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceSpec.
func (in *ClusterInstanceSpec) DeepCopy() *ClusterInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstanceStatus) DeepCopyInto(out *ClusterInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterDeploymentRef != nil {
		in, out := &in.ClusterDeploymentRef, &out.ClusterDeploymentRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.DeploymentConditions != nil {
		in, out := &in.DeploymentConditions, &out.DeploymentConditions
		*out = make([]hivev1.ClusterDeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManifestsRendered != nil {
		in, out := &in.ManifestsRendered, &out.ManifestsRendered
		*out = make([]ManifestReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Reinstall != nil {
		in, out := &in.Reinstall, &out.Reinstall
		*out = new(ReinstallStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(PausedStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstanceStatus.
func (in *ClusterInstanceStatus) DeepCopy() *ClusterInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkEntry) DeepCopyInto(out *ClusterNetworkEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkEntry.
func (in *ClusterNetworkEntry) DeepCopy() *ClusterNetworkEntry {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryption) DeepCopyInto(out *DiskEncryption) {
	*out = *in
	if in.Tang != nil {
		in, out := &in.Tang, &out.Tang
		*out = make([]TangConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryption.
func (in *DiskEncryption) DeepCopy() *DiskEncryption {
	if in == nil {
		return nil
	}
	out := new(DiskEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostRef) DeepCopyInto(out *HostRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostRef.
func (in *HostRef) DeepCopy() *HostRef {
	if in == nil {
		return nil
	}
	out := new(HostRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineNetworkEntry) DeepCopyInto(out *MachineNetworkEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineNetworkEntry.
func (in *MachineNetworkEntry) DeepCopy() *MachineNetworkEntry {
	if in == nil {
		return nil
	}
	out := new(MachineNetworkEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestReference) DeepCopyInto(out *ManifestReference) {
	*out = *in
	if in.APIGroup != nil {
		in, out := &in.APIGroup, &out.APIGroup
		*out = new(string)
		**out = **in
	}
	in.LastAppliedTime.DeepCopyInto(&out.LastAppliedTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestReference.
func (in *ManifestReference) DeepCopy() *ManifestReference {
	if in == nil {
		return nil
	}
	out := new(ManifestReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	out.BmcCredentialsName = in.BmcCredentialsName
	if in.RootDeviceHints != nil {
		in, out := &in.RootDeviceHints, &out.RootDeviceHints
		*out = new(metal3_iov1alpha1.RootDeviceHints)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNetwork != nil {
		in, out := &in.NodeNetwork, &out.NodeNetwork
		*out = new(v1beta1.NMStateConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostRef != nil {
		in, out := &in.HostRef, &out.HostRef
		*out = new(HostRef)
		**out = **in
	}
	if in.ExtraAnnotations != nil {
		in, out := &in.ExtraAnnotations, &out.ExtraAnnotations
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.SuppressedManifests != nil {
		in, out := &in.SuppressedManifests, &out.SuppressedManifests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PruneManifests != nil {
		in, out := &in.PruneManifests, &out.PruneManifests
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.TemplateRefs != nil {
		in, out := &in.TemplateRefs, &out.TemplateRefs
		*out = make([]TemplateRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PausedStatus) DeepCopyInto(out *PausedStatus) {
	*out = *in
	in.TimeSet.DeepCopyInto(&out.TimeSet)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PausedStatus.
func (in *PausedStatus) DeepCopy() *PausedStatus {
	if in == nil {
		return nil
	}
	out := new(PausedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReinstallHistory) DeepCopyInto(out *ReinstallHistory) {
	*out = *in
	in.RequestStartTime.DeepCopyInto(&out.RequestStartTime)
	in.RequestEndTime.DeepCopyInto(&out.RequestEndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReinstallHistory.
func (in *ReinstallHistory) DeepCopy() *ReinstallHistory {
	if in == nil {
		return nil
	}
	out := new(ReinstallHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReinstallSpec) DeepCopyInto(out *ReinstallSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReinstallSpec.
func (in *ReinstallSpec) DeepCopy() *ReinstallSpec {
	if in == nil {
		return nil
	}
	out := new(ReinstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReinstallStatus) DeepCopyInto(out *ReinstallStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RequestStartTime.DeepCopyInto(&out.RequestStartTime)
	in.RequestEndTime.DeepCopyInto(&out.RequestEndTime)
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]ReinstallHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReinstallStatus.
func (in *ReinstallStatus) DeepCopy() *ReinstallStatus {
	if in == nil {
		return nil
	}
	out := new(ReinstallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceNetworkEntry) DeepCopyInto(out *ServiceNetworkEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceNetworkEntry.
func (in *ServiceNetworkEntry) DeepCopy() *ServiceNetworkEntry {
	if in == nil {
		return nil
	}
	out := new(ServiceNetworkEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangConfig) DeepCopyInto(out *TangConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangConfig.
func (in *TangConfig) DeepCopy() *TangConfig {
	if in == nil {
		return nil
	}
	out := new(TangConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateRef) DeepCopyInto(out *TemplateRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateRef.
func (in *TemplateRef) DeepCopy() *TemplateRef {
	if in == nil {
		return nil
	}
	out := new(TemplateRef)
	in.DeepCopyInto(out)
	return out
}
