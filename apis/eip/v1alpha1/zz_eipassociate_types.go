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

type EipAssociateObservation struct {

	// The ID of EIP.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// The band width of the public address.
	BandWidth *float64 `json:"bandWidth,omitempty" tf:"band_width,omitempty"`

	// the ID of the BWS which the EIP associated.
	BandWidthShareID *string `json:"bandWidthShareId,omitempty" tf:"band_width_share_id,omitempty"`

	// creation time of the EIP.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP version of the EIP.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The id of the instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The type of the instance.Valid Values:'Ipfwd', 'Slb'.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// InternetGateway ID.
	InternetGatewayID *string `json:"internetGatewayId,omitempty" tf:"internet_gateway_id,omitempty"`

	// BWS EIP.
	IsBandWidthShare *string `json:"isBandWidthShare,omitempty" tf:"is_band_width_share,omitempty"`

	// The id of the line.
	LineID *string `json:"lineId,omitempty" tf:"line_id,omitempty"`

	// The id of the network interface.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The id of the project.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The Elastic IP address.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// state of the EIP.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type EipAssociateParameters struct {

	// The ID of EIP.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/eip/v1alpha1.EIP
	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// Reference to a EIP in eip to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDRef *v1.Reference `json:"allocationIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDSelector *v1.Selector `json:"allocationIdSelector,omitempty" tf:"-"`

	// The id of the instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The type of the instance.Valid Values:'Ipfwd', 'Slb'.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The id of the network interface.
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

// EipAssociateSpec defines the desired state of EipAssociate
type EipAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EipAssociateParameters `json:"forProvider"`
}

// EipAssociateStatus defines the observed state of EipAssociate.
type EipAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EipAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociate is the Schema for the EipAssociates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type EipAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceId)",message="instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceType)",message="instanceType is a required parameter"
	Spec   EipAssociateSpec   `json:"spec"`
	Status EipAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociateList contains a list of EipAssociates
type EipAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EipAssociate `json:"items"`
}

// Repository type metadata.
var (
	EipAssociate_Kind             = "EipAssociate"
	EipAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EipAssociate_Kind}.String()
	EipAssociate_KindAPIVersion   = EipAssociate_Kind + "." + CRDGroupVersion.String()
	EipAssociate_GroupVersionKind = CRDGroupVersion.WithKind(EipAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&EipAssociate{}, &EipAssociateList{})
}
