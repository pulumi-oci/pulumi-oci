// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Deploy Stage resource in Oracle Cloud Infrastructure Devops service.
//
// Creates a new deployment stage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/devops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := devops.NewDeployStage(ctx, "testDeployStage", &devops.DeployStageArgs{
// 			DeployPipelineId: pulumi.Any(oci_devops_deploy_pipeline.Test_deploy_pipeline.Id),
// 			DeployStagePredecessorCollection: &devops.DeployStageDeployStagePredecessorCollectionArgs{
// 				Items: devops.DeployStageDeployStagePredecessorCollectionItemArray{
// 					&devops.DeployStageDeployStagePredecessorCollectionItemArgs{
// 						Id: pulumi.Any(_var.Deploy_stage_deploy_stage_predecessor_collection_items_id),
// 					},
// 				},
// 			},
// 			DeployStageType: pulumi.Any(_var.Deploy_stage_deploy_stage_type),
// 			ApprovalPolicy: &devops.DeployStageApprovalPolicyArgs{
// 				ApprovalPolicyType:        pulumi.Any(_var.Deploy_stage_approval_policy_approval_policy_type),
// 				NumberOfApprovalsRequired: pulumi.Any(_var.Deploy_stage_approval_policy_number_of_approvals_required),
// 			},
// 			BlueBackendIps: &devops.DeployStageBlueBackendIpsArgs{
// 				Items: pulumi.Any(_var.Deploy_stage_blue_backend_ips_items),
// 			},
// 			ComputeInstanceGroupDeployEnvironmentId: pulumi.Any(oci_devops_deploy_environment.Test_deploy_environment.Id),
// 			Config:                                  pulumi.Any(_var.Deploy_stage_config),
// 			DefinedTags: pulumi.AnyMap{
// 				"foo-namespace.bar-key": pulumi.Any("value"),
// 			},
// 			DeployArtifactId:               pulumi.Any(oci_devops_deploy_artifact.Test_deploy_artifact.Id),
// 			DeployArtifactIds:              pulumi.Any(_var.Deploy_stage_deploy_artifact_ids),
// 			DeploymentSpecDeployArtifactId: pulumi.Any(oci_devops_deploy_artifact.Test_deploy_artifact.Id),
// 			Description:                    pulumi.Any(_var.Deploy_stage_description),
// 			DisplayName:                    pulumi.Any(_var.Deploy_stage_display_name),
// 			DockerImageDeployArtifactId:    pulumi.Any(oci_devops_deploy_artifact.Test_deploy_artifact.Id),
// 			FailurePolicy: &devops.DeployStageFailurePolicyArgs{
// 				PolicyType:        pulumi.Any(_var.Deploy_stage_failure_policy_policy_type),
// 				FailureCount:      pulumi.Any(_var.Deploy_stage_failure_policy_failure_count),
// 				FailurePercentage: pulumi.Any(_var.Deploy_stage_failure_policy_failure_percentage),
// 			},
// 			FreeformTags: pulumi.AnyMap{
// 				"bar-key": pulumi.Any("value"),
// 			},
// 			FunctionDeployEnvironmentId: pulumi.Any(oci_devops_deploy_environment.Test_deploy_environment.Id),
// 			FunctionTimeoutInSeconds:    pulumi.Any(_var.Deploy_stage_function_timeout_in_seconds),
// 			GreenBackendIps: &devops.DeployStageGreenBackendIpsArgs{
// 				Items: pulumi.Any(_var.Deploy_stage_green_backend_ips_items),
// 			},
// 			IsAsync:                             pulumi.Any(_var.Deploy_stage_is_async),
// 			IsValidationEnabled:                 pulumi.Any(_var.Deploy_stage_is_validation_enabled),
// 			KubernetesManifestDeployArtifactIds: pulumi.Any(_var.Deploy_stage_kubernetes_manifest_deploy_artifact_ids),
// 			LoadBalancerConfig: &devops.DeployStageLoadBalancerConfigArgs{
// 				BackendPort:    pulumi.Any(_var.Deploy_stage_load_balancer_config_backend_port),
// 				ListenerName:   pulumi.Any(oci_load_balancer_listener.Test_listener.Name),
// 				LoadBalancerId: pulumi.Any(oci_load_balancer_load_balancer.Test_load_balancer.Id),
// 			},
// 			MaxMemoryInMbs:                pulumi.Any(_var.Deploy_stage_max_memory_in_mbs),
// 			Namespace:                     pulumi.Any(_var.Deploy_stage_namespace),
// 			OkeClusterDeployEnvironmentId: pulumi.Any(oci_devops_deploy_environment.Test_deploy_environment.Id),
// 			RollbackPolicy: &devops.DeployStageRollbackPolicyArgs{
// 				PolicyType: pulumi.Any(_var.Deploy_stage_rollback_policy_policy_type),
// 			},
// 			RolloutPolicy: &devops.DeployStageRolloutPolicyArgs{
// 				PolicyType:          pulumi.Any(_var.Deploy_stage_rollout_policy_policy_type),
// 				BatchCount:          pulumi.Any(_var.Deploy_stage_rollout_policy_batch_count),
// 				BatchDelayInSeconds: pulumi.Any(_var.Deploy_stage_rollout_policy_batch_delay_in_seconds),
// 				BatchPercentage:     pulumi.Any(_var.Deploy_stage_rollout_policy_batch_percentage),
// 				RampLimitPercent:    pulumi.Any(_var.Deploy_stage_rollout_policy_ramp_limit_percent),
// 			},
// 			TrafficShiftTarget: pulumi.Any(_var.Deploy_stage_traffic_shift_target),
// 			WaitCriteria: &devops.DeployStageWaitCriteriaArgs{
// 				WaitDuration: pulumi.Any(_var.Deploy_stage_wait_criteria_wait_duration),
// 				WaitType:     pulumi.Any(_var.Deploy_stage_wait_criteria_wait_type),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// DeployStages can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:devops/deployStage:DeployStage test_deploy_stage "id"
// ```
type DeployStage struct {
	pulumi.CustomResourceState

	// (Updatable) Specifies the approval policy.
	ApprovalPolicy DeployStageApprovalPolicyOutput `pulumi:"approvalPolicy"`
	// (Updatable) Collection of backend environment IP addresses.
	BlueBackendIps DeployStageBlueBackendIpsOutput `pulumi:"blueBackendIps"`
	// The OCID of a compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) A compute instance group environment OCID for rolling deployment.
	ComputeInstanceGroupDeployEnvironmentId pulumi.StringOutput `pulumi:"computeInstanceGroupDeployEnvironmentId"`
	// (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
	Config pulumi.MapOutput `pulumi:"config"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) Optional binary artifact OCID user may provide to this stage.
	DeployArtifactId pulumi.StringPtrOutput `pulumi:"deployArtifactId"`
	// (Updatable) Additional file artifact OCIDs.
	DeployArtifactIds pulumi.StringArrayOutput `pulumi:"deployArtifactIds"`
	// The OCID of a pipeline.
	DeployPipelineId pulumi.StringOutput `pulumi:"deployPipelineId"`
	// (Updatable) Collection containing the predecessors of a stage.
	DeployStagePredecessorCollection DeployStageDeployStagePredecessorCollectionOutput `pulumi:"deployStagePredecessorCollection"`
	// (Updatable) Deployment stage type.
	DeployStageType pulumi.StringOutput `pulumi:"deployStageType"`
	// (Updatable) The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId pulumi.StringOutput `pulumi:"deploymentSpecDeployArtifactId"`
	// (Updatable) Optional description about the deployment stage.
	Description pulumi.StringOutput `pulumi:"description"`
	// (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) A Docker image artifact OCID.
	DockerImageDeployArtifactId pulumi.StringOutput `pulumi:"dockerImageDeployArtifactId"`
	// (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	FailurePolicy DeployStageFailurePolicyOutput `pulumi:"failurePolicy"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// (Updatable) Function environment OCID.
	FunctionDeployEnvironmentId pulumi.StringOutput `pulumi:"functionDeployEnvironmentId"`
	// (Updatable) Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds pulumi.IntOutput `pulumi:"functionTimeoutInSeconds"`
	// (Updatable) Collection of backend environment IP addresses.
	GreenBackendIps DeployStageGreenBackendIpsOutput `pulumi:"greenBackendIps"`
	// (Updatable) A boolean flag specifies whether this stage executes asynchronously.
	IsAsync pulumi.BoolOutput `pulumi:"isAsync"`
	// (Updatable) A boolean flag specifies whether the invoked function should be validated.
	IsValidationEnabled pulumi.BoolOutput `pulumi:"isValidationEnabled"`
	// (Updatable) List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
	KubernetesManifestDeployArtifactIds pulumi.StringArrayOutput `pulumi:"kubernetesManifestDeployArtifactIds"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// (Updatable) Specifies config for load balancer traffic shift stages.
	LoadBalancerConfig DeployStageLoadBalancerConfigOutput `pulumi:"loadBalancerConfig"`
	// (Updatable) Maximum usable memory for the Function (in MB).
	MaxMemoryInMbs pulumi.StringOutput `pulumi:"maxMemoryInMbs"`
	// (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// (Updatable) Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId pulumi.StringOutput `pulumi:"okeClusterDeployEnvironmentId"`
	// The OCID of a project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	RollbackPolicy DeployStageRollbackPolicyOutput `pulumi:"rollbackPolicy"`
	// (Updatable) Description of rollout policy for load balancer traffic shift stage.
	RolloutPolicy DeployStageRolloutPolicyOutput `pulumi:"rolloutPolicy"`
	// The current state of the deployment stage.
	State pulumi.StringOutput `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapOutput `pulumi:"systemTags"`
	// Time the deployment stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Time the deployment stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
	// (Updatable) Specifies the target or destination backend set.
	TrafficShiftTarget pulumi.StringOutput `pulumi:"trafficShiftTarget"`
	// (Updatable) Specifies wait criteria for the Wait stage.
	WaitCriteria DeployStageWaitCriteriaOutput `pulumi:"waitCriteria"`
}

