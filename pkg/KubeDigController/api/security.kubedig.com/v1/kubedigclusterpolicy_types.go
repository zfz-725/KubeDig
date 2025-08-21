// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of KubeDig

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MatchExpressionsType struct {
	// +kubebuilder:validation:Enum=namespace
	Key string `json:"key,omitempty"`

	// +kubebuilder:validation:Enum=In;NotIn
	Operator string `json:"operator,omitempty"`

	Values []string `json:"values,omitempty"`
}

type NsSelectorType struct {
	MatchExpressions []MatchExpressionsType `json:"matchExpressions,omitempty"`
}

// KubeDigClusterPolicySpec defines the desired state of KubeDigClusterPolicy
type KubeDigClusterPolicySpec struct {
	Selector NsSelectorType `json:"selector,omitempty"`

	Process      ProcessType      `json:"process,omitempty"`
	File         FileType         `json:"file,omitempty"`
	Network      NetworkType      `json:"network,omitempty"`
	Capabilities CapabilitiesType `json:"capabilities,omitempty"`
	Syscalls     SyscallsType     `json:"syscalls,omitempty"`

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

// KubeDigClusterPolicyStatus defines the observed state of KubeDigCLusterPolicy
type KubeDigClusterPolicyStatus struct {
	PolicyStatus string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeDigClusterPolicy is the Schema for the kubedigclusterpolicies API
// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:shortName=csp,scope="Cluster"
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Action",type=string,JSONPath=`.spec.action`,priority=10
// +kubebuilder:printcolumn:name="Selector",type=string,JSONPath=`.spec.selector.matchExpressions`,priority=10
type KubeDigClusterPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeDigClusterPolicySpec   `json:"spec,omitempty"`
	Status KubeDigClusterPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeDigClusterPolicyList contains a list of KubeDigClusterPolicy
type KubeDigClusterPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeDigClusterPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeDigClusterPolicy{}, &KubeDigClusterPolicyList{})
}
