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

type VolumeAttachObservation struct {

	// The availability zone in which the EBS volume resides.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The time when the EBS volume was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Specifies whether to delete the EBS volume when the KEC instance to which it is attached is deleted. Default value: false.
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty" tf:"delete_with_instance,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the KEC instance to which the EBS volume is to be attached.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the project.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The capacity of the EBS volume, in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The category of the EBS volume.
	VolumeCategory *string `json:"volumeCategory,omitempty" tf:"volume_category,omitempty"`

	// The description of the EBS volume.
	VolumeDesc *string `json:"volumeDesc,omitempty" tf:"volume_desc,omitempty"`

	// The ID of the EBS volume.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// The name of the EBS volume.
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`

	// The status of the EBS volume.
	VolumeStatus *string `json:"volumeStatus,omitempty" tf:"volume_status,omitempty"`

	// The type of the EBS volume.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type VolumeAttachParameters struct {

	// Specifies whether to delete the EBS volume when the KEC instance to which it is attached is deleted. Default value: false.
	// +kubebuilder:validation:Optional
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty" tf:"delete_with_instance,omitempty"`

	// The ID of the KEC instance to which the EBS volume is to be attached.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the EBS volume.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Volume
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume in ebs to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume in ebs to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

// VolumeAttachSpec defines the desired state of VolumeAttach
type VolumeAttachSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeAttachParameters `json:"forProvider"`
}

// VolumeAttachStatus defines the observed state of VolumeAttach.
type VolumeAttachStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeAttachObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttach is the Schema for the VolumeAttachs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type VolumeAttach struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceId)",message="instanceId is a required parameter"
	Spec   VolumeAttachSpec   `json:"spec"`
	Status VolumeAttachStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachList contains a list of VolumeAttachs
type VolumeAttachList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeAttach `json:"items"`
}

// Repository type metadata.
var (
	VolumeAttach_Kind             = "VolumeAttach"
	VolumeAttach_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeAttach_Kind}.String()
	VolumeAttach_KindAPIVersion   = VolumeAttach_Kind + "." + CRDGroupVersion.String()
	VolumeAttach_GroupVersionKind = CRDGroupVersion.WithKind(VolumeAttach_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeAttach{}, &VolumeAttachList{})
}