// NewDeployStage registers a new resource with the given unique name, arguments, and options.
func NewDeployStage(ctx *pulumi.Context,
	name string, args *DeployStageArgs, opts ...pulumi.ResourceOption) (*DeployStage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployPipelineId == nil {
		return nil, errors.New("invalid value for required argument 'DeployPipelineId'")
	}
	if args.DeployStagePredecessorCollection == nil {
		return nil, errors.New("invalid value for required argument 'DeployStagePredecessorCollection'")
	}
	if args.DeployStageType == nil {
		return nil, errors.New("invalid value for required argument 'DeployStageType'")
	}
	var resource DeployStage
	err := ctx.RegisterResource("oci:devops/deployStage:DeployStage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployStage gets an existing DeployStage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployStageState, opts ...pulumi.ResourceOption) (*DeployStage, error) {
	var resource DeployStage
	err := ctx.ReadResource("oci:devops/deployStage:DeployStage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployStage resources.
type deployStageState struct {
	// (Updatable) Specifies the approval policy.
	ApprovalPolicy *DeployStageApprovalPolicy `pulumi:"approvalPolicy"`
	// (Updatable) Collection of backend environment IP addresses.
	BlueBackendIps *DeployStageBlueBackendIps `pulumi:"blueBackendIps"`
	// The OCID of a compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) A compute instance group environment OCID for rolling deployment.
	ComputeInstanceGroupDeployEnvironmentId *string `pulumi:"computeInstanceGroupDeployEnvironmentId"`
	// (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
	Config map[string]interface{} `pulumi:"config"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Optional binary artifact OCID user may provide to this stage.
	DeployArtifactId *string `pulumi:"deployArtifactId"`
	// (Updatable) Additional file artifact OCIDs.
	DeployArtifactIds []string `pulumi:"deployArtifactIds"`
	// The OCID of a pipeline.
	DeployPipelineId *string `pulumi:"deployPipelineId"`
	// (Updatable) Collection containing the predecessors of a stage.
	DeployStagePredecessorCollection *DeployStageDeployStagePredecessorCollection `pulumi:"deployStagePredecessorCollection"`
	// (Updatable) Deployment stage type.
	DeployStageType *string `pulumi:"deployStageType"`
	// (Updatable) The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId *string `pulumi:"deploymentSpecDeployArtifactId"`
	// (Updatable) Optional description about the deployment stage.
	Description *string `pulumi:"description"`
	// (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) A Docker image artifact OCID.
	DockerImageDeployArtifactId *string `pulumi:"dockerImageDeployArtifactId"`
	// (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	FailurePolicy *DeployStageFailurePolicy `pulumi:"failurePolicy"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) Function environment OCID.
	FunctionDeployEnvironmentId *string `pulumi:"functionDeployEnvironmentId"`
	// (Updatable) Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds *int `pulumi:"functionTimeoutInSeconds"`
	// (Updatable) Collection of backend environment IP addresses.
	GreenBackendIps *DeployStageGreenBackendIps `pulumi:"greenBackendIps"`
	// (Updatable) A boolean flag specifies whether this stage executes asynchronously.
	IsAsync *bool `pulumi:"isAsync"`
	// (Updatable) A boolean flag specifies whether the invoked function should be validated.
	IsValidationEnabled *bool `pulumi:"isValidationEnabled"`
	// (Updatable) List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
	KubernetesManifestDeployArtifactIds []string `pulumi:"kubernetesManifestDeployArtifactIds"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// (Updatable) Specifies config for load balancer traffic shift stages.
	LoadBalancerConfig *DeployStageLoadBalancerConfig `pulumi:"loadBalancerConfig"`
	// (Updatable) Maximum usable memory for the Function (in MB).
	MaxMemoryInMbs *string `pulumi:"maxMemoryInMbs"`
	// (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace *string `pulumi:"namespace"`
	// (Updatable) Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId *string `pulumi:"okeClusterDeployEnvironmentId"`
	// The OCID of a project.
	ProjectId *string `pulumi:"projectId"`
	// (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	RollbackPolicy *DeployStageRollbackPolicy `pulumi:"rollbackPolicy"`
	// (Updatable) Description of rollout policy for load balancer traffic shift stage.
	RolloutPolicy *DeployStageRolloutPolicy `pulumi:"rolloutPolicy"`
	// The current state of the deployment stage.
	State *string `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// Time the deployment stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *string `pulumi:"timeCreated"`
	// Time the deployment stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *string `pulumi:"timeUpdated"`
	// (Updatable) Specifies the target or destination backend set.
	TrafficShiftTarget *string `pulumi:"trafficShiftTarget"`
	// (Updatable) Specifies wait criteria for the Wait stage.
	WaitCriteria *DeployStageWaitCriteria `pulumi:"waitCriteria"`
}

type DeployStageState struct {
	// (Updatable) Specifies the approval policy.
	ApprovalPolicy DeployStageApprovalPolicyPtrInput
	// (Updatable) Collection of backend environment IP addresses.
	BlueBackendIps DeployStageBlueBackendIpsPtrInput
	// The OCID of a compartment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) A compute instance group environment OCID for rolling deployment.
	ComputeInstanceGroupDeployEnvironmentId pulumi.StringPtrInput
	// (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
	Config pulumi.MapInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Optional binary artifact OCID user may provide to this stage.
	DeployArtifactId pulumi.StringPtrInput
	// (Updatable) Additional file artifact OCIDs.
	DeployArtifactIds pulumi.StringArrayInput
	// The OCID of a pipeline.
	DeployPipelineId pulumi.StringPtrInput
	// (Updatable) Collection containing the predecessors of a stage.
	DeployStagePredecessorCollection DeployStageDeployStagePredecessorCollectionPtrInput
	// (Updatable) Deployment stage type.
	DeployStageType pulumi.StringPtrInput
	// (Updatable) The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId pulumi.StringPtrInput
	// (Updatable) Optional description about the deployment stage.
	Description pulumi.StringPtrInput
	// (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) A Docker image artifact OCID.
	DockerImageDeployArtifactId pulumi.StringPtrInput
	// (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	FailurePolicy DeployStageFailurePolicyPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// (Updatable) Function environment OCID.
	FunctionDeployEnvironmentId pulumi.StringPtrInput
	// (Updatable) Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds pulumi.IntPtrInput
	// (Updatable) Collection of backend environment IP addresses.
	GreenBackendIps DeployStageGreenBackendIpsPtrInput
	// (Updatable) A boolean flag specifies whether this stage executes asynchronously.
	IsAsync pulumi.BoolPtrInput
	// (Updatable) A boolean flag specifies whether the invoked function should be validated.
	IsValidationEnabled pulumi.BoolPtrInput
	// (Updatable) List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
	KubernetesManifestDeployArtifactIds pulumi.StringArrayInput
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails pulumi.StringPtrInput
	// (Updatable) Specifies config for load balancer traffic shift stages.
	LoadBalancerConfig DeployStageLoadBalancerConfigPtrInput
	// (Updatable) Maximum usable memory for the Function (in MB).
	MaxMemoryInMbs pulumi.StringPtrInput
	// (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace pulumi.StringPtrInput
	// (Updatable) Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId pulumi.StringPtrInput
	// The OCID of a project.
	ProjectId pulumi.StringPtrInput
	// (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	RollbackPolicy DeployStageRollbackPolicyPtrInput
	// (Updatable) Description of rollout policy for load balancer traffic shift stage.
	RolloutPolicy DeployStageRolloutPolicyPtrInput
	// The current state of the deployment stage.
	State pulumi.StringPtrInput
	// Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapInput
	// Time the deployment stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated pulumi.StringPtrInput
	// Time the deployment stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated pulumi.StringPtrInput
	// (Updatable) Specifies the target or destination backend set.
	TrafficShiftTarget pulumi.StringPtrInput
	// (Updatable) Specifies wait criteria for the Wait stage.
	WaitCriteria DeployStageWaitCriteriaPtrInput
}

func (DeployStageState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployStageState)(nil)).Elem()
}

