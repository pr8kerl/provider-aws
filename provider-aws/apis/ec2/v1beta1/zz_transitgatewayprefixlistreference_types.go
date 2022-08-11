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

type TransitGatewayPrefixListReferenceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PrefixListOwnerID *string `json:"prefixListOwnerId,omitempty" tf:"prefix_list_owner_id,omitempty"`
}

type TransitGatewayPrefixListReferenceParameters struct {

	// +kubebuilder:validation:Optional
	Blackhole *bool `json:"blackhole,omitempty" tf:"blackhole,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// +kubebuilder:validation:Optional
	PrefixListIDRef *v1.Reference `json:"prefixListIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrefixListIDSelector *v1.Selector `json:"prefixListIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.TransitGatewayVPCAttachment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("association_default_route_table_id",true)
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDRef *v1.Reference `json:"transitGatewayRouteTableIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDSelector *v1.Selector `json:"transitGatewayRouteTableIdSelector,omitempty" tf:"-"`
}

// TransitGatewayPrefixListReferenceSpec defines the desired state of TransitGatewayPrefixListReference
type TransitGatewayPrefixListReferenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPrefixListReferenceParameters `json:"forProvider"`
}

// TransitGatewayPrefixListReferenceStatus defines the observed state of TransitGatewayPrefixListReference.
type TransitGatewayPrefixListReferenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPrefixListReferenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPrefixListReference is the Schema for the TransitGatewayPrefixListReferences API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayPrefixListReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayPrefixListReferenceSpec   `json:"spec"`
	Status            TransitGatewayPrefixListReferenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPrefixListReferenceList contains a list of TransitGatewayPrefixListReferences
type TransitGatewayPrefixListReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPrefixListReference `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPrefixListReference_Kind             = "TransitGatewayPrefixListReference"
	TransitGatewayPrefixListReference_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayPrefixListReference_Kind}.String()
	TransitGatewayPrefixListReference_KindAPIVersion   = TransitGatewayPrefixListReference_Kind + "." + CRDGroupVersion.String()
	TransitGatewayPrefixListReference_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayPrefixListReference_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPrefixListReference{}, &TransitGatewayPrefixListReferenceList{})
}
