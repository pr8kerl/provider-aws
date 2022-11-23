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

type ConfigurationObservation struct {

	// Configuration block for the Athena Engine Versioning. For more information, see Athena Engine Versioning. See Engine Version below.
	// +kubebuilder:validation:Optional
	EngineVersion []EngineVersionObservation `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
}

type ConfigurationParameters struct {

	// Integer for the upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan. Must be at least 10485760.
	// +kubebuilder:validation:Optional
	BytesScannedCutoffPerQuery *float64 `json:"bytesScannedCutoffPerQuery,omitempty" tf:"bytes_scanned_cutoff_per_query,omitempty"`

	// Boolean whether the settings for the workgroup override client-side settings. For more information, see Workgroup Settings Override Client-Side Settings. Defaults to true.
	// +kubebuilder:validation:Optional
	EnforceWorkgroupConfiguration *bool `json:"enforceWorkgroupConfiguration,omitempty" tf:"enforce_workgroup_configuration,omitempty"`

	// Configuration block for the Athena Engine Versioning. For more information, see Athena Engine Versioning. See Engine Version below.
	// +kubebuilder:validation:Optional
	EngineVersion []EngineVersionParameters `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Boolean whether Amazon CloudWatch metrics are enabled for the workgroup. Defaults to true.
	// +kubebuilder:validation:Optional
	PublishCloudwatchMetricsEnabled *bool `json:"publishCloudwatchMetricsEnabled,omitempty" tf:"publish_cloudwatch_metrics_enabled,omitempty"`

	// If set to true , allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false , workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. The default is false . For more information about Requester Pays buckets, see Requester Pays Buckets in the Amazon Simple Storage Service Developer Guide.
	// +kubebuilder:validation:Optional
	RequesterPaysEnabled *bool `json:"requesterPaysEnabled,omitempty" tf:"requester_pays_enabled,omitempty"`

	// Configuration block with result settings. See Result Configuration below.
	// +kubebuilder:validation:Optional
	ResultConfiguration []ResultConfigurationParameters `json:"resultConfiguration,omitempty" tf:"result_configuration,omitempty"`
}

type EngineVersionObservation struct {

	// The engine version on which the query runs. If selected_engine_version is set to AUTO, the effective engine version is chosen by Athena.
	EffectiveEngineVersion *string `json:"effectiveEngineVersion,omitempty" tf:"effective_engine_version,omitempty"`
}

type EngineVersionParameters struct {

	// Requested engine version. Defaults to AUTO.
	// +kubebuilder:validation:Optional
	SelectedEngineVersion *string `json:"selectedEngineVersion,omitempty" tf:"selected_engine_version,omitempty"`
}

type ResultConfigurationACLConfigurationObservation struct {
}

type ResultConfigurationACLConfigurationParameters struct {

	// Amazon S3 canned ACL that Athena should specify when storing query results. Valid value is BUCKET_OWNER_FULL_CONTROL.
	// +kubebuilder:validation:Required
	S3ACLOption *string `json:"s3AclOption" tf:"s3_acl_option,omitempty"`
}

type ResultConfigurationEncryptionConfigurationObservation struct {
}

type ResultConfigurationEncryptionConfigurationParameters struct {

	// Whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used. If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
	// +kubebuilder:validation:Optional
	EncryptionOption *string `json:"encryptionOption,omitempty" tf:"encryption_option,omitempty"`

	// For SSE_KMS and CSE_KMS, this is the KMS key ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`
}

type ResultConfigurationObservation struct {
}

type ResultConfigurationParameters struct {

	// That an Amazon S3 canned ACL should be set to control ownership of stored query results. See ACL Configuration below.
	// +kubebuilder:validation:Optional
	ACLConfiguration []ResultConfigurationACLConfigurationParameters `json:"aclConfiguration,omitempty" tf:"acl_configuration,omitempty"`

	// Configuration block with encryption settings. See Encryption Configuration below.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []ResultConfigurationEncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// AWS account ID that you expect to be the owner of the Amazon S3 bucket.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. For more information, see Queries and Query Result Files.
	// +kubebuilder:validation:Optional
	OutputLocation *string `json:"outputLocation,omitempty" tf:"output_location,omitempty"`
}

type WorkgroupObservation struct {

	// ARN of the workgroup
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block with various settings for the workgroup. Documented below.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Workgroup name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkgroupParameters struct {

	// Configuration block with various settings for the workgroup. Documented below.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Description of the workgroup.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// State of the workgroup. Valid values are DISABLED or ENABLED. Defaults to ENABLED.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkgroupSpec defines the desired state of Workgroup
type WorkgroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkgroupParameters `json:"forProvider"`
}

// WorkgroupStatus defines the observed state of Workgroup.
type WorkgroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workgroup is the Schema for the Workgroups API. Manages an Athena Workgroup.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Workgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkgroupSpec   `json:"spec"`
	Status            WorkgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkgroupList contains a list of Workgroups
type WorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workgroup `json:"items"`
}

// Repository type metadata.
var (
	Workgroup_Kind             = "Workgroup"
	Workgroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workgroup_Kind}.String()
	Workgroup_KindAPIVersion   = Workgroup_Kind + "." + CRDGroupVersion.String()
	Workgroup_GroupVersionKind = CRDGroupVersion.WithKind(Workgroup_Kind)
)

func init() {
	SchemeBuilder.Register(&Workgroup{}, &WorkgroupList{})
}