type deployStageArgs struct {
	// (Updatable) Specifies the approval policy.
	ApprovalPolicy *DeployStageApprovalPolicy `pulumi:"approvalPolicy"`
	// (Updatable) Collection of backend environment IP addresses.
	BlueBackendIps *DeployStageBlueBackendIps `pulumi:"blueBackendIps"`
	// (Updatable) A compute instance group environment OCID for rolling deployment.
	ComputeInstanceGroupDeployEnvironmentId *string `pulumi:"computeInstanceGroupDeployEnvironmentId"`
	// (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
	Config map[string]interface{} `pulumi:"config"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Optional binary artifact OCID user may provide to this stage.
	DeployArtifactId *string `pulumi:"deployArtifactId"`
	// (Updatable) Additional file artifact OCIDs.
	DeployArtifactIds []string `pulumi:"deployArtifactIds"`
	// The OCID of a pipeline.
	DeployPipelineId string `pulumi:"deployPipelineId"`
	// (Updatable) Collection containing the predecessors of a stage.
	DeployStagePredecessorCollection DeployStageDeployStagePredecessorCollection `pulumi:"deployStagePredecessorCollection"`
	// (Updatable) Deployment stage type.
	DeployStageType string `pulumi:"deployStageType"`
	// (Updatable) The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId *string `pulumi:"deploymentSpecDeployArtifactId"`
	// (Updatable) Optional description about the deployment stage.
	Description *string `pulumi:"description"`
	// (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) A Docker image artifact OCID.
	DockerImageDeployArtifactId *string `pulumi:"dockerImageDeployArtifactId"`
	// (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	FailurePolicy *DeployStageFailurePolicy `pulumi:"failurePolicy"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) Function environment OCID.
	FunctionDeployEnvironmentId *string `pulumi:"functionDeployEnvironmentId"`
	// (Updatable) Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds *int `pulumi:"functionTimeoutInSeconds"`
	// (Updatable) Collection of backend environment IP addresses.
	GreenBackendIps *DeployStageGreenBackendIps `pulumi:"greenBackendIps"`
	// (Updatable) A boolean flag specifies whether this stage executes asynchronously.
	IsAsync *bool `pulumi:"isAsync"`
	// (Updatable) A boolean flag specifies whether the invoked function should be validated.
	IsValidationEnabled *bool `pulumi:"isValidationEnabled"`
	// (Updatable) List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
	KubernetesManifestDeployArtifactIds []string `pulumi:"kubernetesManifestDeployArtifactIds"`
	// (Updatable) Specifies config for load balancer traffic shift stages.
	LoadBalancerConfig *DeployStageLoadBalancerConfig `pulumi:"loadBalancerConfig"`
	// (Updatable) Maximum usable memory for the Function (in MB).
	MaxMemoryInMbs *string `pulumi:"maxMemoryInMbs"`
	// (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace *string `pulumi:"namespace"`
	// (Updatable) Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId *string `pulumi:"okeClusterDeployEnvironmentId"`
	// (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	RollbackPolicy *DeployStageRollbackPolicy `pulumi:"rollbackPolicy"`
	// (Updatable) Description of rollout policy for load balancer traffic shift stage.
	RolloutPolicy *DeployStageRolloutPolicy `pulumi:"rolloutPolicy"`
	// (Updatable) Specifies the target or destination backend set.
	TrafficShiftTarget *string `pulumi:"trafficShiftTarget"`
	// (Updatable) Specifies wait criteria for the Wait stage.
	WaitCriteria *DeployStageWaitCriteria `pulumi:"waitCriteria"`
}

