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

type SecurityGroupEntriesObservation struct {

	// The cidr block of security group rule.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The description of the entry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction of the entry, valid values:'in', 'out'.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// ICMP code.The required if protocol type is 'icmp'.
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// ICMP type.The required if protocol type is 'icmp'.
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// Port rule start port for TCP or UDP protocol.The required if protocol type is 'tcp' or 'udp'.
	PortRangeFrom *float64 `json:"portRangeFrom,omitempty" tf:"port_range_from,omitempty"`

	// Port rule end port for TCP or UDP protocol.The required if protocol type is 'tcp' or 'udp'.
	PortRangeTo *float64 `json:"portRangeTo,omitempty" tf:"port_range_to,omitempty"`

	// The protocol of the entry, valid values: 'ip', 'tcp', 'udp', 'icmp'.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The ID of the entry.
	SecurityGroupEntryID *string `json:"securityGroupEntryId,omitempty" tf:"security_group_entry_id,omitempty"`
}

type SecurityGroupEntriesParameters struct {

	// The cidr block of security group rule.
	// +kubebuilder:validation:Required
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block,omitempty"`

	// The description of the entry.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction of the entry, valid values:'in', 'out'.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// ICMP code.The required if protocol type is 'icmp'.
	// +kubebuilder:validation:Optional
	IcmpCode *float64 `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// ICMP type.The required if protocol type is 'icmp'.
	// +kubebuilder:validation:Optional
	IcmpType *float64 `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// Port rule start port for TCP or UDP protocol.The required if protocol type is 'tcp' or 'udp'.
	// +kubebuilder:validation:Optional
	PortRangeFrom *float64 `json:"portRangeFrom,omitempty" tf:"port_range_from,omitempty"`

	// Port rule end port for TCP or UDP protocol.The required if protocol type is 'tcp' or 'udp'.
	// +kubebuilder:validation:Optional
	PortRangeTo *float64 `json:"portRangeTo,omitempty" tf:"port_range_to,omitempty"`

	// The protocol of the entry, valid values: 'ip', 'tcp', 'udp', 'icmp'.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type SecurityGroupObservation struct {

	// The time of creation of security group.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SecurityGroupEntries []SecurityGroupEntriesObservation `json:"securityGroupEntries,omitempty" tf:"security_group_entries,omitempty"`

	// The ID of the security group.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// The name of the security group.
	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`

	// The Id of the vpc.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SecurityGroupParameters struct {

	// +kubebuilder:validation:Optional
	SecurityGroupEntries []SecurityGroupEntriesParameters `json:"securityGroupEntries,omitempty" tf:"security_group_entries,omitempty"`

	// The name of the security group.
	// +kubebuilder:validation:Optional
	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`

	// The Id of the vpc.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}