//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.General.DeepCopyInto(&out.General)
	out.ConfMgmt = in.ConfMgmt
	in.Dashboards.DeepCopyInto(&out.Dashboards)
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(Security)
		(*in).DeepCopyInto(*out)
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.ComponentsStatus != nil {
		in, out := &in.ComponentsStatus, &out.ComponentsStatus
		*out = make([]ComponentStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfMgmt) DeepCopyInto(out *ConfMgmt) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfMgmt.
func (in *ConfMgmt) DeepCopy() *ConfMgmt {
	if in == nil {
		return nil
	}
	out := new(ConfMgmt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardsConfig) DeepCopyInto(out *DashboardsConfig) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(DashboardsTlsConfig)
		**out = **in
	}
	out.OpensearchCredentialsSecret = in.OpensearchCredentialsSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardsConfig.
func (in *DashboardsConfig) DeepCopy() *DashboardsConfig {
	if in == nil {
		return nil
	}
	out := new(DashboardsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardsTlsConfig) DeepCopyInto(out *DashboardsTlsConfig) {
	*out = *in
	out.CertificateConfig = in.CertificateConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardsTlsConfig.
func (in *DashboardsTlsConfig) DeepCopy() *DashboardsTlsConfig {
	if in == nil {
		return nil
	}
	out := new(DashboardsTlsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralConfig) DeepCopyInto(out *GeneralConfig) {
	*out = *in
	if in.DefaultRepo != nil {
		in, out := &in.DefaultRepo, &out.DefaultRepo
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralConfig.
func (in *GeneralConfig) DeepCopy() *GeneralConfig {
	if in == nil {
		return nil
	}
	out := new(GeneralConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSearchCluster) DeepCopyInto(out *OpenSearchCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSearchCluster.
func (in *OpenSearchCluster) DeepCopy() *OpenSearchCluster {
	if in == nil {
		return nil
	}
	out := new(OpenSearchCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenSearchCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSearchClusterList) DeepCopyInto(out *OpenSearchClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenSearchCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSearchClusterList.
func (in *OpenSearchClusterList) DeepCopy() *OpenSearchClusterList {
	if in == nil {
		return nil
	}
	out := new(OpenSearchClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenSearchClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCSource) DeepCopyInto(out *PVCSource) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]corev1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCSource.
func (in *PVCSource) DeepCopy() *PVCSource {
	if in == nil {
		return nil
	}
	out := new(PVCSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfig) DeepCopyInto(out *PersistenceConfig) {
	*out = *in
	in.PersistenceSource.DeepCopyInto(&out.PersistenceSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfig.
func (in *PersistenceConfig) DeepCopy() *PersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSource) DeepCopyInto(out *PersistenceSource) {
	*out = *in
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(PVCSource)
		(*in).DeepCopyInto(*out)
	}
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(corev1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(corev1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSource.
func (in *PersistenceSource) DeepCopy() *PersistenceSource {
	if in == nil {
		return nil
	}
	out := new(PersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Security) DeepCopyInto(out *Security) {
	*out = *in
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(TlsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(SecurityConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Security.
func (in *Security) DeepCopy() *Security {
	if in == nil {
		return nil
	}
	out := new(Security)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfig) DeepCopyInto(out *SecurityConfig) {
	*out = *in
	out.SecurityconfigSecret = in.SecurityconfigSecret
	out.AdminSecret = in.AdminSecret
	out.AdminCredentialsSecret = in.AdminCredentialsSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfig.
func (in *SecurityConfig) DeepCopy() *SecurityConfig {
	if in == nil {
		return nil
	}
	out := new(SecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsCertificateConfig) DeepCopyInto(out *TlsCertificateConfig) {
	*out = *in
	out.Secret = in.Secret
	out.CaSecret = in.CaSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsCertificateConfig.
func (in *TlsCertificateConfig) DeepCopy() *TlsCertificateConfig {
	if in == nil {
		return nil
	}
	out := new(TlsCertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsConfig) DeepCopyInto(out *TlsConfig) {
	*out = *in
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(TlsConfigTransport)
		(*in).DeepCopyInto(*out)
	}
	if in.Http != nil {
		in, out := &in.Http, &out.Http
		*out = new(TlsConfigHttp)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsConfig.
func (in *TlsConfig) DeepCopy() *TlsConfig {
	if in == nil {
		return nil
	}
	out := new(TlsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsConfigHttp) DeepCopyInto(out *TlsConfigHttp) {
	*out = *in
	out.CertificateConfig = in.CertificateConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsConfigHttp.
func (in *TlsConfigHttp) DeepCopy() *TlsConfigHttp {
	if in == nil {
		return nil
	}
	out := new(TlsConfigHttp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsConfigTransport) DeepCopyInto(out *TlsConfigTransport) {
	*out = *in
	out.CertificateConfig = in.CertificateConfig
	if in.NodesDn != nil {
		in, out := &in.NodesDn, &out.NodesDn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdminDn != nil {
		in, out := &in.AdminDn, &out.AdminDn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsConfigTransport.
func (in *TlsConfigTransport) DeepCopy() *TlsConfigTransport {
	if in == nil {
		return nil
	}
	out := new(TlsConfigTransport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsSecret) DeepCopyInto(out *TlsSecret) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsSecret.
func (in *TlsSecret) DeepCopy() *TlsSecret {
	if in == nil {
		return nil
	}
	out := new(TlsSecret)
	in.DeepCopyInto(out)
	return out
}