// The set of arguments for constructing a DeployStage resource.
type DeployStageArgs struct {
	// (Updatable) Specifies the approval policy.
	ApprovalPolicy DeployStageApprovalPolicyPtrInput
	// (Updatable) Collection of backend environment IP addresses.
	BlueBackendIps DeployStageBlueBackendIpsPtrInput
	// (Updatable) A compute instance group environment OCID for rolling deployment.
	ComputeInstanceGroupDeployEnvironmentId pulumi.StringPtrInput
	// (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
	Config pulumi.MapInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Optional binary artifact OCID user may provide to this stage.
	DeployArtifactId pulumi.StringPtrInput
	// (Updatable) Additional file artifact OCIDs.
	DeployArtifactIds pulumi.StringArrayInput
	// The OCID of a pipeline.
	DeployPipelineId pulumi.StringInput
	// (Updatable) Collection containing the predecessors of a stage.
	DeployStagePredecessorCollection DeployStageDeployStagePredecessorCollectionInput
	// (Updatable) Deployment stage type.
	DeployStageType pulumi.StringInput
	// (Updatable) The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId pulumi.StringPtrInput
	// (Updatable) Optional description about the deployment stage.
	Description pulumi.StringPtrInput
	// (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) A Docker image artifact OCID.
	DockerImageDeployArtifactId pulumi.StringPtrInput
	// (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	FailurePolicy DeployStageFailurePolicyPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// (Updatable) Function environment OCID.
	FunctionDeployEnvironmentId pulumi.StringPtrInput
	// (Updatable) Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds pulumi.IntPtrInput
	// (Updatable) Collection of backend environment IP addresses.
	GreenBackendIps DeployStageGreenBackendIpsPtrInput
	// (Updatable) A boolean flag specifies whether this stage executes asynchronously.
	IsAsync pulumi.BoolPtrInput
	// (Updatable) A boolean flag specifies whether the invoked function should be validated.
	IsValidationEnabled pulumi.BoolPtrInput
	// (Updatable) List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
	KubernetesManifestDeployArtifactIds pulumi.StringArrayInput
	// (Updatable) Specifies config for load balancer traffic shift stages.
	LoadBalancerConfig DeployStageLoadBalancerConfigPtrInput
	// (Updatable) Maximum usable memory for the Function (in MB).
	MaxMemoryInMbs pulumi.StringPtrInput
	// (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace pulumi.StringPtrInput
	// (Updatable) Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId pulumi.StringPtrInput
	// (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	RollbackPolicy DeployStageRollbackPolicyPtrInput
	// (Updatable) Description of rollout policy for load balancer traffic shift stage.
	RolloutPolicy DeployStageRolloutPolicyPtrInput
	// (Updatable) Specifies the target or destination backend set.
	TrafficShiftTarget pulumi.StringPtrInput
	// (Updatable) Specifies wait criteria for the Wait stage.
	WaitCriteria DeployStageWaitCriteriaPtrInput
}

