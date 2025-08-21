// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubeDigHostPolicySpec defines the desired state of KubeDigHostPolicy
type KubeDigHostPolicySpec struct {
	NodeSelector NodeSelectorType `json:"nodeSelector"`

	Process      ProcessType          `json:"process,omitempty"`
	File         FileType             `json:"file,omitempty"`
	Network      HostNetworkType      `json:"network,omitempty"`
	Capabilities HostCapabilitiesType `json:"capabilities,omitempty"`
	Syscalls     SyscallsType         `json:"syscalls,omitempty"`

	AppArmor string `json:"apparmor,omitempty"`

	// +kubebuilder:validation:optional
	Severity SeverityType `json:"severity,omitempty"`
	// +kubebuilder:validation:optional
	Tags []string `json:"tags,omitempty"`
	// +kubebuilder:validation:optional
	Message string `json:"message,omitempty"`
	// +kubebuilder:validation:optional
	Action ActionType `json:"action,omitempty"`
}

// KubeDigHostPolicyStatus defines the observed state of KubeDigHostPolicy
type KubeDigHostPolicyStatus struct {
	PolicyStatus string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeDigHostPolicy is the Schema for the kubedighostpolicies API
// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster,shortName=hsp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Action",type=string,JSONPath=`.spec.action`,priority=10
// +kubebuilder:printcolumn:name="Selector",type=string,JSONPath=`.spec.nodeSelector.matchLabels`,priority=10
type KubeDigHostPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeDigHostPolicySpec   `json:"spec,omitempty"`
	Status KubeDigHostPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeDigHostPolicyList contains a list of KubeDigHostPolicy
type KubeDigHostPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeDigHostPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeDigHostPolicy{}, &KubeDigHostPolicyList{})
}
