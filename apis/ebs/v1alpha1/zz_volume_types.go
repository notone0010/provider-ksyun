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

type VolumeObservation struct {

	// The availability zone in which the EBS volume resides.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The billing mode of the EBS volume. Valid values: 'HourlyInstantSettlement', 'Daily'.
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// The time when the EBS volume was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the KEC instance to which the EBS volume is to be attached.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies whether to expand the capacity of the EBS volume online, default is true.
	OnlineResize *bool `json:"onlineResize,omitempty" tf:"online_resize,omitempty"`

	// The ID of the project.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The capacity of the EBS volume, in GB. Value range: [10, 32000], Default is 10.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// When the cloud disk snapshot opens, the snapshot id is entered.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The category to which the EBS volume belongs. Valid values: 'system' and 'data'.
	VolumeCategory *string `json:"volumeCategory,omitempty" tf:"volume_category,omitempty"`

	// The description of the EBS volume.
	VolumeDesc *string `json:"volumeDesc,omitempty" tf:"volume_desc,omitempty"`

	// The name of the EBS volume.
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`

	// The status of the EBS volume.
	VolumeStatus *string `json:"volumeStatus,omitempty" tf:"volume_status,omitempty"`

	// The type of the EBS volume. Valid values:ESSD_PL0/ESSD_PL1/ESSD_PL2/ESSD_PL3/SSD3.0/EHDD, default is `SSD3.0`.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type VolumeParameters struct {

	// The availability zone in which the EBS volume resides.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The billing mode of the EBS volume. Valid values: 'HourlyInstantSettlement', 'Daily'.
	// +kubebuilder:validation:Optional
	ChargeType *string `json:"chargeType,omitempty" tf:"charge_type,omitempty"`

	// Specifies whether to expand the capacity of the EBS volume online, default is true.
	// +kubebuilder:validation:Optional
	OnlineResize *bool `json:"onlineResize,omitempty" tf:"online_resize,omitempty"`

	// The ID of the project.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The capacity of the EBS volume, in GB. Value range: [10, 32000], Default is 10.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// When the cloud disk snapshot opens, the snapshot id is entered.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Snapshot
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Reference to a Snapshot in ebs to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDRef *v1.Reference `json:"snapshotIdRef,omitempty" tf:"-"`

	// Selector for a Snapshot in ebs to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDSelector *v1.Selector `json:"snapshotIdSelector,omitempty" tf:"-"`

	// The description of the EBS volume.
	// +kubebuilder:validation:Optional
	VolumeDesc *string `json:"volumeDesc,omitempty" tf:"volume_desc,omitempty"`

	// The name of the EBS volume.
	// +kubebuilder:validation:Optional
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`

	// The type of the EBS volume. Valid values:ESSD_PL0/ESSD_PL1/ESSD_PL2/ESSD_PL3/SSD3.0/EHDD, default is `SSD3.0`.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Volume is the Schema for the Volumes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.availabilityZone)",message="availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.chargeType)",message="chargeType is a required parameter"
	Spec   VolumeSpec   `json:"spec"`
	Status VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
