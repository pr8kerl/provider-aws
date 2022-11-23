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

type BucketCorsConfigurationCorsRuleObservation struct {
}

type BucketCorsConfigurationCorsRuleParameters struct {

	// Set of Headers that are specified in the Access-Control-Request-Headers header.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Set of HTTP methods that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, and DELETE.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Set of origins you want customers to be able to access the bucket from.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Set of headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type BucketCorsConfigurationObservation struct {

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketCorsConfigurationParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Set of origins and methods (cross-origin access that you want to allow) documented below. You can configure up to 100 rules.
	// +kubebuilder:validation:Required
	CorsRule []BucketCorsConfigurationCorsRuleParameters `json:"corsRule" tf:"cors_rule,omitempty"`

	// The account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BucketCorsConfigurationSpec defines the desired state of BucketCorsConfiguration
type BucketCorsConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketCorsConfigurationParameters `json:"forProvider"`
}

// BucketCorsConfigurationStatus defines the observed state of BucketCorsConfiguration.
type BucketCorsConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketCorsConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketCorsConfiguration is the Schema for the BucketCorsConfigurations API. Provides an S3 bucket CORS configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketCorsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketCorsConfigurationSpec   `json:"spec"`
	Status            BucketCorsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketCorsConfigurationList contains a list of BucketCorsConfigurations
type BucketCorsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketCorsConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketCorsConfiguration_Kind             = "BucketCorsConfiguration"
	BucketCorsConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketCorsConfiguration_Kind}.String()
	BucketCorsConfiguration_KindAPIVersion   = BucketCorsConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketCorsConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketCorsConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketCorsConfiguration{}, &BucketCorsConfigurationList{})
}
