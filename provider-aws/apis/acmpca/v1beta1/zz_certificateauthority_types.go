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

type CertificateAuthorityConfigurationObservation struct {
}

type CertificateAuthorityConfigurationParameters struct {

	// +kubebuilder:validation:Required
	KeyAlgorithm *string `json:"keyAlgorithm" tf:"key_algorithm,omitempty"`

	// +kubebuilder:validation:Required
	SigningAlgorithm *string `json:"signingAlgorithm" tf:"signing_algorithm,omitempty"`

	// +kubebuilder:validation:Required
	Subject []SubjectParameters `json:"subject" tf:"subject,omitempty"`
}

type CertificateAuthorityObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	CertificateSigningRequest *string `json:"certificateSigningRequest,omitempty" tf:"certificate_signing_request,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NotAfter *string `json:"notAfter,omitempty" tf:"not_after,omitempty"`

	NotBefore *string `json:"notBefore,omitempty" tf:"not_before,omitempty"`

	Serial *string `json:"serial,omitempty" tf:"serial,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CertificateAuthorityParameters struct {

	// +kubebuilder:validation:Required
	CertificateAuthorityConfiguration []CertificateAuthorityConfigurationParameters `json:"certificateAuthorityConfiguration" tf:"certificate_authority_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PermanentDeletionTimeInDays *float64 `json:"permanentDeletionTimeInDays,omitempty" tf:"permanent_deletion_time_in_days,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RevocationConfiguration []RevocationConfigurationParameters `json:"revocationConfiguration,omitempty" tf:"revocation_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CrlConfigurationObservation struct {
}

type CrlConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CustomCname *string `json:"customCname,omitempty" tf:"custom_cname,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	ExpirationInDays *float64 `json:"expirationInDays" tf:"expiration_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	S3ObjectACL *string `json:"s3ObjectAcl,omitempty" tf:"s3_object_acl,omitempty"`
}

type RevocationConfigurationObservation struct {
}

type RevocationConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CrlConfiguration []CrlConfigurationParameters `json:"crlConfiguration,omitempty" tf:"crl_configuration,omitempty"`
}

type SubjectObservation struct {
}

type SubjectParameters struct {

	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// +kubebuilder:validation:Optional
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier,omitempty" tf:"distinguished_name_qualifier,omitempty"`

	// +kubebuilder:validation:Optional
	GenerationQualifier *string `json:"generationQualifier,omitempty" tf:"generation_qualifier,omitempty"`

	// +kubebuilder:validation:Optional
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// +kubebuilder:validation:Optional
	Initials *string `json:"initials,omitempty" tf:"initials,omitempty"`

	// +kubebuilder:validation:Optional
	Locality *string `json:"locality,omitempty" tf:"locality,omitempty"`

	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`

	// +kubebuilder:validation:Optional
	Pseudonym *string `json:"pseudonym,omitempty" tf:"pseudonym,omitempty"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// CertificateAuthoritySpec defines the desired state of CertificateAuthority
type CertificateAuthoritySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateAuthorityParameters `json:"forProvider"`
}

// CertificateAuthorityStatus defines the observed state of CertificateAuthority.
type CertificateAuthorityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateAuthorityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthority is the Schema for the CertificateAuthoritys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateAuthoritySpec   `json:"spec"`
	Status            CertificateAuthorityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthorityList contains a list of CertificateAuthoritys
type CertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateAuthority `json:"items"`
}

// Repository type metadata.
var (
	CertificateAuthority_Kind             = "CertificateAuthority"
	CertificateAuthority_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateAuthority_Kind}.String()
	CertificateAuthority_KindAPIVersion   = CertificateAuthority_Kind + "." + CRDGroupVersion.String()
	CertificateAuthority_GroupVersionKind = CRDGroupVersion.WithKind(CertificateAuthority_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateAuthority{}, &CertificateAuthorityList{})
}
