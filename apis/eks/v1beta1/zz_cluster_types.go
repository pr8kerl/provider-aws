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

type CertificateAuthorityObservation struct {

	// Base64 encoded certificate data required to communicate with your cluster. Add this to the certificate-authority-data section of the kubeconfig file for your cluster.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`
}

type CertificateAuthorityParameters struct {
}

type ClusterObservation struct {

	// ARN of the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Attribute block containing certificate-authority-data for your cluster. Detailed below.
	CertificateAuthority []CertificateAuthorityObservation `json:"certificateAuthority,omitempty" tf:"certificate_authority,omitempty"`

	// The ID of your local Amazon EKS cluster on the AWS Outpost. This attribute isn't available for an AWS EKS cluster on AWS cloud.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Unix epoch timestamp in seconds for when the cluster was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Endpoint for your Kubernetes API server.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Name of the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Configuration block with kubernetes network configuration for the cluster. Detailed below.
	// +kubebuilder:validation:Optional
	KubernetesNetworkConfig []KubernetesNetworkConfigObservation `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config,omitempty"`

	// Platform version for the cluster.
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`

	// Status of the EKS cluster. One of CREATING, ACTIVE, DELETING, FAILED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see Cluster VPC Considerations and Cluster Security Group Considerations in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
	// +kubebuilder:validation:Required
	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ClusterParameters struct {

	// List of the desired control plane logging to enable. For more information, see Amazon EKS Control Plane Logging.
	// +kubebuilder:validation:Optional
	EnabledClusterLogTypes []*string `json:"enabledClusterLogTypes,omitempty" tf:"enabled_cluster_log_types,omitempty"`

	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	// +kubebuilder:validation:Optional
	EncryptionConfig []EncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// Configuration block with kubernetes network configuration for the cluster. Detailed below.
	// +kubebuilder:validation:Optional
	KubernetesNetworkConfig []KubernetesNetworkConfigParameters `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config,omitempty"`

	// Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn't available for creating Amazon EKS clusters on the AWS cloud.
	// +kubebuilder:validation:Optional
	OutpostConfig []OutpostConfigParameters `json:"outpostConfig,omitempty" tf:"outpost_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding depends_on if using the aws_iam_role_policy resource or aws_iam_role_policy_attachment resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see Cluster VPC Considerations and Cluster Security Group Considerations in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
	// +kubebuilder:validation:Required
	VPCConfig []VPCConfigParameters `json:"vpcConfig" tf:"vpc_config,omitempty"`

	// –  Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ControlPlanePlacementObservation struct {
}

type ControlPlanePlacementParameters struct {

	// The name of the placement group for the Kubernetes control plane instances. This setting can't be changed after cluster creation.
	// +kubebuilder:validation:Required
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`
}

type EncryptionConfigObservation struct {
}

type EncryptionConfigParameters struct {

	// Configuration block with provider for encryption. Detailed below.
	// +kubebuilder:validation:Required
	Provider []ProviderParameters `json:"provider" tf:"provider,omitempty"`

	// List of strings with resources to be encrypted. Valid values: secrets.
	// +kubebuilder:validation:Required
	Resources []*string `json:"resources" tf:"resources,omitempty"`
}

type IdentityObservation struct {

	// Nested block containing OpenID Connect identity provider information for the cluster. Detailed below.
	Oidc []OidcObservation `json:"oidc,omitempty" tf:"oidc,omitempty"`
}

type IdentityParameters struct {
}

type KubernetesNetworkConfigObservation struct {

	// The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified ipv6 for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can't specify a custom IPv6 CIDR block when you create the cluster.
	ServiceIPv6Cidr *string `json:"serviceIpv6Cidr,omitempty" tf:"service_ipv6_cidr,omitempty"`
}

type KubernetesNetworkConfigParameters struct {

	// The IP family used to assign Kubernetes pod and service addresses. Valid values are ipv4 (default) and ipv6. You can only specify an IP family when you create a cluster, changing this value will force a new cluster to be created.
	// +kubebuilder:validation:Optional
	IPFamily *string `json:"ipFamily,omitempty" tf:"ip_family,omitempty"`

	// The CIDR block to assign Kubernetes pod and service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. You can only specify a custom CIDR block when you create a cluster, changing this value will force a new cluster to be created. The block must meet the following requirements:
	// +kubebuilder:validation:Optional
	ServiceIPv4Cidr *string `json:"serviceIpv4Cidr,omitempty" tf:"service_ipv4_cidr,omitempty"`
}

type OidcObservation struct {

	// Issuer URL for the OpenID Connect identity provider.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type OidcParameters struct {
}

type OutpostConfigObservation struct {
}

type OutpostConfigParameters struct {

	// The Amazon EC2 instance type that you want to use for your local Amazon EKS cluster on Outposts. The instance type that you specify is used for all Kubernetes control plane instances. The instance type can't be changed after cluster creation. Choose an instance type based on the number of nodes that your cluster will have. If your cluster will have:
	// +kubebuilder:validation:Required
	ControlPlaneInstanceType *string `json:"controlPlaneInstanceType" tf:"control_plane_instance_type,omitempty"`

	// An object representing the placement configuration for all the control plane instances of your local Amazon EKS cluster on AWS Outpost.
	// The following arguments are supported in the control_plane_placement configuration block:
	// +kubebuilder:validation:Optional
	ControlPlanePlacement []ControlPlanePlacementParameters `json:"controlPlanePlacement,omitempty" tf:"control_plane_placement,omitempty"`

	// The ARN of the Outpost that you want to use for your local Amazon EKS cluster on Outposts. This argument is a list of arns, but only a single Outpost ARN is supported currently.
	// +kubebuilder:validation:Required
	OutpostArns []*string `json:"outpostArns" tf:"outpost_arns,omitempty"`
}

type ProviderObservation struct {
}

type ProviderParameters struct {

	// ARN of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide.
	// +kubebuilder:validation:Required
	KeyArn *string `json:"keyArn" tf:"key_arn,omitempty"`
}

type VPCConfigObservation struct {

	// Cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control-plane-to-data-plane communication.
	ClusterSecurityGroupID *string `json:"clusterSecurityGroupId,omitempty" tf:"cluster_security_group_id,omitempty"`

	// ID of the VPC associated with your cluster.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigParameters struct {

	// Whether the Amazon EKS private API server endpoint is enabled. Default is false.
	// +kubebuilder:validation:Optional
	EndpointPrivateAccess *bool `json:"endpointPrivateAccess,omitempty" tf:"endpoint_private_access,omitempty"`

	// Whether the Amazon EKS public API server endpoint is enabled. Default is true.
	// +kubebuilder:validation:Optional
	EndpointPublicAccess *bool `json:"endpointPublicAccess,omitempty" tf:"endpoint_public_access,omitempty"`

	// List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with 0.0.0.0/0.
	// +kubebuilder:validation:Optional
	PublicAccessCidrs []*string `json:"publicAccessCidrs,omitempty" tf:"public_access_cidrs,omitempty"`

	// account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Manages an EKS Cluster
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