func (DeployStageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployStageArgs)(nil)).Elem()
}

type DeployStageInput interface {
	pulumi.Input

	ToDeployStageOutput() DeployStageOutput
	ToDeployStageOutputWithContext(ctx context.Context) DeployStageOutput
}

func (*DeployStage) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployStage)(nil))
}

func (i *DeployStage) ToDeployStageOutput() DeployStageOutput {
	return i.ToDeployStageOutputWithContext(context.Background())
}

func (i *DeployStage) ToDeployStageOutputWithContext(ctx context.Context) DeployStageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployStageOutput)
}

func (i *DeployStage) ToDeployStagePtrOutput() DeployStagePtrOutput {
	return i.ToDeployStagePtrOutputWithContext(context.Background())
}

func (i *DeployStage) ToDeployStagePtrOutputWithContext(ctx context.Context) DeployStagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployStagePtrOutput)
}

type DeployStagePtrInput interface {
	pulumi.Input

	ToDeployStagePtrOutput() DeployStagePtrOutput
	ToDeployStagePtrOutputWithContext(ctx context.Context) DeployStagePtrOutput
}

type deployStagePtrType DeployStageArgs

func (*deployStagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployStage)(nil))
}

func (i *deployStagePtrType) ToDeployStagePtrOutput() DeployStagePtrOutput {
	return i.ToDeployStagePtrOutputWithContext(context.Background())
}

