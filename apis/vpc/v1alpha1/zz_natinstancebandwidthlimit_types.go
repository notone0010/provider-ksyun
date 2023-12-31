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

type NatInstanceBandwidthLimitObservation struct {

	// The bandwidth limit of network interface that belong to instance.
	BandwidthLimit *float64 `json:"bandwidthLimit,omitempty" tf:"bandwidth_limit,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// the type of instance. Values: epc or kec.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The id of the Nat.
	NATID *string `json:"natId,omitempty" tf:"nat_id,omitempty"`

	// The id of network interface that belong to instance.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// the private ip of network interface.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`
}

type NatInstanceBandwidthLimitParameters struct {

	// The bandwidth limit of network interface that belong to instance.
	// +kubebuilder:validation:Optional
	BandwidthLimit *float64 `json:"bandwidthLimit,omitempty" tf:"bandwidth_limit,omitempty"`

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

	// The id of network interface that belong to instance.
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

// NatInstanceBandwidthLimitSpec defines the desired state of NatInstanceBandwidthLimit
type NatInstanceBandwidthLimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NatInstanceBandwidthLimitParameters `json:"forProvider"`
}

// NatInstanceBandwidthLimitStatus defines the observed state of NatInstanceBandwidthLimit.
type NatInstanceBandwidthLimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NatInstanceBandwidthLimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NatInstanceBandwidthLimit is the Schema for the NatInstanceBandwidthLimits API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type NatInstanceBandwidthLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bandwidthLimit)",message="bandwidthLimit is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.networkInterfaceId)",message="networkInterfaceId is a required parameter"
	Spec   NatInstanceBandwidthLimitSpec   `json:"spec"`
	Status NatInstanceBandwidthLimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatInstanceBandwidthLimitList contains a list of NatInstanceBandwidthLimits
type NatInstanceBandwidthLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatInstanceBandwidthLimit `json:"items"`
}

// Repository type metadata.
var (
	NatInstanceBandwidthLimit_Kind             = "NatInstanceBandwidthLimit"
	NatInstanceBandwidthLimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NatInstanceBandwidthLimit_Kind}.String()
	NatInstanceBandwidthLimit_KindAPIVersion   = NatInstanceBandwidthLimit_Kind + "." + CRDGroupVersion.String()
	NatInstanceBandwidthLimit_GroupVersionKind = CRDGroupVersion.WithKind(NatInstanceBandwidthLimit_Kind)
)

func init() {
	SchemeBuilder.Register(&NatInstanceBandwidthLimit{}, &NatInstanceBandwidthLimitList{})
}
