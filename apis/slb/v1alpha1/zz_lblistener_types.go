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

type LbListenerHealthCheckObservation struct {

	// ID of the health check.
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	// Status maintained by health examination.Valid Values:'start', 'stop'.
	HealthCheckState *string `json:"healthCheckState,omitempty" tf:"health_check_state,omitempty"`

	// Health threshold.Valid Values:1-10. Default is 5.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The service host name of the health check, which is available only for the HTTP or HTTPS health check.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Interval of health examination.Valid Values:1-3600. Default is 5.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Whether the host name is default or not.
	IsDefaultHostName *bool `json:"isDefaultHostName,omitempty" tf:"is_default_host_name,omitempty"`

	// The id of the listener.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// protocol of the listener.
	ListenerProtocol *string `json:"listenerProtocol,omitempty" tf:"listener_protocol,omitempty"`

	// Health check timeout.Valid Values:1-3600. Default is 4.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Link to HTTP type listener health check.
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	// Unhealthy threshold.Valid Values:1-10. Default is 4.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LbListenerHealthCheckParameters struct {

	// Status maintained by health examination.Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	HealthCheckState *string `json:"healthCheckState,omitempty" tf:"health_check_state,omitempty"`

	// Health threshold.Valid Values:1-10. Default is 5.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The service host name of the health check, which is available only for the HTTP or HTTPS health check.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Interval of health examination.Valid Values:1-3600. Default is 5.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Whether the host name is default or not.
	// +kubebuilder:validation:Optional
	IsDefaultHostName *bool `json:"isDefaultHostName,omitempty" tf:"is_default_host_name,omitempty"`

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

type LbListenerObservation struct {

	// The ID of certificate.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// whether enable to HTTP2.
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`

	// HTTP protocol, valid values:'HTTP1.0','HTTP1.1'.
	HTTPProtocol *string `json:"httpProtocol,omitempty" tf:"http_protocol,omitempty"`

	HealthCheck []LbListenerHealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of listener.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// The name of listener.
	ListenerName *string `json:"listenerName,omitempty" tf:"listener_name,omitempty"`

	// The protocol port of listener.
	ListenerPort *float64 `json:"listenerPort,omitempty" tf:"listener_port,omitempty"`

	// The protocol of listener.Valid Values:'TCP', 'UDP', 'HTTP', 'HTTPS'.
	ListenerProtocol *string `json:"listenerProtocol,omitempty" tf:"listener_protocol,omitempty"`

	// The state of listener.Valid Values:'start', 'stop'.
	ListenerState *string `json:"listenerState,omitempty" tf:"listener_state,omitempty"`

	// The ID of LB ACL.
	LoadBalancerACLID *string `json:"loadBalancerAclId,omitempty" tf:"load_balancer_acl_id,omitempty"`

	// The ID of the LB.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Forwarding mode of listener.Valid Values:'RoundRobin', 'LeastConnections', 'MasterSlave', 'QUIC_CID'.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The ID of the redirect listener.
	RedirectListenerID *string `json:"redirectListenerId,omitempty" tf:"redirect_listener_id,omitempty"`

	Session []SessionObservation `json:"session,omitempty" tf:"session,omitempty"`

	// TLS cipher policy, valid values:'TlsCipherPolicy1.0','TlsCipherPolicy1.1','TlsCipherPolicy1.2','TlsCipherPolicy1.2-strict'.
	TLSCipherPolicy *string `json:"tlsCipherPolicy,omitempty" tf:"tls_cipher_policy,omitempty"`
}

type LbListenerParameters struct {

	// The ID of certificate.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// whether enable to HTTP2.
	// +kubebuilder:validation:Optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`

	// HTTP protocol, valid values:'HTTP1.0','HTTP1.1'.
	// +kubebuilder:validation:Optional
	HTTPProtocol *string `json:"httpProtocol,omitempty" tf:"http_protocol,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheck []LbListenerHealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The name of listener.
	// +kubebuilder:validation:Optional
	ListenerName *string `json:"listenerName,omitempty" tf:"listener_name,omitempty"`

	// The protocol port of listener.
	// +kubebuilder:validation:Optional
	ListenerPort *float64 `json:"listenerPort,omitempty" tf:"listener_port,omitempty"`

	// The protocol of listener.Valid Values:'TCP', 'UDP', 'HTTP', 'HTTPS'.
	// +kubebuilder:validation:Optional
	ListenerProtocol *string `json:"listenerProtocol,omitempty" tf:"listener_protocol,omitempty"`

	// The state of listener.Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	ListenerState *string `json:"listenerState,omitempty" tf:"listener_state,omitempty"`

	// The ID of LB ACL.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbAcl
	// +kubebuilder:validation:Optional
	LoadBalancerACLID *string `json:"loadBalancerAclId,omitempty" tf:"load_balancer_acl_id,omitempty"`

	// Reference to a LbAcl in slb to populate loadBalancerAclId.
	// +kubebuilder:validation:Optional
	LoadBalancerACLIDRef *v1.Reference `json:"loadBalancerAclIdRef,omitempty" tf:"-"`

	// Selector for a LbAcl in slb to populate loadBalancerAclId.
	// +kubebuilder:validation:Optional
	LoadBalancerACLIDSelector *v1.Selector `json:"loadBalancerAclIdSelector,omitempty" tf:"-"`

	// The ID of the LB.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.Lb
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a Lb in slb to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a Lb in slb to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// Forwarding mode of listener.Valid Values:'RoundRobin', 'LeastConnections', 'MasterSlave', 'QUIC_CID'.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The ID of the redirect listener.
	// +kubebuilder:validation:Optional
	RedirectListenerID *string `json:"redirectListenerId,omitempty" tf:"redirect_listener_id,omitempty"`

	// +kubebuilder:validation:Optional
	Session []SessionParameters `json:"session,omitempty" tf:"session,omitempty"`

	// TLS cipher policy, valid values:'TlsCipherPolicy1.0','TlsCipherPolicy1.1','TlsCipherPolicy1.2','TlsCipherPolicy1.2-strict'.
	// +kubebuilder:validation:Optional
	TLSCipherPolicy *string `json:"tlsCipherPolicy,omitempty" tf:"tls_cipher_policy,omitempty"`
}

type SessionObservation struct {

	// The name of cookie.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The type of cookie, valid values: 'ImplantCookie','RewriteCookie'.
	CookieType *string `json:"cookieType,omitempty" tf:"cookie_type,omitempty"`

	// Session hold timeout.Valid Values:1-86400.
	SessionPersistencePeriod *float64 `json:"sessionPersistencePeriod,omitempty" tf:"session_persistence_period,omitempty"`

	// The state of session.Valid Values:'start', 'stop'.
	SessionState *string `json:"sessionState,omitempty" tf:"session_state,omitempty"`
}

type SessionParameters struct {

	// The name of cookie.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The type of cookie, valid values: 'ImplantCookie','RewriteCookie'.
	// +kubebuilder:validation:Optional
	CookieType *string `json:"cookieType,omitempty" tf:"cookie_type,omitempty"`

	// Session hold timeout.Valid Values:1-86400.
	// +kubebuilder:validation:Optional
	SessionPersistencePeriod *float64 `json:"sessionPersistencePeriod,omitempty" tf:"session_persistence_period,omitempty"`

	// The state of session.Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	SessionState *string `json:"sessionState,omitempty" tf:"session_state,omitempty"`
}

// LbListenerSpec defines the desired state of LbListener
type LbListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbListenerParameters `json:"forProvider"`
}

// LbListenerStatus defines the observed state of LbListener.
type LbListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbListener is the Schema for the LbListeners API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type LbListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.listenerPort)",message="listenerPort is a required parameter"
	Spec   LbListenerSpec   `json:"spec"`
	Status LbListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbListenerList contains a list of LbListeners
type LbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbListener `json:"items"`
}

// Repository type metadata.
var (
	LbListener_Kind             = "LbListener"
	LbListener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LbListener_Kind}.String()
	LbListener_KindAPIVersion   = LbListener_Kind + "." + CRDGroupVersion.String()
	LbListener_GroupVersionKind = CRDGroupVersion.WithKind(LbListener_Kind)
)

func init() {
	SchemeBuilder.Register(&LbListener{}, &LbListenerList{})
}