func (i *deployStagePtrType) ToDeployStagePtrOutputWithContext(ctx context.Context) DeployStagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployStagePtrOutput)
}

// DeployStageArrayInput is an input type that accepts DeployStageArray and DeployStageArrayOutput values.
// You can construct a concrete instance of `DeployStageArrayInput` via:
//
//          DeployStageArray{ DeployStageArgs{...} }
type DeployStageArrayInput interface {
	pulumi.Input

	ToDeployStageArrayOutput() DeployStageArrayOutput
	ToDeployStageArrayOutputWithContext(context.Context) DeployStageArrayOutput
}

type DeployStageArray []DeployStageInput

func (DeployStageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployStage)(nil)).Elem()
}

func (i DeployStageArray) ToDeployStageArrayOutput() DeployStageArrayOutput {
	return i.ToDeployStageArrayOutputWithContext(context.Background())
}

func (i DeployStageArray) ToDeployStageArrayOutputWithContext(ctx context.Context) DeployStageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployStageArrayOutput)
}

// DeployStageMapInput is an input type that accepts DeployStageMap and DeployStageMapOutput values.
// You can construct a concrete instance of `DeployStageMapInput` via:
//
//          DeployStageMap{ "key": DeployStageArgs{...} }
type DeployStageMapInput interface {
	pulumi.Input

	ToDeployStageMapOutput() DeployStageMapOutput
	ToDeployStageMapOutputWithContext(context.Context) DeployStageMapOutput
}

