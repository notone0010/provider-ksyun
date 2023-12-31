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

type BwsAssociateObservation struct {

	// ID of the EIP.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// bandwidth value.
	BandWidth *float64 `json:"bandWidth,omitempty" tf:"band_width,omitempty"`

	// ID of the BWS.
	BandWidthShareID *string `json:"bandWidthShareId,omitempty" tf:"band_width_share_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BwsAssociateParameters struct {

	// ID of the EIP.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/eip/v1alpha1.EIP
	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// Reference to a EIP in eip to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDRef *v1.Reference `json:"allocationIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDSelector *v1.Selector `json:"allocationIdSelector,omitempty" tf:"-"`

	// ID of the BWS.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/eip/v1alpha1.Bws
	// +kubebuilder:validation:Optional
	BandWidthShareID *string `json:"bandWidthShareId,omitempty" tf:"band_width_share_id,omitempty"`

	// Reference to a Bws in eip to populate bandWidthShareId.
	// +kubebuilder:validation:Optional
	BandWidthShareIDRef *v1.Reference `json:"bandWidthShareIdRef,omitempty" tf:"-"`

	// Selector for a Bws in eip to populate bandWidthShareId.
	// +kubebuilder:validation:Optional
	BandWidthShareIDSelector *v1.Selector `json:"bandWidthShareIdSelector,omitempty" tf:"-"`
}

// BwsAssociateSpec defines the desired state of BwsAssociate
type BwsAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BwsAssociateParameters `json:"forProvider"`
}

// BwsAssociateStatus defines the observed state of BwsAssociate.
type BwsAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BwsAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BwsAssociate is the Schema for the BwsAssociates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type BwsAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BwsAssociateSpec   `json:"spec"`
	Status            BwsAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BwsAssociateList contains a list of BwsAssociates
type BwsAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BwsAssociate `json:"items"`
}

// Repository type metadata.
var (
	BwsAssociate_Kind             = "BwsAssociate"
	BwsAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BwsAssociate_Kind}.String()
	BwsAssociate_KindAPIVersion   = BwsAssociate_Kind + "." + CRDGroupVersion.String()
	BwsAssociate_GroupVersionKind = CRDGroupVersion.WithKind(BwsAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&BwsAssociate{}, &BwsAssociateList{})
}
