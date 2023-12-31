/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NatAssociateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the Nat.
	NATID *string `json:"natId,omitempty" tf:"nat_id,omitempty"`

	// The id of network interface that belong to instance. Notes: Because of there is one resource in the association, conflict with `subnet_id`.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The id of the Subnet. Notes: Because of there is one resource in the association, conflict with `network_interface_id`.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NatAssociateParameters struct {

	// The id of the Nat.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Nat
	// +kubebuilder:validation:Optional
	NATID *string `json:"natId,omitempty" tf:"nat_id,omitempty"`

	// Reference to a Nat in vpc to populate natId.
	// +kubebuilder:validation:Optional
	NATIDRef *v1.Reference `json:"natIdRef,omitempty" tf:"-"`

	// Selector for a Nat in vpc to populate natId.
	// +kubebuilder:validation:Optional
	NATIDSelector *v1.Selector `json:"natIdSelector,omitempty" tf:"-"`

	// The id of network interface that belong to instance. Notes: Because of there is one resource in the association, conflict with `subnet_id`.
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The id of the Subnet. Notes: Because of there is one resource in the association, conflict with `network_interface_id`.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// NatAssociateSpec defines the desired state of NatAssociate
type NatAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NatAssociateParameters `json:"forProvider"`
}

// NatAssociateStatus defines the observed state of NatAssociate.
type NatAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NatAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NatAssociate is the Schema for the NatAssociates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type NatAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatAssociateSpec   `json:"spec"`
	Status            NatAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatAssociateList contains a list of NatAssociates
type NatAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatAssociate `json:"items"`
}

// Repository type metadata.
var (
	NatAssociate_Kind             = "NatAssociate"
	NatAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NatAssociate_Kind}.String()
	NatAssociate_KindAPIVersion   = NatAssociate_Kind + "." + CRDGroupVersion.String()
	NatAssociate_GroupVersionKind = CRDGroupVersion.WithKind(NatAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&NatAssociate{}, &NatAssociateList{})
}
