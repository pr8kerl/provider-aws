/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionObservation struct {

	// The ARN of the connection.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy,omitempty"`

	// The ID of the connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	// Boolean value indicating whether the connection supports MAC Security (MACsec).
	MacsecCapable *bool `json:"macsecCapable,omitempty" tf:"macsec_capable,omitempty"`

	// The ID of the AWS account that owns the connection.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// The MAC Security (MACsec) port link status of the connection.
	PortEncryptionStatus *string `json:"portEncryptionStatus,omitempty" tf:"port_encryption_status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VLAN ID.
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type ConnectionParameters struct {

	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	// +kubebuilder:validation:Required
	Bandwidth *string `json:"bandwidth" tf:"bandwidth,omitempty"`

	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are no_encrypt, should_encrypt, and must_encrypt.
	// +kubebuilder:validation:Optional
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// The AWS Direct Connect location where the connection is located. See DescribeLocations for the list of AWS Direct Connect locations. Use locationCode.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the connection.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the service provider associated with the connection.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See MACsec prerequisites for more information about MAC Security (MACsec) prerequisites. Default value: false.
	// +kubebuilder:validation:Optional
	RequestMacsec *bool `json:"requestMacsec,omitempty" tf:"request_macsec,omitempty"`

	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Provides a Connection of Direct Connect.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
