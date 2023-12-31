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

type AutoSnapshotVolumeAssociationObservation struct {

	// The id of the volume.
	AttachVolumeID *string `json:"attachVolumeId,omitempty" tf:"attach_volume_id,omitempty"`

	// The id of the auto_snapshot_policy_id.
	AutoSnapshotPolicyID *string `json:"autoSnapshotPolicyId,omitempty" tf:"auto_snapshot_policy_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AutoSnapshotVolumeAssociationParameters struct {

	// The id of the volume.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/ebs/v1alpha1.Volume
	// +kubebuilder:validation:Optional
	AttachVolumeID *string `json:"attachVolumeId,omitempty" tf:"attach_volume_id,omitempty"`

	// Reference to a Volume in ebs to populate attachVolumeId.
	// +kubebuilder:validation:Optional
	AttachVolumeIDRef *v1.Reference `json:"attachVolumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume in ebs to populate attachVolumeId.
	// +kubebuilder:validation:Optional
	AttachVolumeIDSelector *v1.Selector `json:"attachVolumeIdSelector,omitempty" tf:"-"`

	// The id of the auto_snapshot_policy_id.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/kec/v1alpha1.AutoSnapshotPolicy
	// +kubebuilder:validation:Optional
	AutoSnapshotPolicyID *string `json:"autoSnapshotPolicyId,omitempty" tf:"auto_snapshot_policy_id,omitempty"`

	// Reference to a AutoSnapshotPolicy in kec to populate autoSnapshotPolicyId.
	// +kubebuilder:validation:Optional
	AutoSnapshotPolicyIDRef *v1.Reference `json:"autoSnapshotPolicyIdRef,omitempty" tf:"-"`

	// Selector for a AutoSnapshotPolicy in kec to populate autoSnapshotPolicyId.
	// +kubebuilder:validation:Optional
	AutoSnapshotPolicyIDSelector *v1.Selector `json:"autoSnapshotPolicyIdSelector,omitempty" tf:"-"`
}

// AutoSnapshotVolumeAssociationSpec defines the desired state of AutoSnapshotVolumeAssociation
type AutoSnapshotVolumeAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoSnapshotVolumeAssociationParameters `json:"forProvider"`
}

// AutoSnapshotVolumeAssociationStatus defines the observed state of AutoSnapshotVolumeAssociation.
type AutoSnapshotVolumeAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoSnapshotVolumeAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoSnapshotVolumeAssociation is the Schema for the AutoSnapshotVolumeAssociations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type AutoSnapshotVolumeAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoSnapshotVolumeAssociationSpec   `json:"spec"`
	Status            AutoSnapshotVolumeAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoSnapshotVolumeAssociationList contains a list of AutoSnapshotVolumeAssociations
type AutoSnapshotVolumeAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoSnapshotVolumeAssociation `json:"items"`
}

// Repository type metadata.
var (
	AutoSnapshotVolumeAssociation_Kind             = "AutoSnapshotVolumeAssociation"
	AutoSnapshotVolumeAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoSnapshotVolumeAssociation_Kind}.String()
	AutoSnapshotVolumeAssociation_KindAPIVersion   = AutoSnapshotVolumeAssociation_Kind + "." + CRDGroupVersion.String()
	AutoSnapshotVolumeAssociation_GroupVersionKind = CRDGroupVersion.WithKind(AutoSnapshotVolumeAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoSnapshotVolumeAssociation{}, &AutoSnapshotVolumeAssociationList{})
}
