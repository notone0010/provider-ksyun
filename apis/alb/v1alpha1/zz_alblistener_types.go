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

type AlbListenerObservation struct {

	// The ID of the ALB.
	ALBID *string `json:"albId,omitempty" tf:"alb_id,omitempty"`

	// The ID of listener.
	ALBListenerID *string `json:"albListenerId,omitempty" tf:"alb_listener_id,omitempty"`

	// The name of the listener.
	ALBListenerName *string `json:"albListenerName,omitempty" tf:"alb_listener_name,omitempty"`

	// The state of listener.Valid Values:'start', 'stop'.
	ALBListenerState *string `json:"albListenerState,omitempty" tf:"alb_listener_state,omitempty"`

	// The ID of certificate.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// whether enable to HTTP2.
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`

	// Backend Protocol, valid values:'HTTP1.0','HTTP1.1'.
	HTTPProtocol *string `json:"httpProtocol,omitempty" tf:"http_protocol,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Forwarding mode of listener. Valid Values:'RoundRobin', 'LeastConnections'.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The protocol port of listener.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol of listener. Valid Values: 'HTTP', 'HTTPS'.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The ID of the redirect ALB listener.
	RedirectALBListenerID *string `json:"redirectAlbListenerId,omitempty" tf:"redirect_alb_listener_id,omitempty"`

	Session []SessionObservation `json:"session,omitempty" tf:"session,omitempty"`

	// TLS cipher policy, valid values:'TlsCipherPolicy1.0','TlsCipherPolicy1.1','TlsCipherPolicy1.2','TlsCipherPolicy1.2-strict','TlsCipherPolicy1.2-moststrict'.
	TLSCipherPolicy *string `json:"tlsCipherPolicy,omitempty" tf:"tls_cipher_policy,omitempty"`
}

type AlbListenerParameters struct {

	// The ID of the ALB.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/alb/v1alpha1.Alb
	// +kubebuilder:validation:Optional
	ALBID *string `json:"albId,omitempty" tf:"alb_id,omitempty"`

	// Reference to a Alb in alb to populate albId.
	// +kubebuilder:validation:Optional
	ALBIDRef *v1.Reference `json:"albIdRef,omitempty" tf:"-"`

	// Selector for a Alb in alb to populate albId.
	// +kubebuilder:validation:Optional
	ALBIDSelector *v1.Selector `json:"albIdSelector,omitempty" tf:"-"`

	// The name of the listener.
	// +kubebuilder:validation:Optional
	ALBListenerName *string `json:"albListenerName,omitempty" tf:"alb_listener_name,omitempty"`

	// The state of listener.Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	ALBListenerState *string `json:"albListenerState,omitempty" tf:"alb_listener_state,omitempty"`

	// The ID of certificate.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// whether enable to HTTP2.
	// +kubebuilder:validation:Optional
	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`

	// Backend Protocol, valid values:'HTTP1.0','HTTP1.1'.
	// +kubebuilder:validation:Optional
	HTTPProtocol *string `json:"httpProtocol,omitempty" tf:"http_protocol,omitempty"`

	// Forwarding mode of listener. Valid Values:'RoundRobin', 'LeastConnections'.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The protocol port of listener.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol of listener. Valid Values: 'HTTP', 'HTTPS'.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The ID of the redirect ALB listener.
	// +kubebuilder:validation:Optional
	RedirectALBListenerID *string `json:"redirectAlbListenerId,omitempty" tf:"redirect_alb_listener_id,omitempty"`

	// +kubebuilder:validation:Optional
	Session []SessionParameters `json:"session,omitempty" tf:"session,omitempty"`

	// TLS cipher policy, valid values:'TlsCipherPolicy1.0','TlsCipherPolicy1.1','TlsCipherPolicy1.2','TlsCipherPolicy1.2-strict','TlsCipherPolicy1.2-moststrict'.
	// +kubebuilder:validation:Optional
	TLSCipherPolicy *string `json:"tlsCipherPolicy,omitempty" tf:"tls_cipher_policy,omitempty"`
}

type SessionObservation struct {

	// The name of cookie.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The type of cookie, valid values: 'ImplantCookie','RewriteCookie'.
	CookieType *string `json:"cookieType,omitempty" tf:"cookie_type,omitempty"`

	// Session hold timeout. Valid Values:1-86400.
	SessionPersistencePeriod *float64 `json:"sessionPersistencePeriod,omitempty" tf:"session_persistence_period,omitempty"`

	// The state of session. Valid Values:'start', 'stop'.
	SessionState *string `json:"sessionState,omitempty" tf:"session_state,omitempty"`
}

type SessionParameters struct {

	// The name of cookie.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The type of cookie, valid values: 'ImplantCookie','RewriteCookie'.
	// +kubebuilder:validation:Optional
	CookieType *string `json:"cookieType,omitempty" tf:"cookie_type,omitempty"`

	// Session hold timeout. Valid Values:1-86400.
	// +kubebuilder:validation:Optional
	SessionPersistencePeriod *float64 `json:"sessionPersistencePeriod,omitempty" tf:"session_persistence_period,omitempty"`

	// The state of session. Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	SessionState *string `json:"sessionState,omitempty" tf:"session_state,omitempty"`
}

// AlbListenerSpec defines the desired state of AlbListener
type AlbListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlbListenerParameters `json:"forProvider"`
}

// AlbListenerStatus defines the observed state of AlbListener.
type AlbListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlbListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListener is the Schema for the AlbListeners API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type AlbListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.port)",message="port is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.protocol)",message="protocol is a required parameter"
	Spec   AlbListenerSpec   `json:"spec"`
	Status AlbListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerList contains a list of AlbListeners
type AlbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListener `json:"items"`
}

// Repository type metadata.
var (
	AlbListener_Kind             = "AlbListener"
	AlbListener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlbListener_Kind}.String()
	AlbListener_KindAPIVersion   = AlbListener_Kind + "." + CRDGroupVersion.String()
	AlbListener_GroupVersionKind = CRDGroupVersion.WithKind(AlbListener_Kind)
)

func init() {
	SchemeBuilder.Register(&AlbListener{}, &AlbListenerList{})
}
