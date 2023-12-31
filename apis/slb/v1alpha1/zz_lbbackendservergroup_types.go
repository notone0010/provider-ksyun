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

type HealthCheckObservation struct {

	// ID of the health check.
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	// Status maintained by health examination.Valid Values:'start', 'stop'.
	HealthCheckState *string `json:"healthCheckState,omitempty" tf:"health_check_state,omitempty"`

	// Health threshold.Valid Values:1-10. Default is 5.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// hostname of the health check.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Interval of health examination.Valid Values:1-3600. Default is 5.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Health check timeout.Valid Values:1-3600. Default is 4.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Link to HTTP type listener health check.
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	// Unhealthy threshold.Valid Values:1-10. Default is 4.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckParameters struct {

	// Status maintained by health examination.Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	HealthCheckState *string `json:"healthCheckState,omitempty" tf:"health_check_state,omitempty"`

	// Health threshold.Valid Values:1-10. Default is 5.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// hostname of the health check.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Interval of health examination.Valid Values:1-3600. Default is 5.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Health check timeout.Valid Values:1-3600. Default is 4.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Link to HTTP type listener health check.
	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	// Unhealthy threshold.Valid Values:1-10. Default is 4.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LbBackendServerGroupObservation struct {

	// ID of the backend server group.
	BackendServerGroupID *string `json:"backendServerGroupId,omitempty" tf:"backend_server_group_id,omitempty"`

	// The name of backend server group. Default: 'backend_server_group'.
	BackendServerGroupName *string `json:"backendServerGroupName,omitempty" tf:"backend_server_group_name,omitempty"`

	// The type of backend server group. Valid values: 'Server', 'Mirror'. Default is 'Server'.
	BackendServerGroupType *string `json:"backendServerGroupType,omitempty" tf:"backend_server_group_type,omitempty"`

	// number of backend servers.
	BackendServerNumber *float64 `json:"backendServerNumber,omitempty" tf:"backend_server_number,omitempty"`

	// creation time of the backend server group.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	HealthCheck []HealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type LbBackendServerGroupParameters struct {

	// The name of backend server group. Default: 'backend_server_group'.
	// +kubebuilder:validation:Optional
	BackendServerGroupName *string `json:"backendServerGroupName,omitempty" tf:"backend_server_group_name,omitempty"`

	// The type of backend server group. Valid values: 'Server', 'Mirror'. Default is 'Server'.
	// +kubebuilder:validation:Optional
	BackendServerGroupType *string `json:"backendServerGroupType,omitempty" tf:"backend_server_group_type,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// ID of the VPC.
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

// LbBackendServerGroupSpec defines the desired state of LbBackendServerGroup
type LbBackendServerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbBackendServerGroupParameters `json:"forProvider"`
}

// LbBackendServerGroupStatus defines the observed state of LbBackendServerGroup.
type LbBackendServerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbBackendServerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbBackendServerGroup is the Schema for the LbBackendServerGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type LbBackendServerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbBackendServerGroupSpec   `json:"spec"`
	Status            LbBackendServerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbBackendServerGroupList contains a list of LbBackendServerGroups
type LbBackendServerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbBackendServerGroup `json:"items"`
}

// Repository type metadata.
var (
	LbBackendServerGroup_Kind             = "LbBackendServerGroup"
	LbBackendServerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LbBackendServerGroup_Kind}.String()
	LbBackendServerGroup_KindAPIVersion   = LbBackendServerGroup_Kind + "." + CRDGroupVersion.String()
	LbBackendServerGroup_GroupVersionKind = CRDGroupVersion.WithKind(LbBackendServerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&LbBackendServerGroup{}, &LbBackendServerGroupList{})
}
