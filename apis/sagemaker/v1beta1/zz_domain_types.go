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

type CanvasAppSettingsObservation struct {
}

type CanvasAppSettingsParameters struct {

	// Time series forecast settings for the Canvas app. see Time Series Forecasting Settings below.
	// +kubebuilder:validation:Optional
	TimeSeriesForecastingSettings []TimeSeriesForecastingSettingsParameters `json:"timeSeriesForecastingSettings,omitempty" tf:"time_series_forecasting_settings,omitempty"`
}

type CustomImageObservation struct {
}

type CustomImageParameters struct {

	// The name of the App Image Config.
	// +kubebuilder:validation:Required
	AppImageConfigName *string `json:"appImageConfigName" tf:"app_image_config_name,omitempty"`

	// The name of the Custom Image.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.ImageVersion
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("image_name",false)
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// Reference to a ImageVersion in sagemaker to populate imageName.
	// +kubebuilder:validation:Optional
	ImageNameRef *v1.Reference `json:"imageNameRef,omitempty" tf:"-"`

	// Selector for a ImageVersion in sagemaker to populate imageName.
	// +kubebuilder:validation:Optional
	ImageNameSelector *v1.Selector `json:"imageNameSelector,omitempty" tf:"-"`

	// The version number of the Custom Image.
	// +kubebuilder:validation:Optional
	ImageVersionNumber *float64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type DefaultResourceSpecObservation struct {
}

type DefaultResourceSpecParameters struct {

	// The instance type that the image version runs on.. For valid values see SageMaker Instance Types.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type DefaultSpaceSettingsObservation struct {
}

type DefaultSpaceSettingsParameters struct {

	// The execution role for the space.
	// +kubebuilder:validation:Required
	ExecutionRole *string `json:"executionRole" tf:"execution_role,omitempty"`

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	// +kubebuilder:validation:Optional
	JupyterServerAppSettings []JupyterServerAppSettingsParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	// +kubebuilder:validation:Optional
	KernelGatewayAppSettings []KernelGatewayAppSettingsParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings,omitempty"`

	// The security groups for the Amazon Virtual Private Cloud that the space uses for communication.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
}

type DefaultUserSettingsJupyterServerAppSettingsCodeRepositoryObservation struct {
}

type DefaultUserSettingsJupyterServerAppSettingsCodeRepositoryParameters struct {

	// The URL of the Git repository.
	// +kubebuilder:validation:Required
	RepositoryURL *string `json:"repositoryUrl" tf:"repository_url,omitempty"`
}

type DefaultUserSettingsJupyterServerAppSettingsObservation struct {
}

type DefaultUserSettingsJupyterServerAppSettingsParameters struct {

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
	// +kubebuilder:validation:Optional
	CodeRepository []DefaultUserSettingsJupyterServerAppSettingsCodeRepositoryParameters `json:"codeRepository,omitempty" tf:"code_repository,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []JupyterServerAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	// +kubebuilder:validation:Optional
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecObservation struct {
}

type DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters struct {

	// The instance type that the image version runs on.. For valid values see SageMaker Instance Types.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type DefaultUserSettingsKernelGatewayAppSettingsObservation struct {
}

type DefaultUserSettingsKernelGatewayAppSettingsParameters struct {

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	// +kubebuilder:validation:Optional
	CustomImage []KernelGatewayAppSettingsCustomImageParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []DefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	// +kubebuilder:validation:Optional
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type DefaultUserSettingsObservation struct {
}

type DefaultUserSettingsParameters struct {

	// The Canvas app settings. See Canvas App Settings below.
	// +kubebuilder:validation:Optional
	CanvasAppSettings []CanvasAppSettingsParameters `json:"canvasAppSettings,omitempty" tf:"canvas_app_settings,omitempty"`

	// The execution role for the space.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ExecutionRole *string `json:"executionRole,omitempty" tf:"execution_role,omitempty"`

	// Reference to a Role in iam to populate executionRole.
	// +kubebuilder:validation:Optional
	ExecutionRoleRef *v1.Reference `json:"executionRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate executionRole.
	// +kubebuilder:validation:Optional
	ExecutionRoleSelector *v1.Selector `json:"executionRoleSelector,omitempty" tf:"-"`

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	// +kubebuilder:validation:Optional
	JupyterServerAppSettings []DefaultUserSettingsJupyterServerAppSettingsParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	// +kubebuilder:validation:Optional
	KernelGatewayAppSettings []DefaultUserSettingsKernelGatewayAppSettingsParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings,omitempty"`

	// The RSession app settings. See RSession App Settings below.
	// +kubebuilder:validation:Optional
	RSessionAppSettings []RSessionAppSettingsParameters `json:"rSessionAppSettings,omitempty" tf:"r_session_app_settings,omitempty"`

	// The security groups for the Amazon Virtual Private Cloud that the space uses for communication.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The sharing settings. See Sharing Settings below.
	// +kubebuilder:validation:Optional
	SharingSettings []SharingSettingsParameters `json:"sharingSettings,omitempty" tf:"sharing_settings,omitempty"`

	// The TensorBoard app settings. See TensorBoard App Settings below.
	// +kubebuilder:validation:Optional
	TensorBoardAppSettings []TensorBoardAppSettingsParameters `json:"tensorBoardAppSettings,omitempty" tf:"tensor_board_app_settings,omitempty"`
}

type DomainObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEFSFileSystemID *string `json:"homeEfsFileSystemId,omitempty" tf:"home_efs_file_system_id,omitempty"`

	// The ID of the Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
	SecurityGroupIDForDomainBoundary *string `json:"securityGroupIdForDomainBoundary,omitempty" tf:"security_group_id_for_domain_boundary,omitempty"`

	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceID *string `json:"singleSignOnManagedApplicationInstanceId,omitempty" tf:"single_sign_on_managed_application_instance_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The domain's URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DomainParameters struct {

	// Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly. Valid values are PublicInternetOnly and VpcOnly.
	// +kubebuilder:validation:Optional
	AppNetworkAccessType *string `json:"appNetworkAccessType,omitempty" tf:"app_network_access_type,omitempty"`

	// The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Valid values are Service and Customer.
	// +kubebuilder:validation:Optional
	AppSecurityGroupManagement *string `json:"appSecurityGroupManagement,omitempty" tf:"app_security_group_management,omitempty"`

	// The mode of authentication that members use to access the domain. Valid values are IAM and SSO.
	// +kubebuilder:validation:Required
	AuthMode *string `json:"authMode" tf:"auth_mode,omitempty"`

	// The default space settings. See Default Space Settings below.
	// +kubebuilder:validation:Optional
	DefaultSpaceSettings []DefaultSpaceSettingsParameters `json:"defaultSpaceSettings,omitempty" tf:"default_space_settings,omitempty"`

	// The default user settings. See Default User Settings below.
	// +kubebuilder:validation:Required
	DefaultUserSettings []DefaultUserSettingsParameters `json:"defaultUserSettings" tf:"default_user_settings,omitempty"`

	// The domain name.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// The domain settings. See Domain Settings below.
	// +kubebuilder:validation:Optional
	DomainSettings []DomainSettingsParameters `json:"domainSettings,omitempty" tf:"domain_settings,omitempty"`

	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The VPC subnets that Studio uses for communication.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type DomainSettingsObservation struct {
}

type DomainSettingsParameters struct {

	// The configuration for attaching a SageMaker user profile name to the execution role as a sts:SourceIdentity key AWS Docs. Valid values are USER_PROFILE_NAME and DISABLED.
	// +kubebuilder:validation:Optional
	ExecutionRoleIdentityConfig *string `json:"executionRoleIdentityConfig,omitempty" tf:"execution_role_identity_config,omitempty"`

	// The security groups for the Amazon Virtual Private Cloud that the Domain uses for communication between Domain-level apps and user apps.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type JupyterServerAppSettingsCodeRepositoryObservation struct {
}

type JupyterServerAppSettingsCodeRepositoryParameters struct {

	// The URL of the Git repository.
	// +kubebuilder:validation:Required
	RepositoryURL *string `json:"repositoryUrl" tf:"repository_url,omitempty"`
}

type JupyterServerAppSettingsDefaultResourceSpecObservation struct {
}

type JupyterServerAppSettingsDefaultResourceSpecParameters struct {

	// The instance type that the image version runs on.. For valid values see SageMaker Instance Types.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type JupyterServerAppSettingsObservation struct {
}

type JupyterServerAppSettingsParameters struct {

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
	// +kubebuilder:validation:Optional
	CodeRepository []JupyterServerAppSettingsCodeRepositoryParameters `json:"codeRepository,omitempty" tf:"code_repository,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []DefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	// +kubebuilder:validation:Optional
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type KernelGatewayAppSettingsCustomImageObservation struct {
}

type KernelGatewayAppSettingsCustomImageParameters struct {

	// The name of the App Image Config.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.AppImageConfig
	// +kubebuilder:validation:Optional
	AppImageConfigName *string `json:"appImageConfigName,omitempty" tf:"app_image_config_name,omitempty"`

	// Reference to a AppImageConfig in sagemaker to populate appImageConfigName.
	// +kubebuilder:validation:Optional
	AppImageConfigNameRef *v1.Reference `json:"appImageConfigNameRef,omitempty" tf:"-"`

	// Selector for a AppImageConfig in sagemaker to populate appImageConfigName.
	// +kubebuilder:validation:Optional
	AppImageConfigNameSelector *v1.Selector `json:"appImageConfigNameSelector,omitempty" tf:"-"`

	// The name of the Custom Image.
	// +kubebuilder:validation:Required
	ImageName *string `json:"imageName" tf:"image_name,omitempty"`

	// The version number of the Custom Image.
	// +kubebuilder:validation:Optional
	ImageVersionNumber *float64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type KernelGatewayAppSettingsDefaultResourceSpecObservation struct {
}

type KernelGatewayAppSettingsDefaultResourceSpecParameters struct {

	// The instance type that the image version runs on.. For valid values see SageMaker Instance Types.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type KernelGatewayAppSettingsObservation struct {
}

type KernelGatewayAppSettingsParameters struct {

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	// +kubebuilder:validation:Optional
	CustomImage []CustomImageParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []KernelGatewayAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	// +kubebuilder:validation:Optional
	LifecycleConfigArns []*string `json:"lifecycleConfigArns,omitempty" tf:"lifecycle_config_arns,omitempty"`
}

type RSessionAppSettingsCustomImageObservation struct {
}

type RSessionAppSettingsCustomImageParameters struct {

	// The name of the App Image Config.
	// +kubebuilder:validation:Required
	AppImageConfigName *string `json:"appImageConfigName" tf:"app_image_config_name,omitempty"`

	// The name of the Custom Image.
	// +kubebuilder:validation:Required
	ImageName *string `json:"imageName" tf:"image_name,omitempty"`

	// The version number of the Custom Image.
	// +kubebuilder:validation:Optional
	ImageVersionNumber *float64 `json:"imageVersionNumber,omitempty" tf:"image_version_number,omitempty"`
}

type RSessionAppSettingsDefaultResourceSpecObservation struct {
}

type RSessionAppSettingsDefaultResourceSpecParameters struct {

	// The instance type that the image version runs on.. For valid values see SageMaker Instance Types.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type RSessionAppSettingsObservation struct {
}

type RSessionAppSettingsParameters struct {

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	// +kubebuilder:validation:Optional
	CustomImage []RSessionAppSettingsCustomImageParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []RSessionAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// The retention policy for data stored on an Amazon Elastic File System (EFS) volume. Valid values are Retain or Delete.  Default value is Retain.
	// +kubebuilder:validation:Optional
	HomeEFSFileSystem *string `json:"homeEfsFileSystem,omitempty" tf:"home_efs_file_system,omitempty"`
}

type SharingSettingsObservation struct {
}

type SharingSettingsParameters struct {

	// Whether to include the notebook cell output when sharing the notebook. The default is Disabled. Valid values are Allowed and Disabled.
	// +kubebuilder:validation:Optional
	NotebookOutputOption *string `json:"notebookOutputOption,omitempty" tf:"notebook_output_option,omitempty"`

	// When notebook_output_option is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.
	// +kubebuilder:validation:Optional
	S3KMSKeyID *string `json:"s3KmsKeyId,omitempty" tf:"s3_kms_key_id,omitempty"`

	// When notebook_output_option is Allowed, the Amazon S3 bucket used to save the notebook cell output.
	// +kubebuilder:validation:Optional
	S3OutputPath *string `json:"s3OutputPath,omitempty" tf:"s3_output_path,omitempty"`
}

type TensorBoardAppSettingsDefaultResourceSpecObservation struct {
}

type TensorBoardAppSettingsDefaultResourceSpecParameters struct {

	// The instance type that the image version runs on.. For valid values see SageMaker Instance Types.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	// +kubebuilder:validation:Optional
	LifecycleConfigArn *string `json:"lifecycleConfigArn,omitempty" tf:"lifecycle_config_arn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	// +kubebuilder:validation:Optional
	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn,omitempty"`

	// The ARN of the image version created on the instance.
	// +kubebuilder:validation:Optional
	SagemakerImageVersionArn *string `json:"sagemakerImageVersionArn,omitempty" tf:"sagemaker_image_version_arn,omitempty"`
}

type TensorBoardAppSettingsObservation struct {
}

type TensorBoardAppSettingsParameters struct {

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	// +kubebuilder:validation:Optional
	DefaultResourceSpec []TensorBoardAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec,omitempty" tf:"default_resource_spec,omitempty"`
}

type TimeSeriesForecastingSettingsObservation struct {
}

type TimeSeriesForecastingSettingsParameters struct {

	// The IAM role that Canvas passes to Amazon Forecast for time series forecasting. By default, Canvas uses the execution role specified in the UserProfile that launches the Canvas app. If an execution role is not specified in the UserProfile, Canvas uses the execution role specified in the Domain that owns the UserProfile. To allow time series forecasting, this IAM role should have the AmazonSageMakerCanvasForecastAccess policy attached and forecast.amazonaws.com added in the trust relationship as a service principal.
	// +kubebuilder:validation:Optional
	AmazonForecastRoleArn *string `json:"amazonForecastRoleArn,omitempty" tf:"amazon_forecast_role_arn,omitempty"`

	// Describes whether time series forecasting is enabled or disabled in the Canvas app. Valid values are ENABLED and DISABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API. Provides a SageMaker Domain resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