type DeployStageMap map[string]DeployStageInput

func (DeployStageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployStage)(nil)).Elem()
}

func (i DeployStageMap) ToDeployStageMapOutput() DeployStageMapOutput {
	return i.ToDeployStageMapOutputWithContext(context.Background())
}

func (i DeployStageMap) ToDeployStageMapOutputWithContext(ctx context.Context) DeployStageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployStageMapOutput)
}

type DeployStageOutput struct {
	*pulumi.OutputState
}

func (DeployStageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployStage)(nil))
}

func (o DeployStageOutput) ToDeployStageOutput() DeployStageOutput {
	return o
}

func (o DeployStageOutput) ToDeployStageOutputWithContext(ctx context.Context) DeployStageOutput {
	return o
}

func (o DeployStageOutput) ToDeployStagePtrOutput() DeployStagePtrOutput {
	return o.ToDeployStagePtrOutputWithContext(context.Background())
}

func (o DeployStageOutput) ToDeployStagePtrOutputWithContext(ctx context.Context) DeployStagePtrOutput {
	return o.ApplyT(func(v DeployStage) *DeployStage {
		return &v
	}).(DeployStagePtrOutput)
}

type DeployStagePtrOutput struct {
	*pulumi.OutputState
}

func (DeployStagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployStage)(nil))
}

func (o DeployStagePtrOutput) ToDeployStagePtrOutput() DeployStagePtrOutput {
	return o
}

func (o DeployStagePtrOutput) ToDeployStagePtrOutputWithContext(ctx context.Context) DeployStagePtrOutput {
	return o
}

type DeployStageArrayOutput struct{ *pulumi.OutputState }

func (DeployStageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeployStage)(nil))
}

func (o DeployStageArrayOutput) ToDeployStageArrayOutput() DeployStageArrayOutput {
	return o
}

func (o DeployStageArrayOutput) ToDeployStageArrayOutputWithContext(ctx context.Context) DeployStageArrayOutput {
	return o
}

func (o DeployStageArrayOutput) Index(i pulumi.IntInput) DeployStageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeployStage {
		return vs[0].([]DeployStage)[vs[1].(int)]
	}).(DeployStageOutput)
}

type DeployStageMapOutput struct{ *pulumi.OutputState }

func (DeployStageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeployStage)(nil))
}

func (o DeployStageMapOutput) ToDeployStageMapOutput() DeployStageMapOutput {
	return o
}

func (o DeployStageMapOutput) ToDeployStageMapOutputWithContext(ctx context.Context) DeployStageMapOutput {
	return o
}

func (o DeployStageMapOutput) MapIndex(k pulumi.StringInput) DeployStageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeployStage {
		return vs[0].(map[string]DeployStage)[vs[1].(string)]
	}).(DeployStageOutput)
}

func init() {
	pulumi.RegisterOutputType(DeployStageOutput{})
	pulumi.RegisterOutputType(DeployStagePtrOutput{})
	pulumi.RegisterOutputType(DeployStageArrayOutput{})
	pulumi.RegisterOutputType(DeployStageMapOutput{})
}
